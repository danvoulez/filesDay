package main

import (
	"fmt"
	"log"
	"os"

	"logline-runtime/auditor"
	"logline-runtime/interpreter"
	"logline-runtime/parser"
	"logline-runtime/plugins"
	"logline-runtime/registry"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Uso: logline-runtime <caminho_do_arquivo_logline_no_db>")
		os.Exit(1)
	}

	dslPath := os.Args[1]
	log.Printf("Iniciando o Runtime LogLine para o arquivo: %s", dslPath