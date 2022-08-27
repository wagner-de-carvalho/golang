package app

import "github.com/urfave/cli"

//Retorna aplicação
func Gerar() *cli.App {
	app := cli.NewApp()
	app.Name = "Aplicação de linha de comando"
	app.Usage = "Busca IPs e nome de servidores na internet"
	return app
}