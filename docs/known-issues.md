# Systems

## MongoDB

| Error message                                            | cause                     | potential solution                                                                   |
| -------------------------------------------------------- | ------------------------- | ------------------------------------------------------------------------------------ |
| `(RoleNotFound) Could not find role: roleRorApi@nhn-ror` | mongo-init.js did not run | stop r.sh and run <code>docker compose down && docker system prune -a && r.sh</code> |
