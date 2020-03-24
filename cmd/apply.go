package cmd

import (
	"github.com/commitdev/kafka-connect-cli/pkg/apply"
	"github.com/spf13/cobra"

	"log"
)

func init() {
	rootCmd.AddCommand(applyCmd)
}

var applyCmd = &cobra.Command{
	Use:   "apply",
	Short: "Apply the configuration changes",
	Run:   Apply,
}

func Apply(cmd *cobra.Command, args []string) {
	err := apply.Apply(kafkaConnectorConfig, kcClient)
	if err != nil {
		log.Fatalf("All or part of the deployment failed. Error: %v", err)
	} else {
		log.Println("Deployment successful.")
	}
}
