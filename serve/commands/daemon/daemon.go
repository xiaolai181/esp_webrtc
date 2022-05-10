package daemon

import (
	"net/http"
	"os"
	"sync"

	"esp_webrtc/routers"

	"github.com/beego/beego/v2/core/logs"
	"github.com/kardianos/service"
)

type Daemon struct {
	config *service.Config
	errs   chan error
}

func NewDaemon() *Daemon {
	config := &service.Config{
		Name:        "leewiki",
		DisplayName: "leewiki service",                       //服务名称
		Description: "A document online management program.", //服务描述
	}
	return &Daemon{
		config: config,
		errs:   make(chan error, 100),
	}
}
func (d *Daemon) Config() *service.Config {
	return d.config
}

var wg sync.WaitGroup

func (d *Daemon) Start(s service.Service) error {
	go d.Run()
	return nil

}
func (d *Daemon) Run() {
	endPoint := "127.0.0.1:8000"
	routers := routers.InitRouter()
	server := &http.Server{
		Addr:    endPoint,
		Handler: routers,
	}
	server.ListenAndServe()
}
func (d *Daemon) Stop(s service.Service) error {
	if service.Interactive() {
		os.Exit(0)
	}
	return nil
}

func Install() {
	d := NewDaemon()
	d.config.Arguments = os.Args[3:]

	s, err := service.New(d, d.config)

	if err != nil {
		logs.Error("Create service error => ", err)
		os.Exit(1)
	}
	err = s.Install()
	if err != nil {
		logs.Error("Install service error:", err)
		os.Exit(1)
	} else {
		logs.Info("Service installed!")
	}

	os.Exit(0)
}

func Uninstall() {
	d := NewDaemon()
	s, err := service.New(d, d.config)

	if err != nil {
		logs.Error("Create service error => ", err)
		os.Exit(1)
	}
	err = s.Uninstall()
	if err != nil {
		logs.Error("Install service error:", err)
		os.Exit(1)
	} else {
		logs.Info("Service uninstalled!")
	}
	os.Exit(0)
}

func Restart() {
	d := NewDaemon()
	s, err := service.New(d, d.config)

	if err != nil {
		logs.Error("Create service error => ", err)
		os.Exit(1)
	}
	err = s.Restart()
	if err != nil {
		logs.Error("Install service error:", err)
		os.Exit(1)
	} else {
		logs.Info("Service Restart!")
	}
	os.Exit(0)
}
