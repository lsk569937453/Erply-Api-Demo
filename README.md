## Erply-Api-Demo
### Install Dependency
```
go mod tidy
go mod vendor
```
### Database Configuration
The project relies on MySQL database, and the configuration of the database is in the file conf/conf.ini.
```
ipAddrees=127.0.0.1
port=3308
username=root
password=elong
```
The default database is erply_ api。 Therefore, you need to create a new one named erply_ api on MySQL. database.

### Compile
Run the following command to compile the project

```
go build
```

### Start

#### Windows
Run the following command to start the project
```
./erply-api.exe

```
#### Linux/MAC
```
./erply-api

```

### Document

The Swagger Document is on http://localhost:9394/docs/index.html.

### Demo Project
The demo project is located in [Demo Project](http://localhost:9394/docs/index.html).


