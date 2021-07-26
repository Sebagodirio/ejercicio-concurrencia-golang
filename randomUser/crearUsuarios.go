package randomUser

import (
	"fmt"
	"os"

	"github.com/Sebagodirio/ejercicio-usuario-asincronos/randomUser/usuario"
)

func CrearUsuariosSincrono() {

	cantidad := 1000000
	file, err := os.Create("usuarioSync.json")

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	usuario.EscribirArchivo(file, cantidad)

	err = file.Close()
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println("Archivo creado exitosamente")
}
