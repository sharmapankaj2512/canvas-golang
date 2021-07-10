package strings

import (
	"github.com/lithammer/dedent"
	"strings"
)

func TrimMargin(s string) string {
	return strings.TrimSpace(dedent.Dedent(s))
}
