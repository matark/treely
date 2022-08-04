package tree

import "fmt"
import "io/fs"
import "io/ioutil"
import "path/filepath"
import "strings"

var Opts struct {
  A bool
  L int
}

var Paths struct {
  Files       []string
  Directories []string
}

type Input struct {
  Name     string  `json:"name"`
  Type     string  `json:"type"`
  Contents []Input `json:"contents"`
}

func Walk(path string) []Input {
  entries, wrong := ioutil.ReadDir(path)

  if wrong == nil {
    return walkFunc(path, entries, 0)
  }

  fmt.Printf("No such file or directory %s\n", path)
  return []Input{}
}

func walkFunc(parentPath string, items []fs.FileInfo, level int) []Input {
  var nodes []Input

  for _, item := range items {
    fullPath := filepath.Join(parentPath, item.Name())

    if !Opts.A && strings.HasPrefix(item.Name(), ".") {
      continue
    }

    if item.IsDir() && level < Opts.L {
      Paths.Directories = append(Paths.Directories, fullPath)
      entries, _ := ioutil.ReadDir(fullPath)
      nodes = append(nodes, Input{
        Name: item.Name(),
        Contents: walkFunc(fullPath, entries, level + 1),
      })
    } else {
      Paths.Files = append(Paths.Files, fullPath)
      nodes = append(nodes, Input{
        Name: item.Name(),
      })
    }
  }

  return nodes
}
