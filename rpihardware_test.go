package rpihardware

import (
	"testing"

	"github.com/pkg/errors"
)

func TestCheck(t *testing.T) {
	//The hardware of the system during the test could work, but it may not be a valid hardware
	_, err := Check()
	if err != nil {
		if !errors.Is(err, ErrHwNotDetected) {
			t.Fatal(err)
		}
	}
}
