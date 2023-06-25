package cmd

import (
	"core/internal/manager"
	"core/pkg/httpServer"
	"core/pkg/logger"
	"github.com/jackc/pgx/v4/pgxpool"
	"go.uber.org/zap"
	"net/http"
	"sync"
)

type App struct {
	Server  *http.Server
	Manager *manager.Manager
}

func InitApp(conn *pgxpool.Pool) *App {
	var once sync.Once

	app := &App{}
	server := httpServer.New()
	once.Do(func() {
		app.Manager = manager.InitManager(conn)
		if err := server.Run(app.Manager.REST.Router.InitEngine()); err != nil {
			logger.Log.Fatal("unable to launch RestAPI: %v", zap.Error(err))
		}
	})

	return app
}
