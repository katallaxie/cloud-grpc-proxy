# protoc-gen-aws-service-proxy

![Main](https://github.com/awslabs/aws-lambda-dart-runtime/workflows/Main/badge.svg?branch=master)
[![License Apache 2](https://img.shields.io/badge/License-Apache2-blue.svg)](https://www.apache.org/licenses/LICENSE-2.0)

The AWS gRPC Service Proxy is an **experimental** project to code generate a proxy from [Protobuf](https://developers.google.com/protocol-buffers) that enables you to use AWS services ([AWS Lambda](https://aws.amazon.com/lambda/), [Amazon DynamoDB](https://aws.amazon.com/dynamodb/)) and more via [gRPC](https://grpc.io/).

> :warning: the API of the project is not yet stable.

## Getting Started

You provide your service definitions in [Protobuf](https://developers.google.com/protocol-buffers).

```
syntax = "proto3";

package proto;

import "google/protobuf/timestamp.proto";
import "api/annotations.proto";

// Example ...
service Example {
  // Insert ...
  rpc Insert(Insert.Request) returns (Insert.Response) {
    option (amazon.api.services).lambda = {
      function_name: "arn:aws:lambda:us-east-2:123456789012:function:my-function:1"
    };
  };

  // Update ...
  rpc Update(Update.Request) returns (Update.Response) {
    option (amazon.api.services).dynamodb.UpdateItem = {
      TableName: "aws-grpc-service-proxy"
      ReturnValues: "ALL_NEW"
      UpdateExpression: "SET #Y = :y, #AT = :t"
      ExpressionAttributeNames: [
        { key: "key", value: "value" }
      ]
    };
  };
}

// Item ...
message Item {
 // UUID ...
  string uuid   = 1;
  // Title ...
  string title  = 2;
  // Body ...
  string body   = 3;
}

// Insert ...
message Insert {
  // Request ...
  message Request {
    Item item = 1;
  }
  // Response ...
  message Response {
    string uuid = 1;
  }
}

// Update ...
message Update {
  // Request ...
  message Request {
    Item item = 1;
  }
  // Response ...
  message Response {
    string uuid = 1;
  }
}
```

The included `protoc-gen-aws-service-proxy` generates a fully functioning proxy from the [Protobuf](https://developers.google.com/protocol-buffers).

## License

[Apache 2.0](/LICENSE)
