// Demonstra nomes de pacote e de identificadores no estilo do Effective Go:
// curtos, sem gaguejar com o nome do pacote, getters sem Get.
package main

import (
	"fmt"

	"github.com/caramelotech/go-labs/examples/12-nomes-de-pacotes/saudacao"
)

func main() {
	fmt.Println(saudacao.Ola("Caramelo"))

	m := saudacao.New("bem-vindo ao lab")
	fmt.Println("texto:", m.Texto())
}
