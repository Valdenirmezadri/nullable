package nullable

// Default holds what every nullable type shares: the value and whether it is set
type Default[T comparable] struct {
	realValue T
	isValid   bool
}

// Get either zero value or value
func (n Default[T]) Get() T {
	return n.realValue
}

// IsNull reports whether the value is null
func (n Default[T]) IsNull() bool {
	return !n.isValid
}

// Set either zero value (null) or value. New, Scan and UnmarshalJSON all go
// through here, so 0/""/zero time is always saved as NULL.
func (n *Default[T]) Set(value T) {
	var zero T
	n.isValid = !isZero(value)
	n.realValue = zero
	if n.isValid {
		n.realValue = value
	}
}

// isZero prefers IsZero() when the type has one: time.Time == compares the location too
func isZero[T comparable](value T) bool {
	if z, ok := any(value).(interface{ IsZero() bool }); ok {
		return z.IsZero()
	}
	var zero T
	return value == zero
}
