package domains

type ErrorBody struct {
	Message string `json:"message"`
}

type SuccessGet struct {
	Result []Metrics `json:"result"`
}
