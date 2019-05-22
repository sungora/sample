package worker

import (
	"fmt"
	"time"

	"github.com/sungora/app/workflow"
)

func init() {
	workflow.TaskAddCron(&One{})
}

// Пример задачи работающей по расписанию
type One struct {
}

// Manager режим работы задачи
func (task *One) Manager() workflow.Manager {
	return workflow.Manager{
		Name:      "SampleTaskOne",
		IsExecute: true,
		Minute:    "*",
		Hour:      "*",
		Day:       "*",
		Month:     "*",
		Week:      "*",
	}
}

// Execute выполняемая задача
func (task *One) Execute() {
	fmt.Println(time.Now().Format(time.RFC3339) + " execute: " + task.Manager().Name)
}
