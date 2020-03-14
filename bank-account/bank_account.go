package account

import (
	"sync"
)

type Account struct {
	open		     bool
	balance			 int64
	mux              sync.Mutex
}

func (a *Account) transaction(f func() (int64, bool)) (int64, bool) {
	if !a.open {
		return 0, false
	}
	return f()
}

func (a *Account) Balance() (int64, bool) {
	return a.transaction(func() (int64, bool) {
		return a.balance, true
	})
}


func (a *Account) Close() (int64, bool) {
	return a.transaction(func() (int64, bool) {
		a.mux.Lock()
		defer a.mux.Unlock()

		balance, ok := a.Balance()
		a.open = false
		return balance, ok
	})
}

func (a *Account) Deposit(v int64) (int64, bool) {
	return a.transaction(func() (int64, bool) {
		a.mux.Lock()
		defer a.mux.Unlock()

		newBalance := a.balance + v
		if newBalance < 0 {
			return 0, false
		}

		a.balance = newBalance
		balance, ok := a.Balance()
		return balance, ok
	})
}

func Open(initialDeposit int64) *Account {
	if initialDeposit < 0 {
		return nil
	}

	return &Account{ balance: initialDeposit, open:true }
}