path "rabbitmq/creds/ror-ms-nhn" {
  capabilities = [ "read", "update"]
}

path "mongodb/creds/ror-ms-nhn" {
  capabilities = [ "read", "update" ]
}

path "mongodb/roles" {
  capabilities = [ "list" ]
}

path "mongodb/static-roles" {
  capabilities = [ "list" ]
}