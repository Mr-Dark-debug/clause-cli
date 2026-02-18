package utils

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

// Version represents a semantic version.
type Version struct {
	Major      int
	Minor      int
	Patch      int
	Prerelease string
	Build      string
}

// ParseVersion parses a version string into a Version struct.
// Supports formats: "1", "1.2", "1.2.3", "1.2.3-alpha", "1.2.3-alpha+build"
func ParseVersion(v string) (*Version, error) {
	v = strings.TrimSpace(v)
	if v == "" {
		return nil, fmt.Errorf("empty version string")
	}

	// Remove 'v' prefix if present
	v = strings.TrimPrefix(v, "v")

	// Split build metadata
	buildIdx := strings.Index(v, "+")
	var build string
	if buildIdx != -1 {
		build = v[buildIdx+1:]
		v = v[:buildIdx]
	}

	// Split prerelease
	preIdx := strings.Index(v, "-")
	var prerelease string
	if preIdx != -1 {
		prerelease = v[preIdx+1:]
		v = v[:preIdx]
	}

	// Parse major.minor.patch
	parts := strings.Split(v, ".")
	if len(parts) == 0 || len(parts) > 3 {
		return nil, fmt.Errorf("invalid version format: %s", v)
	}

	ver := &Version{}

	major, err := strconv.Atoi(parts[0])
	if err != nil {
		return nil, fmt.Errorf("invalid major version: %s", parts[0])
	}
	ver.Major = major

	if len(parts) > 1 {
		minor, err := strconv.Atoi(parts[1])
		if err != nil {
			return nil, fmt.Errorf("invalid minor version: %s", parts[1])
		}
		ver.Minor = minor
	}

	if len(parts) > 2 {
		patch, err := strconv.Atoi(parts[2])
		if err != nil {
			return nil, fmt.Errorf("invalid patch version: %s", parts[2])
		}
		ver.Patch = patch
	}

	ver.Prerelease = prerelease
	ver.Build = build

	return ver, nil
}

// MustParseVersion parses a version string and panics on error.
func MustParseVersion(v string) *Version {
	ver, err := ParseVersion(v)
	if err != nil {
		panic(err)
	}
	return ver
}

// String returns the string representation of the version.
func (v *Version) String() string {
	s := fmt.Sprintf("%d.%d.%d", v.Major, v.Minor, v.Patch)
	if v.Prerelease != "" {
		s += "-" + v.Prerelease
	}
	if v.Build != "" {
		s += "+" + v.Build
	}
	return s
}

// ShortString returns the version without build metadata.
func (v *Version) ShortString() string {
	s := fmt.Sprintf("%d.%d.%d", v.Major, v.Minor, v.Patch)
	if v.Prerelease != "" {
		s += "-" + v.Prerelease
	}
	return s
}

// Compare compares two versions.
// Returns -1 if v < other, 0 if v == other, 1 if v > other.
func (v *Version) Compare(other *Version) int {
	if v.Major != other.Major {
		if v.Major < other.Major {
			return -1
		}
		return 1
	}

	if v.Minor != other.Minor {
		if v.Minor < other.Minor {
			return -1
		}
		return 1
	}

	if v.Patch != other.Patch {
		if v.Patch < other.Patch {
			return -1
		}
		return 1
	}

	// Compare prerelease
	// A version without prerelease is greater than one with
	if v.Prerelease == "" && other.Prerelease != "" {
		return 1
	}
	if v.Prerelease != "" && other.Prerelease == "" {
		return -1
	}
	if v.Prerelease != other.Prerelease {
		return comparePrerelease(v.Prerelease, other.Prerelease)
	}

	return 0
}

// comparePrerelease compares prerelease strings according to semver.
func comparePrerelease(a, b string) int {
	partsA := strings.Split(a, ".")
	partsB := strings.Split(b, ".")

	maxLen := len(partsA)
	if len(partsB) > maxLen {
		maxLen = len(partsB)
	}

	for i := 0; i < maxLen; i++ {
		if i >= len(partsA) {
			return -1
		}
		if i >= len(partsB) {
			return 1
		}

		partA := partsA[i]
		partB := partsB[i]

		// Try numeric comparison
		numA, errA := strconv.Atoi(partA)
		numB, errB := strconv.Atoi(partB)

		if errA == nil && errB == nil {
			if numA != numB {
				if numA < numB {
					return -1
				}
				return 1
			}
		} else {
			// String comparison
			if partA != partB {
				if partA < partB {
					return -1
				}
				return 1
			}
		}
	}

	return 0
}

