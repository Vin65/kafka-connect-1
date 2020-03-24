package list

import (
	"encoding/json"
	"github.com/commitdev/kafka-connect-cli/config"
	"github.com/commitdev/kafka-connect-cli/pkg/client"
	"github.com/k0kubun/pp"
	"github.com/tidwall/pretty"
)

func List(config *config.KafkaConnectorConfig, client client.KafkaConnectClient) error {
	results, err := client.List()
	if err != nil {
		return err
	}

	j, _ := json.Marshal(results)
	formatted := pretty.Color(pretty.Pretty(j), nil)
	println(string(formatted))

	pp.Printf("Total configs: %v", len(results))

	return nil

}
