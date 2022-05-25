package bytes2src

import (
	"bytes"
	"strings"
	"testing"
)

func TestDumpString(t *testing.T) {
	type test struct {
		input []byte
		want  string
	}
	tests := []test{
		{
			input: []byte{0x00, 0x00, 0x00, 0x00, 0x00, 0x00},
			want: `= []byte{
0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
}`,
		},
		{
			input: []byte{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00},
			want: `= []byte{
0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
0x00,
}`,
		},
	}

	for _, test := range tests {
		br := bytes.NewReader(test.input)
		s, err := DumpString(br, Go, 8)
		if err != nil {
			t.Errorf("failed to dump string: %v", err)
		}

		if strings.Compare(*s, test.want) != 0 {
			t.Errorf("strings do not match: WANT \n%s\nGOT \n%s\n", test.want, *s)
			for i, c := range test.want {
				if c != []rune(*s)[i] {
					t.Logf("Character %d: want %c, got %c", i, c, []rune(*s)[i])
				}
			}
		}
	}

}
