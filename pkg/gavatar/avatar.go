package gavatar

import (
	"fmt"
	"math/rand"
)

func GenRandomAvatar(endpoint string) string {
	index := rand.Intn(70) + 1
	return fmt.Sprintf("%s/images/avatar/default/avatar_%d.png", endpoint, index)
}
