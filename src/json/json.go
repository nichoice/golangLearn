package main

import (
	"encoding/json"
	"fmt"
)

type People struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

type Pepleslice struct {
	Peoples []People `json:"peoples"`
}

func main() {
	var p Pepleslice
	p.Peoples = append(p.Peoples, People{Name: "test2", Age: 16})
	p.Peoples = append(p.Peoples, People{Name: "test1", Age: 18})
	b, err := json.Marshal(p)
	if err != nil {
		fmt.Println("JSON ERR: ", err)
	}
	fmt.Println(string(b))
}
