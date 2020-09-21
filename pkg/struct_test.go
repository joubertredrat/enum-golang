package pkg

import (
	"encoding/json"
	"testing"

	"gopkg.in/go-playground/assert.v1"

	"github.com/joubertredrat/enum-golang/pkg/enum"
)

func TestPersonStruct(t *testing.T) {
	nameExpected := "John"
	ageExpected := uint(20)
	var avatarColorExpected enum.AvatarColor = 1
	var stateExpected enum.UF = 1

	person := Person{
		Name:        nameExpected,
		Age:         ageExpected,
		AvatarColor: avatarColorExpected,
		State:       stateExpected,
	}

	assert.Equal(t, nameExpected, person.Name)
	assert.Equal(t, ageExpected, person.Age)
	assert.Equal(t, avatarColorExpected, person.AvatarColor)
	assert.Equal(t, stateExpected, person.State)
}

func TestPersonUnmarshalJSON(t *testing.T) {
	type testCase struct {
		testName            string
		jsonData            []byte
		nameExpected        string
		ageExpected         uint
		avatarColorExpected string
		stateExpected       string
	}

	tests := []testCase{
		{
			testName: "Test Unmarshal Jão",
			jsonData: []byte(`
				{
					"name": "Jão",
					"age": 17,
					"avatarColorData": "blue",
					"uf": "MG"
				}
			`),
			nameExpected:        "Jão",
			ageExpected:         uint(17),
			avatarColorExpected: "blue",
			stateExpected:       "MG",
		},
		{
			testName: "Test Unmarshal Marias",
			jsonData: []byte(`
				{
					"name": "Marias",
					"age": 28,
					"avatarColorData": "yellow",
					"uf": "RJ"
				}
			`),
			nameExpected:        "Marias",
			ageExpected:         uint(28),
			avatarColorExpected: "yellow",
			stateExpected:       "",
		},
		{
			testName: "Test Unmarshal Zé",
			jsonData: []byte(`
				{
					"name": "Zé",
					"age": 30,
					"avatarColorData": "black",
					"uf": "ES"
				}
			`),
			nameExpected:        "Zé",
			ageExpected:         uint(30),
			avatarColorExpected: "",
			stateExpected:       "ES",
		},
	}

	for _, test := range tests {
		t.Run(test.testName, func(t *testing.T) {
			t.Parallel()

			var person Person
			json.Unmarshal(test.jsonData, &person)

			assert.Equal(t, test.nameExpected, person.Name)
			assert.Equal(t, test.ageExpected, person.Age)
			assert.Equal(t, test.avatarColorExpected, person.AvatarColor.GetValue())
			assert.Equal(t, test.stateExpected, person.State.GetValue())
		})
	}
}

func TestPersonMarshalJSON(t *testing.T) {
	type testCase struct {
		testName     string
		person       func() Person
		jsonExpected string
	}

	tests := []testCase{
		{
			testName: "Test Marshal Jão",
			person: func() Person {
				var avatarColorExpected enum.AvatarColor = 1
				var stateExpected enum.UF = 1

				return Person{
					Name:        "Jão",
					Age:         uint(17),
					AvatarColor: avatarColorExpected,
					State:       stateExpected,
				}
			},
			jsonExpected: `{"name":"Jão","age":17,"avatarColorData":"blue","uf":"MG"}`,
		},
		{
			testName: "Test Marshal Marias",
			person: func() Person {
				var avatarColorExpected enum.AvatarColor = 2
				var stateExpected enum.UF = 5

				return Person{
					Name:        "Marias",
					Age:         uint(28),
					AvatarColor: avatarColorExpected,
					State:       stateExpected,
				}
			},
			jsonExpected: `{"name":"Marias","age":28,"avatarColorData":"yellow","uf":""}`,
		},
		{
			testName: "Test Marshal Zé",
			person: func() Person {
				var avatarColorExpected enum.AvatarColor = 10
				var stateExpected enum.UF = 2

				return Person{
					Name:        "Zé",
					Age:         uint(30),
					AvatarColor: avatarColorExpected,
					State:       stateExpected,
				}
			},
			jsonExpected: `{"name":"Zé","age":30,"avatarColorData":"","uf":"ES"}`,
		},
	}

	for _, test := range tests {
		t.Run(test.testName, func(t *testing.T) {
			t.Parallel()

			person := test.person()
			jsonBytes, _ := json.Marshal(person)

			assert.Equal(t, test.jsonExpected, string(jsonBytes))
		})
	}
}
