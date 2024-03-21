# Dice

```go
package main

import (
	"fmt"
	"math/rand"
)

func main() {
	var pair, vPair, howMany int
	maxRollings := 100000

	fmt.Println("Reading pair of dice value...")
	fmt.Scanf("%d", &vPair)

	for i := 0; i < maxRollings; i++ {
		pair = rand.Intn(6) + rand.Intn(6) + 2

		if pair == vPair {
			howMany++
		}
	}

	fmt.Printf("Probability of rolling a %d is %4g\n", vPair, float32(howMany)/float32(maxRollings))
}
```

Este código em Go simula o lançamento de 2 dados (representados por um par de números entre 2 e 12) um grande número de vezes (100000) e calcula a probabilidade de obter um valor específico que o usuário digitou como entrada. Essa técnica é conhecida por Monte Carlo, uma técnica estatística que utiliza métodos de amostragem aleatória para obter resultados numéricos aproximados em problemas complexos, especialmente aqueles envolvendo probabilidades.

## Explicação

Declaramos 4 variáveis int: pair, vPair, howMany e maxRollings.

- pair vai ser definido com a soma dos números de dois dados (dois números aleatórios entre 1 e 6).
- vPair vai ser definido pelo "chute" do usuário
- howMany vai contar quantas vezes o chute do usuário saiu
- maxRollings tá definindo quantas vezes os dados serão jogados

Importamos o pacote rand que está definido em `math/rand` e utilizamos ele para gerar um número aleatório.

```go
rand.Intn(6)
```

O código acima gera um número aleatório entre 0 e 5. Para simular um dado, precisamos de um número aleatório entre 1 e 6, por isso:

```go
rand.Intn(6) + 1
```

Como queremos simular dois dados, precisamos do código acima duas vezes:

```go
rand.Intn(6) + 1 + rand.Intn(6) + 1

// ou

rand.Intn(6) + rand.Intn(6) + 2
```

`howMany` é acrescida de 1 toda vez que o par de valores gerados pelos dados é igual ao chute do usuário.

Ao final, calculamos a probabilidade de aparecer o valor escolhido pelo usuário dentro de `maxRollings` partidas.