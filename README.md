# http-archetype
This http adapter archetype will be used for testing on Dockers's world.

In this repository i will be pushing me job and eventually, when this is done, this app will be into a Docker container to make me better at this engine.

The idea is to create 2 Docker containers to interact with each other like this: 

Case 1 (Creating a transaction)
Http-Adapter->Post Transaction->Database-Adapter/Database-Adapter->Response->Http-Adapter

Case 2 (Asking for a transaction)
Http-Adapter->Get Transaction->Database-Adapter/Database-Adapter->Transaction->Http-Adapter