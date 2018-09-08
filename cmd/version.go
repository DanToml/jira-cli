package cmd

import (
	"fmt"

	"github.com/endocrimes/jira-cli/internal/version"
	"github.com/spf13/cobra"
)

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version and exit",
	RunE: func(cmd *cobra.Command, args []string) error {
		fmt.Println("https://github.com/endocrimes/jira-cli")
		fmt.Printf("%s-%s\n", version.VERSION, version.GITCOMMIT)
		return nil
	},
}
