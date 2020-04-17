package main

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	stripe "github.com/stripe/stripe-go"
	"github.com/stripe/stripe/checkout/session"
)

const (
	recurringSKU = "plan_FZzqRxfQtAx34O"
)

func main() {
	stripe.Key = os.Getenv("STRIPE_SECRET_KEY")
	lambda.Start(handler)
}

type donationRequest struct {
	Count     int  // SKU units, which are a euro each
	Recurring bool // monthly, or one time only?
}

func handler(ctx context.Context, request events.APIGatewayProxyRequest) (*events.APIGatewayProxyResponse, error) {
	// We expect a request body. Unmarshal it.
	var req donationRequest
	if err := json.Unmarshal([]byte(request.Body), &req); err != nil {
		return nil, fmt.Errorf("unmarshal request", req.Count)
	}

	// Trivial validation
	if req.Count <= 0 || req.Count > 1000 {
		return nil, fmt.Errorf("count %d out of range", req.Count)
	}

	// Get the Stripe session
	session, err := session.New(paramsFor(req))
	if err != nil {
		return nil, fmt.Errorf("creating Stripe session: %w", err)
	}

	// Respond with the session ID. The frontend can now redirect to the
	// corresponding Checkout page.
	bs, _ := json.Marshal(map[string]string{"sessionID": session.ID})
	return &events.APIGatewayProxyResponse{
		StatusCode: http.StatusOK,
		Headers:    map[string]string{"Content-Type": "application/javascript;charset=utf-8"},
		Body:       string(bs),
	}, nil
}

// paramsFor returns Checkout session parameters matching the given request
// (amount and recurrency).
func paramsFor(req donationRequest) *stripe.CheckoutSessionParams {
	var params *stripe.CheckoutSessionParams

	// Params for the request
	if req.Recurring {
		params = recurringParams(req.Count)
	} else {
		params = oneTimeParams(req.Count)
	}

	// Common settings
	params.PaymentMethodTypes = stripe.StringSlice([]string{"card"})
	params.SuccessURL = stripe.String("https://syncthing.net/donations/success/")
	params.CancelURL = stripe.String("https://syncthing.net/donations/cancelled/")

	return params
}

// oneTimeParams returns Checkout session parameters for a one time donation
func oneTimeParams(count int) *stripe.CheckoutSessionParams {
	return &stripe.CheckoutSessionParams{
		LineItems: []*stripe.CheckoutSessionLineItemParams{
			{
				Name:        stripe.String("Donation"),
				Description: stripe.String("One time donation to the Syncthing Foundation"),
				Amount:      stripe.Int64(int64(count) * 100), // cents
				Currency:    stripe.String(string(stripe.CurrencyEUR)),
				Quantity:    stripe.Int64(1),
			},
		},
	}
}

// oneTimeParams returns Checkout session parameters for a recurring
// (monthly) donation
func recurringParams(count int) *stripe.CheckoutSessionParams {
	return &stripe.CheckoutSessionParams{
		SubscriptionData: &stripe.CheckoutSessionSubscriptionDataParams{
			Items: []*stripe.CheckoutSessionSubscriptionDataItemsParams{
				&stripe.CheckoutSessionSubscriptionDataItemsParams{
					Plan:     stripe.String(recurringSKU),
					Quantity: stripe.Int64(int64(count)),
				},
			},
		},
	}
}
