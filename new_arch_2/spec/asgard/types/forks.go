package types

// Domain is a unique identifier for signatures, 2 identical pieces of data signed with different domains will result in different sigs
type Domain [4]byte

var (
	GenesisMainnet = Domain{0x0, 0x0, 0x0, 0x0}
	PrimusTestnet  = Domain{0x0, 0x0, 0x1, 0x0}
	ShifuTestnet   = Domain{0x0, 0x0, 0x2, 0x0}
	ShifuV2Testnet = Domain{0x0, 0x0, 0x2, 0x1}
	V3Testnet      = Domain{0x0, 0x0, 0x3, 0x1}
)

type SignatureType [4]byte

var (
	QBFTSignatureType    SignatureType = [4]byte{1, 0, 0, 0}
	PartialSignatureType SignatureType = [4]byte{2, 0, 0, 0}
	DKGSignatureType     SignatureType = [4]byte{3, 0, 0, 0}
)
