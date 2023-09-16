package util

import "errors"

const (
	BUCKET_SHIFT      = 16
	MAX_BUCKET_INDEX  = 1 << 16
	MAX_BUCKET_ID_CNT = 10000000
)

var (
	ErrInvalidBucketIndex = errors.New("err invalid bucket index")
	ErrShmNotMap          = errors.New("err shm not map")
	ErrInvalidIDRange     = errors.New("err invalid id range")
	ErrInvalidBucketIDCnt = errors.New("err invalid bucket id cnt")
	ErrAppIDNotFound      = errors.New("err app not found")
	ErrInvalidInitCfg     = errors.New("err invalid init config")
)
