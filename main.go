package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

type BrasilAPIResponse struct {
	CEP          string `json:"cep"`
	State        string `json:"state"`
	City         string `json:"city"`
	Neighborhood string `json:"neighborhood"`
	Street       string `json:"street"`
}

type ViaCEPResponse struct {
	CEP         string `json:"cep"`
	Logradouro  string `json:"logradouro"`
	Complemento string `json:"complemento"`
	Bairro      string `json:"bairro"`
	Localidade  string `json:"localidade"`
	UF          string `json:"uf"`
}

func requestAPI(url string, ch chan<- map[string]interface{}, apiName string) {
	client := http.Client{}

	resp, err := client.Get(url)
	if err != nil {
		ch <- map[string]interface{}{
			"error":   err,
			"apiName": apiName,
		}
		return
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		ch <- map[string]interface{}{
			"error":   err,
			"apiName": apiName,
		}
		return
	}

	ch <- map[string]interface{}{
		"body":    body,
		"apiName": apiName,
	}
}

func main() {
	cep := "01153000"

	ch := make(chan map[string]interface{}, 2)

	go requestAPI("https://brasilapi.com.br/api/cep/v1/"+cep, ch, "BrasilAPI")
	go requestAPI("http://viacep.com.br/ws/"+cep+"/json/", ch, "ViaCEP")

	select {
	case result := <-ch:
		if err, ok := result["error"]; ok {
			fmt.Printf("Erro ao acessar %s: %v\n", result["apiName"], err)
		} else {
			apiName := result["apiName"].(string)
			body := result["body"].([]byte)
			fmt.Printf("Resposta da %s:\n", apiName)

			if apiName == "BrasilAPI" {
				var data BrasilAPIResponse
				if err := json.Unmarshal(body, &data); err != nil {
					fmt.Printf("Erro ao decodificar resposta: %v\n", err)
					return
				}
				fmt.Printf("CEP: %s\nEstado: %s\nCidade: %s\nBairro: %s\nRua: %s\n",
					data.CEP, data.State, data.City, data.Neighborhood, data.Street)
			} else if apiName == "ViaCEP" {
				var data ViaCEPResponse
				if err := json.Unmarshal(body, &data); err != nil {
					fmt.Printf("Erro ao decodificar resposta: %v\n", err)
					return
				}
				fmt.Printf("CEP: %s\nEstado: %s\nCidade: %s\nBairro: %s\nRua: %s\n",
					data.CEP, data.UF, data.Localidade, data.Bairro, data.Logradouro)
			}
		}
	case <-time.After(1 * time.Second):
		fmt.Println("Erro de timeout: nenhuma resposta em 1 segundo.")
	}
}
