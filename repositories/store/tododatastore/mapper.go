package tododatastore

import "todoingo/models"

func mapTodoData(data *TodoData) *models.TodoData {

	if data == nil {
		return nil
	}

	return &models.TodoData{
		Id:          data.Id.Hex(),
		Task:        data.Task,
		IsCompleted: data.IsCompleted,
		UserId:      data.UserId.Hex(),
		IsDeleted:   data.IsDeleted,
	}
}

func mapTodosData(data []TodoData) []models.TodoData {
	todos := make([]models.TodoData, len(data))

	for index, _ := range data {
		todos[index] = *mapTodoData(&data[index])
	}

	return todos
}
