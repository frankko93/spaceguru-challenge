# spaceguru-challenge
Challenge BackEnd para Space Guru

## Requisitos

- Tener instalado Go
- Clonar el repositorio en su $GOPATH correspondiente
- Levantar la APP con `go run src/api/main.go`
- Ejecutar tests con `go test -v ./src/api/controllers -run Test`
- En el archivo `postmanEndpoints.js` se encuentran las pegagas a la APP


## Descripción del ejercicio

En el ejercicio de **Space Guru** se trabaja con `usuarios` y `conductores`.

### A continuación se detallan los requerimientos:
- Crear un proyecto de go usando módulos y tu librería HTTP de preferencia (en Space Guru usamos gin)
- rear una aplicación para gestionar los conductores que debe permitir:
  - Autenticar y autorizar usuarios admin
  - Autenticar conductores
  - Crear un nuevo conductor (junto con las credenciales de autenticación del mismo)
  - Obtener todos los conductores con paginación
  - Obtener todos los conductores que no estén realizando un viaje (tenés la libertad de elegir cómo identificar si ese conductor ya se encuentra realizando un viaje)
Cada ruta debe autenticarse vía un middleware que use el esquema de autenticación de preferencia (jwt, basic auth, custom token, etc)
Usar una base de datos relacional y agregar en un archivo .sql los scripts de creación de las tablas que uses (ej. mysql)
El servicio debe considerar los siguientes principios de diseño para su implementación:
Manejo de código de estado 5xx y panics
Al menos un patrón de diseño de estructura como (DDD, MVC, etc)
Principios REST (códigos de estados y verbos correcto así como convención de rutas)
Separar la capa de presentación y datos (requerimiento mínimo)
Al menos un endpoint debe tener pruebas unitarias, el resto son solo necesarias si querés demostrar tus capacidades
Al finalizar todo subí tu código a un repositorio y envíanos el link
