package repositories

import (
	"context"
	"fmt"
	"strconv"

	models "github.com/EmmanSkout/TaskManager/models"
	redis "github.com/redis/go-redis/v9"
)

var client *redis.Client
var ctx = context.Background()

func InitializeClient() {
	client = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

}

func ModifyTask(task models.Task) {
	client.HSet(ctx, "Task-"+strconv.Itoa(task.ID), "Name", task.Name)
	client.HSet(ctx, "Task-"+strconv.Itoa(task.ID), "Description", task.Description)
	client.HSet(ctx, "Task-"+strconv.Itoa(task.ID), "Complete", strconv.FormatBool(task.Complete))
	client.HSet(ctx, "Task-"+strconv.Itoa(task.ID), "Date", task.Date)
}

func AddTask(task models.Task) error {
	complete := strconv.FormatBool(task.Complete)
	taskMap := map[string]string{"ID": fmt.Sprint(task.ID), "Name": task.Name, "Description": task.Description, "Date": task.Date, "Complete": complete}
	for k, v := range taskMap {
		err := client.HSet(ctx, "Task-"+fmt.Sprint(task.ID), k, v).Err()
		if err != nil {
			panic(err)
		}
	}
	return nil
}

func GetTasks() []models.Task {
	iter := client.Scan(ctx, 0, "Task-*", 0).Iterator()
	var tasks []models.Task
	for iter.Next(ctx) {
		key := iter.Val()
		value, err := client.HGetAll(ctx, key).Result()
		if err != nil {
			continue
		}
		taskId, _ := strconv.Atoi(value["ID"])
		taskComplete, _ := strconv.ParseBool(value["Complete"])
		task := models.Task{
			ID:          taskId,
			Name:        value["Name"],
			Description: value["Description"],
			Date:        value["Date"],
			Complete:    taskComplete,
		}
		tasks = append(tasks, task)
	}
	return tasks
}
