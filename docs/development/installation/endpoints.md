# Endpoints

| Service       | What                  | Url                                                         | ReadMe link                                                                           | Comment                                     |
| ------------- | --------------------- | ----------------------------------------------------------- | ------------------------------------------------------------------------------------- | ------------------------------------------- |
| DEX           | Authentication        | www: http://localhost:5556, grpc api: http://localhost:5557 | [dex doc](https://dexidp.io/docs/) [docker hub](https://hub.docker.com/r/bitnami/dex) | Reachable from inside and outside of docker |
| Openldap      | Mocking users         | http://localhost:389                                        |                                                                                       |                                             |
| MongoDB       | Document database     | mongodb://localhost:27017                                   |                                                                                       |                                             |
| Mongo-Express | Gui for document base | http://localhost:8081                                       |                                                                                       |                                             |
| RabbitMQ      | Message bus           | GUI: http://localhost:15672, amqp port: localhost:5672      |                                                                                       |                                             |
| Vault         | Secrets handling      | GUI: http://localhost:8200                                  |                                                                                       |                                             |
| Valkey        | Cache                 | GUI: http://localhost:6379                                  |                                                                                       |                                             |
