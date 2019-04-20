package worker

import (
	"fmt"
	"time"

	"github.com/sungora/app/workflow"
)

func init() {
	workflow.TaskAddCron(&Two{})
}

// Пример задачи работающей по расписанию
type Two struct {
}

// Manager режим работы задачи
func (task *Two) Manager() workflow.Manager {
	return workflow.Manager{
		Name:      "SampleTaskTwo",
		IsExecute: true,
		Minute:    "*/2",
		Hour:      "*",
		Day:       "*",
		Month:     "*",
		Week:      "*",
	}
}

// Execute выполняемая задача
func (task *Two) Execute() {
	fmt.Println(time.Now().Format(time.RFC3339) + " execute: " + task.Manager().Name)
}
