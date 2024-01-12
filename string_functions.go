package main

import (
    "fmt"
    s "strings"
)

func main()  {
    p := fmt.Println
    p("containts", s.Contains("test", "es"))
    p("count", s.Count("test", "t"))
    p("hasprefix", s.HasPrefix("test", "te"))
    p("hassuffix", s.HasSuffix("test", "st"))
    p("index", s.Index("test", "t"))
    p("join", s.Join([]string{"foo", "bar", "baz"}, "/"))
    p("repeat", s.Repeat("/", 4))
    p("replace 1", s.Replace("foo", "o", "a", 1))
    p("replace 0 ", s.Replace("foo", "o", "a", 0))
    p("replace -1", s.Replace("foo", "o", "a", -1))
    p("split", s.Split("a-b-c", "-"))
    p("lower", s.ToLower("LKDSJFSDL"))
    p("upper", s.ToLower("lkdsjfsdl"))
}

