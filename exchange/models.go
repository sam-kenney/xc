package exchange

type Response struct {
	Result          string              `json:"result"`
	ConversionRates *map[string]float64 `json:"conversion_rates"`
	ErrorType       *string             `json:"error-type"`
}

func (r Response) IsError() bool {
	return r.Result == "error"
}
