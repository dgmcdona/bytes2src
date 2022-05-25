# bytes2src

This is a simple utility for converting binary data into a source code
representation in multiple languages.

Usage of bytes2src:
  -f string
    	input file from which to read bytes
  -l string
    	programming language for source code output
  -o string
    	output to file <path>
  -w int
    	width of the source code array output (default 8)
  -x	interpret input data as hex dump

## Examples
```
$ cat /dev/urandom | head -c 1 | xxd | bytes2src -x -l golang
= []byte{
0xc6,
}

$ cat /dev/urandom | head -c 64 | bytes2src -w 16 -l csharp
= new []byte{
0x87, 0x8c, 0x99, 0x9c, 0x78, 0xad, 0x58, 0x20, 0xe7, 0xd4, 0x8a, 0x6c, 0x03, 0xe2, 0x1f, 0x39,
0x83, 0xf1, 0xc2, 0xf8, 0x56, 0x31, 0xec, 0x85, 0x1b, 0x78, 0x0d, 0x6c, 0x8a, 0xad, 0x0d, 0xf6,
0x43, 0x68, 0x80, 0x33, 0xc6, 0xd7, 0xdc, 0x24, 0xf7, 0x29, 0x60, 0x53, 0xa7, 0xa2, 0x62, 0xa3,
0xc4, 0x1c, 0x3f, 0x9e, 0x90, 0x4e, 0x7a, 0xd9, 0x29, 0x6e, 0xa8, 0xc2, 0xc7, 0xca, 0xdd, 0x25
}
```
