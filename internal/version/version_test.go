package version

import (
	"testing"
)

/*
* @Author: zouyx
* @Email: 1003941268@qq.com
* @Date:   2025/9/9 下午2:52
* @Package:
 */
func FuzzNewVersion(f *testing.F) {
	testcases := []string{"go1.16.14", " ", "......", "1", "1.2.3-beta.1", "1.2.3+foo", "2.3.4-alpha.1+bar", "lorem ipsum"}

	for _, tc := range testcases {
		f.Add(tc)
	}

	f.Fuzz(func(t *testing.T, a string) {
		v, _ := NewVersion(a)
		if v != nil {
			t.Log(v.String())

		}
	})
}
func TestValidateMetadata(t *testing.T) {
	tests := []struct {
		meta     string
		expected error
	}{
		{"foo", nil},
		{"alpha.1", nil},
		{"alpha.01", nil},
		{"foo☃︎", ErrInvalidMetadata},
		{"alpha.0-1", nil},
		{"al-pha.1Phe70CgWe050H9K1mJwRUqTNQXZRERwLOEg37wpXUb4JgzgaD5YkL52ABnoyiE", nil},
	}

	for _, tc := range tests {
		if err := validateMetadata(tc.meta); err != tc.expected {
			t.Errorf("Unexpected error %q for build %q", err, tc.meta)
		}
	}
}
