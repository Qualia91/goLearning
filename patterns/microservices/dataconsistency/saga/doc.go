/*Saga Pattern.

Service initiates a saga request to the saga log (database).
Saga execution coordinator (SEC) listens to log and actions requests.
Sends requests in request is requested, and compensation requests if something failed (rollback request)


Following example:
Initiator service (9000) receives get and sends a message to saga log service (9009).
SEC listens to messages sent into log service and sends requests out to serviceOne (9001) and serviceTwo (9002).
serviceOne sends a complete back, serviceTwo send s a fail back.
SEC sends compensation request out to both services to get them to roll back.*/

package saga
