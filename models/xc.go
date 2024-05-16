package models

// CLI arguments and options.
type Xc struct {
	Quantity float64 `arg:"positional" help:"The amount of money to valuate the exchange"`
	From     string  `arg:"positional" help:"The currency to exchange from"`
	To       string  `arg:"required" help:"The currency to exchange to"`
}
