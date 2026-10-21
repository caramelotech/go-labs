// Package saudacao mostra nomes curtos: o pacote é o prefixo natural
// dos identificadores exportados (saudacao.New, saudacao.Ola).
package saudacao

// Mensagem carrega um texto de saudação.
// Não se chama SaudacaoMensagem - quem importa já escreve saudacao.Mensagem.
type Mensagem struct {
	texto string
}

// New cria uma Mensagem. Clientes veem saudacao.New, não saudacao.NewMensagem.
func New(texto string) Mensagem {
	return Mensagem{texto: texto}
}

// Texto devolve o conteúdo. Em Go o getter não leva prefixo Get.
func (m Mensagem) Texto() string {
	return m.texto
}

// Ola monta uma saudação. Lê-se bem: saudacao.Ola("Caramelo").
func Ola(nome string) string {
	return "Olá, " + nome + "!"
}
