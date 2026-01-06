# ror-docs

ror-docs en en mkdocs-material side som genereres automatisk ved push til `docs/**/*` og `mkdocs.yaml`. mkdocs.yaml inneholder config og mapping av filer under /docs.

## Mapping av filer

under `map:` i `mkdocs.yaml`defineres mapping av filer, første nivå er headere, undernivå er trestruktur i sidemenyen.

eks:

```yaml
nav:
    - ROR:
          - index.md
          - components.md
          - design.md
          - getting-started.md
    - Clients:
          - ror-admin:
                - ror-admin/index.md
          - ror-cli:
                - ror-cli/index.md
                - ror-cli/auth-flow.md
```

## Kodenær dokumentasjon

Hvis dokumentasjonen er lagt sammen med koden kan den automatisk integreres med ror-docs ved å legg til en kopieringskommando i filen `/cmd/docs/collectdocs.sh`. Husk å mappe filen i `mkdocs.yaml`

eks:

```bash
#!/bin/bash

#API
cp cmd/api/ReadMe.md docs/ror-api/index.md
```
