package main

import (
	"log"

	"github.com/aditya109/temporal-guide/sample-app/greeting"
	"github.com/aditya109/temporal-guide/sample-app/outsource"
	"go.temporal.io/sdk/client"
	"go.temporal.io/sdk/worker"
)

// Worker logic

func main() {
	// temporal client -> which communicates with the Temporal Cluster
	c, err := client.Dial(client.Options{})
	if err != nil {
		log.Fatalln("Unable to create client", err)
	}
	defer c.Close()

	// name of task queue -> which is maintained by the Temporal Server and polled by the Worker. (greeting-tasks)
	w := worker.New(c, "greeting-tasks", worker.Options{})

	// FQN of Workflow Defintion function
	w.RegisterWorkflow(greeting.GreetSomeone)
	w.RegisterActivity(outsource.GreetInSpanish)
	w.RegisterActivity(outsource.GreetInSpanish2)
	w.RegisterWorkflow(greeting.GreetSomeoneInSpanish)
	w.RegisterWorkflow(greeting.GreetSomeoneInSpanish2)

	err = w.Run(worker.InterruptCh())
	if err != nil {
		log.Fatalln("Unable to start worker", err)
	}
}
