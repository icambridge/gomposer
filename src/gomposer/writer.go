package gomposer

import (
	"encoding/json"
	"fmt"
	"os"
)

func WriteLock(lock *Lock) {

	b, err := json.Marshal(lock)

	if err != nil {
		fmt.Println("error:", err)
	}

	f, err := os.Create("composer.test.lock")

	f.Write(b)
}
