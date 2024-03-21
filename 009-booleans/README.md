# Booleans

```go
package main

import (
	"fmt"
	"strings"
)

func main() {
	columns := []bool{true, false}
	rows := []bool{true, false}

	fmt.Println("And Table")
	header := fmt.Sprintf("|%3s    | %3s    | p && q |", "p", "q")
	header = header + "\n" + strings.Repeat("-", len(header))
	fmt.Println(header)

	for _, p := range columns {
		for _, q := range rows {
			fmt.Printf("|%6t | %6t | %6t |\n", p, q, p && q)
		}
	}

	fmt.Println("\nOr Table")
	header = fmt.Sprintf("|%3s    | %3s    | p || q |", "p", "q")
	header = header + "\n" + strings.Repeat("-", len(header))
	fmt.Println(header)

	for _, p := range columns {
		for _, q := range rows {
			fmt.Printf("|%6t | %6t | %6t |\n", p, q, p || q)
		}
	}
}

```

## Explicação

O código acima gera a tabela verdade do and e do or. Ele inicia com a declaração de dois arrays booleanos para representarem as colunas e as linhas das tabelas.

Em seguida ele imprime o título e o header da tabela.

Veja que o header é formatado utilizando a função `Sprintf`. Essa função funciona de forma semelhante ao printf, porém retorna a string formatada, possibilitando armazenar numa variável.

Outra coisa interessante é o uso dos _widths_ para gerarmos colunas sempre com o mesmo tamanho. `width` é o número utilizado no formatador do `Printf` e do `Sprintf`, `%3t` significa que um espaçamento de 3 antes de imprimir o valor de uma variável booleana.

Na linha abaixo do header, usamos o método `string.Repeat` para repetir o caracter "-" a quantidade de vezes que teria o tamanho do header.

Já o for, preenche a tabela para cada combinação de valor da tabela verdade, ou seja, das variáveis p e q.