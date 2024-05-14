package main

import (
	"context"
	"fmt"
)

func main() {
	ctx := context.Background()
	ctx = context.WithValue(ctx, "token", "passwords")
	bookHotel(ctx, "Hotel")
}

func bookHotel(ctx context.Context, _ string) {
	token := ctx.Value("token")
	fmt.Println(token)
}