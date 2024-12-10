package appconstants

import "time"

// TokenExpirationDuration defines the duration for which the token is valid.
const TokenExpiration = 7 * 24 * time.Hour

const (
	ContextKeyAuthenticatedSession = "authenticatedSession"
	ContextKeyAuthenticatedUser    = "authenticatedUser"
)
