package models

import "context"

var ctx = context.TODO()

func GetTasks() ([]string, error) {
	return client.LRange(ctx, "tasks", 0, 10).Result()
}

func PostTask(task string) error {
	return client.LPush(ctx, "tasks", task).Err()
}
