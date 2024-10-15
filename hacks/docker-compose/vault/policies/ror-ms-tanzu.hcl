path "rabbitmq/creds/ror-ms-tanzu" {
  capabilities = [ "read", "update"]
}
path "secret/data/v1.0/ror/config/ms-tanzu" {
  capabilities = [ "create", "read", "update", "list" ] 
}

path "database/creds/redis-ror-ms-tanzu-role" {
  capabilities = [ "read", "update"]
} 