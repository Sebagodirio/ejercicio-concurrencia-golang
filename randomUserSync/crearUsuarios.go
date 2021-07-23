package randomUserSync

import (
	"fmt"
	"os"

	"github.com/SebaGodirio/ejercicio-usuarios-asincronos/randomUserSync/user"
)

func CrearUsuario() {
	c := make(chan struct{})
	cantidad := 1000000
	file, err := os.Create("usuariosAsync.json")
	numeroDeGoroutines := 10

	cantidad = cantidad / numeroDeGoroutines
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	for i := 0; i < 10; i++ {
		go user.EscribirArchivo(file, c, cantidad)
	}
	contador := 0
	for contador < numeroDeGoroutines {
		<-c //Esperar informacion del canal
		fmt.Println("Termino uno")
		contador++
	}
	err = file.Close()
	if err != nil {
		fmt.Println("Algun error")
	}
	fmt.Println("Archivo creado exitosamente")
}
