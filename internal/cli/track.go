package cli

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/spf13/cobra"
	"jonopens/sitemapper/internal/models"
	"jonopens/sitemapper/pkg/sitemap"
)

var (
	trackName   string
	trackUserID string
)

var trackCmd = &cobra.Command{
	Use:   "track <url|file>",
	Short: "Track a sitemap by saving a snapshot to the database",
	Long: `Parse a sitemap and save a snapshot to the database for historical tracking.
This allows you to compare sitemaps over time using the compare command.
Example:
  sitemapper track https://example.com/sitemap.xml --name "example-v1"
  sitemapper track ./sitemap.xml --name "local-snapshot"`,
	Args: cobra.ExactArgs(1),
	RunE: runTrack,
}

func init() {
	trackCmd.Flags().StringVar(&trackName, "name", "", "name for this snapshot (optional)")
	trackCmd.Flags().StringVar(&trackUserID, "user-id", "", "user ID (defaults to config default_user_id)")
}

func runTrack(cmd *cobra.Command, args []string) error {
	ctx := GetContext()
	source := args[0]
	
	// Use default user ID from config if not provided
	if trackUserID == "" {
		trackUserID = ctx.Config.DefaultUserID
	}
	
	ctx.Formatter.Info(fmt.Sprintf("Tracking sitemap from: %s", source))
	
	// Read and parse sitemap
	data, err := readSitemapSource(source)
	if err != nil {
		ctx.Formatter.Error(fmt.Sprintf("Failed to read sitemap: %v", err))
		return err
	}
	
	parser := sitemap.NewParser()
	sitemapType, err := parser.DetectType(data)
	if err != nil {
		ctx.Formatter.Error(fmt.Sprintf("Failed to detect sitemap type: %v", err))
		return err
	}
	
	if sitemapType != "sitemap" {
		return fmt.Errorf("currently only regular sitemaps are supported for tracking, got: %s", sitemapType)
	}
	
	sm, err := parser.Parse(data)
	if err != nil {
		ctx.Formatter.Error(fmt.Sprintf("Failed to parse sitemap: %v", err))
		return err
	}
	
	// Validate sitemap
	validator := sitemap.NewValidator()
	if err := validator.Validate(sm); err != nil {
		ctx.Formatter.Warning(fmt.Sprintf("Sitemap validation warning: %v", err))
	}
	
	ctx.Formatter.Info(fmt.Sprintf("Parsed %d URLs", len(sm.URLs)))
	
	// Save to database
	reportID, err := saveSitemapSnapshot(ctx, sm, source, trackUserID, trackName)
	if err != nil {
		ctx.Formatter.Error(fmt.Sprintf("Failed to save snapshot: %v", err))
		return err
	}
	
	ctx.Formatter.Success(fmt.Sprintf("Snapshot saved with ID: %s", reportID))
	
	// Output report details
	if ctx.Config.OutputFormat == "json" {
		result := map[string]interface{}{
			"report_id":  reportID,
			"source":     source,
			"name":       trackName,
			"user_id":    trackUserID,
			"url_count":  len(sm.URLs),
			"created_at": time.Now().Format(time.RFC3339),
		}
		return ctx.Formatter.Print(result)
	}
	
	fmt.Printf("\nSnapshot Details:\n")
	fmt.Printf("  Report ID: %s\n", reportID)
	fmt.Printf("  Source:    %s\n", source)
	if trackName != "" {
		fmt.Printf("  Name:      %s\n", trackName)
	}
	fmt.Printf("  User ID:   %s\n", trackUserID)
	fmt.Printf("  URLs:      %d\n", len(sm.URLs))
	fmt.Printf("  Created:   %s\n", time.Now().Format("2006-01-02 15:04:05"))
	fmt.Println()
	
	ctx.Formatter.Info(fmt.Sprintf("Use 'sitemapper report get %s' to view this report", reportID))
	
	return nil
}

func saveSitemapSnapshot(ctx *CLIContext, sm *sitemap.Sitemap, source, userID, name string) (string, error) {
	contextBg := context.Background()
	
	// Generate report ID
	reportID := uuid.New().String()
	
	// Count valid entries
	validCount := 0
	for _, u := range sm.URLs {
		validator := sitemap.NewValidator()
		if err := validator.ValidateURL(&u); err == nil {
			validCount++
		}
	}
	
	// Create report
	report := &models.Report{
		ID:                reportID,
		UserID:            userID,
		EntryCount:        len(sm.URLs),
		StoredEntryCount:  len(sm.URLs),
		ValidEntryCount:   validCount,
		InvalidEntryCount: len(sm.URLs) - validCount,
		LiveEntryCount:    0, // Not checking liveness yet
		DownEntryCount:    0,
		GroupingCount:     0,
		UngroupedCount:    len(sm.URLs),
		ChildSitemapCount: 0,
		IsFullyStored:     true,
		SamplingStrategy:  models.SamplingStrategyNone,
		SamplingRate:      nil,
		CreatedAt:         time.Now(),
		UpdatedAt:         time.Now(),
	}
	
	// Start transaction
	tx, err := ctx.DB.BeginTx(contextBg)
	if err != nil {
		return "", fmt.Errorf("failed to begin transaction: %w", err)
	}
	defer tx.Rollback()
	
	// Save report
	if err := tx.Reports().Create(contextBg, report); err != nil {
		return "", fmt.Errorf("failed to create report: %w", err)
	}
	
	// Save entries
	for _, u := range sm.URLs {
		entry := &models.Entry{
			ID:              uuid.New().String(),
			ReportID:        reportID,
			GroupingID:      nil,
			Type:            models.EntryTypeURL,
			URL:             u.Loc,
			LastModified:    parseLastMod(u.LastMod),
			ChangeFreq:      stringPtr(u.ChangeFreq),
			Priority:        float64Ptr(u.Priority),
			IsValid:         true,
			ValidationError: nil,
			SelectionReason: models.SelectionReasonFullStorage,
			CreatedAt:       time.Now(),
			UpdatedAt:       time.Now(),
		}
		
		// Validate entry
		validator := sitemap.NewValidator()
		if err := validator.ValidateURL(&u); err != nil {
			entry.IsValid = false
			errMsg := err.Error()
			entry.ValidationError = &errMsg
		}
		
		if err := tx.Entries().Create(contextBg, entry); err != nil {
			return "", fmt.Errorf("failed to create entry: %w", err)
		}
	}
	
	// Commit transaction
	if err := tx.Commit(); err != nil {
		return "", fmt.Errorf("failed to commit transaction: %w", err)
	}
	
	return reportID, nil
}

func parseLastMod(lastMod string) *time.Time {
	if lastMod == "" {
		return nil
	}
	
	// Try common date formats
	formats := []string{
		"2006-01-02",
		"2006-01-02T15:04:05Z",
		"2006-01-02T15:04:05-07:00",
		time.RFC3339,
	}
	
	for _, format := range formats {
		if t, err := time.Parse(format, lastMod); err == nil {
			return &t
		}
	}
	
	return nil
}

func stringPtr(s string) *string {
	if s == "" {
		return nil
	}
	return &s
}

func float64Ptr(f float64) *float64 {
	if f == 0 {
		return nil
	}
	return &f
}

