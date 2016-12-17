# govendor
Vendoring for go projects

govendor is a tool for vendoring Go projects, focuses on simplicity (KISS).

When run, it reads from a `deps` file in the root of the repo and creates a `vendor` folder into which it pulls the commit specified in `deps`.

The following three example `deps` file lines are all valid and cover all functionality: repo with hash, comment and repo at diferent path with hash.

`git@github.com:davecgh/go-spew.git commit 5215b55f46b2b919f50a1df0eaa5886afe4e3b3d`    
`# or`   
`git@github.com:go-tomb/tomb.git -> gopkg.in/tomb.v2<repo> -> commit 14b3d72120e8d10ea6e6b7f87f7175734b1faab8`     
