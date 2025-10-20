package main

import "fmt"

func main() {

	//declaration maps
	var nilMap map[string]int

	fmt.Println(nilMap)

	teams := map[string][]string{
		"Orcas":   {"Fred", "Ralph", "Bijou"},
		"Lions":   {"Sarah", "Jones", "Billie"},
		"Kittens": {"Waldo", "Raul", "Ze"},
	}

	fmt.Println(teams)

	//Reading and writing an Maps

	totalWins := map[string]int{}
	totalWins["Orcas"] = 1
	totalWins["Lions"] = 2
	fmt.Println(totalWins["Orcas"])
	fmt.Println(totalWins["Kittens"])
	totalWins["Kittens"]++
	fmt.Println(totalWins["Kittens"])
	totalWins["Lions"] = 3
	fmt.Println(totalWins["Lions"])

	//comma ok idom

	m := map[string]int{
		"hello": 5,
		"world": 0,
	}

	v, ok := m["hello"]
	fmt.Println(v, ok)

	vW, okW := m["world"]
	fmt.Println(vW, okW)

	vU, okU := m["goodbye"]
	fmt.Println(vU, okU)

	//define a set as map

	inset := map[int]bool{}
	vals := []int{5, 10, 2, 5, 8, 7, 3, 9, 1, 2, 10}

	for _, v := range vals {
		inset[v] = true
	}

	fmt.Println(len(vals), len(inset))

	fmt.Println(inset[5])
	fmt.Println(inset[500])

	if inset[100] {
		fmt.Println("100 is in inset")
	}

}
