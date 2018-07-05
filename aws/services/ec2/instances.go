package ec2

import (
	"github.com/aws/aws-sdk-go/service/ec2"
)

func check(e error) {

	if e != nil {
		panic(e)
	}
}

func ListInstances(ec2 ec2.EC2) {

}
