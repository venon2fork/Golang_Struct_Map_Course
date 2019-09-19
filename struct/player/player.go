package player

import (
	"fmt"
)

type Adress struct {
	City string
}
type Avatar struct {
	Url string
}

type PointsVie struct {
	Point int
	Vie   int
}

type Player struct {
	Nom, Prenom, Alias string
	Age                int
	Login              string
	Avatar             Avatar
	password           string
	Adress
	PointsVie
	Amis
	Post
}
type Post struct {
	Title     string
	Soustitre string
	Texte     string
	published bool
	Day       int
	Comment   string
	Like      int
	View      int
}
type Amis struct {
	Amis string
}

func (p PointsVie) CalculPoint() int {
	return (p.Point / 10) * p.Vie
}

func (p *PointsVie) String() string {
	return fmt.Sprintf("Point ==> %v  vie ===> %v", p.Vie, p.Point)
}

func New(Nom string) Player {
	return Player{
		Nom:      Nom,
		password: "DefaultPassword",
	}
}
