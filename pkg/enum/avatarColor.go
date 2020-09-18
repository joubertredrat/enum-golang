package enum

//AvatarColor AvatarColor
type AvatarColor EnumType

//const const
const (
	AvatarColorKeyRed EnumKey = iota
	AvatarColorKeyBlue
	AvatarColorKeyYellow
	AvatarColorKeyGreen
	AvatarColorValueRed    EnumValue = "red"
	AvatarColorValueBlue   EnumValue = "blue"
	AvatarColorValueYellow EnumValue = "yellow"
	AvatarColorValueGreen  EnumValue = "green"
)

//GetKey GetKey
func (a AvatarColor) GetKey() int {
	return enumGetKey(EnumType(a))
}

//GetValue GetValue
func (a AvatarColor) GetValue() string {
	return enumGetValue(EnumType(a), a.GetAvailableValues())
}

//IsValid IsValid
func (a AvatarColor) IsValid() bool {
	return enumIsValid(EnumType(a), a.GetAvailableValues())
}

//GetAvailableValues GetAvailableValues
func (a AvatarColor) GetAvailableValues() map[EnumKey]EnumValue {
	return map[EnumKey]EnumValue{
		AvatarColorKeyRed:    AvatarColorValueRed,
		AvatarColorKeyBlue:   AvatarColorValueBlue,
		AvatarColorKeyYellow: AvatarColorValueYellow,
		AvatarColorKeyGreen:  AvatarColorValueGreen,
	}
}

//MarshalJSON MarshalJSON
func (a AvatarColor) MarshalJSON() ([]byte, error) {
	return enumMarshalJSON(&a)
}

//UnmarshalJSON UnmarshalJSON
func (a *AvatarColor) UnmarshalJSON(b []byte) error {
	enumType := enumUnmarshalJSON(b, a.GetAvailableValues())
	*a = AvatarColor(enumType)

	return nil
}
