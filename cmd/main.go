package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/RustyDaemon/goenvlist/internal"
	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

var (
	options = internal.NewOptions()

	rootCmd = &cobra.Command{
		Use:   "goenvlist",
		Short: "List environment variables with filtering options",
		Long: `goenvlist is a command-line tool for displaying environment variables.

		Examples:
		  goenvlist               # List all environment variables
		  goenvlist -s            # Show common environment variables
		  goenvlist -p            # Show only PATH variable
		  goenvlist -r		      # Show raw output without formatting
		  goenvlist -f HOME,USER  # Show only HOME and USER variables`,
		RunE: func(cmd *cobra.Command, args []string) error {
			return run()
		},
	}
)

func init() {
	rootCmd.PersistentFlags().BoolVarP(&options.Simple, "simple", "s", false, "Show only common environment variables")
	rootCmd.PersistentFlags().BoolVarP(&options.Path, "path", "p", false, "Show only PATH variable")
	rootCmd.PersistentFlags().BoolVarP(&options.Raw, "raw", "r", false, "Show unformatted output (useful for scripting)")
	rootCmd.PersistentFlags().StringSliceVarP(&options.Filter, "filter", "f", []string{}, "Show only specific variables (comma-separated)")
}

func run() error {
	variables, err := internal.GetEnvironment()
	if err != nil {
		return fmt.Errorf("error getting environment variables: %w", err)
	}

	formatter := internal.NewFormatter(options, variables)
	formatter.Display()

	return nil
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		if !strings.Contains(err.Error(), "help requested") {
			color.Red("Error: %v", err)
		}
		os.Exit(1)
	}
}
