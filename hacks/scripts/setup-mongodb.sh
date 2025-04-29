vault secrets enable -path=database database

echo "# Setting secrets for mongodb"

echo "Creating a role for ror-api"
vault write database/roles/mongodb-ror-api \
    db_name=mongodb \
    creation_statements='{ "db": "nhn-ror", "roles": [{ "role": "roleRorApi"}] }' \
    default_ttl="1h" \
    max_ttl="24h"

vault write database/roles/mongodb-ror-ms-switchboard \
    db_name=mongodb \
    creation_statements='{ "db": "nhn-ror", "roles": [{ "role": "roleRorApi"}] }' \
    default_ttl="1h" \
    max_ttl="24h"

echo "Creating a role for ror-ms-audit"
vault write database/roles/mongodb-ror-ms-audit \
    db_name=mongodb \
    creation_statements='{ "db": "nhn-ror", "roles": [{ "role": "roleRorApi"}] }' \
    default_ttl="1h" \
    max_ttl="24h"

echo "Setting up connection to MongoDB"
echo "Creating a role for ror-ms-nhn"
vault write database/roles/mongodb-ror-ms-nhn \
    db_name=mongodb \
    creation_statements='{ "db": "nhn-ror", "roles": [{ "role": "roleRorApi"}] }' \
    default_ttl="1h" \
    max_ttl="24h"

echo "Creating a role for ror-ms-tanzu"
vault write database/roles/mongodb-ror-ms-tanzu \
    db_name=mongodb \
    creation_statements='{ "db": "nhn-ror", "roles": [{ "role": "roleRorApi"}] }' \
    default_ttl="1h" \
    max_ttl="24h"

echo "Creating a role for ror-ms-tanzu-auth"
vault write database/roles/mongodb-ror-ms-tanzu-auth \
    db_name=mongodb \
    creation_statements='{ "db": "nhn-ror", "roles": [{ "role": "roleRorApi"}] }' \
    default_ttl="1h" \
    max_ttl="24h"

echo "Creating a role for superadmin"
vault write database/roles/mongodb-superadmin \
    db_name=mongodb \
    creation_statements='{ "db": "admin", "roles": [{"role": "root"}] }' \
    default_ttl="1h" \
    max_ttl="24h"

vault write database/config/mongodb \
      plugin_name=mongodb-database-plugin \
      allowed_roles="mongodb-ror-api,mongodb-ror-ms-switchboard,mongodb-ror-ms-audit,mongodb-ror-ms-nhn,mongodb-ror-ms-tanzu,mongodb-ror-ms-tanzu-auth,mongodb-superadmin" \
      connection_url="mongodb://{{username}}:{{password}}@{{host}}:{{port}}/" \
      username="root" \
      password="" \
      host="ror-sharded-mongodb-mongodb-sharded.ror-mongodb.svc.managedcluster.local" \
      port="27017" \



# clusters
db.clusters.createIndex({_id: 1  }, {name: '_id_' })
db.clusters.createIndex({clusterid: 1  }, {name: 'clusterid_1' })
db.clusters.createIndex({workspace: 1  }, {name: 'workspace_1' })
db.clusters.createIndex({datacenter: 1  }, {name: 'datacenter_1' })
db.clusters.createIndex({'state.resources.kind': 1, 'state.resources.apiversion': 1  }, {name: 'state.resources.kind_1_state.resources.apiversion_1' })
db.clusters.createIndex({'state.resources.metadata.name': 1, 'state.resources.metadata.namespace': 1  }, {name: 'state.resources.metadata.name_1_state.resources.metadata.namespace_1' })

# apikeys
db.apikeys.createIndex({hash: 1 }, {name: 'hash_1', unique: true })
db.apikeys.createIndex({created: 1 }, {name: 'created_1' })
db.apikeys.createIndex({identifier: 1 }, {name: 'identifier_1' })
db.apikeys.createIndex({identifier: 1, type: 1}, {name: 'identifier_1_type_1' })


# acl

db.acl.createIndex({scope: 1, subject: 1})
db.acl.createIndex({group: 1})
db.acl.createIndex({scope: 1, subject: 1, group: 1})


# resources
db.resources.createIndex({ _id: 1 }, {name: '_id_' })
db.resources.createIndex({ clusterid: 1, kind: 1, apiversion: 1} ,{name: 'clusterid_1_kind_1_apiversion_1'})
db.resources.createIndex({ uid: 1 }, {name: 'uid_1', unique: true })
db.resources.createIndex({apiversion: 1,kind: 1,internal: 1,'owner.scope': 1,'owner.subject': 1 },{name: 'apiversion_1_kind_1_internal_1_owner.scope_1_owner.subject_1'})
db.resources.createIndex({kind: 1, apiversion: 1 },{name: 'kind_1_apiversion_1'})

# resourcesv2
db.resourcesv2.createIndex({ _id: 1 }, {name: '_id_' })
db.resourcesv2.createIndex({ uid: 1 }, {name: 'uid_1', unique: true })
db.resourcesv2.createIndex({ 'typemeta.kind': 1 , 'typemeta.apiversion': 1},{name: 'typemeta.kind_1_typemeta.apiversion_1'})
db.resourcesv2.createIndex({ 'metadata.uid': 1 }, {name: 'metadata.uid_1'})
db.resourcesv2.createIndex({ 'rormeta.ownerref.scope': 1, 'rormeta.ownerref.subject': 1 }, {name: 'rormeta.ownerref.scope_1_rormeta.ownerref.subject_1'})


