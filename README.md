# playground (0.1 alpha)

The playground contains some server application, to test my libraries and architecture.

## tested libraries/application
- assertion (test tool)
- restclient (http rest client)
- pqx (postgres database lib)
- gox (build tool)


## architecture proposal
- [root]
    - [cmd]
        - [progs]
        // small programs to generate code
    - [clientlib] 
    // this folder contains a client lib implementation for every rest endpoint (that is used in production)
        - client.go
        - dto.go
    - [internal]
        - [data]
            - [enum] (access to **nothing**)
            // folder contains all enum types of your system
            - [dto] (access to [enum])
            // contains all dtos needed inside the server and for rest endpoints needed in the server(web)
            **(this folder should not have any dependencies)**
                - [?dto]
                // if you have to split up the folder (too much content) 
                -> should split up by topic
            - [dao] (access to [sys])
            // this folder contains every database action and the table entries.
                - [migration]
                // folder that contains all migration files 
                - [?dao] 
                // if you have to split up the folder (too much content) 
                -> should split up by topic not entity
                -> if you need a join between two topics in sql, do the join not in the sql, do it outside 
                -> less coupling -> later on you can extract the topics better in separate services
                - [internal]
                // contains all private [dao] content -> db variable, global db functions
        - [public] (access to **nothing**)
        // contains files regarding web development (templates, css, javascript, images,...)
        // shouldn't contains any go files
        - [service] (access to [sys], [data], [public])
        // contains some logic, that is too big for the rest endpoint files.
            - [?service] 
            // if you have to split up he folder (too much content)
        - [sys] (access to **nothing**)
        // this folder contains the config, paths, website urls, maps and system relevant structs (logging)
        **(this folder should not have any dependencies)**
        - [test] (this folder should only be accessed by a `?_test` package)
        // this folder contains everything that you need for testing.
            - [testAPI]
            // folder contains some rest endpoints, that are needed for testing, but shouldn't be in production
            - init.go
            // contains methods to init your test environment (open db or start server)
            - testclient.go 
            // this file contains for every client method a testclient method
            // also contains for every rest endpoint that is used inside the system (website) a client method
        - [util] (access to [data], [sys], [public])
        // contains some helper functions that are not [service] specific, e.g. converter, 
        - init.go
        - ....go rest-endpoint files 
        // these files contains the rest endpoint logic, every method in this files is a rest endpoint, private methods should be handled in [util] or [service]
    - main.go
    - config files