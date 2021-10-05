package main

import "fmt"

type MyType struct {
	Name string
}

func main() {
	t := MyType{
		Name: "test",
	}
	fmt.Printf("%p\n", &t)
	p := &t
	fmt.Printf("%p\n", &p)
	// for i := 0; i < 3; i++ {
	// 	fmt.Println(i)
	// }
	// fillString := "hello world"
	// fmt.Println(fillString)
	// for i, c := range fillString {
	// 	fmt.Println(i, string(c))
	// }
	// i := 0
	// fmt.Print(i)

	// myarray := [3]int{1, 2, 3}
	// fmt.Println(myarray)

	// for i := range myarray {
	// 	// fmt.Print(i)
	// }

	// myslice := new([]int)
	// fmt.Println(myslice)

	// string_array := [5]string{"I", "am", "stupid", "and", "weak"}
	// fmt.Println(&string_array)
	// for index, s := range string_array {
	// 	fmt.Println(index, s)
	// 	if s == "stupid" {
	// 		string_array[index] = "smart"
	// 	} else if s == "weak" {
	// 		string_array[index] = "strong"
	// 	}
	// }
	// fmt.Println(string_array)
}
