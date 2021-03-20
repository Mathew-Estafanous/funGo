package model

import "testing"

func TestModelSlice_Equals(t *testing.T) {
	type test struct {
		name   string
		models [2]ModelSlice
		want   bool
	}

	table := []test{
		{
			name: "Both slices are same length and same values, should return true.",
			models: [2]ModelSlice{
				{
					ModelInt(1), ModelInt(2), ModelInt(3),
				},
				{
					ModelInt(1), ModelInt(2), ModelInt(3),
				},
			},
			want: true,
		},
		{
			name: "Both slices are of different length even with same values, should return false.",
			models: [2]ModelSlice{
				{
					ModelInt(1), ModelInt(2), ModelInt(3),
				},
				{
					ModelInt(1), ModelInt(2),
				},
			},
			want: false,
		},
		{
			name: "Both slices are same length but different values, should return false.",
			models: [2]ModelSlice{
				{
					ModelInt(1), ModelInt(2), ModelInt(3),
				},
				{
					ModelInt(1), ModelInt(2), ModelInt(4),
				},
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
