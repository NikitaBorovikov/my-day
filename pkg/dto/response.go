package dto

type ErrorResponse struct {
	Error string `json:"error"`
}

type OKResponse struct {
	Data interface{} `json:"data"`
}

func NewErrorResponse(err error) ErrorResponse {
	return ErrorResponse{
		Error: err.Error(),
	}
}

func NewOKResponse(data interface{}) OKResponse {
	return OKResponse{
		Data: data,
	}
}
