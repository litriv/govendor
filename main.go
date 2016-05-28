package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

func main() {

}

type dep struct {
	Repo, Commit string
}

// dir returns the dir from the Repo
// git@github.com:litriv/govendor.git -> github.com/litriv/govendor
func (dep *dep) dir() string {
	// drop the prefix ^.+@ and the postfix \/.+$
	ret := dep.Repo[strings.Index(dep.Repo, "@")+1 : strings.LastIndex(dep.Repo, "/")]
	// change : for /
	return strings.Replace(ret, ":", "/", 1)
}

// parse parses the input and returns a slice of dependencies
func parse(in io.Reader) []*dep {
	var deps []*dep
	data, err := ioutil.ReadAll(in)
	exitOn(err)
	json.Unmarshal(data, deps)
	return deps
}

// vendor takes an input source and 
// a target dir (for testing purposes, so it doesn't interfere with this utility's 
// own "vendor" package)
func vendor(in io.Reader, dir string) {
	origDir, err := os.Getwd()
	exitOn(err)
	for _, dep := range parse(in) {
		dir = filepath.Join(dir, dep.dir())
		err := os.Mkdir(dir, 777)
		exitOn(err)
		os.Chdir(dir)
		checkout(dep)
	}

	err = os.Chdir(origDir)
	exitOn(err)
}

// checkout clones the repo into the current working dir and checks out the right commit 
func checkout(dep *dep) {
	cmd := exec.Command("git", "clone", "repo")
	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()
	if err != nil {
		fmt.Fprintln(os.Stderr, out)
		exitOn(err)
	}
	
	cmd = exec.Command("git", "checkout", dep.Commit)
	err := cmd.Run()
	if err != nil {
		fmt.Fprintln(os.Stderr, out)
		exitOn(err)
	}
}

func exitOn(err error) {
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
