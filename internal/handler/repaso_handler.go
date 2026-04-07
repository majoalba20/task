package handler

import (
	"context"
	"errors"
	"fmt"
	"os"
	"time"
)

// Struct
type Perro struct {
	Raza  string
	Color string
}

type Gato struct {
	Edad uint
}

func DataPerro(p Perro) {
	fmt.Printf("El perro es de raza: %s, y color: %s\n", p.Raza, p.Color)
}

// Interfaz
type Animal interface {
	Sonido() string
}

func (p Perro) Sonido() string {
	return "Guau"
}

func (g Gato) Sonido() string {
	return "Miau"
}

func HacerSonido(a Animal) {
	fmt.Println(a.Sonido())
}

// Declaracion de variables
func DeclaracionVariables() {
	var nombre string = "Pepito"
	var nombre1 = "Juan"
	nombre2 := "Jaime"

	fmt.Println(nombre)
	fmt.Println(nombre1)
	fmt.Println(nombre2)

	nombre2 = "Lola"
	fmt.Println(nombre2)
}

// punteros
func Punteros() {
	edad := 25
	p := &edad

	fmt.Println("edad:", edad)
	fmt.Println("puntero:", p)         //direccion en memoria
	fmt.Println("valor apuntado:", *p) //valor que hay en la direccion de memoria

	// Sin punteros se hace una copia!!!!
	num := 10
	cambiar(num) // Se pasa una copia de num mas no el original
	fmt.Println("sin punteros:", num)

	// Con punteros
	num1 := 10
	cambiar1(&num1)
	fmt.Println("con punteros:", num1)

}

func cambiar(x int) {
	x = 100
}

func cambiar1(x *int) {
	*x = 100
}

// Goroutines
func GoRoutines() {
	// Puede terminar antes de que la goroutine alcance a ejecutarse.
	go saludar()
	// Por eso se adicionan esperas
	time.Sleep(1 * time.Second) // Esperar un segundo antes de seguir
	fmt.Println("Hola desde la funcion de GoRoutine")
}

func saludar() {
	fmt.Println("Hola desde una goroutine")
}

// channels
func Channels() {
	ch := make(chan int) // un canal de enteros

	go func() {
		ch <- 42 // manda "42" al canal
	}()

	// recibe de un channel
	// bloquea el programa hasta que alguien mande algo
	valor := <-ch
	fmt.Println(valor)
}

// context cancelacion
func ContextCancelado() {
	// contexto cancelable
	ctx, cancel := context.WithCancel(context.Background())
	go trabajoCancelado(ctx)
	time.Sleep(2 * time.Second)
	cancel()
	time.Sleep(1 * time.Second)
}

func trabajoCancelado(ctx context.Context) {
	for {
		select {
		// Esta pendiente de que se cancele el contexto
		case <-ctx.Done():
			fmt.Println("Trabajo cancelado")
			return
		default:
			fmt.Println("Trabajando...")
			time.Sleep(500 * time.Millisecond)
		}
	}
}

// context timeout
func ContextTimeOut() {
	// contexto con timeout
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	trabajo(ctx)
}

func trabajo(ctx context.Context) {
	select {
	case <-time.After(5 * time.Second):
		fmt.Println("Trabajo terminado")
	case <-ctx.Done():
		fmt.Println("Trabajo cancelado por timeout")
	}
}

// manejo de errores
func ManejoError() {
	archivo, err := os.Open("datos.txt")
	if err != nil {
		err := errors.New("No se pudo abrir el archivo")
		fmt.Println("Error:", err)
		return
	}
	defer archivo.Close()

	fmt.Println("Archivo abierto:", archivo.Name())
}

func Imprimir[T any](x T) {
	fmt.Println(x)
}

func Defer() {
	defer fmt.Println("Esto va al final")

	fmt.Println("Primero esto")
	fmt.Println("Luego esto")
}
