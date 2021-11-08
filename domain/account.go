package domain

import "time"

type Transaction struct {
	date     time.Time
	category Category
	payee    Payee
	amount   Money
}

type Account struct {
	name         string
	transactions []*Transaction
}

func (a *Account) AddTransaction(t *Transaction) {
	a.transactions = append(a.transactions, t)
}

func (a *Account) Balance() Money {
	balance := Money(0)
	for _, t := range a.transactions {
		balance += t.amount
	}
	return balance
}

func (a *Account) Name() string {
	return a.name
}
