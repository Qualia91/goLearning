/*Concurrency patterns.

This is a collection of useful concurrency patterns in GoLang. as well as notes made
from the book "Concurrency in Go: Tools and techniques for developers"

Common Issues:
- Race conditions: 2 or more operations must execute in the correct order, which is not guaranteed.
- Atomicity: Within the context it is operating, its is indivisible or uniteruptable. For instance i++ is made
		of 3 statements = retrieve i, increment value, write i. Each operation by itself is atomic, but the combination
		might not be. In the case of i being used in a go routing, and outside of a go routine, i++ would not ba atomic.
- Memory access synchronization: Reading and writing to the same variable in multiple threads will produce different results.
- Deadlocks: All processes are waiting on another one, meaning nothing can run.
		Coffman conditions for a deadlock:
			- Mutual Exclusion: A concurrent process holds exclusive rights to a resource at any one time.
			- Wait for condition: A concurrent process must hold a resource and wait for adition resource.
			- No preemption: A resource held by a process can only be help by that process.
			- Circular wait: A concurrent process must be waiting on a chain of concurrent processes that in turn wait on the original process.
- Livelock: Actively preform tasks that do nothing to move the program forward. Common reason: 2 or more concurrent process
		are attempting to prevent a deadlock without coordination. Part of a larger set of problems called Starvation.
- Startvation: A process can not get all the resources it needs to preform work.
- People: Some tips to follow when writing code:
		- Add comments to concurrent functions that include who is responsible for concurrency, how the problem space is
		  mapped to concurrency primitives and who is responsible for synchronization.
		- Take a functional approach (no side effects)

Communicating Sequential Processes (CSP):
- also see process calculus
- Channels, go routines and select statements in go are based on this theory (shown in a programming language
	of the same name) and paper

When to use primitives (mutex's), and when to use channels:
- Are you trying to transfer ownship of data: Passing data from one thread to another - use channels
- Are you trying to guard the internal state of a struct: Making a variable thread safe with a lock - use primitives
- Are you trying to coordinate multiple pieces of logic: Channels as its easier to extend
- Is it a performance-critical section: Channels use memory access sync and so can only be slower than primitives,
	but may indicate program needs restructuring

Go's philosophy on concurrency: Aim for simplicity, use channels where possible, and treat goroutines like a free resource

Go routines are a special case of coroutines: Coroutines are nonpreemptive (they cannot be interuputed). Go routines
can be interrupted when they have become blocked.

Goroutines are hosted on a implimenetation of a M:N scheduler: M green threads are mapped to N OS threads. Goroutines are
the scheduled onto the green threads.

Sync Package Objects:
- WaitGroup
- Mutex
- Cond
- Once
- Pool

Select statements are non-deterministic.


runtime.NumCPU() = number of CPU cores on machine

runtime.GOMAXPROCS() = Number of threads that will host worker queues
	Reason to change this: To find race conditions more easily


Patterns:
Confinement: Ensuring data is only ever available from one concurrent process.
	2 types	defer wg.Done()
		var buff bytes.Buffer
		for _, b := range data {
		fmt.Fprintf(&buff, "%c", b)
		}
		fmt.Println(buff.String()):
	- Ad hoc: Achieve confinement through a convention
	- Lexical: Use lexical scoping to expose only the correct data and concurrency primitives. This
		involves creating objects within functions, returning the objects from functions after a routine
		is created that needs to use that object, then running the other routine and passing in the object returned
		from the first function.
Fo-Select loop: Select nested in for loop. Reasons for use:
	- Sending iteration variables out on a channel
	- Looping infinitely waiting to be stopped




*/

package concurrency
