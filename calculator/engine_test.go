package calculator_test

import (
	"log"
	"os"
	"terminal-calculator-tdd-go/calculator"
	"testing"

	"github.com/stretchr/testify/require"
)

// main method will be run before any of other test in this package.
// we're only able to define one TestMain function per package
func TestMain(m *testing.M) {
	setup()

	e := m.Run()

	teardown()

	os.Exit(e)
}

func setup() {
	log.Println("setting up")
}

func teardown() {
	log.Println("tearing down")
}

func init() {
	log.Println("init setup")
}

func TestAdd(t *testing.T) {
	defer func() {
		log.Println("deferred tearing down")
	}()

	ce := calculator.Engine{}

	actAssert := func(x, y, want float64) {
		// act
		got := ce.Add(x, y)

		// Assert
		require.Equal(t, want, got)
	}

	t.Run("success", func(t *testing.T) {
		// Arrange
		expectedResult := 2.0

		// Act
		res := ce.Add(1.0, 1.0)

		// Assert
		require.Equal(t, expectedResult, res)
	})

	t.Run("negative number", func(t *testing.T) {
		x := -1.0
		y := -2.0

		actAssert(x, y, -3.0)
	})
}

func BenchmarkAdd(b *testing.B) {
	ce := calculator.Engine{}

	for i := 0; i < b.N; i++ {
		ce.Add(1.0, 4.0)
	}
}
