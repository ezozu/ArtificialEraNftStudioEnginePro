// cmd/artificialeranftstudioenginepro/main.go
package main

import (
"flag"
"log"
"os"

"artificialeranftstudioenginepro/internal/artificialeranftstudioenginepro"
)

func main() {
verbose := flag.Bool("verbose", false, "Enable verbose logging")
flag.Parse()

app := artificialeranftstudioenginepro.NewApp(*verbose)
if err := app.Run(); err != nil {
log.Fatal(err)
}
}
