package landingai

// HTTP status codes used in the SDK
const (
	StatusOK                  = 200
	StatusPartialContent      = 206
	StatusBadRequest          = 400
	StatusUnauthorized        = 401
	StatusPaymentRequired     = 402
	StatusUnprocessableEntity = 422
	StatusTooManyRequests     = 429
	StatusInternalServerError = 500
	StatusGatewayTimeout      = 504
)
