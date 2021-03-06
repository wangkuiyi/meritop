//The tree structure is basically assume that all the task forms a tree.
//Also the tree structure stays the same between epochs.

package meritop

type TreeTopology struct {
	fanout, numberOfTasks uint64
	taskID                uint64
	parents, children     []uint64
}

func (t *TreeTopology) SetTaskID(taskID uint64) {
	t.taskID = taskID
	// Not the most efficient way to create parents and children, but
	// since this is not on critical path, we are ok.
	t.parents = make([]uint64, 0, 1)
	t.children = make([]uint64, 0, t.fanout)

	for index := uint64(1); index <= t.numberOfTasks; index++ {
		parentID := (index - 1) / t.fanout
		if index == taskID {
			t.parents = append(t.parents, parentID)
			if len(t.parents) > 1 {
				panic("unexpcted number of partents for a tree topology")
			}
		}
		if parentID == taskID {
			t.children = append(t.children, index)
		}
	}
}

func (t TreeTopology) GetParents(epochID uint64) []uint64 {
	return t.parents
}

func (t TreeTopology) GetChildren(epochID uint64) []uint64 {
	return t.children
}

func (t TreeTopology) NumberOfTasks() uint64 {
	return t.numberOfTasks
}

// Creates a new tree topology with given fanout and number of tasks.
// This will be called each
func NewTreeTopology(fanout, numberOfTasks uint64) *TreeTopology {
	m := &TreeTopology{
		fanout:        fanout,
		numberOfTasks: numberOfTasks,
	}
	return m
}
