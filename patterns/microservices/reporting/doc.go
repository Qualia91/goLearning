/*Providing Reporting.

Solutions:
- Reporting Service: Provides data for reporting by calling other services. Each service has a separate
		service that deals with	reporting for that service.
- Reporting Data Push App: Client services push reporting information to reporting service, then reporting service makes
		report when needed from local data.
- Reporting event subscribers: Using a message broker so each service can send reports to a queue, and reporting systems
		subscribe to them. By combining this with event sourcing (only changes are saved), reporting can be done
		on any time through history.
- Using backup databases for reporting: Databases are backed up. You can use these to do reporting.
- ETL and data-wharehouses: Extraction, transform and load step that takes in databases across microservice, the saves it to
		datawarehouse that reporting uses.
*/

package reporting
