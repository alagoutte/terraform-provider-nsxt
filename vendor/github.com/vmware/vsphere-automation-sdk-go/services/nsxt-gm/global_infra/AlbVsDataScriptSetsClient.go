// Copyright © 2019-2021 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: BSD-2-Clause

// Auto generated code. DO NOT EDIT.

// Interface file for service: AlbVsDataScriptSets
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

type AlbVsDataScriptSetsClient interface {

	// Delete the ALBVSDataScriptSet along with all the entities contained by this ALBVSDataScriptSet. This is a deprecated API. It is recommennded to use NSX Advanced Load Balancer (Avi) Controller UI or API directly instead of NSX-T ALB Policy UI and API.
	//
	// Deprecated: This API element is deprecated.
	//
	// @param albVsdatascriptsetIdParam ALBVSDataScriptSet ID (required)
	// @param forceParam Force delete the resource even if it is being used somewhere (optional, default to false)
	//
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	Delete(albVsdatascriptsetIdParam string, forceParam *bool) error

	// Read a ALBVSDataScriptSet. This is a deprecated API. It is recommennded to use NSX Advanced Load Balancer (Avi) Controller UI or API directly instead of NSX-T ALB Policy UI and API.
	//
	// Deprecated: This API element is deprecated.
	//
	// @param albVsdatascriptsetIdParam ALBVSDataScriptSet ID (required)
	// @return com.vmware.nsx_global_policy.model.ALBVSDataScriptSet
	//
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	Get(albVsdatascriptsetIdParam string) (nsx_global_policyModel.ALBVSDataScriptSet, error)

	// Paginated list of all ALBVSDataScriptSet for infra. This is a deprecated API. It is recommennded to use NSX Advanced Load Balancer (Avi) Controller UI or API directly instead of NSX-T ALB Policy UI and API.
	//
	// Deprecated: This API element is deprecated.
	//
	// @param cursorParam Opaque cursor to be used for getting next page of records (supplied by current result page) (optional)
	// @param includeMarkForDeleteObjectsParam Include objects that are marked for deletion in results (optional, default to false)
	// @param includedFieldsParam Comma separated list of fields that should be included in query result (optional)
	// @param pageSizeParam Maximum number of results to return in this page (server may return fewer) (optional, default to 1000)
	// @param sortAscendingParam (optional)
	// @param sortByParam Field by which records are sorted (optional)
	// @return com.vmware.nsx_global_policy.model.ALBVSDataScriptSetApiResponse
	//
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	List(cursorParam *string, includeMarkForDeleteObjectsParam *bool, includedFieldsParam *string, pageSizeParam *int64, sortAscendingParam *bool, sortByParam *string) (nsx_global_policyModel.ALBVSDataScriptSetApiResponse, error)

	// If a ALBvsdatascriptset with the alb-vsdatascriptset-id is not already present, create a new ALBvsdatascriptset. If it already exists, update the ALBvsdatascriptset. This is a full replace. This is a deprecated API. It is recommennded to use NSX Advanced Load Balancer (Avi) Controller UI or API directly instead of NSX-T ALB Policy UI and API.
	//
	// Deprecated: This API element is deprecated.
	//
	// @param albVsdatascriptsetIdParam ALBvsdatascriptset ID (required)
	// @param aLBVSDataScriptSetParam (required)
	//
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	Patch(albVsdatascriptsetIdParam string, aLBVSDataScriptSetParam nsx_global_policyModel.ALBVSDataScriptSet) error

	// If a ALBVSDataScriptSet with the alb-VSDataScriptSet-id is not already present, create a new ALBVSDataScriptSet. If it already exists, update the ALBVSDataScriptSet. This is a full replace. This is a deprecated API. It is recommennded to use NSX Advanced Load Balancer (Avi) Controller UI or API directly instead of NSX-T ALB Policy UI and API.
	//
	// Deprecated: This API element is deprecated.
	//
	// @param albVsdatascriptsetIdParam ALBVSDataScriptSet ID (required)
	// @param aLBVSDataScriptSetParam (required)
	// @return com.vmware.nsx_global_policy.model.ALBVSDataScriptSet
	//
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	Update(albVsdatascriptsetIdParam string, aLBVSDataScriptSetParam nsx_global_policyModel.ALBVSDataScriptSet) (nsx_global_policyModel.ALBVSDataScriptSet, error)
}

type albVsDataScriptSetsClient struct {
	connector           vapiProtocolClient_.Connector
	interfaceDefinition vapiCore_.InterfaceDefinition
	errorsBindingMap    map[string]vapiBindings_.BindingType
}

