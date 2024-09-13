// Teleport
// Copyright (C) 2024 Gravitational, Inc.
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU Affero General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU Affero General Public License for more details.
//
// You should have received a copy of the GNU Affero General Public License
// along with this program.  If not, see <http://www.gnu.org/licenses/>.

package backend

import (
	"bytes"
	"cmp"
	"fmt"
	"slices"
	"strings"
)

// Key is the unique identifier for an [Item].
type Key []string

// Separator is used as a separator between key parts
const Separator = "/"

// NewKey joins parts into path separated by Separator,
// makes sure path always starts with Separator ("/")
func NewKey(parts ...string) Key {
	return parts
}

// KeyFromString creates a Key from its textual representation.
func KeyFromString(key string) Key {
	if key == "" {
		return Key{}
	}

	k := strings.Split(key, Separator)
	if k[0] == "" && string(key[0]) == Separator {
		return k[1:]
	}
	return k
}

// ExactKey is like Key, except a Separator is appended to the result
// path of Key. This is to ensure range matching of a path will only
// math child paths and not other paths that have the resulting path
// as a prefix.
func ExactKey(parts ...string) Key {
	return append(NewKey(parts...), "")
}

// String returns the textual representation of the key with
// each component concatenated together via the [Separator].
func (k Key) String() string {
	if len(k) == 0 {
		return ""
	}

	return Separator + strings.Join(k, Separator)
}

// HasPrefix reports whether the key begins with prefix.
func (k Key) HasPrefix(prefix Key) bool {
	if len(prefix) == 0 || (len(prefix) == 1 && prefix[0] == "") {
		return true
	}

	prefixLen := len(prefix)
	if prefix.isExactKey() {
		prefixLen--
	}

	kLen := len(k)
	if k.isExactKey() {
		kLen--
	}

	n := min(kLen, prefixLen)

	for i := 0; i < n; i++ {
		if k[i] != prefix[i] {
			return false
		}
	}

	return prefixLen <= kLen
}

// TrimPrefix returns the key without the provided leading prefix string.
// If the key doesn't start with prefix, it is returned unchanged.
func (k Key) TrimPrefix(prefix Key) Key {
	if len(k) == 0 {
		return k
	}

	n := min(len(k), len(prefix))

	lastIndex := -1
	for i := 0; i < n; i++ {
		if k[i] != prefix[i] {
			break
		}
		lastIndex = i
	}

	switch lastIndex {
	case len(k) - 1:
		return Key{}
	case -1:
		return k
	default:
		return k[lastIndex+1:]
	}
}

func (k Key) PrependKey(p Key) Key {
	return append(slices.Clone(p), slices.Clone(k)...)
}

func (k Key) AppendKey(p Key) Key {
	return append(slices.Clone(k), slices.Clone(p)...)
}

func (k Key) ExactKey() Key {
	if k[len(k)-1] != "" {
		return ExactKey(k...)
	}

	return k
}

func (k Key) isExactKey() bool {
	return len(k) > 1 && k[len(k)-1] == ""
}

// HasSuffix reports whether the key ends with suffix.
func (k Key) HasSuffix(suffix Key) bool {
	for i, j := len(k)-1, len(suffix)-1; i >= 0 && j >= 0; i, j = i-1, j-1 {
		if k[i] != suffix[j] && i == len(k)-1 {
			return false
		}
	}

	return len(suffix) <= len(k)
}

// TrimSuffix returns the key without the provided trailing suffix string.
// If the key doesn't end with suffix, it is returned unchanged.
func (k Key) TrimSuffix(suffix Key) Key {
	if len(k) == 0 {
		return k
	}

	lastIndex := len(k)
	for i, j := len(k)-1, len(suffix)-1; i >= 0 && j >= 0; i, j = i-1, j-1 {
		if k[i] != suffix[j] && i == len(k)-1 {
			break
		}

		lastIndex = i
	}

	switch lastIndex {
	case 0:
		return Key{}
	case len(k):
		return k
	default:
		return k[:lastIndex]
	}
}

func (k Key) Components() []string {
	return k
}

func (k Key) Compare(o Key) int {
	if len(o) == 1 && o[0] == string(noEnd) {
		return -1
	}

	n := min(len(k), len(o))
	for i := 0; i < n; i++ {
		compare := strings.Compare(k[i], o[i])
		if compare != 0 {
			return compare
		}
	}

	switch len(k) {
	case len(o):
		return 0
	case 0:
		return -1
	default:
		if len(o) == 0 {
			return 1
		}

		return cmp.Compare(len(k), len(o))
	}
}

// Scan implement sql.Scanner, allowing a [Key] to
// be directly retrieved from sql backends without
// an intermediary object.
func (k *Key) Scan(scan any) error {
	switch key := scan.(type) {
	case []byte:
		if len(key) == 0 {
			return nil
		}

		raw := bytes.Clone(key)
		split := strings.Split(string(raw), Separator)
		*k = split[1:]
	case string:
		if key == "" {
			return nil
		}

		split := strings.Split(key, Separator)
		*k = split[1:]
	default:
		return fmt.Errorf("invalid Key type %T", scan)
	}

	return nil
}
