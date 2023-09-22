This repository offers an array of pre-compiled resources, serving as a demonstrative showcase of AI capabilities.

## How to use

```go
package main

import "fmt"
import "github.com/proxoar/talk-demo-resource/demo"

func main() {
	pool, err := Pool()
	if err != nil {
		panice(err)
	}
	if pool.Len() == 0 {
		panice(err)
	}
	fmt.Println("resource pool len", pool.Len())
	for _, re := range pool.List() {
		fmt.Println(re.Name)
		fmt.Println(re.Text)
		fmt.Println(len(re.Audio))
	}
}

```