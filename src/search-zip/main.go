package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

type ViaCEP struct {
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

func main() {
	for _, zip := range os.Args {
		
		req, err := http.Get("viacep.com.br/ws/" + zip + "/json")

		if err != nil {
			fmt.Fprintf(os.Stderr, "Error with request: %v\v", err)
		}

		defer req.Body.Close()

		res, err := io.ReadAll(req.Body)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error reading the response: %v\n", err)
		}

		var data ViaCEP
		err = json.Unmarshal(res, &data)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error unmarshaling the response: %v\n", err)
		}

		file, err := os.Create("city.txt")
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error creating the file: %v\n", err)
		}

		defer file.Close()
		_, err = file.WriteString(fmt.Sprintf("Zip: %v, Location: %v\n", data.Cep, data.Localidade))
	}
}