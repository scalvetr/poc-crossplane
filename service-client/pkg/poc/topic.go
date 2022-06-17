package poc

import (
	"bytes"
	"encoding/json"
	"errors"
	"io/ioutil"
	"log"
)

type Topic struct {
	Name       string `json:"name"`
	Partitions int    `json:"partitions"`
	Status     string `json:"status,omitempty"`
}

type Error struct {
	Error   string `json:"error"`
	Message string `json:"message"`
	Status  int    `json:"status"`
}

func (c *Client) GetTopics() ([]Topic, error) {
	// GET /topics
	res, err := c.Get(c.BaseURL + "/topics")
	if err != nil {
		return nil, err
	}

	log.Println(res)
	// Read the body
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	if res.StatusCode != 200 {
		var data_obj Error
		var errorBody []byte
		res.Body.Read(errorBody)
		err = json.Unmarshal(errorBody, &data_obj)
		if err != nil {
			return nil, errors.New("unexpected status: " + res.Status)
		}
		return nil, errors.New(data_obj.Error + ": " + data_obj.Message)
	}

	// unmarshall the response
	var data_obj []Topic
	err = json.Unmarshal(body, &data_obj)
	if err != nil {
		return nil, err
	}

	return data_obj, nil
}

func (c *Client) GetTopic(topicName string) (*Topic, error) {
	// GET /topics
	res, err := c.Get(c.BaseURL + "/topics/" + topicName)
	if err != nil {
		return nil, err
	}

	log.Println(res)
	// Read the body
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	if res.StatusCode != 200 {
		var data_obj Error
		var errorBody []byte
		res.Body.Read(errorBody)
		err = json.Unmarshal(errorBody, &data_obj)
		if err != nil {
			return nil, errors.New("unexpected status: " + res.Status)
		}
		return nil, errors.New(data_obj.Error + ": " + data_obj.Message)
	}

	// unmarshall the response
	var data_obj *Topic
	err = json.Unmarshal(body, &data_obj)
	if err != nil {
		return nil, err
	}

	return data_obj, nil
}

func (c *Client) CreateTopic(t Topic) error {
	body, err := json.Marshal(t)
	if err != nil {
		return err
	}
	postBody := bytes.NewBuffer(body)

	// POST /topics
	res, err := c.Post(c.BaseURL+"/topics", postBody)
	log.Println(res)
	if err != nil {
		return err
	}
	if res.StatusCode != 201 {
		var data_obj Error
		var errorBody []byte
		res.Body.Read(errorBody)
		err = json.Unmarshal(errorBody, &data_obj)
		if err != nil {
			return errors.New("unexpected status: " + res.Status)
		}
		return errors.New(data_obj.Error + ": " + data_obj.Message)
	}
	return nil
}

func (c *Client) UpdateTopic(t Topic) error {
	body, err := json.Marshal(t)
	if err != nil {
		return err
	}
	postBody := bytes.NewBuffer(body)

	// POST /topics
	res, err := c.Put(c.BaseURL+"/topics/"+t.Name, postBody)
	log.Println(res)
	if err != nil {
		return err
	}
	if res.StatusCode != 204 {
		var data_obj Error
		var errorBody []byte
		res.Body.Read(errorBody)
		err = json.Unmarshal(errorBody, &data_obj)
		if err != nil {
			return errors.New("unexpected status: " + res.Status)
		}
		return errors.New(data_obj.Error + ": " + data_obj.Message)
	}
	return nil
}

func (c *Client) DeleteTopic(topicName string) error {
	// DELETE /topics
	res, err := c.Delete(c.BaseURL + "/topics/" + topicName)
	log.Println(res)
	if err != nil {
		return err
	}
	if res.StatusCode != 204 {
		var data_obj Error
		var errorBody []byte
		res.Body.Read(errorBody)
		err = json.Unmarshal(errorBody, &data_obj)
		if err != nil {
			return errors.New("unexpected status: " + res.Status)
		}
		return errors.New(data_obj.Error + ": " + data_obj.Message)
	}
	return nil
}
