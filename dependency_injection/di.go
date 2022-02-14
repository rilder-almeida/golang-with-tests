package dependencyinjection

import (
	"fmt"
	"io"
	"net/http"
)

func Greet(writer io.Writer, name string) {
	fmt.Fprintf(writer, "Hello, %s", name)
}

func MyGreeterHandler(writer http.ResponseWriter, r *http.Request) {
	Greet(writer, "world")
}

// func main() {
// 	log.Fatal(http.ListenAndServe(":5000", http.HandlerFunc(MyGreeterHandler)))
// }
