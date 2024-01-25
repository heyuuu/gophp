package standard

import (
	"fmt"
	"github.com/heyuuu/gophp/php/zpp"
	"math/rand"
	"time"
)

func ZifUniqid(_ zpp.Opt, prefix string, moreEntropy bool) string {
	now := time.Now()
	sec := now.Second()
	usec := now.Nanosecond() % int(time.Second) / int(time.Microsecond)

	if moreEntropy {
		return fmt.Sprintf("%s%08x%05x%.8f", prefix, sec, usec, rand.Float64())
	} else {
		return fmt.Sprintf("%s%08x%05x", prefix, sec, usec)
	}
}
