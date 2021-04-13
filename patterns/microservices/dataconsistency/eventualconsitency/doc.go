/*Eventual Consitency.

Message to update database is done over message broker.

In example, message broker is RabbitMQ. To start rabbitMQ:
docker run -it --rm --name rabbitmq -p 5672:5672 -p 15672:15672 rabbitmq:3-management
*/

package eventualconsitency
