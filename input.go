package bytes2src

import (
	"bufio"
	"fmt"
	"io"
	"regexp"
	"strconv"
	"strings"
)

var byteMatcher = regexp.MustCompile("[0-9a-fA-F]{2}")

func ReadHexDump(r io.Reader) (int, []byte, error) {
	s := bufio.NewScanner(r)

	var parsedBytes []byte

	n := 0

	for s.Scan() {
		line := s.Text()

		cols := strings.Fields(line)

		if len(cols) < 3 {
			continue
		}

		splitBytes := func(s string) []string {
			rs := []rune(s)
			var bts []string

			if len(s)%2 != 0 {
				return nil
			}

			for x := 0; x < len(s); x += 2 {
				hexPair := string(rs[x : x+2])
				if !byteMatcher.MatchString(hexPair) {
					return nil
				}
				bts = append(bts, hexPair)
			}

			return bts
		}

		for _, col := range cols[1 : len(cols)-1] {
			if vs := splitBytes(col); vs != nil {
				for _, val := range vs {
					v, err := strconv.ParseUint(val, 16, 8)
					if err != nil {
						return 0, nil, fmt.Errorf("could not parse int from match: %w", err)
					}
					parsedBytes = append(parsedBytes, byte(v))
					n++
				}
			}
		}
	}

	return n, parsedBytes, nil
}
