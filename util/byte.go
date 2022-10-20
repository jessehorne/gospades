package util

func ReverseBytes(b []byte) []byte {
	reversed := make([]byte, len(b))
	count := len(b) - 1
	for i := 0; i < len(b); i++ {
		reversed[count] = b[i]
		count -= 1
	}

	return reversed
}
