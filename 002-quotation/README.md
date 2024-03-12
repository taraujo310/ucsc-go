# Quotation

```go
package main

import (
	"fmt"

	"rsc.io/quote"
)

func main() {
	fmt.Println(quote.Go())
}
```

## Explicação

No site [pkg.go.dev](pkg.go.dev) é possível encontrar vários módulos publicados em que os pacotes podem ser utilizados no código. No caso, estamos usando o pacote __quote__ que está incluído no pacote __rsc.io/quote__.

Para utilizar este módulo, é necessário adicionar um novo module requirements e go.sum para autenticação do módulo. Para isso, basta usar o comando:

```bash
go mod tidy
```

Este comando irá buscar e fazer download do pacote importado no código.