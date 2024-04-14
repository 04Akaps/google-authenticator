package app

import (
	"os"
	"os/signal"
	"syscall"

	"demo-scrapping/authenticator"
	"demo-scrapping/config"
	"demo-scrapping/network"
	"demo-scrapping/repository"
	"demo-scrapping/service"
)

type App struct {
	cfg *config.Config

	network *network.Network

	authenticator authenticator.AuthenticatorImpl
	repository    repository.RepositoryImpl
	service       service.ServiceImpl

	stop chan struct{}
}

func NewApp(cfg *config.Config) *App {
	a := &App{
		cfg:  cfg,
		stop: make(chan struct{}),
	}

	var err error

	if a.authenticator, err = authenticator.NewAuthenticator(cfg); err != nil {
		panic(err)
	} else if a.repository, err = repository.NewRepositry(cfg); err != nil {
		panic(err)
	}

	a.service = service.NewService(cfg, a.repository)
	a.network = network.NewNetwork(cfg, a.service, a.authenticator)

	channel := make(chan os.Signal, 1)
	signal.Notify(channel, syscall.SIGINT)

	go func() {
		<-channel
		a.exit()
	}()

	return a
}

func (a *App) Wait() {
	<-a.stop
	os.Exit(1)
}

func (a *App) exit() {
	a.stop <- struct{}{}
}

func (a *App) Run() {
	a.network.Run()
}
