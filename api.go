package shm_bitmap

import "github.com/LSivan/shm_bitmap/internal/config"

type Interface interface {
	Get(id int64)
	Set(id int64)
}

func New(cfg config.Unmarshaler) (Interface, error) {

	err := cfg.Unmarshal(&config.Config)



	return nil, err
}
