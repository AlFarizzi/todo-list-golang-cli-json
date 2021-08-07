package service

import (
	"fmt"
	"math/rand"
	"testing"
	"time"

	// . "todo-list-sederhana/model"

	"github.com/stretchr/testify/assert"
)

func TestInputTodoSuccess(t *testing.T) {
	fmt.Println(rand.Intn(1000000))
	result, err := InputTodo(rand.Intn(100000), "Fariz", "Makan", "Belum Selesai", time.Now().Format(time.RFC3339))
	assert.IsType(t, TodoList, result, "Nilai Tidak Sesuai")
	assert.Nil(t, err, "Ada Error -> Test Gagal")
}

func TestInputTodoFailed(t *testing.T) {
	_, err := InputTodo(rand.Intn(100000), "", "", "Belum Selesai", time.Now().Format(time.RFC3339))
	assert.NotNil(t, err, "Tidak Ada Error -> Test Gagal")
}

func TestUpdateTodoSuccess(t *testing.T) {
	result, err := UpdateTodo(27887, "Fariz", "Makan", "Selesai", time.Now().Format(time.RFC3339))
	assert.Equal(t, "Berhasil", result, "Gagal Update")
	assert.Nil(t, err, "Ada Error")
}

func TestDeleteTodoSuccess(t *testing.T) {
	result, err := DeleteTodo(27887)
	assert.Equal(t, "Berhasil", result, "Tidak Berhasil")
	assert.Nil(t, err, "Tidak Nil")
}

func TestReadTodoSuccess(t *testing.T) {
	result, err := ReadTodo()
	assert.Equal(t, "Berhasil", result, "Tidak Berhasil")
	assert.Nil(t, err, "Tidak Nil")
}

// func TestUpdateTodoFailed(t *testing.T) {
// 	result, err := UpdateTodo(1, "Fariz", "Makan", "Belum Selesai", time.Now().Format(time.RFC3339))
// 	assert.Equal(t, "Gagal", result, "Berhasil Update")
// 	assert.NotNil(t, err, "Tidak Ada Error")
// }
