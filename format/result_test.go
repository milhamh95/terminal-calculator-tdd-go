package format_test

import (
	"fmt"
	"terminal-calculator-tdd-go/format"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestResult(t *testing.T) {
	// Arrange
	result := 5.55
	expr := "2%3"

	// Act
	wrappedRes := format.Result(expr, result)

	// Assert.
	require.Contains(t, wrappedRes, expr)
	require.Contains(t, wrappedRes, fmt.Sprint(result))
}
