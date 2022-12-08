package main
import (
   "fmt"
    "log"
    "bufio"
    "os"
    "strconv"
    "strings"
)

type graph struct {
n int // numero di vertici
adiacenti []* linkedList
}
type listNode struct {
	item int
	next *listNode
}
type linkedList struct {
  head *listNode
}

func newNode(val int) *listNode {
	node := new(listNode) //new alloca lo spazio per un node, N.B node è un puntatore
	(*node).item = val //dereferenzio puntatore anche se in go non è necessario
											// qua sto dicendo "(ciò a cui punta node).item = val"
	return node        //ritorna il *puntatore* al nodo (new ritorna un puntatore)
}
func NewGraph(n int) *graph{
  var a []*linkedList
  //fmt.Println(l)
  for i:=0;i<n;i++{
    l := &linkedList{nil}
    a=append(a,l)
  }
return &graph{n,a}
}
func addNewNode(l *linkedList, val int) {
	node := newNode(val)
	node.next = (*l).head
	(*l).head = node
}
func NewEdge(g *graph,v1 int ,v2 int){
addNewNode(g.adiacenti[v1],v2)
}
func printList(l linkedList) {
	var p *listNode
	p = l.head
	for p != nil {
		fmt.Print((*p).item, " ")
		p = p.next
	}
}
func PrintVertices(g *graph, v int){
printList(*g.adiacenti[v])
}

func PrintGraph(g *graph,n int){
  for i:=0;i<5;i++{
  fmt.Print(i," è adiacente a ")
  PrintVertices(g,i)
  fmt.Println()
  }
}
func main() {
  f,err:= os.Open("input.txt")
  if err!=nil{
      log.Fatal(err)
    }

  defer f.Close()
  scanner := bufio.NewScanner(f)
  n,_:=strconv.Atoi(os.Args[1])
  g:=NewGraph(n)
  for scanner.Scan(){
    if scanner.Text()==""{
      break
    }
  arr:=strings.Split(scanner.Text()," ")
  z,_:=strconv.Atoi(arr[0])
  x,_:=strconv.Atoi(arr[1])
  NewEdge(g,z,x)
  }
PrintGraph(g,n)
}
