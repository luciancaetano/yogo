package cli

import (
	"errors"
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

const newTPL = `
version: 1.0.0

endpoints:
  - name: GetUser
    method: GET
    path: /users/{id}
    responses:
      200:
        id: int
        name: string
        email: string
      404:
        error: string
  - name: CreateUser
    method: POST
    path: /users
    request:
      name: string
      email: string
    responses:
      201:
        id: int
      400:
        error: string

generators:
  - name: go
    package: apiclient
    output: ./gen/generate_api.go
`

func newCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "new",
		Short: "Create a new yogo.yml configuration file",
		RunE: func(cmd *cobra.Command, args []string) error {
			const fileName = "yogo.yml"

			if _, err := os.Stat(fileName); err == nil {
				return errors.New("yogo.yml already exists, will not overwrite")
			} else if !os.IsNotExist(err) {
				return fmt.Errorf("failed to check if yogo.yml exists: %w", err)
			}

			if err := os.WriteFile(fileName, []byte(newTPL), 0644); err != nil {
				return fmt.Errorf("failed to write yogo.yml: %w", err)
			}

			fmt.Println("Created new yogo.yml file")
			return nil
		},
	}

	return cmd
}
