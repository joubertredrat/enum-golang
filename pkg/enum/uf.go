package enum

//UF UF
type UF EnumType

//const const
const (
	UFKeySP EnumKey = iota
	UFKeyMG
	UFKeyES
	UFValueSP EnumValue = "SP"
	UFValueMG EnumValue = "MG"
	UFValueES EnumValue = "ES"
)

//GetKey GetKey
func (u UF) GetKey() int {
	return enumGetKey(EnumType(u))
}

//GetValue GetValue
func (u UF) GetValue() string {
	return enumGetValue(EnumType(u), u.GetAvailableValues())
}

//IsValid IsValid
func (u UF) IsValid() bool {
	return enumIsValid(EnumType(u), u.GetAvailableValues())
}

//GetAvailableValues GetAvailableValues
func (u UF) GetAvailableValues() map[EnumKey]EnumValue {
	return map[EnumKey]EnumValue{
		UFKeySP: UFValueSP,
		UFKeyMG: UFValueMG,
		UFKeyES: UFValueES,
	}
}

//MarshalJSON MarshalJSON
func (u UF) MarshalJSON() ([]byte, error) {
	return enumMarshalJSON(&u)
}

//UnmarshalJSON UnmarshalJSON
func (u *UF) UnmarshalJSON(b []byte) error {
	enumType := enumUnmarshalJSON(b, u.GetAvailableValues())
	*u = UF(enumType)

	return nil
}
