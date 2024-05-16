package exchange

// Important parts of the response from the exchange API.
// Response definition here if needed in future
// https://www.exchangerate-api.com/docs/standard-requests
type Response struct {
	Result          string              `json:"result"`
	ConversionRates *map[string]float64 `json:"conversion_rates"`
	ErrorType       *string             `json:"error-type"`
}

// Whether the response returned an error.
func (r Response) IsError() bool {
	return r.Result == "error"
}
