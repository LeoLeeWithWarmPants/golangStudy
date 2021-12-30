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

func scanFinanceDetail(optionType int8) (float64, string) {
	var amount float64
	var remark string
	if optionType == 1 {
		fmt.Println("请输入本次收入金额")
	} else if optionType == 2 {
		fmt.Println("请输入本次支出金额")
	}
	fmt.Scanln(&amount)
	if optionType == 1 {
		fmt.Println("请输入本次收入备注")
	} else if optionType == 2 {
		fmt.Println("请输入本次支出备注")
	}
	fmt.Scanln(&remark)
	return amount, remark
}

func Income() {
	var incomeAmount float64
	var remark string
	PrintTitle(constant.MenuItemNames[1])

	incomeAmount, remark = scanFinanceDetail(1)

	fd := models.FinanceDetail{Type: 1, Amount: incomeAmount, Remark: remark}
	accountOption(fd)
}

func Expenditure() {
	var incomeAmount float64
	var remark string
	PrintTitle(constant.MenuItemNames[2])

	incomeAmount, remark = scanFinanceDetail(2)

	fd := models.FinanceDetail{Type: 2, Amount: incomeAmount, Remark: remark}
	accountOption(fd)
}

func accountOption(ao models.AccountOption) {
	ao.DoOption(account)
}
