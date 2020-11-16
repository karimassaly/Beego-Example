# Beego-Example
My RESTful API using the Beego Framework in Go

Things that should be edited: 
  - In the docker-compose file please fill these fields : - POSTGRES_USER=
            - POSTGRES_PASSWORD=
            - POSTGRES_DB=
    - In the database/database.go im using postgres as databse, please fill these fields (line 21) : 	orm.RegisterDataBase("default", "postgres", "user= password= host= port= dbname= sslmode=disable")


To compile: 

- go get -u github.com/beego/bee
- $ export PATH=$PATH:$(go env GOPATH)/bin
- bee run


By the way for the registration i did not hash the passwords
