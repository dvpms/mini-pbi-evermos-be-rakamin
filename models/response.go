package models

// Response represents standard JSON response across all API endpoints
type Response struct {
	Status  bool        `json:"status"`
	Message string      `json:"message"`
	Errors  interface{} `json:"errors"`
	Data    interface{} `json:"data"`
}

// SuccessResponse creates a successful response
func SuccessResponse(message string, data interface{}) Response {
	return Response{
		Status:  true,
		Message: message,
		Errors:  nil,
		Data:    data,
	}
}

// ErrorResponse creates an error response
func ErrorResponse(message string, errors ...string) Response {
	var errList []string
	if len(errors) > 0 {
		errList = errors
	}
	return Response{
		Status:  false,
		Message: message,
		Errors:  errList,
		Data:    nil,
	}
}

// PaginatedData defines standard paginated data wrapper
type PaginatedData struct {
	Page  int         `json:"page"`
	Limit int         `json:"limit"`
	Data  interface{} `json:"data"`
}

// NewPaginatedData creates a new PaginatedData object
func NewPaginatedData(page, limit int, data interface{}) PaginatedData {
	if data == nil {
		data = []interface{}{}
	}
	return PaginatedData{
		Page:  page,
		Limit: limit,
		Data:  data,
	}
}
