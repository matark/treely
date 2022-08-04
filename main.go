package main

import "bufio"
import "flag"
import "fmt"
import "io"
import "os"
import "encoding/json"
import "neweos.de/treely/tree"

const Version string = "1.0.1"

func main() {
  flag.BoolVar(&tree.Opts.A, "a", false, "all")
  flag.IntVar(&tree.Opts.L, "L", 256, "level")
  flag.Parse()

  fileInfo, _ := os.Stdin.Stat()

  if fileInfo.Mode() & os.ModeCharDevice != 0 {
    nodes := tree.Walk(flag.Arg(0))

    fmt.Println(tree.Tree(nodes))
    tree.Report(tree.Info{
      Files: len(tree.Paths.Files),
      Directories: len(tree.Paths.Directories),
    })

    os.Exit(0)
  }

  var result []json.RawMessage
  var info tree.Info

  reader := bufio.NewReader(os.Stdin)
  output := ""

  for {
    input, _, wrong := reader.ReadRune()

    if wrong != nil && wrong == io.EOF {
      break
    }

    output = output + string(input)
  }

  json.Unmarshal([]byte(output), &result)
  json.Unmarshal(result[1], &info)

  nodes := tree.Parse(fmt.Sprintf("[%s]", result[0]))
  fmt.Println(tree.Tree(nodes[0].Contents))
  tree.Report(info)
}
