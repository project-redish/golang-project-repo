package echo

import "fmt"

// WithoutTime execute print command logic
func WithoutTime(msg string){

	fmt.Println(msg)
}

// WithTimes execute print command logic
func WithTimes(msg string,echoTimes int){
	for i := 0; i < echoTimes; i++ {
        fmt.Println(msg)
      }
	}