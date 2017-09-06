package coffeeorbeer

import "testing"

func TestCoffeeOrBeer(t *testing.T) {
	for _, testcase := range testCases {
		observed := CoffeeOrBeer(testcase.input)
		if observed != testcase.expected {
			t.Fatalf("CoffeeOrBeer(%02d:%02d) = %s, want %s (%s)",
				testcase.input.Hour(),
				testcase.input.Minute(),
				observed,
				testcase.expected,
				testcase.description)
		}
	}
}
