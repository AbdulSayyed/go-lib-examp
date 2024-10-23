package modb

import "github.com/AbdulSayyed/go-lib-examp/internal/moda"

// SayHelloB calls moda.SayHelloA and adds its own message
func SayHelloB() string {
    return moda.SayHelloA() + " and Hello from Module B"
}
