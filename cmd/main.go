package main

import example "github.com/joaosoft/example/app"

func main() {
	if _, err := example.NewExample(); err != nil {
		panic(err)
	}
}
