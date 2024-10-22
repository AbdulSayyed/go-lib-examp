package modc

import "github.com/AbdulSayyed/go-lib-examp/modb"

// SayHelloC calls modb.SayHelloB and adds its own message
func SayHelloC() string {
    return modb.SayHelloB() + " and Hello from Module C"
}


