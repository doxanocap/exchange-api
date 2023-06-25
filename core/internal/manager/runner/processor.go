package runner

import (
	"core/internal/manager/interfaces"
	"core/internal/repository"
	"core/internal/rest"
	"core/internal/service"
	"github.com/jackc/pgx/v4/pgxpool"
)

func InitServices(repo *interfaces.Repository) *interfaces.Service {
	return &interfaces.Service{
		IServiceAuth:    service.NewAuthService(repo),
		IServiceHandler: service.NewHandlerService(repo),
	}
}

func InitRepository(db *pgxpool.Pool) *interfaces.Repository {
	return &interfaces.Repository{
		IRepoClient: repository.NewClientModel(db),
		IRepoChat:   repository.NewChatModel(db),
	}
}

func InitREST(services *interfaces.Service) *interfaces.REST {
	return &interfaces.REST{
		Router: rest.InitRouter(services),
	}
}
