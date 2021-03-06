/*
 * Copyright © 2020 Mateusz Kyc
 */

package cmd

import (
	"fmt"
	"github.com/epiphany-platform/cli/pkg/configuration"
	"github.com/epiphany-platform/cli/pkg/environment"
	"github.com/google/uuid"
	"github.com/spf13/cobra"
)

// environmentsInfoCmd represents the info command
var environmentsInfoCmd = &cobra.Command{
	Use:   "info",
	Short: "Displays information about currently selected environment",
	Long:  `TODO`,
	PreRun: func(cmd *cobra.Command, args []string) {
		debug("environments info called")
	},
	Run: func(cmd *cobra.Command, args []string) {
		config, err := configuration.GetConfig()
		if err != nil {
			errGetConfig(err)
		}
		if config.CurrentEnvironment == uuid.Nil {
			errNilEnvironment()
		}
		environment, err := environment.Get(config.CurrentEnvironment)
		if err != nil {
			errGetEnvironmentDetails(err)
		}
		fmt.Print(environment.String())
	},
}

func init() {
	environmentsCmd.AddCommand(environmentsInfoCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// environmentsInfoCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// environmentsInfoCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
