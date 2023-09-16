//go:build linux

package shm

import (
	"syscall"
	"unsafe"
)

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

const (
	IPC_RMID  = 0
	IPC_SET   = 1
	IPC_STAT  = 2
	IPC_INFO  = 3
	IPC_CREAT = 00001000
	IPC_EXCL  = 00002000
)

func ShmCreate(shmKey uintptr, shmSize uint32) (uintptr, error) {
	shmid, _, err := syscall.Syscall(syscall.SYS_SHMGET, shmKey, uintptr(shmSize), IPC_CREAT|IPC_EXCL|0666)
	if err != 0 {
		return shmid, err
	}
	return shmid, nil
}

func ShmCtrl(shmid uintptr, ds *ShmIdDs) error {

	_, _, err := syscall.Syscall(syscall.SYS_SHMCTL, shmid, IPC_STAT, uintptr(unsafe.Pointer(ds)))
	if err != 0 {
		return err
	}
	return nil
}

func ShmAt(shmid uintptr) (uintptr, error) {

	shmaddr, _, err := syscall.Syscall(syscall.SYS_SHMAT, shmid, 0, 0)
	if err != 0 {
		return 0, err
	}
	return shmaddr, nil
}

func ShmDt(shmaddr uintptr) error {
	_, _, err := syscall.Syscall(syscall.SYS_SHMDT, shmaddr, 0, 0)
	if err != 0 {
		return err
	}
	return nil
}

func ShmGet(shmKey uintptr) (uintptr, error) {

	shmid, _, err := syscall.Syscall(syscall.SYS_SHMGET, shmKey, 0, 0666)
	if err != 0 {
		return shmid, err
	}
	return shmid, nil
}

// reference:
//  https://my.oschina.net/tenghui0425/blog/1118882
