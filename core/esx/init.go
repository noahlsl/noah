package esx

import (
	"fmt"

	es "github.com/elastic/go-elasticsearch/v7"
)

func (c *Cfg) NewClient() *es.Client {

	var urls []string
	for _, addr := range c.Address {
		url := fmt.Sprintf("http://%s", addr)
		if c.TLS != 0 {
			url = fmt.Sprintf("https://%s", addr)
		}
		urls = append(urls, url)
	}

	client, err := es.NewClient(es.Config{
		Addresses: urls,
		Username:  c.Username,
		Password:  c.Password,
	})
	if err != nil {
		panic(err)
	}

	return client
}
