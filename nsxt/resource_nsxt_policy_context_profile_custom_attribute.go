/* Copyright © 2023 VMware, Inc. All Rights Reserved.
   SPDX-License-Identifier: MPL-2.0 */

package nsxt

import (
	"fmt"
	"log"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
	"github.com/vmware/vsphere-automation-sdk-go/lib/vapi/std/errors"
	"github.com/vmware/vsphere-automation-sdk-go/runtime/protocol/client"
	gm_infra "github.com/vmware/vsphere-automation-sdk-go/services/nsxt-gm/global_infra/context_profiles/custom_attributes"
	gm_model "github.com/vmware/vsphere-automation-sdk-go/services/nsxt-gm/model"
	infra "github.com/vmware/vsphere-automation-sdk-go/services/nsxt/infra/context_profiles/custom_attributes"
	"github.com/vmware/vsphere-automation-sdk-go/services/nsxt/model"
)

var customAttributeKeys = []string{
	model.PolicyCustomAttributes_KEY_DOMAIN_NAME,
	model.PolicyCustomAttributes_KEY_CUSTOM_URL,
}

func splitCustomAttributeID(id string) (string, string) {
	s := strings.Split(id, "~")
	return s[0], s[1]
}

func makeCustomAttributeID(key string, attribute string) string {
	return fmt.Sprintf("%s~%s", key, attribute)
}

func resourceNsxtPolicyContextProfileCustomAttribute() *schema.Resource {
	return &schema.Resource{
		Create: resourceNsxtPolicyContextProfileCustomAttributeCreate,
		Read:   resourceNsxtPolicyContextProfileCustomAttributeRead,
		Delete: resourceNsxtPolicyContextProfileCustomAttributeDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"key": {
				Type:         schema.TypeString,
				Description:  "Key for attribute",
				Required:     true,
				ForceNew:     true,
				ValidateFunc: validation.StringInSlice(customAttributeKeys, false),
			},
			"attribute": {
				Type:        schema.TypeString,
				Description: "Custom Attribute",
				Required:    true,
				ForceNew:    true,
			},
		},
	}
}

func resourceNsxtPolicyContextProfileCustomAttributeExists(id string, connector client.Connector, isGlobalmodel bool) (bool, error) {
	var err error
	var attrList model.PolicyContextProfileListResult

	key, attribute := splitCustomAttributeID(id)
	source := model.PolicyCustomAttributes_ATTRIBUTE_SOURCE_CUSTOM
	if isGlobalmodel {
		var gmAttrList gm_model.PolicyContextProfileListResult
		client := gm_infra.NewDefaultClient(connector)
		gmAttrList, err = client.List(&key, &source, nil, nil, nil, nil, nil, nil)

		al, err1 := convertModelBindingType(gmAttrList, gm_model.PolicyContextProfileListResultBindingType(), model.PolicyContextProfileListResultBindingType())
		if err1 != nil {
			return false, err1
		}

		attrList = al.(model.PolicyContextProfileListResult)
	} else {
		client := infra.NewDefaultClient(connector)
		attrList, err = client.List(&key, &source, nil, nil, nil, nil, nil, nil)
		if err != nil {
			return false, err
		}
	}

	if *attrList.ResultCount == 0 {
		return false, nil
	}
	if isNotFoundError(err) {
		return false, nil
	}

	if err == nil {
		for _, c := range attrList.Results {
			for _, a := range c.Attributes {
				if *a.Key == key {
					for _, v := range a.Value {
						if v == attribute {
							return true, nil
						}
					}
				}
			}
		}
	}

	return false, err
}

func resourceNsxtPolicyContextProfileCustomAttributeRead(d *schema.ResourceData, m interface{}) error {
	id := d.Id()
	key, attribute := splitCustomAttributeID(id)

	log.Printf("[INFO] Reading ContextProfileCustomAttribute with ID %s", d.Id())

	connector := getPolicyConnector(m)
	exists, err := resourceNsxtPolicyContextProfileCustomAttributeExists(id, connector, isPolicyGlobalManager(m))
	if err != nil {
		return err
	}
	if !exists {
		return errors.NotFound{}
	}
	d.Set("key", key)
	d.Set("attribute", attribute)
	return nil
}

func resourceNsxtPolicyContextProfileCustomAttributeCreate(d *schema.ResourceData, m interface{}) error {
	var err error
	key := d.Get("key").(string)
	attribute := d.Get("attribute").(string)
	attributes := []string{attribute}
	log.Printf("[INFO] Creating ContextProfileCustomAttribute with ID %s", attribute)

	connector := getPolicyConnector(m)

	dataTypeString := model.PolicyCustomAttributes_DATATYPE_STRING
	obj := model.PolicyCustomAttributes{
		Datatype: &dataTypeString,
		Key:      &key,
		Value:    attributes,
	}

	// PATCH the resource
	if isPolicyGlobalManager(m) {
		gmObj, err1 := convertModelBindingType(obj, model.PolicyCustomAttributesBindingType(), gm_model.PolicyCustomAttributesBindingType())
		if err1 != nil {
			return err1
		}

		client := gm_infra.NewDefaultClient(connector)
		err = client.Create(gmObj.(gm_model.PolicyCustomAttributes), "add")
	} else {
		client := infra.NewDefaultClient(connector)
		err = client.Create(obj, "add")
	}
	if err != nil {
		return handleCreateError("ContextProfileCustomAttribute", attribute, err)
	}

	d.Set("key", key)
	d.Set("attribute", attribute)
	d.SetId(makeCustomAttributeID(key, attribute))

	return resourceNsxtPolicyContextProfileCustomAttributeRead(d, m)
}

func resourceNsxtPolicyContextProfileCustomAttributeDelete(d *schema.ResourceData, m interface{}) error {
	key, attribute := splitCustomAttributeID(d.Id())
	log.Printf("[INFO] Deleting ContextProfileCustomAttribute with ID %s", attribute)
	attributes := []string{attribute}
	err := resourceNsxtPolicyContextProfileCustomAttributeRead(d, m)

	if err != nil {
		return err
	}

	connector := getPolicyConnector(m)

	dataTypeString := model.PolicyCustomAttributes_DATATYPE_STRING
	obj := model.PolicyCustomAttributes{
		Datatype: &dataTypeString,
		Key:      &key,
		Value:    attributes,
	}

	// PATCH the resource
	if isPolicyGlobalManager(m) {
		gmObj, err1 := convertModelBindingType(obj, model.PolicyCustomAttributesBindingType(), gm_model.PolicyCustomAttributesBindingType())
		if err1 != nil {
			return err1
		}

		client := gm_infra.NewDefaultClient(connector)
		err = client.Create(gmObj.(gm_model.PolicyCustomAttributes), "remove")
	} else {
		client := infra.NewDefaultClient(connector)
		err = client.Create(obj, "remove")
	}

	if err != nil {
		return handleDeleteError("ContextProfileCustomAttribute", attribute, err)
	}
	return err
}
