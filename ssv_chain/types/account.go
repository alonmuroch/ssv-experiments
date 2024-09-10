package types

import (
	"bytes"
	"fmt"
)

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

// GetBalance returns balance by token address or nil if not found
func (account *Account) GetBalance(address []byte, network [4]byte) *Balance {
	for _, b := range account.Balances {
		if bytes.Equal(b.Network[:], network[:]) && bytes.Equal(b.TokenAddress, address) {
			return b
		}
	}
	return nil
}

// SufficientBalance returns true if balance sufficient (bigger) for token
func (account *Account) SufficientBalance(balance uint64, tokenAddress []byte, network [4]byte) bool {
	panic("implement")
}

// ReduceBalance reduces balance for token, if balance available and not locked. Returns error if not sufficient balance
func (account *Account) ReduceBalance(balance uint64, tokenAddress []byte, network [4]byte) error {
	b := account.GetBalance(tokenAddress, network)
	if b == nil {
		return fmt.Errorf("balance not found")
	}

	if b.Amount < balance {
		return fmt.Errorf("insufficient balance")
	}
	b.Amount -= balance
	return nil
}

// LockBalance locks balance for token, if balanceToLock <= (balance - locked)
func (account *Account) LockBalance(balanceToLock uint64, tokenAddress []byte, network [4]byte) error {
	panic("implement")
}

// ReleaseBalance releases balance for token, if balanceToLock <= locked
func (account *Account) ReleaseBalance(balanceToRelease uint64, tokenAddress []byte, network [4]byte) error {
	panic("implement")
}

// AddBalance adds to balance (if exists) or creates new balance for account
func (account *Account) AddBalance(balance uint64, tokenAddress []byte, network [4]byte) {
	panic("implement")
}
