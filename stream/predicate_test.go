package stream

import (
	. "github.com/Mathew-Estafanous/funGo/model"
	"testing"
)

func TestPredicate_And(t *testing.T) {
	type test struct {
		name string
		pred [2]Predicate
		want bool
	}

	andTests := []test {
		{
			name: "Both predicates return true, which means true is expected.",
			pred: [2]Predicate{
				func(m Model) bool {
					return  true
				},
				func(m Model) bool {
					return true
				},
			},
			want: true,
		},
		{
			name: "One predicate returns false and other true, which means false is expected.",
			pred: [2]Predicate{
				func(m Model) bool {
					return true
				},
				func(m Model) bool {
					return false
				},
			},
			want: false,
		},
		{
			name: "Both predicates return false, which means false is expected.",
			pred: [2]Predicate{
				func(m Model) bool {
					return false
				},
				func(m Model) bool {
					return false
				},
			},
			want: false,
		},
	}

	for _, te := range andTests {
		andPred := te.pred[0].And(te.pred[1])
		if andPred(nil) != te.want {
			t.Error(te.name)
		}
	}
}

func TestPredicate_Not(t *testing.T) {
	type test struct {
		name string
		pred Predicate
		want bool
	}

	notTests := []test {
		{
			name: "Given predicate returns false, which means true is expected.",
			pred: func(m Model) bool {
				return false
			},
			want: true,
		},
		{
			name: "Given predicate returns true, which means false is expected.",
			pred: func(m Model) bool {
				return true
			},
			want: false,
		},
	}

	for _, te := range notTests {
		notPred := te.pred.Not()
		if notPred(nil) != te.want {
			t.Error(te.name)
		}
	}
}

func TestPredicate_Or(t *testing.T) {
	type test struct {
		name string
		pred [2]Predicate
		want bool
	}

	orTests := []test {
		{
			name: "Both predicates return true, which means true is expected.",
			pred: [2]Predicate{
				func(m Model) bool {
					return  true
				},
				func(m Model) bool {
					return true
				},
			},
			want: true,
		},
		{
			name: "One predicate returns false and other true, which means true is expected.",
			pred: [2]Predicate{
				func(m Model) bool {
					return true
				},
				func(m Model) bool {
					return false
				},
			},
			want: true,
		},
		{
			name: "Both predicates return false, which means false is expected.",
			pred: [2]Predicate{
				func(m Model) bool {
					return false
				},
				func(m Model) bool {
					return false
				},
			},
			want: false,
		},
	}

	for _, te := range orTests {
		orPred := te.pred[0].Or(te.pred[1])
		if orPred(nil) != te.want {
			t.Error(te.name)
		}
	}
}