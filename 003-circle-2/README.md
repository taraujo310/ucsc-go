# Circle Area

```go
package main

import (
	"fmt"
	"math"
)

func main() {
	var area, radius float32

	fmt.Print("Input radius: ")
	fmt.Scanf("%f", &radius)

	area = math.Pi * radius * radius

	fmt.Printf("Area = %f\n", area)
}
```

## Explicação

Este exemplo é semelhante ao exemplo anterior porém utiliza a constante `math.Pi` do pacote math.