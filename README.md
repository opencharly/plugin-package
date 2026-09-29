# plugin-package

The `package` typed-step state-provision verb for OpenCharly — probe and install
a system package across distros.

The plugin is a host-coupled verb on the SDK/kit contract (`CheckVerbProvider` +
`ProvisionActor` + `StepProvider`), so it is **compiled-in only**.

## What it provides

| Capability | Surface |
|---|---|
| `verb:package` | the `package:` check/provision typed step |

- **CHECK** — `rpm -q` / `dpkg -s` / `pacman -Q` probe, plus an optional version
  match, via the live check engine.
- **ACT (runtime)** — render the `dnf` / `apt-get` / `pacman` install.
- **ACT (build/deploy)** — lower into a `SystemPackagesStep`; the host
  materializes the kit descriptor (format + cross-distro name via
  `kit.ResolvePackageName`).

## How to use it

Compose the plugin candy in a box or check bed's `candy:` list:

```yaml
- '@github.com/opencharly/plugin-package/candy/plugin-package:<tag>'
```

Then author the verb in a plan:

```yaml
- check: bash is installed
  package:
    package: bash
    installed: true
  context: [runtime]
```

| Field | Meaning |
|---|---|
| `package` | the package name (also the scalar-sugar primary) |
| `installed` | assert installed (default) or absent |
| `version` | one or more acceptable versions |
| `package_map` | per-distro name overrides |

## Layout

- `candy/plugin-package/` — the plugin module: `plugin.go` (provider + meta),
  `schema/package.cue` (the self-contained `#PackageInput`),
  `params/cue_types_gen.go`, and `cmd/serve/main.go`.
- `charly.yml` — the root project manifest (`discover: candy`).
- `.github/workflows/tag-on-merge.yml` — CalVer tag + `CHANGELOG/` on merge.

## Related

- Owning skill: `/charly-internals:plugin` — the plugin/provider model,
  host-coupled verbs, and compiled-in placement (the candy carries no `skill:`
  entity of its own; the gap is tracked in
  [opencharly/opencharly#291](https://github.com/opencharly/opencharly/issues/291)).
- `/charly-check:check` — the declarative check-step surface the `package:` verb
  is authored through.
- [`opencharly/charly`](https://github.com/opencharly/charly) — the charly CLI.
