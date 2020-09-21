package enum

import (
	"testing"

	"gopkg.in/go-playground/assert.v1"
)

func TestAvatarColor(t *testing.T) {
	type testCase struct {
		testName                string
		enum                    func() AvatarColor
		keyExpected             int
		valueExpected           string
		isValidExpected         bool
		availableValuesExpected map[EnumKey]EnumValue
	}

	tests := []testCase{
		{
			testName: "Test valid AvatarColor enum",
			enum: func() AvatarColor {
				var a AvatarColor = 0
				return a
			},
			keyExpected:     1,
			valueExpected:   "MG",
			isValidExpected: true,
			availableValuesExpected: map[EnumKey]EnumValue{
				AvatarColorKeyRed:    AvatarColorValueRed,
				AvatarColorKeyBlue:   AvatarColorValueBlue,
				AvatarColorKeyYellow: AvatarColorValueYellow,
				AvatarColorKeyGreen:  AvatarColorValueGreen,
			},
		},
		{
			testName: "Test invalid AvatarColor enum",
			enum: func() AvatarColor {
				var a AvatarColor = 10
				return a
			},
			keyExpected:     10,
			valueExpected:   "",
			isValidExpected: false,
			availableValuesExpected: map[EnumKey]EnumValue{
				AvatarColorKeyRed:    AvatarColorValueRed,
				AvatarColorKeyBlue:   AvatarColorValueBlue,
				AvatarColorKeyYellow: AvatarColorValueYellow,
				AvatarColorKeyGreen:  AvatarColorValueGreen,
			},
		},
	}

	for _, test := range tests {
		t.Run(test.testName, func(t *testing.T) {
			t.Parallel()

			uf := test.enum()

			assert.Equal(t, test.keyExpected, uf.GetKey())
			assert.Equal(t, test.valueExpected, uf.GetValue())
			assert.Equal(t, test.isValidExpected, uf.IsValid())
			assert.Equal(t, test.availableValuesExpected, uf.GetAvailableValues())
		})
	}
}
