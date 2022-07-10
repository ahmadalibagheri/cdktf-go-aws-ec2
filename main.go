package main

import (
	"cdk.tf/go/stack/generated/aws"
	"cdk.tf/go/stack/generated/aws/ec2"
	"cdk.tf/go/stack/generated/aws/vpc"
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/aws/jsii-runtime-go"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

const (
	team    = "DevOps"
	company = "YourTeam"
)

func NewMyStack(scope constructs.Construct, id string) cdktf.TerraformStack {
	stack := cdktf.NewTerraformStack(scope, &id)
	aws.NewAwsProvider(stack, jsii.String("AWS"), &aws.AwsProviderConfig{
		Region: jsii.String("us-east-1"),
	})

	sg := vpc.NewSecurityGroup(stack, jsii.String("security-group"), &vpc.SecurityGroupConfig{
		Name:        jsii.String("CDKtf-TypeScript-Demo-sg"),
		Description: jsii.String("Allow traffic to the instance"),
		Ingress: []vpc.SecurityGroupIngress{
			vpc.SecurityGroupIngress{
				Protocol:   jsii.String("tcp"),
				FromPort:   jsii.Number(80),
				ToPort:     jsii.Number(80),
				CidrBlocks: &[]*string{jsii.String("0.0.0.0/0")},
			},
			vpc.SecurityGroupIngress{
				Protocol:   jsii.String("tcp"),
				FromPort:   jsii.Number(22),
				ToPort:     jsii.Number(22),
				CidrBlocks: &[]*string{jsii.String("0.0.0.0/0")},
			},
			vpc.SecurityGroupIngress{
				Protocol:   jsii.String("tcp"),
				FromPort:   jsii.Number(443),
				ToPort:     jsii.Number(443),
				CidrBlocks: &[]*string{jsii.String("0.0.0.0/0")},
			},
		},
		Egress: []vpc.SecurityGroupEgress{
			vpc.SecurityGroupEgress{
				Protocol: jsii.String("-1"),
				FromPort: jsii.Number(0),
				ToPort:   jsii.Number(0),
			},
		},

		Tags: &map[string]*string{
			"Name":    jsii.String("Security-Group-Golang-Ec2"),
			"Team":    jsii.String(team),
			"Company": jsii.String(company),
		},
	})

	ec2.NewInstance(stack, jsii.String("ec2-instance"), &ec2.InstanceConfig{
		Ami:                 jsii.String("ami-03d315ad33b9d49c4"),
		InstanceType:        jsii.String("t2.micro"),
		KeyName:             jsii.String("DevOps"),
		VpcSecurityGroupIds: &[]*string{sg.Id()},
		Tags: &map[string]*string{
			"Name":    jsii.String("Ec2-Instance-Golang"),
			"Team":    jsii.String(team),
			"Company": jsii.String(company),
		},
	})

	return stack
}

func main() {
	app := cdktf.NewApp(nil)

	stack := NewMyStack(app, "cdktf-go-aws-ec2")
	cdktf.NewRemoteBackend(stack, &cdktf.RemoteBackendProps{
		Hostname:     jsii.String("app.terraform.io"),
		Organization: jsii.String("jigsaw373"),
		Workspaces:   cdktf.NewNamedRemoteWorkspace(jsii.String("cdktf-go-aws-ec2")),
	})

	app.Synth()
}
