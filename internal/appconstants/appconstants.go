package appconstants

import (
	"time"
)

// This function is used to set the expiration date of a cookie to one hour ago.
// This is useful when we want to delete a cookie.
// Time now minus one hour is current time minus one hour.
func TimeNowMinusOneHour() time.Time {
	return time.Now().Add(-1 * time.Hour)
}

// Duration of seven days in hours. Output: 168h0m0s
// 168h0m0s represents a duration of 168 hours, 0 minutes, and 0 seconds.
const SevenDays = 7 * 24 * time.Hour

// TokenExpirationDuration defines the duration for which the token is valid.
const TokenExpiration = SevenDays

const (
	ContextKeyAuthenticatedSession = "authenticatedSession"
	ContextKeyAuthenticatedUser    = "authenticatedUser"
)

// Forms
const MinTwoCharacters = 2
const MaxFiftyCharacters = 50
const MaxTwoHundredFiftyFiveCharacters = 255
