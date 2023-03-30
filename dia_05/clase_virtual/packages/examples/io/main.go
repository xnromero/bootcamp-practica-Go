package main

import (
	"io"
	"log"
	"os"
	//"strings"
)

func main() {

	var s = "Este texto es copiado"

	/*r := strings.NewReader(s)
	_, err := io.Copy(os.Stdout, r)

	if err != nil {
		log.Fatal(err)
	}*/

	_, err := io.WriteString(os.Stdout, s)
	if err != nil {
		log.Fatal(err)
	}
}
