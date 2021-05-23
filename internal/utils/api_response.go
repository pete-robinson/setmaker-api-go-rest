package utils

type ServiceResponse struct {
	code    int
	payload interface{}
}

// func NewApiResponse(code int, payload interface{}) (ServiceResponse, error) {
// 	a := ServiceResponse{
// 		code: code,
// 		payload: payload,
// 	}

// 	if (ServiceResponse{}) == a {
// 		return errors.New("An error occurred creating the response struct")
// 	}
// }
