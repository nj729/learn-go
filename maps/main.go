package main

import "fmt"

func main() {
	// 1st Method
	colors := map[string]string{
		"red":   "#ff000",
		"green": "#65400",
		"white": "ffffff",
		"black": "000000",
	}
	printMap(colors)
	fmt.Println(colors)
	// // 2nd Method
	// var colors2 map[string]string
	// fmt.Println(colors2)
	// // 3rd Method
	// colors3 := make(map[string]string)
	// fmt.Println(colors3)
	// colors3["blue"] = "#651222"
	// fmt.Println(colors3)
	// delete(colors, "green")
	// fmt.Println(colors)
}

func printMap(c map[string]string) {
	for key, value := range c {
		fmt.Printf("Color: %s , Hex: %s\n", key, value)
	}
}
