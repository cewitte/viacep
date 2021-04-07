package main

import (
	"fmt"
	"log"

	"github.com/cewitte/viacep/models"
)

func main() {

	addr := models.Address{}
	zip := "01452-000"
	err := addr.Fill(&zip)
	if err != nil {
		log.Panic(err)
	}
	fmt.Println("---------------------------------------------------------")
	fmt.Printf("Street: \t\t%s\n", addr.Logradouro)
	fmt.Printf("Neighborhood: \t\t%s\n", addr.Bairro)
	fmt.Printf("City: \t\t\t%s\n", addr.Localidade)
	fmt.Printf("State: \t\t\t%s\n", addr.Uf)
	fmt.Printf("Zip Code: \t\t%s\n", addr.Cep)
	fmt.Println("---------------------------------------------------------")
}
