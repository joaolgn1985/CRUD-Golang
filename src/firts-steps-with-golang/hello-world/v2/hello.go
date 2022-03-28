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

const monitoring = 2
const delay = 5

func main() {
	showIntroduction()

	for {
		showMenu()

		comando := readComand()

		switch comando {
		case 1:
			startMonitoring()
		case 2:
			fmt.Println("Showing logs...")
			printLogs()
		case 0:
			fmt.Println("Exiting of the system...")
			os.Exit(0)
		default:
			fmt.Println("Don't know this command")
			os.Exit(-1)
		}
	}
}

func showIntroduction() {
	name := "Joao Lobo"
	version := 1.2
	fmt.Println("Hi, sr.", name)
	fmt.Println("This system is in version", version)
}

func showMenu() {
	fmt.Println("1 - Starting Monitoring")
	fmt.Println("2 - Show Logs")
	fmt.Println("0 - Exit of the system")
}

func readComand() int {
	var comandRead int
	fmt.Scan(&comandRead)
	fmt.Println("The command chosen was", comandRead)
	fmt.Println("")

	return comandRead
}

func startMonitoring() {
	fmt.Println("Monitoring...")

	sites := readSitesOfFile()

	for i := 0; i < monitoring; i++ {
		for i, site := range sites {
			fmt.Println("Testing website", i, ":", site)
			testSite(site)
		}
		time.Sleep(delay * time.Second)
		fmt.Println("")
	}

	fmt.Println("")
}

func testSite(site string) {
	resp, err := http.Get(site)

	if err != nil {
		fmt.Println("An error has occurred:", err)
	}

	if resp.StatusCode == 200 {
		fmt.Println("Site:", site, "Has been uploaded successfully!")
		registryLog(site, true)
	} else {
		fmt.Println("Site:", site, "It has problems. Status Code:", resp.StatusCode)
		registryLog(site, false)
	}
}

func readSitesOfFile() []string {
	var sites []string

	file, err := os.Open("files.txt")
	if err != nil {
		fmt.Println("An error has occurred", err)
	}
	reader := bufio.NewReader(file)
	for {
		line, err := reader.ReadString('\n')
		line = strings.TrimSpace(line)

		sites = append(sites, line)

		fmt.Println(line)
		if err == io.EOF {
			break
		}
	}
	file.Close()
	return sites
}

func registryLog(site string, status bool) {

	file, err := os.OpenFile("log.txt", os.O_RDWR | os.O_CREATE | os.O_APPEND, 0666 )

	if err != nil {
		fmt.Println(err)
	}

	file.WriteString(time.Now().Format("02/01/2006 15:04:05") + " - " + site + "- online: " + strconv.FormatBool(status) + "\n")

	file.Close()
}

func printLogs(){
	file, err := ioutil.ReadFile("log.txt")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(file))
}