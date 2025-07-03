package main

import "fmt"

type T struct {
	User User
}

type User struct {
	Name string
	Age  int
}

func main() {
	t := T{User: User{Name: "user1", Age: 1}}
	fmt.Println(t)

	fmt.Println(t.User)

	fmt.Println(t.User.Name)

	// fmt.Println(t.Name)
	// 省略した場合
	// type T struct {
	// 	User
	// }
	//output:
}
