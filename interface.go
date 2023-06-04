package testenv

import (
	"fmt"

	"gopkg.in/gomisc/execs.v1"
)

type (
	EnvHandler func(ctx Context)

	Component func(ctx Context, image string, background bool) ComponentOption

	ComponentOption interface {
		fmt.Stringer
		Prepare() (execs.Member, error)
	}

	Environment interface {
		Context() Context
		BeforeRun(EnvHandler)
		Close()
	}

	DeferFunc func(ctx Context)
)
