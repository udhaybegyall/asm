// +build !amd64

package mem

func ContainsByte(haystack []byte, needle byte) bool {
	return containsGeneric(haystack, needle)
}
