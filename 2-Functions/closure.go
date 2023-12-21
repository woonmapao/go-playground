package main

func concatter() func(string) string {
	doc := ""
	return func(word string) string {
		doc += word + " "
		return doc
	}
}

// to use :
func main() {
	// letsUseAggregator := concatter()
	// letsUseAggregator("Pao")
	// letsUseAggregator("and")
	// letsUseAggregator("Mon")
	// letsUseAggregator("went")
	// letsUseAggregator("up")
	// letsUseAggregator("the")
	// letsUseAggregator("hill")
	// fmt.Println(letsUseAggregator("test final result"))
	// Pao and Mon went up the hill test final result
}

// func printReport(messages []string) {
// 	for _, message := range messages {
// 		printCostReport(func(msg string) int {
// 			return len(msg) * 2
// 		}, message)
// 	}

// }
