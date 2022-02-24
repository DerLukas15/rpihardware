package rpihardware

import (
	"encoding/binary"
	"os"
	"unsafe"

	"github.com/pkg/errors"
)

func detect() (*Hardware, error) {
	fcontent, err := os.ReadFile("/proc/device-tree/system/linux,revision")
	if err != nil {
		return nil, errors.Wrap(err, "detect arm64")
	}
	var rev uint32
	var i = 0x0100
	ptr := unsafe.Pointer(&i)
	if 0x01 == *(*byte)(ptr) {
		rev = binary.BigEndian.Uint32(fcontent)
	} else if 0x00 == *(*byte)(ptr) {
		rev = binary.LittleEndian.Uint32(fcontent)
	} else {
		return nil, errors.Wrap(ErrUndefinedEndian, "detect arm64")
	}
	for curCounter, curEntry := range HardwareList {
		if rev == uint32(curEntry.Version) {
			res := &HardwareList[curCounter]
			return res, nil
		}
	}
	return nil, errors.Wrap(ErrHwNotDetected, "detect arm64")
}
