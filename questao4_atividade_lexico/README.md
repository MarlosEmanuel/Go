## 📄 README do Analisador Léxico

Esta pasta contém os arquivos necessários para o **Analisador Léxico**, construído usando a ferramenta **Flex** (Fast Lexical Analyzer). O analisador é projetado para processar *tokens* de entrada e classificá-los como **válidos** ou **inválidos** com base nas regras definidas.

-----

## 📂 Estrutura dos Arquivos

A pasta contém os seguintes arquivos principais:

  * **`lexico2.l`**: O arquivo de definição de regras do **Flex** (extensão `.l`). Ele contém as **definições do léxico** (padrões de expressões regulares) e as ações em C correspondentes. Este é o arquivo fonte primário do analisador.
  * **`lex.yy.c`**: O **código fonte em C** gerado pelo compilador `flex` a partir do `lexico2.l`. Ele contém a lógica do analisador léxico.
  * **`analisador`**: O **executável compilado** a partir do `lex.yy.c`. É a versão pronta para uso do analisador. (Compilado em ambiente **Linux**).

### 🧪 Arquivos de Teste

Estes arquivos de texto são exemplos de entradas para testar o analisador:

  * **`validos.txt`**: Contém exemplos de entradas (tokens/strings) que são **válidas** de acordo com as regras léxicas definidas.
  * **`invalidos.txt`**: Contém exemplos de entradas que são **inválidas** de acordo com as regras léxicas.
  * **`misturados.txt`**: Contém uma mistura de entradas **válidas** e **inválidas** para um teste mais abrangente.

-----

## 🚀 Como Executar o Analisador

O analisador pode ser executado de duas maneiras: usando o executável pré-compilado (`analisador`) ou compilando o código fonte C você mesmo.

### 1\. Usando o Executável Pré-Compilado (`analisador`)

O arquivo `analisador` é um binário **Linux**. Se estiver usando **Linux** ou **WSL (Windows Subsystem for Linux)**, você pode executá-lo diretamente.

-----

## 📝 Como Testar

Para testar o analisador, use o seu executável (`analisador` em Linux/WSL):

1.  **Teste com `validos.txt`:** O analisador deve relatar que **todos** os tokens são válidos.
    ```bash
    ./analisador  validos.txt
    ```
2.  **Teste com `invalidos.txt`:** O analisador deve relatar que **todos** os tokens são inválidos.
    ```bash
    ./analisador  invalidos.txt
    ```
3.  **Teste com `misturados.txt`:** O analisador deve identificar corretamente quais tokens são válidos e quais são inválidos/erros léxicos.
    ```bash
    ./analisador  misturados.txt
    ```