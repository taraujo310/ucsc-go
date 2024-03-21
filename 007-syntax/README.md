# Syntax

Go é uma linguagem compilada de baixo nível. Para rodar um programa em Go, é necessário que ele esteja sintaticamente correto. Um programa sintaticamente correto não significa não significa que está correto mas, sim, que está executável.

## Package

Um programa em Go começa a definição de um pacote utilizando a keyword `package`. Um pacote permite o agrupamento de funcionalidades relacionadas em arquivos separados dentro de um mesmo diretório, proporcionando uma maior modularidade ao código. _Identificadores_ definidos dentro de um pacote podem ser exportados (i.e., visíveis fora desse pacote) caso tenham seu primeiro caractér com a letra maiúscula.

Um pacote com o nome `main` é especial: ele diz ao compilador que a execução de um código iniciará nesse pacote. O pacote `main` deve conter a função main que será o ponto de partida do programa.

```go
package fmt
package main
```

## Import

Um programa usualmente fará uso de pacotes externos que definem funcionalidades criadas por outras pessoas. Para isso, utilizamos a keyword `import`.

```go
import "fmt"

import (
  "fmt"
  "time"
  "os"

  "caminho/do/meu/pacote" // Importando um pacote do próprio projeto
)
```

É possível utilizar alias para simplificar o nome de um pacote ou para evitar conflito entre identificadores.

```go
package main

import f "fmt"

func main() {
  f.Println("Hello World")
}
```

## Declaração de Variável

A declaração de variável em Go utiliza a seguinte sintaxe:

```go
var nomeDaVariavel tipoDaVariavel
```

É padrão em Go utilizar o estilo camelCase para escrita de identificadores. Por exemplo:

```go
var age int
var companyName string
```

É possível definir a variável no momento da declaração:

```go
var age int = 18
var active bool = true
```

Também é possível declarar uma lista de variáveis em uma única declaração:

```go
var x, y, z int
```

O nome da variável deve seguir a regra dos identificadores definida na gramática da linguagem. No caso, deve iniciar com letra ou sublinhado, seguido por letras, dígitos ou sublinhados.

Um identificador iniciado por `_` costuma ser lido como privado.

Mas nem sempre é necessário declarar o tipo da variável. Ele pode ser inferido pelo tipo do valor que está definindo a variável:

```go
name := "Thiago" // name é do tipo string pois "Thiago" é do tipo string
age := 31 // age é do tipo int pois 31 é do tipo int
```

Também é possível declarar uma variável constante utilizando a keyword `const`:

```go
const pi float64 = 3.14159
```

Variáveis declaradas e não definidas são automaticamente definidas pelo compilador.

- Uma variável numérica é automaticamente definida com o número 0.
- Uma variável booleana é automaticamente definida como false.

## Tipos de Variáveis

### 1. Tipos Primitivos:
- **int**: Tipo inteiro para números inteiros.
- **int8, int16, int32, int64**: Tipos inteiros com tamanhos específicos em bits (8, 16, 32 e 64 bits, respectivamente).
- **uint**: Tipo inteiro sem sinal para números inteiros positivos.
- **uint8, uint16, uint32, uint64**: Tipos inteiros sem sinal com tamanhos específicos em bits.
- **float32, float64**: Tipos de ponto flutuante para números com ponto decimal.
- **bool**: Tipo booleano para valores verdadeiro (true) ou falso (false).
- **string**: Tipo para armazenar cadeias de caracteres (sequências de caracteres).
- **complex64**, *complex128**: Tipo para números complexos com partes real e imaginária de 32 bits e 64 bits respectivamente.
- **uintptr**: Tipo inteiro sem sinal grande o suficiente para armazenar o valor não interpretado de um ponteiro.

### 2. Tipos Compostos:
- **array**: Coleção fixa de elementos do mesmo tipo. Exemplo:

```go
var numeros [5]int
```
- **slice**: Coleção dinâmica de elementos do mesmo tipo. Exemplo:

```go
var numeros []int
```
- **map**: Estrutura de dados chave-valor. Exemplo:

```go
var alunos map[string]int
```
- **struct**: Tipo de dados que agrupa campos com diferentes tipos. Exemplo:

```go
type Pessoa struct {
  Nome string;
  Idade int;
}
```

### 3. Outros Tipos:
- **pointer**: Tipo que armazena o endereço de memória de uma variável. Exemplo:

```go
var ptr *int
```

- **interface**: Tipo que define um conjunto de métodos que uma variável pode implementar. Exemplo:

```go
type Animal interface {
  EmitirSom()
}
```
- **func**: Tipo que representa uma função. Exemplo:

```go
var f func(int) int.
```

### 4. Tipo Vazio
- **nil**: Valor nulo em Go usado para representar a ausência de valor em ponteiros, interfaces, slices, mapas, canais e funções.

### 4. Tipos definidos pelo desenvolvedor

São tipos definidos pelo desenvolvedor aqueles com a keyword `type` seguido pelo nome do tipo e sua definição

```go
type Kilogram float64
```

## Variáveis Globais

Uma variável flobal é acessível por todo o pacote em que está definida. Sua declaração fica fora de qualquer função.

```go
package main

import "fmt"

var variavelGlobal int = 100

func main() {
  fmt.Println("Variável global: ", variavelGlobal)
}
```

Porém, o usuo de variáveis globais não é recomendado:

Variáveis globais têm escopo indefinido, o que significa que podem ser acessadas e modificadas por qualquer parte do código dentro do pacote. Isso pode levar a efeitos colaterais indesejados e tornar o código mais difícil de entender e depurar.

Em programas concorrentes, o uso de variáveis globais pode resultar em condições de corrida (race conditions) se não forem tratadas corretamente com mecanismos de sincronização, como mutexes. Isso pode levar a bugs difíceis de reproduzir e corrigir.

O uso excessivo de variáveis globais pode tornar o código menos legível e mais difícil de entender para outros desenvolvedores que estão trabalhando no mesmo projeto. Isso pode aumentar a curva de aprendizado e dificultar a manutenção do código no futuro.

## Escopo de Uma Variável

O escopo definidade a visibilidade e a acessibilidade de um identificador.

O **escopo global** é o nível mais amplo de visibilidade, onde uma variável global fica definida.

Além do escopo global, um escopo mais restrito é definido por um bloco de código limitado por `{}`.

Identificadores declarados entre `{}` será acessível apenas dentro desse bloco. Os blocos possuem visibilidade hierárquica, o que significa que um identificador definido num bloco mais externo é visível dentro de um bloco mais interno, porém o contrário não é verdade.

```go
func main() {
  var maisExterna int = 5

  {
    var maisInterna int = 10

    fmt.Println("maisExterna: ", 5)
  }

  // esse código não compila pois essa variável não existe nesse escopo
  fmt.Println("maisInterna: ", maisInterna)
}
```

Um identificador pode ser "redeclarado" em escopo mais interno:

```go
func main() {
  var numero int = 5

  {
    var numero int = 10

    fmt.Println("número: ", numero)
  }

  fmt.Println("número: ", numero)
}
```

Que resultará em:

```bash
número 10
número 5
```

## Palavras Reservadas

**break, default, func, interface, select, case, defer, go, map, struct, chan, else, goto, package, switch, const, fallthrough, if, range, type, continue, for, import, return, var**