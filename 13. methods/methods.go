package main

import "fmt"

type User struct {
	name string
	email string
	status bool
}

func (u User) getStatus(){
	fmt.Println("Status is", u.status)
}

func (u User) newEmail(){
	u.email = "new@"
	fmt.Println("new email is", u.email)
}

func main(){
	a := User{name: "a", email: "a@", status: true}
	fmt.Println(a)
	a.getStatus()
	a.newEmail()
	fmt.Println(a)
}