// Less returns true if v < other.
func (v *Version) Less(other *Version) bool {
	return v.Compare(other) < 0
}

// Greater returns true if v > other.
func (v *Version) Greater(other *Version) bool {
	return v.Compare(other) > 0
}

// Equal returns true if v == other.
func (v *Version) Equal(other *Version) bool {
	return v.Compare(other) == 0
}

// GTE returns true if v >= other.
func (v *Version) GTE(other *Version) bool {
	return v.Compare(other) >= 0
}

// LTE returns true if v <= other.
func (v *Version) LTE(other *Version) bool {
	return v.Compare(other) <= 0
}

// CompareVersions compares two version strings.
// Returns -1 if v1 < v2, 0 if v1 == v2, 1 if v1 > v2.
func CompareVersions(v1, v2 string) (int, error) {
	ver1, err := ParseVersion(v1)
	if err != nil {
		return 0, fmt.Errorf("invalid first version: %w", err)
	}

	ver2, err := ParseVersion(v2)
	if err != nil {
		return 0, fmt.Errorf("invalid second version: %w", err)
	}

	return ver1.Compare(ver2), nil
}

// IsNewer returns true if current < latest.
func IsNewer(current, latest string) (bool, error) {
	ver1, err := ParseVersion(current)
	if err != nil {
		return false, err
	}

	ver2, err := ParseVersion(latest)
	if err != nil {
		return false, err
	}

	return ver1.Less(ver2), nil
}

// IsOlder returns true if current > latest.
func IsOlder(current, latest string) (bool, error) {
	ver1, err := ParseVersion(current)
	if err != nil {
		return false, err
	}

	ver2, err := ParseVersion(latest)
	if err != nil {
		return false, err
	}

	return ver1.Greater(ver2), nil
}

// VersionsEqual returns true if two versions are equal.
func VersionsEqual(v1, v2 string) (bool, error) {
	ver1, err := ParseVersion(v1)
	if err != nil {
		return false, err
	}

	ver2, err := ParseVersion(v2)
	if err != nil {
		return false, err
	}

	return ver1.Equal(ver2), nil
}

// BumpMajor increments the major version.
func (v *Version) BumpMajor() *Version {
	return &Version{
		Major: v.Major + 1,
		Minor: 0,
		Patch: 0,
	}
}

// BumpMinor increments the minor version.
func (v *Version) BumpMinor() *Version {
	return &Version{
		Major: v.Major,
		Minor: v.Minor + 1,
		Patch: 0,
	}
}

// BumpPatch increments the patch version.
func (v *Version) BumpPatch() *Version {
	return &Version{
		Major: v.Major,
		Minor: v.Minor,
		Patch: v.Patch + 1,
	}
}

// WithPrerelease returns a copy with the given prerelease.
func (v *Version) WithPrerelease(prerelease string) *Version {
	return &Version{
		Major:      v.Major,
		Minor:      v.Minor,
		Patch:      v.Patch,
		Prerelease: prerelease,
	}
}

// WithBuild returns a copy with the given build metadata.
func (v *Version) WithBuild(build string) *Version {
	return &Version{
		Major:      v.Major,
		Minor:      v.Minor,
		Patch:      v.Patch,
		Prerelease: v.Prerelease,
		Build:      build,
	}
}

// IsPrerelease returns true if the version has a prerelease tag.
func (v *Version) IsPrerelease() bool {
	return v.Prerelease != ""
}

// IsStable returns true if the version has no prerelease tag.
func (v *Version) IsStable() bool {
	return v.Prerelease == ""
}

// IsZero returns true if this is version 0.0.0.
func (v *Version) IsZero() bool {
	return v.Major == 0 && v.Minor == 0 && v.Patch == 0
}

// MajorVersion returns just the major version as a string.
func (v *Version) MajorVersion() string {
	return strconv.Itoa(v.Major)
}

// MinorVersion returns just the minor version as a string.
func (v *Version) MinorVersion() string {
	return strconv.Itoa(v.Minor)
}

// PatchVersion returns just the patch version as a string.
func (v *Version) PatchVersion() string {
	return strconv.Itoa(v.Patch)
}

