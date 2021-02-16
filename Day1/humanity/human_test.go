package humanity

import "testing"

func TestNewHumanFromCSV(t *testing.T) {
	tests := []struct {
		in  []string
		out *Human
		err bool
	}{
		{
			in:  []string{"jean", "34", "france"},
			out: &Human{Name: "jean", Age: 34, Country: "france"},
			err: false,
		},
		{
			in:  []string{},
			out: nil,
			err: true,
		},
		{
			in:  []string{"str"},
			out: nil,
			err: true,
		},
	}

	for _, test := range tests {
		human, err := NewHumanFromCSV(test.in)
		if test.err && err != nil {
			continue
		}
		if err != nil {
			t.Errorf("Error when converting line [%s]: %v", test.in, err)
			continue
		}
		if !(test.out.Name == human.Name && test.out.Age == human.Age && test.out.Country == human.Country) {
			t.Errorf("Error when converting line [%s]:\n(exp)%v != (got)%v", test.in, test.out, human)
		}
	}
}
