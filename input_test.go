package bytes2src

import (
	"strings"
	"testing"
)

var testDumps = map[int]string{
	16: `
00000000: af76 0b83 ecc5 03f3 7967 a785 c029 afa9  .v......yg...)..
	`,
	32: `
00000000: 10bb 7b26 a155 6691 e830 51cb 0bda 3073  ..{&.Uf..0Q...0s
00000010: 9bd8 744f 6f3a 2e3d aeaa 4bf5 6b23 c94b  ..tOo:.=..K.k#.K
	`,
	37: `
00000000: 441c a9dc c99f 19ea 7743 5b33 2dfe b300  D.......wC[3-...
00000010: b8ac a960 de93 bbac c2d4 53f7 0384 1ac6  ..........S.....
00000020: 8650 195c 4d                             .P.\M
	`,
	1: `
00000000: 8f                                       .
	`,
}

func TestReadHexdump(t *testing.T) {
	for sz, xxd := range testDumps {
		n, _, err := ReadHexDump(strings.NewReader(xxd))
		if err != nil {
			t.Errorf("Error reading hex dump: %v", err)
		}

		if n != sz {
			t.Errorf("wrong number of bytes parsed: want %d, got %d", sz, n)
		}
	}

}
