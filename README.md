# Go Built-in    

As funções e tipos "built-in" podem ser usados em qualquer lugar do código sem a necessidade de importar um pacote específico. Eles são parte integrante da linguagem Go e estão sempre disponíveis para uso.

## Funções Built-in

make(): Usada para criar mapas, slice e canais.  
new(): Alocar memória para um novo valor de tipo.  
len(): Retorna o tamanho de uma slice, mapa, string, array ou canal.  
cap(): Retorna a capacidade de um slice.
append(): Adiciona elementos a um slice.
close(): Fecha um canal.  
panic(): Causa uma paralisação imediata do programa.  
recover(): Recupera de uma panificação.  
copy(): Copia elementos de um slice para outro.
delete(): Usado para excluir elementos de um mapa.
clear(): Usado para limpar elementos de um slice ou mapa.
max(): Retorna o maior valor de um tipo.
min(): Retorna o menor valor de um tipo.

## Tipos Built-in

int, int8, int16, int32, int64: Tipos inteiros com diferentes tamanhos.  
uint, uint8, uint16, uint32, uint64: Tipos inteiros não assinados com diferentes tamanhos.  
float32, float64: Tipos de ponto flutuante de precisão simples e dupla.  
string: Tipo de sequência de caracteres.  
bool: Tipo booleano.  
byte: Alias para uint8.  
rune: Tipo para representar um caractere Unicode.  
error: Interface para representar erros.  