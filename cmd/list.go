package cmd

import (
	"github.com/commitdev/kafka-connect/pkg/list"
	"github.com/spf13/cobra"
	"log"
)

func init() {
	rootCmd.AddCommand(listCmd)
}

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List all current kafka connect configs.",
	Run:   List,
}

func List(cmd *cobra.Command, args []string) {

	err := list.List(kafkaConnectorConfig, kcClient)
	if err != nil {
		log.Fatal(err)
	}
}
