package test

import (
	"encoding/json"
	"errors"
	"math"
	"reflect"

	"github.com/hashicorp/test/internal/constraints"
)

// Nil asserts a is nil.
func Nil(t T, a any) {
	t.Helper()

	if a != nil {
		t.Fatalf("expected to be nil; is not nil")
	}
}

// Nilf asserts a is nil, using a custom error message.
func Nilf(t T, a any, msg string, args ...any) {
	t.Helper()

	if a != nil {
		t.Fatalf(msg, args...)
	}
}

// NotNil asserts a is not nil.
func NotNil(t T, a any) {
	t.Helper()

	if a == nil {
		t.Fatalf("expected to not be nil; is nil")
	}
}

// NotNilf asserts a is not nil, using a custom error message.
func NotNilf(t T, a any, msg string, args ...any) {
	t.Helper()

	if a == nil {
		t.Fatalf(msg, args...)
	}
}

// True asserts that condition is true.
func True(t T, condition bool) {
	t.Helper()

	if !condition {
		t.Fatalf("expected condition to be true; is false")
	}
}

// Truef asserts condition is true, using a custom error message.
func Truef(t T, condition bool, msg string, args ...any) {
	t.Helper()

	if !condition {
		t.Fatalf(msg, args...)
	}
}

// False asserts condition is false.
func False(t T, condition bool) {
	t.Helper()

	if condition {
		t.Fatalf("expected condition to be false; is true")
	}
}

// Falsef asserts condition is false, using a custom error message.
func Falsef(t T, condition bool, msg string, args ...any) {
	t.Helper()

	if condition {
		t.Fatalf(msg, args...)
	}
}

// Error asserts err is a non-nil error.
func Error(t T, err error) {
	t.Helper()

	if err == nil {
		t.Fatalf("expected non-nil error; is nil")
	}
}

// ErrorIs asserts err
func ErrorIs(t T, err error, target error) {
	t.Helper()

	if !errors.Is(err, target) {
		t.Fatalf("expected %v errors.Is %v", err, target)
	}
}

// NoError asserts err is a nil error.
func NoError(t T, err error) {
	t.Helper()

	if err != nil {
		t.Fatalf("expected nil error, got %q", err)
	}
}

// Eq asserts a == b.
func Eq[C comparable](t T, a, b C) {
	t.Helper()

	if a != b {
		t.Fatalf("expected %v == %v", a, b)
	}
}

// Eqf asserts a == b, using a custom error message.
func Eqf[C comparable](t T, a, b C, msg string, args ...any) {
	t.Helper()

	if a != b {
		t.Fatalf(msg, args...)
	}
}

// NotEq asserts a != b.
func NotEq[C comparable](t T, a, b C) {
	t.Helper()

	if a == b {
		t.Fatalf("expected %v != %v", a, b)
	}
}

// NotEqf asserts a != b, using a custom error message.
func NotEqf[C comparable](t T, a, b C, msg string, args ...any) {
	t.Helper()

	if a == b {
		t.Fatalf(msg, args...)
	}
}

// EqJSON asserts a and b are equivalent JSON.
func EqJSON(t T, a, b string) {
	t.Helper()

	var expA, expB interface{}

	if err := json.Unmarshal([]byte(a), &expA); err != nil {
		t.Fatalf("failed to unmarshal first argument as json: %v", err)
	}

	if err := json.Unmarshal([]byte(b), &expB); err != nil {
		t.Fatalf("failed to unmarshal second argument as json: %v", err)
	}

	if !reflect.DeepEqual(expA, expB) {
		jsonA, _ := json.Marshal(expA)
		jsonB, _ := json.Marshal(expB)
		t.Fatalf("json strings are not the same; %s vs. %s", jsonA, jsonB)
	}
}

// Equals asserts a.Equals(b).
func Equals[E EqualsFunc[E]](t T, a, b E) {
	t.Helper()

	if !a.Equals(b) {
		t.Fatalf("expected to be equal: %v, %v", a, b)
	}
}

// Equalsf asserts a.Equal(b), using a custom error message.
func Equalsf[E EqualsFunc[E]](t T, a, b E, msg string, args ...any) {
	t.Helper()

	if !a.Equals(b) {
		t.Fatalf(msg, args...)
	}
}

// NotEquals asserts !a.Equals(b).
func NotEquals[E EqualsFunc[E]](t T, a, b E) {
	t.Helper()

	if a.Equals(b) {
		t.Fatalf("expected to be not equal: %v, %v", a, b)
	}
}

// NotEqualsf asserts !a.Equals(b), using a custom error message.
func NotEqualsf[E EqualsFunc[E]](t T, a, b E, msg string, args ...any) {
	t.Helper()

	if a.Equals(b) {
		t.Fatalf(msg, args...)
	}
}

// Lesser asserts a.Less(b).
func Lesser[L LessFunc[L]](t T, a, b L) {
	t.Helper()

	if !a.Less(b) {
		t.Fatalf("expected to be less; %v, %v", a, b)
	}
}

// Empty asserts slice is empty.
func Empty[A any](t T, slice []A) {
	t.Helper()

	if len(slice) != 0 {
		t.Fatalf("expected slice to be empty; is len %d", len(slice))
	}
}

// Len asserts slice is of length n.
func Len[A any](t T, n int, slice []A) {
	t.Helper()

	l := len(slice)
	if l != n {
		t.Fatalf("expected slice to be length %d; is %d", n, l)
	}
}

