# Releases

This page describes the release process for the perses/spec repo. It's pretty simple:

- 1. Do a commit "Release vX.Y.Z" that upgrades the version number in [package.json](./ts/package.json) + [pom.xml](./java/pom.xml)
- 2. Once the release commit is on the upstream's `main` branch, go to the repo on `GitHub UI > Releases > Draft a new release`, fill the form appropriately & submit. The CI then handles automatically the rest of the release process (e.g publishing the modules)