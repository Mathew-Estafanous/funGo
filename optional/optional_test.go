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

func TestOptional_Get(t *testing.T) {
	m := FakeModel(5)
	opt := Optional{model: m, empty: false}

	if result, _ := opt.Get(); !result.Equals(m) {
		t.Errorf("Received result value exected %v but received %v", m, result)
	}

	emptyOpt := Optional{model: nil, empty: true}
	if _, err := emptyOpt.Get(); err == nil {
		t.Error("Received no error when getting from an empty optional.")
	}
}

func TestOptional_IsEmpty(t *testing.T) {
	emptyOpt := Optional{model: nil, empty: true}
	if !emptyOpt.IsEmpty() {
		t.Error("Empty optional returned false for empty when it should be true.")
	}
}

func TestOptional_GetOrElse(t *testing.T) {
	m := FakeModel(5)
	opt := Optional{model: m, empty: false}
	if result := opt.GetOrElse(nil); !result.Equals(m) {
		t.Errorf("Received %v instead of %v as the Get model.", result, m)
	}

	emptyOpt := Optional{model: nil, empty: true}
	if result := emptyOpt.GetOrElse(m); !result.Equals(m) {
		t.Errorf("Received %v instead of %v as the OrElse model.", result, m)
	}
}

func TestOptional_IfNotPresent(t *testing.T) {
	m := FakeModel(5)
	opt := Optional{model: m, empty: false}
	opt.IfNotPresent(func() {
		t.Errorf("Optional called func when optional was not-empty.")
	})

	called := false
	emptyOpt := Optional{model: nil, empty: true}
	emptyOpt.IfNotPresent(func() {
		called = true
	})
	if !called {
		t.Errorf("Optional didn't call func when optional was empty.")
	}
}

func TestOptional_IfPresent(t *testing.T) {
	called := false
	m := FakeModel(5)
	opt := Optional{model: m, empty: false}
	opt.IfPresent(func(value Model) {
		called = true
	})
	if !called {
		t.Errorf("Optional didn't called func when optional wasn't empty.")
	}

	called = false
	emptyOpt := Optional{model: nil, empty: true}
	emptyOpt.IfPresent(func(value Model) {
		t.Errorf("Optional called function when it was empty.")
	})
}