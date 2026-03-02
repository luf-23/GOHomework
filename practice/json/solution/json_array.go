package solution

import (
	json2 "encoding/json"
	"fmt"
)

const jsonArray = `[
  {
    "name": "张三",
    "age": 18,
    "sex": "男",
    "hobbies": ["football", "swimming"]
  },
  {
    "name": "李四",
    "age": 19,
    "sex": "女",
    "hobbies": ["swimming", "reading"]
  }
]`

func JsonArray() {
	persons := make([]Person, 10)

	err := json2.Unmarshal([]byte(jsonArray), &persons)

	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println(persons)

	jsonDt, err2 := json2.Marshal(persons)

	if err2 != nil {
		fmt.Println("Error:", err2)
		return
	}

	fmt.Println(string(jsonDt))

}
