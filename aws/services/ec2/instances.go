package ec2

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws/client"
	"github.com/awslabs/aws-sdk-go/aws"
	"github.com/awslabs/aws-sdk-go/gen/ec2"
)

func check(e error) {

	if e != nil {
		panic(e)
	}
}

/**
Credit to https://gist.github.com/stephen-mw/9f289d724c4cfd3c88f2#file-list_ec2_instances-md
 */
func ListInstances(ec2 ec2.EC2) {

	// Only grab instances that are running or just started
	filters := []ec2.Filter{
		ec2.Filter{
			aws.String("instance-state-name"),
			[]string{"running", "pending"},
		},
	}
	request := ec2.DescribeInstancesRequest{Filters: filters}
	result, err := client.DescribeInstances(&request)
	check(err)

	for _, reservation := range result.Reservations {
		for _, instance := range reservation.Instances {
			fmt.Println(*instance.InstanceID)
		}
	}
}
