package manager

import (
	"core/internal/manager/interfaces"
	"core/internal/manager/runner"
	"github.com/jackc/pgx/v4/pgxpool"
	"sync"
)

type Manager struct {
	REST    *interfaces.REST
	Service *interfaces.Service
	Repo    *interfaces.Repository
}

func InitManager(pool *pgxpool.Pool) *Manager {
	cpu := &Manager{}

	var once sync.Once
	once.Do(func() {
		cpu.Repo = runner.InitRepository(pool)
		cpu.Service = runner.InitServices(cpu.Repo)
		cpu.REST = runner.InitREST(cpu.Service)

	})

	return cpu
}
