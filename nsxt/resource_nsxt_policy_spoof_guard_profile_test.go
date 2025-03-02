/* Copyright © 2020 VMware, Inc. All Rights Reserved.
   SPDX-License-Identifier: MPL-2.0 */

package nsxt

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

var accTestPolicySpoofGuardProfileCreateAttributes = map[string]string{
	"display_name":              getAccTestResourceName(),
	"description":               "terraform created",
	"address_binding_allowlist": "true",
}

var accTestPolicySpoofGuardProfileUpdateAttributes = map[string]string{
	"display_name":              getAccTestResourceName(),
	"description":               "terraform updated",
	"address_binding_allowlist": "false",
}

func TestAccResourceNsxtPolicySpoofGuardProfile_basic(t *testing.T) {
	testResourceName := "nsxt_policy_spoof_guard_profile.test"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		CheckDestroy: func(state *terraform.State) error {
			return testAccNsxtPolicySpoofGuardProfileCheckDestroy(state, accTestPolicySpoofGuardProfileUpdateAttributes["display_name"])
		},
		Steps: []resource.TestStep{
			{
				Config: testAccNsxtPolicySpoofGuardProfileTemplate(true),
				Check: resource.ComposeTestCheckFunc(
					testAccNsxtPolicySpoofGuardProfileExists(accTestPolicySpoofGuardProfileCreateAttributes["display_name"], testResourceName),
					resource.TestCheckResourceAttr(testResourceName, "display_name", accTestPolicySpoofGuardProfileCreateAttributes["display_name"]),
					resource.TestCheckResourceAttr(testResourceName, "description", accTestPolicySpoofGuardProfileCreateAttributes["description"]),
					resource.TestCheckResourceAttr(testResourceName, "address_binding_allowlist", accTestPolicySpoofGuardProfileCreateAttributes["address_binding_allowlist"]),

					resource.TestCheckResourceAttrSet(testResourceName, "nsx_id"),
					resource.TestCheckResourceAttrSet(testResourceName, "path"),
					resource.TestCheckResourceAttrSet(testResourceName, "revision"),
					resource.TestCheckResourceAttr(testResourceName, "tag.#", "1"),
				),
			},
			{
				Config: testAccNsxtPolicySpoofGuardProfileTemplate(false),
				Check: resource.ComposeTestCheckFunc(
					testAccNsxtPolicySpoofGuardProfileExists(accTestPolicySpoofGuardProfileUpdateAttributes["display_name"], testResourceName),
					resource.TestCheckResourceAttr(testResourceName, "display_name", accTestPolicySpoofGuardProfileUpdateAttributes["display_name"]),
					resource.TestCheckResourceAttr(testResourceName, "description", accTestPolicySpoofGuardProfileUpdateAttributes["description"]),
					resource.TestCheckResourceAttr(testResourceName, "address_binding_allowlist", accTestPolicySpoofGuardProfileUpdateAttributes["address_binding_allowlist"]),

					resource.TestCheckResourceAttrSet(testResourceName, "nsx_id"),
					resource.TestCheckResourceAttrSet(testResourceName, "path"),
					resource.TestCheckResourceAttrSet(testResourceName, "revision"),
					resource.TestCheckResourceAttr(testResourceName, "tag.#", "1"),
				),
			},
			{
				Config: testAccNsxtPolicySpoofGuardProfileMinimalistic(),
				Check: resource.ComposeTestCheckFunc(
					testAccNsxtPolicySpoofGuardProfileExists(accTestPolicySpoofGuardProfileCreateAttributes["display_name"], testResourceName),
					resource.TestCheckResourceAttr(testResourceName, "description", ""),
					resource.TestCheckResourceAttrSet(testResourceName, "nsx_id"),
					resource.TestCheckResourceAttrSet(testResourceName, "path"),
					resource.TestCheckResourceAttrSet(testResourceName, "revision"),
					resource.TestCheckResourceAttr(testResourceName, "tag.#", "0"),
				),
			},
		},
	})
}

func TestAccResourceNsxtPolicySpoofGuardProfile_importBasic(t *testing.T) {
	name := getAccTestResourceName()
	testResourceName := "nsxt_policy_spoof_guard_profile.test"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		CheckDestroy: func(state *terraform.State) error {
			return testAccNsxtPolicySpoofGuardProfileCheckDestroy(state, name)
		},
		Steps: []resource.TestStep{
			{
				Config: testAccNsxtPolicySpoofGuardProfileMinimalistic(),
			},
			{
				ResourceName:      testResourceName,
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testAccNsxtPolicySpoofGuardProfileExists(displayName string, resourceName string) resource.TestCheckFunc {
	return func(state *terraform.State) error {

		connector := getPolicyConnector(testAccProvider.Meta().(nsxtClients))

		rs, ok := state.RootModule().Resources[resourceName]
		if !ok {
			return fmt.Errorf("Policy SpoofGuardProfile resource %s not found in resources", resourceName)
		}

		resourceID := rs.Primary.ID
		if resourceID == "" {
			return fmt.Errorf("Policy SpoofGuardProfile resource ID not set in resources")
		}

		exists, err := resourceNsxtPolicySpoofGuardProfileExists(resourceID, connector, testAccIsGlobalManager())
		if err != nil {
			return err
		}
		if !exists {
			return fmt.Errorf("Policy SpoofGuardProfile %s does not exist", resourceID)
		}

		return nil
	}
}

func testAccNsxtPolicySpoofGuardProfileCheckDestroy(state *terraform.State, displayName string) error {
	connector := getPolicyConnector(testAccProvider.Meta().(nsxtClients))
	for _, rs := range state.RootModule().Resources {

		if rs.Type != "nsxt_policy_spoof_guard_profile" {
			continue
		}

		resourceID := rs.Primary.Attributes["id"]
		exists, err := resourceNsxtPolicySpoofGuardProfileExists(resourceID, connector, testAccIsGlobalManager())
		if err == nil {
			return err
		}

		if exists {
			return fmt.Errorf("Policy SpoofGuardProfile %s still exists", displayName)
		}
	}
	return nil
}

func testAccNsxtPolicySpoofGuardProfileTemplate(createFlow bool) string {
	var attrMap map[string]string
	if createFlow {
		attrMap = accTestPolicySpoofGuardProfileCreateAttributes
	} else {
		attrMap = accTestPolicySpoofGuardProfileUpdateAttributes
	}
	return fmt.Sprintf(`
resource "nsxt_policy_spoof_guard_profile" "test" {
  display_name = "%s"
  description  = "%s"

  address_binding_allowlist = %s

  tag {
    scope = "scope1"
    tag   = "tag1"
  }
}`, attrMap["display_name"], attrMap["description"], attrMap["address_binding_allowlist"])
}

func testAccNsxtPolicySpoofGuardProfileMinimalistic() string {
	return fmt.Sprintf(`
resource "nsxt_policy_spoof_guard_profile" "test" {
  display_name = "%s"

}`, accTestPolicySpoofGuardProfileUpdateAttributes["display_name"])
}
