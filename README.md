# Estagiario
Programa que permite utilizar chatGPT 3 para gerar comandos shellscript e já executa-lo logo em seguida na sua máquina.

# CUIDADO!!!!
Qualquer coisa que pedir para ser executado, irá executar sem perdir confirmação!!!
Utilize por sua conta e risco!!!!
## Configuração
Crie uma variavel de ambiente com sua chat do chatGPT chamada ```API_KEY_CHAT_GPT```

### Exemplo de criação para linux:
Abra o arquivo ```.bashrc```:
```sh
sudo vim ~/.bashrc
```
Adicione a variàvel com sua chave no final do arquivo
```sh
export API_KEY_CHAT_GPT=sua-chave-aqui
```

### Executando a aplicação
Para executar uma aplicação em golang basta executar o comando abaixo:

```sh
go run *.go
```

## Exemplo de utilização
```sh
Estagiário: Que é hoje?

-----------------------------------------------------------------------
- Hoje é 08/04/2023
 
-----------------------------------------------------------------------


Estagiário: Verifique e me fale se a pasta teste existe

-----------------------------------------------------------------------
- A pasta teste existe
 
-----------------------------------------------------------------------


Estagiário: Remova a pasta

-----------------------------------------------------------------------
- Tarefa executada com sucesso
-----------------------------------------------------------------------


Estagiário: Crie a pasta chamada teste e dentro dessa pasta crie um projeto em nodejs utilizando gerenciador de pacotes npm cuja nome do projeto será teste. Dentro crie um arquivo chamado index.js com uma função que imprima a frase Hello World usando console.log

-----------------------------------------------------------------------
- Wrote to /home/tayron/estagiario/teste/package.json:

{
  "name": "teste",
  "version": "1.0.0",
  "description": "",
  "main": "index.js",
  "scripts": {
    "test": "echo \"Error: no test specified\" && exit 1"
  },
  "keywords": [],
  "author": "",
  "license": "ISC"
}


Hello World
 
-----------------------------------------------------------------------
```

Logo em seguida poderá ver que foi criado o diretório chamado teste com a seguinte estrutura
```sh
- teste
 - index.js
 - package.json
 ```

### Como utilizar nos seus projetos
Crie o um executável e coloque dentro da pasta de seus projetos, porém não versione, adicione ele na lista de arquivos ignorados.
O diretório raiz será onde o executável estiver sendo chamado.

#### Comando para criar executável
O comando abaixo irá gerar um executável chamado ```estagiario```
```sh
go build -o estagiario -ldflags "-s -w" main.go
```

Podendo ser utilizado da seguinte forma no linux:
```sh
./estagiario 

Estagiário: Que dia é hoje?

-----------------------------------------------------------------------
- Hoje é sábado
 
-----------------------------------------------------------------------


Estagiário:
```