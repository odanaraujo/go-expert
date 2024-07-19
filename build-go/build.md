Nós temos a `RUNTIME` do go, que é o código do go. Ou seja, tudo que precisamos para rodar o projeto, está dentro do `RUNTIME`. Se pegarmos o `RUNTIME` + o nosso código, teremos o nosso arquivo `BINÁRIO`.


Para buildar um projeto go, basta utilizar o comando `go build` que ele vai gerar um binário.
Com esse arquivo gerado, qualquer pessoa pode executar o projeto na máquina. 
Caso seja linux, é preciso setá a variável `GOOS` da seguinte forma: `GOOS=linux go build main.go`

