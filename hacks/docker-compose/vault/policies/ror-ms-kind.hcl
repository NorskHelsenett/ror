path "rabbitmq/creds/ror-ms-kind" {
  capabilities = [ "read", "update"]
}
path "secret/data/v1.0/ror/config/ms-kind" {
  capabilities = [ "create", "read", "update", "list" ] 
}
