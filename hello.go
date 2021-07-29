package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
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
			readLog()
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
	fmt.Println("")
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
	fmt.Println("")
	fmt.Println("Iniciando monitoramento...")
	sites := readTxtFileAndReturnValues()

	for i := 0; i < monitoring; i++ {
		fmt.Println("----------------")

		for i, site := range sites {
			fmt.Println(i, "- Nome do site:", site)
			testSite(site)
		}
		time.Sleep(delay * time.Second)

	}

}

func testSite(site string) {

	resp, err := http.Get(site)

	if err != nil {
		fmt.Println("Ocorreu um erro", err)
	}

	if resp.StatusCode == 200 {
		writeLog(site, true)
		fmt.Println("Site:", site, "está funcionando corretamente!")
	} else {
		writeLog(site, false)
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

	file.Close()

	return sites
}

func writeLog(site string, status bool) {

	file, err := os.OpenFile("log.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)

	if err != nil {
		fmt.Println("Ocorreu um erro", err)
	}

	file.WriteString(time.Now().Format("02/01/2006 15:04:05") + " - O site " + site + " está funcionando corretamente e o status da requisição foi: " + strconv.FormatBool(status) + "\n")

	file.Close()
}

func readLog() {

	fmt.Println("Exibindo logs...")

	file, err := ioutil.ReadFile("log.txt")

	if err != nil {
		fmt.Println("Ocorreu um erro", err)
	}

	fmt.Println(string(file))
}
