package main

import "github.com/igor-ferreira-almeida/user-crud/infrastructure/server"

// @title User Crud API
// @version 1.0
// @description This is a generic user crud.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host petstore.swagger.io
// @BasePath /
func main() {
	server.Start()
}
