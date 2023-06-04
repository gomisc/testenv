package components

import (
	"os"

	"gopkg.in/gomisc/containers.v1"
	"gopkg.in/gomisc/errors.v1"
	"gopkg.in/gomisc/execs.v1"
	"gopkg.in/gomisc/iorw.v1"
	"gopkg.in/gomisc/network.v1/ports"

	"gopkg.in/gomisc/testenv.v1"
)

var _ testenv.ComponentOption = (*kafkaImpl)(nil)

const (
	KafkaDefaultImage = "bitnami/kafka:latest"
	KafkaListenPort   = 9092
)

type kafkaImpl struct {
	*containers.BaseContainer
}

func Kafka(ctx testenv.Context, image string, background bool) testenv.ComponentOption {
	kafka := &kafkaImpl{
		BaseContainer: containers.NewBaseContainer(ctx.Client(), ctx.Network(), ctx.Controller()),
	}

	name := testenv.DefaultNamer("kafka")
	kafka.Name = name
	kafka.Ctx = ctx.Context()
	kafka.Background = background
	kafka.Autoremove = !background
	kafka.StartTimeout = testenv.DefaultStartTimeout
	kafka.Ports = containers.PortBinds{
		{
			Name:      ports.ListenPort,
			Container: containers.NewPort(ctx.PortsAllocator().NextPort(), "tcp"),
			Host:      KafkaListenPort,
		},
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
