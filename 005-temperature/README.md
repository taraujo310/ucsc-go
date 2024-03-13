# Temperature - Loops

## Sintaxe For

```go

for <declaration expression>; <conditional expression>; <incremental expression> {
  // code
}

```

Sendo que cada uma dessas expressões são opcionais. E assim temos:

### Infinite Loop

```go
for {
  // code
}
```

### While Loop

```go
var i = 1;

for i <= 3 {
  // code

  i += 1
}
```

### For Loop

```go
for i = 1; i <= 3; i++ {
  // code
}
```

### break & continue

Go aceita as palavras reservadas `break` para interrupção do loop e `continue` para seguir para a próxima iteração do loop

### For-Each loop

Veremos um pouco mais dessa sintaxe posteriomente mas podemos ter `for-each loop`, i.e., iterando através de uma lista:

```go
strings := []string{"hello", "world"}

for i, s := range strings {
  fmt.Println(i, s)
}
```

```bash
0 hello
1 world
```

## Exemplo

```go
package main

import "fmt"

func main() {
	var from = 0.0
	var to = 250.0
	var by = 10.0
	var fahrenheit = from
	var centigrade float32

	fmt.Println("|   Fahrenheit\t| Centigrade\t|")

	for from <= to {
		centigrade = 5.0/9.0 * (fahrenheit - 32)

		fmt.Printf("|\t%g ºF\t|  %.2f ºC\t|\n", fahrenheit, centigrade)

		fahrenheit += by
		from += by
	}
}
```

### Exemplo em C

Esse é um exemplo ligeiramente adaptado do livro original do Kernighan and Ritchie __The C Programming Language__. Utiliza loop para calcular uma tabela de 25 possibilidades de temperatura em Fahrenheit e seu equivalente em Centígrados. O exemplo em C ficará abaixo para comparação.


```c
#include <stdio.h>

main() {
  float fahrenheit, centigrade;
  float from, to, by;

  from = 0.0;
  to = 250.0;
  by = 10.0;

  fahrenheit = from;

  while (from <= to) {
    centigrade = 5.0/9.0 * (fahrenheit-32);

    printf("%f\t%f\n", fahrenheit, centigrade);

    fahrenheit += by;
    from += by;
  }
}
```

## Explicação

A estrutura for está fazendo um loop desde 0 até 250 (`for from <= to`), variando o iterador a cada 10 (`from += by`).

O placeholder `%.2f` especifica a impressão de um valor float com apenas duas casas decimais.

Perceba que as variáveis `from` e `by` poderiam ter sido declaradas como inteiras. Porém, a variável `fahrenheit` é do tipo float32 e seus valores são sempre definidos através de `from` e `by`. Isso obrigaria a fazer uma quantidade de conversões explícitas que complicaria a legibilidade do código. Para tal, definimos todas como float32.

Caso seja o desejo de melhorar o uso de memória, trabalharíamos com int32 em `from` e `by` e conversões explícitas.

Outra questão relacionada aos tipos é a divisão `5.0/9.0`. Caso tivesse feito `5/9` o resultado seria em inteiro, ou seja, 0.

Perceba que no exemplo original, K&R utilizam o `while loop`. Em Go, `while` não existe pois você pode usar o `for loop` como while, que é como estamos utilizando nesse exemplo.

Em Go, as chaves são obrigatórias em todas as estruturas de controle de fluxo.

## Exercício Sugerido

1. Modifique o programa para receber valores para from, to e by.