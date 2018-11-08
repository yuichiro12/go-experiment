package tests

import (
	"reflect"
	"testing"

	"github.com/davecgh/go-spew/spew"
)

func Test_Cannot_Set_Dynamically_Nested_Map_to_Interface(t *testing.T) {
	keys := []string{"a", "b", "c"}
	var i interface{}
	v := "something"
	actual := make(map[string]interface{})
	actual["a"] = make(map[string]interface{})
	actual["a"].(map[string]interface{})["b"] = make(map[string]interface{})
	actual["a"].(map[string]interface{})["b"].(map[string]interface{})["c"] = make(map[string]interface{})
	actual["a"].(map[string]interface{})["b"].(map[string]interface{})["c"] = "something"
	setValue(keys, i, v)
	t.Fatalf(`
[expected]:
%v
[actual]:
%v`, spew.Sdump(actual), spew.Sdump(i))
}

func Test_can_Set_Value_to_Interface(t *testing.T) {
	var i interface{}
	value := "hello"
	set(&i, value)
	spew.Sdump(i)
}

func set(i interface{}, v interface{}) {
	i = v
}

func setValue(keys []string, p interface{}, v interface{}) {
	if len(keys) > 1 {
		if m, ok := p.(map[string]interface{}); ok {
			if _, ok := m[keys[0]]; ok {
				// すでにそのキーを持ってる場合
				setValue(keys[1:], m[keys[0]], v)
			} else {
				m[keys[0]] = createNestedValue(keys[1:], v)
			}
		} else {
			val := createNestedValue(keys, v)
			p = reflect.ValueOf(val).Interface()

		}
	} else {
		return
	}
}

func createNestedValue(keys []string, v interface{}) interface{} {
	m := make(map[string]interface{})
	if len(keys) > 1 {
		m[keys[0]] = createNestedValue(keys[1:], m)
	} else {
		m[keys[0]] = v
	}
	return m
}
