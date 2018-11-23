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
	pg, err := goja.Compile(script, "", false)
	if err != nil {
		log.Fatalln(err)
	}
	v, err := vm.RunProgram(pg)
	if err != nil {
		log.Fatalln(err)
	}
	spew.Dump()
	v, err = vm.RunString(script)
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

func TestCompile(t *testing.T) {
	vm := goja.New()
	script, err := readFile("goja_test.js")
	if err != nil {
		t.Fatal(err)
	}
	har, err := readFile("golang.org.har")
	har = har
	if err != nil {
		t.Fatal(err)
	}
	pgm, err := goja.Compile("", script, false)
	vm.Set("json", har)
	vm.Set("func", "onEndScenario")
	if err != nil {
		t.Fatal(err)
	}
	v, err := vm.RunProgram(pgm)
	if err != nil {
		t.Fatal(err)
	}
	spew.Dump(v.String())
}