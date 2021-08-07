package service

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"time"
	. "todo-list-sederhana/model"
)

type Lists []Todo

var TodoList Lists

func (l *Lists) Filter(id int, author string, todo string, done string, created_at string) (Lists, string, error) {
	var updatedTodo Lists
	for _, i := range *l {
		if i.Id != id {
			updatedTodo = append(updatedTodo, i)
		} else {
			updatedTodo = append(updatedTodo, Todo{Id: id, Author: author, Todo: todo, Done: done, Created_At: created_at})
		}
	}
	return updatedTodo, "Berhasil", nil
}

func (l *Lists) DeleteTodo(id int) *Lists {
	var NewTodoList Lists
	for _, i := range *l {
		if i.Id != id {
			fmt.Println(i.Id, id)
			NewTodoList = append(NewTodoList, i)
		}
	}

	return &NewTodoList
}

func InputTodo(id int, author string, todo string, done string, created_at string) (Lists, error) {
	if len(author) == 0 || len(todo) == 0 || len(done) == 0 || len(created_at) == 0 {
		return nil, errors.New("Data Kamu Tidak Valid")
	}

	_, exist := os.Stat("todo.json")

	if os.IsNotExist(exist) == true {
		TodoList = append(TodoList, Todo{Id: id, Author: author, Todo: todo, Done: done, Created_At: created_at})
		file, err := os.Create("todo.json")
		if err != nil {
			return TodoList, err
		}
		defer file.Close()
		json, _ := json.Marshal(TodoList)
		file.Write(json)
		file.Sync()
	} else {
		file, _ := os.OpenFile("todo.json", os.O_WRONLY, 0644)
		reader, _ := ioutil.ReadFile("todo.json")
		defer file.Close()

		json.Unmarshal([]byte(reader), &TodoList)
		TodoList = append(TodoList, Todo{Id: id, Author: author, Todo: todo, Done: done, Created_At: created_at})
		json, err := json.Marshal(TodoList)
		if err != nil {
			fmt.Println(err.Error(), "Atas")
		}
		_, err = file.Write(json)
		if err != nil {
			fmt.Println(err.Error(), "bawah")
		}
		file.Sync()
	}
	return TodoList, nil
}

func UpdateTodo(id int, author string, todo string, done string, created_at string) (string, error) {
	_, err := os.Stat("todo.json")

	if os.IsNotExist(err) == true {
		return "Gagal", errors.New("Kamu Tidak Memiliki Todo List")
	} else {
		reader, err := ioutil.ReadFile("todo.json")
		if err != nil {
			return "Gagal", err
		}

		openFile, err := os.OpenFile("todo.json", os.O_WRONLY, 0644)
		if err != nil {
			return "Gagal", err
		}
		defer openFile.Close()

		json.Unmarshal([]byte(reader), &TodoList)
		updatedToko, _, err := TodoList.Filter(id, author, todo, done, created_at)
		if err != nil {
			return err.Error(), err
		}

		result, err := json.Marshal(updatedToko)
		if err != nil {
			return err.Error(), err
		}
		os.Remove("todo.json")
		openFile, _ = os.Create("todo.json")
		defer openFile.Close()
		time.Sleep(time.Second * 5)
		openFile.Write(result)
		openFile.Sync()
	}
	return "Berhasil", nil
}

func DeleteTodo(id int) (string, error) {
	openFile, err := os.OpenFile("todo.json", os.O_WRONLY, 0644)
	defer openFile.Close()
	if err != nil {
		return "Gagal", err
	}

	reader, err := ioutil.ReadFile("todo.json")
	if err != nil {
		return "Gagal", err
	}

	err = json.Unmarshal([]byte(reader), &TodoList)
	if err != nil {
		return "Gagal", err
	}

	new := *TodoList.DeleteTodo(id)
	os.Remove("todo.json")
	newFile, err := os.Create("todo.json")
	if err != nil {
		return "Gagal", err
	}
	if len(new) > 0 {
		json, _ := json.Marshal(new)
		newFile.Write(json)
	} else {
		newFile.Write([]byte(""))
	}
	newFile.Sync()

	return "Berhasil", nil
}

func ReadTodo() (string, error) {
	reader, err := ioutil.ReadFile("todo.json")
	if err != nil {
		return "Gagal", err
	}
	fmt.Println(string(reader))
	return "Berhasil", nil
}
