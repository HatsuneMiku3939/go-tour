package main

import (
  "io"
  "os"
  "strings"
)

type rot13Reader struct {
  r io.Reader
}

func (rot *rot13Reader) Read(p []byte) (n int, err error) {
  lut := map[byte]byte {
    'A' : 'N',
    'B' : 'O',
    'C' : 'P',
    'D' : 'Q',
    'E' : 'R',
    'F' : 'S',
    'G' : 'T',
    'H' : 'U',
    'I' : 'V',
    'J' : 'W',
    'K' : 'X',
    'L' : 'Y',
    'M' : 'Z',
    'N' : 'A',
    'O' : 'B',
    'P' : 'C',
    'Q' : 'D',
    'R' : 'E',
    'S' : 'F',
    'T' : 'G',
    'U' : 'H',
    'V' : 'I',
    'W' : 'J',
    'X' : 'K',
    'Y' : 'L',
    'Z' : 'M',
    'a' : 'n',
    'b' : 'o',
    'c' : 'p',
    'd' : 'q',
    'e' : 'r',
    'f' : 's',
    'g' : 't',
    'h' : 'u',
    'i' : 'v',
    'j' : 'w',
    'k' : 'x',
    'l' : 'y',
    'm' : 'z',
    'n' : 'a',
    'o' : 'b',
    'p' : 'c',
    'q' : 'd',
    'r' : 'e',
    's' : 'f',
    't' : 'g',
    'u' : 'h',
    'v' : 'i',
    'w' : 'j',
    'x' : 'k',
    'y' : 'l',
    'z' : 'm',
  }

  n = 0
  for {
    buffer := make([]byte, 1024)
    buffer_size, buffer_error := rot.r.Read(buffer)

    err = buffer_error
    if buffer_error != nil {
      break
    }

    for i := 0; i < buffer_size; i++ {
      p[i] = lut[buffer[i]]
    }

    n += buffer_size
  }

  return
}

func main() {
  s := strings.NewReader("Lbh penpxrq gur pbqr!")
  r := rot13Reader {s}
  io.Copy(os.Stdout, &r)
}
