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
Preventing Goroutine Leaks: If a go routine is responsible for creating a goroutine, it is also responsible for ensuring it can stop.
	- Goroutine receiving on a channel: Also pass in a done channel, and use that to cancel go routine if need be via select statement.
	- Goroutine blocked on read of channel: Also pass in done channel.
Or Channel: Recursive function to combine multiple channels and block until any channel fires.
Error Handling: Create structures with result and error, and return this instead. The place that needs the result of
		a concurrent process will also probably care about any errors received. This is very idiomatic with the way
		most functions return a result and an error, and after such functions you will usually see an if err statement.
Pipeline: A series of things that take in data, do a transform, and pass it back out. Each transform is a stage of the pipeline.
		- A stage consumes and returns the same type
		- Very functional programming like (Building up higher order functions by composing functions)
Generators: Transforms a list of things into a channel containing them things. Used to parrallelise iteration/tasks on a list
Fan-out: Starting multiple goroutines to handle input from one pipeline
Fan-in: Combining multiple results into one channel
Or-Done channel: range over channel safely with ways to cancel loop
tee-channel: Pass in a channel to read from, and it returns 2 channels that will get the same value
Bridge-channel: Destructures a channel of channels into a simple channel
Queues: Allows decoupling of stages:
		Little Law: L = Lam * W
		- L = Average number of units in the system
		- Lam = average arrival rate of units
		- W = average time a unit spends in a system
		- Example 1: How many requests per second our pipeline can handle (Lam):
			L = 3 stages all processing a request = 3r
			W = 1 second
			Lam = 3 => We can handle 3 requests per second
		- Example 2: How big does out queue need to be to handle 100,000 requests per second.
			Pipeline has 3 stages: L = Lr - 3r
			Lam = 100,000
			Each request takes 1ms
			Lr - 3r = 100,000r/s * 0.0001s
			Lr - 3r = 10r
			Lr = 7r
			L = 7 => Our queue needs to be 7 units long
Context Package:
- The allow for done data to be passed to child goroutines
- Contexts are immutable
- To change the context for a child (add in a cancel function), you use the functions on the current conext, and pass along the resultant new context
- Instances of context are meant to flow through the programs call-graph
- Never store context as member variables, always pass them in as contexts internally change
- Creating an empty context:
	- Background(): Normal way of creating a default context
	- TODO(): Used when you don't know what context to use, but must not reach production
	- can pass key value pairs to context value, but they must be comparable and safe top access by multiple go routines
		- it as advised you create a custom key and value type that everything that uses ctv values uses to manage type safety
		- Types of data should be scoped and scoping should be scrutinized by the team as a whole. An example:
			- If data is process data, don't transit across api boundaries (vice verse)
			- Data should be immutable
			- Mostly use simple types
			- Data should be data and not types or methods
			- Don't pass data that drives functionality
	- Takeway: The done/cancel stuff is really good. It means cancels can be littered throughout for different things.
		The value store should be used at own risk, and sparingly.


*/

package concurrency
