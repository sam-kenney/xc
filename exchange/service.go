package exchange

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"strings"

	"sam.kenney.blog/xc/models"
)

// Display the converted rate for the given quantity and currencies.
func Get(xc models.Xc, apiKey string) error {
	if apiKey == "" {
		return errors.New("Exchange API key not built into bin")
	}

	from := strings.Trim(strings.ToUpper(xc.From), " ")
	to := strings.Trim(strings.ToUpper(xc.To), " ")
	qty := xc.Quantity

	resp, err := request(from, apiKey)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	res, err := getRates(*resp)
	if err != nil {
		return nil
	}

	rates := *res
	rate, ok := rates[to]
	if !ok {
		return errors.New(fmt.Sprintf("Currency %s not found", to))
	}

	calc := qty * rate
	fmt.Printf("%.2f %s = %.2f %s\n", qty, from, calc, to)

	return nil
}

// Make the request to the API.
func request(from string, apiKey string) (*http.Response, error) {
	url := fmt.Sprintf("https://v6.exchangerate-api.com/v6/%s/latest/%s", apiKey, from)

	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

// Return the map of currencies to exchange rates from a response body.
func getRates(resp http.Response) (*map[string]float64, error) {
	data := Response{}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	json.Unmarshal(body, &data)

	if data.IsError() {
		return nil, errors.New(fmt.Sprintf("Error: %s", *data.ErrorType))
	}

	return data.ConversionRates, nil
}
