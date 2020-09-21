package enum

import (
	"testing"

	"gopkg.in/go-playground/assert.v1"
)

func TestUF(t *testing.T) {
	type testCase struct {
		testName                string
		enum                    func() UF
		keyExpected             int
		valueExpected           string
		isValidExpected         bool
		availableValuesExpected map[EnumKey]EnumValue
	}

	tests := []testCase{
		{
			testName: "Test valid UF enum",
			enum: func() UF {
				var uf UF = 1
				return uf
			},
			keyExpected:     1,
			valueExpected:   "MG",
			isValidExpected: true,
			availableValuesExpected: map[EnumKey]EnumValue{
				UFKeySP: UFValueSP,
				UFKeyMG: UFValueMG,
				UFKeyES: UFValueES,
			},
		},
		{
			testName: "Test invalid UF enum",
			enum: func() UF {
				var uf UF = 5
				return uf
			},
			keyExpected:     5,
			valueExpected:   "",
			isValidExpected: false,
			availableValuesExpected: map[EnumKey]EnumValue{
				UFKeySP: UFValueSP,
				UFKeyMG: UFValueMG,
				UFKeyES: UFValueES,
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
