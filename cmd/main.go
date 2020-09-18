package main

import (
	"encoding/json"
	"fmt"

	"github.com/joubertredrat/enum-golang/pkg"
)

func main() {
	bytes := []byte(`[
		{
			"name": "Jão",
			"age": 17,
			"avatarColorData": "blue",
			"uf": "MG"
		},
		{
			"name": "Marias",
			"age": 28,
			"avatarColorData": "yellow",
			"uf": "RJ"
		},
		{
			"name": "Zé",
			"age": 30,
			"avatarColorData": "black",
			"uf": "ES"
		}
	]`)

	fmt.Printf("%+v\n", string(bytes))

	var people []pkg.Person
	json.Unmarshal(bytes, &people)

	fmt.Printf("%+v\n\n\n", people)

	for _, person := range people {
		fmt.Printf("%+v, %+v\n", person.Name, person.Age)
		fmt.Printf("  AvatarColor.GetKey(): %+v\n", person.AvatarColor.GetKey())
		fmt.Printf("  AvatarColor.GetValue(): %+v\n", person.AvatarColor.GetValue())
		fmt.Printf("  AvatarColor.IsValid(): %+v\n", person.AvatarColor.IsValid())
		fmt.Printf("  State.GetKey(): %+v\n", person.State.GetKey())
		fmt.Printf("  State.GetValue(): %+v\n", person.State.GetValue())
		fmt.Printf("  State.IsValid(): %+v\n", person.State.IsValid())
	}

	json, _ := json.Marshal(people)
	fmt.Printf("\n\n\n%+v\n", string(json))
}
