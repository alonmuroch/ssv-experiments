package module

// Key represents the cryptographic key used by a module
type Key [4]byte

var (
	BLS = Key{0x0, 0x0, 0x0, 0x0}
	RSA = Key{0x0, 0x0, 0x1, 0x0}
)
