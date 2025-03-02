// Copyright © 2019-2021 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: BSD-2-Clause

// Auto generated code. DO NOT EDIT.

// Interface file for service: AlbHardwareSecurityModuleGroups
// Used by client-side stubs.

package global_infra

import (
	vapiStdErrors_ "github.com/vmware/vsphere-automation-sdk-go/lib/vapi/std/errors"
	vapiBindings_ "github.com/vmware/vsphere-automation-sdk-go/runtime/bindings"
	vapiCore_ "github.com/vmware/vsphere-automation-sdk-go/runtime/core"
	vapiProtocolClient_ "github.com/vmware/vsphere-automation-sdk-go/runtime/protocol/client"
	nsx_global_policyModel "github.com/vmware/vsphere-automation-sdk-go/services/nsxt-gm/model"
)

const _ = vapiCore_.SupportedByRuntimeVersion2

type AlbHardwareSecurityModuleGroupsClient interface {

	// Delete the ALBHardwareSecurityModuleGroup along with all the entities contained by this ALBHardwareSecurityModuleGroup. This is a deprecated API. It is recommennded to use NSX Advanced Load Balancer (Avi) Controller UI or API directly instead of NSX-T ALB Policy UI and API.
	//
	// Deprecated: This API element is deprecated.
	//
	// @param albHardwaresecuritymodulegroupIdParam ALBHardwareSecurityModuleGroup ID (required)
	// @param forceParam Force delete the resource even if it is being used somewhere (optional, default to false)
	//
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	Delete(albHardwaresecuritymodulegroupIdParam string, forceParam *bool) error

	// Read a ALBHardwareSecurityModuleGroup. This is a deprecated API. It is recommennded to use NSX Advanced Load Balancer (Avi) Controller UI or API directly instead of NSX-T ALB Policy UI and API.
	//
	// Deprecated: This API element is deprecated.
	//
	// @param albHardwaresecuritymodulegroupIdParam ALBHardwareSecurityModuleGroup ID (required)
	// @return com.vmware.nsx_global_policy.model.ALBHardwareSecurityModuleGroup
	//
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	Get(albHardwaresecuritymodulegroupIdParam string) (nsx_global_policyModel.ALBHardwareSecurityModuleGroup, error)

	// Paginated list of all ALBHardwareSecurityModuleGroup for infra. This is a deprecated API. It is recommennded to use NSX Advanced Load Balancer (Avi) Controller UI or API directly instead of NSX-T ALB Policy UI and API.
	//
	// Deprecated: This API element is deprecated.
	//
	// @param cursorParam Opaque cursor to be used for getting next page of records (supplied by current result page) (optional)
	// @param includeMarkForDeleteObjectsParam Include objects that are marked for deletion in results (optional, default to false)
	// @param includedFieldsParam Comma separated list of fields that should be included in query result (optional)
	// @param pageSizeParam Maximum number of results to return in this page (server may return fewer) (optional, default to 1000)
	// @param sortAscendingParam (optional)
	// @param sortByParam Field by which records are sorted (optional)
	// @return com.vmware.nsx_global_policy.model.ALBHardwareSecurityModuleGroupApiResponse
	//
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	List(cursorParam *string, includeMarkForDeleteObjectsParam *bool, includedFieldsParam *string, pageSizeParam *int64, sortAscendingParam *bool, sortByParam *string) (nsx_global_policyModel.ALBHardwareSecurityModuleGroupApiResponse, error)

	// If a ALBhardwaresecuritymodulegroup with the alb-hardwaresecuritymodulegroup-id is not already present, create a new ALBhardwaresecuritymodulegroup. If it already exists, update the ALBhardwaresecuritymodulegroup. This is a full replace. This is a deprecated API. It is recommennded to use NSX Advanced Load Balancer (Avi) Controller UI or API directly instead of NSX-T ALB Policy UI and API.
	//
	// Deprecated: This API element is deprecated.
	//
	// @param albHardwaresecuritymodulegroupIdParam ALBhardwaresecuritymodulegroup ID (required)
	// @param aLBHardwareSecurityModuleGroupParam (required)
	//
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	Patch(albHardwaresecuritymodulegroupIdParam string, aLBHardwareSecurityModuleGroupParam nsx_global_policyModel.ALBHardwareSecurityModuleGroup) error

	// If a ALBHardwareSecurityModuleGroup with the alb-HardwareSecurityModuleGroup-id is not already present, create a new ALBHardwareSecurityModuleGroup. If it already exists, update the ALBHardwareSecurityModuleGroup. This is a full replace. This is a deprecated API. It is recommennded to use NSX Advanced Load Balancer (Avi) Controller UI or API directly instead of NSX-T ALB Policy UI and API.
	//
	// Deprecated: This API element is deprecated.
	//
	// @param albHardwaresecuritymodulegroupIdParam ALBHardwareSecurityModuleGroup ID (required)
	// @param aLBHardwareSecurityModuleGroupParam (required)
	// @return com.vmware.nsx_global_policy.model.ALBHardwareSecurityModuleGroup
	//
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	Update(albHardwaresecuritymodulegroupIdParam string, aLBHardwareSecurityModuleGroupParam nsx_global_policyModel.ALBHardwareSecurityModuleGroup) (nsx_global_policyModel.ALBHardwareSecurityModuleGroup, error)
}