func NewAlbVsDataScriptSetsClient(connector vapiProtocolClient_.Connector) *albVsDataScriptSetsClient {
	interfaceIdentifier := vapiCore_.NewInterfaceIdentifier("com.vmware.nsx_global_policy.global_infra.alb_vs_data_script_sets")
	methodIdentifiers := map[string]vapiCore_.MethodIdentifier{
		"delete": vapiCore_.NewMethodIdentifier(interfaceIdentifier, "delete"),
		"get":    vapiCore_.NewMethodIdentifier(interfaceIdentifier, "get"),
		"list":   vapiCore_.NewMethodIdentifier(interfaceIdentifier, "list"),
		"patch":  vapiCore_.NewMethodIdentifier(interfaceIdentifier, "patch"),
		"update": vapiCore_.NewMethodIdentifier(interfaceIdentifier, "update"),
	}
	interfaceDefinition := vapiCore_.NewInterfaceDefinition(interfaceIdentifier, methodIdentifiers)
	errorsBindingMap := make(map[string]vapiBindings_.BindingType)

	aIface := albVsDataScriptSetsClient{interfaceDefinition: interfaceDefinition, errorsBindingMap: errorsBindingMap, connector: connector}
	return &aIface
}

func (aIface *albVsDataScriptSetsClient) GetErrorBindingType(errorName string) vapiBindings_.BindingType {
	if entry, ok := aIface.errorsBindingMap[errorName]; ok {
		return entry
	}
	return vapiStdErrors_.ERROR_BINDINGS_MAP[errorName]
}

