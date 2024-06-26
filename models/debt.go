package main

type DebtStatus int8

const (
	DebtStatusPending DebtStatus = iota
	DebtStatusPaid
)

type Debt struct {
	Id        string
	Name      string
	Currency  string
	ClientId  string
	Date      string
	CreatedAt string
	UpdatedAt string
	Amount    float64
	Status    DebtStatus
}

type DebtService interface {
	create(debt Debt) (Debt, error)
	list() ([]Debt, error)
	get(id string) (Debt, error)
	update(id string, debt Debt) (Debt, error)
	delete(id string) error
	syncronize() error
}
