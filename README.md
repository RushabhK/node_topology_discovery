Implementation of Network Topology Discovery.

# Design Philosophy

### Setup
* A UDP Server would run on each machine to serve the requests from neighboring nodes.
* Each machine would make request for discovery to the neighboring nodes via a UDP connection.

### Algorithm for discovery on a machine M
* Initialize the Visited Nodes with the M's IP and port
* Traverse through all the neighboring nodes on M
* Initialize the response with the M's IP:PORT as the key, with its neighbors and Machine Info as the value
* If the neighboring node is not visited
    * Create a UDP connection to the neighboring node
    * Request for all the nodes the neighbor can discover other than the already visited ones
    * Add all the nodes from response to the visited nodes
* Accumulate the responses in the result
* Transform the result with using the Machine names
* Write the result to a file

### Algorithm of the UDP server to serve the request on a machine M
* Get the list of all the visited nodes from the request
* Add M to the list of visited nodes
* Initialize the response with the M's IP:PORT as the key, with its neighbors and Machine Info as the value
* Traverse through all the neighboring nodes on M
* If the neighboring node is not visited
    * Create a UDP connection to the neighboring node
    * Request for all the nodes the neighbor can discover other than the already visited ones
    * Add all the nodes from response to the visited nodes
* Accumulate the responses and send it to the requested machine