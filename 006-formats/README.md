# Formats

O pacote `fmt` provê as funções `Print`, `Println` e `Printf`.

- `Print` imprime um texto
- `Println` imprime um texto + nova linha
- `Printf` imprime um texto formatado

## A sintaxe de `Println`:

```go
var x = 55
fmt.Println("x value is ", x)
```

Output:

```bash
x value is 55
| // cursor
```

A formatação, no `Println` acontece de forma sequencial dos seus argumentos e a cada variável o tipo é inferido e formatado da forma correta.

## A sintaxe do `Printf`:

```go
var x = 55
var y = 33
fmt.Printf("x value is %d and y value is %d\n", x, y)
```

Output:

```bash
x value is 55 and y value is 33
| // cursor
```

O argumento `"x value is %d\n"` é uma string de controle de formatação. O placeholder `%d` indica que uma variável do tipo inteiro listada nos argumentos seguintes será impressa.

A ordem dos placeholders importa! No caso, o primeiro placeholder formatará a primeira variável listada (`x`) e o segundo placeholder formatará a segunda variável listada (`y`).

Para que o cursor vá para a linha de baixo, é necessário imprimir uma nova linha. Para isso, foi utilizado `\n`.

## Lista de Placeholders

| Placeholder | Uso                       | Exemplo               |
| ------------| ------------------------- | --------------------- |
| %d          | int base 10               | 55                    |
| %o          | int base 8                | 67 (55 base 8)        |
| %x          | int base 16               | 43 (55 base 16)       |
| %b          | int base 2                | 110111 (55 base 2)    |
| %f          | float                     | 12.345                |
| %g / %e     | float notação científica  | 1.2345 e+01           |
| %c          | charater                  | A                     |
| %s          | string                    | A string with blanks  |