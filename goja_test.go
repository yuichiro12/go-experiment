package tests

import (
	"io/ioutil"
	"log"
	"testing"

	"github.com/davecgh/go-spew/spew"

	"github.com/dop251/goja"
)

func TestGojaFunc(t *testing.T) {
	vm := goja.New()
	script, err := readFile("goja_test.js")
	if err != nil {
		log.Fatalln(err)
	}
	har, err := readFile("golang.org.har")
	if err != nil {
		log.Fatalln(err)
	}
	vm.Set("json", har)
	_, err = vm.RunString(script)
	if err != nil {
		log.Fatalln(err)
	}
	spew.Dump(vm.Get("jsonstr"))

}
func readFile(f string) (string, error) {
	buf, err := ioutil.ReadFile(f)
	if err != nil {
		return "", err
	}
	return string(buf), nil
}
