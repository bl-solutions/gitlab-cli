package cmd

import (
	"fmt"
	"github.com/bl-solutions/gitlab-cli/gitlab"
	"github.com/spf13/viper"
	"strconv"

	"github.com/spf13/cobra"
)

// putCmd represents the put command
var putCmd = &cobra.Command{
	Use:   "put",
	Short: "Put variables of a project",
	Long: `Put variables of a project.

Examples:
· To put variables of project 389 from 389-vars.json (default filename):
    gitlab put 389

· To put variables of project 389 on vars.json file:
    gitlab put 389 -f vars.json`,
	Run: func(cmd *cobra.Command, args []string) {
		url := viper.GetString("gitlab.url")
		token := viper.GetString("gitlab.token")
		projectId, err := strconv.Atoi(args[0])
		cobra.CheckErr(err)

		client, err := gitlab.InitClient(url, token, projectId)
		cobra.CheckErr(err)

		if filename == "" {
			filename = fmt.Sprintf("%d-vars.json", projectId)
		}

		gitlab.HandlePutVariables(client, projectId, filename)
	},
}

func init() {
	rootCmd.AddCommand(putCmd)
}
