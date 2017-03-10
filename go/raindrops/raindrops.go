package raindrops

import (
	"bytes"
	"strconv"
)

const testVersion = 2

func Convert(number int) string {
	var message bytes.Buffer
	messages := map[int]string{
		3: "Pling",
		5: "Plang",
		7: "Plong",
	}

	for key := range messages {
		if number%key == 0 {
			message.WriteString(messages[key])
		}
	}

	if message.String() == "" {
		return strconv.Itoa(number)
	}

	return message.String()
}
