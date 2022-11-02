package testenv

import (
	"fmt"
	"strings"

	"git.eth4.dev/golibs/coders/base58"
	"git.eth4.dev/golibs/coders/uuid"
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