// VersionRange represents a version range constraint.
type VersionRange struct {
	Min     *Version
	Max     *Version
	MinIncl bool // true for >=, false for >
	MaxIncl bool // true for <=, false for <
}

// ParseVersionRange parses a version range string.
// Examples: ">=1.0.0", ">=1.0.0 <2.0.0", "^1.2.3", "~1.2.3"
func ParseVersionRange(s string) (*VersionRange, error) {
	s = strings.TrimSpace(s)

	// Handle caret (^) - compatible with version
	if strings.HasPrefix(s, "^") {
		v, err := ParseVersion(s[1:])
		if err != nil {
			return nil, err
		}
		return &VersionRange{
			Min:     v,
			Max:     v.BumpMajor(),
			MinIncl: true,
			MaxIncl: false,
		}, nil
	}

	// Handle tilde (~) - approximately equivalent
	if strings.HasPrefix(s, "~") {
		v, err := ParseVersion(s[1:])
		if err != nil {
			return nil, err
		}
		return &VersionRange{
			Min:     v,
			Max:     v.BumpMinor(),
			MinIncl: true,
			MaxIncl: false,
		}, nil
	}

	// Handle comparison operators
	re := regexp.MustCompile(`^(>=|<=|>|<|=)?\s*(.+)$`)
	matches := re.FindStringSubmatch(s)
	if matches == nil {
		return nil, fmt.Errorf("invalid version range: %s", s)
	}

	op := matches[1]
	verStr := matches[2]

	v, err := ParseVersion(verStr)
	if err != nil {
		return nil, err
	}

	switch op {
	case ">=":
		return &VersionRange{Min: v, MinIncl: true}, nil
	case ">":
		return &VersionRange{Min: v, MinIncl: false}, nil
	case "<=":
		return &VersionRange{Max: v, MaxIncl: true}, nil
	case "<":
		return &VersionRange{Max: v, MaxIncl: false}, nil
	case "=", "":
		return &VersionRange{Min: v, Max: v, MinIncl: true, MaxIncl: true}, nil
	default:
		return nil, fmt.Errorf("unknown operator: %s", op)
	}
}

// Contains checks if a version is within the range.
func (r *VersionRange) Contains(v *Version) bool {
	if r.Min != nil {
		cmp := r.Min.Compare(v)
		if r.MinIncl && cmp > 0 {
			return false
		}
		if !r.MinIncl && cmp >= 0 {
			return false
		}
	}

	if r.Max != nil {
		cmp := r.Max.Compare(v)
		if r.MaxIncl && cmp < 0 {
			return false
		}
		if !r.MaxIncl && cmp <= 0 {
			return false
		}
	}

	return true
}

// String returns the string representation of the range.
func (r *VersionRange) String() string {
	var parts []string

	if r.Min != nil {
		if r.MinIncl {
			parts = append(parts, ">="+r.Min.String())
		} else {
			parts = append(parts, ">"+r.Min.String())
		}
	}

	if r.Max != nil {
		if r.MaxIncl {
			parts = append(parts, "<="+r.Max.String())
		} else {
			parts = append(parts, "<"+r.Max.String())
		}
	}

	return strings.Join(parts, " ")
}

// FindLatest finds the latest version from a slice.
func FindLatest(versions []string) (string, error) {
	if len(versions) == 0 {
		return "", fmt.Errorf("no versions provided")
	}

	var latest *Version
	var latestStr string

	for _, v := range versions {
		ver, err := ParseVersion(v)
		if err != nil {
			continue
		}

		if latest == nil || ver.Greater(latest) {
			latest = ver
			latestStr = v
		}
	}

	if latest == nil {
		return "", fmt.Errorf("no valid versions found")
	}

	return latestStr, nil
}

// SortVersions sorts a slice of version strings.
func SortVersions(versions []string) error {
	parsed := make([]*Version, len(versions))
	for i, v := range versions {
		ver, err := ParseVersion(v)
		if err != nil {
			return fmt.Errorf("invalid version %q: %w", v, err)
		}
		parsed[i] = ver
	}

	// Simple bubble sort (good enough for small slices)
	for i := 0; i < len(parsed)-1; i++ {
		for j := i + 1; j < len(parsed); j++ {
			if parsed[i].Greater(parsed[j]) {
				parsed[i], parsed[j] = parsed[j], parsed[i]
				versions[i], versions[j] = versions[j], versions[i]
			}
		}
	}

	return nil
}
