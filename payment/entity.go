package payment

type Transaction struct {
	ID     int
	Amount int // dibuat agar tidak cycle error
}