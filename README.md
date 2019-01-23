# Gokv - A Light-Weight Key-Value store written in Go

Gokv is used to store, retrieve and remove key-value pairs into a JSON file.

## Getting Started

These instructions will get you a copy of the project up and running on your local machine for development and testing purposes. See deployment for notes on how to deploy the project on a live system.

### Installing

To install Gokv use:

```
go get github.com/klaatu01/gokv
```

### Example
Program 'lwkv.go'

```Go
package main

import (
  gokv "github.com/klaatu01/gokv"
  "fmt"
  "flag"
)

func main(){
  key := flag.String("k", "NIL", "Key")
  val := flag.String("v", "NIL", "Value")
  remove := flag.Bool("r", false, "Remove")
  flag.Parse()
  gokv.SetPath("./db.json")
  if(*key != "NIL" && *val == "NIL"){
    if(*remove == false){
      fmt.Println(gokv.Get(*key))
    } else {
      gokv.Remove(*key)
    }
  } else if (*key != "NIL" && *val != "NIL"){
    gokv.Set(*key, *val)
  } else {
    fmt.Println("Must have 1 or 2 arguments")
  }
}
```
Usage
`go build lwkv.go`
`./lwkv -k=foo -v=bar`
`./lwkv -k=foo`
bar
`./lwkv -r -k=foo`


## License

This project is licensed under the MIT License - see the [LICENSE.md](LICENSE.md) file for details
