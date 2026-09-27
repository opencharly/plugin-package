package pkgverb

import (
	"testing"

	"github.com/opencharly/spec/spec"
)

// TestPackageVerb_MaterializeStep_FallbackName pins the cross-distro resolver's fallback contract
// at the materialization boundary: with NO `package_map` (or a map that names no running-distro
// tag), the declared package name is used VERBATIM. It FAILS if the materializer starts dropping or
// rewriting a package on a distro the map does not cover — the cross-distro behaviour this candy
// owns via kit.ResolvePackageName. (The sibling TestPackageVerb_StepProvider covers the mapped case.)
func TestPackageVerb_MaterializeStep_FallbackName(t *testing.T) {
	// no package_map at all
	s := (verb{}).MaterializeStep(
		&spec.Op{PluginInput: map[string]any{"package": "htop"}}, "", "tools", "deb", []string{"debian:13", "debian"})
	ps := s.(*spec.SystemPackagesStep)
	if len(ps.Packages) != 1 || ps.Packages[0] != "htop" {
		t.Fatalf("fallback Packages = %v, want [htop]", ps.Packages)
	}

	// a package_map naming a DIFFERENT distro must not override the running distro's name
	s2 := (verb{}).MaterializeStep(
		&spec.Op{PluginInput: map[string]any{"package": "htop", "package_map": map[string]any{"fedora": "htop-fedora"}}},
		"", "tools", "deb", []string{"debian:13", "debian"})
	p2 := s2.(*spec.SystemPackagesStep)
	if len(p2.Packages) != 1 || p2.Packages[0] != "htop" {
		t.Fatalf("non-matching map Packages = %v, want [htop]", p2.Packages)
	}
}
