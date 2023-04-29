import { Construct } from "constructs";
import {
    BasePathMapping,
    DomainName,
    RestApi,
    TokenAuthorizer,
} from "aws-cdk-lib/aws-apigateway";
import { IFunction } from "aws-cdk-lib/aws-lambda";

export class ApiGatewayConstruct extends Construct {
    private readonly _api: RestApi;

    constructor(scope: Construct, id: string, func: IFunction) {
        super(scope, id);

        const authorizer = new TokenAuthorizer(this, "TokenAuthorizer", {
            authorizerName: "BearTokenAuthorizer",
            handler: func,
        });

        this._api = new RestApi(this, "RestApi", {
            description: "Sample API",
            restApiName: "Sample API",
            deployOptions: {
                stageName: `main`,
            },
            defaultMethodOptions: {
                authorizer: authorizer,
            },
        });
    }

    get api(): RestApi {
        return this._api;
    }
}
