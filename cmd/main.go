package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"os"
	"regexp"
	"strconv"
)

var (
	// Used for flags.
	cfgFile  string
	filename string

	rootCmd = &cobra.Command{
		Use: "gitlab",
		Args: func(cmd *cobra.Command, args []string) error {
			// Check number of arguments
			if err := cobra.ExactArgs(2)(cmd, args); err != nil {
				return err
			}

			// Check if first arg is a valid action
			regex := regexp.MustCompile("^(get|put|flush)$")
			if !regex.MatchString(args[0]) {
				return fmt.Errorf("invalid action specified: %v", args[0])
			}

			// Check if second arg is an int
			if _, err := strconv.Atoi(args[1]); err != nil {
				return fmt.Errorf("invalid project id specified: %v", args[1])
			}

			return nil
		},
		Run: func(cmd *cobra.Command, args []string) {
		},
	}
)

// Execute executes the root command.
func Execute() error {
	return rootCmd.Execute()
}

func init() {
	cobra.OnInitialize(initConfig)
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is ~/.gitlab.yaml)")
	rootCmd.PersistentFlags().StringVarP(&filename, "filename", "f", "", "json file used")
}

func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory
		home, err := os.UserHomeDir()
		cobra.CheckErr(err)

		// Search config in home directory with name ".cobra" (without extension).
		viper.AddConfigPath(home)
		viper.AddConfigPath(".")
		viper.SetConfigType("yaml")
		viper.SetConfigName(".gitlab")
	}

	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	}
}
