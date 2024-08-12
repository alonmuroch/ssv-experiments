package types

type Account struct {
	// Address is an ethereum address controlling the account
	Address []byte
	// Nonce is a nonce of the account, starting with 0
	Nonce uint64
	// Balance amount of non-withdrawable balance in SSV tokens
	Balance uint64
	// Withdrawable amount of withdrawable balance in SSV tokens
	Withdrawable uint64
}
