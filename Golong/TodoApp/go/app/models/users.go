package models

import (
	"fmt"
	"log"
	"time"
)

type User struct {
	ID        int
	UUID      string
	Name      string
	Email     string
	Password  string
	CreatedAt time.Time
}

func (u *User) CreateUser() (err error) {
	cmd := `insert into users (uuid, name, email, password created_at) values (?, ?, ?, ?, ?)`

	fmt.Println("----------------------------------")
	fmt.Println(cmd)
	fmt.Println(createUUID())
	fmt.Println(u.Name)
	fmt.Println(u.Email)
	fmt.Println(Encrypt(u.Password))
	fmt.Println(time.Now())
	fmt.Println("----------------------------------")

	_, err = Db.Exec(cmd,
		createUUID(),
		u.Name,
		u.Email,
		Encrypt(u.Password),
		time.Now(),
	)

	if err != nil {
		log.Fatalln(err)
	}
	return err
}
