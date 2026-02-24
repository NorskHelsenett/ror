# ROR Docs

ROR Docs is an mkdocs-material ppage that generates automatically when changes are pushed to `docs/*` and `mkdocs.yaml`.

`mkdocs.yaml` holds the configuration and the mapping of files that's available under `docs/`.

To run docs locally you'll need python installed and run the following in the root of the repository:

```bash
pip install -r mkdocs-requirements.txt
mkdocs serve
```

## Mapping of files

Under `map:` in `mkdocs.yaml` it defines the header, articles and under articles based on the indentation.

Example:

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
