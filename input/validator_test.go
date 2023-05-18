package input_test

import (
	"terminal-calculator-tdd-go/input"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCheckInput(t *testing.T) {
	validOperators := []string{"+"}
	t.Run("valid case", func(t *testing.T) {
		v := setup(t, validOperators)
		err := v.CheckInput(validOperators[0], []float64{2.5, 3.5})
		require.NoError(t, err)
	})

	t.Run("invalid operator", func(t *testing.T) {
		v := setup(t, validOperators)
		err := v.CheckInput("-", []float64{2.5, 3.5})
		require.NotNil(t, err)
	})

	t.Run("invalid operand count", func(t *testing.T) {
		v := setup(t, validOperators)
		err := v.CheckInput("-", []float64{2.5})
		require.NotNil(t, err)
	})
}

func setup(t *testing.T, validOps []string) *input.Validator {
	t.Helper()
	return input.NewValidator(2, validOps)
}
