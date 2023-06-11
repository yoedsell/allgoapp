package main

import "myapp/Routes"

// type Student struct {
// 	Name string `Json : "name"`
// 	Age  int    `Json : "age"`
// }

// var jsonData = `{
// 	"name" : "Tshering",
// 	"age" : 20
// }`

// func structToJson(stud Student) {
// 	sbyte, err := json.Marshal(stud)
// 	_ = err

// 	fmt.Println(string(sbyte))
// }

// func jsonToStruct(jsondata string) {
// 	var stud Student

// 	err := json.Unmarshal([]byte(jsondata), &stud)
// 	_ = err

// 	fmt.Printf("%+v", stud)
// }

func main() {
	Routes.InitializeRoutes()

	// Student1 := Student{
	// 	Name: "Pema",
	// 	Age:  20,
	// }

	// structToJson(Student1)

	// jsonToStruct(jsonData)

}
