package robotname 

import (
	"fmt"
	"math/rand"
	"time"
)

var name_lookup = map[string]int{}

type Robot struct {

	name string
}

const letters = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"

func (r *Robot) GenerateName() (string, error) {

	prefix := make([]rune, 2)
	prefix[0] = rune(letters[rand.Intn(26)])
	prefix[1] = rune(letters[rand.Intn(26)])
	suffix := fmt.Sprintf("%d", rand.Intn(1000))
	r.name = string(prefix) + suffix
	name_lookup[r.name] = 0
	return r.name, nil
}

func (r *Robot) Reset() {
	r.GenerateName() 
}

func (r *Robot) Name() (string, error) {
	
	if r.name == "" {
		r.GenerateName()
	}
	return r.name, nil
}

func init() {
	rand.Seed(time.Now().UnixNano()* int64(maxNames))
}