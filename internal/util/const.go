package util

import "errors"

const (
	BUCKET_SHIFT = 16
	MAX_BUCKET_INDEX = 1 << 16
)

var (
	ErrInvalidBucketIndex = errors.New("err invalid bucket index")
	ErrShmNotMap = errors.New("err shm not map")
	ErrGenShmKey = errors.New("failed to gen ken")
)
