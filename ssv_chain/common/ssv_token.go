package common

const (
	SSVDecimals = 18
	VBit        = 1
	VKBit       = 1000 * VBit
	VMBit       = 1000 * VKBit // 2^20
	VGBit       = 1000 * VMBit
	VTBit       = 1000 * VGBit
	VPBit       = 1000 * VTBit
	VEBit       = 1000 * VPBit
	MaxSSV      = 15000000
)

// helper functions for VGBit denomination
const (
	// VGBitOneSSV represents 1 whole SSV in VGBit denomination
	VGBitOneSSV           = uint64(1000000000)
	VGBitTenthSSV         = VGBitOneSSV / 10
	VGBitOneHundredthSSV  = VGBitOneSSV / 100
	VGBitOneThousandthSSV = VGBitOneSSV / 1000
	VGBitTenThousandthSSV = VGBitOneSSV / 10000
)
