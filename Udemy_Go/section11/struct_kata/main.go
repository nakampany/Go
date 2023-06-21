package main

import "fmt"

type User struct {
	Name string
	Age  int
}

// ポインタのUserをスライスとして格納できる
type Users []*User

func main() {
	user1 := User{Name: "a", Age: 1}
	user2 := User{Name: "b", Age: 2}
	user3 := User{Name: "c", Age: 3}
	user4 := User{Name: "d", Age: 4}
	user5 := User{Name: "e", Age: 5}

	users := Users{}

	users = append(users, &user1)

	users = append(users, &user1)
	users = append(users, &user2)
	users = append(users, &user3)
	users = append(users, &user4, &user5)

	// 表示の時は、構造体のスライスなので、、、
	fmt.Println(users) // [0xc000010030]

	for _, u := range users {
		fmt.Println(u)
	}
	// &{a 1}
	// &{b 2}
	// &{c 3}
	// &{d 4}
	// &{e 5}

	users2 := make([]*User, 0)
	users2 = append(users2, &user1)
	users2 = append(users2, &user2)
	users2 = append(users2, &user3)

	for _, u := range users2 {
		fmt.Println(u)
	}

	// map
	m := map[int]User{
		1: {Name: "a", Age: 1},
		2: {Name: "b", Age: 2},
	}
	fmt.Println(m)

	n := map[User]string{
		{Name: "a", Age: 1}: "Tokyo",
		{Name: "b", Age: 2}: "Fukui",
	}
	fmt.Println(n)

}
