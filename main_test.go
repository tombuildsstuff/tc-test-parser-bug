package main

import (
	"fmt"
	"log"
	"os"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
)

func TestPassing(t *testing.T) {
	os.Stderr.WriteString("stderr - from the passing test\n")
	log.Printf("stdout - from the passing test")
	t.Log("This should always pass")
}

func TestSkipped(t *testing.T) {
	os.Stderr.WriteString("stderr - from the skipped test\n")
	log.Printf("stdout - from the skipped test")
	t.Log("This should be skipped")
	t.Skip("Skipping because we should skip it")
}

func TestFailed(t *testing.T) {
	os.Stderr.WriteString("stderr - from the failed test\n")
	log.Printf("stdout - from the failed test")
	t.Log("This should always fail")
	t.Fail()
}

func TestTerraformRunner(t *testing.T) {
	var print = func(words string, instance int) {
		os.Stderr.WriteString(fmt.Sprintf("stderr - %d - %s\n", instance, words))
		log.Printf("stdout - %d, %s", instance, words)
	}

	var runner = resource.Test
	if os.Getenv("PARALLEL") != "" {
		runner = resource.ParallelTest
	}

	for i := 0; i < 5; i++ {
		t.Run(fmt.Sprintf("Instance %d", i+1), func(t *testing.T) {
			number := i+1
			runner(t, resource.TestCase{
				PreCheck:     func() {
					print("precheck", number)
				},
				Providers:    map[string]terraform.ResourceProvider{
					"rick": &schema.Provider{
						DataSourcesMap: map[string]*schema.Resource{
							"rick_and_morty": {
								Read: func(data *schema.ResourceData, i interface{}) error {
									data.Set("test", "bar")
									return nil
								},
								Schema: map[string]*schema.Schema{
									"test": {
										Type:  schema.TypeString,
										Computed:  true,
									},
								},
							},
						},
						ResourcesMap: make(map[string]*schema.Resource),
					},
				},
				CheckDestroy: func(state *terraform.State) error {
					print("destroy", number)
					return nil
				},
				Steps: []resource.TestStep{
					{
						Config: `data "rick_and_morty" "test" { }`,
						Check: resource.ComposeTestCheckFunc(
							func(state *terraform.State) error {
								print("check func", number)
								return nil
							},
						),
					},
				},
			})
		})
	}
}