// Lenf asserts slice is of length n, using a custom error message.
func Lenf[A any](t T, n int, slice []A, msg string, args ...any) {
	t.Helper()

	l := len(slice)
	if l != n {
		t.Fatalf(msg, args...)
	}
}

func contains[C comparable](slice []C, item C) bool {
	found := false
	for i := 0; i < len(slice); i++ {
		if slice[i] == item {
			found = true
			break
		}
	}
	return found
}

// Contains asserts item exists in slice.
func Contains[C comparable](t T, slice []C, item C) {
	t.Helper()

	if !contains(slice, item) {
		t.Fatalf("expected slice to contain %v but does not: %v", item, slice)
	}
}

// Containsf asserts item exists in slice, using a custom error message.
func Containsf[C comparable](t T, slice []C, item C, msg string, args ...any) {
	t.Helper()

	if !contains(slice, item) {
		t.Fatalf(msg, args...)
	}
}

// Less asserts a < b.
func Less[O constraints.Ordered](t T, a, b O) {
	t.Helper()

	if !(a < b) {
		t.Fatalf("expected %v < %v", a, b)
	}
}

// Lessf asserts a < b, using a custom error message.
func Lessf[O constraints.Ordered](t T, a, b O, msg string, args ...any) {
	t.Helper()

	if !(a < b) {
		t.Fatalf(msg, args...)
	}
}

// LessEq asserts a <= b.
func LessEq[O constraints.Ordered](t T, a, b O) {
	t.Helper()

	if !(a <= b) {
		t.Fatalf("expected %v <= %v", a, b)
	}
}

// LessEqf asserts a <= b, using a custom error message.
func LessEqf[O constraints.Ordered](t T, a, b O, msg string, args ...any) {
	t.Helper()

	if !(a <= b) {
		t.Fatalf(msg, args...)
	}
}

// Greater asserts a > b.
func Greater[O constraints.Ordered](t T, a, b O) {
	t.Helper()

	if !(a > b) {
		t.Fatalf("expected %v > %v", a, b)
	}
}

// Greaterf asserts a > b, using a custom error message.
func Greaterf[O constraints.Ordered](t T, a, b O, msg string, args ...any) {
	t.Helper()

	if !(a > b) {
		t.Fatalf(msg, args...)
	}
}

// GreaterEq asserts a >= b.
func GreaterEq[O constraints.Ordered](t T, a, b O) {
	t.Helper()

	if !(a >= b) {
		t.Fatalf("expected %v >= %v", a, b)
	}
}

// GreaterEqf asserts a >= b, using a custom error message.
func GreaterEqf[O constraints.Ordered](t T, a, b O, msg string, args ...any) {
	t.Helper()

	if !(a >= b) {
		t.Fatalf(msg, args...)
	}
}

// Number is float, integer, or complex.
type Number interface {
	constraints.Ordered
	constraints.Float | constraints.Integer | constraints.Complex
}

// Numeric returns false if n is Inf/NaN.
//
// Always returns true for integral values.
func Numeric[N Number](n N) bool {
	check := func(f float64) bool {
		if math.IsNaN(f) {
			return false
		} else if math.IsInf(f, 0) {
			return false
		}
		return true
	}
	return check(float64(n))
}

// InDelta asserts a and b are within delta of each other.
func InDelta[N Number](t T, a, b, delta N) {
	t.Helper()

	var zero N

	if !Numeric(delta) {
		t.Fatalf("delta must be numeric; got %v", delta)
	}

	if delta <= zero {
		t.Fatalf("delta must be positive; got %v", delta)
	}

	if !Numeric(a) {
		t.Fatalf("first argument must be numeric; got %v", a)
	}

	if !Numeric(b) {
		t.Fatalf("second argument must be numeric; got %v", b)
	}

	difference := a - b
	if difference < -delta || difference > delta {
		t.Fatalf("%v and %v not within %v", a, b, delta)
	}
}

// InDeltaSlice asserts each element a[n] is within delta of b[n].
func InDeltaSlice[N Number](t T, a, b []N, delta N) {
	t.Helper()

	if len(a) != len(b) {
		t.Fatalf("slices not of same length; %d != %d", len(a), len(b))
	}

	for i := 0; i < len(a); i++ {
		InDelta(t, a[i], b[i], delta)
	}
}

// MapEq asserts a contains the same key:value pairs as b.
func MapEq[K comparable, V any](t T, a, b map[K]V) {
	t.Helper()

	if len(a) != len(b) {
		t.Fatalf("maps are different size; %d vs %d; %v != %v", len(a), len(b), a, b)
	}
	for key, valueA := range a {
		valueB, exists := b[key]
		if !exists {
			t.Fatalf("map keys are different; %v != %v", a, b)
		}

		if !reflect.DeepEqual(valueA, valueB) {
			t.Fatalf("value for key %v different; %v != %v", key, a, b)
		}
	}
}

// MapLen asserts map is of size n.
func MapLen[K comparable, V any](t T, n int, m map[K]V) {
	t.Helper()

	s := len(m)
	if s != n {
		t.Fatalf("expected map to be size %d; is %d", n, s)
	}
}

// MapLenf asserts map is of size n, using a custom error message.
func MapLenf[K comparable, V any](t T, n int, m map[K]V, msg string, args ...any) {
	t.Helper()

	s := len(m)
	if s != n {
		t.Fatalf(msg, args...)
	}
}

// MapEmpty asserts map is empty.
func MapEmpty[K comparable, V any](t T, m map[K]V) {
	t.Helper()

	if l := len(m); l > 0 {
		t.Fatalf("expected map to be empty; is size %d", l)
	}
}
