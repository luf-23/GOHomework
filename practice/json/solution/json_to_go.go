package solution

import (
	json2 "encoding/json"
	"fmt"
)

const jsonData = `{"name":"John Doe","age":30,"hobbies":["reading","writing","coding"]}`

func JsonToGo() {
	var person Person

	err := json2.Unmarshal([]byte(jsonData), &person)

	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println(person)
}
