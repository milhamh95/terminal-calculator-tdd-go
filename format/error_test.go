package format_test

import (
	"errors"
	"terminal-calculator-tdd-go/format"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestError(t *testing.T) {
	// Arrange
	expectedErr := errors.New("error msg")
	expr := "2%3"

	// Act
	wrappedErr := format.Error(expr, expectedErr)

	// Assert.
	require.NotNil(t, wrappedErr)
	require.ErrorContains(t, wrappedErr, expectedErr.Error())
	require.ErrorContains(t, wrappedErr, expr)
}
