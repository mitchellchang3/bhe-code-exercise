package sieve

type Sieve interface {
	NthPrime(n int64) int64
}

type MySieve struct{}

func NewSieve() Sieve {
	return MySieve{}
}

func (s MySieve) NthPrime(req int64) int64 {
	// Initialize default values
	primeMap := make(map[int64]int64)
	currVal := int64(2)
	recentPrimeNum := int64(0)
	primeNumCount := int64(0)

	// Loop until we found the req's amount of numbers
	for primeNumCount < req+1 {
		// If the current value is not in the map, increment the prime number count and add its square into the map.
		// The reason why we add its square is because all the values in between are already multiples of previous numbers.
		// For example, the prime number 11 squared is 121. All multiples of 11 in between are already covered by multiples of 2, 3, 5 and 7 so we don't need to check for those.
		if val, ok := primeMap[currVal]; !ok {
			primeNumCount++
			recentPrimeNum = currVal

			primeMap[currVal*currVal] = currVal

			// If the current value is in the map already, don't increment the prime number count and increase the current map's value to the future value.
			// For example, if current value is 5, add 10 to the map since that is 5's future multiple.
			// We might have to loop until we find an available spot in the map.
			// For example, when we try to move 2's multiple up to 12, it's actually already covered by 3's multiple. So, we'll have to bump it up again to 14.
		} else {
			newMarker := currVal + val
			for {
				if _, ok2 := primeMap[newMarker]; !ok2 && newMarker > currVal {
					primeMap[newMarker] = val
					// We no longer need the older multiples so we remove them as we go.
					delete(primeMap, currVal)
					break
				}

				newMarker += val
			}
		}

		currVal++
	}

	// Return the last value found
	return recentPrimeNum
}
