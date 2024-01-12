package main

import (
	"path/filepath"
	"strings"
)


func main()  {
    p := filepath.Join("dir1", "dir2", "dir3")
    println(p)
    println(filepath.Join("dir//", "foo"))
    println(filepath.Join("dir/../dir1", "foo"))
    println(filepath.Dir(p))
    println(filepath.Base(p))
    println(filepath.IsAbs("foo/bar"))
    println(filepath.IsAbs("/foo/bar"))
    filename := "foo.json"
    ext := filepath.Ext(filename)
    println(ext)
    println(strings.TrimSuffix(filename, ext))
    rel, err := filepath.Rel("/b", "a/b/f/e")
    if err != nil {
        println(err.Error())
    }
    println(rel == "")
    rele, err := filepath.Rel("a/b", "a/b/f/e")
    if err != nil {
        println(err.Error())
    }
    println(rele)
}
