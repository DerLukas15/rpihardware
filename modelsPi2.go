package rpihardware

const (
	physAddrBaseRPi2  uint32 = 0x3f000000
	videocoreBaseRPi2 uint32 = 0xc0000000
)

var (
	model2List = []Hardware{
		//
		// Model Zero 2 W
		//
		{
			Version:       0x902120,
			RPiType:       RPiType2,
			PhysAddrBase:  physAddrBaseRPi2,
			VideocoreBase: videocoreBaseRPi2,
			OscFreq:       oscFreqRPi1,
			Desc:          "Pi Zero 2 W v1.0",
		},
		//
		// Pi 2 Model B
		//
		{
			Version:       0xa01041,
			RPiType:       RPiType2,
			PhysAddrBase:  physAddrBaseRPi2,
			VideocoreBase: videocoreBaseRPi2,
			OscFreq:       oscFreqRPi1,
			Desc:          "Pi 2",
		},
		{
			Version:       0xa01040,
			RPiType:       RPiType2,
			PhysAddrBase:  physAddrBaseRPi2,
			VideocoreBase: videocoreBaseRPi2,
			OscFreq:       oscFreqRPi1,
			Desc:          "Pi 2",
		},
		{
			Version:       0xa21041,
			RPiType:       RPiType2,
			PhysAddrBase:  physAddrBaseRPi2,
			VideocoreBase: videocoreBaseRPi2,
			OscFreq:       oscFreqRPi1,
			Desc:          "Pi 2",
		},
		//
		// Pi 2 with BCM2837
		//
		{
			Version:       0xa22042,
			RPiType:       RPiType2,
			PhysAddrBase:  physAddrBaseRPi2,
			VideocoreBase: videocoreBaseRPi2,
			OscFreq:       oscFreqRPi1,
			Desc:          "Pi 2",
		},
		//
		// Pi 3 Model B
		//
		{
			Version:       0xa020d3,
			RPiType:       RPiType2,
			PhysAddrBase:  physAddrBaseRPi2,
			VideocoreBase: videocoreBaseRPi2,
			OscFreq:       oscFreqRPi1,
			Desc:          "Pi 3 B+",
		},
		{
			Version:       0xa02082,
			RPiType:       RPiType2,
			PhysAddrBase:  physAddrBaseRPi2,
			VideocoreBase: videocoreBaseRPi2,
			OscFreq:       oscFreqRPi1,
			Desc:          "Pi 3",
		},
		{
			Version:       0xa02083,
			RPiType:       RPiType2,
			PhysAddrBase:  physAddrBaseRPi2,
			VideocoreBase: videocoreBaseRPi2,
			OscFreq:       oscFreqRPi1,
			Desc:          "Pi 3",
		},
		{
			Version:       0xa22082,
			RPiType:       RPiType2,
			PhysAddrBase:  physAddrBaseRPi2,
			VideocoreBase: videocoreBaseRPi2,
			OscFreq:       oscFreqRPi1,
			Desc:          "Pi 3",
		},
		{
			Version:       0xa22083,
			RPiType:       RPiType2,
			PhysAddrBase:  physAddrBaseRPi2,
			VideocoreBase: videocoreBaseRPi2,
			OscFreq:       oscFreqRPi1,
			Desc:          "Pi 3",
		},
		{
			Version:       0x9020e0,
			RPiType:       RPiType2,
			PhysAddrBase:  physAddrBaseRPi2,
			VideocoreBase: videocoreBaseRPi2,
			OscFreq:       oscFreqRPi1,
			Desc:          "Model 3 A+",
		},

		//
		// Pi Compute Module 3
		//
		{
			Version:       0xa020a0,
			RPiType:       RPiType2,
			PhysAddrBase:  physAddrBaseRPi2,
			VideocoreBase: videocoreBaseRPi2,
			OscFreq:       oscFreqRPi1,
			Desc:          "Compute Module 3/L3",
		},
		//
		// Pi Compute Module 3+
		//
		{
			Version:       0xa02100,
			RPiType:       RPiType2,
			PhysAddrBase:  physAddrBaseRPi2,
			VideocoreBase: videocoreBaseRPi2,
			OscFreq:       oscFreqRPi1,
			Desc:          "Compute Module 3+",
		},
	}
)
