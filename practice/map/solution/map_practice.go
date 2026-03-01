package solution

import "fmt"

/*
1. 创建一个存储学生信息的map，键为学生姓名(string)，值为年龄(int)
2. 添加5个学生信息：张三(18)、李四(19)、王五(20)、赵六(21)、钱七(22)
3. 查询并打印"李四"的年龄
4. 修改"王五"的年龄为25
5. 删除"赵六"的信息
6. 遍历打印所有剩余学生的信息
7. 统计map中学生的总人数
*/

func MapPractice01() {
	stuInfo := make(map[string]int)
	stuInfo["张三"] = 18
	stuInfo["李四"] = 19
	stuInfo["王五"] = 20
	stuInfo["赵六"] = 21
	stuInfo["钱七"] = 22
	fmt.Println(stuInfo["李四"])
	stuInfo["王五"] = 25
	delete(stuInfo, "赵六")
	for k, v := range stuInfo {
		fmt.Println(k, v)
	}
	fmt.Println(len(stuInfo))
}
