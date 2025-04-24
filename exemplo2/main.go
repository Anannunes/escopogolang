package main 
import "fmt" 
var nomeEscola = "Escola Técnica SENAI" // variável global 
func main() { 
nome := "Ana" 
idade:= 16 
mensagem:= boasVindas (nome) 
fmt.Println(mensagem) 
status := verificaMaioridade(idade) 
fmt.Println(status) 
}