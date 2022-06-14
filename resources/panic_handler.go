package resources

//PanicHandlerResponse is a struct that holds the detail of a Panic event
type PanicHandlerResponse struct {
	Service    string `json:"service"`
	URL        string `json:"url.omitempty"`
	Method     string `json:"method,omitempty"`
	Env        string `json:"env"`
	StatusCode int64  `json:"status_code"`
	Status     string `json:"status,omitempty"`
	Context    string `json:"context,omitempty"`
	Message    string `json:"message"`
}
