# New Exchange
An exchange marketplace written in Go

### What is an exchange?
A marketplace where buyers ans sellers interact to trade listed products at an agreeable price.

### What its he goal of the project?
- Learning about exchanges.
- Learning go language in the process
- Building something interesting.
- Possibly providing an alternative implementation for proprietary software.

### What is required from an exchange?
We need some way to have orders placed.
We need some way to match orders.
We need some way to notify buyers and sellers of trades.
We need some way for new products to be added.
We need some way for traders to be regirstered.

#### Placing orders
*We acccept the following:*
product: the unique identifier of the product the order relates to.
type: bid | aks (bid for buying and ask for selling)
qty: number of products the order is for.
amount: value of a single qty of the product in cents. This will be mutliplied by qty to get to total order value.

### How do you test the exchange? 
TBD


### V1
A working exchange, with listed products, where an order can be placed, matched and confirmed. The exchange will use a simple matching algorithm for starters.

