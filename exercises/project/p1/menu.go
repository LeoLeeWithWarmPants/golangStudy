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
	(*menu).Title = constant.MenuTitle
	if (*menu).MenuItems == nil {
		(*menu).MenuItems = make([]models.MenuItem, 4)
	}
	for idx, menuItemName := range constant.MenuItemNames {
		menuItemTemp := models.MenuItem{Rank: idx + 1, ItemName: menuItemName}
		(*menu).MenuItems[idx] = menuItemTemp
	}
}

func PrintMenuText(menu *models.Menu) {
	fmt.Printf("%s%s%s\n", constant.CutOffLine, menu.Title, constant.CutOffLine)
	if menu.MenuItems != nil && len(menu.MenuItems) > 0 {
		for _, ele := range menu.MenuItems {
			fmt.Printf("%s%d.%s\n", constant.Spaces, ele.Rank, ele.ItemName)
		}
	}
}

func PrintMenu() {
	buildMenu()
	PrintMenuText(menu)
}

func main() {
	PrintMenu()
}
