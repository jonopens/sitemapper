package cli

import (
	"fmt"
	"io"
	"net/url"
	"os"
	"time"

	"github.com/spf13/cobra"
	"jonopens/sitemapper/pkg/http"
	"jonopens/sitemapper/pkg/sitemap"
)

var (
	parseValidate  bool
	parseShowStats bool
)

var parseCmd = &cobra.Command{
	Use:   "parse <url|file>",
	Short: "Parse and validate a sitemap from file or URL",
	Long: `Parse and optionally validate an XML sitemap.
Supports both local files and remote URLs.
Can detect sitemap type (sitemap vs sitemap index) and show statistics.`,
	Args: cobra.ExactArgs(1),
	RunE: runParse,
}

func init() {
	parseCmd.Flags().BoolVar(&parseValidate, "validate", false, "validate sitemap structure")
	parseCmd.Flags().BoolVar(&parseShowStats, "show-stats", false, "show sitemap statistics")
}

func runParse(cmd *cobra.Command, args []string) error {
	ctx := GetContext()
	source := args[0]
	
	ctx.Formatter.Info(fmt.Sprintf("Parsing sitemap from: %s", source))
	
	// Read sitemap data
	data, err := readSitemapSource(source)
	if err != nil {
		ctx.Formatter.Error(fmt.Sprintf("Failed to read sitemap: %v", err))
		return err
	}
	
	// Parse sitemap
	parser := sitemap.NewParser()
	sitemapType, err := parser.DetectType(data)
	if err != nil {
		ctx.Formatter.Error(fmt.Sprintf("Failed to detect sitemap type: %v", err))
		return err
	}
	
	ctx.Formatter.Info(fmt.Sprintf("Detected type: %s", sitemapType))
	
	switch sitemapType {
	case "sitemap":
		return parseSitemap(ctx, parser, data)
	case "index":
		return parseSitemapIndex(ctx, parser, data)
	default:
		return fmt.Errorf("unknown sitemap type: %s", sitemapType)
	}
}

func parseSitemap(ctx *CLIContext, parser *sitemap.Parser, data []byte) error {
	sm, err := parser.Parse(data)
	if err != nil {
		ctx.Formatter.Error(fmt.Sprintf("Failed to parse sitemap: %v", err))
		return err
	}
	
	// Validate if requested
	if parseValidate {
		validator := sitemap.NewValidator()
		if err := validator.Validate(sm); err != nil {
			ctx.Formatter.Error(fmt.Sprintf("Validation failed: %v", err))
			return err
		}
		ctx.Formatter.Success("Sitemap is valid")
	}
	
	// Show stats if requested
	if parseShowStats {
		showSitemapStats(ctx, sm)
	}
	
	// Output sitemap data
	if ctx.Config.OutputFormat == "json" {
		return ctx.Formatter.Print(sm)
	}
	
	ctx.Formatter.Success(fmt.Sprintf("Successfully parsed sitemap with %d URLs", len(sm.URLs)))
	
	// Show sample URLs in table format
	if len(sm.URLs) > 0 {
		fmt.Println("\nSample URLs (first 10):")
		rows := [][]string{
			{"URL", "Last Modified", "Change Freq", "Priority"},
		}
		
		limit := 10
		if len(sm.URLs) < limit {
			limit = len(sm.URLs)
		}
		
		for i := 0; i < limit; i++ {
			u := sm.URLs[i]
			priority := ""
			if u.Priority > 0 {
				priority = fmt.Sprintf("%.1f", u.Priority)
			}
			rows = append(rows, []string{
				truncate(u.Loc, 60),
				u.LastMod,
				u.ChangeFreq,
				priority,
			})
		}
		
		ctx.Formatter.Print(rows)
	}
	
	return nil
}

func parseSitemapIndex(ctx *CLIContext, parser *sitemap.Parser, data []byte) error {
	index, err := parser.ParseIndex(data)
	if err != nil {
		ctx.Formatter.Error(fmt.Sprintf("Failed to parse sitemap index: %v", err))
		return err
	}
	
	// Output index data
	if ctx.Config.OutputFormat == "json" {
		return ctx.Formatter.Print(index)
	}
	
	ctx.Formatter.Success(fmt.Sprintf("Successfully parsed sitemap index with %d sitemaps", len(index.Sitemaps)))
	
	return nil
}

func showSitemapStats(ctx *CLIContext, sm *sitemap.Sitemap) {
	stats := map[string]interface{}{
		"Total URLs":   len(sm.URLs),
		"With LastMod": countWithLastMod(sm),
		"With Priority": countWithPriority(sm),
		"With ChangeFreq": countWithChangeFreq(sm),
	}
	
	fmt.Println("\nStatistics:")
	for key, val := range stats {
		fmt.Printf("  %s: %v\n", key, val)
	}
}

func countWithLastMod(sm *sitemap.Sitemap) int {
	count := 0
	for _, u := range sm.URLs {
		if u.LastMod != "" {
			count++
		}
	}
	return count
}

func countWithPriority(sm *sitemap.Sitemap) int {
	count := 0
	for _, u := range sm.URLs {
		if u.Priority > 0 {
			count++
		}
	}
	return count
}

func countWithChangeFreq(sm *sitemap.Sitemap) int {
	count := 0
	for _, u := range sm.URLs {
		if u.ChangeFreq != "" {
			count++
		}
	}
	return count
}

func readSitemapSource(source string) ([]byte, error) {
	// Check if source is a URL
	if u, err := url.Parse(source); err == nil && (u.Scheme == "http" || u.Scheme == "https") {
		// Fetch from URL
		client := http.NewRetryClient(3, 30*time.Second)
		resp, err := client.Get(source)
		if err != nil {
			return nil, fmt.Errorf("failed to fetch URL: %w", err)
		}
		defer resp.Body.Close()
		
		if resp.StatusCode != 200 {
			return nil, fmt.Errorf("unexpected status code: %d", resp.StatusCode)
		}
		
		return io.ReadAll(resp.Body)
	}
	
	// Read from file
	return os.ReadFile(source)
}

func truncate(s string, maxLen int) string {
	if len(s) <= maxLen {
		return s
	}
	return s[:maxLen-3] + "..."
}

