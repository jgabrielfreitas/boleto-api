package main

import (
	"fmt"

	"os"

	"bitbucket.org/mundipagg/boletoapi/api"
	"bitbucket.org/mundipagg/boletoapi/config"
	"bitbucket.org/mundipagg/boletoapi/log"
)

// BB
type bbObj struct {
	convenio   convenioBb
	boleto     boletoBb
	pagador    pagadorBb
	controleBb controleBb
}

type boletoBb struct {
	codigoModalidadeTitulo               int
	dataEmissaoTitulo                    string
	dataVencimentoTitulo                 string
	valorOriginalTitulo                  int64
	codigoTipoDesconto                   int16
	codigoTipoJuroMora                   int16
	codigoTipoMulta                      int16
	codigoAceiteTitulo                   string
	codigoTipoTitulo                     int
	indicadorPermissaoRecebimentoParcial string
	textoNumeroTituloCliente             string
}

type pagadorBb struct {
	codigoTipoInscricaoPagador   int16
	numeroInscricaoPagadorNúmero int
	nomePagador                  string
	textoEnderecoPagador         string
	numeroCepPagador             int
	nomeMunicipioPagador         string
	nomeBairroPagador            string
	siglaUfPagador               string
}

type controleBb struct {
	codigoChaveUsuario         string // PIC X(08) WTF?
	codigoTipoCanalSolicitacao int
}

type convenioBb struct {
	numeroConvenio         int
	numeroCarteira         int
	numeroVariacaoCarteira int
}

// Caixa
type caixaObj struct {
	unidade             string
	identificadorOrigem string
	nossoNumero         int64
	tipoEspecie         string // talvez não seja obrigatório
	caixaAuth           caixaAuth
}

type caixaAuth struct {
	codigoCedente    string
	nossoNumero      string
	dataDeVencimento string
	valor            int64
	cpfCnpj          string
}

// Santander
type santanderObj struct {
	dados     dadosSantander
	expiracao int
	sistema   string
}

type dadosSantander struct {
	convenioSantander convenioSantander
	pagadorSantander  pagadorSantander
	tituloSantander   tituloSantander
	mensagem          string
}

type convenioSantander struct {
	codBanco    string
	codConvenio string
}

type pagadorSantander struct {
	tpDoc    int
	numDoc   int64
	nome     string
	endereco string
	bairro   string
	cidade   string
	uf       string
	cep      int
}

type tituloSantander struct {
	nossoNumero        int64
	seuNumero          string
	dataVencimento     string
	dataEmissao        string
	especie            string
	valorNominal       int64
	pcMulta            int
	qtDiasMulta        int
	pcJuro             int
	tpDesc             int
	valorDesconto      int64
	dataLimiteDesconto string
	valorAbatimento    int64
	tpProtesto         int
	qtDiasProtesto     int
	tituloQtDiasBaixa  int
}

func main() {
	logo()
	main1()
	err := log.Install()
	if err != nil {
		fmt.Println("Log SEQ Fails")
		os.Exit(-1)
	}
	api.InstallRestAPI()
}

func logo() {
	l := `
  ____        _      _                      _ 
 |  _ \      | |    | |         /\         (_)
 | |_) | ___ | | ___| |_ ___   /  \   _ __  _ 
 |  _ < / _ \| |/ _ \ __/ _ \ / /\ \ | '_ \| |
 | |_) | (_) | |  __/ || (_) / ____ \| |_) | |
 |____/ \___/|_|\___|\__\___/_/    \_\ .__/|_|
                                     | |      
                                     |_|      
	`
	fmt.Println(l)
	fmt.Println("Version: " + config.Get().Version)
}
