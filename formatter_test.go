package waterwheel

import (
	"fmt"
	"testing"
	"time"
)

func TestSimpleFormatter(t *testing.T) {
	t.Parallel()

	buf := []byte{}
	tm := time.Now()
	r := &Record{
		Level:   Info,
		Time:    tm,
		Message: "TestSimpleFormatter",
	}

	SimpleFormatter(r, &buf)
	exp := fmt.Sprintf("INFO     %s.%06d TestSimpleFormatter\n", tm.UTC().Format("2006-01-02T15:04:05"), tm.Nanosecond()/1e3)
	if string(buf) != exp {
		t.Fatalf("mismatched formatting: expected %s, got %s", exp, string(buf))
	}
}
