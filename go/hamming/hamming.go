package hamming

import (
	"errors"
)

const testVersion = 5

var ErrDifferentGenomes = errors.New("Diff genomes!")

func Distance(a, b string) (distance int, _ error) {
	if len(a) != len(b) {
		return -1, ErrDifferentGenomes
	}

	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			distance++
		}
	}

	return distance, nil
}
