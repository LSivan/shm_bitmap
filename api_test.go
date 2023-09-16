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
					// Store 30 ID's bit in per segment
					// Total count of ID is 199=200-1
					// So need 7 segments (199/30 ≈ 6.63 < 7)
					// size of every segment is 240=30*8
					appID: {
						IDBegin:     1,
						IDEnd:       200,
						BucketIdCnt: 30,
						BitSize:     8,
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

			ids := []int64{1, 15, 29, 33, 89}
			mapBit, err := got.BatchGet(appID, ids)
			t.Logf("got.BatchGet((%d, %v) res=%d err=%v \n", appID, ids, mapBit, err)
		})
	}
}

/**
$ ipcs -m

------ Shared Memory Segments --------
key        shmid      owner      perms      bytes      nattch     status
0x00020000 2490379    root       666        240        0
0x00020001 2523148    root       666        240        0
0x00020002 2555917    root       666        240        0
0x00020003 2588686    root       666        240        0
0x00020004 2621455    root       666        240        0
0x00020005 2654224    root       666        240        0
0x00020006 2686993    root       666        240        0

*/
