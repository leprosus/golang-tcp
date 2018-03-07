package tcp

import "fmt"

type Version struct {
	Major uint
	Minor uint
}

func (ver *Version) String() string {
	return fmt.Sprintf("%d.%d", ver.Major, ver.Minor)
}

func (ver *Version) IsEqual(other Version) bool {
	return ver.Major == other.Major &&
		ver.Minor == other.Minor
}

func (ver *Version) IsLess(other Version) bool {
	return ver.Major < other.Major ||
		(ver.Major == other.Major && ver.Minor < other.Minor)
}

func (ver *Version) IsMore(other Version) bool {
	return ver.Major > other.Major ||
		(ver.Major == other.Major && ver.Minor > other.Minor)
}
