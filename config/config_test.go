package config

import (
	"encoding/json"
	"io/ioutil"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestKafkaConnectorConfigInstance_MergeAll(t *testing.T) {
	config, _ := LoadConfig("../example/kafka-connect.yaml")

	expectedValues, _ := ioutil.ReadFile("../test-files/test_merge_all_output.json")
	var expectedResults []KafkaConnectorConfigWrapper
	json.Unmarshal([]byte(expectedValues), &expectedResults)

	result := config.MergeAll()
	r1, _ := json.Marshal(result[0])
	e1, _ := json.Marshal(expectedResults[0])
	r2, _ := json.Marshal(result[1])
	e2, _ := json.Marshal(expectedResults[1])

	assert.Equal(t, r1, e1)
	assert.Equal(t, r2, e2)
	assert.Equal(t, 4, len(result))

}
