package bucket

import (
	"github.com/LSivan/shm_bitmap/internal/config"
	"github.com/LSivan/shm_bitmap/internal/shm"
	"github.com/LSivan/shm_bitmap/internal/util"
	"reflect"
	"unsafe"
)

type shmInfo struct {
	key  string
	id   uintptr
	data []interface{}
}

type Bucket struct {
	shmInfo
	appID uint64
}

func (b *Bucket) GenShmKey(bucketIdx int) (uintptr, error) {
	if bucketIdx >= util.MAX_BUCKET_INDEX {
		return 0, util.ErrInvalidBucketIndex
	}

	shmKey := uintptr(b.appID<<util.BUCKET_SHIFT) + uintptr(bucketIdx)
	return shmKey, nil
}

func (b *Bucket) GetAndCreate(bucketIdx int) (uintptr, error) {
	shmKey, err := b.GenShmKey(bucketIdx)
	if err != nil {
		return shmKey, err
	}
	shmId, err := shm.ShmGet(shmKey)
	if err != nil {
		shmId, err = shm.ShmCreate(shmKey, config.Config.BucketIdCnt*uint32(config.Config.BitSize))
		if err != nil {
			return shmId, err
		}

		shmAddr, err := shm.ShmAt(shmId)
		if err != nil {
			return shmId, err
		}
		slice := (*reflect.SliceHeader)(unsafe.Pointer(&b.data))
		slice.Len = int(config.Config.BucketIdCnt)
		slice.Cap = int(config.Config.BucketIdCnt)
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
		slice.Len = int(config.Config.BucketIdCnt)
		slice.Cap = int(config.Config.BucketIdCnt)
		slice.Data = shmAddr
		return shmId, nil
	}
}

func (b *Bucket) SetBit(id int64, bit uint64) error {
	idx := id % int64(config.Config.BucketIdCnt)
	if len(b.data) <= 0 {
		return util.ErrShmNotMap
	}
	if idx >= int64(len(b.data)) {
		return util.ErrInvalidBucketIndex
	}
	b.data[idx] = bit
	return nil
}

func (b *Bucket) GetBit(id int64) (interface{}, error) {
	idx := id % int64(config.Config.BucketIdCnt)
	if len(b.data) <= 0 {
		return nil, util.ErrShmNotMap
	}
	if idx >= int64(len(b.data)) {
		return nil, util.ErrInvalidBucketIndex
	}
	return b.data[idx], nil
}
