package modb

import "github.com/AbdulSayyed/go-lib-examp/moda"

// SayHelloB calls moda.SayHelloA and adds its own message
func SayHelloB() string {
    return moda.SayHelloA() + " and Hello from Module B"
}
