package account

import (
	"sync"
)

type Account struct {

	balance int
	open bool
	mux sync.Mutex
}

func (account *Account) Balance() (balance int, ok bool) {

	account.mux.Lock()
	defer account.mux.Unlock()

	return account.balance, account.open
}

func (account *Account) update(balance int, open bool) {

	account.mux.Lock()
	defer account.mux.Unlock()
	account.balance, account.open = balance, open
}

func Open(amt int) (account *Account) {

	if amt >= 0 {

		account = &Account{}
		account.update(amt, true)
	} else {
		account = nil
	}
	return 
}

func (account *Account) Close() (payout int, ok bool) {

	account.mux.Lock()
	defer account.mux.Unlock()

	if !account.open {
		payout, ok = 0, false
		return
	}

	payout, ok = account.balance, true
	account.balance, account.open = 0, false	
	return
}

func (account *Account) Open() (open bool) {

	account.mux.Lock()
	account.mux.Unlock()
	open = account.open
	return
}


func (account *Account) Deposit(amt int) (balance int, ok bool) {

	account.mux.Lock()
	defer account.mux.Unlock()

	if !account.open {
		return
	}

	if new_balance := account.balance+amt; new_balance >= 0 {
		account.balance = new_balance
		balance, ok = account.balance, true
	} else {
		ok = false
	}
	return
} 



