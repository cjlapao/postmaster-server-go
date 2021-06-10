package nginx

import (
	"fmt"
)

type Event struct {
	WorkerConnections int
}

type Http struct {
	Include string
	Index   string
}

func (e *Event) ToString() string {
	result := "events {\n"
	if e.WorkerConnections > 0 {
		result += fmt.Sprintf("  worker_connections %v;\n", e.WorkerConnections)
	}
	result += "}\n"
	return result
}
