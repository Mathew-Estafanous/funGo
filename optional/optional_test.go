package optional

import "testing"

// Simple Test model that implements the Model interface
// to ensure there are minimal external dependencies.
type FakeModel int
func (t FakeModel) Equals(m Model) bool {
	return t == m
}

func TestOptionalOf(t *testing.T) {
	m := FakeModel(5)
	opt := OptionalOf(m)
	if opt.empty {
		t.Errorf("Created optional returned empty as false instead of true.")
	}
	if !opt.model.Equals(m) {
		t.Errorf("Created optoinal value is %v instead of %v", opt.model, m)
	}

	emptyOpt := OptionalOf(nil)
	if !emptyOpt.empty {
		t.Errorf("Created optional is not empty as expected.")
	}
	if emptyOpt.model != nil {
		t.Errorf("Created optional model value is %v instead of nil.", opt.model)
	}
}

func TestOptionalEmpty(t *testing.T) {
	o := OptionalEmpty()
	if !o.empty {
		t.Error("The created optional is not empty, as expected.")
	}

	if o.model != nil {
		t.Errorf("The created optional model is %v instead of nil.", o.model)
	}
}