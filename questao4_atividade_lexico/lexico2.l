%{
#include <stdio.h>
#include <string.h>
%}

%option noyywrap
%option nodefault

/* --- DEFINIÇÕES REGULARES --- */

/* 1. Símbolo: $ OU Letras seguidas opcionalmente de $ */
SIMBOLO         (\$|[a-zA-Z]+\$?)

/* 2. Dígito Inicial: 0 sozinho OU 1-9 seguido de até 2 dígitos */
DIGITO_INICIAL  (0|[1-9][0-9]{0,2})

/* 3. Milhar: Grupo de ponto seguido de 3 dígitos (pode repetir) */
MILHAR          (\.[0-9]{3})*

/* 4. Inteiro: Concatenação do início com os milhares */
INTEIRO         {DIGITO_INICIAL}{MILHAR}

/* 5. Fração: Vírgula seguida de pelo menos 2 dígitos */
FRACAO          ,[0-9]{2,}

/* 6. Número Positivo Base */
NUM_POSITIVO    {INTEIRO}{FRACAO}

/* 7. Variações de Negativo */
/* Menos na frente */
NUM_NEG_SINAL   \-{NUM_POSITIVO}
/* Entre parênteses (escapando os parênteses) */
NUM_NEG_PAR     \({NUM_POSITIVO}\)

/* 8. Número Completo (Positivo ou as variações negativas) */
NUM_COMPLETO    ({NUM_POSITIVO}|{NUM_NEG_SINAL}|{NUM_NEG_PAR})

/* 9. MOEDA FINAL (Símbolo colado no número, sem espaços) */
MOEDA_FINAL     {SIMBOLO}{NUM_COMPLETO}

%%

    /* --- REGRAS DE TRANSIÇÃO --- */

    /* Caso 1a: Moeda válida seguida de quebra de linha */
{MOEDA_FINAL}(\r|\n|\r\n) {
    /* Removemos a quebra de linha visualmente apenas para o print ficar limpo */
    yytext[strcspn(yytext, "\r\n")] = 0;
    printf("VALIDO:    %s\n", yytext);
}

    /* Caso 1b: Moeda válida no final do arquivo (sem quebra de linha depois) */
{MOEDA_FINAL} {
    printf("VALIDO:    %s\n", yytext);
}

    /* Caso Especial: Fim de Arquivo */
<<EOF>> {
    return 0;
}

    /* Caso 2: Linhas em branco (ignoramos) */
[ \t]*(\r|\n|\r\n) {
    /* Ignora linhas vazias */
}

    /* Caso 3: Qualquer outra coisa que não casou acima é Inválido (com quebra de linha) */
.*(\r|\n|\r\n) {
    yytext[strcspn(yytext, "\r\n")] = 0;
    printf("INVALIDO: %s\n", yytext);
}

    /* Caso 3b: Qualquer outra coisa inválida no final do arquivo */
.+ {
    printf("INVALIDO: %s\n", yytext);
}

%%

/* --- CÓDIGO DO USUÁRIO --- */
int main(int argc, char **argv){
    /* Se passar um arquivo como argumento, lê ele. Senão lê do teclado. */
    if(argc > 1) {
        if(!(yyin = fopen(argv[1], "r"))) {
            perror(argv[1]);
            return 1;
        }
    }
    yylex();
    return 0;
}
