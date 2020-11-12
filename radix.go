package main

// ParseMixedRadix interprets mixed radix digits in base 10.
// Each position on a mixed radix numeral system may have varying numerical bases and is a multiple of the next smaller one.
func ParseMixedRadix(radix []int, digits []int) int {
	val := 0
	multiple := 1
	for i := len(radix) - 1; i >= 0; i-- {
		qty := digits[i] * multiple
		multiple *= radix[i]
		val += qty
	}
	return val
}

// FormatInt returns the mixed radix representation of i.
func FormatInt(radix []int, i int64) []int {
	digits := make([]int, len(radix))
	var multiple int64
	multiple = 1
	for r := len(radix) - 1; r > 0; r-- {
		digits[r] = int((i / multiple) % int64(radix[r]))
		multiple *= int64(radix[r])
	}
	digits[0] = int(i / multiple)
	return digits
}
