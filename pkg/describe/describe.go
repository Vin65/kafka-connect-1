package describe

import (
	"encoding/json"
	"github.com/commitdev/kafka-connect-cli/pkg/client"
	"github.com/tidwall/pretty"
)

func Describe(kclient client.KafkaConnectClient, configName string) error {
	result, err := kclient.GetStatus(configName)
	if err != nil {
		return err
	}

	j, _ := json.Marshal(result)
	formatted := pretty.Color(pretty.Pretty(j), nil)
	println(string(formatted))
	return nil
}