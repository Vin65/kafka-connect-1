# Kafka Connect CLI

This is a helper cli that will take in configuration and deploy kafka connect configuration for multiple instances. The reason for its existence is because with the current kafka connect API it is very tediuos to deploy a configuration for multiple instances of a database.

This helps:
* Avoid human error by manually hand writing many configurations
* Reduce replicated work as you can define one configuration for multiple databases
* Added upsert capability so if the connector exists it will update it in place

## Installation

You can install this cli by running the following command:
```
    go get github.com/commitdev/kafka-connect-cli
```

## Configuration
The configuration works in 'layers'. There are 3 defined layers:
* globals
* databases
* tables

'globals' define the base layer for all connectors, this is then overridden and merged with any configuation in the the 'databases' layer. Finally, 'tables' the final layer will take the result of 'globals' and 'databases' and override and merge its own changes. The total number of configurations that will be produced will be  the amount of 'databases' multiplied by the amount of 'tables'.

The configuration under each layer uses the standard kafka connect configuration. The only custom configuration is the 'version' this is attached to each connector name. The connector name format will be `TABLE-DATABASE-VERSION`. Example configuration can be found in the example directory.

## Quick start
In order to use tool you must first define a configuration, you can find an example of this in example/kafka-connect.yaml. Once this is done you can now run:
```
    # user / pass are optional - if used they are appened in basic auth headers (-u & -p)

    # Spits out what the configuration that would applied would look like
    kafka-connect -a KAFKA_CONNECT_ADDRESS spit -c kafka-connect.yaml

    # Applies the configuration, will create if not present and update if it is present
    kafka-connect -a KAFKA_CONNECT_ADDRESS apply -c kafka-connect.yaml
```

There are also other helper commands such as:
* list - lists availble connectors on the kafka connect cluster
* delete - will delete the connectors that would have been generated from the config provided to this tool
* describe - will describe a connector, its status etc.
* restart - will restart a specific connector or onnectors that would have been generated from the config provided to this tool

## Maintainers
* Pritesh Patel
