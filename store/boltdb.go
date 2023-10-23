package store

import (
	"autoapi/graph/model"
	"autoapi/util"
	"encoding/json"

	"github.com/boltdb/bolt"
)

const BucketName = "ServerConfigs"

type BoltStore struct {
	dbPath string
}

func NewBoltStore(dbPath string) ServerConfigStorer {
	return &BoltStore{dbPath: dbPath}
}

func (s *BoltStore) CreateService(config *model.ServerConfig) (*model.ServerConfig, error) {
	db, err := bolt.Open(s.dbPath, 0600, nil)
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
		if err == nil {
			err = util.CreateKustomizations(config)
			if err != nil {
				return err
			}
		} else {
			return err
		}

		return b.Put([]byte(config.ID), configBytes)
	})

	if err != nil {
		return nil, err
	}

	return config, nil
}

func (s *BoltStore) GetAllConfigs() ([]*model.ServerConfig, error) {
	db, err := bolt.Open(s.dbPath, 0600, nil)
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

func (s *BoltStore) DeleteConfig(id string) error {
	db, err := bolt.Open(s.dbPath, 0600, nil)
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

func (s *BoltStore) DeleteAllConfigs() error {
	db, err := bolt.Open(s.dbPath, 0600, nil)
	if err != nil {
		return err
	}
	defer db.Close()

	return db.Update(func(tx *bolt.Tx) error {
		// Delete the entire bucket which will effectively remove all the configurations.
		return tx.DeleteBucket([]byte(BucketName))
	})
}
