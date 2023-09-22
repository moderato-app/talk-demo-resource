This repository offers an array of pre-compiled resources, serving as a demonstrative showcase of AI capabilities.

## How to use

```go
package main

import (
	"fmt"
	
	demo "github.com/proxoar/talk-demo-resource"
)

func main() {
	pool, err := demo.NewResourcePool()
	if err != nil {
		panic(err)
	}
	if pool.Len() == 0 {
		panic(err)
	}
	fmt.Println("resource pool len", pool.Len())
	for _, re := range pool.List() {
		fmt.Println(re.Name)
		fmt.Println(re.Text)
		fmt.Println(len(re.Audio))
	}
}

```