// Copyright © 2019-2021 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: BSD-2-Clause

// Auto generated code. DO NOT EDIT.

// Interface file for service: AlbDnsPolicies
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

type AlbDnsPoliciesClient interface {

	// Delete the ALBDnsPolicy along with all the entities contained by this ALBDnsPolicy. This is a deprecated API. It is recommennded to use NSX Advanced Load Balancer (Avi) Controller UI or API directly instead of NSX-T ALB Policy UI and API.
	//
	// Deprecated: This API element is deprecated.
	//
	// @param albDnspolicyIdParam ALBDnsPolicy ID (required)
	// @param forceParam Force delete the resource even if it is being used somewhere (optional, default to false)
	//
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	Delete(albDnspolicyIdParam string, forceParam *bool) error

	// Read a ALBDnsPolicy. This is a deprecated API. It is recommennded to use NSX Advanced Load Balancer (Avi) Controller UI or API directly instead of NSX-T ALB Policy UI and API.
	//
	// Deprecated: This API element is deprecated.
	//
	// @param albDnspolicyIdParam ALBDnsPolicy ID (required)
	// @return com.vmware.nsx_global_policy.model.ALBDnsPolicy
	//
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	Get(albDnspolicyIdParam string) (nsx_global_policyModel.ALBDnsPolicy, error)

	// Paginated list of all ALBDnsPolicy for infra. This is a deprecated API. It is recommennded to use NSX Advanced Load Balancer (Avi) Controller UI or API directly instead of NSX-T ALB Policy UI and API.
	//
	// Deprecated: This API element is deprecated.
	//
	// @param cursorParam Opaque cursor to be used for getting next page of records (supplied by current result page) (optional)
	// @param includeMarkForDeleteObjectsParam Include objects that are marked for deletion in results (optional, default to false)
	// @param includedFieldsParam Comma separated list of fields that should be included in query result (optional)
	// @param pageSizeParam Maximum number of results to return in this page (server may return fewer) (optional, default to 1000)
	// @param sortAscendingParam (optional)
	// @param sortByParam Field by which records are sorted (optional)
	// @return com.vmware.nsx_global_policy.model.ALBDnsPolicyApiResponse
	//
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	List(cursorParam *string, includeMarkForDeleteObjectsParam *bool, includedFieldsParam *string, pageSizeParam *int64, sortAscendingParam *bool, sortByParam *string) (nsx_global_policyModel.ALBDnsPolicyApiResponse, error)

	// If a ALBdnspolicy with the alb-dnspolicy-id is not already present, create a new ALBdnspolicy. If it already exists, update the ALBdnspolicy. This is a full replace. This is a deprecated API. It is recommennded to use NSX Advanced Load Balancer (Avi) Controller UI or API directly instead of NSX-T ALB Policy UI and API.
	//
	// Deprecated: This API element is deprecated.
	//
	// @param albDnspolicyIdParam ALBdnspolicy ID (required)
	// @param aLBDnsPolicyParam (required)
	//
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	Patch(albDnspolicyIdParam string, aLBDnsPolicyParam nsx_global_policyModel.ALBDnsPolicy) error

	// If a ALBDnsPolicy with the alb-DnsPolicy-id is not already present, create a new ALBDnsPolicy. If it already exists, update the ALBDnsPolicy. This is a full replace. This is a deprecated API. It is recommennded to use NSX Advanced Load Balancer (Avi) Controller UI or API directly instead of NSX-T ALB Policy UI and API.
	//
	// Deprecated: This API element is deprecated.
	//
	// @param albDnspolicyIdParam ALBDnsPolicy ID (required)
	// @param aLBDnsPolicyParam (required)
	// @return com.vmware.nsx_global_policy.model.ALBDnsPolicy
	//
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	Update(albDnspolicyIdParam string, aLBDnsPolicyParam nsx_global_policyModel.ALBDnsPolicy) (nsx_global_policyModel.ALBDnsPolicy, error)
}

type albDnsPoliciesClient struct {
	connector           vapiProtocolClient_.Connector
	interfaceDefinition vapiCore_.InterfaceDefinition
	errorsBindingMap    map[string]vapiBindings_.BindingType
}

