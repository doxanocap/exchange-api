package repository
                  
type ErrorResponse struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}

func (err *ErrorResponse) IsError() bool {
	if err.Message != "" || err.Status >= 300 {
		return true
	}
	return false
}
