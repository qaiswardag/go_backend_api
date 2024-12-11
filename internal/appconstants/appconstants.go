package appconstants

import "time"

// This function is used to set the expiration date of a cookie to one hour ago.
// This is useful when we want to delete a cookie.
// Time now minus one hour is current time minus one hour.
func TimeNowMinusOneHour() time.Time {
	return time.Now().Add(-1 * time.Hour)
}

// SevenDays defines the duration of seven days.
const SevenDays = 7 * 24 * time.Hour

// TokenExpirationDuration defines the duration for which the token is valid.
const TokenExpiration = SevenDays

const (
	ContextKeyAuthenticatedSession = "authenticatedSession"
	ContextKeyAuthenticatedUser    = "authenticatedUser"
)
