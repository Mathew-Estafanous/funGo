package model

import "testing"

func TestModelByte_Equals(t *testing.T) {
	type test struct {
		name string
		models [2]ModelByte
		want bool
	}

	table := []test {
		{
			name: "Both bytes are equal and should return true.",
			models: [2]ModelByte{
				ModelByte('a'),
				ModelByte('a'),
			},
			want: true,
		},
		{
			name: "Both bytes are different and should return false.",
			models: [2]ModelByte{
				ModelByte('a'),
				ModelByte('b'),
			},
			want: false,
		},
	}

	for _, te := range table {
		model1, model2 := te.models[0], te.models[1]
		if model1.Equals(model2) != te.want {
			t.Errorf(te.name)
		}
	}
}