package repos

import (
	"fmt"
	"io"
	"log"
	"os"

	envservice "github.com/kyromoto/go-ddns-broker/src/services/env-service"
	"gopkg.in/yaml.v3"
)

type ClientRepository struct {
	Clients []Client `yaml:""`
}

type Client struct {
	Username  string `yaml:""`
	Password  string `yaml:""`
	Providers []Provider
	Actions   []Action
}

type Provider struct {
	Name string
}

type Action struct {
	Name     string
	Provider string
}

func OpenClientRepo() (*ClientRepository, error) {
	filename := envservice.ClientRepoDataFilename()

	log.Printf("repo filename: %v", filename)

	data, err := os.OpenFile(filename, os.O_RDONLY|os.O_CREATE, 0666)

	if err != nil {
		return nil, err
	}

	defer data.Close()

	bytes, err := io.ReadAll(data)

	if err != nil {
		return nil, err
	}

	repo := ClientRepository{}

	err = yaml.Unmarshal(bytes, &repo)

	if err != nil {
		return nil, err
	}

	return &repo, nil
}

func GetClientdByUsername(username string) (*Client, error) {

	repo, err := OpenClientRepo()

	if err != nil {
		return nil, err
	}

	for _, client := range repo.Clients {
		if client.Username != username {
			continue
		}

		return &client, nil
	}

	return nil, fmt.Errorf("client with username %v not found", username)
}
