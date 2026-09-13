package helpers

type Response struct {
	Success bool        `json:"success"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
	Error   interface{} `json:"error,omitempty"`
}

type PaginationMeta struct {
	CurrentPage int   `json:"current_page"`
	PageSize    int   `json:"page_size"`
	TotalPages  int   `json:"total_pages"`
	TotalItems  int64 `json:"total_items"`
}

type PaginationResponse struct {
	Success bool           `json:"success"`
	Message string         `json:"message"`
	Data    interface{}    `json:"data"`
	Meta    PaginationMeta `json:"meta"`
}

// Success mengembalikan struktur JSON untuk response sukses
func Success(message string, data interface{}) Response {
	return Response{
		Success: true,
		Message: message,
		Data:    data,
	}
}

// Error mengembalikan struktur JSON untuk response gagal
func Error(message string, err interface{}) Response {
	return Response{
		Success: false,
		Message: message,
		Error:   err,
	}
}

// Pagination mengembalikan struktur JSON untuk response data dengan pagination
func Pagination(message string, data interface{}, meta PaginationMeta) PaginationResponse {
	return PaginationResponse{
		Success: true,
		Message: message,
		Data:    data,
		Meta:    meta,
	}
}
