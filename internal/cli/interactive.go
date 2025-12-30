package cli

import (
	"fmt"
	"os"
	"strings"

	prompt "github.com/c-bata/go-prompt"
	"github.com/spf13/cobra"
)

var interactiveCmd = &cobra.Command{
	Use:   "interactive",
	Short: "Launch interactive shell mode",
	Long: `Launch an interactive REPL shell for running multiple commands.
Type 'help' for available commands, 'exit' or 'quit' to exit.`,
	RunE: runInteractive,
}

func runInteractive(cmd *cobra.Command, args []string) error {
	ctx := GetContext()
	
	ctx.Formatter.Success("Welcome to Sitemapper Interactive Mode")
	fmt.Println("Type 'help' for available commands, 'exit' or 'quit' to exit.")
	fmt.Println()
	
	// Run interactive prompt
	p := prompt.New(
		executor,
		completer,
		prompt.OptionPrefix("sitemapper> "),
		prompt.OptionTitle("sitemapper interactive"),
		prompt.OptionPrefixTextColor(prompt.Green),
		prompt.OptionSuggestionBGColor(prompt.DarkGray),
		prompt.OptionSuggestionTextColor(prompt.White),
		prompt.OptionSelectedSuggestionBGColor(prompt.LightGray),
		prompt.OptionSelectedSuggestionTextColor(prompt.Black),
	)
	p.Run()
	
	return nil
}

func executor(input string) {
	input = strings.TrimSpace(input)
	
	// Handle exit commands
	if input == "exit" || input == "quit" {
		fmt.Println("Goodbye!")
		os.Exit(0)
	}
	
	// Handle empty input
	if input == "" {
		return
	}
	
	// Parse command and execute
	args := strings.Fields(input)
	if len(args) == 0 {
		return
	}
	
	// Reset root command flags for each execution
	rootCmd.SetArgs(args)
	
	// Execute the command
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
	}
	
	fmt.Println() // Add spacing between commands
}

func completer(d prompt.Document) []prompt.Suggest {
	suggestions := []prompt.Suggest{
		// Main commands
		{Text: "parse", Description: "Parse and validate a sitemap"},
		{Text: "compare", Description: "Compare two sitemaps"},
		{Text: "track", Description: "Track a sitemap snapshot"},
		{Text: "report", Description: "Manage reports"},
		{Text: "grouping", Description: "Manage groupings"},
		{Text: "help", Description: "Show help information"},
		{Text: "exit", Description: "Exit interactive mode"},
		{Text: "quit", Description: "Exit interactive mode"},
		
		// Report subcommands
		{Text: "report list", Description: "List all reports"},
		{Text: "report get", Description: "Get a specific report"},
		
		// Grouping subcommands
		{Text: "grouping list", Description: "List all groupings"},
		{Text: "grouping create", Description: "Create a new grouping"},
		
		// Common flags
		{Text: "--help", Description: "Show help for a command"},
		{Text: "--format", Description: "Output format (json, table, text)"},
		{Text: "--no-color", Description: "Disable colored output"},
		{Text: "--validate", Description: "Validate sitemap (parse command)"},
		{Text: "--show-stats", Description: "Show statistics (parse command)"},
		{Text: "--name", Description: "Name for snapshot or grouping"},
		{Text: "--user-id", Description: "User ID"},
	}
	
	return prompt.FilterHasPrefix(suggestions, d.GetWordBeforeCursor(), true)
}

