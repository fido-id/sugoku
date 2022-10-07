package main

import (
	"bufio"
	"fmt"
	"github.com/mauroartizzu/sugoku/internal/board"
	"math/rand"
	"os"
	"strconv"
	"testing"
	"time"
)

func TestMain(m *testing.M) {
	//log.Println("Do stuff BEFORE the tests!")
	exitVal := m.Run()
	//log.Println("Do stuff AFTER the tests!")

	os.Exit(exitVal)
}

func getRandomSudoku() board.Board {
	files, _ := os.ReadDir("samples/")

	rand.Seed(time.Now().Unix())
	fileIndex := rand.Intn(len(files))
	file, err := os.Open(fmt.Sprintf("samples/%s", files[fileIndex].Name()))
	if err != nil {
		fmt.Println("Impossibile aprire il file:", os.Args[1])
		panic(err)
	}

	scan := bufio.NewScanner(file)

	scan.Split(bufio.ScanLines)

	matrix := make([][]int, 0)

	for scan.Scan() {
		row := make([]int, 0)
		for _, c := range scan.Text() {
			result, _ := strconv.Atoi(string(c))
			row = append(row, result)
		}
		matrix = append(matrix, row)
	}

	b := board.Board{}

	b.New(matrix)

	return b
}
