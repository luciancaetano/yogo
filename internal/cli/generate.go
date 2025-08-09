package cli

import (
	"fmt"
	gogenerator "github/luciancaetano/yogo/internal/go-generator"
	"github/luciancaetano/yogo/internal/yogofile"
	"go/format"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
)

func generateCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "generate",
		Short: "Generate client code from yogo.yml spec",
		RunE: func(cmd *cobra.Command, args []string) error {
			const inputFile = "yogo.yml"

			yogo := yogofile.New()

			data, err := os.ReadFile(inputFile)
			if err != nil {
				return err
			}
			if err := yogo.Unmarshal(data); err != nil {
				return err
			}
			if err := yogo.Validate(); err != nil {
				return err
			}

			if len(yogo.Generators) == 0 {
				return fmt.Errorf("no generator configuration found")
			}

			if !yogo.ContainsGenerator("go") {
				return fmt.Errorf("no 'go' generator configuration found")
			}

			gen := yogo.GetGenerator("go")

			if gen.Package == "" {
				return fmt.Errorf("package name is required in generator configuration")
			}

			if gen.Output == "" {
				return fmt.Errorf("output file name is required in generator configuration")
			}

			outputDir := filepath.Dir(gen.Output)
			if err := os.MkdirAll(outputDir, 0755); err != nil {
				return fmt.Errorf("failed to create output directory: %w", err)
			}

			code, err := gogenerator.Generate(yogo)
			if err != nil {
				return fmt.Errorf("generate error: %w", err)
			}

			formattedCode, err := format.Source([]byte(code))
			if err != nil {
				return fmt.Errorf("error formatting generated code: %w", err)
			}

			if err := os.WriteFile(gen.Output, []byte(formattedCode), 0644); err != nil {
				return fmt.Errorf("write file error: %w", err)
			}

			fmt.Println("Client code generated at", gen.Output)
			return nil
		},
	}

	return cmd
}
