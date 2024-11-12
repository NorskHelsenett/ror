path "rabbitmq/creds/ror-tanzu-agent" {
  capabilities = [ "read", "update"]
}
path "secret/data/v1.0/ror/tanzu/agent" {
  capabilities = [ "create", "read", "update", "list" ] 
}

path "mongodb/creds/ror-tanzu-agent" {
  capabilities = [ "read", "update" ]
}

path "mongodb/roles" {
  capabilities = [ "list" ]
}

path "mongodb/static-roles" {
  capabilities = [ "list" ]
}

path "database/creds/redis-ror-tanzu-agent-role" {
  capabilities = [ "read", "update"]
} 