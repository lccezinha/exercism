package hamming

import (
	"errors"
	"strings"
)

const testVersion = 5

func Distance(a, b string) (int, error) {
	var distance int
	splitedA := strings.Split(a, "")
	splitedB := strings.Split(b, "")
	lenA := len(splitedA)
	lenB := len(splitedB)
	genomesSizes := lenA - lenB

	if genomesSizes != 0 {
		return genomesSizes, errors.New("Different genomes.")
	}

	for i := 0; i < lenA; i++ {
		if splitedA[i] != splitedB[i] {
			distance += 1
		}
	}

	return distance, nil
}
