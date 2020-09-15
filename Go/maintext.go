package main

import(
        _ "github.com/go-sql-driver/mysql"      // La librería que nos permite conectar a MySQL
        "database/sql"                          // Interactuar con bases de datos
        "fmt"                                   // Imprimir mensajes y esas cosas
        "log"
)
func obtenerBaseDeDatos() (db *sql.DB, e error) {
        usuario := "docker"
        pass := "docker"
        host := "tcp(db:3306)" //db=127.0.0.1
        nombreBaseDeDatos := "test_db"
        // Debe tener la forma usuario:contraseña@host/nombreBaseDeDatos
        db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@%s/%s", usuario, pass, host, nombreBaseDeDatos))
        if err != nil {
                return nil, err
        }
        return db, nil
}

type Contacto struct {
        Nombre, Direccion, CorreoElectronico string
        Id                                   int
}

func insertar(c Contacto) (e error) {
        db, err := obtenerBaseDeDatos()
        if err != nil {
                return err
        }
       // defer db.Close()

        // Preparamos para prevenir inyecciones SQL
        sentenciaPreparada, err := db.Prepare("INSERT INTO agenda (nombre, direccion, correo_electronico) VALUES(?, ?, ?)")
        if err != nil {
                return err
        }
        //defer sentenciaPreparada.Close()
        // Ejecutar sentencia, un valor por cada '?'
        _, err = sentenciaPreparada.Exec(c.Nombre, c.Direccion, c.CorreoElectronico)
        if err != nil {
                return err
        }
        return nil
}


func main() {
        cnn, err := obtenerBaseDeDatos() //sql.Open("mysql", "docker:docker@tcp(db:3306)/test_db")
        if err != nil {
                log.Fatal(err)
        }
        //insertar
        //var c Contacto
        c := Contacto{
                Nombre:            "Angel Sullon ",
                Direccion:         "Calle Sin Nombre #123",
                CorreoElectronico: "c@gmail.me",
        }
        insertar(c)      
        //
        // query all data
        rows, e := cnn.Query("select * from agenda")
        //ErrorCheck(e)
        if e != nil {
                log.Fatal(e)
        }
         
        contactos := []Contacto{}
        // declare empty post variable
        
         
        // iterate over rows
        for rows.Next() {
            c2 := Contacto{}
            rows.Scan(&c2.Id, &c2.Nombre, &c2.Direccion, &c2.CorreoElectronico)
            //ErrorCheck(e)
            contactos = append(contactos, c2)
            //fmt.Println(c2)
        }
        for _, contacto := range contactos {
                fmt.Printf("%v\n", contacto.Nombre)
        }



        id := 1
        var name string

        if err := cnn.QueryRow("SELECT name FROM test_tb WHERE id = ? LIMIT 1", id).Scan(&name); err != nil {
                log.Fatal(err)
        }

        fmt.Println(id, "hi ",name)
}


