package main

import (
	"github.com/rutmanz/hackhour-go/pkg/cli/cmd"
)

func main() {
	cmd.Execute()
	// client := api.NewHackHourClient(os.Getenv("HACKHOUR_API_KEY"))
	// out, err := client.SessionStart("test")
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Printf("%+v\n", out)
}
