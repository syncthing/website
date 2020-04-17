package main

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"time"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	stripe "github.com/stripe/stripe-go"
	"github.com/stripe/stripe/checkout/session"
)

const (
	recurringSKU             = "plan_FZzqRxfQtAx34O"
	captchaValidationTimeout = 5 * time.Second
)

var (
	stripeSecretKey    = os.Getenv("STRIPE_SECRET_KEY")
	recaptchaSecretKey = os.Getenv("RECAPTCHA_SECRET_KEY")
)

func main() {
	stripe.Key = stripeSecretKey
	lambda.Start(handler)
}

type donationRequest struct {
	Count     int    // SKU units, which are a euro each
	Recurring bool   // monthly, or one time only?
	Captcha   string // magic token
}

func handler(ctx context.Context, request events.APIGatewayProxyRequest) (*events.APIGatewayProxyResponse, error) {
	// We expect a request body. Unmarshal it.
	var req donationRequest
	if err := json.Unmarshal([]byte(request.Body), &req); err != nil {
		return nil, fmt.Errorf("unmarshal request: %w", err)
	}

	// Trivial validation
	if req.Captcha == "" {
		return nil, errors.New("missing captcha token")
	}
	if req.Count <= 0 || req.Count > 1000 {
		return nil, fmt.Errorf("count %d out of range", req.Count)
	}

	// Captcha validation
	captchaCtx, cancel := context.WithTimeout(ctx, captchaValidationTimeout)
	defer cancel()
	if err := verifyRecaptcha(captchaCtx, req.Captcha); err != nil {
		return nil, fmt.Errorf("captcha: %w", err)
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
				{
					Plan:     stripe.String(recurringSKU),
					Quantity: stripe.Int64(int64(count)),
				},
			},
		},
	}
}

// verifyRecaptcha checks the given token against Google's verification
// service and returns nil if it checks out.
func verifyRecaptcha(ctx context.Context, token string) error {
	req, err := http.NewRequest(http.MethodPost, "https://www.google.com/recaptcha/api/siteverify", nil)
	if err != nil {
		return err
	}
	req = req.WithContext(ctx)

	q := req.URL.Query()
	q.Set("secret", recaptchaSecretKey)
	q.Set("response", token)
	req.URL.RawQuery = q.Encode()

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return fmt.Errorf("create request: %w", err)
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("HTTP %d: %s", resp.StatusCode, resp.Status)
	}
	bs, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("read response: %w", err)
	}

	var captchaResponse struct {
		Success bool
	}
	if err := json.Unmarshal(bs, &captchaResponse); err != nil {
		return fmt.Errorf("unmarshal response: %w", err)
	}
	if !captchaResponse.Success {
		return errors.New("check failed")
	}

	return nil
}
