package maps

import "fmt"

type Product struct {
	id    string
	title string
	price float64
}

func main() { // map is a different kind of data structure, you can also use to group data together.
	websites := map[string]string{
		"Google":              "https://google.com",
		"Amazon Web Services": "https://aws.com",
	}
	fmt.Println(websites)
	fmt.Println(websites["Amazon Web Services"])

	websites["LinkedIn"] = "https://linkedin.com"
	fmt.Println(websites)

	delete(websites, "Google")
	fmt.Println(websites)
}
