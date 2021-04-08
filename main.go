package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/cewitte/viacep/cepmodel"
)

func main() {
	zip := flag.String("cep", "01452-000", "-cep=\"01452-000\"")
	flag.Parse()
	addr, err := cepmodel.New(zip)

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
