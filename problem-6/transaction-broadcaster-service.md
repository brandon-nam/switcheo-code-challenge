# Transaction Broadcaster Service Design 

A transaction broadcaster service's role is to transfer the signed transaction data to a blockchain node, from which it will validate the data and propage to other nodes in the blockchain network.

To implement the broadcaster service with the given specification, we can approach the design of the architecture through following the workflow of a POST request to /transaction/broadcast. 

As the client makes a POST request, the client will receive a status code 200 if the job has been added successfully to a message queue that connects the broadcaster service with the blockchain node. If it has not been added successfully, then status code of 4xx will be sent to the client to retry requesting.

The use of message queue allows decoupling and horizontal scaling of the worker nodes, which allows for faster .

Since we are dealing with sensitive data, we should use SSL or TLS encryption for in-transit encryption. 

After the message is added to the message queue, the blockchain node that is subscribed to the message queue will pull the messages from the message queues and make RPC calls. When succeeded, the blockchain node will send a completion message to another queue that is attached with the service that deals with the database to store and update the job status. The job statuses can then be viewed by another service for visualization. When failed to process, the job can be sent to dead letter queue, send the failure status to the job status queue, and the blockchain node can periodically check if there is anything in the dead letter queue to re-process it. When the job processing goes over 30 seconds, we can timeout and put the job into dead letter queue, send a timeout status to the job status queue, and similarly retry processing periodically. 

I believe the use of database is more preferred than memory for status tracking since it seems according to the specification that a user would like to keep all successful / unsuccessful transaction history, which sounds like it requires long term storage. 

As an additional requirement, if we want longer storage of the failed jobs, we can attach a service that stores the jobs in the dead letter queue into a database, then the transaction broadcast service can periodically check the database for failed attempts that need to be re-processed. This way, even if the broadcaster service restarts unexpectedly, it should still eventually broadcast the failed attempts. 

If an admin wants to, at any point in time, retry a failed broadcast manually, the admin can check the status of the job through the methods described above, and send another request through the same API endpoint. If this is done, we would need to check the dead letter queue and dequeue the job to prevent from processing the job twice. We can add a field to the message that is sent to the queue to distinguish if the message is for a manual retry or for proper requests, then the blockchain node, when it receives the manual retry message, check the long term dead letter storage DB to see if there is one already, and delete it if exists. If not, it means the message has been processed or is being processed so we abort the processing.

To conclude, to solve the decoupling and horizontal scaling we could use message queues. For status tracking we can send messages to another service that stores the statuses of the jobs into the database that can be queried by the user ID to give the list of successful / unsuccessful transactions by a user. For automatic retries we can use dead letter queues for short-term retries. For long-term retries to prevent loss of letters when the system is down, we can attach a service that stores the contents in the dead letter queue into the DB. For manual retries, we can check the dead letter DB and prevent duplicate message processing.


