package main

import "example/service/cmd"

//	@title						Example Service - Open API
//	@securityDefinitions.apikey	Bearer
//	@in							header
//	@name						Authorization
func main() {
	cmd.Execute()
}