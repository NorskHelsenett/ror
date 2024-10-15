path "rabbitmq/creds/ror-ms-tanzu-auth" {
  capabilities = [ "read", "update"]
}
#path "secret/data/v1.0/ror/config/ms-tanzu-auth" {
#  capabilities = [ "create", "read", "update", "list" ] 
#}

path "mongodb/creds/ror-ms-tanzu" {
  capabilities = [ "read", "update" ]
}

path "mongodb/roles" {
  capabilities = [ "list" ]
}

path "mongodb/static-roles" {
  capabilities = [ "list" ]
}

path "database/creds/redis-ror-ms-tanzu-auth-role" {
  capabilities = [ "read", "update"]
} 