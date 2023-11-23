package main

// fmt = pacote de formatação e concatenação
// net/http = pacote de criação de servidores http online
import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

// Função que retorna 2 valores de tipos diferentes
func soma(x int, y int) (int, bool) {
	if x > 10 {
		return x + y, true
	} else {
		return x + y, false
	}
}

// Função para criar um Endpoint de uma API
func home(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello World!"))

	course_2 := Course{
		Name:        "Golang Course",
		Description: "Let's learn golang",
		Price:       100,
	}

	json.NewEncoder(w).Encode(course_2)
}

// Golang não tem classes nem objetos, e sim Estruturas (Structures)
// Course = Público (não é necessário digitar public)
// course = Privado (não é necessário digitar private)
type Course struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Price       int    `json:"price"`
}

// Criação de métodos, que é uma função que faz parte de uma struct
// GetFullInfo() é o nome do seu método
// C course serve como se fosse um "as", e é necessário especificar a struct que ele faz parte
func (c Course) GetFullInfo() string {
	return fmt.Sprintf("Name: %s, Description: %s, Price: %d", c.Name, c.Description, c.Price)
}

// Laço de Repetição FOR em Golang
func counter() {
	for i := 0; i < 10; i++ {
		fmt.Println(i)
		time.Sleep(time.Second)
	}
}

// Simulador de um Worker dentro de um seridor WEB
func worker(workerId int, data chan int) {
	for x := range data { // Vai ficar lendo tudo que cair no canal data
		fmt.Printf("Worker %d received %d\n", workerId, x)
		time.Sleep(time.Second)
	}
}

func main() {
	a := "Hello world"
	// o := só ocorre na primeira vez, na hora de declarar a variável
	// Go tem tipagem forte, ele verifica qual é o tipo da variavel que foi criada
	// Não é possível alterar o tipo dela futuramente

	resultado, status := soma(10, 20)
	// Tem como retornar 2 valores em uma unica função, como nesse exemplo
	// Também é possível declarar 2 variáveis simultaneamente

	println(a)
	fmt.Println("O resultado é", resultado, "e o status é", status)

	// Como criar um servidor HTTP
	http.HandleFunc("/", home)
	http.ListenAndServe(":8080", nil)

	// Criação de uma variável baseada em uma structure
	// Ocorre uma validaçao da tipagem igual no typescript
	course := Course{
		Name:        "Golang Course",
		Description: "Let's learn golang",
		Price:       100,
	}

	course.Price = 200

	fmt.Println(course.GetFullInfo())
	fmt.Println(course)
	fmt.Println(course.Name)

	// Go Routines (Sistema de Threads, Carretel)
	go counter() // Thread 1
	go counter() // Thread 2
	counter()

	// Sistema de Channels (Aprimorar o sistema de Routines)
	// Evitar Race Conditions e Sincronismo

	channel := make(chan string)
	// Esse canal vai fazer uma comunicaçao entre uma thread e outra thread
	// Vai fazer um paralelo entre cada thread
	go func() {
		channel <- "Hello World"
	}()
	// Esvaziar o channel pegando o valor que a thread passou
	fmt.Println(<-channel)

	// O "i" vai ser jogado no canal e o canal vai passar para a Thread 2
	// Vai ler o canal, vai aguardar esvaziar o canal, e passar o valor do I para o canal novamente
	channel_worker := make(chan int)
	go worker(1, channel_worker) // Thread 2
	go worker(2, channel_worker) // Thread 3
	// Se usar 2 go routines, vai agilizar o FOR, ao invés de 10 seg vai ser só 5
	// Workers funcionam como funcionarios, quanto mais trabalhando mais consumo de memória
	// Isso se deve ao fato de estar utilizando o mesmo channel
	// Porém maior velocidade de processamento de dados (2kb por worker)

	// Toda vez que o canal é lido é esvaziado, quando está cheio não consigo enviar valores
	for i := 0; i < 10; i++ {
		channel_worker <- i
	}

	// Como usar varios workers
	for i:= 0; i < 10000; i++ {
		go worker(i, channel_worker)
	}

	// Usando 10 mil workers para 100 mil requisiçÕes, vai ficar muito mais rápido a runtime da função
	// Vai usar apenas 19b ao invés de usar o core normal do processador, que exigira 1GB
	for i := 0; i < 10000; i++ {
		channel_worker <- i
	}

	// O GO Routine é semelhante ao concurrentBag, não corre o risco de race condition
	// Nenhum worker vai mudar diretamente um valor e encerrar outro processo
}
