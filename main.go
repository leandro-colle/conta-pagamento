package main

import (
	"conta-pagamento/database"
	"conta-pagamento/routes"
)

func main() {
	database.Connection()
	routes.HandleRequest()
}
