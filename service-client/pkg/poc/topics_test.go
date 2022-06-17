package poc

import (
	"errors"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"net/http"
	"testing"

	"log"
)

func TestTopics(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Topics CRUD Suite")
}

//var apiBaseURL = "http://service-api.localtest.me"
var apiBaseURL = "http://localhost:8080/v1"
var tenant = "tenant1"

var _ = Describe("Topics", func() {
	var client *Client
	BeforeEach(func() {
		By("Initializing the client with a base URL")
		client = &Client{
			Client: http.Client{
				// tune settings
			},
			BaseURL: apiBaseURL,
			Tenant:  tenant,
		}
		Expect(client.BaseURL).To(Equal(apiBaseURL))
	})
	It("Test CreateTopic", func() {
		topic := Topic{Name: "test1", Status: "CREATED", Partitions: 4}
		log.Printf("\nTopic: %v", topic)
		err := client.CreateTopic(topic)
		Expect(err).NotTo(HaveOccurred())
	})
	It("Test GetTopics", func() {
		topics, err := client.GetTopics()
		Expect(err).NotTo(HaveOccurred())
		log.Printf("\nTopics: %v", topics)
	})
	It("Test UpdateTopic", func() {
		topic := Topic{Name: "test1", Status: "CREATED", Partitions: 5}
		err := client.UpdateTopic(topic)
		Expect(err).NotTo(HaveOccurred())
	})
	It("Test GetTopic", func() {
		topic, err := client.GetTopic("test1")
		Expect(err).NotTo(HaveOccurred())
		log.Printf("\nTopic: %v", topic)
	})
	It("Test DeleteTopic", func() {
		err := client.DeleteTopic("test1")
		Expect(err).NotTo(HaveOccurred())
	})
	It("Test DeleteTopic unexisting", func() {
		err := client.DeleteTopic("test_unexisting")
		Expect(err).To(Equal(errors.New("unexpected status: 404 Not Found")))
	})
	It("Test UpdateTopic unexisting", func() {
		topic := Topic{Name: "test_unexisting", Status: "CREATED", Partitions: 5}
		err := client.UpdateTopic(topic)
		Expect(err).To(Equal(errors.New("unexpected status: 404 Not Found")))
	})
	It("Test GetTopic unexisting", func() {
		_, err := client.GetTopic("test_unexisting")
		Expect(err).To(Equal(errors.New("unexpected status: 404 Not Found")))
	})

})
