package tests

import (
	"encoding/json"
	"fmt"
	"log"
	"runtime"
	"sync"
	"testing"
	"time"

	"github.com/davecgh/go-spew/spew"
)

func init() {
	runtime.GOMAXPROCS(0)
}

func Test_Set_String_To_MapStringInterface(t *testing.T) {
	st := time.Now()
	m := make(map[string]interface{})
	v := "c"
	m["a"] = v
	cnt := 0
	for {
		go func() {
			defer func() {
				if r := recover(); r != nil {
					fmt.Printf("%v", r)
				}
			}()
			mm := make(map[string]interface{})
			for k, v := range m {
				mm[k] = v
			}
			cnt++
			setStringToMapStringInterface(mm)
		}()
		elapsed := time.Since(st)
		if elapsed >= 5*time.Second {
			fmt.Printf("%d goroutines called\n", cnt)
			break
		}
	}
}

func setStringToMapStringInterface(m map[string]interface{}) {
	m["a"] = "c"
}

func Test_Set_String_To_MapStringMapStringInterface(t *testing.T) {
	st := time.Now()
	m := make(map[string]interface{})
	v := make(map[string]interface{})
	v["b"] = "c"
	m["a"] = v
	cnt := 0
	mut := sync.Mutex{}
	for {
		go func() {
			defer func() {
				if r := recover(); r != nil {
					fmt.Printf("%v", r)
				}
			}()
			mm := make(map[string]interface{})
			for k, v := range m {
				mm[k] = v
			}
			cnt++
			setStringToMapStringMapStringInterface(mm, &mut)
		}()
		elapsed := time.Since(st)
		spew.Dump(elapsed)
		if elapsed >= 5*time.Second {
			fmt.Printf("%d goroutines called\n", cnt)
			break
		}
	}
}

func setStringToMapStringMapStringInterface(m map[string]interface{}, mut *sync.Mutex) {
	for {
		mut.Lock()
		m["a"].(map[string]interface{})["b"] = "c"
		mut.Unlock()
	}
}

func Test_Set_String_To_MapStringMapStringInterface_DeepCopy(t *testing.T) {
	st := time.Now()
	m := make(map[string]interface{})
	v := make(map[string]interface{})
	v["b"] = "c"
	m["a"] = v
	cnt := 0
	for {
		go func() {
			mm, err := deepCopy(m)
			if err != nil {
				t.Fatal(err)
			}
			setStringToMapStringMapStringInterface_DeepCopy(mm)
		}()
		cnt++
		elapsed := time.Since(st)
		spew.Dump(elapsed)
		if elapsed >= 5*time.Second {
			fmt.Printf("%d goroutines called\n", cnt)
			break
		}
	}
}

func deepCopy(m map[string]interface{}) (map[string]interface{}, error) {
	v := make(map[string]interface{})
	jsonStr, err := json.Marshal(m)
	if err != nil {
		return nil, err
	}
	if err := json.Unmarshal(jsonStr, &v); err != nil {
		return nil, err
	}
	return v, nil
}

func setStringToMapStringMapStringInterface_DeepCopy(m map[string]interface{}) {
	for {
		m["a"].(map[string]interface{})["b"] = "d"
	}
}

func Test_Set_String_To_MapStringMapStringInterface_Cause_Panic_AND_CANNOT_RECOVER(t *testing.T) {
	st := time.Now()
	m := make(map[string]interface{})
	v := make(map[string]interface{})
	v["b"] = "c"
	m["a"] = v
	for {
		go func() {
			defer func() {
				if r := recover(); r != nil {
					fmt.Printf("%v", r)
				}
			}()
			mm := make(map[string]interface{})
			// コピーしても内側のmap[string]interface{}が参照渡しなので、エラー
			for k, v := range m {
				mm[k] = v
			}
			// パフォーマンスを揃えるためjson.Marshal & Unmarshalする
			vv := make(map[string]interface{})
			j, err := json.Marshal(mm)
			if err != nil {
				log.Fatal(err)
			}
			json.Unmarshal(j, vv)
			setStringToMapStringMapStringInterface_Cause_Panic(mm)
		}()
		elapsed := time.Since(st)
		if elapsed >= 5*time.Second {
			break
		}
	}
}

func setStringToMapStringMapStringInterface_Cause_Panic(m map[string]interface{}) {
	for {
		m["a"].(map[string]interface{})["b"] = "c"
	}
}
