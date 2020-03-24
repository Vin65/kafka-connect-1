package cmd

import (
	"errors"
	"github.com/commitdev/kafka-connect-cli/pkg/delete"
	"github.com/spf13/cobra"
	"log"
)

func init() {
	rootCmd.AddCommand(deleteCmd)
}

var deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete a config from kafka connect.",
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) < 1 && kafkaConnectorConfig == nil {
			return errors.New("requires name of connector as an argument or a supplied config")
		}
		return nil
	},
	Run: Delete,
}

func Delete(cmd *cobra.Command, args []string) {
	if len(args) > 0 {
		name := args[0]
		err := delete.Delete(kcClient, name)
		if err != nil {
			log.Printf("Failed to delete %s. Error: %v", name, err)
			return
		}

	} else {
		err := delete.DeleteFromConfig(kafkaConnectorConfig, kcClient)
		if err != nil {
			log.Printf("Failed to delete from config. Error: %v", err)
			return
		}

	}

	log.Printf("Successfully deleted connector(s).")

}
