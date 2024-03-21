package main

import "fmt"

type (
	Database interface {
		Insert() error
		Update() error
	}

	PostgresDB struct{}
	MockDb     struct{}
)

func (p *PostgresDB) Insert() error {
	fmt.Println("PostgresDB Insert")
	return nil
}

func (p *PostgresDB) Update() error {
	fmt.Println("PostgresDB Update")
	return nil
}

func (p *MockDb) Insert() error {
	fmt.Println("MockDb Insert")
	return nil
}

func (p *MockDb) Update() error {
	fmt.Println("MockDb Update")
	return nil
}

func InsertPlayerItem(db Database) error {
	return db.Insert()
}
func UpdatePlayerItem(db Database) error {
	return db.Update()
}

func main() {
	postgres := &PostgresDB{}
	mock := &MockDb{}

	InsertPlayerItem(postgres)
	UpdatePlayerItem(postgres)

	InsertPlayerItem(mock)
	UpdatePlayerItem(mock)
}

// import (
// 	"encoding/json"
// 	"fmt"
// )

// type Player struct {
// 	UserName string `json:"username"`
// 	Level    uint   `json:"level"`
// 	Status   string `json:"status"`
// 	Class    string `json:"class"`
// }

// func (p Player) GetUserName() string {
// 	return p.UserName
// }

// func (p *Player) LevelUp() {
// 	p.Level++
// }

// func main() {

// 	player1 := Player{
// 		UserName: "player1",
// 		Level:    99,
// 		Status:   "Active",
// 		Class:    "Warrior",
// 	}

// 	jsonData, _ := json.Marshal(&player1)
// 	fmt.Println(string(jsonData))

// 	fmt.Println("Username : " + player1.GetUserName())

// 	player1.LevelUp()

// 	fmt.Println("Level : ", player1.Level)

// 	// b := new(int)
// 	// c := new(*int)

// 	// a := 1
// 	// b = &a
// 	// c = &b

// 	// fmt.Println("address a ", &a)
// 	// fmt.Println("address b ", &b)
// 	// fmt.Println("address c ", &c)

// 	// a := make(map[int]int)
// 	// a := map[int]int{
// 	// 	1: 1,
// 	// 	2: 2,
// 	// 	3: 3,
// 	// }

// 	// if _, ok := a[3]; ok {
// 	// 	fmt.Println("OK")
// 	// } else {
// 	// 	fmt.Println("NOT OK")
// 	// }

// 	// dfs(graph, 1, make(map[int]bool))

// }

// var graph = map[int][]int{
// 	1: {2, 3},
// 	2: {4, 5},
// 	3: {6},
// }

// func dfs(graph map[int][]int, vertex int, visited map[int]bool) {
// 	visited[vertex] = true
// 	for _, v := range graph[vertex] {
// 		fmt.Printf("-> %d", v)
// 		dfs(graph, v, visited)
// 	}
// }
