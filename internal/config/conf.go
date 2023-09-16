package config

// Cfg is config of Bucket. And Bucket is a segment of shared memory
// Size of bucket will be BucketSize = BucketIdCnt * BitSize
// Size of shared memory will be like SHR = BucketSize * BucketCnt
type Cfg struct {
	BucketCnt   int    // numbers of bucket
	BucketIdCnt uint32 // numbers of id in every bucket
	BitSize     uint   // sizeof(bit). maybe 1,2,4,8
}
