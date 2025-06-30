package scrapper

import (
	"encoding/json"
	"fmt"
	"strconv"
)

// FlexibleInt is a type that can handle both integers and string representations of integers.
type FlexibleInt int

// UnmarshalJSON implements the json.Unmarshaler interface.
func (fi *FlexibleInt) UnmarshalJSON(data []byte) error {
	// Try to unmarshal as an int first
	var num int
	if err := json.Unmarshal(data, &num); err == nil {
		*fi = FlexibleInt(num)
		return nil
	}

	// If that fails, try to unmarshal as a string
	var str string
	if err := json.Unmarshal(data, &str); err != nil {
		return err // If both fail, return an error
	}

	// Convert the string to an int
	parsedNum, err := strconv.Atoi(str)
	if err != nil {
		return fmt.Errorf("value should be a number or a string representation of a number")
	}

	*fi = FlexibleInt(parsedNum)
	return nil
}

// // Example struct that uses FlexibleInt
// type Example struct {
// 	Value FlexibleInt `json:"value"`
// }

// func main() {
// 	jsonStr1 := `{"value": 42}`
// 	jsonStr2 := `{"value": "42"}`

// 	var example1 Example
// 	if err := json.Unmarshal([]byte(jsonStr1), &example1); err != nil {
// 		fmt.Println("Error unmarshalling JSON:", err)
// 	} else {
// 		fmt.Printf("Example 1: %d\n", example1.Value)
// 	}

// 	var example2 Example
// 	if err := json.Unmarshal([]byte(jsonStr2), &example2); err != nil {
// 		fmt.Println("Error unmarshalling JSON:", err)
// 	} else {
// 		fmt.Printf("Example 2: %d\n", example2.Value)
// 	}
// }
