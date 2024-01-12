package main

import (
	"io/fs"
	"os"
	"path/filepath"
)

func check(e error) {
    if e != nil {
        panic(e)
    }
}

func main()  {
    err := os.Mkdir("subdir", 0755)
    defer os.RemoveAll("subdir")
    check(err)
    createEmptyFile := func(name string) {
        d := []byte("")
        check(os.WriteFile(name, d, 0644))
    }
    createEmptyFile("subdir/foo")
    err = os.MkdirAll("subdir/parent/child/", 0755)
    check(err)

    createEmptyFile("subdir/parent/file2")
    createEmptyFile("subdir/parent/file3")
    createEmptyFile("subdir/parent/child/file4")

    c, err := os.ReadDir("subdir/parent")
    check(err)

    for _, entry := range c {
        println(" ", entry.Name(), entry.IsDir())
    }
     
    err = os.Chdir("subdir/parent/child")
    check(err)
    c, err = os.ReadDir(".")
    println("foo")
    for _, entry := range c {
        println(" ", entry.Name(), entry.IsDir())
    }
    err = os.Chdir("../../..")
    check(err)
    println("visiting subdir")

    err = filepath.WalkDir("subdir", visit)
}

func visit(path string, d fs.DirEntry, err error) error {
    if err != nil {
        return err
    }
    println(" ", path, d.IsDir())
    return nil
}
