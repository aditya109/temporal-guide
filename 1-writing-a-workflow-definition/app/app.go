package main

import (
	"context"
	"log"
	"os"

	"github.com/aditya109/temporal-guide/1-writing-a-workflow-definition/app/outsource"

	"go.temporal.io/sdk/client"
)

func main() {
	c, err := client.Dial(client.Options{})
	if err != nil {
		log.Fatalln("Unable to create client", err)
	}

	options := client.StartWorkflowOptions{
		ID:        "my-second-workflow",
		TaskQueue: "greeting-tasks",
	}

	we, err := c.ExecuteWorkflow(context.TODO(), options, outsource.GreetInSpanish, os.Args[1])
	if err != nil {
		log.Fatalln("Unable to execute workflow", err)
	}
	log.Println("Started workflow", "WorkflowID", we.GetID(), "RunID", we.GetRunID())

	var result string
	err = we.Get(context.Background(), &result)
	if err != nil {
		log.Fatalln("Unable get workflow result", err)
	}
	log.Println("Workflow result:", result)
	defer c.Close()
}
