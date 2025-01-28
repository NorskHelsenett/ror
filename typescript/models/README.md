# ROR NPM

Npm package for ROR (Release Operate Report) project: https://github.com/norskHelsenett/ror

## Prerequisite

Npm publish token in environment variables

## Build

```bash
npm i
```

## Test locally

- Update version number in `package.json`
- Run

```bash
npm pack
```

A new file will appear beside ``package.json`. named something like: rork8s-ror-resources-<version number>.tgz

To use in another `package.json`
Replace dependency url, while testing to example this:

```json
"@rork8s/ror-resources": "file:../<ref to repo path>/typescript/models/rork8s-ror-resources-0.0.0.tgz",
```

## Publish

- Update version number in `package.json`
- Remember to add npm token before publishing
- Run

```bash
npm publish
```
