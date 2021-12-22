package cryptocurrency

import (
	"errors"
	"strings"
)

// In the best case scenario, this function should be migrated to a go-proto-validator
func validateCryptocurrencyMessage(cryptocurrency *CryptocurrencyMessage) error {
	if len(cryptocurrency.Symbol) < 3 {
		return errors.New("symbol must be at least 3 characters")
	}
	if len(cryptocurrency.Name) < 3 {
		return errors.New("name must be at least 3 characters")
	}
	if len(cryptocurrency.IconURL) == 0 || strings.HasPrefix("https://", cryptocurrency.IconURL) {
		return errors.New("icon cannot be empty and must have https")
	}
	return nil
}
