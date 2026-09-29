# AGENTS.md — plugin-package

Standalone plugin repo serving the `package` typed-step state-provision verb
(`verb:package`). The plugin is a Go module at `candy/plugin-package/` (module
path `github.com/opencharly/plugin-package/candy/plugin-package`); the root
`charly.yml` only declares `discover: candy` so the repo is a project and its
candy is scanned.

Canonical files:

- `candy/plugin-package/charly.yml` — the `plugin-package:` candy entity
  (`plugin:` block, `plan:` check).
- `candy/plugin-package/` — the Go source: `plugin.go`, `schema/package.cue`,
  `params/cue_types_gen.go`, `cmd/serve/main.go`.
- `.github/workflows/tag-on-merge.yml` — CalVer tag + `CHANGELOG/` on merge.
- `README.md` — user overview only; never agent guidance.

## Load these skills first (R0)

- `/charly-internals:plugin` — the plugin authoring reference: the `plugin:`
  block, the unified Provider model, host-coupled verbs and compiled-in
  placement, the per-plugin CUE-schema contract.
- `/charly-internals:install-plan` — the `InstallPlan` IR and the
  `SystemPackagesStep` this verb lowers into at build/deploy time.
- `/charly-check:check` — the declarative check-step surface the `package:` verb
  is authored through.
- `/charly-internals:git-workflow` — before any git/PR action.

## Build / validate / test

- `go build ./...` in `candy/plugin-package/` — compile the plugin module.
- `go test ./...` in `candy/plugin-package/` — the plugin's Go tests
  (`step_provider_coverage_test.go` locks the StepProvider surface).
- `charly box validate` at the repo root — the structural check (the candy +
  `plugin:` block, CUE schema).
- The merge gate is the **org-wide** `charly/pr-validator` (required check
  `validate / validate`, defined in `opencharly/.github`); this repo has **no**
  per-repo candy gate.

## Modify this repo

- Edit the `plugin-package:` candy entity, the Go source, and
  `schema/package.cue` **together** — the schema is the single source for the
  verb's `params/` struct.
- Keep it **compiled-in**: it is a host-coupled verb on the SDK/kit contract
  (`CheckVerbProvider` + `ProvisionActor` + `StepProvider`); do not describe it as
  a general out-of-process plugin.

## Landing

Load `/charly-internals:git-workflow` before any git/PR action; it owns the
landing mechanics. The authoritative rulebook is the umbrella `AGENTS.md` in
`opencharly/opencharly` and `charly/AGENTS.md` in the charly repo.
