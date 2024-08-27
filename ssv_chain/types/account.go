package types

type Account struct {
	// Address is an L1 address controlling the account
	Address []byte `ssz-max:"128"`
	// Nonce is a nonce of the account, starting with 0
	Nonce uint64
	// Balances per token and network
	Balances []*Balance `ssz-max:"128"`
	// Withdrawable amount of withdrawable balance in SSV tokens
	Withdrawable uint64
}

// BalanceByTokenAddress returns balance by token address or nil if not found
func (account *Account) BalanceByTokenAddress(address []byte) *Balance {
	panic("implement")
}

// SufficientBalance returns true if balance sufficient (bigger) for token
func (account *Account) SufficientBalance(balance uint64, tokenAddress []byte, network [4]byte) bool {
	panic("implement")
}

// ReduceBalance reduces balance for token, returns error if not sufficient balance
func (account *Account) ReduceBalance(balance uint64, tokenAddress []byte, network [4]byte) error {
	panic("implement")
}

func (account *Account) DepositBalance(newBalance *Balance) {
	panic("implement")
}
