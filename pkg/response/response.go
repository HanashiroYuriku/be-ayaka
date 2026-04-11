package response

type Status string

const (
	StatusSuccess       Status = "SUCCESS"
	InvalidRequest      Status = "INVALID_REQUEST"
	DataNotFound        Status = "DATA_NOT_FOUND"
	InternalServerError Status = "SERVER_ERROR"
	DataConflict        Status = "DATA_CONFLICT"
	Unauthorized        Status = "UNAUTHORIZED"
)

func (s Status) String() string {
	return string(s)
}

// Response is the standard response format for all API responses
type Response struct {
	Status           string      `json:"status"`
	Message          string      `json:"message"`
	Data             interface{} `json:"data,omitempty"`
	TraceID          string      `json:"traceId,omitempty"`
	StatusCodeClient *int        `json:"statusCodeClient,omitempty"`
	ErrorList        interface{} `json:"errorList,omitempty"`
	ErrorField       interface{} `json:"errorField,omitempty"`
}

// NewSuccessResponse creates a new success response with the given data
func NewSuccessResponse(data interface{}) *Response {
	return &Response{
		Status:  StatusSuccess.String(),
		Message: "Request successful",
		Data:    data,
	}
}

// NewErrorResponse creates a new error response with the given message and optional error details
func NewErrorResponse(status Status, message string, id string) *Response {
	return &Response{
		Status:  status.String(),
		Message: message,
		TraceID: id,
	}
}

// NewErrorFieldResponse creates a new error response with the given message and error field details
func NewErrorFieldResponse(status Status, message string, err interface{}) *Response {
	return &Response{
		Status:     status.String(),
		Message:    message,
		ErrorField: ParseErrorField(err.(string)),
	}
}
