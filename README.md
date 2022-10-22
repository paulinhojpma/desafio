
# Producers Data

Producers Data is an application that processes transaction data from sellers/affiliates, persists in a related database and presents the data in a list.

  ## Start application

After having cloned the project on your machine, **Docker** and **Docker Compose** must be installed, after that run the command in the project root:

```bash
$ docker compose up
```
Wait to create and run all containers and you can access the application, with the URL.
```bash
localhost:3000/transactions
```

# Basic Operation

After acess the URL, you will se an image like that:

![Tela inicial](https://drive.google.com/uc?export=view&id=1CTZVXMrCg6bSFsyR51BBaCojkO_PjBfX)
After put a valid file on the form, hit the submit button and the screen below will show up:

![Tela Tabela](https://drive.google.com/uc?export=view&id=1B4g_nwwokqssWS69w7kNrwhZdwst2yHR)





# API Routes 

  This section will show the backend api.

## Create Transactions

### Request  

```POST localhost:8890/transactions
Request Headers: Content-Type multipart/form-data ```

  Request Body: A txt file, example: 
12022-01-15T19:20:30-03:00CURSO DE BEM-ESTAR            0000012750JOSE CARLOS
12021-12-03T11:46:02-03:00DOMINANDO INVESTIMENTOS       0000050000MARIA CANDIDA
22022-01-16T14:13:54-03:00CURSO DE BEM-ESTAR            0000012750THIAGO OLIVEIRA
32022-01-16T14:13:54-03:00CURSO DE BEM-ESTAR            0000004500THIAGO OLIVEIRA
42022-01-16T14:13:54-03:00CURSO DE BEM-ESTAR            0000004500JOSE CARLOS
12022-01-22T08:59:13-03:00DOMINANDO INVESTIMENTOS       0000050000MARIA CANDIDA


```
  

### Response
  
#### Sucess
HTTP/1.1 201 CREATED

 ```json 
{"codResponse":201,"message":"Transactions Created"}
```

#### Fail
HTTP/1.1 400 BAD REQUEST

 ```json 
{"resource":"Error on unpack file","code":"400","message":"Invalid format file","idOperation":""}
```
  
 



## Retrieve transactions data

### Request  

```GET localhost:8890/transactions```  

### Response
  
#### Sucess
HTTP/1.1 200 OK

 ```json 
{"codResponse":200,"message":"Transações retornadas com sucesso","producers":[{"ID":36,"Name":"JOSE CARLOS","Transactions":[{"ID":101,"CreatedAt":"2022-10-22T00:00:00Z","UpdatedAt":"2022-10-22T00:00:00Z","DeletedAt":null,"TypeID":1,"Type":{"Type":1,"Description":"Venda produtor","Nature":"Entrada"},"Date":"2022-01-15T00:00:00Z","Product":"CURSO DE BEM-ESTAR","Value":127.5,"ProducerID":36},{"ID":102,"CreatedAt":"2022-10-22T00:00:00Z","UpdatedAt":"2022-10-22T00:00:00Z","DeletedAt":null,"TypeID":4,"Type":{"Type":4,"Description":"Comissão recebida","Nature":"Entrada"},"Date":"2022-01-16T00:00:00Z","Product":"CURSO DE BEM-ESTAR","Value":45,"ProducerID":36},{"ID":103,"CreatedAt":"2022-10-22T00:00:00Z","UpdatedAt":"2022-10-22T00:00:00Z","DeletedAt":null,"TypeID":1,"Type":{"Type":1,"Description":"Venda produtor","Nature":"Entrada"},"Date":"2022-03-01T00:0 ...
```

#### Fail
HTTP/1.1 500 INTERNAL SERVER ERROR

 ```json 
{"resource":"There is no transactions","code":"400","message":"Error on retrieve transactions","idOperation":""}
```
