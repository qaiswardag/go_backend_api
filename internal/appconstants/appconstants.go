package appconstants

import "time"

const SevenDays = 7 * 24 * time.Hour

// TokenExpirationDuration defines the duration for which the token is valid.
const TokenExpiration = SevenDays

const (
	ContextKeyAuthenticatedSession = "authenticatedSession"
	ContextKeyAuthenticatedUser    = "authenticatedUser"
)
