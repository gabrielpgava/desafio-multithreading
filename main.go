package main

import (
	"context"
	"fmt"
	"io"
	"math/rand"
	"net/http"
	"time"
)



func main(){

	cep :="01153000"


	result := make(chan string)

	go brasilApi(cep, result)
	go viaCep(cep, result)


	fmt.Println("A api mais rapida foi:")
	fmt.Println(<-result)

}


func brasilApi(cep string, result chan string) (string, error) {
	time.Sleep(time.Duration(1+rand.Intn(2)) * time.Second)

	ctx, cancel:= context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, "https://brasilapi.com.br/api/cep/v1/"+cep, nil)
	if err != nil {
		fmt.Println("Erro ao criar o request:", err)
		return "", err
	}


	data, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Println("Erro ao fazer a requisição:", err)
		return "", err
	}
	defer data.Body.Close()

	body, err := io.ReadAll(data.Body)
	if err != nil {
		fmt.Println("Erro ao ler o corpo da resposta:", err)
		return "", err
	}

	texto := "Resposta da API BrasilAPI: " + string(body)
	result <- texto

	return "", nil
}

func viaCep(cep string, result chan string) (string, error) {
	time.Sleep(time.Duration(1+rand.Intn(2)) * time.Second)

	ctx, cancel:= context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, "http://viacep.com.br/ws/"+cep+"/json/", nil)
	if err != nil {
		fmt.Println("Erro ao criar o request:", err)
		return "", err
	}


	data, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Println("Erro ao fazer a requisição:", err)
		return "", err
	}
	defer data.Body.Close()

	body, err := io.ReadAll(data.Body)
	if err != nil {
		fmt.Println("Erro ao ler o corpo da resposta:", err)
		return "", err
	}


	texto := "Resposta da API ViaCep: " + string(body)
	result <- texto

	return "", nil
}