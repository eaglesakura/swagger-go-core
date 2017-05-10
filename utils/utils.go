package utils

import (
	"strings"
	"reflect"
	"encoding/json"
	"fmt"
)

func AddPath(base string, local string) string {
	if strings.LastIndex(base, "/") == (len(base) - 1) {
		base = base[0:len(base) - 1]
	}

	if strings.Index(local, "/") == 0 {
		local = local[1:]
	}
	return base + "/" + local
}

func ParameterToString(ptr interface{}) string {
	ref := reflect.ValueOf(ptr)
	if ref.IsNil() {
		return ""
	}

	value := ref.Elem()
	switch value.Kind() {
	case reflect.Struct, reflect.Array:{
		buf, _ := json.Marshal(ref.Interface())
		return string(buf)
	}
	default:{
		return fmt.Sprintf("%v", value.Interface())
	}
	}
}