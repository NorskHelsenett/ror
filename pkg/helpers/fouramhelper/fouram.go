package fouramhelper

import "time"

// fourAm returns the next occurance of 4 AM
//
// Four AM lyrics by Scooter
// I got the classy swag
// Now that's why I get flagged
// More like a taxi cab
// Where are the chicks to check? (Down the floor)
//
// You light up another cigarette and I pour the wine
// It's four o'clock in the morning and it's starting to get light
func FourAm() time.Time {
	now := time.Now()
	var timeToPourTheWine time.Time
	if now.Hour() >= 0 && now.Hour() <= 3 {
		timeToPourTheWine = time.Date(now.Year(), now.Month(), now.Day(), 4, 0, 0, 0, time.Now().Location())
	} else {
		timeToPourTheWine = time.Date(now.Year(), now.Month(), now.Day(), 4, 0, 0, 0, time.Now().Location()).AddDate(0, 0, 1)
	}
	return timeToPourTheWine
}
