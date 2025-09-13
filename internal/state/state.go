package state

import "github.com/bmccarson/gator/internal/config"

type DataStore struct {
	CurrentConfig config.Config
}

func Init(cfg config.Config) DataStore {
	return DataStore{
		CurrentConfig: cfg,
	}
}
