# Inicio

Para el motor de base de datos hacemos uso de MySql mediante docker.
En el root de la carpeta se puede encontrar en archivo docker-compose.yaml con el cual inicar la base de datos, este al iniciarse genera la base e inserta datos de prueba
```sh
docker-compose up
```

# Endpoints

#### Dentist

###### GET: localhost:8080/dentist/34
```
curl --location 'localhost:8080/dentist/34' 
```

###### POST : localhost:8080/dentist/
```
curl --location 'localhost:8080/dentist/' \
--header 'Content-Type: application/json' \
--data '{
	"apellido" : "maria",
	"nombre" : "jose",
	"matricula" : "3123123"
}'
```
###### PUT : localhost:8080/dentist/34
```
curl --location --request PUT 'localhost:8080/dentist/34' \
--header 'Content-Type: application/json' \
--data '{
	"apellido" : "Ernesto",
	"nombre" : "jose",
	"matricula" : "3123123"
}'
```
###### PATCH : localhost:8080/dentist/34
```
curl --location --request PATCH 'localhost:8080/dentist/34' \
--header 'Content-Type: application/json' \
--data '{
	"apellido" : "Medicus"
}'
```
###### DELETE : localhost:8080/dentist/34
```
curl --location --request DELETE 'localhost:8080/dentist/34'
```

#### Patient

###### GET: localhost:8080/patient/1
```
curl --location 'localhost:8080/patient/1'
```

###### POST : localhost:8080/patient/
```
curl --location 'localhost:8080/patient/' \
--header 'Content-Type: application/json' \
--data '{
    "apellido": "Martha",
    "nombre": "Martinez",
    "domicilio": "Calle Falsa 1234",
    "dni": "87654321",
    "fecha_alta": "2023-01-01"
}'
```
###### PUT : localhost:8080/patient/6
```
curl --location --request PUT 'localhost:8080/patient/6' \
--header 'Content-Type: application/json' \
--data '{
    "apellido": "Marta",
    "nombre": "Martinez",
    "domicilio": "Calle Falsa 1234",
    "dni": "87654321",
    "fecha_alta": "2023-01-01"
}'
```
###### PATCH : localhost:8080/patient/6
```
curl --location --request PATCH 'localhost:8080/patient/6' \
--header 'Content-Type: application/json' \
--data '{
    "apellido": "Marta",
    "nombre": "Martinez",
    "domicilio": "Calle 1234",
    "dni": "87654321",
    "fecha_alta": "2023-01-01"
}'
```
###### DELETE : localhost:8080/patient/6
```
curl --location --request DELETE 'localhost:8080/patient/6'
```
