package Etc

import (
	"math/rand"
	"time"
)

// RandMap ...
// return random key and val from map
func RandMap(m map[string]string) (k string, v string) {
	i := rand.Intn(len(m))
	for k := range m {
		if i == 0 {
			return k, m[k]
		}
		i--
	}
	panic("never")
}

/* from go 1.10~
func RandArray() {
	words := strings.Fields("ink runs from the corners of my mouth")
	rand.Shuffle(len(words), func(i, j int) {
		words[i], words[j] = words[j], words[i]
	})
	fmt.Println(words)
}
*/

// RandArray ...
// pick random value from array
func RandArray(array []int) int {
	s := rand.NewSource(time.Now().Unix())
	r := rand.New(s) // initialize local pseudorandom generator
	randint := r.Intn(len(array))

	return randint
}
