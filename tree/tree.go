package tree

import "encoding/json"

var style = struct {
  Indent     string
  Bar        string
  Branch     string
  LastBranch string
}{
  "    ",
  "│   ",
  "├── ",
  "└── ",
}

func Parse(input string) []Input {
  var nodes []Input
  wrong := json.Unmarshal([]byte(input), &nodes)

  if wrong == nil {
    return nodes
  }

  return []Input{}
}

func Tree(nodes []Input) string {
  return treeFunc(nodes, "")
}

func treeFunc(nodes []Input, prefix string) string {
  var result string = ""

  for index, node := range nodes {
    indent := style.Bar
    branch := style.Branch

    if index + 1 >= len(nodes) {
      indent = style.Indent
      branch = style.LastBranch
    }

    result = result + prefix + treeBar(index) + "\n"
    result = result + prefix + branch + node.Name + "\n"

    if len(node.Contents) == 0 {} else {
      result = result + treeFunc(node.Contents, prefix + indent)
    }
  }

  return result
}

func treeBar(index int) string {
  bar := style.Bar

  if index == 0 {
    bar = "."
  }

  return bar
}
