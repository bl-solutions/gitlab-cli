package cmd

import (
	"fmt"
	"github.com/bl-solutions/gitlab-cli/gitlab"
	"github.com/spf13/viper"
	"strconv"

	"github.com/spf13/cobra"
)

// getCmd represents the get command
var getCmd = &cobra.Command{
	Use:   "get",
	Short: "Get variables of a project",
	Long: `Get variables of a project.

Examples:
· To get variables of project 389 on 389-vars.json file (default filename):
    gitlab get 389

· To get variables of project 389 on vars.json file:
    gitlab get 389 -f vars.json`,
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

		gitlab.HandleGetVariables(client, projectId, filename)
	},
}

func init() {
	rootCmd.AddCommand(getCmd)
}
