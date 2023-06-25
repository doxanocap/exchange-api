package interfaces

type IServiceAuth interface {
	//SignOut(header http.Header) repository.ErrorResponse
	//SignIn(body []byte) (repository.AuthResponse, error)
	//SignUp(body []byte) (repository.AuthResponse, error)
	//RefreshTokens(header http.Header) repository.ErrorResponse
	//UserValidation(header http.Header) (repository.AuthUser, repository.ErrorResponse)
}

type IServiceHandler interface {
	//ExchangersData(body []byte) ([]repository.ExchangersResponse, repository.ErrorResponse)
	//CurrenciesData(body []byte) ([]repository.CurrenciesResponse, repository.ErrorResponse)
}

type IServicePool interface {
	//Start() error
	//HandleWebSocketConn()
}

type IServiceChat interface {
}

type IServiceClient interface {
	//NewClient(context *gin.Context, pool *PoolService, conn *websocket.Conn) *ClientType
}

type Service struct {
	IServiceAuth
	IServiceChat
	IServicePool
	IServiceHandler
	IServiceClient
}