func NewAlbDnsPoliciesClient(connector vapiProtocolClient_.Connector) *albDnsPoliciesClient {
	interfaceIdentifier := vapiCore_.NewInterfaceIdentifier("com.vmware.nsx_global_policy.global_infra.alb_dns_policies")
	methodIdentifiers := map[string]vapiCore_.MethodIdentifier{
		"delete": vapiCore_.NewMethodIdentifier(interfaceIdentifier, "delete"),
		"get":    vapiCore_.NewMethodIdentifier(interfaceIdentifier, "get"),
		"list":   vapiCore_.NewMethodIdentifier(interfaceIdentifier, "list"),
		"patch":  vapiCore_.NewMethodIdentifier(interfaceIdentifier, "patch"),
		"update": vapiCore_.NewMethodIdentifier(interfaceIdentifier, "update"),
	}
	interfaceDefinition := vapiCore_.NewInterfaceDefinition(interfaceIdentifier, methodIdentifiers)
	errorsBindingMap := make(map[string]vapiBindings_.BindingType)

	aIface := albDnsPoliciesClient{interfaceDefinition: interfaceDefinition, errorsBindingMap: errorsBindingMap, connector: connector}
	return &aIface
}

func (aIface *albDnsPoliciesClient) GetErrorBindingType(errorName string) vapiBindings_.BindingType {
	if entry, ok := aIface.errorsBindingMap[errorName]; ok {
		return entry
	}
	return vapiStdErrors_.ERROR_BINDINGS_MAP[errorName]
}

