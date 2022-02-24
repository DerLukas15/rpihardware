//Package rpihardware determines Raspberry Pi hardware and corresponding information during runtime.
package rpihardware

import "github.com/pkg/errors"

//Version type for hardware version
type Version uint32

//RPiType type for Raspberry Pi Type.
type RPiType uint32

//Possible values for RPiType
const (
	RPiTypeUnknown RPiType = 0
	RPiType1       RPiType = 1
	RPiType2       RPiType = 2
	RPiType4       RPiType = 3
)

//Errors
var (
	ErrHwNotDetected   = errors.New("Hardware undefined")
	ErrUndefinedEndian = errors.New("Endian undefined")
)

//Hardware contians all information about a Raspberry Pi hardware.
type Hardware struct {
	Version       Version //Hardware version
	RPiType       RPiType //Raspberry Pi type
	PhysAddrBase  uint32  //Base for physical memory address
	Desc          string  //Description of the hardware
	OscFreq       uint32  //Frequency of the build in oscillator
	VideocoreBase uint32  //Base for videocore
}

//HardwareList contains all defined hardwares which are used for the check.
var HardwareList = []Hardware{}

func init() {
	HardwareList = append(HardwareList, model4List...)
	HardwareList = append(HardwareList, model2List...)
	HardwareList = append(HardwareList, model1List...)
}

//Check determines the currently used hardware and will return a pointer to the definition.
//Check will return 'ErrHwNotDetected' if no hardware could be determined.
func Check() (*Hardware, error) {
	return detect()
}
