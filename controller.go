package meritop

import "github.com/coreos/go-etcd/etcd"

// This is the controller of a job.
// A job needs controller to setup etcd data layout, request
// cluster containers, etc. to setup framework to run.
type controller struct {
	name       string
	etcdclient *etcd.Client
	topology   Topology
}

func (c *controller) initEtcdLayout() (err error) {
	// initiate etcd data layout
	for i := uint64(0); i < c.topology.NumberOfTasks(); i++ {
		key := MakeTaskMasterPath(c.name, i)
		c.etcdclient.Create(key, "", 0)
	}
	return
}

func (c *controller) destroyEtcdLayout() (err error) {
	c.etcdclient.Delete("/", true)
	return
}
