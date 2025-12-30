package cli

import (
	"fmt"

	"github.com/spf13/cobra"
	"jonopens/sitemapper/internal/cli/output"
	"jonopens/sitemapper/internal/config"
	"jonopens/sitemapper/internal/database"
	"jonopens/sitemapper/internal/repositories"
)

// CLIContext holds shared CLI state
type CLIContext struct {
	Config    *config.Config
	DB        repositories.Database
	Formatter *output.Formatter
}

var (
	cfgFile      string
	outputFormat string
	noColor      bool
	cliCtx       *CLIContext
)

// rootCmd represents the base command
var rootCmd = &cobra.Command{
	Use:   "sitemapper",
	Short: "A tool for tracking and comparing XML sitemaps",
	Long: `Sitemapper is a CLI tool for long-running XML sitemap comparison.
It helps you track sitemap changes over time, either on a schedule or ad hoc,
and allows for manual source file upload or fetch from a URL.`,
	PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
		// Skip initialization for help and interactive commands
		if cmd.Name() == "help" || cmd.Name() == "completion" {
			return nil
		}
		
		// Load configuration
		cfg, err := config.LoadConfigWithViper(cfgFile)
		if err != nil {
			return fmt.Errorf("failed to load config: %w", err)
		}
		
		// Override config with CLI flags
		if outputFormat != "" {
			cfg.OutputFormat = outputFormat
		}
		if noColor {
			cfg.ColorOutput = false
		}
		
		// Initialize database
		db, err := database.NewDatabase(cfg)
		if err != nil {
			return fmt.Errorf("failed to initialize database: %w", err)
		}
		
		// Create formatter
		formatter := output.NewFormatter(output.Format(cfg.OutputFormat), cfg.ColorOutput)
		
		// Store context
		cliCtx = &CLIContext{
			Config:    cfg,
			DB:        db,
			Formatter: formatter,
		}
		
		return nil
	},
	PersistentPostRunE: func(cmd *cobra.Command, args []string) error {
		// Clean up database connection
		if cliCtx != nil && cliCtx.DB != nil {
			return cliCtx.DB.Close()
		}
		return nil
	},
}

// Execute runs the root command
func Execute() error {
	return rootCmd.Execute()
}

func init() {
	// Global flags
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is ./configs/config.yaml)")
	rootCmd.PersistentFlags().StringVarP(&outputFormat, "format", "f", "", "output format (json, table, text)")
	rootCmd.PersistentFlags().BoolVar(&noColor, "no-color", false, "disable colored output")
	
	// Add subcommands
	rootCmd.AddCommand(parseCmd)
	rootCmd.AddCommand(compareCmd)
	rootCmd.AddCommand(trackCmd)
	rootCmd.AddCommand(reportCmd)
	rootCmd.AddCommand(groupingCmd)
	rootCmd.AddCommand(interactiveCmd)
}

// GetContext returns the CLI context
func GetContext() *CLIContext {
	return cliCtx
}

