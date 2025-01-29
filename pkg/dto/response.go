package dto

type Response struct {
	Data interface{} `json:"data"`
}

func NewResponse(data interface{}) Response {
	return Response{
		Data: data,
	}
}
