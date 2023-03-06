package publicgit

import (
	"fmt"

	"github.com/google/uuid"
)

func Example() {
	privateKeyAWS := uuid.NewString()
	publicKeyAWS := uuid.NewString()
	fmt.Println(privateKeyAWS)
	fmt.Println(publicKeyAWS)
}
