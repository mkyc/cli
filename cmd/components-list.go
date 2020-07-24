/*
 * Copyright © 2020 Mateusz Kyc
 */

package cmd

import (
	"fmt"
	"github.com/mkyc/epiphany-wrapper-poc/pkg/repository"
	"github.com/spf13/cobra"
)

// componentsListCmd represents the list command
var componentsListCmd = &cobra.Command{
	Use:   "list",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("list called")
		for _, c := range repository.GetRepository().Components {
			for _, v := range c.Versions {
				fmt.Printf("%s %s", c.Name, v.Version)
			}
		}
	},
}

func init() {
	componentsCmd.AddCommand(componentsListCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// componentsListCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// componentsListCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
