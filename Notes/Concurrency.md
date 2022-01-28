## Concurrency and Parallelism 

- **Concurrent programming** 
  -   an application is a process running on a machine(a process is an independently executing entity that runs in its own address space) 
  -   a process is composed of one or more operating system threads(light weighted process) --> threads are concurrently executing entities that share the same address space  
- **Execution thread** 
  -   an execution thread is a sequence of executable statements and it may or may not interact with other threads(threads often share some varibles with other threads) 
  -   a thread can be blocked; access a shared variable or resource used currently by another thread. waiting for the result or completion of another thread 
  -   an application program often be divided into multiple processes and potentially many threads 

| Parallel programming       | Concurrent programming  |
| -------------------------- | ------------------------------------------ |
| process running on different processor or cores simultaneously  | concurrent means expressing program structure such that is organized into independently executing actions|
| /                          |it also target a single processor, the process run in turn over time according to a schedule|

## Goroutines 

## Channels 
