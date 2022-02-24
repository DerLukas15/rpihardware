package rpihardware

const (
	physAddrBaseRPi1  uint32 = 0x20000000
	videocoreBaseRPi1 uint32 = 0x40000000
	oscFreqRPi1       uint32 = 19200000
)

var (
	model1List = []Hardware{
		//
		// Model B Rev 1.0
		//
		{
			Version:       0x02,
			RPiType:       RPiType1,
			PhysAddrBase:  physAddrBaseRPi1,
			VideocoreBase: videocoreBaseRPi1,
			OscFreq:       oscFreqRPi1,
			Desc:          "Model B",
		},
		{
			Version:       0x03,
			RPiType:       RPiType1,
			PhysAddrBase:  physAddrBaseRPi1,
			VideocoreBase: videocoreBaseRPi1,
			OscFreq:       oscFreqRPi1,
			Desc:          "Model B",
		},

		//
		// Model B Rev 2.0
		//
		{
			Version:       0x04,
			RPiType:       RPiType1,
			PhysAddrBase:  physAddrBaseRPi1,
			VideocoreBase: videocoreBaseRPi1,
			OscFreq:       oscFreqRPi1,
			Desc:          "Model B",
		},
		{
			Version:       0x05,
			RPiType:       RPiType1,
			PhysAddrBase:  physAddrBaseRPi1,
			VideocoreBase: videocoreBaseRPi1,
			OscFreq:       oscFreqRPi1,
			Desc:          "Model B",
		},
		{
			Version:       0x06,
			RPiType:       RPiType1,
			PhysAddrBase:  physAddrBaseRPi1,
			VideocoreBase: videocoreBaseRPi1,
			OscFreq:       oscFreqRPi1,
			Desc:          "Model B",
		},

		//
		// Model A
		//
		{
			Version:       0x07,
			RPiType:       RPiType1,
			PhysAddrBase:  physAddrBaseRPi1,
			VideocoreBase: videocoreBaseRPi1,
			OscFreq:       oscFreqRPi1,
			Desc:          "Model A",
		},
		{
			Version:       0x08,
			RPiType:       RPiType1,
			PhysAddrBase:  physAddrBaseRPi1,
			VideocoreBase: videocoreBaseRPi1,
			OscFreq:       oscFreqRPi1,
			Desc:          "Model A",
		},
		{
			Version:       0x09,
			RPiType:       RPiType1,
			PhysAddrBase:  physAddrBaseRPi1,
			VideocoreBase: videocoreBaseRPi1,
			OscFreq:       oscFreqRPi1,
			Desc:          "Model A",
		},

		//
		// Model B
		//
		{
			Version:       0x0d,
			RPiType:       RPiType1,
			PhysAddrBase:  physAddrBaseRPi1,
			VideocoreBase: videocoreBaseRPi1,
			OscFreq:       oscFreqRPi1,
			Desc:          "Model B",
		},
		{
			Version:       0x0e,
			RPiType:       RPiType1,
			PhysAddrBase:  physAddrBaseRPi1,
			VideocoreBase: videocoreBaseRPi1,
			OscFreq:       oscFreqRPi1,
			Desc:          "Model B",
		},
		{
			Version:       0x0f,
			RPiType:       RPiType1,
			PhysAddrBase:  physAddrBaseRPi1,
			VideocoreBase: videocoreBaseRPi1,
			OscFreq:       oscFreqRPi1,
			Desc:          "Model B",
		},

		//
		// Model B+
		//
		{
			Version:       0x10,
			RPiType:       RPiType1,
			PhysAddrBase:  physAddrBaseRPi1,
			VideocoreBase: videocoreBaseRPi1,
			OscFreq:       oscFreqRPi1,
			Desc:          "Model B+",
		},
		{
			Version:       0x13,
			RPiType:       RPiType1,
			PhysAddrBase:  physAddrBaseRPi1,
			VideocoreBase: videocoreBaseRPi1,
			OscFreq:       oscFreqRPi1,
			Desc:          "Model B+",
		},
		{
			Version:       0x900032,
			RPiType:       RPiType1,
			PhysAddrBase:  physAddrBaseRPi1,
			VideocoreBase: videocoreBaseRPi1,
			OscFreq:       oscFreqRPi1,
			Desc:          "Model B+",
		},

		//
		// Compute Module
		//
		{
			Version:       0x11,
			RPiType:       RPiType1,
			PhysAddrBase:  physAddrBaseRPi1,
			VideocoreBase: videocoreBaseRPi1,
			OscFreq:       oscFreqRPi1,
			Desc:          "Compute Module 1",
		},
		{
			Version:       0x14,
			RPiType:       RPiType1,
			PhysAddrBase:  physAddrBaseRPi1,
			VideocoreBase: videocoreBaseRPi1,
			OscFreq:       oscFreqRPi1,
			Desc:          "Compute Module 1",
		},

		//
		// Pi Zero
		//
		{
			Version:       0x900092,
			RPiType:       RPiType1,
			PhysAddrBase:  physAddrBaseRPi1,
			VideocoreBase: videocoreBaseRPi1,
			OscFreq:       oscFreqRPi1,
			Desc:          "Pi Zero v1.2",
		},
		{
			Version:       0x900093,
			RPiType:       RPiType1,
			PhysAddrBase:  physAddrBaseRPi1,
			VideocoreBase: videocoreBaseRPi1,
			OscFreq:       oscFreqRPi1,
			Desc:          "Pi Zero v1.3",
		},
		{
			Version:       0x920093,
			RPiType:       RPiType1,
			PhysAddrBase:  physAddrBaseRPi1,
			VideocoreBase: videocoreBaseRPi1,
			OscFreq:       oscFreqRPi1,
			Desc:          "Pi Zero v1.3",
		},
		{
			Version:       0x9200c1,
			RPiType:       RPiType1,
			PhysAddrBase:  physAddrBaseRPi1,
			VideocoreBase: videocoreBaseRPi1,
			OscFreq:       oscFreqRPi1,
			Desc:          "Pi Zero W v1.1",
		},
		{
			Version:       0x9000c1,
			RPiType:       RPiType1,
			PhysAddrBase:  physAddrBaseRPi1,
			VideocoreBase: videocoreBaseRPi1,
			OscFreq:       oscFreqRPi1,
			Desc:          "Pi Zero W v1.1",
		},
		//
		// Model A+
		//
		{
			Version:       0x12,
			RPiType:       RPiType1,
			PhysAddrBase:  physAddrBaseRPi1,
			VideocoreBase: videocoreBaseRPi1,
			OscFreq:       oscFreqRPi1,
			Desc:          "Model A+",
		},
		{
			Version:       0x15,
			RPiType:       RPiType1,
			PhysAddrBase:  physAddrBaseRPi1,
			VideocoreBase: videocoreBaseRPi1,
			OscFreq:       oscFreqRPi1,
			Desc:          "Model A+",
		},
		{
			Version:       0x900021,
			RPiType:       RPiType1,
			PhysAddrBase:  physAddrBaseRPi1,
			VideocoreBase: videocoreBaseRPi1,
			OscFreq:       oscFreqRPi1,
			Desc:          "Model A+",
		},
	}
)
