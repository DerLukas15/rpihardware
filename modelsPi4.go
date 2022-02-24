package rpihardware

const (
	physAddrBaseRPi4 uint32 = 0xfe000000
	oscFreqRPi4      uint32 = 54000000
)

var (
	model4List = []Hardware{
		//
		// Raspberry Pi 400
		//
		{
			Version:       0xc03130,
			RPiType:       RPiType4,
			PhysAddrBase:  physAddrBaseRPi4,
			VideocoreBase: videocoreBaseRPi2,
			OscFreq:       oscFreqRPi4,
			Desc:          "Pi 400 - 4GB v1.0",
		},
		{
			Version:       0xc03131,
			RPiType:       RPiType4,
			PhysAddrBase:  physAddrBaseRPi4,
			VideocoreBase: videocoreBaseRPi2,
			OscFreq:       oscFreqRPi4,
			Desc:          "Pi 400 - 4GB v1.1",
		},
		//
		// Raspberry Pi 4
		//
		{
			Version:       0xA03111,
			RPiType:       RPiType4,
			PhysAddrBase:  physAddrBaseRPi4,
			VideocoreBase: videocoreBaseRPi2,
			OscFreq:       oscFreqRPi4,
			Desc:          "Pi 4 Model B - 1GB v1.1",
		},
		{
			Version:       0xB03111,
			RPiType:       RPiType4,
			PhysAddrBase:  physAddrBaseRPi4,
			VideocoreBase: videocoreBaseRPi2,
			OscFreq:       oscFreqRPi4,
			Desc:          "Pi 4 Model B - 2GB v.1.1",
		},
		{
			Version:       0xC03111,
			RPiType:       RPiType4,
			PhysAddrBase:  physAddrBaseRPi4,
			VideocoreBase: videocoreBaseRPi2,
			Desc:          "Pi 4 Model B - 4GB v1.1",
		},
		{
			Version:       0xA03112,
			RPiType:       RPiType4,
			PhysAddrBase:  physAddrBaseRPi4,
			VideocoreBase: videocoreBaseRPi2,
			OscFreq:       oscFreqRPi4,
			Desc:          "Pi 4 Model B - 1GB v1.2",
		},
		{
			Version:       0xB03112,
			RPiType:       RPiType4,
			PhysAddrBase:  physAddrBaseRPi4,
			VideocoreBase: videocoreBaseRPi2,
			OscFreq:       oscFreqRPi4,
			Desc:          "Pi 4 Model B - 2GB v.1.2",
		},
		{
			Version:       0xC03112,
			RPiType:       RPiType4,
			PhysAddrBase:  physAddrBaseRPi4,
			VideocoreBase: videocoreBaseRPi2,
			OscFreq:       oscFreqRPi4,
			Desc:          "Pi 4 Model B - 4GB v1.2",
		},
		{
			Version:       0xb03114,
			RPiType:       RPiType4,
			PhysAddrBase:  physAddrBaseRPi4,
			VideocoreBase: videocoreBaseRPi2,
			OscFreq:       oscFreqRPi4,
			Desc:          "Pi 4 Model B - 2GB v1.4",
		},
		{
			Version:       0xD03114,
			RPiType:       RPiType4,
			PhysAddrBase:  physAddrBaseRPi4,
			VideocoreBase: videocoreBaseRPi2,
			OscFreq:       oscFreqRPi4,
			Desc:          "Pi 4 Model B - 8GB v1.4",
		},
		{
			Version:       0xc03114,
			RPiType:       RPiType4,
			PhysAddrBase:  physAddrBaseRPi4,
			VideocoreBase: videocoreBaseRPi2,
			OscFreq:       oscFreqRPi4,
			Desc:          "Pi 4 Model B - 4GB v1.4",
		},
		{
			Version:       0xb03115,
			RPiType:       RPiType4,
			PhysAddrBase:  physAddrBaseRPi4,
			VideocoreBase: videocoreBaseRPi2,
			OscFreq:       oscFreqRPi4,
			Desc:          "Pi 4 Model B - 2GB v1.5",
		},
		//
		// Compute Module 4
		//
		{
			Version:       0xa03140,
			RPiType:       RPiType4,
			PhysAddrBase:  physAddrBaseRPi4,
			VideocoreBase: videocoreBaseRPi2,
			OscFreq:       oscFreqRPi4,
			Desc:          "Compute Module 4 v1.0 eMMC",
		},
		{
			Version:       0xb03140,
			RPiType:       RPiType4,
			PhysAddrBase:  physAddrBaseRPi4,
			VideocoreBase: videocoreBaseRPi2,
			OscFreq:       oscFreqRPi4,
			Desc:          "Compute Module 4 v1.0 Lite",
		},
		{
			Version:       0xc03140,
			RPiType:       RPiType4,
			PhysAddrBase:  physAddrBaseRPi4,
			VideocoreBase: videocoreBaseRPi2,
			OscFreq:       oscFreqRPi4,
			Desc:          "Compute Module 4 v1.0 WiFi",
		},
		{
			Version:       0xd03140,
			RPiType:       RPiType4,
			PhysAddrBase:  physAddrBaseRPi4,
			VideocoreBase: videocoreBaseRPi2,
			OscFreq:       oscFreqRPi4,
			Desc:          "Compute Module 4 v1.0 WiFi 8GB",
		},
	}
)
