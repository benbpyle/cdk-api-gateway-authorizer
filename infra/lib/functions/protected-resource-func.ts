import { Construct } from "constructs";
import { GoFunction } from "@aws-cdk/aws-lambda-go-alpha";
import { Duration } from "aws-cdk-lib";
import * as path from "path";
import {
    LambdaIntegration,
    Resource,
    RestApi,
} from "aws-cdk-lib/aws-apigateway";

export class ProtectedResource extends Construct {
    private readonly _func: GoFunction;

    constructor(scope: Construct, id: string, api: RestApi) {
        super(scope, id);

        this._func = new GoFunction(this, `ProtectedResource`, {
            entry: path.join(__dirname, `../../../src/protected-resource`),
            functionName: `protected-resource-func`,
            timeout: Duration.seconds(30),
        });

        api.root.addMethod(
            "GET",
            new LambdaIntegration(this._func, {
                proxy: true,
            })
        );
    }

    get function(): GoFunction {
        return this._func;
    }
}
