package main

import (
	"fmt"
	"golangStudyProject/exercises/project/p1/constant"
	"golangStudyProject/exercises/project/p1/models"
)

var account *models.Account = new(models.Account)

func printFinanceDetailTitle() {
	fmt.Println(constant.FinanceDetailTile)
}

func PrintFinanceDetail(title string) {
	PrintTitle(title)
	printFinanceDetailTitle()
	PrintFinanceDetail0()
}

func PrintFinanceDetail0() {
	if account == nil {
		account = new(models.Account)
		account.Amount = 0.0
	}
	if account.FinanceDetail == nil {
		account.FinanceDetail = make([]models.FinanceDetail, 0)
	}
	if len(account.FinanceDetail) > 0 {
		account.PrintAccountDetail()
	}
}

func Income() {
	var incomeAmount float64
	var remark string
	PrintTitle(constant.MenuItemNames[1])
	fmt.Println("请输入本次收入金额")
	fmt.Scanln(&incomeAmount)
	fmt.Println("请输入本次收入备注")
	fmt.Scanln(&remark)

	fd := models.FinanceDetail{Type: 1, Amount: incomeAmount, Remark: remark}
	income0(fd)
}

func income0(ao models.AccountOption) {
	ao.DoOption(account)
}
