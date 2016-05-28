package main

import (
	"testing"
	"os"
	"fmt"
)

func TestDir(t *testing.T) {
	dep := &dep{Repo: "git@github.com:litriv/govendor.git"}
	got := dep.dir()
	want := "github.com/litriv"
	if got != want {
		t.Fatalf("Wanted %q but got %q", want, got)
	}
}

func TestVendor(t *testing.T) {
	f, err := os.Open("test.json")
	if err != nil {
		t.Fatal(err)
	}
	dir := os.TempDir()
	fmt.Println(dir)
	checkout(f, dir)
	
}
