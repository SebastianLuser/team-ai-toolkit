package errors

import (
	"fmt"
	"testing"
)

func TestSentinelErrors_AreDistinct(t *testing.T) {
	errs := []error{ErrNotFound, ErrValidation, ErrUnauthorized, ErrForbidden, ErrDuplicate, ErrConflict}

	for i, a := range errs {
		for j, b := range errs {
			if i != j && Is(a, b) {
				t.Errorf("errors should be distinct: %v == %v", a, b)
			}
		}
	}
}

func TestWrappedError_IsDetected(t *testing.T) {
	wrapped := fmt.Errorf("repo.GetUser(id=42): %w", ErrNotFound)

	if !Is(wrapped, ErrNotFound) {
		t.Error("wrapped error should match ErrNotFound")
	}
	if Is(wrapped, ErrValidation) {
		t.Error("wrapped error should not match ErrValidation")
	}
}

func TestDoubleWrapped_IsDetected(t *testing.T) {
	inner := fmt.Errorf("db query failed: %w", ErrNotFound)
	outer := fmt.Errorf("usecase.GetDocument: %w", inner)

	if !Is(outer, ErrNotFound) {
		t.Error("double-wrapped error should match ErrNotFound")
	}
}
