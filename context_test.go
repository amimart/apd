package apd

import (
	"testing"
)

func TestLn(t *testing.T) {
	var p34Context = &Context{
		Precision:   34,
		MaxExponent: MaxExponent,
		MinExponent: MinExponent,
		Traps:       DefaultTraps,
		Rounding:    RoundDown,
	}
	p, _, err := NewFromString("1.6285091944505809264504560045920167")
	if err != nil {
		panic(err)
	}
	var ln Decimal
	p34Context.Ln(&ln, p)

	if ln.String() != "0.4876649916811116824516548471782886" {
		t.Fatalf("expected '%s', got '%s'", "0.4876649916811116824516548471782886", ln.String())
	}
}
