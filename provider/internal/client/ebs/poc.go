package poc

import (
	"github.com/scalvetr/provider-poc/service-client/pkg/poc"
	"net/http"
)

type POCService struct {
	POCClient poc.Client
}

func NewService(tenant string, baseUrl string, _ []byte) (POCService, error) {
	return POCService{
		POCClient: poc.Client{
			Tenant:  tenant,
			BaseURL: baseUrl,
			Client:  http.Client{},
		},
	}, nil
}

func (c *POCService) GetTopics() ([]poc.Topic, error) {
	return c.POCClient.GetTopics()
}

func (c *POCService) GetTopic(topicName string) (*poc.Topic, error) {
	return c.POCClient.GetTopic(topicName)
}

func (c *POCService) CreateTopic(t poc.Topic) error {
	return c.POCClient.CreateTopic(t)
}

func (c *POCService) UpdateTopic(t poc.Topic) error {
	return c.POCClient.UpdateTopic(t)
}

func (c *POCService) DeleteTopic(topicName string) error {
	return c.POCClient.DeleteTopic(topicName)
}
