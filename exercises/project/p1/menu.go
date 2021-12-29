package main

import (
	"fmt"
	"golangStudyProject/exercises/project/p1/constant"
	"golangStudyProject/exercises/project/p1/models"
)

var menu *models.Menu

func buildMenu() {
	if menu == nil {
		menu = new(models.Menu)
	}
	menu.Title = constant.MenuTitle
}

func PrintMenuText(menu *models.Menu) {
	fmt.Printf("%s%s%s\n", constant.CutOffLine, menu.Title, constant.CutOffLine)
}

func PrintMenu() {
	buildMenu()
	PrintMenuText(menu)
}

func PrintMenuItems(menuItems *models.MenuItem) {

}

func main() {
	PrintMenu()
}
