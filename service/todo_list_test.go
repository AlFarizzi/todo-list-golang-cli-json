package service

import (
	"fmt"
	"math/rand"
	"testing"
	"time"

	// . "todo-list-sederhana/model"

	"github.com/stretchr/testify/assert"
)

func BenchmarkInputTodoSuccess(b *testing.B) {
	data := []struct {
		nama string
		todo string
		done string
	}{
		{"Fariz", "Makan", "Belum Selesai"},
		{"Fariz", "Lorem Ipsum is simply dummy text of the printing and typesetting industry. Lorem Ipsum has been the industry's standard dummy text ever since the 1500s, when an unknown printer took a galley of type and scrambled it to make a type specimen book. It has survived not only five centuries, but also the leap into electronic typesetting, remaining essentially unchanged. It was popularised in the 1960s with the release of Letraset sheets containing Lorem Ipsum passages, and more recently with desktop publishing software like Aldus PageMaker including versions of Lorem Ipsum", "Belum Selesai"},
		{"Fariz", "Lorem Ipsum is simply dummy text of the printing and typesetting industry. Lorem Ipsum has been the industry's standard dummy text ever since the 1500s, when an unknown printer took a galley of type and scrambled it to make a type specimen book. It has survived not only five centuries, but also the leap into electronic typesetting, remaining essentially unchanged. It was popularised in the 1960s with the release of Letraset sheets containing Lorem Ipsum passages, and more recently with desktop publishing software like Aldus PageMaker including versions of Lorem Ipsum", "Belum Selesai"},
	}

	for _, d := range data {
		b.Run(d.nama, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				InputTodo(rand.Intn(100000), d.nama, d.todo, d.done, time.Now().Format(time.RFC3339))
			}
		})
	}
}

func BenchmarkTestUpdate(b *testing.B) {
	data := []struct {
		id   int
		nama string
		todo string
		done string
	}{
		{31515, "Fariz", "Makan", "Selesai"},
		{81351, "Fariz", "Lorem Ipsum is simply dummy text of the printing and typesetting industry. Lorem Ipsum has been the industry's standard dummy text ever since the 1500s, when an unknown printer took a galley of type and scrambled it to make a type specimen book. It has survived not only five centuries, but also the leap into electronic typesetting, remaining essentially unchanged. It was popularised in the 1960s with the release of Letraset sheets containing Lorem Ipsum passages, and more recently with desktop publishing software like Aldus PageMaker including versions of Lorem Ipsum", "Selesai"},
		{61598, "Fariz", "Lorem Ipsum is simply dummy text of the printing and typesetting industry. Lorem Ipsum has been the industry's standard dummy text ever since the 1500s, when an unknown printer took a galley of type and scrambled it to make a type specimen book. It has survived not only five centuries, but also the leap into electronic typesetting, remaining essentially unchanged. It was popularised in the 1960s with the release of Letraset sheets containing Lorem Ipsum passages, and more recently with desktop publishing software like Aldus PageMaker including versions of Lorem Ipsum", "Selesai"},
	}

	for _, d := range data {
		b.Run(d.nama, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				UpdateTodo(d.id, d.nama, d.todo, d.done, time.Now().Format(time.RFC3339))
			}
		})
	}

}

func BenchmarkDeleteTodo(b *testing.B) {
	data := []int{98081, 27887, 31847}
	for _, d := range data {
		b.Run("Hai", func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				DeleteTodo(d)
			}
		})
	}
}

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
