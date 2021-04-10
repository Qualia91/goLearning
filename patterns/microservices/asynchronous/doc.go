/*Asynchronous Microservices.

Why:
Specific components talk to each other in async way.
Reasons:
- Need to start a task but don't need a response straight away.
- Long running tasks
- Decouple clients and service
- Better user experience (no waiting for loading)

How:
- Event Based: Intermediate message queue with events passed through it
- Async API Calls: Request/Ack using callbacks. Meaning request is made but work isn't received until done, just and ack is


Event Based
Examples:
- RabbitMQ
- ZMQ
- KubeMQ
- Kafka


*/
package asynchronous
