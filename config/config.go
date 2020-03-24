package config

import (
	"fmt"
	"github.com/imdario/mergo"
	"github.com/jinzhu/configor"
)

const DBName = "database.dbname"
const TableName = "table.whitelist"

type KafkaConnectorConfigInstance map[string]interface{}

type KafkaConnectorConfigWrapper struct {
	Name   string                       `json:"name"`
	Config KafkaConnectorConfigInstance `json:"config"`
}

func (c KafkaConnectorConfigInstance) copy() KafkaConnectorConfigInstance {
	newMap := make(map[string]interface{})
	for key, value := range c {
		newMap[key] = value
	}
	return newMap
}

type KafkaConnectorConfig struct {
	Version        string                         `yaml:"version"`
	GlobalConfig   KafkaConnectorConfigInstance   `yaml:"globals"`
	DatabaseConfig []KafkaConnectorConfigInstance `yaml:"databases"`
	TableConfig    []KafkaConnectorConfigInstance `yaml:"tables"`
}

func (c *KafkaConnectorConfig) MergeAll() []KafkaConnectorConfigWrapper {
	var instances = []KafkaConnectorConfigWrapper{}

	for _, db := range c.DatabaseConfig {
		configInstance := c.GlobalConfig
		mergo.MergeWithOverwrite(&configInstance, db)
		for _, table := range c.TableConfig {
			mergo.MergeWithOverwrite(&configInstance, table)
			dbName := db[DBName]
			tableName := table[TableName]
			connectorName := fmt.Sprintf("%s-%s-v%s", tableName, dbName, c.Version)

			wrap := KafkaConnectorConfigWrapper{
				Name:   connectorName,
				Config: configInstance.copy(),
			}
			instances = append(instances, wrap)

		}
	}

	return instances
}

func LoadConfig(path string) (*KafkaConnectorConfig, error) {

	config := KafkaConnectorConfig{}
	err := configor.Load(&config, path)
	return &config, err
}
