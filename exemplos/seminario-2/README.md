# Go Seminars: Fundamentos da Linguagem

Aqui contém uma série de pequenos programas Go (`.go`) utilizados para demonstrar os **tipos de dados**, **estruturas de dados**, **expressões/operadores**, **estruturas de controle de fluxo** e **escopo/vinculação** na linguagem Go (Golang).

Cada bloco de código fornecido no seu conjunto é um programa Go completo e independente (cada um contendo o `package main` e a `func main`).

## 🚀 Como Executar os Códigos

Para executar qualquer um desses programas, você precisará ter o **Go instalado** em sua máquina.

### Execução de um Bloco de Código

Abra o terminal na pasta onde os arquivos foram salvos e use o comando `go run`:

```bash
# Por exemplo:
go run 01_tipos_inteiros.go
```

-----

## 📝 Explicação dos Códigos

Abaixo está um resumo do que cada arquivo demonstra:

### 1\. Tipos de Dados Primitivos

| Arquivo Sugerido | Tópicos Principais | Descrição |
| :--- | :--- | :--- |
| **`tipos_inteiros.go`** | `int8`, `int`, `uintN`, `rune`, `byte` | Demonstra a diferença entre inteiros **com sinal** (`intN`) e **sem sinal** (`uintN`), mostrando seus limites de min/max. Explica **`rune`** (alias para `int32`, usado para Unicode/caracteres) e **`byte`** (alias para `uint8`, usado para dados). |
| **`ponto_flutuante.go`** | `float32`, `float64`, `const` | Mostra a diferença de precisão entre `float32` e `float64` (o padrão de Go). Utiliza **`const`** para valores fixos. |
| **`booleanos_strings.go`** | `bool`, `string`, `len()` | Demonstra a declaração de variáveis **`bool`** (`true`/`false`) e o tipo **`string`**, incluindo concatenação e como obter o tamanho em bytes (`len`). |

-----

### 2\. Estruturas de Dados e Referências

| Arquivo Sugerido | Tópicos Principais | Descrição |
| :--- | :--- | :--- |
| **`estruturas_de_dados.go`** | **Arrays**, **Slices**, **Structs**, **Maps**, **Ponteiros** (`*`, `&`) | Abrange **Arrays** (tamanho fixo), **Slices** (dinâmicos, com `append`), **Structs** (dados customizados), **Maps** (chave-valor, com `make` e `delete`) e o conceito de **Ponteiros** (`*` para valor, `&` para endereço de memória). |

-----

### 3\. Escopo, Precedência e Visibilidade

| Arquivo Sugerido | Tópicos Principais | Descrição |
| :--- | :--- | :--- |
| **`expressoes_e_precedencia.go`** | Operadores Aritméticos, Lógicos e Bitwise, Precedência, Ausência de Sobrecarga | Detalha a ordem de avaliação (**precedência**) dos operadores em Go (níveis 5 ao 1). Demonstra a regra de **não permitir a sobrecarga de operadores**, exigindo funções explícitas para operações complexas. |
| **`escopo_e_visibilidade.go`** | **Case Sensitive**, Variáveis Públicas/Privadas, Escopo de Bloco | Explica que Go é **Case Sensitive**. Mostra a regra de **visibilidade** (Exportação): **Letra Inicial Maiúscula** torna um identificador **Público** (exportado); **Minúscula** o torna **Privado**. Demonstra escopo de função e bloco. |

-----

### 4\. Estruturas de Controle de Fluxo

| Arquivo Sugerido | Tópicos Principais | Descrição |
| :--- | :--- | :--- |
| **`controle_de_fluxo.go`** | `if/else` com inicializador, `for` (3 tipos), `for range`, `switch` | Demonstra o `if` com declaração de variável restrita ao escopo, as três formas do **`for`** (o único loop), o **`for range`** para iteração e o **`switch`** (simples e sem expressão). |
| **`comandos_de_repeticao.go`** | `for` Clássico, `for` While, `for` Infinito, `break`, `continue`, `_` | Foca nos comandos de repetição, usando **`break`** para sair de loops e **`continue`** para pular iterações. Mostra o uso do `_` (underscore) no `for range` para ignorar valores retornados. |
| **`concorrencia_e_controle.go`** | **`go`**, **`chan`**, **`select`**, **`defer`**, **Interfaces**, **Métodos**, **`goto`** | Cobre tópicos avançados: **`go`** (Goroutines para concorrência), **`chan`** (Canais), **`select`** (Para comunicação não bloqueante), **`defer`** (Execução atrasada LIFO), **Métodos**, **Interfaces** e o comando **`goto`**. |