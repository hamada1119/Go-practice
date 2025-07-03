package main

import "fmt"

type User struct {
	Name string
	Age  int
}

func main() {
	m := map[int]User{
		1: {Name: "User1", Age: 10},
		2: {Name: "User2", Age: 20},
	}

	fmt.Println(m)

	m2 := map[User]string{
		{Name: "User1", Age: 10}: "Tokyo",
		{Name: "User2", Age: 20}: "LA",
	}

	fmt.Println(m2)

	m3 := make(map[int]User)
	fmt.Println(m3)
	m3[1] = User{Name: "User3"}
	fmt.Println(m3)

	for i, j := range m {
		fmt.Println(i, j)
	}

}
