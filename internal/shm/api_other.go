//go:build !linux

package shm

// TODO support other platform
// Currently, only Linux is supported

type ShmIdDs struct {
	IpcPerm struct {
		key     uint32
		uid     uint32
		gid     uint32
		cuid    uint32
		cgid    uint32
		mode    uint32
		pad1    uint16
		seq     uint16
		pad2    uint16
		unused1 uint
		unused2 uint
	}
	ShmSegsz   uint32
	ShmAtime   uint64
	ShmDtime   uint64
	ShmCtime   uint64
	ShmCpid    uint32
	ShmLpid    uint32
	ShmNattch  uint16
	ShmUnused  uint16
	ShmUnused2 uintptr
	ShmUnused3 uintptr
}

func ShmCreate(shmKey uintptr, shmSize uint32) (uintptr, error) {
	panic("implement me")
	return 0, nil
}

func ShmCtrl(shmid uintptr, ds *ShmIdDs) error {
	panic("implement me")
	return nil
}

func ShmAt(shmid uintptr) (uintptr, error) {
	panic("implement me")
	return 0, nil
}

func ShmDt(shmaddr uintptr) error {
	panic("implement me")
	return nil
}

func ShmGet(shmKey uintptr) (uintptr, error) {
	panic("implement me")
	return 0, nil
}
