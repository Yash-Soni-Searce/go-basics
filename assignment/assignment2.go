package main

// import (
// 	"fmt"
// 	"os"
// )

// func main() {
// 	fmt.Println(os.Args[1])
// 	a := os.Args[1]
// 	x,err:= os.ReadFile(a)
// 	if err!= nil {
// 		fmt.Println(err)
// 	}
// 	fmt.Println(string(x))
// }
import (
    "fmt"
    "io"
    "os"
)

func main() {
    inputFile := os.Args[1]
    file, err := os.Open(inputFile) // For read access.
    if err != nil {
        fmt.Print("encountered error while trying to open the file")
    }
    io.Copy(os.Stdout, file)
}