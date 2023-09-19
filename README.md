# grupo-siete-go
trabajo final de la materia go.

# API de Reserva de Turnos Odontológicos

Esta API ha sido desarrollada para gestionar la reserva de turnos en una clínica odontológica. Permite a los usuarios administrar información sobre odontólogos, pacientes y turnos, así como garantizar la seguridad en las operaciones sensibles.

## Guia para correr en local
- Para levantar la base de datos SQL, ejecutar Docker desde la raiz del proyecto con el comando : 
```
docker compose up -d
```
- Para levantar la api, desde la terminal pararse en raiz/cmd y ejecutar el comando: 
```
go run main.go
```
- Si se quiere levantar la aplicacion desde el IDE con el Run, agregar en la configuracion del Run el archivo .env
- Para leer la documentacion de los endpoints, ir a la URL: http://localhost:8080/docs/index.html

#### En caso de hacer cambios en los endpoints, es necesario actualizar la documentacion de Swagger
- Eliminar la carpeta raiz/docs
- Correr en terminal 
```
go install github.com/swaggo/swag/cmd/swag@latest
```
- Y para generar la nueva documentacion: 
```
swag init -g cmd/main.go
```

## Funcionalidades Principales

### Administración de Odontólogos (Dentistas)

- **POST /dentistas:** Agregar Dentista
- **GET /dentistas/{id}:** Obtener Dentista por ID
- **PUT /dentistas/{id}:** Actualizar Dentista
- **PATCH /dentistas/{id}:** Actualizar Dentista (parcial)
- **DELETE /dentistas/{id}:** Eliminar Dentista

### Administración de Pacientes

- **POST /pacientes:** Agregar Paciente
- **GET /pacientes/{id}:** Obtener Paciente por ID
- **PUT/PATCH /pacientes/{id}:** Actualizar Paciente
- **DELETE /pacientes/{id}:** Eliminar Paciente

### Gestión de Turnos

- **POST /turnos:** Agregar Turno
- **GET /turnos/{id}:** Obtener Turno por ID
- **PUT /turnos/{id}:** Actualizar Turno
- **PATCH /turnos/{id}:** Actualizar Turno (parcial)
- **DELETE /turnos/{id}:** Eliminar Turno
- **POST /turnos/paciente/{dni}/dentista/{matricula}:** Agregar Turno por DNI y Matrícula
- **GET /turnos/paciente/{dni}:** Obtener Turnos por DNI del Paciente

### Seguridad

- La API implementa seguridad mediante middleware para operaciones sensibles (POST, PUT, PATCH y DELETE), lo que requiere autenticación.

### Documentación

- La documentación de la API está disponible a través de Swagger, lo que facilita su comprensión y uso.

## Consideraciones Técnicas

La aplicación sigue una estructura de diseño orientado a paquetes con las siguientes capas:

1. **Capa de Entidades de Negocio:** Contiene las clases que representan las entidades principales, como Dentista, Paciente y Turno.

2. **Capa de Acceso a Datos (Repository):** Gestiona la interacción con la base de datos. Se puede utilizar una base de datos relacional (por ejemplo, H2 o MySQL) o no relacional (por ejemplo, MongoDB).

3. **Capa de Acceso a Datos (base de datos):** Donde reside la base de datos del sistema.

4. **Capa de Servicio:** Contiene la lógica de negocio y coordina las operaciones entre la capa de acceso a datos y la capa de controladores.

5. **Capa de Controladores (Handlers):** Aquí se implementan los controladores que manejan las solicitudes HTTP y gestionan las respuestas.

Este proyecto te ofrece una solución completa para la gestión de turnos en una clínica odontológica. ¡Aprovecha esta API para mejorar la eficiencia de tu clínica y brindar un mejor servicio a tus pacientes!