type albHardwareSecurityModuleGroupsClient struct {
	connector           vapiProtocolClient_.Connector
	interfaceDefinition vapiCore_.InterfaceDefinition
	errorsBindingMap    map[string]vapiBindings_.BindingType
}

func NewAlbHardwareSecurityModuleGroupsClient(connector vapiProtocolClient_.Connector) *albHardwareSecurityModuleGroupsClient {
	interfaceIdentifier := vapiCore_.NewInterfaceIdentifier("com.vmware.nsx_global_policy.global_infra.alb_hardware_security_module_groups")
	methodIdentifiers := map[string]vapiCore_.MethodIdentifier{
		"delete": vapiCore_.NewMethodIdentifier(interfaceIdentifier, "delete"),
		"get":    vapiCore_.NewMethodIdentifier(interfaceIdentifier, "get"),
		"list":   vapiCore_.NewMethodIdentifier(interfaceIdentifier, "list"),
		"patch":  vapiCore_.NewMethodIdentifier(interfaceIdentifier, "patch"),
		"update": vapiCore_.NewMethodIdentifier(interfaceIdentifier, "update"),
	}
	interfaceDefinition := vapiCore_.NewInterfaceDefinition(interfaceIdentifier, methodIdentifiers)
	errorsBindingMap := make(map[string]vapiBindings_.BindingType)

	aIface := albHardwareSecurityModuleGroupsClient{interfaceDefinition: interfaceDefinition, errorsBindingMap: errorsBindingMap, connector: connector}
	return &aIface
}

func (aIface *albHardwareSecurityModuleGroupsClient) GetErrorBindingType(errorName string) vapiBindings_.BindingType {
	if entry, ok := aIface.errorsBindingMap[errorName]; ok {
		return entry
	}
	return vapiStdErrors_.ERROR_BINDINGS_MAP[errorName]
}

func (aIface *albHardwareSecurityModuleGroupsClient) Delete(albHardwaresecuritymodulegroupIdParam string, forceParam *bool) error {
	typeConverter := aIface.connector.TypeConverter()
	executionContext := aIface.connector.NewExecutionContext()
	operationRestMetaData := albHardwareSecurityModuleGroupsDeleteRestMetadata()
	executionContext.SetConnectionMetadata(vapiCore_.RESTMetadataKey, operationRestMetaData)
	executionContext.SetConnectionMetadata(vapiCore_.ResponseTypeKey, vapiCore_.NewResponseType(true, false))

	sv := vapiBindings_.NewStructValueBuilder(albHardwareSecurityModuleGroupsDeleteInputType(), typeConverter)
	sv.AddStructField("AlbHardwaresecuritymodulegroupId", albHardwaresecuritymodulegroupIdParam)
	sv.AddStructField("Force", forceParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		return vapiBindings_.VAPIerrorsToError(inputError)
	}

	methodResult := aIface.connector.GetApiProvider().Invoke("com.vmware.nsx_global_policy.global_infra.alb_hardware_security_module_groups", "delete", inputDataValue, executionContext)
	if methodResult.IsSuccess() {
		return nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), aIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return vapiBindings_.VAPIerrorsToError(errorInError)
		}
		return methodError.(error)
	}
}

