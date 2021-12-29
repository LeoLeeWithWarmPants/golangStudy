package models

type Menu struct {
	Title     string
	menuItems []MenuItem
}

type MenuItem struct {
	rank     int
	itemName string
}
