package maps

import "fmt"

func main() {

	websites := map[string]string{
		"Google":              "https://google.com",
		"Amazon web services": "https://aws.com",
	}
	// access by key
	fmt.Println(websites["Google"])
	fmt.Println(websites)
	websites["LinkedIn"] = "https://linkedin.com"

	delete(websites, "Google")
	fmt.Println("New List===>", websites)
}