func (aIface *albHardwareSecurityModuleGroupsClient) Get(albHardwaresecuritymodulegroupIdParam string) (nsx_global_policyModel.ALBHardwareSecurityModuleGroup, error) {
	typeConverter := aIface.connector.TypeConverter()
	executionContext := aIface.connector.NewExecutionContext()
	operationRestMetaData := albHardwareSecurityModuleGroupsGetRestMetadata()
	executionContext.SetConnectionMetadata(vapiCore_.RESTMetadataKey, operationRestMetaData)
	executionContext.SetConnectionMetadata(vapiCore_.ResponseTypeKey, vapiCore_.NewResponseType(true, false))

	sv := vapiBindings_.NewStructValueBuilder(albHardwareSecurityModuleGroupsGetInputType(), typeConverter)
	sv.AddStructField("AlbHardwaresecuritymodulegroupId", albHardwaresecuritymodulegroupIdParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput nsx_global_policyModel.ALBHardwareSecurityModuleGroup
		return emptyOutput, vapiBindings_.VAPIerrorsToError(inputError)
	}

	methodResult := aIface.connector.GetApiProvider().Invoke("com.vmware.nsx_global_policy.global_infra.alb_hardware_security_module_groups", "get", inputDataValue, executionContext)
	var emptyOutput nsx_global_policyModel.ALBHardwareSecurityModuleGroup
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), AlbHardwareSecurityModuleGroupsGetOutputType())
		if errorInOutput != nil {
			return emptyOutput, vapiBindings_.VAPIerrorsToError(errorInOutput)
		}
		return output.(nsx_global_policyModel.ALBHardwareSecurityModuleGroup), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), aIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, vapiBindings_.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (aIface *albHardwareSecurityModuleGroupsClient) List(cursorParam *string, includeMarkForDeleteObjectsParam *bool, includedFieldsParam *string, pageSizeParam *int64, sortAscendingParam *bool, sortByParam *string) (nsx_global_policyModel.ALBHardwareSecurityModuleGroupApiResponse, error) {
	typeConverter := aIface.connector.TypeConverter()
	executionContext := aIface.connector.NewExecutionContext()
	operationRestMetaData := albHardwareSecurityModuleGroupsListRestMetadata()
	executionContext.SetConnectionMetadata(vapiCore_.RESTMetadataKey, operationRestMetaData)
	executionContext.SetConnectionMetadata(vapiCore_.ResponseTypeKey, vapiCore_.NewResponseType(true, false))

	sv := vapiBindings_.NewStructValueBuilder(albHardwareSecurityModuleGroupsListInputType(), typeConverter)
	sv.AddStructField("Cursor", cursorParam)
	sv.AddStructField("IncludeMarkForDeleteObjects", includeMarkForDeleteObjectsParam)
	sv.AddStructField("IncludedFields", includedFieldsParam)
	sv.AddStructField("PageSize", pageSizeParam)
	sv.AddStructField("SortAscending", sortAscendingParam)
	sv.AddStructField("SortBy", sortByParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput nsx_global_policyModel.ALBHardwareSecurityModuleGroupApiResponse
		return emptyOutput, vapiBindings_.VAPIerrorsToError(inputError)
	}

	methodResult := aIface.connector.GetApiProvider().Invoke("com.vmware.nsx_global_policy.global_infra.alb_hardware_security_module_groups", "list", inputDataValue, executionContext)
	var emptyOutput nsx_global_policyModel.ALBHardwareSecurityModuleGroupApiResponse
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), AlbHardwareSecurityModuleGroupsListOutputType())
		if errorInOutput != nil {
			return emptyOutput, vapiBindings_.VAPIerrorsToError(errorInOutput)
		}
		return output.(nsx_global_policyModel.ALBHardwareSecurityModuleGroupApiResponse), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), aIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, vapiBindings_.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (aIface *albHardwareSecurityModuleGroupsClient) Patch(albHardwaresecuritymodulegroupIdParam string, aLBHardwareSecurityModuleGroupParam nsx_global_policyModel.ALBHardwareSecurityModuleGroup) error {
	typeConverter := aIface.connector.TypeConverter()
	executionContext := aIface.connector.NewExecutionContext()
	operationRestMetaData := albHardwareSecurityModuleGroupsPatchRestMetadata()
	executionContext.SetConnectionMetadata(vapiCore_.RESTMetadataKey, operationRestMetaData)
	executionContext.SetConnectionMetadata(vapiCore_.ResponseTypeKey, vapiCore_.NewResponseType(true, false))

	sv := vapiBindings_.NewStructValueBuilder(albHardwareSecurityModuleGroupsPatchInputType(), typeConverter)
	sv.AddStructField("AlbHardwaresecuritymodulegroupId", albHardwaresecuritymodulegroupIdParam)
	sv.AddStructField("ALBHardwareSecurityModuleGroup", aLBHardwareSecurityModuleGroupParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		return vapiBindings_.VAPIerrorsToError(inputError)
	}

	methodResult := aIface.connector.GetApiProvider().Invoke("com.vmware.nsx_global_policy.global_infra.alb_hardware_security_module_groups", "patch", inputDataValue, executionContext)
	if methodResult.IsSuccess() {
		return nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), aIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return vapiBindings_.VAPIerrorsToError(errorInError)
		}
		return methodError.(error)
	}
}

