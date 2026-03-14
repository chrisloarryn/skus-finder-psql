package constants

import "regexp"

const (
	ProductSKUPattern      = `^FAL-\d{7,8}$`
	ProductSKUMinLength    = 11
	ProductSKUMaxLength    = 12
	ProductTextMinLength   = 3
	ProductTextMaxLength   = 50
	ProductPriceMin        = 1.00
	ProductPriceMax        = 99999999.00
	NegativePriceThreshold = 0
)

var ProductSKURegexp = regexp.MustCompile(ProductSKUPattern)
