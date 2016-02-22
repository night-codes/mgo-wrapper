# mgo-wrapper (mongo)
Wrapper for package mgo with auto reconnect

## How To Use
Set environment variable DBHOST like "localhost:27017" if you want to use non-standard port or host

```go
package main

type (
    obj  map[string]interface{}
    prod struct{
        // ....
    }
)

import (
    "github.com/mirrr/mgo-wrapper"
    "fmt"
)

func main() {
    products := []prod{}
    mongo.DB("myDataBase").C("products").Find(obj{}).All(&products)
    fmt.Println(products)
}
```

## License
DO WHAT THE FUCK YOU WANT TO PUBLIC LICENSE
Version 2, December 2004

Copyright (C) 2016 Oleksiy Chechel <alex.mirrr@gmail.com>

Everyone is permitted to copy and distribute verbatim or modified
copies of this license document, and changing it is allowed as long
as the name is changed.

DO WHAT THE FUCK YOU WANT TO PUBLIC LICENSE
TERMS AND CONDITIONS FOR COPYING, DISTRIBUTION AND MODIFICATION

 0. You just DO WHAT THE FUCK YOU WANT TO.
