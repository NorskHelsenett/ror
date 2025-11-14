# Systems

## MongoDB

| Error message                                                                 | cause                     | potential solution                                                                   |
| ----------------------------------------------------------------------------- | ------------------------- | ------------------------------------------------------------------------------------ |
| `(RoleNotFound) Could not find role: roleRorApi@nhn-ror`                      | mongo-init.js did not run | stop r.sh and run <code>docker compose down && docker system prune -a && r.sh</code> |
| `/usr/local/bin/docker-entrypoint.sh: ignoring /docker-entrypoint-initdb.d/*` | mongo-init.js did not run | check the permissions on hacks/data/docker-compose/mongodb/entrypoint is sufficient  |
