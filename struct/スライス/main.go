package main

import "fmt"

type User struct {
	Name string
	Age  int
}

type Users []*User

func main() {
	user1 := User{Name: "user1", Age: 10}
	user2 := User{Name: "user2", Age: 20}
	user3 := User{Name: "user3", Age: 30}
	user4 := User{Name: "user4", Age: 40}

	users := Users{}

	users = append(users, &user1)
	users = append(users, &user2)
	users = append(users, &user3)
	users = append(users, &user4)

	for i, j := range users {
		fmt.Println(i, *j)
	}

	users2 := make([]*User, 0)
	users2 = append(users2, &user1, &user2)

	for i, j := range users2 {
		fmt.Println(i, j)
	}
}
