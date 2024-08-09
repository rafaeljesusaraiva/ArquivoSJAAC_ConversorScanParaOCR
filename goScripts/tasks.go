package goScripts

// TaskProgress represents the progress of a single task
type TaskProgress struct {
	Name     string  // name of the task
	Progress float64 // progress of the task (0-100)
}

// OverallProgress represents the overall progress of all tasks
type OverallProgress struct {
	Tasks         []TaskProgress // progress of each task
	TotalProgress float64        // overall progress (0-100)
	CurrentTask   int            // index of the current task
}

// AddTask adds a new task to the overall progress
func (op *OverallProgress) AddTask(name string) (id int) {
	newTask := TaskProgress{
		Name:     name,
		Progress: 0,
	}
	op.Tasks = append(op.Tasks, newTask)

	return len(op.Tasks) - 1
}

// UpdateTaskProgress updates the progress of a task
func (op *OverallProgress) UpdateTaskProgress(taskIndex int, progress float64) {
	if taskIndex < 0 || taskIndex >= len(op.Tasks) {
		panic("task index out of range")
	}
	op.Tasks[taskIndex].Progress = progress
}

func (op *OverallProgress) CalculateTasksCompleted() int {
	tasksCompleted := 0
	for _, task := range op.Tasks {
		if task.Progress == 100 {
			tasksCompleted++
		}
	}
	return tasksCompleted
}

// Reset resets the overall progress and all tasks to their initial state
func (op *OverallProgress) Reset() {
	op.Tasks = []TaskProgress{}
	op.TotalProgress = 0
	op.CurrentTask = 0
}
