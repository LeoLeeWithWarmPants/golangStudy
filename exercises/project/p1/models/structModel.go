package models

import (
	"fmt"
	"golangStudyProject/exercises/project/p1/constant"
)

type Menu struct {
	Title     string
	MenuItems []MenuItem
}

type MenuItem struct {
	Rank     int
	ItemName string
}

type Account struct {
	Amount        float64 //账户总余额
	FinanceDetail []FinanceDetail
}

func (account *Account) PrintAccountDetail() {
	for _, ele := range account.FinanceDetail {
		if ele.Type == 1 {
			fmt.Printf("收入%s%g%s%g%s%s\n", constant.Spaces3, ele.CurrentAmount, constant.Spaces3, ele.Amount, constant.Spaces3, ele.Remark)
		} else if ele.Type == 2 {
			fmt.Printf("支出%s%g%s%g%s%s\n", constant.Spaces3, ele.CurrentAmount, constant.Spaces3, ele.Amount, constant.Spaces3, ele.Remark)
		}
	}
}

type AccountOption interface {
	DoOption(account *Account)
}

type FinanceDetail struct {
	Type          int8 // 1-收入 2-支出
	Amount        float64
	CurrentAmount float64 //此时的账户余额
	Remark        string
}

func (financeDetail FinanceDetail) DoOption(account *Account) {
	if account == nil {
		account = new(Account)
		account.Amount = 0.0
	}
	if account.FinanceDetail == nil {
		account.FinanceDetail = make([]FinanceDetail, 0)
	}

	switch financeDetail.Type {
	case 1:
		account.Amount += financeDetail.Amount
	case 2:
		if (account.Amount - financeDetail.Amount) < 0 {
			fmt.Println("账户余额不足以支出")
		}
		account.Amount -= financeDetail.Amount
	}
	financeDetail.CurrentAmount = account.Amount
	account.FinanceDetail = append(account.FinanceDetail, financeDetail)
}
