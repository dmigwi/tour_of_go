package greetings_test

import (
	"regexp"
	"testing"

	"github.com/dmigwi/tour_of_go/greetings"
)

func TestHello(t *testing.T) {
	name := "Daniel"
	want := regexp.MustCompile(`\b` + name + `\b`)

	message, err := greetings.Hello(name)
	if !want.MatchString(message) || err != nil {
		t.Errorf(`Hello("Daniel") = %q, %v, want match for %#q, nil`, message, err, want)
	}
}

// TestHelloEmpty calls greetings.Hello with an empty string,
// checking for an error.
func TestHelloEmpty(t *testing.T) {
	msg, err := greetings.Hello("")
	if msg != "" || err == nil {
		t.Errorf(`Hello("") = %q, %v, want "", error`, msg, err)
	}
}
