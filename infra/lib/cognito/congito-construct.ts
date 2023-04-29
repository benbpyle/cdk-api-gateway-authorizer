import { Construct } from "constructs";

import * as cdk from "aws-cdk-lib";
import * as cognito from "aws-cdk-lib/aws-cognito";
import * as lambda from "aws-cdk-lib/aws-lambda";
import { Duration } from "aws-cdk-lib";

export class CognitoConstruct extends Construct {
    private readonly _pool: cognito.UserPool;
    constructor(scope: Construct, id: string) {
        super(scope, id);

        this._pool = new cognito.UserPool(this, "SamplePool", {
            userPoolName: "SamplePool",
            selfSignUpEnabled: false,
            signInAliases: {
                email: true,
                username: true,
                preferredUsername: true,
            },
            autoVerify: {
                email: false,
            },
            standardAttributes: {
                email: {
                    required: true,
                    mutable: true,
                },
            },
            customAttributes: {
                isAdmin: new cognito.StringAttribute({ mutable: true }),
            },
            passwordPolicy: {
                minLength: 8,
                requireLowercase: true,
                requireDigits: true,
                requireUppercase: true,
                requireSymbols: true,
            },
            accountRecovery: cognito.AccountRecovery.EMAIL_ONLY,
            removalPolicy: cdk.RemovalPolicy.DESTROY,
        });

        this._pool.addClient("sample-client", {
            userPoolClientName: "sample-client",
            authFlows: {
                adminUserPassword: true,
                custom: true,
                userPassword: true,
                userSrp: false,
            },
            idTokenValidity: Duration.minutes(60),
            refreshTokenValidity: Duration.days(30),
            accessTokenValidity: Duration.minutes(60),
        });
    }

    get pool(): cognito.UserPool {
        return this._pool;
    }
}
