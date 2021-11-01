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
- Cada ruta debe autenticarse vía un middleware que use el esquema de autenticación de preferencia (jwt, basic auth, custom token, etc)
- Usar una base de datos relacional y agregar en un archivo .sql los scripts de creación de las tablas que uses (ej. mysql)
- El servicio debe considerar los siguientes principios de diseño para su implementación:
  - Manejo de código de estado 5xx y panics
  - Al menos un patrón de diseño de estructura como (DDD, MVC, etc)
  - Principios REST (códigos de estados y verbos correcto así como convención de rutas)
  - Separar la capa de presentación y datos (requerimiento mínimo)
- Al menos un endpoint debe tener pruebas unitarias, el resto son solo necesarias si querés demostrar tus capacidades
- Al finalizar todo subí tu código a un repositorio y envíanos el link


#### URLs Usuarios

1. **Crear un nuevo usuario.**

Por ejemplo:

_Request_:
```
POST /users
```
_Body_:

```json
{
	"email": "spaceguru@test.com",
	"name": "franco",
	"surname": "aballay",
  "type": "driver",
  "password": "1234"
}
```

_Response_:
```json
{
    "id": "e8adcf81-71d3-4720-ab54-f4354e408bd4",
    "name": "franco",
    "surname": "aballay",
    "type": "driver",
    "email": "spaceguru@test.com",
    "password": "1234"
}
```

2. **Autenticar un usuario.**

La respuesta se usa para `x-auth` en los headers en cada ruta que necesite autenticación.

Por ejemplo:

_Request_:
```
POST /users/login
```
_Body_:

```json
{
	"email": "spaceguru@test.com",
	"password": "1234"
}
```

_Response_:
```json
{
    "token": "xxxxxx.yyyyyyy.zzzzzz"
}
```


3. **[adicional] Crear una nueva ruta.**

La respuesta se usa para `x-auth` en los headers en cada ruta que necesite autenticación.

Donde:

- `status`: Es el filtro por el estado de la propiedad; es un parámetro opcional y su valor default es `ALL`. Puede tener los siguientes valores:
	- `in_progress`: Significa que el viaje está activo.
	- `finished`: Significa que el viaje está finalizado.
	- `cancelled`: Significa que el viaje está cancelado.

Por ejemplo:

_Request_:

```
HEADER x-auth: xxxxx.yyyy.zzzz [token de sesión]
```
```
POST /travel```
_Body_:

```json
{
	"user_id": "e8adcf81-71d3-4720-ab54-f4354e408bd4", //id del usuario creado
	"vehicle_id": "1",
	"status": "in_progress",
  "route": "aaaaaaaa"
}
```

_Response_:
```json
{
    "id": 1,
    "user_id": "f4597fc7-bf8a-4b80-93e8-32cd8ad857bc",
    "vehicle_id": "1",
    "status": "in_progress",
    "route": "aaaaaaaa"
}
```

4. **Buscar conductores**

Los resultados vienen paginados, y se representan de la siguiente manera.

_Request_:
  
```
GET /users/drivers?status={status}&page={pageNumber}&pageSize={pageSize}
```

Donde:

- `status`: Es el filtro por el estado de la propiedad; es un parámetro opcional y su valor default es `ALL`. Puede tener los siguientes valores:
	- `free`: Retorna todos los conductores que no tienen un viaje en curso.
  - `in_travel`: Retorna todos los conductores que tienen un viaje en curso.
- `page`: Es el número de la página; es un parámetro opcional y su valor default es 1
- `pageSize`: Es el tamaño solicitado de resultados en la página. Es un parámetro opcional, su valor default es 10, y su valor máximo es 20.

_Response_:

La respuesta debe seguir la siguiente estructura de campos:

- `page`: La página actual de los resultados
- `pageSize`: El máximo número de resultados retornado por página
- `total`: El número total de propiedades
- `totalPages`: El número total de páginas que contienen resultados para la búsqueda hecha.
- `data`: Un array con los objetos conteniendo las propiedades solicitadas en el request. Para ver en detalle la estructura de una `propiedad`, por favor avance a la siguiente sección.

```json
{
	"page": 1,
	"pageSize": 10,
	"totalPages": 1,
	"total": 1,
	"data": [
		{
      "id": "ce6ac213-21c8-410e-a299-ad7967fb8034",
      "name": "franco",
      "surname": "aballay",
      "type": "driver",
      "email": "spaceguru2@test.com",
      "password": "1234",
      "updatedAt": "0001-01-01T00:00:00Z",
      "createdAt": "0001-01-01T00:00:00Z"
    }
	]
}
```
