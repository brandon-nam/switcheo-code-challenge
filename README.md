# switcheo-code-challenge

## Problem 5 Consensus Breaking Change 

To check the consensus breaking change, checkout to <code> consensus-breaking-change</code> branch. 

A consensus breaking change is the modification that requires every node to be updated with so that the way of processing transactions are still validated. 

My change in the x/ledger/keeper/ledger.go of re-using the ID of the transaction is an example of consensus-breaking change. Before the change the ID of the transaction was not re-used by not decrementing the ID when the transaction is removed, but with the change now the ID is decremented everytime a transaction is removed. This change definitely affects the way of processing and accessing the transactions. Now the same ID integer can refer to two different states, which means it breaks the consensus. 


## Problem 5 Commands
Here is the list of cli commands available for my blockchain application. 

### Create Ledger 
<code> ledgerd tx ledger create-ledger \<title: string> \<body: string> \<cost: uint> --from nam </code>

### Read & List All Ledger & List Cost-Filtered Ledger

**Read a ledger by its ID**

<code> ledgerd query ledger show-ledger \<id: uint> </code>

Note that the ID of the ledger is by the LedgerCount, and the LedgerCount does not get decremented once removed. Thus, the newly created Ledger entity will not have the ID of the deleted Ledger.

**List All Ledgers**

<code>ledgerd query ledger list-ledger</code>

**List Cost-Filtered Ledgers**

<code>ledgerd query ledger list-ledgers-cost-filter \<min_cost: uint> \<max_cost: uint></code>

### Update Ledger

<code>ledgerd tx ledger update-ledger \<title: string> \<body: string> \<cost: uint> \<id: uint> --from nam</code>

### Delete Ledger

<code>ledgerd tx ledger delete-ledger \<id: uint> --from nam</code>


