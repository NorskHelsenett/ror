path "rabbitmq/creds/ror-ms-auth" {
  capabilities = [ "read", "update"]
}
path "secret/data/v1.0/ror/dex/*" {
  capabilities = [ "create", "read", "update", "list" ] 
}
path "secret/data/v1.0/ror/clusters/+/dex/*" {
  capabilities = [ "create", "read", "update", "list" ] 
}
