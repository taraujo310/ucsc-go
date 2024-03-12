# Hello World

```go
package main

import "fmt"

func main() {
	fmt.Println("Hello, World!")
}
```

## Explicação

- Declaração de um pacote principal, uma forma de agrupar funções.
  - O pacote é feito de todos os arquivos no mesmo diretório
- Importa o pacote fmt para formatações e impressão de texto
- A função Println não aceita especificações de formatação
- Implementa a função main que será chamada quando o pacote main for executado

## Execução

```bash
go run .
```