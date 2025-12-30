package cli

import (
	"context"
	"fmt"

	"github.com/spf13/cobra"
	"jonopens/sitemapper/internal/repositories"
)

var reportCmd = &cobra.Command{
	Use:   "report",
	Short: "Manage reports",
	Long:  `List, view, and manage sitemap reports.`,
}

var reportListCmd = &cobra.Command{
	Use:   "list",
	Short: "List all reports",
	Long:  `List all reports for a user.`,
	RunE:  runReportList,
}

var reportGetCmd = &cobra.Command{
	Use:   "get <report-id>",
	Short: "Get a specific report",
	Long:  `Display details for a specific report.`,
	Args:  cobra.ExactArgs(1),
	RunE:  runReportGet,
}

var (
	reportListUserID string
	reportListLimit  int
)

func init() {
	// List command flags
	reportListCmd.Flags().StringVar(&reportListUserID, "user-id", "", "filter by user ID (defaults to config default_user_id)")
	reportListCmd.Flags().IntVar(&reportListLimit, "limit", 50, "maximum number of reports to list")
	
	// Add subcommands
	reportCmd.AddCommand(reportListCmd)
	reportCmd.AddCommand(reportGetCmd)
}

func runReportList(cmd *cobra.Command, args []string) error {
	ctx := GetContext()
	
	// Use default user ID from config if not provided
	if reportListUserID == "" {
		reportListUserID = ctx.Config.DefaultUserID
	}
	
	ctx.Formatter.Info(fmt.Sprintf("Listing reports for user: %s", reportListUserID))
	
	// Get reports from database
	reportRepo := ctx.DB.Reports()
	reports, err := reportRepo.List(context.Background(), repositories.ReportFilters{
		UserID: reportListUserID,
		Limit:  reportListLimit,
	})
	if err != nil {
		ctx.Formatter.Error(fmt.Sprintf("Failed to list reports: %v", err))
		return err
	}
	
	if len(reports) == 0 {
		ctx.Formatter.Info("No reports found")
		return nil
	}
	
	// Output results
	if ctx.Config.OutputFormat == "json" {
		return ctx.Formatter.Print(reports)
	}
	
	// Print as table
	fmt.Printf("\nFound %d report(s):\n\n", len(reports))
	
	rows := [][]string{
		{"ID", "User ID", "URLs", "Valid", "Invalid", "Created"},
	}
	
	for _, report := range reports {
		rows = append(rows, []string{
			truncate(report.ID, 20),
			truncate(report.UserID, 15),
			fmt.Sprintf("%d", report.EntryCount),
			fmt.Sprintf("%d", report.ValidEntryCount),
			fmt.Sprintf("%d", report.InvalidEntryCount),
			report.CreatedAt.Format("2006-01-02 15:04"),
		})
	}
	
	ctx.Formatter.Print(rows)
	
	fmt.Printf("\nTotal: %d report(s)\n", len(reports))
	
	return nil
}

func runReportGet(cmd *cobra.Command, args []string) error {
	ctx := GetContext()
	reportID := args[0]
	
	ctx.Formatter.Info(fmt.Sprintf("Fetching report: %s", reportID))
	
	// Get report from database
	reportRepo := ctx.DB.Reports()
	report, err := reportRepo.GetByID(context.Background(), reportID)
	if err != nil {
		ctx.Formatter.Error(fmt.Sprintf("Failed to get report: %v", err))
		return err
	}
	
	if report == nil {
		ctx.Formatter.Error("Report not found")
		return fmt.Errorf("report not found: %s", reportID)
	}
	
	// Output results
	if ctx.Config.OutputFormat == "json" {
		return ctx.Formatter.Print(report)
	}
	
	// Print report details
	fmt.Printf("\nReport Details:\n")
	fmt.Printf("  ID:                %s\n", report.ID)
	fmt.Printf("  User ID:           %s\n", report.UserID)
	fmt.Printf("\n")
	fmt.Printf("Entry Counts:\n")
	fmt.Printf("  Total Entries:     %d\n", report.EntryCount)
	fmt.Printf("  Stored Entries:    %d\n", report.StoredEntryCount)
	fmt.Printf("  Valid Entries:     %d\n", report.ValidEntryCount)
	fmt.Printf("  Invalid Entries:   %d\n", report.InvalidEntryCount)
	
	if report.LiveEntryCount > 0 || report.DownEntryCount > 0 {
		fmt.Printf("  Live Entries:      %d\n", report.LiveEntryCount)
		fmt.Printf("  Down Entries:      %d\n", report.DownEntryCount)
	}
	
	fmt.Printf("\n")
	fmt.Printf("Grouping:\n")
	fmt.Printf("  Grouping Count:    %d\n", report.GroupingCount)
	fmt.Printf("  Ungrouped Count:   %d\n", report.UngroupedCount)
	
	if report.ChildSitemapCount > 0 {
		fmt.Printf("\n")
		fmt.Printf("Structure:\n")
		fmt.Printf("  Child Sitemaps:    %d\n", report.ChildSitemapCount)
	}
	
	fmt.Printf("\n")
	fmt.Printf("Sampling:\n")
	fmt.Printf("  Fully Stored:      %t\n", report.IsFullyStored)
	fmt.Printf("  Sampling Strategy: %s\n", report.SamplingStrategy)
	if report.SamplingRate != nil {
		fmt.Printf("  Sampling Rate:     %.2f%%\n", *report.SamplingRate*100)
	}
	
	fmt.Printf("\n")
	fmt.Printf("Timestamps:\n")
	fmt.Printf("  Created:           %s\n", report.CreatedAt.Format("2006-01-02 15:04:05"))
	fmt.Printf("  Updated:           %s\n", report.UpdatedAt.Format("2006-01-02 15:04:05"))
	
	// Get and display sample entries
	ctx.Formatter.Info("\nFetching sample entries...")
	entryRepo := ctx.DB.Entries()
	entries, err := entryRepo.List(context.Background(), repositories.EntryFilters{
		ReportID: reportID,
		Limit:    10,
	})
	if err == nil && len(entries) > 0 {
		fmt.Printf("\nSample Entries (first %d):\n\n", len(entries))
		
		entryRows := [][]string{
			{"URL", "Valid", "Type"},
		}
		
		for _, entry := range entries {
			validStr := "✓"
			if !entry.IsValid {
				validStr = "✗"
			}
			entryRows = append(entryRows, []string{
				truncate(entry.URL, 70),
				validStr,
				string(entry.Type),
			})
		}
		
		ctx.Formatter.Print(entryRows)
	}
	
	fmt.Println()
	
	return nil
}

