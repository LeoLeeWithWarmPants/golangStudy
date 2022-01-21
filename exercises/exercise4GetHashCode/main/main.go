package main

import (
	"fmt"
	"math"
	"strings"
)

func hashForJava(str string) int32 {
	var hashCode int32 = 0
	valueArray := strToCharArray(str)
	if len(str) > 0 {
		for i := 0; i < len(valueArray); i++ {
			hashCode = 31*hashCode + int32(valueArray[i])
		}
	}
	return hashCode
}

func strToCharArray(str string) []byte {
	return []byte(str)
}

func getResult(strKey string, salt int32) int32 {
	hashResult := hashForJava(strKey)
	result := int32(math.Abs(float64(hashResult))) % salt
	return result
}

func printStrKeyReminder() {
	fmt.Println("喂！请输入分表键：")
}

func printSaltReminder() {
	fmt.Println("喂！请输入分表数：")
}

func printResult(result int32) {
	fmt.Println(result)
}

func printContinueText() {
	fmt.Println("输入666继续计算")
}

func logic(strKey string, salt int32) {
	var continueText string
	for {
		printStrKeyReminder()
		fmt.Scanln(&strKey)
		printSaltReminder()
		fmt.Scanln(&salt)
		printResult(getResult(strKey, salt))
		printContinueText()
		fmt.Scanln(&continueText)
		if strings.Compare(continueText, "666") == 0 {

		} else {
			break
		}
	}
}

func main() {
	var strKey string
	var salt int32
	logic(strKey, salt)
}
