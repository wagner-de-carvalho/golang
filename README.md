# Golang

O **Go** começou em setembro de 2007, quando **Robert Griesemer**, **Ken Thompson** e eu começamos a discutir uma nova linguagem para lidar com os desafios de engenharia que nós e nossos colegas do **Google** enfrentávamos em nosso trabalho diário.

Quando lançamos o **Go** ao público pela primeira vez em novembro de 2009, não sabíamos se seria amplamente adotado ou se poderia influenciar futuros idiomas. Olhando para trás à partir de 2020, o **Go** teve sucesso nos dois sentidos: é amplamente usado dentro e fora do **Google**, e suas abordagens à simultaneidade de rede e engenharia de software tiveram um efeito notável em outras linguagens e suas ferramentas.

**Go** acabou por ter um alcance muito mais amplo do que esperávamos. Seu crescimento no setor foi fenomenal e impulsionou muitos projetos no **Google**.
- ***Rob Pike***

## Comandos
- `go mod init nome-do-modulo` = cria um módulo.
- `go run {arquivo}.go` = executa o arquivo.
- `./nome-do-modulo` = executa o módulo.

## Testes
- `go test` = executa teste automatizado.
- `go test ./...` = executar testes em todos os pacotes.
- `go test -v` = executar testes em modo verboso.
- `go test -cover` = exibir porcentagem de cobertura de testes.
- `go test -coverprofile nome-do-arquivo.txt` = gera arquivo de cobertura de testes.
- `t.Parallel()` = rodar teste em paralelo.
