# temporal-guide

## Concepts

### What is Temporal ?

Temporal is a platform that guarantees the durable execution of your application code. 

- It allows you to develop as if failures don't even exist. 
- Your application will run reliably even if it encounters problems, such as network outages or server crashes, which would be catastrophic for a typical application. 
- Temporal handles these types of problems, allowing you to focus on the business logic, instead of writing application code to detect and recover from failures.

#### Workflows

Temporal applications are built using an abstraction called Workflows. You'll develop those Workflows by writing code in a programming language. The code you write is the same code that will be executed at runtime. 

Temporal Workflows are resilient. They can run and keeping running for years even if the underlying infrastructure fails.

If the application crashes, Temporal will automatically recreate its pre-failure state so it can continue where it left off.

**Definition**: Conceptually, a workflow defines a sequence of steps. With Temporal, those steps are defined by writing code, known as a *Workflow Definition*, and are carried out by running that code, which results in a *Workflow Execution*.

Example 1: Expense Report Workflow Example

![Expense Report Workflow Example](https://cdn.talentlms.com/temporal/1664480399_expense-report-workflow-diagram.png?Policy=eyJTdGF0ZW1lbnQiOlt7IlJlc291cmNlIjoiaHR0cHM6XC9cL2Nkbi50YWxlbnRsbXMuY29tXC90ZW1wb3JhbFwvMTY2NDQ4MDM5OV9leHBlbnNlLXJlcG9ydC13b3JrZmxvdy1kaWFncmFtLnBuZyIsIkNvbmRpdGlvbiI6eyJEYXRlTGVzc1RoYW4iOnsiQVdTOkVwb2NoVGltZSI6MTY3MzA0OTYwMH19fV19&Signature=F-EMX%2FW4kvwI1OagV0RaZTmnEw%2F7O25eufUmGYqdE9FGT9Ndy-ANAT07AX50UihqJGm9EzH9gjP8lPAFg1lRHxTd7dhHQAd6uCvH-RIG4GrCJZE1-cJC1ZbkKipuhshoORSHySi3GwYOIP66zD21kNxa02C6urIm3%2FaqpqpuUc5u5nxr0Swx%2Fkjji51XmrlO5lWzVJkF482bcqp5jEZGJb4Y6JVoVk3T9RDCmNGOumvmIHX4UZv2A%2Fmqx1Hj97kbE1eODKD--7pSiO8h9ohKnJNvZV9pp6qjcKz%2FqCh0wDBldkG%2FiCktEkp6%2F9s3Bipqimu5sr44ZLkA3w3jLHhBZA__&Key-Pair-Id=APKAJDCWVQTW4P3KI3XA)

<u>POI</u>

- The above is an example of a potentially long-running process.
- There are conditional logic.
- There are cycles present in the workflows.
- There are multiple independent touch-points.
- There are also involvement of external systems. 

Example 2 : Money Transfer

![](https://cdn.talentlms.com/temporal/1664480380_money-transfer.png?Policy=eyJTdGF0ZW1lbnQiOlt7IlJlc291cmNlIjoiaHR0cHM6XC9cL2Nkbi50YWxlbnRsbXMuY29tXC90ZW1wb3JhbFwvMTY2NDQ4MDM4MF9tb25leS10cmFuc2Zlci5wbmciLCJDb25kaXRpb24iOnsiRGF0ZUxlc3NUaGFuIjp7IkFXUzpFcG9jaFRpbWUiOjE2NzMwNDk2MDB9fX1dfQ__&Signature=lqGMiED3ia%2FdzZ5msXJrzfKKFqPecUBV7bnxE3FCjvIZbtHg6l4GQlkGXe4xWmOhlbyzlOJz%2Fk3zmW-b09lEXs8D%2FdJ-7uOiSTw9onRInHi3lBrK7yRHB9L%2FPi35aYkoQRctRd74AuTUaImL1trJEoWhe7h-Mrgzz8ijJxEJ8X7e%2FQskS-B06S4wbkrDdK91fBNYExGPkpWHc87yewAJSmO6A8zXGFSkf27akBUk2G%2FR9iR4nX9uIPfpNWE8QijdV--h75FVKKUD0LoEpC57oaWJj08wvD7EOag-Y4dyZ64NCrOvDAVxgOJoCssQvYs%2FOTa88xKJiSvrou29L37PCQ__&Key-Pair-Id=APKAJDCWVQTW4P3KI3XA)

The first step is to withdraw money from the employer’s bank account and the second is to deposit the same amount into the employee’s bank account. There are two important constraints for doing this correctly. First, you must execute *both* the withdrawal and the deposit. Second, you must execute each of them *exactly once*.

More broadly, the reimbursement is just a transfer of money between two accounts. There are plenty of other use cases for this same workflow. In fact, millions of people every day depend on it when they use services like Square, Stripe, Western Union, PayPal, Swish, or Apple Pay.

This workflow would typically involve multiple accounts accessed through some type of remote procedure calls, making it a distributed system. As with any distributed system, it could fail for many reasons, including server failure or a network outage. If this workflow wasn't built on Temporal, the consequences could be catastrophic. If it happened to fail in the middle of this function, between the withdrawal and deposit steps, then the account balances would be incorrect. Worse yet, the current state would be lost, so restarting the application would repeat the withdrawal again, rather than resuming with the deposit. As a developer, it's your responsibility to detect and mitigate those failures, but if you're using Temporal Workflows, you can let the platform handle these types of failures.

##### Temporal Architecture

###### Temporal Server

> *If someone mentions Temporal without additional context, they're probably referring to the Temporal Platform.*

![](https://cdn.talentlms.com/temporal/1664480472_temporal-server-diagram.png?Policy=eyJTdGF0ZW1lbnQiOlt7IlJlc291cmNlIjoiaHR0cHM6XC9cL2Nkbi50YWxlbnRsbXMuY29tXC90ZW1wb3JhbFwvMTY2NDQ4MDQ3Ml90ZW1wb3JhbC1zZXJ2ZXItZGlhZ3JhbS5wbmciLCJDb25kaXRpb24iOnsiRGF0ZUxlc3NUaGFuIjp7IkFXUzpFcG9jaFRpbWUiOjE2NzMwNDk2MDB9fX1dfQ__&Signature=dwLRwk0HPYIGKEITMIKTl%2F52FM5fAWJ7JJ8ah--NUmtOVHNMcmm0r8RK0iw1TepOs4S3xZcJbse1qzTYUQ4FWJpeP8ZNRyOLKhSGRmUbbknby7nrXZ-7xFDGWZ37hwuC3WvM0NZ09xaYKu1uRbOmT0ENJmKy2fW8QwWMWvJua2ILnKGaOM9mtg1PpKhwdgDfmLA0AtmY8UhStVbwvhjWzoXFc4K0y-K5VwIZ7sY97Qw8BU9SXF2varF5aygd8xL2ewUlrm9XbS7FlMyBjeA96gNq7rXiB4PGYxNjB3K3s3NwPP8afwBXv8-Msm9W5Vq8Wrxq9RmTbjAz6y0z8jty8w__&Key-Pair-Id=APKAJDCWVQTW4P3KI3XA)

The **Temporal Server** consists of a front-end service, plus several back-end services that work together to manage the execution of your application code. All of those services are horizontally scalable and a production environment will typically run multiple instances of each, deployed across multiple machines, to increase performance and availability.

**Note:** The front-end service that is part of the Temporal Server acts as an API gateway. In other words, it is a front-end for clients, not end users (end users will interact with the CLI or Web UI).

###### Temporal Cluster

The Temporal Server is an essential part of the overall system, but requires additional components for operation. The complete system is known as the **Temporal Cluster**, which is a deployment of the Temporal Server software on some number of machines, plus the additional components used with it.

![](https://cdn.talentlms.com/temporal/1665506196_temporal-cluster-diagram.png?Policy=eyJTdGF0ZW1lbnQiOlt7IlJlc291cmNlIjoiaHR0cHM6XC9cL2Nkbi50YWxlbnRsbXMuY29tXC90ZW1wb3JhbFwvMTY2NTUwNjE5Nl90ZW1wb3JhbC1jbHVzdGVyLWRpYWdyYW0ucG5nIiwiQ29uZGl0aW9uIjp7IkRhdGVMZXNzVGhhbiI6eyJBV1M6RXBvY2hUaW1lIjoxNjczMDQ5NjAwfX19XX0_&Signature=Cqw8G4wzXCLIS-8yEjqUZbsmJNqSGHB7qp-bPu7t-zfiYCGJoHLVfArDxge9-RsjhLL-4HC%2F8e%2FLjFQ2oFzV8qhfhiBioioA39T9sqUgbhlB1WypRAT2QNrfH%2F6U5YpI78RFDDacO3eRk6h6aDgWXUp4q63st6PPE1TEP5xOTSBmExl4RzG18Agbg72-Leid13T7daerXy884oxRhyYbS8aAGQIpybq4DlV9fJlFmZFsZPQ3X8szqR4RO-tIZVHg8Suo8j7RMk%2FiyvHuATeYiKp1hKctLXg7%2Fp%2FrALLAmiNFSWdit4JN0RqWu0LQK%2F3niF8S%2FMK96euOyX4oi4JIUw__&Key-Pair-Id=APKAJDCWVQTW4P3KI3XA)

1. The only *required* component is a database such as `Apache Cassandra`, `PostgreSQL`, or `MySQL`.
2. The Temporal Cluster tracks the current state of every execution of your Workflows, as well as also maintains a history of all Events that occur during their executions, which it uses to reconstruct the current state in case of failure. It persists this and other information, such as details related to durable timers and queues to the database.
3. `Elasticsearch` is an optional component. It is not necessary for basic operation, but adding it will give me advanced searching, sorting and filtering capabilities for information about current and recent Workflow Executions.
4. Two other which are often used with Temporal -
   - `Prometheus` to collect metrics from Temporal
   - `Grafana` is used to create dashboards based on those metrics.

###### Workers

*The Temporal Cluster does not execute your code.* While the platform guarantees the durable execution of your code, it achieves this through orchestration. The execution of your application code is external to the cluster, and in typical deployments, takes place on a separate set of servers, potentially running in a different data centre than the Temporal Cluster.

The entity responsible for executing your code is known as a Worker, and it's common to run Workers on multiple servers, since this increases both the scalability and availability of your application. The Worker, which is part of your application, communicates with the Temporal Cluster to manage the execution of your Workflows.

 ![](https://cdn.talentlms.com/temporal/1664480638_temporal-platform-diagram.png?Policy=eyJTdGF0ZW1lbnQiOlt7IlJlc291cmNlIjoiaHR0cHM6XC9cL2Nkbi50YWxlbnRsbXMuY29tXC90ZW1wb3JhbFwvMTY2NDQ4MDYzOF90ZW1wb3JhbC1wbGF0Zm9ybS1kaWFncmFtLnBuZyIsIkNvbmRpdGlvbiI6eyJEYXRlTGVzc1RoYW4iOnsiQVdTOkVwb2NoVGltZSI6MTY3MzEzNjAwMH19fV19&Signature=msUyHvochU8T%2FnYJ4VqJQywzxn3Yil-FQqi2gNRAMI-exh044FhMIsyoLgX9JTEq3OVPQ5N9XyBEXTBDm-kD5OUP7ZnFwSJosuTg6UggJ4kaX9qDde1mjPx-PuwoELYy4DveCudIO3zGKbSrR9dlUKzCQ34DjpdjpulSkFV8ALeU3bBCx%2FNswVDgxXIzRSrWmPEaVidieJkcIr2HmKxxqmcjXMpAFZ0Xc9rzQ6JyapHNYHA4vn1Ou71XJFF46MtH4o5rcvtyDbhPguqTN8Y3lRD60Sxr6TsHa4kV02wePz6vzSMQvSgAsqrdgloc-L%2FDeEXNb43NFlMH6E6Q9-11Rw__&Key-Pair-Id=APKAJDCWVQTW4P3KI3XA)

- The application will contain the code used to initialise the Worker, the Workflow and other functions that comprise your business logic, and possibly also code used to start or check the status of the Workflow. 
- At runtime, you'll need everything needed to execute the application, which will include any libraries or other dependencies referenced in your code, on each machine where at least one Worker process will run. 
- Temporal uses `gRPC` for communication, so each machine running a Worker will require connectivity to the front-end service on the Temporal cluster.

#### Namespaces

Logical segregation of workflows.

Use `tctl` to create additional namespaces.

### How to set up Temporal Clusters ?

1. **Self-hosted**

   For local development, we can use Docker Compose and [Temporalite](https://github.com/temporalio/temporalite).

   For production, we can orchestrate it using `Kubernetes`.

2. **Temporal Cloud**

   Temporal Cloud uses consumption-based pricing, so you only pay for what you use, and you can see your current and past usage at any time from the web interface.

### Getting started with Temporal Development

#### Writing a Workflow Definition

The code that makes up the workflow is called the Workflow Definition.

Temporal does not impose any rules for naming its Workflow, but that being name of the Temporal Workflow is called Workflow Type. (name=Type=name-of-the-go-function, can be renamed) 

#### Values must be serializable

Temporal maintains information about current and past Workflow Executions, which can be used for historical and empirical debugging later.

In order for Temporal to store the Workflow's input and output, data used in input parameters and return values must be serializable. By default, Temporal can handle null or binary values, as well as any data that can be serialized using Go's support for JSON or Protocol Buffers.  

<span style="color:darkorange">This means that most of the types you'd typically use in a function, such as integers and floating point numbers, boolean values, and strings, are all handled automatically, as are structs composed from these types, but types such as **channels**, **functions**, or **unsafe pointers** are prohibited as either input parameters or return values.</span>

#### Data Confidentiality

Although the input parameters and return values are stored as part of the Event History of your Workflow Executions, you can create a custom Data Converter to encrypt the data as it enters the Temporal Cluster and decrypt it upon exit, thereby maintaining the confidentiality of any sensitive data used as input or output of your applications.

#### Avoid Passing Large Amounts of Data

Because the Event History contains the input and output, which is also sent across the network from the application to the Temporal Cluster, you'll have a better performance if you limit the amount of data sent.

To protect against unexpected failures caused by sending or storing too much data, the Temporal Server imposes various limits beyond which it will emit warnings or errors, depending on the severity.

https://docs.temporal.io/kb/temporal-platform-limits-sheet

#### The Role of a Worker

As mentioned earlier, Workers execute your Workflow code. The Worker itself is provided by the Temporal SDK, but your application will include code to configure and run it. When that code executes, the Worker establishes a persistent connection to the Temporal Cluster and begins polling a Task Queue on the Cluster, seeking work to perform.

#### Worker Initialization Explanation

```go
package main

import (
	"log"

	"github.com/aditya109/temporal-guide/sample-app/greeting"
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

	err = w.Run(worker.InterruptCh())
	if err != nil {
		log.Fatalln("Unable to start worker", err)
	}
}
```



There are typically three things you need in order to configure a Worker:

1. **A Temporal Client**, which used to communicate with the Temporal Cluster. The first line in the `main` function creates a client, while the next few lines check that creating it didn't produce any errors and ensure that it will be closed when no longer needed. The code used to create the client will vary from what's shown here if you're using Temporal Cloud or a self-hosted cluster, as this will also include the address and port number of the frontend service and credentials used for authentication.
2. **The name of a Task Queue**, which is maintained by the Temporal Server and polled by the Worker. In this example, the task queue name is `greeting-tasks`. This value is supplied, along with the client, when creating the Worker.
3. **The fully-qualified name of the Workflow Definition function**, which is used in the call to `RegisterWorkflow`. Every Workflow Definition function must be registered with at least one Worker for execution to proceed, but you may register multiple of these functions with any given Worker.

Once you have finished configuring the Worker, you can call its `Run` function to start it. The Worker will then begin a "long poll" on the specified task queue. If you start the Worker from a terminal using a program like the one shown above, don't be surprised if you see nothing more than a few lines of output. This is the expected behaviour and the program isn't stuck, it's just busy polling the task queue and working on the tasks that it has accepted from the Temporal Cluster.

##### The Lifetime of a Worker

1. The lifetime of the Worker and the duration of a Workflow Execution are unrelated. 
2. The `Run` function used to start this Worker is a blocking function that doesn't stop unless it is terminated or encounters a fatal error. 
   - The Worker's process may last for days, weeks, or longer -
     - If the Workflows it handles are relatively short, then a single Worker might execute thousands or even millions of them during its lifetime. 
     - On the other hand, a Workflow can run for years, while the server where a Worker process is running might be rebooted after a few months by an administrator doing maintenance. 
     - If the Workflow Type was registered with other workers, one or more of them will automatically continue where the original Worker left off. 
     - If there are no other Workers available, then the Workflow Execution will continue where it left off as soon as the original Worker is restarted. In either case, the downtime will not cause Workflow Execution to fail.

### How to start a workflow ?

#### Using tctl to Start a Workflow 

```bash
> tctl workflow start \
    --workflow_type GreetSomeone \
    --taskqueue greeting-tasks \
    --workflow_id my-first-workflow \
    --input '"Donna"'
    
Started Workflow Id: my-first-workflow, run Id: fcead8d0-f907-4968-b5b7-9ed45f5c8c0e


```

> The command also specifies a Workflow ID, which is optional, but recommended. This is a user-defined identifier, which typically has some business meaning, so an expense reporting workflow might have a Workflow ID that identifies the expense report or the employee who submitted it. If omitted, a UUID will be automatically assigned as the Workflow ID.

##### What happened when the command was run ?

```bash
❯ tctl workflow observe --workflow_id my-first-workflow
Progress:
  1, 2023-01-07T16:03:58Z, WorkflowExecutionStarted
  2, 2023-01-07T16:03:58Z, WorkflowTaskScheduled
  3, 2023-01-07T16:03:58Z, WorkflowTaskStarted
  4, 2023-01-07T16:03:58Z, WorkflowTaskCompleted
  5, 2023-01-07T16:03:58Z, WorkflowExecutionCompleted

Result:
  Run Time: 1 seconds
  Status: COMPLETED
  Output: ["Hello Donna !"]
```

##### What happens when the workflow fails ?

```go
❯ tctl workflow observe --workflow_id my-first-workflow
Progress:
  1, 2023-01-07T16:10:07Z, WorkflowExecutionStarted
  2, 2023-01-07T16:10:07Z, WorkflowTaskScheduled
  3, 2023-01-07T16:10:07Z, WorkflowTaskStarted
  4, 2023-01-07T16:10:07Z, WorkflowTaskCompleted
  5, 2023-01-07T16:10:07Z, WorkflowExecutionFailed

Result:
  Run Time: 1 seconds
  Status: FAILED
  Failure: &Failure{Message:Workflow exited,Source:GoSDK,StackTrace:,Cause:nil,FailureType:Failure_ApplicationFailureInfo,}
```

```bash
❯ tctl wf show --workflow_id my-first-workflow
  1  WorkflowExecutionStarted    {WorkflowType:{Name:GreetSomeone},                            
                                  ParentInitiatedEventId:0, TaskQueue:{Name:greeting-tasks,     
                                  Kind:Normal}, Input:["Donna"],                                
                                  WorkflowExecutionTimeout:0s, WorkflowRunTimeout:0s,           
                                  WorkflowTaskTimeout:10s, Initiator:Unspecified,               
                                  OriginalExecutionRunId:6c9bd15a-eede-4eba-af17-7f798df49db8,  
                                  Identity:104592@firestik@,                                    
                                  FirstExecutionRunId:6c9bd15a-eede-4eba-af17-7f798df49db8,     
                                  Attempt:1, FirstWorkflowTaskBackoff:0s,                       
                                  ParentInitiatedEventVersion:0}                                
  2  WorkflowTaskScheduled       {TaskQueue:{Name:greeting-tasks,                              
                                  Kind:Normal},                                                 
                                  StartToCloseTimeout:10s,                                      
                                  Attempt:1}                                                    
  3  WorkflowTaskStarted         {ScheduledEventId:2, Identity:88961@firestik@,                
                                  RequestId:b9638542-162e-49ef-839d-a3541194c079,               
                                  SuggestContinueAsNew:false, HistorySizeBytes:0}               
  4  WorkflowTaskCompleted       {ScheduledEventId:2, StartedEventId:3,                        
                                  Identity:88961@firestik@,                                     
                                  BinaryChecksum:c7e23d7e59fc99ccf1ad2d4fdbad44b3}              
  5  WorkflowExecutionCompleted  {Result:["Hello                                               
                                  Donna !"],                                                    
                                  WorkflowTaskCompletedEventId:4}  
```

#### Using Application Code

```go
package main

import(
    "context"
    "log"
    "app"
    "os"
    "go.temporal.io/sdk/client"
)

func main() {
    c, err := client.Dial(client.Options{})
    if err != nil {
        log.Fatalln("Unable to create client", err)
    }
    defer c.Close()

    options := client.StartWorkflowOptions{
       ID:        "my-first-workflow",
       TaskQueue: "greeting-tasks",
    }

    we, err := c.ExecuteWorkflow(context.Background(), options, app.GreetSomeone, os.Args[1])
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
}

/**
❯ go run app.go Tom
2023/01/07 22:02:29 INFO  No logger configured for temporal client. Created default one.
2023/01/07 22:02:29 Started workflow WorkflowID my-first-workflow RunID b288349a-ea7b-4743-bb32-2c7fd07eda4b
2023/01/07 22:02:29 Unable get workflow result workflow execution error (type: GreetSomeone, workflowID: my-first-workflow, runID: b288349a-ea7b-4743-bb32-2c7fd07eda4b): Workflow exited

❯ go run app.go Donna
2023/01/07 22:02:55 INFO  No logger configured for temporal client. Created default one.
2023/01/07 22:02:55 Started workflow WorkflowID my-first-workflow RunID 6c9bd15a-eede-4eba-af17-7f798df49db8
2023/01/07 22:02:55 Workflow result: Hello Donna !

*/
```

### Making changes to a Workflow

Backwards compatibility is an important consideration in Temporal. 

#### Input Parameters and Return Values

In general, you should avoid changing the number or types of input parameters and return values for your Workflow. Temporal recommends that your Workflow Function takes a single input parameter, a `struct`, rather than multiple input parameters.

#### Determinism

Also, your Workflow must be [deterministic](https://docs.temporal.io/concepts/what-is-a-workflow-definition/#deterministic-constraints). Temporal has a specific definition for this, but understanding it requires more detailed knowledge of Workflow Execution, so we can generalize for now. You can view determinism as a requirement that each execution of a given Workflow must produce the same output, given the same input. This means that you shouldn't do things like work with random numbers in your Workflow code. If you need to do things such as working with random numbers, the SDK provides safe alternatives. Activities, which you'll learn and use later in this course, provide a safe way to perform operations that interact with the outside world. Such operations, which might access files, databases, or network services, are inherently non-deterministic.

You can use the [Temporal Workflow Check](https://github.com/temporalio/sdk-go/tree/master/contrib/tools/workflowcheck) tool to analyze the Workflow Definitions you've written in Go and report typical cases of non-determinism in your code.

#### Versioning

Since Workflow Executions might run for months or years, it's possible that you'll need to make major changes to a Workflow Definition while there are already executions running based on the current definition of that Workflow. If these changes do not affect the deterministic nature of the Workflow, you can simply deploy them. However, you can use the SDK's ["Versioning"](https://docs.temporal.io/go/versioning/) feature to identify when a non-deterministic change is introduced, which allows older executions to use the original code and new executions to use the modified version.

<span style="color:darkorange">**After making changes to your application, you'll need to deploy them to the server. However, those changes won't take effect until you restart the Workers.** </span>

### Activities

#### What are Activities ?

In Temporal, you can use Activities to encapsulate business logic that is prone to failure. Unlike the Workflow Definition, there is no requirement for an Activity Definition to be deterministic.

In general, any operation that introduces the possibility of failure should be done as part of an Activity, rather than as part of the Workflow directly. While Activities are executed as part of Workflow Execution, they have an important characteristic; they're retired if they fail. If you have an extensive Workflow that needs to access a service, and that service happens to become unavailable, you don't want to re-run the entire Workflow. Instead, you just want to retry the part that failed, so you can define that code in an Activity and reference it in your Workflow Definition. The code within that Activity Definition will be executed, retried if necessary, and the Workflow will continue its progress once the Activity completes successfully.

#### Activity Definition

Just as a Workflow Definition is an exportable Go function, an Activity Definition is also an exportable Go function, and has the same rules about types allowed as input parameters and return values as the Workflow Definition. The function must return a value of type `Error`, but it may also return another value of any allowed type. 

> *Although Temporal does not require a specific first parameter for your Activity Definition, as it does for Workflow Definitions, we recommend making `context.Context` the first parameter in your Activity Definition. This will allow you to take advantage of some additional features, such as heartbeating that improves failure detection for long-running Activities.* 

```go
func GreetInSpanish(ctx context.Context, name string) (string, error) {
	base := "http://localhost:9999/get-spanish-greeting?name=%s"
	url := fmt.Sprintf(base, url.QueryEscape(name))

	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	translation := string(body)
	status := resp.StatusCode
	if status >= 400 {
		message := fmt.Sprintf("HTTP Error %d: %s", status, translation)
		return "", errors.New(message)
	}

	return translation, nil
}
```

##### Registering an Activity

```go
func main() {
    c, err := client.Dial(client.Options{})
    if err != nil {
        log.Fatalln("Unable to create client", err)
    }
    defer c.Close()

    w := worker.New(c, "greeting-tasks", worker.Options{})

    w.RegisterWorkflow(app.GreetSomeone)
    w.RegisterActivity(app.GreetInSpanish)

    err = w.Run(worker.InterruptCh())
    if err != nil {
        log.Fatalln("Unable to start worker", err)
    }
}
```

##### Executing an Activity

###### Specifying Activity Options

The first step to executing an Activity as part of your Workflow is to specify the options that govern its execution:

```
options := workflow.ActivityOptions{
    StartToCloseTimeout: time.Second * 5,
}Copy Code
```

Crucially, this includes a Start-to-Close timeout, which we recommend that you always set. Its value should be longer than the maximum amount of time you think the execution of the Activity should take. This allows the Temporal Cluster to detect to a Worker that crashed, in which case it will consider that attempt failed, and will create another task that a different Worker could pick up. There are [other types of timeouts](https://docs.temporal.io/tags/timeouts) that you can potentially set here, but they're less frequently used and not relevant to this course.

Essentially, you run an Activity via a Workflow.

```go
func GreetSomeone(ctx workflow.Context, name string) (string, error) {
    options := workflow.ActivityOptions{
        StartToCloseTimeout: time.Second * 5,
    }
    ctx = workflow.WithActivityOptions(ctx, options)

    var spanishGreeting string
    err := workflow.ExecuteActivity(ctx, GreetInSpanish, name).Get(ctx, &spanishGreeting)
    if err != nil {
        return "", err
    }

    return spanishGreeting, nil
}
```

#### How to handle Activity failure ?

##### Default Behavior

Temporal's default behavior is to automatically retry an Activity, with a short delay between each attempt, until it either succeeds or is canceled. That means that intermittent failures *require no action on your part*. When a subsequent request succeeds, your code will resume as if the failure never occurred. However, that behavior may not always be desirable, so Temporal allows you to customize it through a custom Retry Policy.

##### Changing the Timing and Number of Retry Attempts

Four properties that determine the timing and number of retries:

![img](https://cdn.talentlms.com/temporal/1664545202_retry-policy-property-table.png?Policy=eyJTdGF0ZW1lbnQiOlt7IlJlc291cmNlIjoiaHR0cHM6XC9cL2Nkbi50YWxlbnRsbXMuY29tXC90ZW1wb3JhbFwvMTY2NDU0NTIwMl9yZXRyeS1wb2xpY3ktcHJvcGVydHktdGFibGUucG5nIiwiQ29uZGl0aW9uIjp7IkRhdGVMZXNzVGhhbiI6eyJBV1M6RXBvY2hUaW1lIjoxNjczMTM2MDAwfX19XX0_&Signature=hCbVwHGlSOCvTwyme50upqu28OVtgqghaTiQzkxCPpSXQpoJizawWLMvPtfLh5muLCfYby8Ofisk793hRiYrBmhpd7TEGVDT8TRAE4lSVUmTl0yHsewfv5J4PrYT0khylMo5ZCEZhJ9ocllLf4xBT%2FTXgvcG9PgCiIC9gMolS0rsgxFd-f2MGf4CzdPcKl8rrRSXBX%2F9Uh0QytLFz5EbUlIRlmivINX4aT1Fdnm6ctw9qHSbWb7W1uxiDhRXTWGPSfLhQNeDBvC3-8iJN64kqq88PJD9m9OrXatH9BLFUY7PJM0XxbQ1ssEbdj51z12PRvIFj-DWiJdzVGcqy84RZQ__&Key-Pair-Id=APKAJDCWVQTW4P3KI3XA)

- The `InitialInterval` property defines how long after the initial failure the first retry will occur. By default, that's one second.
- The `BackoffCoefficient` is a multiplier, applied to the `InitialInterval` value, that's used to calculate the delay between each subsequent attempt. Assuming that you use the defaults for both properties, that means there will be a retry after 1 second, another after 2 seconds, then 4 seconds, 8 seconds and so on.
- The `MaximumInterval` puts a limit on that delay, and by default it's 100 times the initial interval, which means that the delays would keep doubling as described, but would never exceed 100 seconds.
- Finally, the `MaximumAttempts` specified the maximum count of retries allowed before marking the Activity as failed, in which case the Workflow can handle the failure according to its business logic.

```go



func GreetSomeone(ctx workflow.Context, name string) (string, error) {
    options := workflow.ActivityOptions{
        StartToCloseTimeout: time.Second * 5,
    }
    ctx = workflow.WithActivityOptions(ctx, options)

    var spanishGreeting string
    err := workflow.ExecuteActivity(ctx, GreetInSpanish, name).Get(ctx, &spanishGreeting)
    if err != nil {
        return "", err
    }

    return spanishGreeting, nil
}
```

