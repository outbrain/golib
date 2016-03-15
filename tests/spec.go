package tests

import (
	"testing"
)

// Spec is an access point to test assertions
type Spec struct {
	t *testing.T
}

// TestingSpec generates a spec. You will want to use it once in a test file
func TestingSpec(t *testing.T) *Spec {
	return &Spec{t: t}
}

// AssertNil expects given value to be nil, or errors
func (spec *Spec) AssertNil(actual interface{}) {
	if actual == nil {
		return
	}
	spec.t.Error("Expected %+v to be nil", actual)
}

// AssertNotNil expects given value to be not nil, or errors
func (spec *Spec) AssertNotNil(actual interface{}) {
	if actual != nil {
		return
	}
	spec.t.Error("Expected %+v to be not nil", actual)
}

// AssertEquals expects given values to be equal (comparison via `==`), or errors
func (spec *Spec) AssertEquals(actual, value interface{}) {
	if actual == value {
		return
	}
	spec.t.Error("Expected %+v, got %+v", value, actual)
}

// AssertNotEquals expects given values to be nonequal (comparison via `==`), or errors
func (spec *Spec) AssertNotEquals(actual, value interface{}) {
	if !(actual == value) {
		return
	}
	spec.t.Error("Expected not %+v", value)
}

// AssertEqualsAny expects given actual to equal (comparison via `==`) at least one of given values, or errors
func (spec *Spec) AssertEqualsAny(actual interface{}, values ...interface{}) {
	for _, value := range values {
		if actual == value {
			return
		}
	}
	spec.t.Error("Expected %+v to equal any of given values", actual)
}

// AssertNotEqualsAny expects given actual to be nonequal (comparison via `==`)tp any of given values, or errors
func (spec *Spec) AssertNotEqualsAny(actual interface{}, values ...interface{}) {
	for _, value := range values {
		if actual == value {
			spec.t.Error("Expected not %+v", value)
		}
	}
}
