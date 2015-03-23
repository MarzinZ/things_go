package main

import (
	"fmt"
	"strconv"
	"encoding/json"
	"io/ioutil"
	"os"
	"github.com/docopt/docopt-go"
)

const (
	SAVEFILE = "todo.json"
)

type ToDoList struct {
	List	[]Item
}

type Item struct {
	Content string
	Finish	bool
}

func (t *ToDoList) Len() int {
	return len(t.List)
}

func (t *ToDoList) Add(item *Item) {
	t.List = append(t.List, *item)
}

func (t *ToDoList) Remove(index int) bool {
	if index < 0 || index >= t.Len() {
		return false
	}
	t.List = append(t.List[:index], t.List[index+1:]...)
	return true
}

func (t *ToDoList) Done(index int) {
	t.List[index].Finish = true
}

func (t ToDoList) String() string {
	var str string
	for index, item := range t.List {
		if item.Finish {
			str = str + "\n" + strconv.Itoa(index) + ": " + item.Content + "\t\tFinish"
		} else {
			str = str + "\n" + strconv.Itoa(index) + ": " + item.Content + "\t\ttodo"
		}
	}
	return str + "\n"
}

func (t *ToDoList) Show() {
	fmt.Println(t)
}

func (t *ToDoList) Update(saveFile string) {
	byteData, _ := json.MarshalIndent(t, "", "\t")
	ioutil.WriteFile(saveFile, byteData, 0666)
}

func (t *ToDoList) Init(saveFile string) {
	if _, err := os.Stat(saveFile); err == nil {
		byteData, _ := ioutil.ReadFile(saveFile)
		_ = json.Unmarshal(byteData, &t)
	}
}

func main() {
		usage := `things_go 0.1
Usage:
  things add <content> 
  things list 
  things rm <index>
  things done <index>
  things -h | --help
  things --version
Options:
  -h --help     Show this screen.
  --version     Show version.`

  	args, _ := docopt.Parse(usage, nil, true, "things_go 0.1", false)
	var things ToDoList
	things.Init(SAVEFILE)
	switch {
	case args["add"]:
		content := args["<content>"].(string)
		item := Item{Content:content, Finish:false}
		things.Add(&item)
	case args["list"]:
		fmt.Println("ToDoList:")
	case args["rm"]:
		index, _ := strconv.Atoi(args["<index>"].(string))
		things.Remove(index)
	case args["done"]:
		index, _ := strconv.Atoi(args["<index>"].(string))
		things.Done(index)
	}
	things.Update(SAVEFILE)
	things.Show()
}