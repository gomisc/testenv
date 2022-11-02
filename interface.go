package testenv

import (
	"fmt"

	"git.eth4.dev/golibs/deps"
	"git.eth4.dev/golibs/execs"
)

type (
	EnvHandler func(ctx deps.ContainersAdapter)

	Component func(ctx deps.ContainersAdapter, image string, background bool) ComponentOption

	ComponentOption interface {
		fmt.Stringer
		Prepare() (execs.Member, error)
	}

	Environment interface {
		Context() deps.ContainersAdapter
		BeforeRun(EnvHandler)
		Close()
	}
)
