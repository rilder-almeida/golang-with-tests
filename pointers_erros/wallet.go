package pointerserros

import "fmt"

type Bitcoin int

type Stringer interface {
	String() string
}

type Wallet struct {
	Balance Bitcoin
}

func (b Bitcoin) String() string {
	return fmt.Sprintf("%d BTC", b)
}

func (w *Wallet) Deposit(amount Bitcoin) {
	w.Balance += amount
}

func (w *Wallet) GetBalance() (Bitcoin, error) {
	return w.Balance, nil
}

func (w *Wallet) Withdraw(amount Bitcoin) error {
	if amount > w.Balance {
		return InsufficientFounds(amount)
	}

	w.Balance -= amount
	return nil
}

func InsufficientFounds(amount Bitcoin) error {
	return fmt.Errorf("%s is insufficient", amount)
}
