package cmd

import (
	"errors"
	"github.com/commitdev/kafka-connect-cli/pkg/describe"
	"github.com/spf13/cobra"
	"log"
)

func init() {
	rootCmd.AddCommand(describeCmd)
}

var describeCmd = &cobra.Command{
	Use:   "describe",
	Short: "Describe config that exists in kafka connect.",
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) < 1 {
			return errors.New("requires name of connector as an argument")
		}
		return nil
	},
	Run: Describe,
}

func Describe(cmd *cobra.Command, args []string) {
	name := args[0]
	err := describe.Describe(kcClient, name)
	if err != nil {
		log.Printf("Failed to describe %s. Error: %v", name, err)
		return
	}

}
