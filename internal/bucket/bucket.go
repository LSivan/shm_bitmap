package bucket

import (
	"github.com/LSivan/shm_bitmap/internal/config"
	"github.com/LSivan/shm_bitmap/internal/shm"
	"github.com/LSivan/shm_bitmap/internal/util"
	"reflect"
	"unsafe"
)

type shmInfo struct {
	//key  string
	id   uintptr
	data []int64
}

// Bucket
// manage a segment of shared memory.
type Bucket struct {
	shmInfo
	master *AppBucket
}

func (b *Bucket) GenShmKey(appID uint32, bucketIdx int) (uintptr, error) {
	if bucketIdx >= util.MAX_BUCKET_INDEX {
		return 0, util.ErrInvalidBucketIndex
	}

	shmKey := uintptr(appID<<util.BUCKET_SHIFT) + uintptr(bucketIdx)
	return shmKey, nil
}

func (b *Bucket) GetAndCreate(appID uint32, bucketIdx int) (uintptr, error) {
	shmKey, err := b.GenShmKey(appID, bucketIdx)
	if err != nil {
		return shmKey, err
	}
	shmId, err := shm.ShmGet(shmKey)
	if err != nil {
		shmId, err = shm.ShmCreate(shmKey, b.master.cfg.BucketIdCnt*uint32(b.master.cfg.BitSize))
		if err != nil {
			return shmId, err
		}

		shmAddr, err := shm.ShmAt(shmId)
		if err != nil {
			return shmId, err
		}
		slice := (*reflect.SliceHeader)(unsafe.Pointer(&b.data))
		slice.Len = int(b.master.cfg.BucketIdCnt)
		slice.Cap = int(b.master.cfg.BucketIdCnt)
		slice.Data = shmAddr
		return shmId, nil
	} else {
		// TODO
		//var shmIDs shm.ShmIdDs
		//err = shm.ShmCtrl(shmId, &shmIDs)
		//if err != nil {
		//	return shmId, err
		//}
		//if shmIDs.ShmSegsz !=

		shmAddr, err := shm.ShmAt(shmId)
		if err != nil {
			return shmId, err
		}
		slice := (*reflect.SliceHeader)(unsafe.Pointer(&b.data))
		slice.Len = int(b.master.cfg.BucketIdCnt)
		slice.Cap = int(b.master.cfg.BucketIdCnt)
		slice.Data = shmAddr
		return shmId, nil
	}
}

func (b *Bucket) SetBit(id, bit int64) error {
	idx := id % int64(b.master.cfg.BucketIdCnt)
	if len(b.data) <= 0 {
		return util.ErrShmNotMap
	}
	if idx >= int64(len(b.data)) {
		return util.ErrInvalidBucketIndex
	}
	b.data[idx] = bit
	return nil
}

func (b *Bucket) GetBit(id int64) (int64, error) {
	idx := id % int64(b.master.cfg.BucketIdCnt)
	if len(b.data) <= 0 {
		return 0, util.ErrShmNotMap
	}
	if idx >= int64(len(b.data)) {
		return 0, util.ErrInvalidBucketIndex
	}
	return b.data[idx], nil
}

type AppBucket struct {
	appId   uint32
	buckets []*Bucket
	cfg     config.Cfg
}

func (ab *AppBucket) BucketIdx(id int64) int {
	return int(id / int64(ab.cfg.BucketIdCnt))
}

func (ab *AppBucket) GetBit(id int64) (int64, error) {
	bucketIdx := ab.BucketIdx(id)

	if bucketIdx >= len(ab.buckets) {
		return 0, util.ErrInvalidBucketIndex
	}
	bucket := ab.buckets[bucketIdx]
	return bucket.GetBit(id)
}

func (ab *AppBucket) SetBit(id, bit int64) error {
	bucketIdx := ab.BucketIdx(id)

	if bucketIdx >= len(ab.buckets) {
		return util.ErrInvalidBucketIndex
	}

	bucket := ab.buckets[bucketIdx]
	return bucket.SetBit(id, bit)
}

func NewAppBucket(appID uint32, cfg config.Cfg) (*AppBucket, error) {
	ab := AppBucket{
		appId: appID,
		cfg:   cfg,
	}

	if ab.cfg.BucketCnt <= 0 {
		return nil, util.ErrInvalidBucketCnt
	}
	if ab.cfg.BucketIdCnt <= 0 || ab.cfg.BucketIdCnt > util.MAX_BUCKET_ID_CNT {
		return nil, util.ErrInvalidBucketIDCnt
	}

	ab.buckets = make([]*Bucket, ab.cfg.BucketCnt)

	for i := 0; i < ab.cfg.BucketCnt; i++ {
		var bucket Bucket

		shmID, err := bucket.GetAndCreate(appID, i)
		if err != nil {
			return nil, err
		}
		bucket.id = shmID

		ab.buckets[i] = &bucket
	}

	return &ab, nil
}
