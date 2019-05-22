package worker

import (
	"fmt"
	"time"

	"github.com/sungora/app/workflow"
)

func init() {
	workflow.TaskAddCron(&Four{})
}

// Пример задачи работающей по расписанию
type Four struct {
}

// Manager режим работы задачи
func (task *Four) Manager() workflow.Manager {
	return workflow.Manager{
		Name:      "SampleTaskFour",
		IsExecute: true,
		Minute:    "*/4",
		Hour:      "*",
		Day:       "*",
		Month:     "*",
		Week:      "*",
	}
}

// Execute выполняемая задача
func (task *Four) Execute() {
	fmt.Println(time.Now().Format(time.RFC3339) + " execute: " + task.Manager().Name)
}
