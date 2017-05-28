# playground (0.1 alpha)

The playground contains some server application, to test my libraries and architecture.

## tested libraries/application
- assertion (test tool)
- restclient (http rest client)
- pqx (postgres database lib)
- gox (build tool)


## architecture proposal
- [root]
    - [client] 
    // this folder contains a client implementation for every rest endpoint (that is used in production)
        - client.go
        - dto.go
    - [internal]
        - [data] 
        // this folder contains every database action and the table entries.
            - [?data] 
            // if you have to split up he folder (too much content)
        - [public] 
        // contains files regarding web development (templates, css, javascript, images,...)
        - [service] 
        // contains some logic, that is too big for the rest endpoint files.
            - [?service] 
            // if you have to split up he folder (too much content)
        - [sys] 
        // this folder contains the config and system relevant structs (this folder should not have any dependencies)
        - [test] 
        // this folder contains everything that you need for testing. 
            - init.go
            // contains methods to init your test environment (open db or start server)
            - testclient.go 
            // this file contains for every client method a testclient method
        - [util] 
        // contains some helper functions that are not [service] specific
        - init.go
        - ....go rest-endpoint files 
        // these files contains the rest endpoint logic, every method in this files is a rest endpoint, private methods should be handled in [util] or [service]
    - main.go
    - config files