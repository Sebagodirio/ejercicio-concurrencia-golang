package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/Sebagodirio/ejercicio-usuario-asincronos/randomUser"
	"github.com/Sebagodirio/ejercicio-usuario-asincronos/randomUserSync"
	"github.com/Sebagodirio/ejercicio-usuario-asincronos/randomUserSyncSinCB"
)

func main() {
	menu :=
		`
	[ 1 ] Crear lista de usuarios sincronamente
	[ 2 ] Crear lista de usuarios asincronamente
	[ 3 ] Crear lista de usuarios asincronamente sin cuello de botella
`
	fmt.Print(menu)

	reader := bufio.NewReader(os.Stdin)

	entrada, _ := reader.ReadString('\n')          // Leer hasta el separador de salto de línea
	eleccion := strings.TrimRight(entrada, "\r\n") // Remover el salto de línea de la entrada del usuario

	switch eleccion {
	case "1":
		randomUser.CrearUsuariosSincrono()
	case "2":
		randomUserSync.CrearUsuario()
	case "3":
		randomUserSyncSinCB.CrearUsuariosSinCB()
	default:
		fmt.Println("Opcion incorrecta")
	}

}
