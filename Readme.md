# DynamoDB Local Testing With GO

A demonstration of how to test DynamoDB using Dynamo Local in a Docker
container.

## How it Works

The Item library is a simple library that uses DynamoDB to store and
retrieve Items (a struct representing the object we wish to store in
Dynamo).

This library required testing. We could mock out Dynamo, but as it's the
the main function of the library, we really want to test against a real
Dynamo table. However we don't want to have to create Dynamo tables in
AWS each time we want to run a test. Fortunately AWS provides a local
version of DynamoDB for us to test against. We can run this in a
container and point out Dynamo Client at hit when we run our tests.

This also has the benefit of allowing us to quickly create Dynamo Tables
and tear them down afterwords (by destroying the container). We create a
new, randomly used Dynamo table for each and every test, which means that
the tests won't clash, for instance if an item you expect not to exist
actually does because a previous test created it.

