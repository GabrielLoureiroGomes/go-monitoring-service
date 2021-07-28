package main

import (
	"bufio"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
	"time"
)

const monitoring = 3
const delay = 5

func main() {

	for {
		showMenu()

		result := scanCommand()

		switch result {
		case 1:
			startMonitoring()
		case 2:
			fmt.Println("Exibindo logs...")
		case 0:
			fmt.Println("Saindo do programa...")
			os.Exit(0)
		default:
			fmt.Println("Comando inválido!")
			os.Exit(-1)
		}
	}

}

func showMenu() {
	fmt.Println("1 - Iniciar monitoramento")
	fmt.Println("2 - Exibir logs")
	fmt.Println("0 - Sair do programa")
}

func scanCommand() int {
	var command int
	fmt.Scan(&command)

	return command
}

func startMonitoring() {
	fmt.Println("Iniciando monitoramento...")
	sites := readTxtFileAndReturnValues()

	for i := 0; i < monitoring; i++ {
		fmt.Println("----------------")

		for i, site := range sites {
			fmt.Println(i, "- Nome do site:", site)
			testaSite(site)
		}
		time.Sleep(delay * time.Second)

	}

}

func testaSite(site string) {

	resp, err := http.Get(site)

	if err != nil {
		fmt.Println("Ocorreu um erro", err)
	}

	if resp.StatusCode == 200 {
		fmt.Println("Site:", site, "está funcionando corretamente!")
	} else {
		fmt.Println("Site:", site, "está fora do ar! Status Code:", resp.StatusCode)
	}
}

func readTxtFileAndReturnValues() []string {
	var sites []string

	file, err := os.Open("sites.txt")

	if err != nil {
		fmt.Println("Ocorreu um erro", err)
	}

	leitor := bufio.NewReader(file)

	for {

		line, err := leitor.ReadString('\n')
		line = strings.TrimSpace(line)

		sites = append(sites, line)

		if err == io.EOF {
			break
		}

	}

	return sites
}
