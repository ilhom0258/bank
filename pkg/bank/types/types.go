package types

// Money - money type
type Money int64

// Currency - currency type
type Currency string

// Constants
const (
	TJS Currency = "TJS"
	USD Currency = "USD"
	RUB Currency = "RUB"
	EUR Currency = "EUR"
)

// PAN - pan code of cards
type PAN string

// Card  - Card structure
type Card struct {
	ID         int
	PAN        PAN
	Balance    Money
	MinBalance Money
	Currency   Currency
	Color      string
	Name       string
	Active     bool
}

//Category - category of payment(book, auto)
type Category string

// Payment - payment
type Payment struct {
	ID       int
	Amount   Money
	Category Category
}
