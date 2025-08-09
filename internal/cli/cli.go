package cli

import (
	"github.com/spf13/cobra"
)

func Run() error {
	rootCmd := &cobra.Command{
		Use:   "yogo",
		Short: "YOGO - Rest API client generator",
	}

	rootCmd.AddCommand(newCmd())
	rootCmd.AddCommand(generateCmd())

	if err := rootCmd.Execute(); err != nil {
		return err
	}
	return nil
}
