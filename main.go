package main

import (
	"database/sql"
	"fmt"

	// "log"
	_ "github.com/denisenkom/go-mssqldb"
	_ "github.com/go-sql-driver/mysql"
)

func crearUsuario(db *sql.DB, nombre string, apellido string, edad int, peso float32) (int64, error) {
	tsql := fmt.Sprintf("INSERT INTO usuario (nombre, apellidos, edad, peso ) VALUES ('%s','%s','%d','%f');",
		nombre, apellido, edad, peso)
	result, err := db.Exec(tsql)
	if err != nil {
		fmt.Println("Error al crear usuario: " + err.Error())
		return -1, err
	}
	return result.LastInsertId()
}

func verUsuarios(db *sql.DB) {

	rows, err := db.Query("SELECT * FROM usuario")
	if err != nil {

		fmt.Println("Error al ejecutar la consulta:", err)
		return
	}
	defer rows.Close()

	for rows.Next() {
		var id, edad int
		var nombre, apellido string
		var peso float32

		err := rows.Scan(&id, &nombre, &apellido, &edad, &peso)
		if err != nil {

			fmt.Println("Error al escanear fila:", err)
			return
		}
		fmt.Println("ID:", id, "Nombre:", nombre, "Apellido:", apellido, "edad:", edad, "peso:", peso)
	}
	if err := rows.Err(); err != nil {
		fmt.Println("Error al obtener resultados:", err)
		return
	}

	if err != nil {
		panic(err.Error())
	}
}

func editarUsuarios(db *sql.DB, id int, nombre string, apellidos string, edad int, peso float32) (int64, error) {

	query := "UPDATE usuario SET nombre = ?, apellidos = ?, edad = ?, peso = ? WHERE id = ?"

	result, err := db.Exec(query, nombre, apellidos, edad, peso, id)
	if err != nil {
		fmt.Println("Error al editar usuario: " + err.Error())
		return -1, err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return -1, err
	}
	fmt.Println("Usuario editado")
	return rowsAffected, nil
}

func eliminarUsuario(db *sql.DB, id int) (int64, error) {

	query := "DELETE FROM usuario WHERE id = ?"

	result, err := db.Exec(query, id)
	if err != nil {
		return -1, err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return -1, err
	}
	fmt.Println("Usuario eliminado")
	return rowsAffected, nil

}

func main() {

	db, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/usuario")
	if err != nil {
		panic(err.Error())
	}

	defer db.Close()

	fmt.Println("Connection Successful")

	fmt.Println("Selecciona una opcion")
	fmt.Println("1.- Usuario")
	var opcion int
	if _, err := fmt.Scan(&opcion); err != nil {
		fmt.Println("  Caracter invalido ", err)
		return
	}

	for {
		fmt.Println("Seleccionaste Usuario")
		fmt.Println("1.- Crear Usuario")
		fmt.Println("2.- Ver Usuarios")
		fmt.Println("3.- Editar Usuario")
		fmt.Println("4.- Eliminar Usuario")
		fmt.Println("5.- Salir")

		var opcion int
		if _, err := fmt.Scan(&opcion); err != nil {
			fmt.Println("Caracter inválido", err)
			return
		}

		switch opcion {
		case 1:
			
			var nombre, apellidos string
			var edad int
			var peso float32

			// Pedir al usuario que ingrese los datos
			
			fmt.Print("Ingrese el nombre: ")
			fmt.Scan(&nombre)

			fmt.Print("Ingrese los apellidos: ")
			fmt.Scan(&apellidos)

			fmt.Print("Ingrese la edad: ")
			fmt.Scan(&edad)

			fmt.Print("Ingrese el peso: ")
			fmt.Scan(&peso)
			idUsuario, err := crearUsuario(db,nombre,apellidos,edad,peso)
			if err != nil {
				fmt.Println("No se pudo crear el usuario:", err.Error())
			}
			fmt.Printf("Usuario creado con el  ID: %d .\n", idUsuario)

		case 2:
			verUsuarios(db)
		case 3:
			var id int
			var nombre, apellidos string
			var edad int
			var peso float32

			// Pedir al usuario que ingrese los datos
			fmt.Print("Ingrese el ID: ")
			fmt.Scan(&id)

			fmt.Print("Ingrese el nombre: ")
			fmt.Scan(&nombre)

			fmt.Print("Ingrese los apellidos: ")
			fmt.Scan(&apellidos)

			fmt.Print("Ingrese la edad: ")
			fmt.Scan(&edad)

			fmt.Print("Ingrese el peso: ")
			fmt.Scan(&peso)

			editarUsuarios(db, id, nombre, apellidos, edad, peso)
		case 4:
			var id int

			fmt.Print("Ingrese el ID: ")
			fmt.Scan(&id)
			eliminarUsuario(db, id)
		case 5:
			fmt.Println("Saliendo...")
			return
		default:
			fmt.Println("Opción inválida, por favor ingrese una opción válida.")
		}
	}

}
