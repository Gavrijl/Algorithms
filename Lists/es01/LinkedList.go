package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

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
func addNewNode(l *linkedList, val int) {
	node := newNode(val)
	node.next = (*l).head
	(*l).head = node
}
func printList(l linkedList) {
	var p *listNode
	p = l.head
	for p != nil {
		fmt.Print((*p).item, " ")
		p = p.next
	}
}
func searchList(l linkedList, val int) (bool, *listNode) {
	var p *listNode
	p := l.head
	for p != nil {
		if (*p).item == val {
			return true, p
		}
		p = p.next
	}
	return false, nil
}

func removeItem(l *linkedList, elem int) bool {
	var curr, prec *listNode
	curr = l.head
	prec = nil
	for curr != nil {
		if (*curr).item == elem {
			if (*prec).next == nil {
				l.head = l.head.next
			} else {
				prec.next = curr.next
			}
			return true
		}
		prec = curr
		curr = curr.next
	}
}
func main() {
	var l linkedList
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {

		r := strings.Fields(scanner.Text())
		if len(r) == 0 {
			continue
		}
		if len(r) == 2 {
			n, _ := strconv.Atoi(r[1])
			switch r[0] {
			case "+":
				pres, _ := searchList(l, n)
				if !pres {
					addNewNode(&l, n)
				}
			case "-":
				pres, _ := searchList(l, n)
				if pres {
					removeItem(&l, n)
				}
			case "?":
				pres, _ := searchList(l, n)
				if pres {
					fmt.Println("Presente")
				} else {
					fmt.Println("Non presente")
				}
			}
		}
		switch r[0] {
		case "c":
			c := 0
			var p *listNode
			p := l.head
			for p != nil {
				c++
			}
			fmt.Println(c)
		case "p":
			printList(l)
		case "o":

		case "f":
			os.Exit(0)
		case "d":
			var p *listNode
			p := l.head
			for p != nil {
				_ := RemoveItem(l, (*p).val)
			}
		}

	}
}
