path "rabbitmq/creds/ror-ms-switchboard" {
  capabilities = [ "read", "update"]
}

path "secret/data/v1.0/ror/config/ror-ms-switchboard" {
  capabilities = [ "create", "read", "update", "list" ] 
}

path "mongodb/creds/ror-ms-switchboard" {
  capabilities = [ "read", "update" ]
}

# Recommended: List all dynamic and static roles
path "mongodb/roles" {
  capabilities = [ "list" ]
}

path "mongodb/static-roles" {
  capabilities = [ "list" ]
}
