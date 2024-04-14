package mdvt

// Fork is a unique identifier for an SSV network fork version, 2 identical pieces of data signed with different domains will result in different sigs
type Fork [4]byte

var (
	GenesisMainnet = Fork{0x0, 0x0, 0x0, 0x0}
	TestnetV1      = Fork{0x0, 0x0, 0x1, 0x0}
)

type SignatureType [4]byte

var (
	QBFTSignatureType    SignatureType = [4]byte{1, 0, 0, 0}
	PartialSignatureType SignatureType = [4]byte{2, 0, 0, 0}
)
