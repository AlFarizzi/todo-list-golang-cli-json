package main

import (
	"flag"
	"fmt"
	"math/rand"
	"time"
	. "todo-list-sederhana/service"
)

func main() {
	mode := flag.String("mode", "tambah", "Masukan Mode")
	author := flag.String("author", "author", "Masukan Nama Author")
	todo := flag.String("todo", "Wake Up", "Masukan TODO List Kamu")
	done := flag.String("done", "Belum Selesai", "Selesai ?")
	todoId := flag.Int("id", 0, "Masukan id Todo")
	created_at := time.Now().Format(time.RFC3339)
	flag.Parse()

	if *mode == "tambah" {
		rand.Seed(time.Now().UnixNano())
		_, err := InputTodo(rand.Intn(1000), *author, *todo, *done, created_at)
		if err != nil {
			fmt.Println(err.Error())
		} else {
			fmt.Println("Berhasil")
		}
	} else if *mode == "update" {
		result, err := UpdateTodo(*todoId, *author, *todo, *done, created_at)
		if err != nil {
			fmt.Println(err.Error())
		} else {
			fmt.Println(result)
		}
	} else if *mode == "hapus" {
		result, err := DeleteTodo(*todoId)
		if err != nil {
			fmt.Println(err.Error())
		} else {
			fmt.Println(result)
		}
	} else if *mode == "baca" {
		ReadTodo()
	}

}
