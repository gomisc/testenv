package testenv

import (
	"context"

	"gopkg.in/gomisc/colors.v1"
	"gopkg.in/gomisc/containers.v1"
	"gopkg.in/gomisc/envs.v1"
	"gopkg.in/gomisc/network.v1/ports"
	"gopkg.in/gomisc/slog.v1"
)

// Context контекст тестовой среды.
type Context interface {
	AddDeferFunc(def DeferFunc)
	Client() containers.Client
	Colors() colors.Generator
	Context() context.Context
	Controller() envs.Controller
	Logger() slog.Logger
	Network() containers.Network
	PortsAllocator() ports.Allocator
	Verbose() int
}
