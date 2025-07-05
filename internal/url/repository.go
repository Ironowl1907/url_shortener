package url

import "math/rand"

var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func RandSeq(n int) string {
	res := make([]rune, n)

	for i := range n {
		res[i] = letters[rand.Intn(len(letters))]
	}

	return string(res)
}
