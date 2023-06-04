package testenv

import (
	"fmt"
	"strings"

	"gopkg.in/gomisc/coders.v1/base58"
	"gopkg.in/gomisc/coders.v1/uuid"
)

func DefaultNamer(prefix ...string) string {
	return strings.Join(
		append(prefix, base58.Encode(uuid.UUIDv4())), "-",
	)
}

func FormatStdoutPrefix(color, name string) string {
	return fmt.Sprintf("\x1b[32m[o]\x1b[%s[%s]\x1b[0m ", color, name)
}

func FormatStderrPrefix(color, name string) string {
	return fmt.Sprintf("\x1b[91m[e]\x1b[%s[%s]\x1b[0m ", color, name)
}
