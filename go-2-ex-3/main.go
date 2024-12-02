package main

import "fmt"

func main() {
	// create a map called "modules"
	modules := map[int]string{
		104: "Introduction to Go",
		117: "Advanced Go",
		346: "Go Concurrency",
	}

	fmt.Println("Modul 104:", modules[104])
	fmt.Println("Modul 117:", modules[117])
	fmt.Println("Modul 346:", modules[346])

	// delete one
	delete(modules, 117)

	// add one
	modules[200] = "Go Testing"

	// replace one
	modules[346] = "Go Concurrency 2"

	fmt.Println(modules)
}
