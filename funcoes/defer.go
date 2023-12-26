package main

import ("fmt"
		"os"
)

func dia1() {
	fmt.Println("Domingo!")
}

func dia2() {
	fmt.Println("segunda-feira!")
}

func main() {

	defer dia2()
	dia1()

	f := createFile("C:/Users/aliss/AppData/Local/Temp/defer.txt")
    defer closeFile(f)
    writeFile(f)

}		
func createFile(p string) *os.File {
    fmt.Println("creating")
    f, err := os.Create(p)
    if err != nil {
        panic(err)
    }
    return f
}
func writeFile(f *os.File) {
    fmt.Println("writing")
    fmt.Fprintln(f, "data")
}
//É importante verificar erros ao fechar um arquivo, mesmo em uma função adiada.

func closeFile(f *os.File) {
    fmt.Println("closing")
    err := f.Close()
    if err != nil {
        fmt.Fprintf(os.Stderr, "error: %v\n", err)
        os.Exit(1)
    }
}


