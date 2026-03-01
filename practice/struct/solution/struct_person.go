package solution

import "fmt"

type Person struct {
	Name string
	Age  int
	Sex  string
}

// ModifyAge1 值类型接收者
func (p Person) ModifyAge1(age int) {
	p.Age = age // 这个修改只是针对副本，不会影响原来的值
}

// ModifyAge2 指针型接收者
func (p *Person) ModifyAge2(age int) {
	p.Age = age // 这个修改会影响原来的值
}

func StructPersonTest() {
	p := Person{
		Age:  18,
		Name: "张三",
		Sex:  "男",
	}
	fmt.Println(p)
	p.ModifyAge1(15)
	fmt.Println(p)
	p.ModifyAge2(17)
	fmt.Println(p)
}
