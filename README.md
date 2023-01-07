# temporal demo

1. Start temporal 
    ```bash
    cd temporal-setup
    docker compose up #you can view it on http://localhost:8080/
    ```

2. Start worker
```
cd 1-writing-a-workflow-definition/app/worker
go run worker.go
```

3. Execute workflows

    - Example workflow executions which are successful
        ```bash
        tctl workflow start --workflow_type GreetSomeoneInSpanish --taskqueue greeting-tasks --workflow_id my-second-workflow --input '"2"'
        ```
    - Example executing a workflow with an activity which fails 
        ```bash
        tctl workflow start --workflow_type GreetSomeoneInSpanish2 --taskqueue greeting-tasks --workflow_id my-second-workflow --input '"2"'
        ```