func (aIface *albDnsPoliciesClient) Delete(albDnspolicyIdParam string, forceParam *bool) error {
	typeConverter := aIface.connector.TypeConverter()
	executionContext := aIface.connector.NewExecutionContext()
	operationRestMetaData := albDnsPoliciesDeleteRestMetadata()
	executionContext.SetConnectionMetadata(vapiCore_.RESTMetadataKey, operationRestMetaData)
	executionContext.SetConnectionMetadata(vapiCore_.ResponseTypeKey, vapiCore_.NewResponseType(true, false))

	sv := vapiBindings_.NewStructValueBuilder(albDnsPoliciesDeleteInputType(), typeConverter)
	sv.AddStructField("AlbDnspolicyId", albDnspolicyIdParam)
	sv.AddStructField("Force", forceParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		return vapiBindings_.VAPIerrorsToError(inputError)
	}

	methodResult := aIface.connector.GetApiProvider().Invoke("com.vmware.nsx_global_policy.global_infra.alb_dns_policies", "delete", inputDataValue, executionContext)
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

func (aIface *albDnsPoliciesClient) Get(albDnspolicyIdParam string) (nsx_global_policyModel.ALBDnsPolicy, error) {
	typeConverter := aIface.connector.TypeConverter()
	executionContext := aIface.connector.NewExecutionContext()
	operationRestMetaData := albDnsPoliciesGetRestMetadata()
	executionContext.SetConnectionMetadata(vapiCore_.RESTMetadataKey, operationRestMetaData)
	executionContext.SetConnectionMetadata(vapiCore_.ResponseTypeKey, vapiCore_.NewResponseType(true, false))

	sv := vapiBindings_.NewStructValueBuilder(albDnsPoliciesGetInputType(), typeConverter)
	sv.AddStructField("AlbDnspolicyId", albDnspolicyIdParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput nsx_global_policyModel.ALBDnsPolicy
		return emptyOutput, vapiBindings_.VAPIerrorsToError(inputError)
	}

	methodResult := aIface.connector.GetApiProvider().Invoke("com.vmware.nsx_global_policy.global_infra.alb_dns_policies", "get", inputDataValue, executionContext)
	var emptyOutput nsx_global_policyModel.ALBDnsPolicy
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), AlbDnsPoliciesGetOutputType())
		if errorInOutput != nil {
			return emptyOutput, vapiBindings_.VAPIerrorsToError(errorInOutput)
		}
		return output.(nsx_global_policyModel.ALBDnsPolicy), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), aIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, vapiBindings_.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (aIface *albDnsPoliciesClient) List(cursorParam *string, includeMarkForDeleteObjectsParam *bool, includedFieldsParam *string, pageSizeParam *int64, sortAscendingParam *bool, sortByParam *string) (nsx_global_policyModel.ALBDnsPolicyApiResponse, error) {
	typeConverter := aIface.connector.TypeConverter()
	executionContext := aIface.connector.NewExecutionContext()
	operationRestMetaData := albDnsPoliciesListRestMetadata()
	executionContext.SetConnectionMetadata(vapiCore_.RESTMetadataKey, operationRestMetaData)
	executionContext.SetConnectionMetadata(vapiCore_.ResponseTypeKey, vapiCore_.NewResponseType(true, false))

	sv := vapiBindings_.NewStructValueBuilder(albDnsPoliciesListInputType(), typeConverter)
	sv.AddStructField("Cursor", cursorParam)
	sv.AddStructField("IncludeMarkForDeleteObjects", includeMarkForDeleteObjectsParam)
	sv.AddStructField("IncludedFields", includedFieldsParam)
	sv.AddStructField("PageSize", pageSizeParam)
	sv.AddStructField("SortAscending", sortAscendingParam)
	sv.AddStructField("SortBy", sortByParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput nsx_global_policyModel.ALBDnsPolicyApiResponse
		return emptyOutput, vapiBindings_.VAPIerrorsToError(inputError)
	}

	methodResult := aIface.connector.GetApiProvider().Invoke("com.vmware.nsx_global_policy.global_infra.alb_dns_policies", "list", inputDataValue, executionContext)
	var emptyOutput nsx_global_policyModel.ALBDnsPolicyApiResponse
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), AlbDnsPoliciesListOutputType())
		if errorInOutput != nil {
			return emptyOutput, vapiBindings_.VAPIerrorsToError(errorInOutput)
		}
		return output.(nsx_global_policyModel.ALBDnsPolicyApiResponse), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), aIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, vapiBindings_.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (aIface *albDnsPoliciesClient) Patch(albDnspolicyIdParam string, aLBDnsPolicyParam nsx_global_policyModel.ALBDnsPolicy) error {
	typeConverter := aIface.connector.TypeConverter()
	executionContext := aIface.connector.NewExecutionContext()
	operationRestMetaData := albDnsPoliciesPatchRestMetadata()
	executionContext.SetConnectionMetadata(vapiCore_.RESTMetadataKey, operationRestMetaData)
	executionContext.SetConnectionMetadata(vapiCore_.ResponseTypeKey, vapiCore_.NewResponseType(true, false))

	sv := vapiBindings_.NewStructValueBuilder(albDnsPoliciesPatchInputType(), typeConverter)
	sv.AddStructField("AlbDnspolicyId", albDnspolicyIdParam)
	sv.AddStructField("ALBDnsPolicy", aLBDnsPolicyParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		return vapiBindings_.VAPIerrorsToError(inputError)
	}

	methodResult := aIface.connector.GetApiProvider().Invoke("com.vmware.nsx_global_policy.global_infra.alb_dns_policies", "patch", inputDataValue, executionContext)
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

func (aIface *albDnsPoliciesClient) Update(albDnspolicyIdParam string, aLBDnsPolicyParam nsx_global_policyModel.ALBDnsPolicy) (nsx_global_policyModel.ALBDnsPolicy, error) {
	typeConverter := aIface.connector.TypeConverter()
	executionContext := aIface.connector.NewExecutionContext()
	operationRestMetaData := albDnsPoliciesUpdateRestMetadata()
	executionContext.SetConnectionMetadata(vapiCore_.RESTMetadataKey, operationRestMetaData)
	executionContext.SetConnectionMetadata(vapiCore_.ResponseTypeKey, vapiCore_.NewResponseType(true, false))

	sv := vapiBindings_.NewStructValueBuilder(albDnsPoliciesUpdateInputType(), typeConverter)
	sv.AddStructField("AlbDnspolicyId", albDnspolicyIdParam)
	sv.AddStructField("ALBDnsPolicy", aLBDnsPolicyParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput nsx_global_policyModel.ALBDnsPolicy
		return emptyOutput, vapiBindings_.VAPIerrorsToError(inputError)
	}

	methodResult := aIface.connector.GetApiProvider().Invoke("com.vmware.nsx_global_policy.global_infra.alb_dns_policies", "update", inputDataValue, executionContext)
	var emptyOutput nsx_global_policyModel.ALBDnsPolicy
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), AlbDnsPoliciesUpdateOutputType())
		if errorInOutput != nil {
			return emptyOutput, vapiBindings_.VAPIerrorsToError(errorInOutput)
		}
		return output.(nsx_global_policyModel.ALBDnsPolicy), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), aIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, vapiBindings_.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}
