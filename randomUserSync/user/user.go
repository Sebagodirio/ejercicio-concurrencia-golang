package user

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"os"
)

type Usuario struct {
	Nombre   string
	Apellido string
	Email    string
}

var apellidos []string
var nombres []string

var usuarios []Usuario

func init() {
	generarApellidos()
	generarNombres()
}

func generarApellidos() {
	apellidos = append(apellidos, "Lopez", "Juarez", "González", "Rodriguez", "Gomez", "Fernandez", "Diaz", "Martinez", "Perez", "Garcia", "Sanchez", "Romero", "Sosa", "Torres", "Alvarez", "Ruiz", "Ramirez", "Flores", "Benitez", "Acosta", "Medina", "Herrera")
}

func generarNombres() {
	nombres = append(nombres, "Sebastian", "Nicolas", "Jorge", "Camila", "Cecilia", "Florencia", "Omar", "Oscar", "Beatriz", "Morena", "Atenea", "Elizabeth", "Fernando", "Carolina", "Olivia", "Micaela", "Rocio", "Tobias", "Gabriel", "Renata", "Lorena", "Fernanda")
}
func generarApellido(nombre, apellido string) string {
	return nombre + "." + apellido + "@gmail.com"
}

func generarUsuarios(cantidad int) []Usuario {
	largoNombres := len(nombres)
	largoApellidos := len(apellidos)
	var nombre string
	var apellido string
	for i := 0; i < cantidad; i++ {
		nombre = nombres[rand.Intn(largoNombres)]
		apellido = apellidos[rand.Intn(largoApellidos)]
		usuarios = append(usuarios, Usuario{
			Nombre:   nombre,
			Apellido: apellido,
			Email:    generarApellido(nombre, apellido),
		})
	}
	return usuarios
}

func EscribirArchivo(file *os.File, c chan struct{}, cantidad int) {
	var user Usuario
	generarUsuarios(cantidad)

	for i := 0; i < cantidad; i++ {
		user = Usuario{
			Nombre:   usuarios[i].Nombre,
			Apellido: usuarios[i].Apellido,
			Email:    usuarios[i].Email,
		}
		datos, _ := json.Marshal(user)            //Lo paso a json
		_, err := file.WriteString(string(datos)) //Casteo a string y lo escribo en el archivo
		if err != nil {
			fmt.Println(err.Error())
			return
		}
	}

	c <- struct{}{}
}
