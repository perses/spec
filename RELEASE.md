# Releases

This page describes the release process for the perses/spec repo. It's pretty simple:

1. Checkout to a new branch named `release-vX.Y.Z` from main.
2. Upgrade the version number in [package.json](./ts/package.json) + [pom.xml](./java/pom.xml)
3. Run `cd ts && npm install`.
4. Commit these changes as a standalone commit ("Prepare release vX.Y.Z") and create a PR.
5. Once the release commit is on the upstream's `main` branch, go to the repo on `GitHub UI > Releases > Draft a new release`, fill the form appropriately & submit. The CI then handles automatically the rest of the release process (e.g publishing the modules)
