# REST API en Go con PostgreSQL, GORM y Gorilla/mux 🚀

Este proyecto es un backend desarrollado en Go que utiliza PostgreSQL como base de datos. Utiliza el ORM GORM para interactuar con la base de datos y Gorilla/mux para manejar las rutas de la API REST. Este proyecto nos permite realizar operaciones CRUD (Crear, Listar y Eliminar) en nuestras entidades.

## Dependencias 📚

Este proyecto utiliza varios paquetes de Go:

1. **GORM:** Un fantástico ORM (Object-Relational Mapper) para Go. Facilita la interacción con nuestra base de datos PostgreSQL.
2. **Gorilla/mux:** Un potente paquete para manejar rutas y solicitudes HTTP en nuestras aplicaciones Go.

## Base de datos 🎲

Para nuestra base de datos, usamos PostgreSQL ejecutándose en un contenedor Docker. Aquí está el comando que usamos para ejecutar la base de datos:

```bash
ddocker run --name some-postgres -e POSTGRES_PASSWORD=admin -e POSTGRES_USER=root -p 5432:5432 -d postgres
```
## Código 🖥️

El archivo `main.go` es el punto de entrada de nuestra aplicación. Aquí establecemos la conexión con la base de datos, definimos nuestras rutas y comenzamos el servidor.

La estructura del proyecto incluye las siguientes carpetas y archivos principales:

- **routes:** Contiene los archivos `index.routes.go`, `users.go` y `tasks.go`, que definen las rutas y manejadores para nuestras operaciones API.
- **models:** Contiene las definiciones de los modelos de datos que estamos utilizando.
- **db:** Contiene el código necesario para conectarse a nuestra base de datos PostgreSQL.

## Cómo iniciar el servidor 🚀

Para iniciar el servidor, simplemente ejecuta el comando `go run main.go` desde la raíz del proyecto. Esto inicia el servidor en el puerto 3000. Asegúrate de que tu base de datos PostgreSQL esté en ejecución antes de iniciar el servidor.

## Usando la API 🌐

Una vez que el servidor está en marcha, puedes interactuar con la API a través de varias rutas. Puedes usar herramientas como Postman o curl para realizar peticiones a la API. Aquí hay algunas de las rutas y métodos HTTP que nuestra API admite, y un ejemplo de cómo usar curl para hacer una petición:

- Obtener todos los usuarios: GET `/users`
- Obtener un usuario específico: GET `/users/{id}`
- Crear un nuevo usuario: POST `/users`

    Ejemplo de petición curl: 
    ```bash
        curl -X POST -H "Content-Type: application/json" -d '{"first_name":"Santiago","last_name": "Vazquez", "email":"john@example.com"}' http://localhost:3000/users
    ```
- Eliminar un usuario: DELETE `/users/{id}`
- Obtener todas las tareas: GET `/tasks`
- Crear una nueva tarea: POST `/tasks`
- Obtener una tarea específica: GET `/tasks/{id}`
- Eliminar una tarea: DELETE `/tasks/{id}`

## Cómo contribuir 🤝

Las contribuciones son siempre bienvenidas. Si encuentras algún problema o tienes alguna idea para mejorar el proyecto, no dudes en abrir un 'issue' o enviar un 'pull request'.

Gracias por visitar y probar nuestro proyecto! 😄🙏🚀

