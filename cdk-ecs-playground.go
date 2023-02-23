package main

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
	ec2 "github.com/aws/aws-cdk-go/awscdk/v2/awsec2"
	ecs "github.com/aws/aws-cdk-go/awscdk/v2/awsecs"
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/aws/jsii-runtime-go"
	"os"
)

type EcsPlayGroundStackProps struct {
	awscdk.StackProps
}

func EcsPlayGroundStack(scope constructs.Construct, id string, props *EcsPlayGroundStackProps) awscdk.Stack {
	var sprops awscdk.StackProps
	if props != nil {
		sprops = props.StackProps
	}
	stack := awscdk.NewStack(scope, &id, &sprops)

	vpc := ec2.Vpc_FromLookup(stack, jsii.String("VPC"), &ec2.VpcLookupOptions{
		IsDefault: jsii.Bool(true),
	})

	ecs.NewCluster(stack, jsii.String("My-CDK-Cluster"), &ecs.ClusterProps{
		ContainerInsights: jsii.Bool(true),
		Vpc:               vpc,
		ClusterName:       jsii.String("CDKClusterPlayGround"),
	})

	return stack
}

func main() {
	defer jsii.Close()

	app := awscdk.NewApp(nil)

	EcsPlayGroundStack(app, "EcsPlayGroundStack", &EcsPlayGroundStackProps{
		awscdk.StackProps{
			Env: env(),
		},
	})

	app.Synth(nil)
}

func env() *awscdk.Environment {

	return &awscdk.Environment{
		Account: jsii.String(os.Getenv("AWS_ACCOUNT")),
		Region:  jsii.String(os.Getenv("AWS_REGION")),
	}
}
