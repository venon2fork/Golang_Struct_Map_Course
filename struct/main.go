package main

import (
	"fmt"
	"jodo_crashcourse/EXO/struct/player"
)

func main() {
	var p1 player.Player
	p1.Age = 22
	p1.Alias = "Golum"
	p1.Login = "golmu@gmail.com"
	fmt.Printf("%v  \n", p1)
	a := player.Avatar{Url: "https://avatar.jpg"}
	fmt.Printf("%v\n", a)

	p3 := player.Player{
		Nom:    "john",
		Prenom: "adi",
		Age:    32,
		Login:  "milo@gmail.com",
		Alias:  "milo",
		Avatar: player.Avatar{
			Url: "https://avatar2.jpg",
		},
	}
	fmt.Printf("%v\n", p3)
	p4 := player.New("dan")
	p4.Avatar = a
	fmt.Printf("%v \n", p4)

	p5 := player.New("koki")
	p5.City = "France"
	p5.Avatar = a
	fmt.Printf("%v \n", p5)

	var p6 player.Player
	p6.Vie = 1
	p6.Point = 100

	fmt.Println(p6.CalculPoint())

}
