package config

// Cfg is config of Bucket. And Bucket is a segment of shared memory
// BucketCnt is (IDEnd - IDBegin) / BucketIdCnt
// Size of bucket will be BucketSize = BucketIdCnt * ByteSize
// Size of shared memory will be like SHR = BucketSize * BucketCnt
type Cfg struct {
	//BucketCnt      int    // numbers of bucket
	IDBegin, IDEnd int64  // use to cal bucket index of ID
	BucketIdCnt    uint32 // numbers of id in every bucket
	ByteSize       uint   // how many byte. 1 byte = 8 bit
}
