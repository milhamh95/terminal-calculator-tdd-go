package input_test

import (
	"fmt"
	"github.com/stretchr/testify/require"
	"terminal-calculator-tdd-go/calculator"
	"terminal-calculator-tdd-go/input"
	"terminal-calculator-tdd-go/mocks"
	"testing"
)

func TestParser(t *testing.T) {
	t.Run("valid input", func(t *testing.T) {
		//Arrange
		expr := "2 + 3"
		operator := "+"
		operands := []float64{2.0, 3.0}
		expectedResult := "2 + 3 = 5.5"

		engine := mocks.NewOperationProcessor(t)
		engine.On("ProcessOperation", &calculator.Operation{
			Expression: expr,
			Operator:   operator,
			Operands:   operands,
		})

		validator := mocks.NewValidationHelper(t)
		validator.On("CheckInput", operator, operands).Return(nil).Once()

		parser := input.NewParser(engine, validator)

		// Act
		result, err := parser.ProcessExpression(expr)

		// Assert
		require.Nil(t, err)
		require.NotNil(t, result)
		require.Contains(t, *result, expectedResult)
		require.Contains(t, *result, expr)

		validator.AssertExpectations(t)
		engine.AssertExpectations(t)

	})

	t.Run("invalid operator", func(t *testing.T) {
		//Arrange
		expr := "2 + 3"
		operator := "+"
		operands := []float64{2.0, 3.0}
		expectedErrMsg := "bad operator"

		engine := mocks.NewOperationProcessor(t)

		validator := mocks.NewValidationHelper(t)
		validator.On("CheckInput", operator, operands).Return(fmt.Errorf(expectedErrMsg)).Once()

		parser := input.NewParser(engine, validator)

		// Act
		result, err := parser.ProcessExpression(expr)

		// Assert
		require.NotNil(t, err)
		require.Nil(t, result)
		require.Contains(t, err.Error(), expectedErrMsg)
		require.Contains(t, err.Error(), expr)

		validator.AssertExpectations(t)
	})
}
