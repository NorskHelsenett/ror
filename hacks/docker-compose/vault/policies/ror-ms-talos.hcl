path "rabbitmq/creds/ror-ms-talos" {
  capabilities = [ "read", "update"]
}
path "secret/data/v1.0/ror/config/ms-talos" {
  capabilities = [ "create", "read", "update", "list" ] 
}
