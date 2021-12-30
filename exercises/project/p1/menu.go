package main

import (
	"fmt"
	"golangStudyProject/exercises/project/p1/constant"
	"golangStudyProject/exercises/project/p1/models"
)

var menu *models.Menu

func buildMainMenu() {
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

func printMainMenuText(menu *models.Menu) {
	PrintTitle(menu.Title)
	if menu.MenuItems != nil && len(menu.MenuItems) > 0 {
		for _, ele := range menu.MenuItems {
			fmt.Printf("%s%d.%s\n", constant.Spaces1+constant.Spaces2, ele.Rank, ele.ItemName)
		}
	}
}

func PrintTitle(menuItemName string) {
	fmt.Printf("%s%s%s\n", constant.CutOffLine, menuItemName, constant.CutOffLine)
}

func PrintMainMenu() {
	buildMainMenu()
	printMainMenuText(menu)
}
