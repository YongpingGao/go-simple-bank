package api

import (
	"github.com/YongpingGao/go-simple-bank/util"
	"github.com/go-playground/validator/v10"
)

var validCurrency validator.Func = func(fl validator.FieldLevel) bool {
	if currency, ok := fl.Field().Interface().(string); ok {
		res := util.IsSupportedCurrency(currency)

		return res
	}
	return false
}
