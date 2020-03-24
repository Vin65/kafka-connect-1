package apply

import (
	"errors"
	"github.com/commitdev/kafka-connect-cli/config"
	"github.com/commitdev/kafka-connect-cli/pkg/client"
	"log"
)

func Apply(config *config.KafkaConnectorConfig, kclient client.KafkaConnectClient) error {
	instances := config.MergeAll()
	if len(instances) == 0 {
		return errors.New("No configuration to deploy!")
	}

	log.Printf("Attempting to apply [%v] instances", len(instances))

	for _, instance := range instances {
		resp, err := kclient.Get(instance.Name)
		if err != nil {
			_, ok := err.(*client.NotFound)
			if ok {
				log.Printf("Attempting to create %v ...", instance.Name)
				err = kclient.Create(&instance)
				if err != nil {
					log.Printf("Failed to create %s. Error: %s", instance.Name, err)
				}
			}

		} else {
			log.Printf("Config exists attempting to update %v ...", resp.Name)
			err = kclient.Update(instance.Name, &instance.Config)
			if err != nil {
				log.Printf("Failed to update %s. Error: %s", instance.Name, err)
			}

		}

	}

	return nil
}
