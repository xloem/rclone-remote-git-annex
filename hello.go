package main

import (
	"fmt"
	"log"
	git "github.com/go-git/go-git/v5"
	gitobject "github.com/go-git/go-git/v5/plumbing/object"
// import git "github.com/go-git/go-git"
//import (
//	jsonrpc "github.com/filecoin-project/go-jsonrpc"
//	lotusapi "github.com/filecoin-project/lotus/api"
//)
)

// git annex has multiple kinds of remotes, many are embedded
// git-annex is a single cli tool with plumbing commands.  it assumes a git repository.
// there is 

func main() {
	var gitpath string
	gitpath = "/home/ubuntu/src/test/"
	r, err := git.PlainOpen(gitpath)
	if err != nil {
		log.Fatal("PlainOpen", err)
	}

	head, err := r.Head()
	if err != nil {
		log.Fatal("Head ", err)
	}

	cIter, err := r.Log(&git.LogOptions{From: head.Hash()})
	if err != nil {
		log.Fatal("Log", err)
	}

	err = cIter.ForEach(func(c *gitobject.Commit) error {
		fmt.Println(c)
		return nil
	})
	if err != nil {
		log.Fatal("ForEach", err)
	}

	fmt.Println("Hello, World!")
}
