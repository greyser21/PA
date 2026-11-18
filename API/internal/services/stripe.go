package services

// StripeService: stub for Stripe payment integration (Mission 1).
// Implement PaymentIntent/Checkout with STRIPE_KEY in production.
type StripeService struct{ APIKey string }

func NewStripeService(apiKey string) *StripeService { return &StripeService{APIKey: apiKey} }

func (s *StripeService) CreatePayment(amountCents int64, currency, description string) (map[string]any, error) {
	return map[string]any{"status": "pending", "provider": "stripe", "amount": amountCents, "currency": currency, "description": description}, nil
}
