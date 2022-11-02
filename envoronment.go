package testenv

import (
	"context"
	"fmt"
	"os"
	"sort"
	"strings"
	"syscall"
	"time"

	"git.eth4.dev/golibs/deps"
	"git.eth4.dev/golibs/execs"
)

const (
	DefaultStartTimeout = time.Minute
)

type testEnvironment struct {
	ctx             deps.ContainersAdapter
	infra, services []ComponentOption
}

func (env *testEnvironment) Context() deps.ContainersAdapter {
	return env.ctx
}

func (env *testEnvironment) BeforeRun(handler EnvHandler) {
	// TODO implement me
	panic("implement me")
}

func (env *testEnvironment) Close() {
	// TODO implement me
	panic("implement me")
}

func (env *testEnvironment) runComponents() {
	var infra, services execs.Members

	for ii := 0; ii < len(env.infra); ii++ {
		im, err := env.infra[ii].Prepare()
		if err != nil {

		}

		infra = append(infra, im)
	}

	for is := 0; is < len(env.services); is++ {
		sm, err := env.services[is].Prepare()
		if err != nil {

		}

		services = append(services, sm)
	}

	p := execs.Start(execs.NewOrdered(
		execs.Member{Name: "infrastructure", Runner: execs.NewParallel(infra...)},
		execs.Member{Name: "services", Runner: execs.NewOrdered(services...)},
	))

	<-p.Ready()

	envars := env.ctx.ConfigController().DumpEnv()
	sort.Strings(envars)

	if _, err := fmt.Fprintln(os.Stdout, strings.Join(envars, "\n")); err != nil {
		env.ctx.Logger().Error("write-stdout", err)
	}

	ctx, cancel := context.WithCancel(env.ctx.Context())
	env.ctx.AddDeferFunc(
		func(ctx deps.Context) {
			cancel()

			for {
				select {
				case err := <-p.Wait():
					if err != nil {
						env.ctx.Logger().Error("stop environment with error", err)
						return
					}

					env.ctx.Logger().Info("environment process shut down")

					return
				case <-time.After(time.Second * 2):
				}
			}
		},
	)

	go func() {
		<-ctx.Done()
		p.Signal(syscall.SIGTERM)
		env.ctx.Logger().Info("got context done, sent SIGTERM to process")
	}()
}
