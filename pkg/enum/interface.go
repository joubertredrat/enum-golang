package enum

import (
	"bytes"
	"encoding/json"
)

//EnumType EnumType
type EnumType int

//EnumKey EnumKey
type EnumKey int

//EnumValue EnumValue
type EnumValue string

//Enum the super enum interface
type Enum interface {
	GetKey() int
	GetValue() string
	IsValid() bool
	GetAvailableValues() map[EnumKey]EnumValue
	MarshalJSON() ([]byte, error)
	UnmarshalJSON(b []byte) error
}

func enumGetKey(e EnumType) int {
	return int(e)
}

func enumGetValue(e EnumType, l map[EnumKey]EnumValue) string {
	if !enumIsValid(e, l) {
		return ""
	}

	enumKey := EnumKey(e)
	enumValue := l[enumKey]

	return string(enumValue)
}

func enumIsValid(e EnumType, l map[EnumKey]EnumValue) bool {
	enumKey := EnumKey(e)
	_, ok := l[enumKey]

	return ok
}

func enumMarshalJSON(e Enum) ([]byte, error) {
	buffer := bytes.NewBufferString(`"`)
	buffer.WriteString(e.GetValue())
	buffer.WriteString(`"`)
	return buffer.Bytes(), nil
}

func enumUnmarshalJSON(b []byte, l map[EnumKey]EnumValue) EnumType {
	var jsonValue string
	err := json.Unmarshal(b, &jsonValue)

	if err != nil {
		return EnumType(-1)
	}

	jsonEnumValue := EnumValue(jsonValue)

	for key, value := range l {
		if jsonEnumValue == value {
			return EnumType(key)
		}
	}

	return EnumType(-1)
}
