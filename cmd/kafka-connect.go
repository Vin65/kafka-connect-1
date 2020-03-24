package cmd

import (
	"github.com/commitdev/kafka-connect-cli/config"
	"github.com/commitdev/kafka-connect-cli/pkg/client"
	"github.com/spf13/cobra"
	"log"
)

var (
	kafkaConnectorConfigFile string
	kafkaConnectorConfig     = &config.KafkaConnectorConfig{}
	rootCmd                  = &cobra.Command{
		Use:              "kafka-connect",
		Short:            "Deployment tool to help deploy kafka connect configuration.",
		PersistentPreRun: PreloadConfig,
	}
	kafkaConnectAddress  string
	kafkaConnectUser     string
	kafkaConnectPassword string
	kcClient             client.KafkaConnectClient
)

func PreloadConfig(cmd *cobra.Command, args []string) {
	if kafkaConnectorConfigFile != "" {
		cfg, err := config.LoadConfig(kafkaConnectorConfigFile)
		if err != nil {
			log.Printf("Error loading config: %v", err)
			return
		}
		kafkaConnectorConfig = cfg
	}

	client, err := client.NewKafkaConnectClient(kafkaConnectAddress, kafkaConnectUser, kafkaConnectPassword)
	if err != nil {
		log.Fatalf("Error: %v", err)
	}

	kcClient = client
}

func Execute() error {

	return rootCmd.Execute()
}

func init() {
	rootCmd.PersistentFlags().StringVarP(&kafkaConnectorConfigFile, "config", "c", "", "Path to kafka connect connect deployer config.")
	rootCmd.PersistentFlags().StringVarP(&kafkaConnectAddress, "address", "a", "", "URL to kafka connect.")
	rootCmd.PersistentFlags().StringVarP(&kafkaConnectUser, "user", "u", "", "User for kafka connect.")
	rootCmd.PersistentFlags().StringVarP(&kafkaConnectPassword, "password", "p", "", "Password for kafka connect.")
}
