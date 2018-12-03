//** Multiline strings **//

// package main
//
// import "fmt"
//
// func main() {
// 	str := `This is
// 	a
// 	string.`
// 	fmt.Println("vim-go")
// 	fmt.Println(str)
// }

//** concatenation strings with bytes.buffer **//
// package main
//
// import (
// 	"bytes"
// 	"fmt"
// )
//
// func main() {
// 	var b bytes.Buffer
//
// 	for i := 0; i < 1000; i++ {
// 		b.WriteString(randString())
// 	}
//
// 	fmt.Println(b.String())
// }
//
// func randString() string {
// 	return "abc-123-"
// }

//** concatenation strings with strings.join **//

// package main
//
// import (
// 	"fmt"
// 	"strings"
// )
//
// func main() {
// 	var strs []string
//
// 	for i := 0; i < 1000; i++ {
// 		strs = append(strs, randString())
// 	}
//
// 	fmt.Println(strings.Join(strs, ""))
// }
//
// func randString() string {
// 	return "abc-123-"
// }

//** convert strings with strconv **//
// package main
//
// import (
// 	"fmt"
// 	"strconv"
// )
//
// func main() {
// 	i := 123
// 	t := strconv.Itoa(i)
// 	fmt.Println(i)
// 	fmt.Println(t)
// }

//** convert strings with fmt.sprintf **//
// package main
//
// import (
// 	"fmt"
// )
//
// func main() {
// 	i := 123
// 	t := fmt.Sprintf("We are currently processing ticket number %d.", i)
// 	u := fmt.Sprintln("We are currently processing ticket number %d.", i)
// 	fmt.Println(i)
// 	fmt.Println(t)
// 	fmt.Println(u)
// }

//** Creating randowm strings **//
// package main
//
// import (
// 	"fmt"
// 	"math/rand"
// 	"time"
// )
//
// func main() {
// 	fmt.Println(RandString(10))
// }
//
// var source = rand.NewSource(time.Now().UnixNano())
//
// const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
//
// func RandString(length int) string {
// 	b := make([]byte, length)
// 	for i := range b {
// 		b[i] = charset[source.Int63()%int64(len(charset))]
// 	}
// 	return string(b)
// }

//** strings.HasPrefix strings.Suffix**//
// package main
//
// import (
// 	"fmt"
// 	"strings"
// )
//
// func main() {
// 	fmt.Println(strings.HasPrefix("something", "some"))
// 	fmt.Println(strings.HasSuffix("something", "thing"))
// }

//** Strings converts to byte slice **//
package main

import (
	"fmt"
)

func main() {
	var s string = "this is a string"
	fmt.Println(s)
	var b []byte
	b = []byte(s)
	fmt.Println(b)
	for i := range b {
		fmt.Println(string(b[i]))
	}
	s = string(b)
	fmt.Println(s)
}
