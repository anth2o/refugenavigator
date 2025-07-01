package scrapper

import (
	"encoding/json"
	"fmt"
	"strconv"
)

// FlexibleInt is a type that can handle both integers and string representations of integers.
type FlexibleInt int

func (fi *FlexibleInt) UnmarshalJSON(data []byte) error {
	var num int
	if err := json.Unmarshal(data, &num); err == nil {
		*fi = FlexibleInt(num)
		return nil
	}

	var str string
	if err := json.Unmarshal(data, &str); err != nil {
		return err
	}

	parsedNum, err := strconv.Atoi(str)
	if err != nil {
		return fmt.Errorf("value should be a number or a string representation of a number")
	}

	*fi = FlexibleInt(parsedNum)
	return nil
}
