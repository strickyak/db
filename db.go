package main

import L "github.com/ipfs/go-ipfs/Godeps/_workspace/src/github.com/syndtr/goleveldb/leveldb"
import opt "github.com/ipfs/go-ipfs/Godeps/_workspace/src/github.com/syndtr/goleveldb/leveldb/opt"

import (
  "flag"
  "fmt"
  "os"
)

var dbpath = flag.String("db", "", "path to database")

func main() {
  flag.Parse()
  a := flag.Args()

  if *dbpath == "" {
    panic("You must set flag --db=database_path")
  }

  t, err := L.OpenFile(*dbpath, &opt.Options{})
  if err != nil { panic(err) }
  defer func() {
    t.Close()
  }()

  if len(a) < 1 {
    Usage()
  }

  switch a[0] {
    case "scan":
      it := t.NewIterator(nil, &opt.ReadOptions{})
      if err != nil { panic(err) }

      it.First()
      for it.Valid() {
        fmt.Printf("%q :: %q\n", it.Key(), it.Value())
        it.Next()
      }
    case "get":
      val, err := t.Get([]byte(a[1]), &opt.ReadOptions{})
      if err != nil { panic(err) }
      fmt.Printf("%q\n", val)
    case "put":
      err := t.Put([]byte(a[1]), []byte(a[2]), &opt.WriteOptions{})
      if err != nil { panic(err) }
    case "del":
      err := t.Delete([]byte(a[1]), &opt.WriteOptions{})
      if err != nil { panic(err) }
    default:
      Usage()
  }
}

func Usage() {
  fmt.Fprintf(os.Stderr, `Usage:
    db scan
    db get Key
    db put Key value
    db del Key
`)
  os.Exit(13)
}
