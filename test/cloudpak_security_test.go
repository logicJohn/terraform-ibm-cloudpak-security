/**************************************************************************************************************
To Write a test file, use following link as a reference

https://github.com/terraform-ibm-modules/terraform-ibm-function/blob/main/test/cloud_function_test.go

***************************************************************************************************************/
package test

import (
	"testing"

	"github.com/gruntwork-io/terratest/modules/terraform"
)

// An example of how to test the Terraform module to create cos instance in examples/instance using Terratest.
func TestAccIBMCP4S(t *testing.T) {
	t.Parallel()

	// Construct the terraform options with default retryable errors to handle the most common retryable errors in
	// terraform testing.
	terraformOptions := terraform.WithDefaultRetryableErrors(t, &terraform.Options{
		// The path to where our Terraform code is located
		TerraformDir: "../examples/cp4s",

		// Variables to pass to our Terraform code using -var options
		Vars: map[string]interface{}{
			"cluster_id":       "",
			"resource_group":   "Default",
			"storageclass":     "ibmc-file-gold-gid",
			"entitled_registry_key":  "", //pragma: allowlist secret
			"entitled_registry_email":  "",
		},
	})

	// At the end of the test, run `terraform destroy` to clean up any resources that were created
	defer terraform.Destroy(t, terraformOptions)

	// This will run `terraform init` and `terraform apply` and fail the test if there are any errors
	terraform.InitAndApply(t, terraformOptions)

	// Run `terraform output` to get the value of an output variable
	endpoint := terraform.Output(t, terraformOptions, "endpoint")
	if len(endpoint) <= 0 {
		t.Fatal("Wrong output")
	}
	fmt.Println("Cloud Pak for Security Console URL", endpoint)
	user := terraform.Output(t, terraformOptions, "user")
	if len(user) <= 0 {
		t.Fatal("Wrong output")
	}
	fmt.Println("Cloud Pak for Security Console User ID", user)
	password := terraform.Output(t, terraformOptions, "password") //pragma: allowlist secret
	if len(password) <= 0 {
		t.Fatal("Wrong output")
	}
	fmt.Println("Cloud Pak for Security Console Password", password)
}