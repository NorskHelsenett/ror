# Required: Get credentials from the database secrets engine for 'ror-api' role.
path "mongodb/creds/ror-api" {
  capabilities = [ "read", "update"]
}
path "rabbitmq/creds/ror-api" {
  capabilities = [ "read", "update"]
}
path "database/creds/redis-ror-api-role" {
  capabilities = [ "read", "update"]
} 
path "secret/data/v1.0/ror/clusters/*" {
  capabilities = [ "create","read", "update", "list" ] 
}
path "secret/data/v1.0/ror/*" {
  capabilities = [ "read", "list" ] 
}
