package stripe

import (
	"github.com/stripe/stripe-go/v76"
)

func isNotFoundError(err error) bool {
	if stripeErr, ok := err.(*stripe.Error); ok {
		switch stripeErr.Code {
		case stripe.ErrorCodeMissing, stripe.ErrorCodeResourceMissing:
			return true
		}
	}
	return false
}