func (aIface *albHardwareSecurityModuleGroupsClient) Update(albHardwaresecuritymodulegroupIdParam string, aLBHardwareSecurityModuleGroupParam nsx_global_policyModel.ALBHardwareSecurityModuleGroup) (nsx_global_policyModel.ALBHardwareSecurityModuleGroup, error) {
	typeConverter := aIface.connector.TypeConverter()
	executionContext := aIface.connector.NewExecutionContext()
	operationRestMetaData := albHardwareSecurityModuleGroupsUpdateRestMetadata()
	executionContext.SetConnectionMetadata(vapiCore_.RESTMetadataKey, operationRestMetaData)
	executionContext.SetConnectionMetadata(vapiCore_.ResponseTypeKey, vapiCore_.NewResponseType(true, false))

	sv := vapiBindings_.NewStructValueBuilder(albHardwareSecurityModuleGroupsUpdateInputType(), typeConverter)
	sv.AddStructField("AlbHardwaresecuritymodulegroupId", albHardwaresecuritymodulegroupIdParam)
	sv.AddStructField("ALBHardwareSecurityModuleGroup", aLBHardwareSecurityModuleGroupParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput nsx_global_policyModel.ALBHardwareSecurityModuleGroup
		return emptyOutput, vapiBindings_.VAPIerrorsToError(inputError)
	}

	methodResult := aIface.connector.GetApiProvider().Invoke("com.vmware.nsx_global_policy.global_infra.alb_hardware_security_module_groups", "update", inputDataValue, executionContext)
	var emptyOutput nsx_global_policyModel.ALBHardwareSecurityModuleGroup
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), AlbHardwareSecurityModuleGroupsUpdateOutputType())
		if errorInOutput != nil {
			return emptyOutput, vapiBindings_.VAPIerrorsToError(errorInOutput)
		}
		return output.(nsx_global_policyModel.ALBHardwareSecurityModuleGroup), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), aIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, vapiBindings_.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}
