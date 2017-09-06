package coffeeorbeer

import (
	"time"
)

var testCases = []struct {
	input       time.Time
	expected    string
	description string
}{
	{time.Date(2017, 8, 4, 16, 01, 0, 0, time.Local), "It's time for beer!", "After 16:00"},
	{time.Date(2017, 8, 4, 15, 59, 0, 0, time.Local), "It's time for coffee!", "Before 16:00"},
	{time.Date(2017, 8, 4, 7, 59, 0, 0, time.Local), "It's time to sober up!", "Before 08:00"},
}
