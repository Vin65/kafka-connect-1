package spit

import (
	"encoding/json"
	"github.com/commitdev/kafka-connect/config"
	"github.com/k0kubun/pp"
	"github.com/tidwall/pretty"
)

func Spit(config *config.KafkaConnectorConfig) {
	instances := config.MergeAll()
	for _, instance := range instances {
		j, _ := json.Marshal(instance)
		formatted := pretty.Color(pretty.Pretty(j), nil)
		println(string(formatted))
	}

	pp.Printf("Total configs generated: %v", len(instances))
}
