package main

import (
	"fmt"
	"strings"
)

// class Post....
type Post struct {
	Titre     string
	Soustitre string
	Article   string
	published bool
}

// fonction
func (p Post) modifiePost() string {
	return fmt.Sprintf("%v....._\n%v....._\n%v....", p.Titre[:25], p.Soustitre, p.Article[:30])
}

func (p Post) publisher() bool {
	return p.published
}

func (p *Post) publish() {
	p.published = true
}

func (p *Post) unpublish() {
	p.published = false
}

func TitleMajuscule(p *Post) {
	p.Titre = strings.ToUpper(p.Titre)
}

// fonction
func (p Post) afficheTitre(e string) []string {
	var x = []string{}
	x = strings.Fields(e)
	return x
}

func main() {

	p2 := Post{
		Titre:     "Go est vraiment bien merci pour ce langage",
		Soustitre: "Pourquoi choissir golang",
		Article: `lorem Le Lorem Ipsum est simplement du faux texte employé dans 
la composition et la mise en page avant impression. Le Lorem Ipsum est le 
faux texte standard de l'imprimerie depuis les années 1500, quand un imprimeur
anonyme assembla ensemble des morceaux de texte pour réaliser un livre spécimen
de polices de texte. Il n'a pas fait que survivre cinq siècles, mais s'est aussi 
adapté à la bureautique informatique, sans que son contenu n'en soit modifié. 
Il a été popularisé dans les années 1960 grâce à la vente de feuilles Letraset 
contenant des passages du Lorem Ipsum, et, plus récemment, par son inclusion dans 
des applications de mise en page de texte, comme Aldus PageMaker`,
	}
	// w := p2.afficheTitre(p2.Titre)
	// fmt.Printf("%v...\n", w[1:3])
	fmt.Println(p2.modifiePost())

	fmt.Println(p2.publisher()) //false
	p2.publish()                // on le modifie pour true
	fmt.Println(p2.publisher()) //true
	fmt.Println(p2.publisher()) //true
	p2.unpublish()              // on le modifie pour false
	fmt.Println(p2.publisher()) //false
	TitleMajuscule(&p2)
	fmt.Println(p2.modifiePost())

}
