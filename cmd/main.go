package main

import (
	"context"
	"penn/internal"
)

func main() {
	internal.Request(context.Background(), internal.HTTPRequester)
}
