package bucket

import "testing"

func TestBucket_GetAndCreate(t *testing.T) {
	type fields struct {
		shmInfo shmInfo
		appID   uint64
	}
	type args struct {
		bucketIdx int
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
	}{
		{
			name:    "test",
			fields:  fields{},
			args:    args{1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := &Bucket{
				shmInfo: tt.fields.shmInfo,
				appID:   tt.fields.appID,
			}
			got, err := b.GetAndCreate(tt.args.bucketIdx)
			t.Logf("GetAndCreate() got = %v, err = %v", got, err)
		})
	}
}
