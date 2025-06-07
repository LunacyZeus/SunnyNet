//go:build !linux

package SunnyProxy

func bindDevice(fd uintptr, ifceName string) error {
	return nil
}

func setMark(fd uintptr, mark int) error {
	return nil
}
