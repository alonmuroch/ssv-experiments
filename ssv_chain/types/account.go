package types

type Account struct {
	// Network on which the account exists
	Network [4]byte `ssz-size:"4"`
	// Address is an L1 address controlling the account
	Address []byte `ssz-max:"128"`
	// Nonce is a nonce of the account, starting with 0
	Nonce uint64
	// Balances per token and network
	Balances []*Balance `ssz-max:"128"`
}

// BalanceByTokenAddress returns balance by token address or nil if not found
func (account *Account) BalanceByTokenAddress(address []byte, network [4]byte) *Balance {
	panic("implement")
}

// SufficientBalance returns true if balance sufficient (bigger) for token
func (account *Account) SufficientBalance(balance uint64, tokenAddress []byte, network [4]byte) bool {
	panic("implement")
}

// ReduceBalance reduces balance for token, if balance available and not locked. Returns error if not sufficient balance
func (account *Account) ReduceBalance(balance uint64, tokenAddress []byte, network [4]byte) error {
	panic("implement")
}

// LockBalance locks balance for token, if balanceToLock <= (balance - locked)
func (account *Account) LockBalance(balanceToLock uint64, tokenAddress []byte, network [4]byte) error {
	panic("implement")
}

// ReleaseBalance releases balance for token, if balanceToLock <= locked
func (account *Account) ReleaseBalance(balanceToRelease uint64, tokenAddress []byte, network [4]byte) error {
	panic("implement")
}

func (account *Account) DepositBalance(newBalance *Balance) {
	panic("implement")
}
