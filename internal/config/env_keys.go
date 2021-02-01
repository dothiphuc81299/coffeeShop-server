package config

// Config ...
type Config struct {
	IsDev bool
	// Database
	Database struct {
		URI            string
		CoffeeShopName string
	}
}
