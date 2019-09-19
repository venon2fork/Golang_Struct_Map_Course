package main

import "fmt"

type Vector struct {
	X, Y int
}
type User struct {
	name string // map["hr"]= "bobette"
}

type Key struct {
	ID   int
	Name string
	//map[key{int,string}]=string
	//exemple m[key{1, "GoogleMoteur"}= "http://www.google.fr"]
}

func main() {
	// Slice
	p := make([]int, 0)
	//Map
	m := make(map[string]int)
	fmt.Println(p)
	fmt.Println(m)

	// creer une map avec une struct
	// v := make(map[Vector]string)
	// fmt.Println(v)
	// v[X{23}] = X
	// fmt.Println(v)

	m = make(map[string]int)
	fmt.Printf("value = %v  len= %v\n", m, len(m))

	m["hello"] = 5
	m["goodbye"] = 10
	fmt.Printf("value = %v  len= %v\n", m, len(m))

	// test la presence d'une dans une map
	k, presenceDuneValeur := m["helo"]
	fmt.Printf("k= %v presenceDuneValeur=%v \n", k, presenceDuneValeur)

	if _, presenceDuneValeur := m["godbye"]; presenceDuneValeur {
		fmt.Println("presence de la valeur")
	} else {
		fmt.Println("non presence de la valeur")

	}

	// effacer une valeur dans une map
	delete(m, "goodbye")
	fmt.Printf("value = %v  len= %v\n", m, len(m))
	// map literal
	m1 := map[string]int{
		"Bob":     10,
		"Alice":   15,
		"Bobette": 20,
		"John":    7,
	}

	for name, age := range m1 {
		fmt.Printf("name=%v, age=%v\n", name, age)
	}

	fmt.Println("Key only")
	i := 1
	for name := range m1 {
		fmt.Printf("name=%v\n", name)
		m1[name] = i
		i++
	}

	fmt.Println(m1)

	// Using a pointer or not has a big impact!
	m2 := map[string]*User{
		"HR":  {"Bob"},
		"CEO": {"Alice"},
	}

	fmt.Println(m2["HR"])

	hr := m2["HR"]
	hr.name = "John"
	fmt.Println(m2["HR"])

	// when using a pointer = just copy the pointer
	// when using a value = copy the entire struct
	m2["CFO"] = &User{"Bobette"}

	res := make(map[Key]string)
	res[Key{1, "aze"}] = "file1"

	ko := Key{2, "ert"}
	res[ko] = "file2"
	fmt.Println(res)
	delete(res, Key{1, "aze"})
	fmt.Println(res)

}
