package cmd

import (
	"errors"
	"github.com/commitdev/kafka-connect/pkg/get"
	"github.com/spf13/cobra"
	"log"
)

func init() {
	rootCmd.AddCommand(getCmd)
}

var getCmd = &cobra.Command{
	Use:   "get",
	Short: "Get a config from kafka connect.",
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) < 1 {
			return errors.New("requires name of connector as an argument")
		}
		return nil
	},
	Run: Get,
}

func Get(cmd *cobra.Command, args []string) {
	name := args[0]
	err := get.Get(kafkaConnectorConfig, kcClient, name)
	if err != nil {
		log.Fatal(err)
	}
}
