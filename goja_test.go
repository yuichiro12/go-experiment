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
	spew.Dump(har)
	if err != nil {
		log.Fatalln(err)
	}
	vm.Set("json", har)
	pg, err := goja.Compile(har, "", false)
	if err != nil {
		log.Fatalln(err)
	}
	vm.RunProgram(pg)
	v, err := vm.RunString(script)
	if err != nil {
		log.Fatalln(err)
	}
	spew.Dump(v.String())
	vm = goja.New()
	vm.Set("json", har)
	v, err = vm.RunString(script)
	if err != nil {
		log.Fatalln(err)
	}
	spew.Dump(v.String())
}
func readFile(f string) (string, error) {
	buf, err := ioutil.ReadFile(f)
	if err != nil {
		return "", err
	}
	return string(buf), nil
}
