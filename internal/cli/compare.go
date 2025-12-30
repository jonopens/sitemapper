package cli

import (
	"context"
	"fmt"

	"github.com/spf13/cobra"
	"jonopens/sitemapper/internal/repositories"
	"jonopens/sitemapper/internal/services"
	"jonopens/sitemapper/pkg/sitemap"
)

var (
	compareShowUnchanged bool
)

var compareCmd = &cobra.Command{
	Use:   "compare <source1> <source2>",
	Short: "Compare two sitemaps and show differences",
	Long: `Compare two sitemaps and display added, removed, and unchanged URLs.
Sources can be URLs, file paths, or report IDs from tracked snapshots.
Examples:
  sitemapper compare file1.xml file2.xml
  sitemapper compare https://example.com/sitemap.xml file.xml
  sitemapper compare report-id-1 report-id-2`,
	Args: cobra.ExactArgs(2),
	RunE: runCompare,
}

func init() {
	compareCmd.Flags().BoolVar(&compareShowUnchanged, "show-unchanged", false, "show unchanged URLs in output")
}

func runCompare(cmd *cobra.Command, args []string) error {
	ctx := GetContext()
	source1 := args[0]
	source2 := args[1]
	
	ctx.Formatter.Info(fmt.Sprintf("Comparing: %s vs %s", source1, source2))
	
	// Load first sitemap
	sitemap1, name1, err := loadSitemapFromSource(ctx, source1)
	if err != nil {
		ctx.Formatter.Error(fmt.Sprintf("Failed to load first sitemap: %v", err))
		return err
	}
	
	// Load second sitemap
	sitemap2, name2, err := loadSitemapFromSource(ctx, source2)
	if err != nil {
		ctx.Formatter.Error(fmt.Sprintf("Failed to load second sitemap: %v", err))
		return err
	}
	
	// Compare sitemaps
	added, removed, unchanged := compareSitemaps(sitemap1, sitemap2)
	
	// Output results
	if ctx.Config.OutputFormat == "json" {
		result := map[string]interface{}{
			"source1":   name1,
			"source2":   name2,
			"added":     len(added),
			"removed":   len(removed),
			"unchanged": len(unchanged),
			"added_urls":   added,
			"removed_urls": removed,
		}
		if compareShowUnchanged {
			result["unchanged_urls"] = unchanged
		}
		return ctx.Formatter.Print(result)
	}
	
	// Print summary
	fmt.Printf("\nComparison Results:\n")
	fmt.Printf("  Source 1: %s (%d URLs)\n", name1, len(sitemap1.URLs))
	fmt.Printf("  Source 2: %s (%d URLs)\n", name2, len(sitemap2.URLs))
	fmt.Printf("\n")
	fmt.Printf("  Added:     %d URLs\n", len(added))
	fmt.Printf("  Removed:   %d URLs\n", len(removed))
	fmt.Printf("  Unchanged: %d URLs\n", len(unchanged))
	fmt.Printf("\n")
	
	// Print differences
	if len(added) > 0 {
		fmt.Println("Added URLs:")
		for i, url := range added {
			if i >= 20 {
				fmt.Printf("  ... and %d more\n", len(added)-20)
				break
			}
			fmt.Printf("  + %s\n", url)
		}
		fmt.Println()
	}
	
	if len(removed) > 0 {
		fmt.Println("Removed URLs:")
		for i, url := range removed {
			if i >= 20 {
				fmt.Printf("  ... and %d more\n", len(removed)-20)
				break
			}
			fmt.Printf("  - %s\n", url)
		}
		fmt.Println()
	}
	
	if compareShowUnchanged && len(unchanged) > 0 {
		fmt.Println("Unchanged URLs (first 10):")
		limit := 10
		if len(unchanged) < limit {
			limit = len(unchanged)
		}
		for i := 0; i < limit; i++ {
			fmt.Printf("  = %s\n", unchanged[i])
		}
		if len(unchanged) > limit {
			fmt.Printf("  ... and %d more\n", len(unchanged)-limit)
		}
		fmt.Println()
	}
	
	// Use colored diff output
	if ctx.Config.ColorOutput {
		ctx.Formatter.PrintDiff(added, removed, nil)
	}
	
	ctx.Formatter.Success("Comparison complete")
	
	return nil
}

func loadSitemapFromSource(ctx *CLIContext, source string) (*sitemap.Sitemap, string, error) {
	// First try to load as a report ID from database
	reportService := services.NewReportService(ctx.DB)
	report, err := reportService.GetReport(context.Background(), source)
	if err == nil && report != nil {
		// Load entries from report
		sm, err := loadSitemapFromReport(ctx, source)
		if err == nil {
			return sm, fmt.Sprintf("Report: %s", source), nil
		}
	}
	
	// Otherwise, load as file or URL
	data, err := readSitemapSource(source)
	if err != nil {
		return nil, "", err
	}
	
	parser := sitemap.NewParser()
	sm, err := parser.Parse(data)
	if err != nil {
		return nil, "", err
	}
	
	return sm, source, nil
}

func loadSitemapFromReport(ctx *CLIContext, reportID string) (*sitemap.Sitemap, error) {
	// Get entries from database for this report
	entryRepo := ctx.DB.Entries()
	entries, err := entryRepo.List(context.Background(), repositories.EntryFilters{
		ReportID: reportID,
	})
	if err != nil {
		return nil, err
	}
	
	// Convert entries to sitemap
	sm := &sitemap.Sitemap{
		URLs: make([]sitemap.URL, 0, len(entries)),
	}
	
	for _, entry := range entries {
		u := sitemap.URL{
			Loc: entry.URL,
		}
		
		if entry.LastModified != nil {
			u.LastMod = entry.LastModified.Format("2006-01-02")
		}
		if entry.ChangeFreq != nil {
			u.ChangeFreq = *entry.ChangeFreq
		}
		if entry.Priority != nil {
			u.Priority = *entry.Priority
		}
		
		sm.URLs = append(sm.URLs, u)
	}
	
	return sm, nil
}

func compareSitemaps(sm1, sm2 *sitemap.Sitemap) (added, removed, unchanged []string) {
	// Build sets of URLs
	urls1 := make(map[string]bool)
	urls2 := make(map[string]bool)
	
	for _, u := range sm1.URLs {
		urls1[u.Loc] = true
	}
	
	for _, u := range sm2.URLs {
		urls2[u.Loc] = true
	}
	
	// Find added URLs (in sm2 but not in sm1)
	for _, u := range sm2.URLs {
		if !urls1[u.Loc] {
			added = append(added, u.Loc)
		}
	}
	
	// Find removed URLs (in sm1 but not in sm2)
	for _, u := range sm1.URLs {
		if !urls2[u.Loc] {
			removed = append(removed, u.Loc)
		}
	}
	
	// Find unchanged URLs (in both)
	for _, u := range sm1.URLs {
		if urls2[u.Loc] {
			unchanged = append(unchanged, u.Loc)
		}
	}
	
	return added, removed, unchanged
}

