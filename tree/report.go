package tree

import "fmt"

type Info struct {
  Files       int `json:"files"`
  Directories int `json:"directories"`
}

func Report(info Info) {
  var words []string = make([]string, 0)
  words = append(words, "directories", "files")

  if info.Directories == 1 {
    words[0] = "directory"
  }

  if info.Files == 1 {
    words[1] = "file"
  }

  fmt.Printf("%d %s, %d %s\n", info.Directories, words[0], info.Files, words[1])
}
