// THIS FILE IS AUTOMATICALLY GENERATED. DO NOT EDIT.

package iam

import (
	"github.com/gophergala2016/goad/cli/Godeps/_workspace/src/github.com/aws/aws-sdk-go/private/waiter"
)

func (c *IAM) WaitUntilInstanceProfileExists(input *GetInstanceProfileInput) error {
	waiterCfg := waiter.Config{
		Operation:   "GetInstanceProfile",
		Delay:       1,
		MaxAttempts: 40,
		Acceptors: []waiter.WaitAcceptor{
			{
				State:    "success",
				Matcher:  "status",
				Argument: "",
				Expected: 200,
			},
			{
				State:    "retry",
				Matcher:  "status",
				Argument: "",
				Expected: 404,
			},
		},
	}

	w := waiter.Waiter{
		Client: c,
		Input:  input,
		Config: waiterCfg,
	}
	return w.Wait()
}

func (c *IAM) WaitUntilUserExists(input *GetUserInput) error {
	waiterCfg := waiter.Config{
		Operation:   "GetUser",
		Delay:       1,
		MaxAttempts: 20,
		Acceptors: []waiter.WaitAcceptor{
			{
				State:    "success",
				Matcher:  "status",
				Argument: "",
				Expected: 200,
			},
			{
				State:    "retry",
				Matcher:  "error",
				Argument: "",
				Expected: "NoSuchEntity",
			},
		},
	}

	w := waiter.Waiter{
		Client: c,
		Input:  input,
		Config: waiterCfg,
	}
	return w.Wait()
}
