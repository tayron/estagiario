# Estagiario
Programa que permite utilizar chatGPT 3 para gerar comandos shellscript e já executa-lo logo em seguida na sua máquina.

# CUIDADO!!!!
Qualquer coisa que pedir para ser executado, irá executar sem perdir confirmação!!!
Utilize por sua conta e risco!!!!
## Configuração
Crie uma variavel de ambiente com sua chat do chatGPT chamada ```API_KEY_CHAT_GPT```
Para gerar sua chave acesse [https://platform.openai.com/account/api-keys](https://platform.openai.com/account/api-keys)

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

### Chamando estagiário para trabalhar nos seus projetos
Em uma pasta de sua preferencia no linux, eu utilizo ```/home/tayron/Scripts/```, crie um arquivo chamado: ```chamar-estagiario.sh``` e dê permissão de execução: ```chmod +x chamar-estagiario.sh```, o conteúdo dele deverá ser o mesmo conteúdo do arquivo [https://github.com/tayron/estagiario/install.sh](install.sh) localizado na raiz deste projeto.

Agora crie um alias no seu sistema operacional seguindo o comando: 
```sh 
sudo vim ~/.bashrc 
```

No final do arquivo adicione o comando: 
```sh 
alias chamarEstagiario="/home/tayron/Scripts/./chamar-estagiario.sh" 
```
**OBS.:** Troque ```/home/tayron/Scripts``` pelo caminho de onde você criou o arquivo shell script.

Feche todos os terminais e abra seu VsCode, abra um projeto qualquer e no terminal do VsCode digite ```chamarEstagiario ```, o resultado deverá ser semelhante ao retorno abaixo:
```sh
Baixando o executável...
  % Total    % Received % Xferd  Average Speed   Time    Time     Time  Current
                                 Dload  Upload   Total   Spent    Left  Speed
  0     0    0     0    0     0      0      0 --:--:-- --:--:-- --:--:--     0
100 6404k  100 6404k    0     0  11.2M      0 --:--:-- --:--:-- --:--:-- 11.2M
Download concluído com sucesso!
```

Uma vez que foi dito que o download foi concluído, na raíz do projeto deve ter sido criado um arquivo chamado estagiário, basta executá-lo ```sh ./estagiário ``` e começar a usar.

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