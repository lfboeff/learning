package autenticacao

import (
	"api_mod/src/config"
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

// CriarToken retorna um token assinado com as permissões do usuário
func CriarToken(usuarioID uint64) (string, error) {

	permissoes := jwt.MapClaims{}

	permissoes["authorized"] = true
	permissoes["exp"] = time.Now().Add(time.Hour * 6).Unix()
	permissoes["usuarioId"] = usuarioID

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, permissoes)

	return token.SignedString(config.SecretKey) // secret

}

// ValidarToken verifica se o token passado na request é válido
func ValidarToken(r *http.Request) error {

	tokenString := extrairToken(r)

	token, err := jwt.Parse(tokenString, retornarChaveDeVerificacao)
	if err != nil {
		return err
	}

	// fmt.Println(token)

	if _, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return nil
	}

	return errors.New("token inválido")
}

// ExtrairUsuarioID retorna o usuarioId que está salvo no token
func ExtrairUsuarioID(r *http.Request) (uint64, error) {

	tokenString := extrairToken(r)

	token, err := jwt.Parse(tokenString, retornarChaveDeVerificacao)
	if err != nil {
		return 0, err
	}

	if permissoes, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {

		usuarioId, err := strconv.ParseUint(fmt.Sprintf("%0.f", permissoes["usuarioId"]), 10, 64)
		if err != nil {
			return 0, err
		}
		return usuarioId, nil
	}

	return 0, errors.New("token inválido")

}

func extrairToken(r *http.Request) string {

	token := r.Header.Get("Authorization") // retorna: "Bearer askdsdgsdgfff(<=token)"

	if len(strings.Split(token, " ")) == 2 {
		return strings.Split(token, " ")[1]
	}

	return ""
}

func retornarChaveDeVerificacao(token *jwt.Token) (interface{}, error) {

	if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
		return nil, fmt.Errorf("método de assinatura inesperado: %v", token.Header["alg"])
	}

	return config.SecretKey, nil
}
