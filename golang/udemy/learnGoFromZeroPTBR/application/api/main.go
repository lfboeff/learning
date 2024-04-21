package main

import (
	"api_mod/src/config"
	"api_mod/src/router"
	"fmt"
	"log"
	"net/http"
)

// func init() {
// 	chave := make([]byte, 64)
// 	if _, err := rand.Read(chave); err != nil {
// 		log.Fatal(err)
// 	}
// 	stringBase64 := base64.StdEncoding.EncodeToString(chave)
// 	fmt.Println(stringBase64)
// }

func main() {

	config.Carregar()

	r := router.Gerar()

	fmt.Printf("Escutando a porta %d...\n", config.Porta)

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", config.Porta), r))
}
