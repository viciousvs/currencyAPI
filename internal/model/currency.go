package model

type Currency struct {
	Id        int     `db:"id"`
	Date      string  `db:"date"`
	TimeStamp int     `db:"time_stamp"`
	Base      string  `db:"base"`
	Rate      string  `db:"rate"`
	Value     float64 `db:"value"`
}
type Currencies struct {
	Date      string             `json:"date" db:"date"`
	Timestamp int                `json:"timestamp" db:"time_stamp"`
	Base      string             `json:"base" db:"base"`
	Rates     map[string]float64 `json:"rates" db:"rates"`
}
