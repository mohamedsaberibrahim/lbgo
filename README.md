# Load-balancer algorithms implementation in Go.

## Algorithms check lists:

### - Static load balancing algorithms
- [x] Round robin: Distributes traffic to a list of servers in rotation using the Domain Name System (DNS). An authoritative name server will have a list of different A records for a domain and provides a different one in response to each DNS query.

- [ ] Weighted round robin: Allows an administrator to assign different weights to each server. Servers deemed able to handle more traffic will receive slightly more. Weighting can be configured within DNS records.

- [ ] IP hash: Combines incoming traffic's source and destination IP addresses and uses a mathematical function to convert it into a hash. Based on the hash, the connection is assigned to a specific server.
### - Dynamic load balancing algorithms
- [ ] Least connection: Checks which servers have the fewest connections open at the time and sends traffic to those servers. This assumes all connections require roughly equal processing power.

- [ ] Weighted least connection: Gives administrators the ability to assign different weights to each server, assuming that some servers can handle more connections than others.

- [ ] Weighted response time: Averages the response time of each server, and combines that with the number of connections each server has open to determine where to send traffic. By sending traffic to the servers with the quickest response time, the algorithm ensures faster service for users.

- [ ] Resources-based: Distributes load based on what resources each server has available at the time. Specialized software (called an "agent") running on each server measures that server's available CPU and memory, and the load balancer queries the agent before distributing traffic to that server.
