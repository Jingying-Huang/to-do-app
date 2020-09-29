package models

import "context"

func GetTasks() ([]string, error) {
	ctx := context.TODO()
	return client.LRange(ctx, "tasks", 0, 10).Result()
}

func PostTask(task string) error {
	ctx := context.TODO()
	return client.LPush(ctx, "tasks", task).Err()
}
