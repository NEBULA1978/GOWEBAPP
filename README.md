# GOWEBAPP

Como crear un servidor Web con Golang

Lo primero ha sido crear el esqueleto de nuestra aplicación con la definición del 'package', importar las dependencias(De momento solo net/http) y crear la función principal con el servidor web:

package main

import (
        "net/http"
        )

func main() {
        http.ListenAndServe(":8080", nil)
        }

Después hemos visto como compilar el proyecto o arrancarlo directamente con 'run'(Mucho mas útil cuando estamos en entorno de desarrollo):

# Compilar el proyecto
go build server.go
./server

# Arrancar el proyecto directamente
go run server.go

Servir contenido estático con Golang
Para ejecutar por consola o en navegador web:

Por consola:
curl 127.0.0.1:8080/static/hola.css

Por navegador:
127.0.0.1:8080/static/hola.css

Nos muestra el contenido de css.

Servir contenido dinámico con Golang

Para ejecutar por consola o en navegador web:

Por consola:
curl 127.0.0.1:8080/info

Por navegador:
127.0.0.1:8080/info

Lectura de parámetros GET y POST con Go

Redirecciones y errores en servidores web con Go

Uso de plantillas en Go

Certificados SSL en Golang


