package main

import "fmt"

//键盘输入
func main() {
	var name string
	var age byte
	var salary float32
	var isMarried bool
	//fmt.Scanln,当程序执行到Scanln，程序会停止在这里，等待控制台的输入，待换行之后提交输入并向下运行
	/*fmt.Println("plz input ur name")
	fmt.Scanln(&name)
	fmt.Println("plz input ur age")
	fmt.Scanln(&age)
	fmt.Println("plz input ur salary")
	fmt.Scanln(&salary)
	fmt.Println("did u married?")
	fmt.Scanln(&isMarried)
	fmt.Printf("name:%s,age:%d,salary:%g,isMarried:%t\n", name, age, salary, isMarried)*/

	//fmt.Scanf
	fmt.Println("plz input ur name, age, salary, isMarried,use blank splice")
	fmt.Scanf("%s %d %g %t", &name, &age, &salary, &isMarried)
	fmt.Printf("name:%s,age:%d,salary:%g,isMarried:%t\n", name, age, salary, isMarried)
}
