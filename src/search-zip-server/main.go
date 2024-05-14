package main

import (
	"encoding/json"
	"io"
	"net/http"
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
	http.HandleFunc("/", SearchZipHandler)
	http.ListenAndServe(":8080", nil)
}

func SearchZipHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	zipParam := r.URL.Query().Get("zip")
	if zipParam == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	zip, err := SearchZip(zipParam)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(zip)
}

func SearchZip(zip string) (*ViaCEP, error) {
	res, err := http.Get("viacep.com.br/ws/" + zip + "/json")
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	var z ViaCEP
	err = json.Unmarshal(body, &z)
	if err != nil {
		return nil, err
	}
	return &z, nil
}