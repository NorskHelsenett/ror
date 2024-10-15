path "rabbitmq/creds/ror-ms-slack" {
  capabilities = [ "read", "update"]
}
path "secret/data/v1.0/ror/config/ms-slack/*" {
  capabilities = [ "create", "read", "update", "list", "list" ] 
}
