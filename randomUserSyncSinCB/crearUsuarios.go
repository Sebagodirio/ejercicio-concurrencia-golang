package randomUserSyncSinCB

import (
	"fmt"
	"os"

	user "github.com/Sebagodirio/ejercicio-usuario-asincronos/randomUserSyncSinCB/usuario"
)

func CrearUsuariosSinCB() {
	cantidad := 1000000
	file, err := os.Create("usuarioAsyncSinCB.json")
	numeroDeGoroutines := 10

	cantidad = cantidad / numeroDeGoroutines
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	c := make(chan struct{})
	doneWrite := make(chan struct{})
	resultados := make(chan string)

	go func() {
		for resultado := range resultados {
			_, err = file.WriteString(resultado)
			if err != nil {
				panic(err)
			}
		}
		doneWrite <- struct{}{}
	}()

	for i := 0; i < numeroDeGoroutines; i++ {
		go user.EscribirArchivoV1(c, resultados, cantidad)
	}
	contador := 0
	for contador < numeroDeGoroutines {
		<-c //Espera que me llegue alg
		fmt.Println("Termino uno")
		contador++
	}
	close(resultados)
	<-doneWrite //Espera que me llegue alg
	fmt.Println("Archivo creado exitosamente")
}
