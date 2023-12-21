package main

import (
	"errors"
)

type user struct {
	name        string
	phoneNumber int
}

func getUserMap(names []string, phoneNumbers []int) (map[string]user, error) {
	if len(names) != len(phoneNumbers) {
		return nil, errors.New("invalid sizes")
	}
	userMap := make(map[string]user)
	for i := 0; i < len(names); i++ {
		userMap[names[i]] = user{
			name:        names[i],
			phoneNumber: phoneNumbers[i],
		}
	}
	return userMap, nil
}

// func main() {

// 	ages := make(map[string]int)
// 	ages["John"] = 37

// 	ages = map[string]int{
// 		"Mon":  5,
// 		"Gink": 23,
// 	}
// 	ages["Pao"] = 22

// 	fmt.Println(ages)
// 	fmt.Println(len(ages))

// 	name6 := []string{
// 		"pao",
// 		"mon",
// 		"mee",
// 		"moo",
// 		"gink",
// 		"gonk",
// 	}

// 	phone6 := []int{
// 		123,
// 		234,
// 		345,
// 		456,
// 		567,
// 		678,
// 	}

// 	phone7 := []int{
// 		123,
// 		234,
// 		345,
// 		456,
// 		567,
// 		678,
// 		789,
// 	}

// 	fmt.Println("===============")
// 	fmt.Println("Testing case1:")
// 	userMaps1, err := getUserMap(name6, phone7)
// 	if err != nil {
// 		fmt.Println(err)
// 	} else {
// 		for _, name := range name6 {
// 			fmt.Printf("key: %v, value:\n", name)
// 			fmt.Println(" - name:", userMaps1[name].name)
// 			fmt.Println(" - number:", userMaps1[name].phoneNumber)
// 		}

// 	}

// 	fmt.Println("===============")
// 	fmt.Println("Testing case2:")
// 	userMaps2, err := getUserMap(name6, phone6)
// 	if err != nil {
// 		fmt.Println(err)
// 	} else {
// 		for _, name := range name6 {
// 			fmt.Printf("key: %v, value:\n", name)
// 			fmt.Println(" - name:", userMaps2[name].name)
// 			fmt.Println(" - number:", userMaps2[name].phoneNumber)
// 		}

// 	}
// }
