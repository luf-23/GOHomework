package main

import (
	"fmt"
	"math/rand"
)

func main() {
	guessDigit()
}

func guessDigit() {
	fmt.Println("欢迎来到猜数字游戏！")
	fmt.Println("规则：")
	fmt.Print("1.计算机将在1-100之间随机选择一个数字\n" +
		"2.您可以选择难度级别（简单、中等、困难），不同难度对应不同的猜测机会\n" +
		"3.请输入您的猜测。\n")
	fmt.Print("请选择难度级别序号：")
	fmt.Print("1.简单（10次机会）\n2.中等（5次机会）\n3.困难（3次机会）\n")
	var level int
	_, _ = fmt.Scanf("%d", &level)
	fmt.Print("输入选择：")
	fmt.Println("开始游戏：")
	number := rand.Int()%100 + 1

	var chance int

	switch level {
	case 1:
		chance = 10
	case 2:
		chance = 5
	case 3:
		chance = 3
	default:
		panic("输入错误")
	}

	var isok bool = false

	for i := 1; i <= chance; i++ {
		fmt.Printf("第%d次猜测，请输入您的数字（1-100）：", i)
		var now int
		_, _ = fmt.Scanf("%d", &now)
		if now < number {
			fmt.Println("您输入的数字偏小")
		} else if now > number {
			fmt.Println("您输入的数字偏大")
		} else {
			fmt.Printf("恭喜，您在第%d次中猜测成功！", i)
			isok = true
			break
		}
	}
	if !isok {
		fmt.Printf("很遗憾，正确数字为%d，您没有猜中，游戏结束！\n", number)
	}

}
