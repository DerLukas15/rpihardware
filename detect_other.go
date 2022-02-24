//go:build !arm64
// +build !arm64

package rpihardware

import (
	"bufio"
	"io"
	"os"
	"strconv"
	"strings"

	"github.com/pkg/errors"
)

const (
	manufacturerMask uint32 = 0xf << 16
	warrantyMask            = 0x3 << 24
)

func detect() (*Hardware, error) {
	f, err := os.Open("/proc/cpuinfo")
	if err != nil {
		return nil, errors.Wrap(err, "detect other")
	}
	defer f.Close()
	return parseCPUInfoOutput(f)
}

func parseCPUInfoOutput(f io.Reader) (*Hardware, error) {
	reader := bufio.NewReader(f)
	for {
		line, err := reader.ReadString('\n')
		if err != nil {
			if errors.Is(err, io.EOF) {
				break
			}
			return nil, errors.Wrap(err, "detect other")
		}
		detectedHardware, err := parseCPUInfoLine(line)
		if err != nil {
			return nil, err
		}
		if detectedHardware != nil {
			return detectedHardware, nil
		}
	}
	return nil, errors.Wrap(ErrHwNotDetected, "detect other")
}

func parseCPUInfoLine(line string) (*Hardware, error) {
	checkStrings := []string{
		"Revision",
		"CPU revision",
	}
	for _, curCheckString := range checkStrings {
		if strings.Contains(line, curCheckString) {
			sections := strings.Split(line, ":")
			revString := strings.TrimSpace(sections[1])
			rev, err := strconv.ParseUint(revString, 16, 32)
			if err != nil {
				return nil, errors.Wrap(err, "detect other")
			}
			for curCounter, curEntry := range HardwareList {
				hwVer := uint32(curEntry.Version) & ^(warrantyMask | manufacturerMask)
				checkRev := uint32(rev) &^ (warrantyMask | manufacturerMask)
				if checkRev == hwVer {
					res := &HardwareList[curCounter]
					return res, nil
				}
			}
		}
	}
	return nil, nil
}
