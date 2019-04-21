package mapreduce

import (
	"time"
)

// schedule starts and waits for all tasks in the given phase (Map or Reduce).
func (mr *Master) schedule(phase jobPhase) {
	var ntasks int
	var nios int // number of inputs (for reduce) or outputs (for map)
	switch phase {
	case mapPhase:
		ntasks = len(mr.files)
		nios = mr.nReduce
	case reducePhase:
		ntasks = mr.nReduce
		nios = len(mr.files)
	}

	debug("Schedule: %v %v tasks (%d I/Os)\n", ntasks, phase, nios)

	// All ntasks tasks have to be scheduled on workers, and only once all of
	// them have been completed successfully should the function return.
	// Remember that workers may fail, and that any given worker may finish
	// multiple tasks.
	//
	// TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO
	//

	stats := make([]bool, ntasks)
	for {
		count := ntasks
		for i := 0; i < ntasks; i++ {
			if !stats[i] {
				args := new(DoTaskArgs)
				args.JobName = mr.jobName
				args.Phase = phase
				args.TaskNumber = i
				args.NumOtherPhase = nios
				if phase == mapPhase {
					args.File = mr.files[i]
				}
				worker := <-mr.registerChannel
				go func(slot int, worker string) {
					ok := call(worker, "Worker.DoTask", &args, new(struct{}))
					if ok {
						stats[slot] = true
						mr.registerChannel <- worker
					} else {
						debug("Schedule: worker %v error when doing task %v\n", worker, args)
					}
					// mr.registerChannel <- worker
				}(i, worker)
			} else {
				count--
			}
		}
		if count == 0 {
			break
		}
		time.Sleep(500 * time.Millisecond)
	}
	debug("Schedule: %v phase done\n", phase)
}
