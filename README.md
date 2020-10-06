# loc2map
Go package to create a map image given a lat and lng

## To Use:
```
package main

import "github.com/forrest321/loc2map"

func main(){
	err := loc2map.Convert(30.330557, -86.164910, "grayton.png")
	if err != nil {
		log.Fatal(err)
	}
}
```

## Example:
See /example/main.go
```
$/prj/path/> cd example
$/prj/path/example/> go build
$/prj/path/example/> ./example
$/prj/path/example/> ls
```

## TODO:
- [x] cli tool
- [ ] hide map search box
