
package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type item struct {
	name  string
	age   int
	hobby []string
}
type vertex struct {
	value  item
	chiave string
}
type graph struct {
	vertices map[string]*vertex
	adj      map[*vertex][]*vertex
}

func NewVertex(nome string, age int, hobby []string, key string) *vertex {
	var i item
	var v vertex
	i.name = nome
	i.age = age
	i.hobby = hobby
	v.value = i
	v.chiave = key
	return &v
}
func graphNew(n int) *graph { //returns the adress of a graph with n vertices
	g := new(graph)
	m := make(map[string]*vertex)
	ad := make(map[*vertex][]*vertex)
	for i := 1; i <= n; i++ {
		k := strconv.Itoa(i)
		m[k] = nil
	}
	(*g).vertices = m
	(*g).adj = ad
	return g
}
func pritnVertices(g *graph) { // prints only the vertices
	for k, v := range (*g).vertices {
		fmt.Println(k, " ", (*v).value.name, " hobbies: ", (*v).value.hobby)
	}
}
func printFollowers(g *graph) { //prints adj list
	for _, v := range (*g).vertices {
		fmt.Print((*v).value.name, " follows:")
		for j := 0; j < len((*g).adj[v]); j++ {
			fmt.Print(" ", (*g.adj[v][j]).value.name)
		}
		fmt.Println()
	}
}

//Scrivete una funzione che data una stringa A stampi gli hobby dell’utente di nome A e
//l’elenco di tutti gli hobby delle persone che seguono A.
func printHobbyFollowers(nome string, g *graph) {
	for _, v := range (*g).vertices {
		if nome == (*v).value.name {
			fmt.Println((*v).value.hobby)
		}

    for j := 0; j < len((*g).adj[v]); j++ {
      if (*g.adj[v][j]).value.name == nome  { // if a person has A in followed list, print that person's hobbies
                                              //if that vertex has A as a followed person print that vertex hobbies
        fmt.Println((*v).value.hobby)
      }
    }
	}
}
func printHobbyFollowed(nome string, g *graph)
func main() {
	g := graphNew(5)
	f, err := os.Open("inputes2.txt") //here i just create the graph
	if err != nil {
		log.Fatal(err)
	}
	sc := bufio.NewScanner(f)
	for sc.Scan() {
		r := strings.Fields(sc.Text())
		age, _ := strconv.Atoi(r[2])
		hobbies := strings.Split(r[3], ",")
		v := NewVertex(r[1], age, hobbies, r[0])
		(*g).vertices[r[0]] = v
	}
	f2, err2 := os.Open("followers.txt")
	if err2 != nil {
		log.Fatal(err2)
	}
	sc2 := bufio.NewScanner(f2)
	for sc2.Scan() {
		riga := strings.Fields(sc2.Text()) //in this part i fill the map, which will be my adj map
		vert := (*g).vertices[riga[0]]
		followers := strings.Split(riga[2], ",")
		for i := 0; i < len(followers); i++ {
			follows := (*g).vertices[followers[i]]
			(*g).adj[vert] = append((*g).adj[vert], follows)
		}
	}
	printHobbyFollowers("Giacomo", g)
}
