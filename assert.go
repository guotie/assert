package assert

import (
	"fmt"
)

func Assert(c bool, msg string) {
	if !c {
		panic(msg)
	}
}

func Assertf(c bool, format string, args ...interface{}) {
	if !c {
		panic(fmt.Sprintf(format, args))
	}
}
