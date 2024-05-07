package main

import (
	"bufio"
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
)

const monitoramentos = 5
const delay = 5

func main() {

	exibeIntroducao()
	for {

		exibemenu()

		// if comando == 1 {
		// 	fmt.Println("Iniciando Monitoramento...")
		// } else if comando == 2 {
		// 	fmt.Println("Exibindo Logs...")
		// } else if comando == 0 {
		// 	fmt.Println("Saindo do Programa...")
		// } else {
		// 	fmt.Println("Opcão Invalida...")
		// }

		// switch não precisa de break
		// comando := leComando() forma curta de declarar uma variável

		comando := leComando()

		switch comando {
		case 1:
			iniciarMonitoramento()
		case 2:
			fmt.Println("Exibindo Logs...")
			imprimeLogs()
		case 0:
			fmt.Println("Saindo do Programa...")
			os.Exit(0)
		default:
			fmt.Println("Opcão Invalida...")
			os.Exit(-1)
		}
	}
}

func exibeIntroducao() {
	nome := "Marcos"
	versao := 1.1
	fmt.Println("Ola, sr.", nome)
	fmt.Println("Este programa está na versão", versao)
}

func leComando() int {
	var comandoLido int
	fmt.Scanf("%d", &comandoLido)
	fmt.Println("O valor da variável comando é:", comandoLido, &comandoLido)

	return comandoLido
}

func exibemenu() {
	fmt.Println("1- Iniciar Monitoramento")
	fmt.Println("2- Exibir Logs")
	fmt.Println("0- Sair do Programa")
}

// func iniciaMonitoramento() {
// 	fmt.Println("Monitorando...")
// 	site := "https://www.alura.com.br"
//espera mais de uma varievel, _ = http.Get(site) se for para ignorar alguma variavel
//onde a função contem 2.
// 	resp, _ := http.Get(site)

// 	fmt.Println("Status:", resp.StatusCode)

// }

func iniciarMonitoramento() {
	fmt.Println("Monitorando...")
	// site com URL inexistente
	// var sites [4]string
	// sites[0] = "https://httpbin.org/status/404"
	// sites[1] = "https://httpbin.org/status/200"
	// sites[2] = "https://httpbin.org/status/500"
	// sites[3] = "https://httpbin.org/status/200"

	// sites := []string{"https://httpbin.org/status/404", "https://httpbin.org/status/200", "https://httpbin.org/status/500", "https://httpbin.org/status/200"}

	// for i := 0; i < len(sites); i++ {

	// }

	sites := leSitesDoArquivo()

	for i := 0; i < monitoramentos; i++ {
		for i, site := range sites {
			fmt.Println("Testando site", i, ":", site)
			testeSite(site)
		}

		time.Sleep(time.Second * delay)
	}

	// site := "https://httpbin.org/status/404" // ou 200

}

func testeSite(site string) {
	resp, err := http.Get(site)

	if err != nil {
		fmt.Println("Ocorreu um erro:", err)
	}

	if resp.StatusCode == 200 {
		fmt.Println("Site:", site, "foi carregado com sucesso!")
		registraLog(site, true)
	} else {
		fmt.Println("Site:", site, "está com problemas. Status Code:", resp.StatusCode)
		registraLog(site, false)
	}
}

func leSitesDoArquivo() []string {

	var sites []string

	arquivo, err := os.Open("sites.txt")
	if err != nil {
		fmt.Println("Ocorreu um erro:", err)
	}

	leitor := bufio.NewReader(arquivo)

	for {
		linha, err := leitor.ReadString('\n')
		linha = strings.TrimSpace(linha)
		sites = append(sites, linha)
		if err == io.EOF {
			break
		}
	}

	arquivo.Close()

	return sites
}

func registraLog(site string, status bool) {

	arquivo, err := os.OpenFile("log.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		fmt.Println("Ocorreu um erro:", err)
	}

	arquivo.WriteString(time.Now().Format("02/01/2006 15:04:05") + " - " + site + " - online: " + strconv.FormatBool(status) + "\n")

	arquivo.Close()
}

func imprimeLogs() {
	arquivo, err := os.ReadFile("log.txt")
	if err != nil {
		fmt.Println("Ocorreu um erro:", err)
	}

	fmt.Println(string(arquivo))
}
