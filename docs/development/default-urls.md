# URLs

| Service       | What                  | Url                                                         | ReadMe link                                                                           | Comment                                     |
| ------------- | --------------------- | ----------------------------------------------------------- | ------------------------------------------------------------------------------------- | ------------------------------------------- |
| DEX           | Authentication        | www: http://localhost:5556, grpc api: http://localhost:5557 | [dex doc](https://dexidp.io/docs/) [docker hub](https://hub.docker.com/r/bitnami/dex) | Reachable from inside and outside of docker |
| Openldap      | Mocking users         | http://localhost:389                                        |                                                                                       |                                             |
| MongoDB       | Document database     | mongodb://localhost:27017                                   |                                                                                       |                                             |
| Mongo-Express | Gui for document base | http://localhost:8081                                       |                                                                                       |                                             |
| RabbitMQ      | Message bus           | GUI: http://localhost:15672, amqp port: localhost:5672      |                                                                                       |                                             |
| Vault         | Secrets handling      | GUI: http://localhost:8200                                  |                                                                                       |                                             |
| Valkey        | Cache                 | GUI: http://localhost:6379                                  |                                                                                       |                                             |

## ROR services

| Service    | What      | Url                    | Port | ReadMe link                                                | Comment                   |
| ---------- | --------- | ---------------------- | ---- | ---------------------------------------------------------- | ------------------------- |
| ROR-Api    | Api       | http://localhost:10000 | 8080 | [ror-api](https://github.com/NorskHelsenett/ror-api)       |                           |
| ROR-WebApp | Web       | http://localhost:11000 | 8090 | [ror-webapp](https://github.com/NorskHelsenett/ror-webapp) |                           |
| ROR-Agent  | K8s Agent | http://localhost:8100  | 8100 | [ror-agent](https://github.com/NorskHelsenett/ror-agent)   | Not run by docker-compose |
