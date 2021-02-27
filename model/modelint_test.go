package model

import "testing"

func TestModelInt_Equals(t *testing.T) {
	type test struct {
		name string
		models [2]ModelInt
		want bool
	}

	table := []test {
		{
			name: "Check if two values that are equal return true",
			models: [2]ModelInt{
				ModelInt(5),
				ModelInt(5),
			},
			want: true,
		},
		{
			name: "Two different values that are not equal should return false",
			models: [2]ModelInt{
				ModelInt(6),
				ModelInt(1),
			},
			want: false,
		},
	}

	for _, te := range table {
		model1, model2 := te.models[0], te.models[1]
		if model1.Equals(model2); !te.want {
			t.Errorf(te.name)
		}
	}
}
