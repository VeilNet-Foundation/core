// router/amr.go
func BuildRandomPath(exclude []string) []Node {
  depth := 3 + rand.Intn(5) // 3–7 узлов
  path := make([]Node, 0, depth)
  
  for i := 0; i < depth; i++ {
    node := selectByReputationAndLatency(exclude)
    path = append(path, node)
    exclude = append(exclude, node.ID)
  }
  
  return path
}