package main

import (
	"bufio"
	"compress/gzip"
	"fmt"
	"os"
)

var manPaths = map[string]string{
	"main":  "/usr/share/man/",
	"local": "/usr/local/share/man/",
}

func main() {

	args := os.Args

	// TODO: create parse args function
	if len(args) < 2 {
		fmt.Println("What manual do you want?")
		fmt.Println("Usage: mano <command> or mano -D <documentation>")
	}

	command := args[1]

	// TODO: verify manual existence
	manFilePath := manPaths["main"] + "man1/" + command + ".1.gz"

	if _, err := os.Stat(manFilePath); err != nil {
		fmt.Println("Error", err)
		return
	}

	// TODO: read file
	gzFile, err := os.Open(manFilePath)
	if err != nil {
		fmt.Println("fail reading man file", err)
		return
	}
	defer gzFile.Close()

	gzReader, err := gzip.NewReader(gzFile)
	if err != nil {
		fmt.Println("Error al crear el lector gzip:", err)
		return
	}
	defer gzReader.Close()

	// TODO: create model
	lectorBuffer := bufio.NewReader(gzReader)
	for {
		linea, err := lectorBuffer.ReadString('\n')
		if err != nil {
			// Verificar si se alcanzó el final del archivo
			if err.Error() == "EOF" {
				break
			}
			fmt.Println("Error al leer el archivo:", err)
			return
		}
		fmt.Print(linea)
	}
}
