path "rabbitmq/creds/ror-ms-audit" {
  capabilities = [ "read", "update"]
}
path "secret/data/v1.0/ror/config/ms-audit" {
  capabilities = [ "create", "read", "update", "list" ] 
}

path "mongodb/creds/ror-ms-audit" {
  capabilities = [ "read", "update" ]
}

path "mongodb/roles" {
  capabilities = [ "list" ]
}

path "mongodb/static-roles" {
  capabilities = [ "list" ]
}