func (aIface *albVsDataScriptSetsClient) Delete(albVsdatascriptsetIdParam string, forceParam *bool) error {
	typeConverter := aIface.connector.TypeConverter()
	executionContext := aIface.connector.NewExecutionContext()
	operationRestMetaData := albVsDataScriptSetsDeleteRestMetadata()
	executionContext.SetConnectionMetadata(vapiCore_.RESTMetadataKey, operationRestMetaData)
	executionContext.SetConnectionMetadata(vapiCore_.ResponseTypeKey, vapiCore_.NewResponseType(true, false))

	sv := vapiBindings_.NewStructValueBuilder(albVsDataScriptSetsDeleteInputType(), typeConverter)
	sv.AddStructField("AlbVsdatascriptsetId", albVsdatascriptsetIdParam)
	sv.AddStructField("Force", forceParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		return vapiBindings_.VAPIerrorsToError(inputError)
	}

	methodResult := aIface.connector.GetApiProvider().Invoke("com.vmware.nsx_global_policy.global_infra.alb_vs_data_script_sets", "delete", inputDataValue, executionContext)
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

func (aIface *albVsDataScriptSetsClient) Get(albVsdatascriptsetIdParam string) (nsx_global_policyModel.ALBVSDataScriptSet, error) {
	typeConverter := aIface.connector.TypeConverter()
	executionContext := aIface.connector.NewExecutionContext()
	operationRestMetaData := albVsDataScriptSetsGetRestMetadata()
	executionContext.SetConnectionMetadata(vapiCore_.RESTMetadataKey, operationRestMetaData)
	executionContext.SetConnectionMetadata(vapiCore_.ResponseTypeKey, vapiCore_.NewResponseType(true, false))

	sv := vapiBindings_.NewStructValueBuilder(albVsDataScriptSetsGetInputType(), typeConverter)
	sv.AddStructField("AlbVsdatascriptsetId", albVsdatascriptsetIdParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput nsx_global_policyModel.ALBVSDataScriptSet
		return emptyOutput, vapiBindings_.VAPIerrorsToError(inputError)
	}

	methodResult := aIface.connector.GetApiProvider().Invoke("com.vmware.nsx_global_policy.global_infra.alb_vs_data_script_sets", "get", inputDataValue, executionContext)
	var emptyOutput nsx_global_policyModel.ALBVSDataScriptSet
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), AlbVsDataScriptSetsGetOutputType())
		if errorInOutput != nil {
			return emptyOutput, vapiBindings_.VAPIerrorsToError(errorInOutput)
		}
		return output.(nsx_global_policyModel.ALBVSDataScriptSet), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), aIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, vapiBindings_.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (aIface *albVsDataScriptSetsClient) List(cursorParam *string, includeMarkForDeleteObjectsParam *bool, includedFieldsParam *string, pageSizeParam *int64, sortAscendingParam *bool, sortByParam *string) (nsx_global_policyModel.ALBVSDataScriptSetApiResponse, error) {
	typeConverter := aIface.connector.TypeConverter()
	executionContext := aIface.connector.NewExecutionContext()
	operationRestMetaData := albVsDataScriptSetsListRestMetadata()
	executionContext.SetConnectionMetadata(vapiCore_.RESTMetadataKey, operationRestMetaData)
	executionContext.SetConnectionMetadata(vapiCore_.ResponseTypeKey, vapiCore_.NewResponseType(true, false))

	sv := vapiBindings_.NewStructValueBuilder(albVsDataScriptSetsListInputType(), typeConverter)
	sv.AddStructField("Cursor", cursorParam)
	sv.AddStructField("IncludeMarkForDeleteObjects", includeMarkForDeleteObjectsParam)
	sv.AddStructField("IncludedFields", includedFieldsParam)
	sv.AddStructField("PageSize", pageSizeParam)
	sv.AddStructField("SortAscending", sortAscendingParam)
	sv.AddStructField("SortBy", sortByParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput nsx_global_policyModel.ALBVSDataScriptSetApiResponse
		return emptyOutput, vapiBindings_.VAPIerrorsToError(inputError)
	}

	methodResult := aIface.connector.GetApiProvider().Invoke("com.vmware.nsx_global_policy.global_infra.alb_vs_data_script_sets", "list", inputDataValue, executionContext)
	var emptyOutput nsx_global_policyModel.ALBVSDataScriptSetApiResponse
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), AlbVsDataScriptSetsListOutputType())
		if errorInOutput != nil {
			return emptyOutput, vapiBindings_.VAPIerrorsToError(errorInOutput)
		}
		return output.(nsx_global_policyModel.ALBVSDataScriptSetApiResponse), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), aIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, vapiBindings_.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (aIface *albVsDataScriptSetsClient) Patch(albVsdatascriptsetIdParam string, aLBVSDataScriptSetParam nsx_global_policyModel.ALBVSDataScriptSet) error {
	typeConverter := aIface.connector.TypeConverter()
	executionContext := aIface.connector.NewExecutionContext()
	operationRestMetaData := albVsDataScriptSetsPatchRestMetadata()
	executionContext.SetConnectionMetadata(vapiCore_.RESTMetadataKey, operationRestMetaData)
	executionContext.SetConnectionMetadata(vapiCore_.ResponseTypeKey, vapiCore_.NewResponseType(true, false))

	sv := vapiBindings_.NewStructValueBuilder(albVsDataScriptSetsPatchInputType(), typeConverter)
	sv.AddStructField("AlbVsdatascriptsetId", albVsdatascriptsetIdParam)
	sv.AddStructField("ALBVSDataScriptSet", aLBVSDataScriptSetParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		return vapiBindings_.VAPIerrorsToError(inputError)
	}

	methodResult := aIface.connector.GetApiProvider().Invoke("com.vmware.nsx_global_policy.global_infra.alb_vs_data_script_sets", "patch", inputDataValue, executionContext)
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

func (aIface *albVsDataScriptSetsClient) Update(albVsdatascriptsetIdParam string, aLBVSDataScriptSetParam nsx_global_policyModel.ALBVSDataScriptSet) (nsx_global_policyModel.ALBVSDataScriptSet, error) {
	typeConverter := aIface.connector.TypeConverter()
	executionContext := aIface.connector.NewExecutionContext()
	operationRestMetaData := albVsDataScriptSetsUpdateRestMetadata()
	executionContext.SetConnectionMetadata(vapiCore_.RESTMetadataKey, operationRestMetaData)
	executionContext.SetConnectionMetadata(vapiCore_.ResponseTypeKey, vapiCore_.NewResponseType(true, false))

	sv := vapiBindings_.NewStructValueBuilder(albVsDataScriptSetsUpdateInputType(), typeConverter)
	sv.AddStructField("AlbVsdatascriptsetId", albVsdatascriptsetIdParam)
	sv.AddStructField("ALBVSDataScriptSet", aLBVSDataScriptSetParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput nsx_global_policyModel.ALBVSDataScriptSet
		return emptyOutput, vapiBindings_.VAPIerrorsToError(inputError)
	}

	methodResult := aIface.connector.GetApiProvider().Invoke("com.vmware.nsx_global_policy.global_infra.alb_vs_data_script_sets", "update", inputDataValue, executionContext)
	var emptyOutput nsx_global_policyModel.ALBVSDataScriptSet
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), AlbVsDataScriptSetsUpdateOutputType())
		if errorInOutput != nil {
			return emptyOutput, vapiBindings_.VAPIerrorsToError(errorInOutput)
		}
		return output.(nsx_global_policyModel.ALBVSDataScriptSet), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), aIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, vapiBindings_.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}
