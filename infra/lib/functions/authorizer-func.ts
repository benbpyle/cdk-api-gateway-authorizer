import { Construct } from "constructs";
import { GoFunction } from "@aws-cdk/aws-lambda-go-alpha";
import { Duration } from "aws-cdk-lib";
import * as path from "path";
import {
    LambdaIntegration,
    Resource,
    RestApi,
} from "aws-cdk-lib/aws-apigateway";

export class AuthorizerFunction extends Construct {
    private readonly _func: GoFunction;

    constructor(scope: Construct, id: string, poolId: string) {
        super(scope, id);

        this._func = new GoFunction(this, `AuthorizerFunc`, {
            entry: path.join(__dirname, `../../../src/authorizer`),
            functionName: `authorizer-func`,
            timeout: Duration.seconds(30),
            environment: {
                USER_POOL_ID: poolId,
            },
        });
    }

    get function(): GoFunction {
        return this._func;
    }
}
