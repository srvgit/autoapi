package util

import (
	"encoding/json"
	"fmt"
	"strings"
	"text/template"
	"time"

	"github.com/google/uuid"
	bolt "go.etcd.io/bbolt"
)

type ServerConfig struct {
	ID                          string `json:"id,omitempty"`
	GraphPackagePath            string `json:"graphPackagePath"`
	GinPackagePath              string `json:"ginPackagePath"`
	GQLGenHandlerPackagePath    string `json:"gqlgenHandlerPackagePath"`
	GQLGenPlaygroundPackagePath string `json:"gqlgenPlaygroundPackagePath"`
	PlaygroundEndpoint          string `json:"playgroundEndpoint"`
	GraphQLEndpoint             string `json:"graphQLEndpoint"`
	RootEndpoint                string `json:"rootEndpoint"`
}

const bucketName = "configurations"

func storeConfig(db *bolt.DB, config *ServerConfig) error {
	return db.Update(func(tx *bolt.Tx) error {
		b, err := tx.CreateBucketIfNotExists([]byte(bucketName))
		if err != nil {
			return err
		}
		id := fmt.Sprintf("%d", time.Now().UnixNano())
		config.ID = id
		encodedConfig, err := json.Marshal(config)
		if err != nil {
			return err
		}
		return b.Put([]byte(id), encodedConfig)
	})
}

func GenerateCodeFromConfig(db *bolt.DB, configID string) (string, error) {
	var config ServerConfig

	err := db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(bucketName))
		if b == nil {
			return fmt.Errorf("bucket %s not found", bucketName)
		}
		data := b.Get([]byte(configID))
		return json.Unmarshal(data, &config)
	})
	if err != nil {
		return "", err
	}

	tmpl, err := template.ParseFiles("template/server.go.tmpl")
	if err != nil {
		return "", err
	}

	var result strings.Builder
	err = tmpl.Execute(&result, config)
	if err != nil {
		return "", err
	}

	return result.String(), nil
}

func GenerateUUID() string {
	return uuid.New().String()
}
