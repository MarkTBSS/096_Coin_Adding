package server

import (
	"fmt"
	"net/http"
	"sync"

	"github.com/MarkTBSS/096_Coin_Adding/config"
	"github.com/MarkTBSS/096_Coin_Adding/databases"
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
)

type echoServer struct {
	app  *echo.Echo
	conf *config.Config
	db   databases.Database
}

func (s *echoServer) healthCheck(pctx echo.Context) error {
	return pctx.String(http.StatusOK, "OK")
}

func (s *echoServer) httpListening() {
	url := fmt.Sprintf(":%d", s.conf.Server.Port)

	if err := s.app.Start(url); err != nil && err != http.ErrServerClosed {
		s.app.Logger.Fatalf("Error: %v", err)
	}
}

func (s *echoServer) Start() {
	s.app.GET("/v1/health", s.healthCheck)
	s.initItemShopRouter()
	s.initItemManagingRouter()
	s.initPlayerCoinRouter()
	s.httpListening()
}

var server *echoServer
var once sync.Once

func NewEchoServer(conf *config.Config, db databases.Database) *echoServer {
	echoApp := echo.New()
	echoApp.Logger.SetLevel(log.DEBUG)
	once.Do(func() {
		server = &echoServer{
			app:  echoApp,
			conf: conf,
			db:   db,
		}
	})
	return server
}
