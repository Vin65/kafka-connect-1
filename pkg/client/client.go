package client

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/commitdev/kafka-connect-cli/config"
	"github.com/commitdev/kafka-connect-cli/pkg/utils"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
)

//KafkaConnectClient implements functions interacting with kafka connect configurations.
type KafkaConnectClient interface {
	List() ([]string, error)
	Create(*config.KafkaConnectorConfigWrapper) error
	Get(name string) (*KafkaConnectGetResponse, error)
	GetStatus(name string) (*KafkaConnectGetStatusResponse, error)
	Update(name string, config *config.KafkaConnectorConfigInstance) error
	Delete(name string) error
}

type kafkaConnectHTTPClient struct {
	base   *url.URL
	client utils.HttpClient
}

//NewKafkaConnectClient creates a new kafka connect client.
func NewKafkaConnectClient(address string, user string, password string) (KafkaConnectClient, error) {
	client := utils.NewHttpClient(user, password)
	baseURL, err := url.Parse(address)
	if err != nil {
		return nil, err
	}

	return &kafkaConnectHTTPClient{
		base:   baseURL,
		client: client,
	}, nil
}

func (c *kafkaConnectHTTPClient) List() ([]string, error) {
	u, _ := url.Parse("/connectors")
	endpoint := c.base.ResolveReference(u).String()
	req, _ := http.NewRequest("GET", endpoint, nil)
	resp, err := c.client.Do(req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	result := make([]string, 0)
	err = json.NewDecoder(resp.Body).Decode(&result)

	return result, err
}

func (c *kafkaConnectHTTPClient) Create(config *config.KafkaConnectorConfigWrapper) error {
	u, _ := url.Parse("/connectors")
	endpoint := c.base.ResolveReference(u).String()
	payload, _ := json.Marshal(config)
	req, _ := http.NewRequest("POST", endpoint, bytes.NewBuffer(payload))
	resp, err := c.client.Do(req)
	if resp.StatusCode != 200 && resp.StatusCode != 201 {
		message, _ := ioutil.ReadAll(resp.Body)

		log.Printf("Status: %v Message: %s", resp.StatusCode, string(message))
	}

	return err

}

type KafkaConnectGetResponse struct {
	Name   string      `json:"name"`
	Config interface{} `json:"config"`
	Tasks  []struct {
		Connector string `json:"connector"`
		Task      int    `json:"task"`
	} `json:"tasks"`
	Type string `json:"type"`
}

type NotFound struct {
	msg string
}

func NewNotFoundError(msg string) *NotFound {
	return &NotFound{msg: msg}
}

func (e *NotFound) Error() string {
	return e.msg
}

func (c *kafkaConnectHTTPClient) Get(name string) (*KafkaConnectGetResponse, error) {
	u, _ := url.Parse(fmt.Sprintf("/connectors/%s", name))
	endpoint := c.base.ResolveReference(u).String()
	req, _ := http.NewRequest("GET", endpoint, nil)
	resp, err := c.client.Do(req)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode == 404 {
		return nil, NewNotFoundError("This config does not exist.")
	}

	defer resp.Body.Close()

	result := KafkaConnectGetResponse{}
	err = json.NewDecoder(resp.Body).Decode(&result)

	return &result, nil
}

type KafkaConnectGetStatusResponse  struct {
	Name      string `json:"name"`
	Connector struct {
		State    string `json:"state"`
		WorkerID string `json:"worker_id"`
	} `json:"connector"`
	Tasks []struct {
		ID       int    `json:"id"`
		State    string `json:"state"`
		WorkerID string `json:"worker_id"`
		Trace    string `json:"trace,omitempty"`
	} `json:"tasks"`
}

func (c *kafkaConnectHTTPClient) GetStatus(name string) (*KafkaConnectGetStatusResponse, error) {
	u, _ := url.Parse(fmt.Sprintf("/connectors/%s/status", name))
	endpoint := c.base.ResolveReference(u).String()
	req, _ := http.NewRequest("GET", endpoint, nil)
	resp, err := c.client.Do(req)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode == 404 {
		return nil, NewNotFoundError("This config does not exist.")
	}

	defer resp.Body.Close()

	result := KafkaConnectGetStatusResponse{}
	err = json.NewDecoder(resp.Body).Decode(&result)

	return &result, nil
}

func (c *kafkaConnectHTTPClient) Update(name string, config *config.KafkaConnectorConfigInstance) error {
	u, _ := url.Parse(fmt.Sprintf("/connectors/%s/config", name))
	endpoint := c.base.ResolveReference(u).String()
	payload, _ := json.Marshal(config)
	req, _ := http.NewRequest("PUT", endpoint, bytes.NewBuffer(payload))
	resp, err := c.client.Do(req)
	if err != nil {
		return err
	}

	if resp.StatusCode == 404 {
		return NewNotFoundError("This config does not exist.")
	}

	return nil

}

func (c *kafkaConnectHTTPClient) Delete(name string) error {
	u, _ := url.Parse(fmt.Sprintf("/connectors/%s", name))
	endpoint := c.base.ResolveReference(u).String()
	req, _ := http.NewRequest("DELETE", endpoint, nil)
	resp, err := c.client.Do(req)
	if err != nil {
		return err
	}

	if resp.StatusCode == 404 {
		return NewNotFoundError("This config does not exist.")
	}

	return nil
}
