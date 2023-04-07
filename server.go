package main

import (
	"fmt"
	"net/http"
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
