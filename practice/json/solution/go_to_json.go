package solution

import (
	json2 "encoding/json"
	"fmt"
)

type Person struct {
	Name    string   `json:"name"`
	Age     int      `json:"age"`
	Sex     string   `json:"sex"`
	Hobbies []string `json:"hobbies"`
}

func GoToJson() {
	person := Person{
		Name:    "张三",
		Age:     18,
		Sex:     "男",
		Hobbies: []string{"football", "basketball"},
	}
	fmt.Println(person)
	// marshal : 序列化
	json, err := json2.Marshal(person)

	if err != nil {
		fmt.Println("error:", err)
		return
	}
	fmt.Println(string(json))
}
