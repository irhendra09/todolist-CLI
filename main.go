package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var model = make([]string, 10)
var scanner = bufio.NewScanner(os.Stdin)

func main() {
	viewShowTodo()
}


func showTodo(){
	fmt.Println("==TODOLIST APP==")
	for i, todo := range model{
		var no = i +1
		if todo != ""{
			fmt.Printf("%d. %s\n", no, todo)
		}
	}
}

func addTodo(todo string){
	isFull := true
	for _, item:= range model{
		if item == "" {
			isFull = false
			break
		}
	}

	if isFull {
		temp := model
		model = make([]string, len(model)*2)
		copy(model, temp)
	}

	for i, item := range model {
		if item == "" {
			model[i] = todo
			break
		}
	}
}

func removeTodo (number int) bool{
	if number - 1 >= len(model) {
		return false
	} else if model[number -1] == "" {
		return false
	} else {
		for i := number -1; i < len(model); i++ {
			if i == len(model)-1 {
				model[i] = ""
			} else {
				model[i] = model [i + 1]
			}
		}
		return true
	}
}

func input (info string) string {
	fmt.Print(info + " : ")
	scanner.Scan()
	data := scanner.Text()
	return data
}

func viewShowTodo() {
	for {
		showTodo()
		fmt.Println("Menu")
		fmt.Println("1. Tambah")
		fmt.Println("2. Hapus")
		fmt.Println("x. Keluar")
		input := input("Pilih")

		switch input {
		case "1":
			viewAddTodo()
		case "2":
			viewRemoveTodo()
		case "x":
			return
		default:
			fmt.Println("Pilihan tidak dimengerti")
		}
	}
}


func viewAddTodo() {
	fmt.Println("Menambahkan Todolist")

	todo := input("Todo ( x Jika Batal)")

	if todo == "x" {
		//batal
	} else {
		addTodo(todo)
	}
}


func viewRemoveTodo() {
	fmt.Println("Menghapus Todolist")
	numberStr := input("Nomor yang dihapus (x Jika Batal)")
	number, err := strconv.Atoi(strings.TrimSpace(numberStr))
	if err != nil {
		fmt.Println("Invalid number")
		return
	}
	success := removeTodo(number)
	if numberStr == "x" {
		//batal
	}else if success {
		fmt.Println("Berhasil menghapus Todo")
	} else {
		fmt.Println("Gagal menghapus Todo")
	}
}