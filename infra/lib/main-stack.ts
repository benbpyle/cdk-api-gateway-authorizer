import { Construct } from "constructs";
import * as cdk from "aws-cdk-lib";
import { ApiGatewayConstruct } from "./api/api-gateway-construct";
import { StackProps } from "aws-cdk-lib";
import { AuthorizerFunction } from "./functions/authorizer-func";
import { CognitoConstruct } from "./cognito/congito-construct";
import { ProtectedResource } from "./functions/protected-resource-func";

export class MainStack extends cdk.Stack {
    constructor(scope: Construct, id: string, props: StackProps) {
        super(scope, id, props);
        const pool = new CognitoConstruct(this, "Cognito");
        const authorizer = new AuthorizerFunction(
            this,
            "AuthorizerFunction",
            pool.pool.userPoolId
        );
        const api = new ApiGatewayConstruct(
            this,
            "ApiGateway",
            authorizer.function
        );
        const protectedResource = new ProtectedResource(
            this,
            "ProtectedResource",
            api.api
        );
    }
}
