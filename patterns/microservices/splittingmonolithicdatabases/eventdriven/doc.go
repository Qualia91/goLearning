/*Event driven db synchronization.

The idea is simple that when a service saves data, they will issue an event that some data has changed that other
services can listen do and act accordingly.

For this example, the client service will save some data and send out that data has changed to a rabbitmq message broker
on an exchange with name data_has_changed.
2 services in the background (service_one and service_two) both care about this data change, and so are bound to the rabbitmq
exchange with name data_has_changed.
The client will save some other_data and send on an exchange called other_data_has_changed. Only service_one is bound to this
data as service_two does not need to know this.*/

package eventdriven
