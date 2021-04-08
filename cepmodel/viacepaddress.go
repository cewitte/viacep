package cepmodel

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/cewitte/viacep/cepcleaners"
)

var viacepURL = "https://viacep.com.br/ws/zipcode/json/"

type ViaCepAddress struct {
	Cep         string `json:"cep"`
	Logradouro  string `json:"logradouro"`
	Complemento string `json:"complemento"`
	Bairro      string `json:"bairro"`
	Localidade  string `json:"localidade"`
	Uf          string `json:"uf"`
	Ibge        string `json:"ibge"`
	Gia         string `json:"gia"`
	Ddd         string `json:"ddd"`
	Siafi       string `json:"siafi"`
}

func NewAddressFromCEP(cep *string) (*ViaCepAddress, error) {
	// Clean zip code first.
	*cep = cepcleaners.ExtractNumbers(*cep)

	customURL := strings.Replace(viacepURL, "zipcode", *cep, -1)
	method := "GET"

	client := &http.Client{}
	req, err := http.NewRequest(method, customURL, nil)

	if err != nil {
		return nil, err
	}

	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	var addr ViaCepAddress
	err = json.Unmarshal(body, &addr)
	if err != nil {
		return nil, err
	}

	return &addr, nil
}
