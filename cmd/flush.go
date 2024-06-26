package cmd

import (
	"github.com/spf13/viper"
	"gitlab/gitlab"
	"strconv"

	"github.com/spf13/cobra"
)

// flushCmd represents the flush command
var flushCmd = &cobra.Command{
	Use:   "flush",
	Short: "Flush all variables of a project",
	Long: `Flush all variables of a project.

Examples:
· To flush all variables of project 389:
    gitlab flush 389`,
	Run: func(cmd *cobra.Command, args []string) {
		url := viper.GetString("gitlab.url")
		token := viper.GetString("gitlab.token")
		projectId, err := strconv.Atoi(args[0])
		cobra.CheckErr(err)

		client, err := gitlab.InitClient(url, token, projectId)
		cobra.CheckErr(err)

		gitlab.HandleFlushVariables(client, projectId)
	},
}

func init() {
	rootCmd.AddCommand(flushCmd)
}
