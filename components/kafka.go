package components

import (
	"os"

	"git.eth4.dev/golibs/deps"
	"git.eth4.dev/golibs/errors"
	"git.eth4.dev/golibs/execs"
	"git.eth4.dev/golibs/iorw"
	"git.eth4.dev/golibs/network/ports"

	"git.eth4.dev/golibs/containers"

	"git.eth4.dev/golibs/testenv"
)

var _ testenv.ComponentOption = (*kafkaImpl)(nil)

const (
	KafkaDefaultImage = "bitnami/kafka:latest"
	KafkaListenPort   = 9092
)

type kafkaImpl struct {
	*containers.BaseContainer
}

func Kafka(ctx deps.ContainersAdapter, image string, background bool) testenv.ComponentOption {
	kafka := &kafkaImpl{
		BaseContainer: containers.NewBaseContainer(ctx.Client(), ctx.Network(), ctx.ConfigController()),
	}

	name := testenv.DefaultNamer("kafka")
	kafka.Name = name
	kafka.Ctx = ctx.Context()
	kafka.Background = background
	kafka.Autoremove = !background
	kafka.StartTimeout = testenv.DefaultStartTimeout
	kafka.Ports = containers.PortBinds{
		{Name: ports.ListenPort, Container: containers.NewPort(ctx.PortsAllocator().NextPort(), "tcp"), Host: KafkaListenPort},
	}

	kafka.BaseContainer.Image = image
	if kafka.BaseContainer.Image == "" {
		kafka.BaseContainer.Image = KafkaDefaultImage
	}

	color := ""
	if ctx.Verbose() >= 1 && !background {
		color = ctx.Colors().NextColor()
		kafka.ErrorStream = iorw.NewPrefixedWriter(
			testenv.FormatStderrPrefix(color, name),
			os.Stderr,
		)
	}

	if ctx.Verbose() >= 10 && !background {
		kafka.OutputStream = iorw.NewPrefixedWriter(
			testenv.FormatStdoutPrefix(color, name),
			os.Stdout,
		)
	}

	return kafka
}

// String реализация fmt.Stringer
func (k *kafkaImpl) String() string {
	return k.GetName()
}

// Prepare подготавливает компонент
func (k *kafkaImpl) Prepare() (execs.Member, error) {
	if err := containers.CheckImages(
		k.GetClient(),
		containers.WithPullImage(k.GetImage()),
	); err != nil {
		return execs.Member{}, errors.Wrap(err, "check kafka image")
	}

	return execs.Member{Name: k.GetName(), Runner: k}, nil
}

func (k *kafkaImpl) Run(signals <-chan os.Signal, ready chan<- struct{}) error {
	// TODO implement me
	panic("implement me")
}
