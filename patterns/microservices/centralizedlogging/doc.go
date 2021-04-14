/*Centralized Logging System.

Logging API service that everything messages legs to.
This validates the logs, stores them and provides UI or portal to view them.

Structure formate suggestions:
- Log levels
- Date and time
- Correlation ID (id that goes with transaction so you can follow it)
- Host and app info
- Actual log message

Best to use shared library to do logging format.
*/

package centralizedlogging
