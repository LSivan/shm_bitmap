package shm_bitmap

import (
	"github.com/LSivan/shm_bitmap/internal/bucket"
	"github.com/LSivan/shm_bitmap/internal/config"
	"github.com/LSivan/shm_bitmap/internal/util"
)

type BitMapApi interface {
	Set(appID uint32, id, bit int64) error
	Get(appID uint32, id int64) (int64, error)
	BatchGet(appID uint32, ids []int64) (map[int64]int64, error)
}

type mgr struct {
	appBuckets map[uint32]*bucket.AppBucket
}

func (m *mgr) getAppBucket(appID uint32) (*bucket.AppBucket, error) {
	ab, ok := m.appBuckets[appID]
	if !ok || ab == nil {
		return ab, util.ErrAppIDNotFound
	}
	return ab, nil
}

func (m *mgr) Set(appID uint32, id, bit int64) error {
	ab, err := m.getAppBucket(appID)
	if err != nil {
		return err
	}
	return ab.SetBit(id, bit)
}

func (m *mgr) Get(appID uint32, id int64) (int64, error) {
	ab, err := m.getAppBucket(appID)
	if err != nil {
		return 0, err
	}

	return ab.GetBit(id)
}

func (m *mgr) BatchGet(appID uint32, ids []int64) (map[int64]int64, error) {
	ab, err := m.getAppBucket(appID)
	if err != nil {
		return map[int64]int64{}, err
	}

	res := make(map[int64]int64, len(ids))

	for _, id := range ids {
		bit, err := ab.GetBit(id)
		if err != nil {
			return res, err
		}
		res[id] = bit
	}
	return res, nil
}

func New(appBucketCfg map[uint32]config.Cfg) (BitMapApi, error) {
	if len(appBucketCfg) <= 0 {
		return nil, util.ErrInvalidInitCfg
	}

	m := mgr{
		appBuckets: map[uint32]*bucket.AppBucket{},
	}

	for appID, conf := range appBucketCfg {
		ab, err := bucket.NewAppBucket(appID, conf)
		if err != nil {
			return nil, err
		}
		m.appBuckets[appID] = ab
	}

	return &m, nil
}
