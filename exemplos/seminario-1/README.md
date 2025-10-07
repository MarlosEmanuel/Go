## Explicação dos Códigos

### Paradigma Funcional — `Funcional/main.go`

```go
package main

import "fmt"

func contador() func() int {
    x := 0
    return func() int {
        x++
        return x
    }
}

func main() {
    c := contador()
    fmt.Println(c())
    fmt.Println(c())
}
```

**Conceito:**
Este código mostra o uso de **funções de ordem superior** e **closures** — recursos típicos do paradigma funcional.
A função `contador` retorna outra função que **mantém o estado interno** da variável `x`, mesmo após a função externa ter sido executada.

**Saída esperada:**

```
1
2
```

Cada chamada de `c()` incrementa o contador, mostrando o encapsulamento funcional de estado.

---

### Paradigma Imperativo — `Imperativo/imperativo.go`

```go
package main

import "fmt"

func main() {
    for i := 0; i < 5; i++ {
        fmt.Println(i)
    }
}
```

**Conceito:**
O paradigma imperativo se baseia em **sequências de instruções** que alteram o estado do programa passo a passo.
O exemplo usa um laço `for` para imprimir números de `0` a `4`, mostrando o controle explícito do fluxo de execução.

**Saída esperada:**

```
0
1
2
3
4
```

---

### Paradigma Orientado a Objetos — `orientado-a-objetos/cachorro.go`

```go
package main

import "fmt"

type Animal struct {
    Nome string
}

func (a Animal) Falar() {
    fmt.Println(a.Nome, "Faz um som.")
}

type Cachorro struct {
    Animal
}

func (c Cachorro) Latir() {
    fmt.Println(c.Nome, "Late!")
}

func main() {
    c := Cachorro{Animal{"Rex"}}
    c.Falar()
    c.Latir()
}
```

**Conceito:**
Go não possui herança tradicional como em Java, mas utiliza **composição de tipos** e **métodos associados**.
O exemplo mostra como criar tipos que **simulam classes** e **comportamentos especializados**, permitindo reuso de código e organização por entidades.

**Saída esperada:**

```
Rex Faz um som.
Rex Late!
```

---

## Conclusão

Os três exemplos mostram como a linguagem Go permite aplicar diferentes paradigmas de programação:

| Paradigma               | Características principais           | Exemplo no código     |
| ----------------------- | ------------------------------------ | --------------------- |
| **Imperativo**          | Controle de fluxo e estado explícito | Laço `for`            |
| **Funcional**           | Funções como valores e closures      | `contador()`          |
| **Orientado a Objetos** | Tipos compostos e métodos associados | `Cachorro` e `Animal` |

Esses exemplos foram utilizados no **Seminário 1** como demonstração prática dos principais paradigmas de programação aplicados ao Go.

---
