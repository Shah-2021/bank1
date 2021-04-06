package types

// Money представляет собой денежную сумму в минимальных единицах (центы, копейки, дирамы и т.д.)
type Money int64

// Category представляет собой категорию, в которой был совершён платёж (авто, аптеки, рестораны и т.д.)
type Category string


// Status представляет собой статус платежа
type Status string

// Предопределенные статусы платежей
const (
	StatusOk Status = "OK"
	StatusFail Status = "FAIL"
	StatusInProgress Status = "INPROGRESS"
) 

// Currency представляет код валюты
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

// Card представляет информацию о платёжной карте
type Card struct{
	ID		int
	PAN		PAN
	Balance Money  // использовали Money
	Currency Currency
	Color	string
	Name	string
	Active	bool
	MinBalance Money
}

// Payment представляет информацию о платеже
type Payment struct{
	ID int
	Amount Money
	Category Category
	Status Status
}

// PaymentSource источник платежа
type PaymentSource struct {
Type string // 'card'
Number string // номер вида '5058 xxxx xxxx 8888'
Balance Money // баланс в дирамах
}