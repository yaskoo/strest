package jsonq

import (
	"encoding/json"
	"fmt"
	"reflect"
	"strconv"
	"strings"
)

// Gets an element from json using the specified path
// Array element use their index e.g `cars.1.model`
func Get(path string, data string) string {
	var v interface{}
	json.Unmarshal([]byte(data), &v)

	parts := strings.Split(path, ".")

	cur := v
	for _, part := range parts {
		kind := reflect.TypeOf(cur).Kind()
		switch kind {
		case reflect.Map:
			temp := cur.(map[string]interface{})
			cur = temp[part]
		case reflect.Slice:
			temp := cur.([]interface{})
			i, err := strconv.Atoi(part)
			if err != nil {
				panic(err)
			}
			cur = temp[i]
		}
	}
	return fmt.Sprintf("%v", cur)
}
