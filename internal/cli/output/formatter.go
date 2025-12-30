package output

import (
	"encoding/json"
	"fmt"
	"io"
	"os"

	"github.com/fatih/color"
	"github.com/olekukonko/tablewriter"
)

// Format represents the output format type
type Format string

const (
	FormatJSON  Format = "json"
	FormatTable Format = "table"
	FormatText  Format = "text"
)

// Formatter handles different output formats
type Formatter struct {
	format      Format
	writer      io.Writer
	colorOutput bool
}

// NewFormatter creates a new output formatter
func NewFormatter(format Format, colorOutput bool) *Formatter {
	return &Formatter{
		format:      format,
		writer:      os.Stdout,
		colorOutput: colorOutput,
	}
}

// SetWriter sets the output writer
func (f *Formatter) SetWriter(w io.Writer) {
	f.writer = w
}

// Print outputs data in the configured format
func (f *Formatter) Print(data interface{}) error {
	switch f.format {
	case FormatJSON:
		return f.printJSON(data)
	case FormatTable:
		return f.printTable(data)
	case FormatText:
		return f.printText(data)
	default:
		return fmt.Errorf("unsupported format: %s", f.format)
	}
}

// printJSON outputs data as formatted JSON
func (f *Formatter) printJSON(data interface{}) error {
	encoder := json.NewEncoder(f.writer)
	encoder.SetIndent("", "  ")
	return encoder.Encode(data)
}

// printTable outputs data as a formatted table
func (f *Formatter) printTable(data interface{}) error {
	table := tablewriter.NewWriter(f.writer)
	
	// Configure table style
	table.SetAutoWrapText(false)
	table.SetAutoFormatHeaders(true)
	table.SetHeaderAlignment(tablewriter.ALIGN_LEFT)
	table.SetAlignment(tablewriter.ALIGN_LEFT)
	table.SetCenterSeparator("")
	table.SetColumnSeparator("")
	table.SetRowSeparator("")
	table.SetHeaderLine(false)
	table.SetBorder(false)
	table.SetTablePadding("\t")
	table.SetNoWhiteSpace(true)
	
	// Handle different data types
	switch v := data.(type) {
	case [][]string:
		table.AppendBulk(v)
		table.Render()
	case map[string]interface{}:
		for key, val := range v {
			table.Append([]string{key, fmt.Sprintf("%v", val)})
		}
		table.Render()
	default:
		// Fallback to JSON for complex types
		return f.printJSON(data)
	}
	
	return nil
}

// printText outputs data as plain text
func (f *Formatter) printText(data interface{}) error {
	_, err := fmt.Fprintln(f.writer, data)
	return err
}

// Success prints a success message
func (f *Formatter) Success(message string) {
	if f.colorOutput {
		color.New(color.FgGreen).Fprintln(f.writer, "✓", message)
	} else {
		fmt.Fprintln(f.writer, "[SUCCESS]", message)
	}
}

// Error prints an error message
func (f *Formatter) Error(message string) {
	if f.colorOutput {
		color.New(color.FgRed).Fprintln(f.writer, "✗", message)
	} else {
		fmt.Fprintln(f.writer, "[ERROR]", message)
	}
}

// Warning prints a warning message
func (f *Formatter) Warning(message string) {
	if f.colorOutput {
		color.New(color.FgYellow).Fprintln(f.writer, "⚠", message)
	} else {
		fmt.Fprintln(f.writer, "[WARNING]", message)
	}
}

// Info prints an info message
func (f *Formatter) Info(message string) {
	if f.colorOutput {
		color.New(color.FgCyan).Fprintln(f.writer, "ℹ", message)
	} else {
		fmt.Fprintln(f.writer, "[INFO]", message)
	}
}

// PrintDiff prints a diff with colors
func (f *Formatter) PrintDiff(added, removed, unchanged []string) {
	if f.colorOutput {
		green := color.New(color.FgGreen)
		red := color.New(color.FgRed)
		
		for _, item := range added {
			green.Fprintf(f.writer, "+ %s\n", item)
		}
		for _, item := range removed {
			red.Fprintf(f.writer, "- %s\n", item)
		}
		for _, item := range unchanged {
			fmt.Fprintf(f.writer, "  %s\n", item)
		}
	} else {
		for _, item := range added {
			fmt.Fprintf(f.writer, "+ %s\n", item)
		}
		for _, item := range removed {
			fmt.Fprintf(f.writer, "- %s\n", item)
		}
		for _, item := range unchanged {
			fmt.Fprintf(f.writer, "  %s\n", item)
		}
	}
}

