### Description

gRPC + Gateway application for storing restaurant orders, allows to view, add, update and delete orders  
server runs on port 8282  
gateway runs on port 8383  
database runs on port 5432  
Use this address to run the application: `localhost:8383/{path}`  

### Table with paths  

|                Path                   |    Method     |                   Description               |           Body example       |
| --------------------------------------|---------------| --------------------------------------------| ---------------------------- |
| `api/orders/{id}`                     |    `DELETE`   |   delete order                              |                              |
| `api/orders`                             |    `POST`     |   create new order                          | ```"order": {"id": "1","Date": "2021-04-02 00:00:00 +0000 UTC","TableNumber": "1","WaiterId": "1","Price": "300","Payment": true}```|
| `api/orders/{id}`                        |    `PUT`      |   update order                              |```"order": {"id": "1","Date": "2021-04-02 00:00:00 +0000 UTC","TableNumber": "1","WaiterId": "1","Price": "300","Payment": true}```|
| `api/orders/{id}`                        |    `GET`      |   get order by `id`                         |                               |
| `api/orders`                             |    `GET`      |   get all orders                            |                               |

### Generate go code from protobuf
To generate go code from brotobuf use command: `protoc -I . user.proto --grpc-gateway_out . --go_out=plugins=grpc:.`
### Run application via docker-compose
To run application via docker-compose use command: `docker-compose -f docker-compose.yaml up`
