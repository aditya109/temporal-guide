package greeting

import (
	"errors"
	"log"
	"time"

	"github.com/aditya109/temporal-guide/1-writing-a-workflow-definition/app/outsource"
	"go.temporal.io/sdk/temporal"
	"go.temporal.io/sdk/workflow"
)

// OLD - Business Logic
// func GreetSomeone(name string) string {
// 	return fmt.Sprintf("Hello %s !", name)
// }

// This is a workflow definition and type = GreetSomeone
func GreetSomeone(ctx workflow.Context, name string) (string, error) {
	log.Println("Workflow: GreetSomeone got triggered.")
	if name == "Donna" {
		return "Hello " + name + " !", nil
	}
	return "", errors.New("workflow exited")
}

func GreetSomeoneInSpanish2(ctx workflow.Context, name string) (string, error) {
	retryPolicy := &temporal.RetryPolicy{
		InitialInterval:    0,
		BackoffCoefficient: 1,
		MaximumInterval:    0 * time.Second,
		MaximumAttempts:    1,
	}
	options := workflow.ActivityOptions{
		StartToCloseTimeout: time.Second * 5,
		RetryPolicy:         retryPolicy,
	}
	ctx = workflow.WithActivityOptions(ctx, options)

	var spanishGreeting string
	err := workflow.ExecuteActivity(ctx, outsource.GreetInSpanish2, name).Get(ctx, &spanishGreeting)
	if err != nil {
		return "", err
	}

	return spanishGreeting, nil
}
func GreetSomeoneInSpanish(ctx workflow.Context, name string) (string, error) {
	retryPolicy := &temporal.RetryPolicy{
		InitialInterval:    1 * time.Second,
		BackoffCoefficient: 2.0,
		MaximumInterval:    time.Second * 3,
		MaximumAttempts:    2,
	}
	options := workflow.ActivityOptions{
		StartToCloseTimeout: time.Second * 5,
		RetryPolicy:         retryPolicy,
	}
	ctx = workflow.WithActivityOptions(ctx, options)

	var spanishGreeting string
	err := workflow.ExecuteActivity(ctx, outsource.GreetInSpanish, name).Get(ctx, &spanishGreeting)
	if err != nil {
		return "", err
	}

	return spanishGreeting, nil
}
