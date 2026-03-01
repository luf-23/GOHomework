package solution

import "fmt"

/*
1. 创建一个包含 10 个整数的切片，初始值为 1 到 10
2. 使用切片操作获取第 3 到第 7 个元素（包含第 7 个）
3. 在切片末尾添加三个新元素：11, 12, 13
4. 删除切片中的第 5 个元素
5. 将切片中的所有元素乘以 2
6. 打印最终切片的内容和容量
*/

func SlicePractice01() {
	sli := make([]int, 10)
	for i := 1; i <= 10; i++ {
		sli[i-1] = i
	}
	fmt.Println("sli:", sli)
	s1 := sli[2:7]
	fmt.Println("s1:", s1)
	s1 = append(s1, 11, 12, 13)
	fmt.Println("s1:", s1)
	s1 = append(s1[:4], s1[6:]...)
	fmt.Println("s1:", s1)
	for i, v := range s1 {
		s1[i] = v * 2
	}
	fmt.Println("s1:", s1)
	fmt.Println(len(s1), cap(s1))
}
