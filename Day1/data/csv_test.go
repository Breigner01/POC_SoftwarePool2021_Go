package data

import "testing"

func TestLineToCSV(t *testing.T) {
	tests := []struct {
		in  string
		out []string
		err bool
	}{
		{
			in:  "jean,34,france",
			out: []string{"jean", "34", "france"},
			err: false,
		},
		{
			in:  "j,j,j,j",
			out: nil,
			err: true,
		},
		{
			in:  "j",
			out: nil,
			err: true,
		},
	}

	for _, test := range tests {
		csv, err := LineToCSV(test.in)
		if test.err && err != nil {
			continue
		}
		if err != nil {
			t.Errorf("Error when converting line [%s]: %v", test.in, err)
			continue
		}
		for i, v := range test.out {
			if v != csv[i] {
				t.Errorf("Error when converting line [%s]:\n(exp)%v != (got)%v", test.in, test.out, csv)
			}
		}
	}
}

func TestReadFileCSV(t *testing.T) {
	tests := []struct {
		in  string
		out []string
		err bool
	}{
		{
			in: "../resources/example.csv",
			out: []string{"Damien,27,France",
				"Tom,5,USA",
				"QuentinV,8,Epicat",
				"Ugo,17,LeCodeDegueu",
				"QuentinF,23,Ukraine",
				"Louis,90,USA",
				"Gregoire,8,Epicat",
				"Ugo,17,LeCodeDegueu"},
			err: false,
		},
		{
			in:  "j,j,j,j",
			out: nil,
			err: true,
		},
		{
			in:  "j",
			out: nil,
			err: true,
		},
	}

	for _, test := range tests {
		csv, err := ReadFileCSV(test.in)
		if test.err && err != nil {
			continue
		}
		if err != nil {
			t.Errorf("Error when converting line [%s]: %v", test.in, err)
			continue
		}
		for i, v := range test.out {
			if v != csv[i] {
				t.Errorf("Error when converting line [%s]:\n(exp)%v != (got)%v", test.in, test.out, csv)
			}
		}
	}
}
