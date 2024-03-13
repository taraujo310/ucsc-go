# Marathon

```go
package main

import "fmt"

func main () {
	var miles, yards int32 = 26, 385
	var kilometers float32

	kilometers = 1.60934 * (float32(miles) + float32(yards)/1760.0)

	fmt.Printf("\nA marathon is %g kilometers.\n\n", kilometers)
}
```

## Explicação

Este programa declara as variáveis miles e yards com valores já inicialmente definidos. Faz a conversão para kilometros e imprime o resultado da conversão.

Este programa poderia ser ter declarado as variáveis miles e yards como float32 diretamente, mas faremos assim para falar sobre conversão de tipos.

`float32(miles)` converte a variável miles para float32. A conversão é obrigatória pois não existe conversão implícita em Go.

`%g` é utilizado para impressão de valores float.