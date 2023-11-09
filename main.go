package main

import (
	f "fmt"
	"os"

	b "github.com/gen2brain/beeep"
)

func main() {
	var usernameWelcome string = "welcome back, " + os.Getenv("USERNAME") + "!"
	err := b.Notify("Welcomer", usernameWelcome, "assets/1f44b.png")
	f.Println("Working!")
	if err != nil {
		panic(err)
	}
}
