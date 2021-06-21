package stripe

import (
	"github.com/stripe/stripe-go"
	"github.com/turbot/steampipe-plugin-sdk/plugin"
)

func isNotFoundErrorWrap() plugin.ErrorPredicate {
	return func(err error) bool {
		if stripeErr, ok := err.(*stripe.Error); ok {
			switch stripeErr.Code {
			case stripe.ErrorCodeMissing, stripe.ErrorCodeResourceMissing:
				return true
			}
		}
		return false
	}
}

func isNotFoundError(err error) bool {
	if stripeErr, ok := err.(*stripe.Error); ok {
		switch stripeErr.Code {
		case stripe.ErrorCodeMissing, stripe.ErrorCodeResourceMissing:
			return true
		}
	}
	return false
}
