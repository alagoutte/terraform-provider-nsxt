// Copyright © 2019-2021 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: BSD-2-Clause

// Auto generated code. DO NOT EDIT.

// Data type definitions file for service: TlsInspectionConfigProfileBindings.
// Includes binding types of a structures and enumerations defined in the service.
// Shared by client-side stubs and server-side skeletons to ensure type
// compatibility.

package tier_1s

import (
	vapiBindings_ "github.com/vmware/vsphere-automation-sdk-go/runtime/bindings"
	vapiData_ "github.com/vmware/vsphere-automation-sdk-go/runtime/data"
	vapiProtocol_ "github.com/vmware/vsphere-automation-sdk-go/runtime/protocol"
	nsx_global_policyModel "github.com/vmware/vsphere-automation-sdk-go/services/nsxt-gm/model"
	"reflect"
)

func tlsInspectionConfigProfileBindingsDeleteInputType() vapiBindings_.StructType {
	fields := make(map[string]vapiBindings_.BindingType)
	fieldNameMap := make(map[string]string)
	fields["tier1_id"] = vapiBindings_.NewStringType()
	fields["tls_inspection_config_profile_binding_id"] = vapiBindings_.NewStringType()
	fieldNameMap["tier1_id"] = "Tier1Id"
	fieldNameMap["tls_inspection_config_profile_binding_id"] = "TlsInspectionConfigProfileBindingId"
	var validators = []vapiBindings_.Validator{}
	return vapiBindings_.NewStructType("operation-input", fields, reflect.TypeOf(vapiData_.StructValue{}), fieldNameMap, validators)
}

func TlsInspectionConfigProfileBindingsDeleteOutputType() vapiBindings_.BindingType {
	return vapiBindings_.NewVoidType()
}

