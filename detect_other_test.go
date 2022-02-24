//go:build !arm64
// +build !arm64

package rpihardware

import (
	"strings"
	"testing"

	"github.com/pkg/errors"
)

func TestCPUInfoParser(t *testing.T) {
	line := "model name	: ARMv6-compatible processor rev 7 (v6l)"
	returnedHW, err := parseCPUInfoLine(line)
	if err != nil {
		t.Fatal(err)
	}
	if returnedHW != nil {
		t.Errorf("Expected no hardware but got '%s'", returnedHW.Desc)
	}
	line = "Revision	: 0010"
	returnedHW, err = parseCPUInfoLine(line)
	if err != nil {
		t.Fatal(err)
	}
	if returnedHW == nil {
		t.Errorf("Expected hardware but got none")
	}
	line = "Revision	: a01040"
	returnedHW, err = parseCPUInfoLine(line)
	if err != nil {
		t.Fatal(err)
	}
	if returnedHW == nil {
		t.Errorf("Expected hardware but got none")
	}
	line = "Revision	: c03131"
	returnedHW, err = parseCPUInfoLine(line)
	if err != nil {
		t.Fatal(err)
	}
	if returnedHW == nil {
		t.Errorf("Expected hardware but got none")
	}
}

func TestParserOutOfRange(t *testing.T) {
	line := "Revision	: 100000000"
	_, err := parseCPUInfoLine(line)
	if err == nil {
		t.Errorf("Expected 'value out of range' error")
	}
}

func TestDetect(t *testing.T) {
	//The hardware of the system during the test could work, but it may not be a valid hardware
	_, err := detect()
	if err != nil {
		if !errors.Is(err, ErrHwNotDetected) {
			t.Fatal(err)
		}
	}
}

func TestParseCPUOutput(t *testing.T) {
	f := strings.NewReader(`
processor	: 0
model name	: Random name
BogoMIPS	: 38.40
Features	: half thumb fastmult vfp edsp neon vfpv3 tls vfpv4 idiva idivt vfpd32 lpae evtstrm crc32
CPU implementer	: 0x41
CPU architecture: 6
CPU variant	: 0x0
CPU part	: 0xd03
CPU revision	: c03131
`)
	returnedHW, err := parseCPUInfoOutput(f)
	if err != nil {
		t.Fatal(err)
	}
	if returnedHW == nil {
		t.Errorf("Expected hardware but got none")
	}
}
