package models

type Menu struct {
	Title     string
	MenuItems []MenuItem
}

type MenuItem struct {
	Rank     int
	ItemName string
}
