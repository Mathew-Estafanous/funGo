package model

import "testing"

func TestModelFloat_Equals(t *testing.T) {
	type test struct {
		name   string
		models [2]ModelFloat
		want   bool
	}

	table := []test{
		{
			name: "Check if two values that are equal return true",
			models: [2]ModelFloat{
				ModelFloat(5),
				ModelFloat(5),
			},
			want: true,
		},
		{
			name: "Two different values that are not equal should return false",
			models: [2]ModelFloat{
				ModelFloat(6),
				ModelFloat(1),
			},
			want: false,
		},
	}

	for _, te := range table {
		model1, model2 := te.models[0], te.models[1]
		if model1.Equals(model2) != te.want {
			t.Error(te.name)
		}
	}
}
