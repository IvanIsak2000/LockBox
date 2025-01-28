package key

import (
	"fmt"
	"math/rand"
	"time"
)

func GenerateApiKey() string {
	rand.Seed(time.Now().Unix())
	return fmt.Sprintf("%x", rand.Int())
	
}