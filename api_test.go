package shm_bitmap

import (
	"github.com/LSivan/shm_bitmap/internal/config"
	"testing"
)

func TestNew(t *testing.T) {
	var appID uint32 = 2
	type args struct {
		appBucketCfg map[uint32]config.Cfg
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "test",
			args: args{
				appBucketCfg: map[uint32]config.Cfg{
					appID: {
						BucketCnt:   5,
						BucketIdCnt: 20,
						BitSize:     4,
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := New(tt.args.appBucketCfg)
			if err != nil {
				t.Errorf("New() error = %v", err)
				return
			}

			for i := int64(1); i <= 100; i++ {
				err = got.Set(appID, i, i+10000)
				if err != nil {
					t.Errorf("got.Set(%d, %d, %d) error = %v", appID, i, i+10000, err)
					return
				}

				var res int64
				res, err = got.Get(appID, i)
				if err != nil {
					t.Errorf("got.Get(%d, %d) error = %v", appID, i, err)
					return
				}

				t.Logf("got.Get(%d, %d) res=%d \n", appID, i, res)
			}

		})
	}
}
