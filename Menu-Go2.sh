#!/bin/bash

echo "Bienvenido al script interactivo para el servidor web en Golang."

while true; do
    echo "Seleccione una opción:"
    echo "1. Compilar y ejecutar el servidor web"
    echo "2. Añadir producto"
    echo "3. Ver productos"
    echo "4. Redireccionar"
    echo "5. Mostrar error"
    echo "6. Salir"

    read -p "Opción: " opcion

    case $opcion in
        1)
            echo "Compilando y ejecutando el servidor web..."
            go build server.go
            ./server
            ;;
        2)
            read -p "Nombre del producto: " producto
            echo "Añadiendo producto..."
            curl -X POST -d "add=$producto" 127.0.0.1:8080/producto
            ;;
        3)
            read -p "Índice del producto: " indice
            echo "Mostrando producto..."
            curl 127.0.0.1:8080/producto?prod=$indice
            ;;
        4)
            echo "Redireccionando..."
            curl 127.0.0.1:8080/redirect
            ;;
        5)
            echo "Mostrando error..."
            curl 127.0.0.1:8080/error
            ;;
        6)
            echo "Saliendo del script. ¡Hasta luego!"
            exit 0
            ;;
        *)
            echo "Opción no válida. Inténtelo de nuevo."
            ;;
    esac
done
