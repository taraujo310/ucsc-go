# Circle Area

```go
package main

import "fmt"

func main() {
	var area, radius float32
	const pi = 3.14159

	fmt.Print("Input radius: ")
	fmt.Scanf("%f", &radius)

	area = pi * radius * radius

	fmt.Printf("Area = %f\n", area)
}
```

## Explicação

- A palavra reservada `var` é utilizada para declaração de variáveis. Já a palavra reservada `const` é utilizada para declarar constantes. `float32` define o tipo da variável. Poderia ser `float64` também.
- A função `fmt.Print` imprime um texto sem pular linha.
- A função `fmt.Scanf` lê um valor pela entrada padrão de dados:
  - A formatação `"%f"` faz com que o Scanf espere por um valor do tipo float
  - No segundo argumento, está colocando o valor float lido e armazenando no mesmo endereço de memória da variável radius
- A função `fmt.Printf` imprime um texto com especificação de formatação e variáveis. No caso, o placeholder %f será substituído pela variáveil `area` e o `\n` vai inserir uma nova linha.
