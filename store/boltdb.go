package store

import (
	"autoapi/graph/model"
	"encoding/json"

	"github.com/boltdb/bolt"
)

const BucketName = "ServerConfigs"

func StoreConfigInDB(config *model.ServerConfig) (*model.ServerConfig, error) {
	db, err := bolt.Open("configurations.db", 0600, nil)
	if err != nil {
		return nil, err
	}
	defer db.Close()

	err = db.Update(func(tx *bolt.Tx) error {
		b, err := tx.CreateBucketIfNotExists([]byte(BucketName))
		if err != nil {
			return err
		}

		configBytes, err := json.Marshal(config)
		if err != nil {
			return err
		}

		return b.Put([]byte(config.ID), configBytes)
	})

	if err != nil {
		return nil, err
	}

	return config, nil
}

func GetAllConfigsFromDB() ([]*model.ServerConfig, error) {
	db, err := bolt.Open("configurations.db", 0600, nil)
	if err != nil {
		return nil, err
	}
	defer db.Close()

	var configs []*model.ServerConfig

	err = db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(BucketName))
		if b == nil {
			return nil // No bucket means no data yet.
		}

		return b.ForEach(func(k, v []byte) error {
			var config model.ServerConfig
			err := json.Unmarshal(v, &config)
			if err != nil {
				return err
			}
			configs = append(configs, &config)
			return nil
		})
	})

	if err != nil {
		return nil, err
	}

	return configs, nil
}

func DeleteConfigFromDB(id string) error {
	db, err := bolt.Open("configurations.db", 0600, nil)
	if err != nil {
		return err
	}
	defer db.Close()

	return db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(BucketName))
		if b == nil {
			return nil // No bucket means no data yet.
		}
		return b.Delete([]byte(id))
	})
}

func DeleteAllConfigsFromDB() error {
	db, err := bolt.Open("configurations.db", 0600, nil)
	if err != nil {
		return err
	}
	defer db.Close()

	return db.Update(func(tx *bolt.Tx) error {
		// Delete the entire bucket which will effectively remove all the configurations.
		return tx.DeleteBucket([]byte(BucketName))
	})
}
