# temporal demo

```bash
# start temporal using docker-compose up, you can view it on http://localhost:8080/

# to start worker
go run worker/worker.go

# example workflow executions

tctl workflow start --workflow_type GreetSomeoneInSpanish --taskqueue greeting-tasks --workflow_id my-second-workflow --input '"2"'

# executing a workflow with an activity which fails 
tctl workflow start --workflow_type GreetSomeoneInSpanish2 --taskqueue greeting-tasks --workflow_id my-second-workflow --input '"2"'

```