func tlsInspectionConfigProfileBindingsDeleteRestMetadata() vapiProtocol_.OperationRestMetadata {
	fields := map[string]vapiBindings_.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]vapiBindings_.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["tier1_id"] = vapiBindings_.NewStringType()
	fields["tls_inspection_config_profile_binding_id"] = vapiBindings_.NewStringType()
	fieldNameMap["tier1_id"] = "Tier1Id"
	fieldNameMap["tls_inspection_config_profile_binding_id"] = "TlsInspectionConfigProfileBindingId"
	paramsTypeMap["tier1_id"] = vapiBindings_.NewStringType()
	paramsTypeMap["tls_inspection_config_profile_binding_id"] = vapiBindings_.NewStringType()
	paramsTypeMap["tier1Id"] = vapiBindings_.NewStringType()
	paramsTypeMap["tlsInspectionConfigProfileBindingId"] = vapiBindings_.NewStringType()
	pathParams["tier1_id"] = "tier1Id"
	pathParams["tls_inspection_config_profile_binding_id"] = "tlsInspectionConfigProfileBindingId"
	resultHeaders := map[string]string{}
	errorHeaders := map[string]map[string]string{}
	return vapiProtocol_.NewOperationRestMetadata(
		fields,
		fieldNameMap,
		paramsTypeMap,
		pathParams,
		queryParams,
		headerParams,
		dispatchHeaderParams,
		bodyFieldsMap,
		"",
		"",
		"DELETE",
		"/global-manager/api/v1/global-infra/tier-1s/{tier1Id}/tls-inspection-config-profile-bindings/{tlsInspectionConfigProfileBindingId}",
		"",
		resultHeaders,
		204,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func tlsInspectionConfigProfileBindingsGetInputType() vapiBindings_.StructType {
	fields := make(map[string]vapiBindings_.BindingType)
	fieldNameMap := make(map[string]string)
	fields["tier1_id"] = vapiBindings_.NewStringType()
	fields["tls_inspection_config_profile_binding_id"] = vapiBindings_.NewStringType()
	fieldNameMap["tier1_id"] = "Tier1Id"
	fieldNameMap["tls_inspection_config_profile_binding_id"] = "TlsInspectionConfigProfileBindingId"
	var validators = []vapiBindings_.Validator{}
	return vapiBindings_.NewStructType("operation-input", fields, reflect.TypeOf(vapiData_.StructValue{}), fieldNameMap, validators)
}

func TlsInspectionConfigProfileBindingsGetOutputType() vapiBindings_.BindingType {
	return vapiBindings_.NewReferenceType(nsx_global_policyModel.TlsConfigProfileBindingMapBindingType)
}

func tlsInspectionConfigProfileBindingsGetRestMetadata() vapiProtocol_.OperationRestMetadata {
	fields := map[string]vapiBindings_.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]vapiBindings_.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["tier1_id"] = vapiBindings_.NewStringType()
	fields["tls_inspection_config_profile_binding_id"] = vapiBindings_.NewStringType()
	fieldNameMap["tier1_id"] = "Tier1Id"
	fieldNameMap["tls_inspection_config_profile_binding_id"] = "TlsInspectionConfigProfileBindingId"
	paramsTypeMap["tier1_id"] = vapiBindings_.NewStringType()
	paramsTypeMap["tls_inspection_config_profile_binding_id"] = vapiBindings_.NewStringType()
	paramsTypeMap["tier1Id"] = vapiBindings_.NewStringType()
	paramsTypeMap["tlsInspectionConfigProfileBindingId"] = vapiBindings_.NewStringType()
	pathParams["tier1_id"] = "tier1Id"
	pathParams["tls_inspection_config_profile_binding_id"] = "tlsInspectionConfigProfileBindingId"
	resultHeaders := map[string]string{}
	errorHeaders := map[string]map[string]string{}
	return vapiProtocol_.NewOperationRestMetadata(
		fields,
		fieldNameMap,
		paramsTypeMap,
		pathParams,
		queryParams,
		headerParams,
		dispatchHeaderParams,
		bodyFieldsMap,
		"",
		"",
		"GET",
		"/global-manager/api/v1/global-infra/tier-1s/{tier1Id}/tls-inspection-config-profile-bindings/{tlsInspectionConfigProfileBindingId}",
		"",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func tlsInspectionConfigProfileBindingsPatchInputType() vapiBindings_.StructType {
	fields := make(map[string]vapiBindings_.BindingType)
	fieldNameMap := make(map[string]string)
	fields["tier1_id"] = vapiBindings_.NewStringType()
	fields["tls_inspection_config_profile_binding_id"] = vapiBindings_.NewStringType()
	fields["tls_config_profile_binding_map"] = vapiBindings_.NewReferenceType(nsx_global_policyModel.TlsConfigProfileBindingMapBindingType)
	fieldNameMap["tier1_id"] = "Tier1Id"
	fieldNameMap["tls_inspection_config_profile_binding_id"] = "TlsInspectionConfigProfileBindingId"
	fieldNameMap["tls_config_profile_binding_map"] = "TlsConfigProfileBindingMap"
	var validators = []vapiBindings_.Validator{}
	return vapiBindings_.NewStructType("operation-input", fields, reflect.TypeOf(vapiData_.StructValue{}), fieldNameMap, validators)
}

func TlsInspectionConfigProfileBindingsPatchOutputType() vapiBindings_.BindingType {
	return vapiBindings_.NewReferenceType(nsx_global_policyModel.TlsConfigProfileBindingMapBindingType)
}

func tlsInspectionConfigProfileBindingsPatchRestMetadata() vapiProtocol_.OperationRestMetadata {
	fields := map[string]vapiBindings_.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]vapiBindings_.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["tier1_id"] = vapiBindings_.NewStringType()
	fields["tls_inspection_config_profile_binding_id"] = vapiBindings_.NewStringType()
	fields["tls_config_profile_binding_map"] = vapiBindings_.NewReferenceType(nsx_global_policyModel.TlsConfigProfileBindingMapBindingType)
	fieldNameMap["tier1_id"] = "Tier1Id"
	fieldNameMap["tls_inspection_config_profile_binding_id"] = "TlsInspectionConfigProfileBindingId"
	fieldNameMap["tls_config_profile_binding_map"] = "TlsConfigProfileBindingMap"
	paramsTypeMap["tier1_id"] = vapiBindings_.NewStringType()
	paramsTypeMap["tls_config_profile_binding_map"] = vapiBindings_.NewReferenceType(nsx_global_policyModel.TlsConfigProfileBindingMapBindingType)
	paramsTypeMap["tls_inspection_config_profile_binding_id"] = vapiBindings_.NewStringType()
	paramsTypeMap["tier1Id"] = vapiBindings_.NewStringType()
	paramsTypeMap["tlsInspectionConfigProfileBindingId"] = vapiBindings_.NewStringType()
	pathParams["tier1_id"] = "tier1Id"
	pathParams["tls_inspection_config_profile_binding_id"] = "tlsInspectionConfigProfileBindingId"
	resultHeaders := map[string]string{}
	errorHeaders := map[string]map[string]string{}
	return vapiProtocol_.NewOperationRestMetadata(
		fields,
		fieldNameMap,
		paramsTypeMap,
		pathParams,
		queryParams,
		headerParams,
		dispatchHeaderParams,
		bodyFieldsMap,
		"",
		"tls_config_profile_binding_map",
		"PATCH",
		"/global-manager/api/v1/global-infra/tier-1s/{tier1Id}/tls-inspection-config-profile-bindings/{tlsInspectionConfigProfileBindingId}",
		"",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func tlsInspectionConfigProfileBindingsUpdateInputType() vapiBindings_.StructType {
	fields := make(map[string]vapiBindings_.BindingType)
	fieldNameMap := make(map[string]string)
	fields["tier1_id"] = vapiBindings_.NewStringType()
	fields["tls_inspection_config_profile_binding_id"] = vapiBindings_.NewStringType()
	fields["tls_config_profile_binding_map"] = vapiBindings_.NewReferenceType(nsx_global_policyModel.TlsConfigProfileBindingMapBindingType)
	fieldNameMap["tier1_id"] = "Tier1Id"
	fieldNameMap["tls_inspection_config_profile_binding_id"] = "TlsInspectionConfigProfileBindingId"
	fieldNameMap["tls_config_profile_binding_map"] = "TlsConfigProfileBindingMap"
	var validators = []vapiBindings_.Validator{}
	return vapiBindings_.NewStructType("operation-input", fields, reflect.TypeOf(vapiData_.StructValue{}), fieldNameMap, validators)
}

func TlsInspectionConfigProfileBindingsUpdateOutputType() vapiBindings_.BindingType {
	return vapiBindings_.NewReferenceType(nsx_global_policyModel.TlsConfigProfileBindingMapBindingType)
}

func tlsInspectionConfigProfileBindingsUpdateRestMetadata() vapiProtocol_.OperationRestMetadata {
	fields := map[string]vapiBindings_.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]vapiBindings_.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["tier1_id"] = vapiBindings_.NewStringType()
	fields["tls_inspection_config_profile_binding_id"] = vapiBindings_.NewStringType()
	fields["tls_config_profile_binding_map"] = vapiBindings_.NewReferenceType(nsx_global_policyModel.TlsConfigProfileBindingMapBindingType)
	fieldNameMap["tier1_id"] = "Tier1Id"
	fieldNameMap["tls_inspection_config_profile_binding_id"] = "TlsInspectionConfigProfileBindingId"
	fieldNameMap["tls_config_profile_binding_map"] = "TlsConfigProfileBindingMap"
	paramsTypeMap["tier1_id"] = vapiBindings_.NewStringType()
	paramsTypeMap["tls_config_profile_binding_map"] = vapiBindings_.NewReferenceType(nsx_global_policyModel.TlsConfigProfileBindingMapBindingType)
	paramsTypeMap["tls_inspection_config_profile_binding_id"] = vapiBindings_.NewStringType()
	paramsTypeMap["tier1Id"] = vapiBindings_.NewStringType()
	paramsTypeMap["tlsInspectionConfigProfileBindingId"] = vapiBindings_.NewStringType()
	pathParams["tier1_id"] = "tier1Id"
	pathParams["tls_inspection_config_profile_binding_id"] = "tlsInspectionConfigProfileBindingId"
	resultHeaders := map[string]string{}
	errorHeaders := map[string]map[string]string{}
	return vapiProtocol_.NewOperationRestMetadata(
		fields,
		fieldNameMap,
		paramsTypeMap,
		pathParams,
		queryParams,
		headerParams,
		dispatchHeaderParams,
		bodyFieldsMap,
		"",
		"tls_config_profile_binding_map",
		"PUT",
		"/global-manager/api/v1/global-infra/tier-1s/{tier1Id}/tls-inspection-config-profile-bindings/{tlsInspectionConfigProfileBindingId}",
		"",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}
