package model

import (
	"testing"
)

func TestModelMap_Equals(t *testing.T) {
	type test struct {
		name string
		models [2]ModelMap
		want bool
	}

	table := []test {
		{
			name: "Both ModelMaps are equal and should return true.",
			models: [2]ModelMap{
				{
					ModelInt(1): ModelInt(1),
					ModelInt(2): ModelInt(2),
				},
				{
					ModelInt(1): ModelInt(1),
					ModelInt(2): ModelInt(2),
				},
			},
			want: true,
		},
		{
			name: "Both ModelMaps have same key but different values, should return false.",
			models: [2]ModelMap{
				{
					ModelInt(1): ModelInt(1),
					ModelInt(2): ModelInt(2),
				},
				{
					ModelInt(1): ModelInt(2),
					ModelInt(2): ModelInt(3),
				},
			},
			want: false,
		},
		{
			name: "Both ModelMaps have different keys and same values, should return false.",
			models: [2]ModelMap{
				{
					ModelInt(1): ModelInt(1),
					ModelInt(2): ModelInt(2),
				},
				{
					ModelInt(2): ModelInt(1),
					ModelInt(3): ModelInt(2),
				},
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