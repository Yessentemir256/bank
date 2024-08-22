package types

// Money представляет собой денежную сумму в минимальных еденицах (центы, копейки, дирамы и.т.д.).
type Money int64

// Category представляет собой категории, в которой был совершен платеж (авто, аптеки, рестораны и.т.д).
type Category string

// Status представляет собой статус платежа.
type Status string

// Предопределенные статусы платежей.
const (
	StatusOk         Status = "OK"
	StatusFail       Status = "FAIL"
	StatusInProgress Status = "INPROGRESS"
)

// Currency представляют код валюты
type Currency string

// Коды валют
const (
	TJS Currency = "TJS"
	RUB Currency = "RUB"
	USD Currency = "USD"
	EUR Currency = "EUR"
)

// PAN представляет номер карты
type PAN string

// Card представляет информацию о платежной карте
type Card struct {
	ID         int
	PAN        PAN
	Balance    Money // использовали Money
	MinBalance Money // использовали Money
	Currency   Currency
	Color      string
	Name       string
	Active     bool
}

// Payment представляет информацию о платеже.
type Payment struct {
	ID       int
	Amount   Money
	Category Category
}
