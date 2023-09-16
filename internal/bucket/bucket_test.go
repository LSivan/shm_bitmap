package bucket

import (
	"github.com/LSivan/shm_bitmap/internal/config"
	"testing"
)

func TestNewAppBucket(t *testing.T) {
	type args struct {
		appID uint32
		cfg   config.Cfg
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "test",
			args: args{
				appID: 1,
				cfg: config.Cfg{
					IDBegin:     0,
					IDEnd:       50,
					BucketIdCnt: 20,
					BitSize:     4,
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewAppBucket(tt.args.appID, tt.args.cfg)
			if err != nil {
				t.Errorf("NewAppBucket() error = %v", err)
				return
			}

			for i := 15; i < 25; i++ {
				bit := 100 + i
				err = got.SetBit(int64(i), int64(bit))
				if err != nil {
					t.Errorf("got.SetBit(%d,%d) error = %v", i, bit, err)
					return
				}

				t.Logf("%+v \n", got.buckets[got.BucketIdx(int64(i))].data)

				v, e := got.GetBit(int64(i))
				t.Logf("b.GetBit(%d) got = %v, err = %v \n", i, v, e)
			}

		})
	}
}
