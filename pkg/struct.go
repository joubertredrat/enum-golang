package pkg

import (
	"github.com/joubertredrat/enum-golang/pkg/enum"
)

//Person Person
type Person struct {
	Name        string           `json:"name"`
	Age         uint             `json:"age"`
	AvatarColor enum.AvatarColor `json:"avatarColorData"`
	State       enum.UF          `json:"uf"`
}
