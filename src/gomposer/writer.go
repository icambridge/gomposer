package gomposer

import (
    "fmt"
    "encoding/json"
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
