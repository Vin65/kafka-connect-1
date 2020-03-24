package get

import (
	"encoding/json"
	"github.com/commitdev/kafka-connect/config"
	"github.com/commitdev/kafka-connect/pkg/client"
	"github.com/tidwall/pretty"
)

func Get(config *config.KafkaConnectorConfig, client client.KafkaConnectClient, name string) error {
	result, err := client.Get(name)
	if err != nil {
		return err
	}

	j, _ := json.Marshal(result)
	formatted := pretty.Color(pretty.Pretty(j), nil)
	println(string(formatted))
	return nil
}
