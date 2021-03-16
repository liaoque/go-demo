package user

import (
	"bytes"
	"encoding/json"
	"fmt"
	"time"
)

type User struct {
	id int	`json:id`
	name string	`json:name`
	Tel string	`json:addr`
	Height float32	`json:height`
	Desc *string	`json:desc`
	Weight *int	`json:weight`
}



type User2 struct {
	Id int	`json:id`
	Name string	`json:name`
	Online bool	`json:is_online,string, omitempty`
	Register time.Time	`json:,omitempty`
}

func (user *User) Id() int {
	return user.id
}

func (user *User) SetId(id int) {
	user.id = id
}

func (user *User) Name() string {
	return user.name
}

func (user *User) SetName(name string) {
	user.name = name
}


func NewUser(id int, name string, tel string, height float32, desc *string, weight *int) *User {
	return &User{id: id, name: name, Tel: tel, Height: height, Desc: desc, Weight: weight}
}

func NewUser2(id int, name string, tel string, height float32, desc *string, weight *int) User {
	return User{id: id, name: name, Tel: tel, Height: height, Desc: desc, Weight: weight}
}

func (user User)String()string  {
	return fmt.Sprintf("User{id:%d, name: %s, tel: %s, height: %e, desc: %s, weight: %d}",
		user.id, user.name, user.Tel, user.Height, user.Height, *user.Desc, *user.Weight,
	)
}


func (user User2)Write(p []byte) (n int, err error){
	fmt.Printf("%T %s\n", p, string(p))

	bufferString := bytes.NewBufferString(string(p))

	var user2 []User2
	//json.NewDecoder(p)
	decoder := json.NewDecoder(bufferString)
	err = decoder.Decode(&user2)
	if err == nil {
		fmt.Println(user2)
	}else {
		fmt.Println(err)
	}
	return 10, nil
}

