package delete

import (
	"github.com/commitdev/kafka-connect-cli/config"
	"github.com/commitdev/kafka-connect-cli/pkg/client"
	"log"
)

func Delete(kclient client.KafkaConnectClient, configName string) error {
	log.Printf("Attempting to delete %s connector", configName)

	return kclient.Delete(configName)
}

func DeleteFromConfig(config *config.KafkaConnectorConfig, kclient client.KafkaConnectClient) error {
	instances := config.MergeAll()

	log.Printf("Attempting to delete [%v] connectors", len(instances))

	for _, instance := range instances {
		err := kclient.Delete(instance.Name)
		if err != nil {
			log.Printf("Failed to delete config: %s", instance.Name)
			return err
		}

	}

	return nil
}
