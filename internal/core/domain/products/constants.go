package products

import "regexp"

const (
	skuPattern             = `^FAL-\d{7,8}$`
	skuMinLength           = 11
	skuMaxLength           = 12
	productTextMinLength   = 3
	productTextMaxLength   = 50
	productPriceMin        = 1.00
	productPriceMax        = 99999999.00
	negativePriceThreshold = 0
)

var skuRegexp = regexp.MustCompile(skuPattern)
