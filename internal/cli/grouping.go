package cli

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/spf13/cobra"
	"jonopens/sitemapper/internal/models"
)

var groupingCmd = &cobra.Command{
	Use:   "grouping",
	Short: "Manage URL groupings",
	Long:  `Create and manage URL groupings for organizing sitemap entries.`,
}

var groupingListCmd = &cobra.Command{
	Use:   "list",
	Short: "List all groupings",
	Long:  `List all URL groupings.`,
	RunE:  runGroupingList,
}

var groupingCreateCmd = &cobra.Command{
	Use:   "create",
	Short: "Create a new grouping",
	Long:  `Create a new URL grouping.`,
	RunE:  runGroupingCreate,
}

var (
	groupingName        string
	groupingDescription string
	groupingUserID      string
)

func init() {
	// Create command flags
	groupingCreateCmd.Flags().StringVar(&groupingName, "name", "", "name for the grouping (required)")
	groupingCreateCmd.Flags().StringVar(&groupingDescription, "description", "", "description for the grouping")
	groupingCreateCmd.Flags().StringVar(&groupingUserID, "user-id", "", "user ID (defaults to config default_user_id)")
	groupingCreateCmd.MarkFlagRequired("name")
	
	// Add subcommands
	groupingCmd.AddCommand(groupingListCmd)
	groupingCmd.AddCommand(groupingCreateCmd)
}

func runGroupingList(cmd *cobra.Command, args []string) error {
	ctx := GetContext()
	
	ctx.Formatter.Info("Listing groupings")
	
	// Get groupings from database
	groupingRepo := ctx.DB.Groupings()
	groupings, err := groupingRepo.List(context.Background())
	if err != nil {
		ctx.Formatter.Error(fmt.Sprintf("Failed to list groupings: %v", err))
		return err
	}
	
	if len(groupings) == 0 {
		ctx.Formatter.Info("No groupings found")
		return nil
	}
	
	// Output results
	if ctx.Config.OutputFormat == "json" {
		return ctx.Formatter.Print(groupings)
	}
	
	// Print as table
	fmt.Printf("\nFound %d grouping(s):\n\n", len(groupings))
	
	rows := [][]string{
		{"ID", "Name", "User ID", "Description", "Created"},
	}
	
	for _, grouping := range groupings {
		desc := ""
		if grouping.Description != nil {
			desc = truncate(*grouping.Description, 30)
		}
		rows = append(rows, []string{
			truncate(grouping.ID, 20),
			truncate(grouping.Name, 20),
			truncate(grouping.UserID, 15),
			desc,
			grouping.CreatedAt.Format("2006-01-02"),
		})
	}
	
	ctx.Formatter.Print(rows)
	
	fmt.Printf("\nTotal: %d grouping(s)\n", len(groupings))
	
	return nil
}

func runGroupingCreate(cmd *cobra.Command, args []string) error {
	ctx := GetContext()
	
	// Use default user ID from config if not provided
	if groupingUserID == "" {
		groupingUserID = ctx.Config.DefaultUserID
	}
	
	ctx.Formatter.Info(fmt.Sprintf("Creating grouping: %s", groupingName))
	
	// Create grouping
	grouping := &models.Group{
		ID:          uuid.New().String(),
		UserID:      groupingUserID,
		Name:        groupingName,
		Description: stringPtrOrNil(groupingDescription),
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}
	
	// Save to database
	groupingRepo := ctx.DB.Groupings()
	if err := groupingRepo.Create(context.Background(), grouping); err != nil {
		ctx.Formatter.Error(fmt.Sprintf("Failed to create grouping: %v", err))
		return err
	}
	
	ctx.Formatter.Success(fmt.Sprintf("Grouping created with ID: %s", grouping.ID))
	
	// Output grouping details
	if ctx.Config.OutputFormat == "json" {
		return ctx.Formatter.Print(grouping)
	}
	
	fmt.Printf("\nGrouping Details:\n")
	fmt.Printf("  ID:          %s\n", grouping.ID)
	fmt.Printf("  Name:        %s\n", grouping.Name)
	fmt.Printf("  User ID:     %s\n", grouping.UserID)
	if grouping.Description != nil {
		fmt.Printf("  Description: %s\n", *grouping.Description)
	}
	fmt.Printf("  Created:     %s\n", grouping.CreatedAt.Format("2006-01-02 15:04:05"))
	fmt.Println()
	
	return nil
}

func stringPtrOrNil(s string) *string {
	if s == "" {
		return nil
	}
	return &s
}

