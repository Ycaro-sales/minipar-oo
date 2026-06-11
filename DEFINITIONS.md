# Linguagem Minipar

## Definição
A Linguagem Minipar é uma linguagem multiparadigma (imperativa, Orientada a objeto e funcional) de alto nivel tipada estáticamente com Coletor de lixo, focada em paralelismo real e comunicação entre computadores.

Inspirada em linguagens como Go, Rust e Python para criar uma linguagem com construções que facilitem implementações de padrões de orientação a objetos e de paralelismo (utilizando o poder da programação funcional para simplificar a implementação desses padrões)

## Funcionalidades
- Interfaces para definir comportamentos de objetos
- Estruturas nativas da linguagem como canais e os blocos par e seq para facilitar o trabalho de comunicações entre processos paralelizados
- Funções como atributos de primeira classe para permitir padrões de programação funcional
- Blocos que retornam para melhor ergonomia durante desenvolvimento

## Palavras Chaves
```
# Definição
let # Declaração de variavel
func # Definição de função
class # Definição de classe
interface # Definição de interface
enum # Definição de enumerador
struct # Definição de struct
type # Definição de tipo
Self

# Booleanos
true # Verdadeiro
false # False

# Condicionais
if
else
switch
in # checa se valor esta dentro de um iteravel

# Iteração
for (for each)
while
do # do while
continue # ir para próxima iteração do loop
pass
break # quebrar iteração do loop
goto # ir para um label

# retorno
return

```

## Tipos padrões
```
# Tipos Primitivos
any
i8
i16
i32
i64

u8
u16
u32
u64

f16
f32
f64

char # 'a'
string # "abc123"

tuple # (tipo, tipo, ...) # literal: (1, "a") imutavel
array # [tipo] # literal: [1, 2, 3] # todos do mesmo tipo
dict # <tipo_chave, tipo_valor> # literal: {a: "a"}
set # {tipo} # literal: {1, 2, 3}
chan #
enum # literal: enum.valor

function # |par[: tipo], par[: tipo], par[: tipo]| [-> tipo_saida] {
}
```

# Fluxo de controle
```
if (cond) {

}

if (cond) {

} else {

}

while (cond) {

}

do {

} while (cond)

for (i in iter) {

}

switch (var) {
	case => {}
	case => {}
	case => {}
}

```

## Definições
```
# entre chaves é opcional

#Variavel
let <id>[: tipo] = <expr>

#Função
func <id>(par: tipo, par: tipo) -> tipo {
	/* code */
}

#Classe
class <id> [implements (<interface id>, <interface id>)]{
	#atributos
	<id>: int = 123
	<id>: int = 123
	<id>: int = 123
	
	#construtor
	<id> {
	
	}
	
	#metodo
	func <id> (par){
	
	}
}

interface <id> {
	func <id>(par: tipo, par:tipo) -> tipo retorno
	func <id>(par: tipo, par:tipo) -> tipo retorno
	func <id>(par: tipo, par:tipo) -> tipo retorno
}
```

## Exemplos

### Classes e Interfaces
```
interface speaker {
	func introduce(self);
}

class Human implements speaker {
	name: str

	func introduce(self){
		print("Hello, my name is", name);
	}
}


class Duck implements speaker {
	func introduce(self){
		print("quack");
	}
}

class Dog implements speaker {
	func introduce(self){
		print("woof woof");
	}
}

func introduce_to_speaker(s: speaker) {
	print("Hello!")
	s.introduce();
}
```

### Funções como cidadões de primeira classe
```
func adder() -> ((int) -> int){
	let sum = 0;
	return func(x: int) -> int{
		sum += x;
		return sum;
	}
}

func main(){
	positive = adder();
	negative = adder();

	print(positive(2)) // 2
	print(positive(2)) // 4

	print(negative(-2)) // -2
	print(negative(-2)) // -4

}
```

Fluxos de controle
```
let array: [[i32;3];3] = [
	[1, 2, 3],
	[4, 5, 6],
	[7, 8, 9]
]
let limite: int = 3

for(i in array){
	print(i)
}

let soma: int = 0
while(soma < limite){
	soma = soma + 1
}
```

## Gramatica
[BNF](./BNF.bnf)

## Tokens
Operadores

Literais

Delimitadores

Comentários

```
```
## Erros padroes

### Lexico

### Sintatico

### Semantico

# Compilador

## Diagrama de classes
