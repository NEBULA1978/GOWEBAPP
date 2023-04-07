package main

import (
	"fmt"
	"net/http"
	"strconv"
)

func main() {
	// Podemos servir contenido estático con Go solamente añadiendo estas dos lí­neas de código donde creamos un servidor de ficheros pasándole la carpeta en la primera lí­nea y configurando la ruta en la segunda lí­nea:
	fs := http.FileServer(http.Dir("./static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))
	/* Definimos la ruta que llamará la función home */
	http.HandleFunc("/", home)

	// Con el segundo método hemos creado la función en el mismo lugar que creamos la ruta y usamos el paquete "fmt"(Que tenéis que agregar en la sección de los imports) para servir el contenido de una forma un poco mas elegante:
	http.HandleFunc("/info", func(w http.ResponseWriter, req *http.Request) {
		// Formateamos datos fmt
		fmt.Fprintln(w, "Host: ", req.Host)
		fmt.Fprintln(w, "URI: ", req.RequestURI)
		fmt.Fprintln(w, "Method: ", req.Method)
		fmt.Fprintln(w, "RemoteAddr: ", req.RemoteAddr)
	})

	http.HandleFunc("/producto", producto)

	// Continuamos para bingo con las redirecciones y los errores(algo tan importante para los SEO :) ):
	http.HandleFunc("/redirect", func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, "/producto", 301)
	})

	http.HandleFunc("/error", func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, "Error chungo", 501)
	})

	// Tenemos dos métodos para añadir cabeceras en nuestras aplicaciones Go: "Add" que añadirá la cabecera tantas veces como queramos y "Set" que solo pondrá una vez la cabecera:
	http.HandleFunc("/cabeceras", func(w http.ResponseWriter, req *http.Request) {
		w.Header().Add("Test", "test1")

		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		fmt.Fprintln(w, "{ \"hola\":1 }")
	})

	http.ListenAndServe(":8080", nil)
}

/* Definimos la función home */
// Hemos usado dos métodos para servir contenido dinámico, el primero, definir una ruta y crear una función a parte que básicamente devuelve el código HTML desde un string:
func home(w http.ResponseWriter, r *http.Request) {
	html := "<html>"
	html += "<body>"
	html += "<h1>Hola Mundo</h1>"
	html += "</body>"
	html += "</html>"
	// Le pasamos a w el string previamente convertido a array de bytes
	w.Write([]byte(html))
}

// Lectura de parámetros GET y POST con Go

// Para la lectura de parámetros GET y POST nos hemos montado un pequeño mantenimiento de productos en una función, aquí­ tenéis el código:
var productos []string

func producto(w http.ResponseWriter, r *http.Request) {
	// Parseamos formulario
	r.ParseForm()
	// Pedimos parametro add
	add, okForm := r.Form["add"]

	if okForm && len(add) == 1 {
		productos = append(productos, string(add[0]))
		w.Write([]byte("Producto añadido correctamente"))

		return
	}
	// Nos devuelve dos variables,el propio parametro y si ha podido recojerlo de la ruta
	prod, ok := r.URL.Query()["prod"]

	if ok && len(prod) == 1 {
		// convertimos el prod1 a asci a string
		pos, err := strconv.Atoi(prod[0])

		if err != nil {
			return
		}

		html := "<html>"
		html += "<body>"
		html += "<h1>Producto " + productos[pos] + "</h1>"
		html += "</body>"
		html += "</html>"

		w.Write([]byte(html))

		return
	}

	html := "<html>"
	html += "<body>"
	html += "<h1>Total Productos " + strconv.Itoa(len(productos)) + "</h1>"
	html += "</body>"
	html += "</html>"
}
