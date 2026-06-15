# Minipar Studio — Guia de Uso
 
> Ambiente de desenvolvimento visual para a linguagem **Minipar**, com compilação integrada via **MCC (Minipar Compiler)**.
 
---
 
## Pré-requisitos
 
Antes de começar, certifique-se de ter instalado:
 
- [Go](https://go.dev/dl/) — para compilar o MCC
- [Node.js e npm](https://nodejs.org/) — para rodar o Minipar Studio
---
 
## 1. Compilar o MCC (Minipar Compiler)
 
O compilador MCC precisa estar compilado antes de iniciar o studio. Acesse a pasta do compilador e execute o build:
 
```bash
cd minipar-oo
go build -o mcc
```
 
Após a execução, o binário `mcc` será gerado dentro da pasta `minipar-oo`. Ele será utilizado pelo studio para compilar os arquivos `.minipar`.
 
---
 
## 2. Instalar dependências do Minipar Studio
 
Com o MCC compilado, acesse a pasta do studio e instale as dependências via npm:
 
```bash
cd minipar-studio
npm install
```
 
> **Requisito:** Node.js e npm precisam estar instalados e disponíveis no PATH do sistema.
 
---
 
## 3. Rodar o Minipar Studio
 
Você tem duas opções para executar o studio:
 
### Modo de desenvolvimento
 
Ideal para desenvolvimento e testes. Ativa hot-reload e logs detalhados:
 
```bash
npm run dev
```
 
### Modo de produção
 
Para uso em ambiente de produção, faça o build primeiro e depois inicie:
 
```bash
npm run build
npm start
```
 
---
 
## Resumo rápido
 
```bash
# 1. Compilar o MCC
cd minipar-oo
go build -o mcc
 
# 2. Instalar dependências do studio
cd minipar-studio
npm install
 
# 3a. Rodar em modo desenvolvimento
npm run dev
 
# 3b. OU rodar em modo produção
npm run build
npm start
```
 
---
 
## Estrutura esperada do projeto
 
```
/
├── minipar-oo/          # Código-fonte do compilador MCC
│   └── mcc              # Binário gerado após `go build -o mcc`
└── minipar-studio/      # Interface do Minipar Studio
    ├── package.json
    └── ...
```
 
---
 
## Problemas comuns
 
| Problema | Possível causa | Solução |
|---|---|---|
| `mcc: command not found` | Binário não compilado ou fora do PATH | Execute `go build -o mcc` dentro de `minipar-oo/` |
| `npm install` falha | Node.js ou npm não instalados | Instale via [nodejs.org](https://nodejs.org/) |
| Studio não encontra o compilador | Caminho do `mcc` incorreto | Verifique se o binário está em `minipar-oo/mcc` |
| Porta já em uso | Outro processo ocupando a porta padrão | Encerre o processo ou altere a porta nas configurações |