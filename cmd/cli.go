package main

import (
	"fmt"
	"os"

	"github.com/rutmanz/hackhour-go/pkg/api"
)

func main() {
	client := api.NewHackHourClient(os.Getenv("HACKHOUR_API_KEY"))
	out, err := client.Status()
	if err != nil {
		panic(err)
	}
	fmt.Printf("%+v\n", out)
}