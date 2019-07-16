package main

import (
	"bufio"
	"os"
	"fmt"
	"log"
	"sort"
)


func checkError(err error){
	if err != nil {
		log.Fatal(err)
	}
}


func lerArquivo(path string) ([]string,error){
	file,err :=  os.Open(path)

	checkError(err)

	defer file.Close()

	linhas := []string{}

	scan := bufio.NewScanner(file)

	for scan.Scan(){
		linhas = append(linhas,scan.Text())
	}

	return linhas, scan.Err()
}


func main(){
	dados, err := lerArquivo("arquivo.txt")
	checkError(err)
	sort.Strings(dados)

	fmt.Println(dados)
}