# AI contributor guide

## Purpose and instruction scope

This repository is the public Perses resource specification. It represents dashboard, datasource, plugin, and related
contracts in Go, CUE, TypeScript, and Java. A model change is complete only when all affected language representations,
runtime validators, generated schemas, documentation, and compatibility expectations agree.

Before editing:

- Read `README.md` for repository purpose, packages, and supported language representations.
- Read the affected language package and neighboring versions of the same contract.
- Inspect generation notices and the root `Makefile` before changing a generated counterpart.

## Architecture map

- `go/`: Go resource types, validation behavior, and tests.
- `cue/`: CUE schemas and validation. Files ending in `_go_gen.cue` are generated from Go; companion `_patch.cue` files
  hold CUE-specific constraints that generation cannot express.
- `ts/`: strict TypeScript contracts and Zod runtime schemas published for UI and plugin consumers.
- `java/`: Java models and validation artifacts built with Maven.

The main `perses` repository consumes these contracts in its API and application. `perses/shared` and `perses/plugins`
consume the TypeScript and schema layers. Do not add product-specific behavior or UI components here.

## Engineering rules

- Treat field names, optionality, defaults, enum values, discriminators, validation, and serialization as public API.
- Preserve backward compatibility for stored dashboards and external clients unless a breaking change is explicit,
  documented, and coordinated across consumers.
- Assess every contract change across Go, generated and patched CUE, TypeScript types, Zod schemas, and Java. Update
  only the relevant representations, but never allow them to diverge silently.
- In TypeScript, keep interfaces and runtime schemas aligned, preserve strict typing, and export public symbols through
  the intended package barrel. Runtime input must be validated rather than trusted through a cast.
- Do not hand-edit `_go_gen.cue` or files with a generated notice. Change the Go source, run the generator, then inspect
  the generated diff. Put non-generated CUE constraints in the matching patch file.
- Add focused positive and negative validation tests for new fields, constraints, defaults, and compatibility behavior.
- New source files need the repository's Apache license header.
- Do not add dependencies, change published versions, raise lint ceilings, or add broad suppressions without an explicit
  task-specific reason.

## Validation

Use the same Go version as in `go.mod`, Node.js from `ts/.nvmrc`, pnpm 12 from `ts/package.json`, and Java version from `pom.xml` where applicable.

For Go and CUE changes, run the relevant checks from the repository root:

```sh
make checkformat
make checkunused
make checkstyle
make go-test
make cue-eval
make cue-test
```

Run `make cue-gen` after changing Go types that produce CUE, and verify that only expected generated files changed.

For TypeScript changes, run from `ts/`:

```sh
pnpm install --frozen-lockfile
pnpm lint
pnpm format:check
pnpm type-check
pnpm build
```

For Java changes, run the relevant Maven tests from `java/`, normally `mvn test` or `mvn install`. If a language
toolchain is unavailable, still validate the other representations and report the missing check clearly.

## Completion checklist

- All affected language representations and runtime validators express the same contract.
- Stored-resource and client compatibility, including omitted fields and defaults, has been considered.
- Generated files came from their source and contain only expected changes.
- New validation behavior has focused positive, negative, and round-trip coverage where applicable.
- Relevant Go, CUE, TypeScript, Java, format, lint, type, and build checks pass.
- The final diff contains no credentials, warning-ceiling increases, version bumps, build output, or unrelated edits.
