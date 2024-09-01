package main

import (
	"fmt"

	"github.com/gkcharlesdevs/cucumber/auth"
	"github.com/gkcharlesdevs/cucumber/hello"
	"github.com/gkcharlesdevs/cucumber/service"
)

var y1 counter = 23

func main(){
	fmt.Println(name)
	fmt.Println(school)
	fmt.Println(pi)

	hellos()
	fmt.Println(y1)
	fmt.Println(hello.Milage)
	service.LoggedIn()
	auth.CheckedIn()
}