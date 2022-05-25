package bytes2src

import (
	"fmt"
	"io"
	"strings"
)

type languageFormat struct {
	Initializer           string   // = new byte[] {
	ByteFmt               string   // 0xff, \xff, etc
	BlockEnds             []string // { "[", "]" }
	LastByteCommaRequired bool
}

type stringDumper struct {
	sb        strings.Builder
	r         byteReadSeeker
	inputSize int
	lf        languageFormat
	nread     int
	width     int
}

type byteReadSeeker interface {
	io.ByteReader
	io.Seeker
}

type Language int

const (
	Go Language = iota
	JavaScript
	CSharp
)

var languageFormats = map[Language]languageFormat{
	Go: {
		Initializer:           "= []byte",
		ByteFmt:               "0x%02x",
		BlockEnds:             []string{"{", "}"},
		LastByteCommaRequired: true,
	},
	JavaScript: {
		Initializer:           "= new UInt8Array[]",
		ByteFmt:               "0x%02x",
		BlockEnds:             []string{"([", "])"},
		LastByteCommaRequired: false,
	},
	CSharp: {
		Initializer:           "= new []byte",
		ByteFmt:               "0x%02x",
		BlockEnds:             []string{"{", "}"},
		LastByteCommaRequired: false,
	},
}

func (sd *stringDumper) Init() error {
	sz, err := sd.r.Seek(0, io.SeekEnd)
	if err != nil {
		return fmt.Errorf("error getting input data size: %w", err)
	}

	sd.inputSize = int(sz)

	if _, err = sd.r.Seek(0, io.SeekStart); err != nil {
		return fmt.Errorf("error seeking back to start of data: %w", err)
	}

	return nil
}

func (sd *stringDumper) WriteHeader() {
	sd.sb.WriteString(sd.lf.Initializer)
	sd.sb.WriteString(sd.lf.BlockEnds[0])
	sd.sb.WriteRune('\n')
}

func (sd *stringDumper) WriteBytes() error {
	for i := 0; i < sd.inputSize; i++ {
		b, err := sd.r.ReadByte()
		if err != nil {
			return fmt.Errorf("failed to read byte of input: %w", err)
		}

		if i%sd.width == 0 && i != 0 {
			sd.sb.WriteRune('\n')
		}

		sd.sb.WriteString(fmt.Sprintf(sd.lf.ByteFmt, b))

		if sd.lf.LastByteCommaRequired || i < sd.inputSize-1 {
			sd.sb.WriteRune(',')
		}

		if i < sd.inputSize-1 && (i+1)%sd.width != 0 {
			sd.sb.WriteRune(' ')
		}
	}

	return nil
}

func (sd *stringDumper) WriteFooter() {
	sd.sb.WriteRune('\n')
	sd.sb.WriteString(sd.lf.BlockEnds[1])
}

func DumpString(r byteReadSeeker, lang Language, width int) (*string, error) {

	sd := &stringDumper{
		r:     r,
		lf:    languageFormats[lang],
		width: width,
	}

	if err := sd.Init(); err != nil {
		return nil, fmt.Errorf("failed to create stringDumper: %w", err)
	}

	sd.WriteHeader()

	sd.WriteBytes()

	sd.WriteFooter()

	s := sd.sb.String()
	return &s, nil
}
