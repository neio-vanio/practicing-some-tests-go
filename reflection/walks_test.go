package reflection

import (
	"reflect"
	"testing"
)

func TestWalks(t *testing.T) {

	cases := []struct {
		Name     string
		Input    interface{}
		Expected []string
	}{
		{
			"Struct with a field string",
			struct {
				Name string
				Nick string
			}{"Almeida", "Carlos"},
			[]string{"Almeida", "Carlos"},
		},
		{
			"Struct with a int",
			struct {
				Name  string
				Idade int
			}{"Vanio", 32},
			[]string{"Vanio"},
		},
		{
			"Campos aninhados",
			Person{
				"Chris",
				Profile{33, "Londres"},
			},
			[]string{"Chris", "Londres"},
		},
		{
			"Campos aninhados",
			&Person{
				"Chris",
				Profile{33, "Londres"},
			},
			[]string{"Chris", "Londres"},
		},
		{
			"Slices",
			[]Profile{
				{33, "Londres"},
				{34, "Reykjavík"},
			},
			[]string{"Londres", "Reykjavík"},
		},
	}

	for _, test := range cases {
		t.Run(test.Name, func(t *testing.T) {

			var result []string

			walks(test.Input, func(arg string) {
				result = append(result, arg)
			})

			if !reflect.DeepEqual(result, test.Expected) {
				t.Errorf("result '%v', expected '%v'", result, test.Expected)
			}

		})
	}
}
