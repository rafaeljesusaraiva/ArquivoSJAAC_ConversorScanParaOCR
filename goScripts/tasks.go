package goScripts

// TaskProgress represents the progress of a single task
type TaskProgress struct {
	Name     string  // name of the task
	Progress float64 // progress of the task (0-100)
}

// OverallProgress represents the overall progress of all tasks
type OverallProgress struct {
	Tasks      []TaskProgress // progress of each task
	TotalTasks int            // total number of tasks to be completed
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

// CalculateMainProgress calculates the main progress of all tasks
func (op *OverallProgress) CalculateMainProgress() float64 {
	totalProgress := 0.0
	for _, task := range op.Tasks {
		totalProgress += task.Progress
	}
	// Add 0 progress for remaining tasks
	totalProgress += 0 * float64(op.TotalTasks-len(op.Tasks))
	return (totalProgress / float64(op.TotalTasks))
}

// Reset resets the overall progress and all tasks to their initial state
func (op *OverallProgress) Reset() {
	op.Tasks = []TaskProgress{}
}
