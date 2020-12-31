package backup

// Copyright (c) Microsoft and contributors.  All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//
// See the License for the specific language governing permissions and
// limitations under the License.
//
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

// AzureFileShareType enumerates the values for azure file share type.
type AzureFileShareType string

const (
	// Invalid ...
	Invalid AzureFileShareType = "Invalid"
	// XSMB ...
	XSMB AzureFileShareType = "XSMB"
	// XSync ...
	XSync AzureFileShareType = "XSync"
)

// PossibleAzureFileShareTypeValues returns an array of possible values for the AzureFileShareType const type.
func PossibleAzureFileShareTypeValues() []AzureFileShareType {
	return []AzureFileShareType{Invalid, XSMB, XSync}
}

// ContainerType enumerates the values for container type.
type ContainerType string

const (
	// ContainerTypeAzureBackupServerContainer ...
	ContainerTypeAzureBackupServerContainer ContainerType = "AzureBackupServerContainer"
	// ContainerTypeAzureSQLContainer ...
	ContainerTypeAzureSQLContainer ContainerType = "AzureSqlContainer"
	// ContainerTypeCluster ...
	ContainerTypeCluster ContainerType = "Cluster"
	// ContainerTypeDPMContainer ...
	ContainerTypeDPMContainer ContainerType = "DPMContainer"
	// ContainerTypeGenericContainer ...
	ContainerTypeGenericContainer ContainerType = "GenericContainer"
	// ContainerTypeIaasVMContainer ...
	ContainerTypeIaasVMContainer ContainerType = "IaasVMContainer"
	// ContainerTypeIaasVMServiceContainer ...
	ContainerTypeIaasVMServiceContainer ContainerType = "IaasVMServiceContainer"
	// ContainerTypeInvalid ...
	ContainerTypeInvalid ContainerType = "Invalid"
	// ContainerTypeMABContainer ...
	ContainerTypeMABContainer ContainerType = "MABContainer"
	// ContainerTypeSQLAGWorkLoadContainer ...
	ContainerTypeSQLAGWorkLoadContainer ContainerType = "SQLAGWorkLoadContainer"
	// ContainerTypeStorageContainer ...
	ContainerTypeStorageContainer ContainerType = "StorageContainer"
	// ContainerTypeUnknown ...
	ContainerTypeUnknown ContainerType = "Unknown"
	// ContainerTypeVCenter ...
	ContainerTypeVCenter ContainerType = "VCenter"
	// ContainerTypeVMAppContainer ...
	ContainerTypeVMAppContainer ContainerType = "VMAppContainer"
	// ContainerTypeWindows ...
	ContainerTypeWindows ContainerType = "Windows"
)

// PossibleContainerTypeValues returns an array of possible values for the ContainerType const type.
func PossibleContainerTypeValues() []ContainerType {
	return []ContainerType{ContainerTypeAzureBackupServerContainer, ContainerTypeAzureSQLContainer, ContainerTypeCluster, ContainerTypeDPMContainer, ContainerTypeGenericContainer, ContainerTypeIaasVMContainer, ContainerTypeIaasVMServiceContainer, ContainerTypeInvalid, ContainerTypeMABContainer, ContainerTypeSQLAGWorkLoadContainer, ContainerTypeStorageContainer, ContainerTypeUnknown, ContainerTypeVCenter, ContainerTypeVMAppContainer, ContainerTypeWindows}
}

// ContainerTypeBasicProtectionContainer enumerates the values for container type basic protection container.
type ContainerTypeBasicProtectionContainer string

const (
	// ContainerTypeAzureBackupServerContainer1 ...
	ContainerTypeAzureBackupServerContainer1 ContainerTypeBasicProtectionContainer = "AzureBackupServerContainer"
	// ContainerTypeAzureSQLContainer1 ...
	ContainerTypeAzureSQLContainer1 ContainerTypeBasicProtectionContainer = "AzureSqlContainer"
	// ContainerTypeAzureWorkloadContainer ...
	ContainerTypeAzureWorkloadContainer ContainerTypeBasicProtectionContainer = "AzureWorkloadContainer"
	// ContainerTypeDPMContainer1 ...
	ContainerTypeDPMContainer1 ContainerTypeBasicProtectionContainer = "DPMContainer"
	// ContainerTypeGenericContainer1 ...
	ContainerTypeGenericContainer1 ContainerTypeBasicProtectionContainer = "GenericContainer"
	// ContainerTypeIaaSVMContainer ...
	ContainerTypeIaaSVMContainer ContainerTypeBasicProtectionContainer = "IaaSVMContainer"
	// ContainerTypeMicrosoftClassicComputevirtualMachines ...
	ContainerTypeMicrosoftClassicComputevirtualMachines ContainerTypeBasicProtectionContainer = "Microsoft.ClassicCompute/virtualMachines"
	// ContainerTypeMicrosoftComputevirtualMachines ...
	ContainerTypeMicrosoftComputevirtualMachines ContainerTypeBasicProtectionContainer = "Microsoft.Compute/virtualMachines"
	// ContainerTypeProtectionContainer ...
	ContainerTypeProtectionContainer ContainerTypeBasicProtectionContainer = "ProtectionContainer"
	// ContainerTypeSQLAGWorkLoadContainer1 ...
	ContainerTypeSQLAGWorkLoadContainer1 ContainerTypeBasicProtectionContainer = "SQLAGWorkLoadContainer"
	// ContainerTypeStorageContainer1 ...
	ContainerTypeStorageContainer1 ContainerTypeBasicProtectionContainer = "StorageContainer"
	// ContainerTypeVMAppContainer1 ...
	ContainerTypeVMAppContainer1 ContainerTypeBasicProtectionContainer = "VMAppContainer"
	// ContainerTypeWindows1 ...
	ContainerTypeWindows1 ContainerTypeBasicProtectionContainer = "Windows"
)

// PossibleContainerTypeBasicProtectionContainerValues returns an array of possible values for the ContainerTypeBasicProtectionContainer const type.
func PossibleContainerTypeBasicProtectionContainerValues() []ContainerTypeBasicProtectionContainer {
	return []ContainerTypeBasicProtectionContainer{ContainerTypeAzureBackupServerContainer1, ContainerTypeAzureSQLContainer1, ContainerTypeAzureWorkloadContainer, ContainerTypeDPMContainer1, ContainerTypeGenericContainer1, ContainerTypeIaaSVMContainer, ContainerTypeMicrosoftClassicComputevirtualMachines, ContainerTypeMicrosoftComputevirtualMachines, ContainerTypeProtectionContainer, ContainerTypeSQLAGWorkLoadContainer1, ContainerTypeStorageContainer1, ContainerTypeVMAppContainer1, ContainerTypeWindows1}
}

// DataSourceType enumerates the values for data source type.
type DataSourceType string

const (
	// DataSourceTypeAzureFileShare ...
	DataSourceTypeAzureFileShare DataSourceType = "AzureFileShare"
	// DataSourceTypeAzureSQLDb ...
	DataSourceTypeAzureSQLDb DataSourceType = "AzureSqlDb"
	// DataSourceTypeClient ...
	DataSourceTypeClient DataSourceType = "Client"
	// DataSourceTypeExchange ...
	DataSourceTypeExchange DataSourceType = "Exchange"
	// DataSourceTypeFileFolder ...
	DataSourceTypeFileFolder DataSourceType = "FileFolder"
	// DataSourceTypeGenericDataSource ...
	DataSourceTypeGenericDataSource DataSourceType = "GenericDataSource"
	// DataSourceTypeInvalid ...
	DataSourceTypeInvalid DataSourceType = "Invalid"
	// DataSourceTypeSAPAseDatabase ...
	DataSourceTypeSAPAseDatabase DataSourceType = "SAPAseDatabase"
	// DataSourceTypeSAPHanaDatabase ...
	DataSourceTypeSAPHanaDatabase DataSourceType = "SAPHanaDatabase"
	// DataSourceTypeSharepoint ...
	DataSourceTypeSharepoint DataSourceType = "Sharepoint"
	// DataSourceTypeSQLDataBase ...
	DataSourceTypeSQLDataBase DataSourceType = "SQLDataBase"
	// DataSourceTypeSQLDB ...
	DataSourceTypeSQLDB DataSourceType = "SQLDB"
	// DataSourceTypeSystemState ...
	DataSourceTypeSystemState DataSourceType = "SystemState"
	// DataSourceTypeVM ...
	DataSourceTypeVM DataSourceType = "VM"
	// DataSourceTypeVMwareVM ...
	DataSourceTypeVMwareVM DataSourceType = "VMwareVM"
)

// PossibleDataSourceTypeValues returns an array of possible values for the DataSourceType const type.
func PossibleDataSourceTypeValues() []DataSourceType {
	return []DataSourceType{DataSourceTypeAzureFileShare, DataSourceTypeAzureSQLDb, DataSourceTypeClient, DataSourceTypeExchange, DataSourceTypeFileFolder, DataSourceTypeGenericDataSource, DataSourceTypeInvalid, DataSourceTypeSAPAseDatabase, DataSourceTypeSAPHanaDatabase, DataSourceTypeSharepoint, DataSourceTypeSQLDataBase, DataSourceTypeSQLDB, DataSourceTypeSystemState, DataSourceTypeVM, DataSourceTypeVMwareVM}
}

// EngineType enumerates the values for engine type.
type EngineType string

const (
	// BackupEngineTypeAzureBackupServerEngine ...
	BackupEngineTypeAzureBackupServerEngine EngineType = "AzureBackupServerEngine"
	// BackupEngineTypeBackupEngineBase ...
	BackupEngineTypeBackupEngineBase EngineType = "BackupEngineBase"
	// BackupEngineTypeDpmBackupEngine ...
	BackupEngineTypeDpmBackupEngine EngineType = "DpmBackupEngine"
)

// PossibleEngineTypeValues returns an array of possible values for the EngineType const type.
func PossibleEngineTypeValues() []EngineType {
	return []EngineType{BackupEngineTypeAzureBackupServerEngine, BackupEngineTypeBackupEngineBase, BackupEngineTypeDpmBackupEngine}
}

// FabricName enumerates the values for fabric name.
type FabricName string

const (
	// FabricNameAzure ...
	FabricNameAzure FabricName = "Azure"
	// FabricNameInvalid ...
	FabricNameInvalid FabricName = "Invalid"
)

// PossibleFabricNameValues returns an array of possible values for the FabricName const type.
func PossibleFabricNameValues() []FabricName {
	return []FabricName{FabricNameAzure, FabricNameInvalid}
}

// FeatureType enumerates the values for feature type.
type FeatureType string

const (
	// FeatureTypeAzureBackupGoals ...
	FeatureTypeAzureBackupGoals FeatureType = "AzureBackupGoals"
	// FeatureTypeAzureVMResourceBackup ...
	FeatureTypeAzureVMResourceBackup FeatureType = "AzureVMResourceBackup"
	// FeatureTypeFeatureSupportRequest ...
	FeatureTypeFeatureSupportRequest FeatureType = "FeatureSupportRequest"
)

// PossibleFeatureTypeValues returns an array of possible values for the FeatureType const type.
func PossibleFeatureTypeValues() []FeatureType {
	return []FeatureType{FeatureTypeAzureBackupGoals, FeatureTypeAzureVMResourceBackup, FeatureTypeFeatureSupportRequest}
}

// InquiryStatus enumerates the values for inquiry status.
type InquiryStatus string

const (
	// InquiryStatusFailed ...
	InquiryStatusFailed InquiryStatus = "Failed"
	// InquiryStatusInvalid ...
	InquiryStatusInvalid InquiryStatus = "Invalid"
	// InquiryStatusSuccess ...
	InquiryStatusSuccess InquiryStatus = "Success"
)

// PossibleInquiryStatusValues returns an array of possible values for the InquiryStatus const type.
func PossibleInquiryStatusValues() []InquiryStatus {
	return []InquiryStatus{InquiryStatusFailed, InquiryStatusInvalid, InquiryStatusSuccess}
}

// IntentItemType enumerates the values for intent item type.
type IntentItemType string

const (
	// IntentItemTypeInvalid ...
	IntentItemTypeInvalid IntentItemType = "Invalid"
	// IntentItemTypeSQLAvailabilityGroupContainer ...
	IntentItemTypeSQLAvailabilityGroupContainer IntentItemType = "SQLAvailabilityGroupContainer"
	// IntentItemTypeSQLInstance ...
	IntentItemTypeSQLInstance IntentItemType = "SQLInstance"
)

// PossibleIntentItemTypeValues returns an array of possible values for the IntentItemType const type.
func PossibleIntentItemTypeValues() []IntentItemType {
	return []IntentItemType{IntentItemTypeInvalid, IntentItemTypeSQLAvailabilityGroupContainer, IntentItemTypeSQLInstance}
}

// ItemType enumerates the values for item type.
type ItemType string

const (
	// ItemTypeAzureFileShare ...
	ItemTypeAzureFileShare ItemType = "AzureFileShare"
	// ItemTypeAzureSQLDb ...
	ItemTypeAzureSQLDb ItemType = "AzureSqlDb"
	// ItemTypeClient ...
	ItemTypeClient ItemType = "Client"
	// ItemTypeExchange ...
	ItemTypeExchange ItemType = "Exchange"
	// ItemTypeFileFolder ...
	ItemTypeFileFolder ItemType = "FileFolder"
	// ItemTypeGenericDataSource ...
	ItemTypeGenericDataSource ItemType = "GenericDataSource"
	// ItemTypeInvalid ...
	ItemTypeInvalid ItemType = "Invalid"
	// ItemTypeSAPAseDatabase ...
	ItemTypeSAPAseDatabase ItemType = "SAPAseDatabase"
	// ItemTypeSAPHanaDatabase ...
	ItemTypeSAPHanaDatabase ItemType = "SAPHanaDatabase"
	// ItemTypeSharepoint ...
	ItemTypeSharepoint ItemType = "Sharepoint"
	// ItemTypeSQLDataBase ...
	ItemTypeSQLDataBase ItemType = "SQLDataBase"
	// ItemTypeSQLDB ...
	ItemTypeSQLDB ItemType = "SQLDB"
	// ItemTypeSystemState ...
	ItemTypeSystemState ItemType = "SystemState"
	// ItemTypeVM ...
	ItemTypeVM ItemType = "VM"
	// ItemTypeVMwareVM ...
	ItemTypeVMwareVM ItemType = "VMwareVM"
)

// PossibleItemTypeValues returns an array of possible values for the ItemType const type.
func PossibleItemTypeValues() []ItemType {
	return []ItemType{ItemTypeAzureFileShare, ItemTypeAzureSQLDb, ItemTypeClient, ItemTypeExchange, ItemTypeFileFolder, ItemTypeGenericDataSource, ItemTypeInvalid, ItemTypeSAPAseDatabase, ItemTypeSAPHanaDatabase, ItemTypeSharepoint, ItemTypeSQLDataBase, ItemTypeSQLDB, ItemTypeSystemState, ItemTypeVM, ItemTypeVMwareVM}
}

// JobSupportedAction enumerates the values for job supported action.
type JobSupportedAction string

const (
	// JobSupportedActionCancellable ...
	JobSupportedActionCancellable JobSupportedAction = "Cancellable"
	// JobSupportedActionInvalid ...
	JobSupportedActionInvalid JobSupportedAction = "Invalid"
	// JobSupportedActionRetriable ...
	JobSupportedActionRetriable JobSupportedAction = "Retriable"
)

// PossibleJobSupportedActionValues returns an array of possible values for the JobSupportedAction const type.
func PossibleJobSupportedActionValues() []JobSupportedAction {
	return []JobSupportedAction{JobSupportedActionCancellable, JobSupportedActionInvalid, JobSupportedActionRetriable}
}

// JobType enumerates the values for job type.
type JobType string

const (
	// JobTypeAzureIaaSVMJob ...
	JobTypeAzureIaaSVMJob JobType = "AzureIaaSVMJob"
	// JobTypeAzureStorageJob ...
	JobTypeAzureStorageJob JobType = "AzureStorageJob"
	// JobTypeAzureWorkloadJob ...
	JobTypeAzureWorkloadJob JobType = "AzureWorkloadJob"
	// JobTypeDpmJob ...
	JobTypeDpmJob JobType = "DpmJob"
	// JobTypeJob ...
	JobTypeJob JobType = "Job"
	// JobTypeMabJob ...
	JobTypeMabJob JobType = "MabJob"
)

// PossibleJobTypeValues returns an array of possible values for the JobType const type.
func PossibleJobTypeValues() []JobType {
	return []JobType{JobTypeAzureIaaSVMJob, JobTypeAzureStorageJob, JobTypeAzureWorkloadJob, JobTypeDpmJob, JobTypeJob, JobTypeMabJob}
}

// MabServerType enumerates the values for mab server type.
type MabServerType string

const (
	// MabServerTypeAzureBackupServerContainer ...
	MabServerTypeAzureBackupServerContainer MabServerType = "AzureBackupServerContainer"
	// MabServerTypeAzureSQLContainer ...
	MabServerTypeAzureSQLContainer MabServerType = "AzureSqlContainer"
	// MabServerTypeCluster ...
	MabServerTypeCluster MabServerType = "Cluster"
	// MabServerTypeDPMContainer ...
	MabServerTypeDPMContainer MabServerType = "DPMContainer"
	// MabServerTypeGenericContainer ...
	MabServerTypeGenericContainer MabServerType = "GenericContainer"
	// MabServerTypeIaasVMContainer ...
	MabServerTypeIaasVMContainer MabServerType = "IaasVMContainer"
	// MabServerTypeIaasVMServiceContainer ...
	MabServerTypeIaasVMServiceContainer MabServerType = "IaasVMServiceContainer"
	// MabServerTypeInvalid ...
	MabServerTypeInvalid MabServerType = "Invalid"
	// MabServerTypeMABContainer ...
	MabServerTypeMABContainer MabServerType = "MABContainer"
	// MabServerTypeSQLAGWorkLoadContainer ...
	MabServerTypeSQLAGWorkLoadContainer MabServerType = "SQLAGWorkLoadContainer"
	// MabServerTypeStorageContainer ...
	MabServerTypeStorageContainer MabServerType = "StorageContainer"
	// MabServerTypeUnknown ...
	MabServerTypeUnknown MabServerType = "Unknown"
	// MabServerTypeVCenter ...
	MabServerTypeVCenter MabServerType = "VCenter"
	// MabServerTypeVMAppContainer ...
	MabServerTypeVMAppContainer MabServerType = "VMAppContainer"
	// MabServerTypeWindows ...
	MabServerTypeWindows MabServerType = "Windows"
)

// PossibleMabServerTypeValues returns an array of possible values for the MabServerType const type.
func PossibleMabServerTypeValues() []MabServerType {
	return []MabServerType{MabServerTypeAzureBackupServerContainer, MabServerTypeAzureSQLContainer, MabServerTypeCluster, MabServerTypeDPMContainer, MabServerTypeGenericContainer, MabServerTypeIaasVMContainer, MabServerTypeIaasVMServiceContainer, MabServerTypeInvalid, MabServerTypeMABContainer, MabServerTypeSQLAGWorkLoadContainer, MabServerTypeStorageContainer, MabServerTypeUnknown, MabServerTypeVCenter, MabServerTypeVMAppContainer, MabServerTypeWindows}
}

// ManagementType enumerates the values for management type.
type ManagementType string

const (
	// ManagementTypeAzureBackupServer ...
	ManagementTypeAzureBackupServer ManagementType = "AzureBackupServer"
	// ManagementTypeAzureIaasVM ...
	ManagementTypeAzureIaasVM ManagementType = "AzureIaasVM"
	// ManagementTypeAzureSQL ...
	ManagementTypeAzureSQL ManagementType = "AzureSql"
	// ManagementTypeAzureStorage ...
	ManagementTypeAzureStorage ManagementType = "AzureStorage"
	// ManagementTypeAzureWorkload ...
	ManagementTypeAzureWorkload ManagementType = "AzureWorkload"
	// ManagementTypeDefaultBackup ...
	ManagementTypeDefaultBackup ManagementType = "DefaultBackup"
	// ManagementTypeDPM ...
	ManagementTypeDPM ManagementType = "DPM"
	// ManagementTypeInvalid ...
	ManagementTypeInvalid ManagementType = "Invalid"
	// ManagementTypeMAB ...
	ManagementTypeMAB ManagementType = "MAB"
)

// PossibleManagementTypeValues returns an array of possible values for the ManagementType const type.
func PossibleManagementTypeValues() []ManagementType {
	return []ManagementType{ManagementTypeAzureBackupServer, ManagementTypeAzureIaasVM, ManagementTypeAzureSQL, ManagementTypeAzureStorage, ManagementTypeAzureWorkload, ManagementTypeDefaultBackup, ManagementTypeDPM, ManagementTypeInvalid, ManagementTypeMAB}
}

// ObjectType enumerates the values for object type.
type ObjectType string

const (
	// ObjectTypeAzureFileShareBackupRequest ...
	ObjectTypeAzureFileShareBackupRequest ObjectType = "AzureFileShareBackupRequest"
	// ObjectTypeAzureWorkloadBackupRequest ...
	ObjectTypeAzureWorkloadBackupRequest ObjectType = "AzureWorkloadBackupRequest"
	// ObjectTypeBackupRequest ...
	ObjectTypeBackupRequest ObjectType = "BackupRequest"
	// ObjectTypeIaasVMBackupRequest ...
	ObjectTypeIaasVMBackupRequest ObjectType = "IaasVMBackupRequest"
)

// PossibleObjectTypeValues returns an array of possible values for the ObjectType const type.
func PossibleObjectTypeValues() []ObjectType {
	return []ObjectType{ObjectTypeAzureFileShareBackupRequest, ObjectTypeAzureWorkloadBackupRequest, ObjectTypeBackupRequest, ObjectTypeIaasVMBackupRequest}
}

// ObjectTypeBasicILRRequest enumerates the values for object type basic ilr request.
type ObjectTypeBasicILRRequest string

const (
	// ObjectTypeAzureFileShareProvisionILRRequest ...
	ObjectTypeAzureFileShareProvisionILRRequest ObjectTypeBasicILRRequest = "AzureFileShareProvisionILRRequest"
	// ObjectTypeIaasVMILRRegistrationRequest ...
	ObjectTypeIaasVMILRRegistrationRequest ObjectTypeBasicILRRequest = "IaasVMILRRegistrationRequest"
	// ObjectTypeILRRequest ...
	ObjectTypeILRRequest ObjectTypeBasicILRRequest = "ILRRequest"
)

// PossibleObjectTypeBasicILRRequestValues returns an array of possible values for the ObjectTypeBasicILRRequest const type.
func PossibleObjectTypeBasicILRRequestValues() []ObjectTypeBasicILRRequest {
	return []ObjectTypeBasicILRRequest{ObjectTypeAzureFileShareProvisionILRRequest, ObjectTypeIaasVMILRRegistrationRequest, ObjectTypeILRRequest}
}

// ObjectTypeBasicOperationStatusExtendedInfo enumerates the values for object type basic operation status
// extended info.
type ObjectTypeBasicOperationStatusExtendedInfo string

const (
	// ObjectTypeOperationStatusExtendedInfo ...
	ObjectTypeOperationStatusExtendedInfo ObjectTypeBasicOperationStatusExtendedInfo = "OperationStatusExtendedInfo"
	// ObjectTypeOperationStatusJobExtendedInfo ...
	ObjectTypeOperationStatusJobExtendedInfo ObjectTypeBasicOperationStatusExtendedInfo = "OperationStatusJobExtendedInfo"
	// ObjectTypeOperationStatusJobsExtendedInfo ...
	ObjectTypeOperationStatusJobsExtendedInfo ObjectTypeBasicOperationStatusExtendedInfo = "OperationStatusJobsExtendedInfo"
	// ObjectTypeOperationStatusProvisionILRExtendedInfo ...
	ObjectTypeOperationStatusProvisionILRExtendedInfo ObjectTypeBasicOperationStatusExtendedInfo = "OperationStatusProvisionILRExtendedInfo"
)

// PossibleObjectTypeBasicOperationStatusExtendedInfoValues returns an array of possible values for the ObjectTypeBasicOperationStatusExtendedInfo const type.
func PossibleObjectTypeBasicOperationStatusExtendedInfoValues() []ObjectTypeBasicOperationStatusExtendedInfo {
	return []ObjectTypeBasicOperationStatusExtendedInfo{ObjectTypeOperationStatusExtendedInfo, ObjectTypeOperationStatusJobExtendedInfo, ObjectTypeOperationStatusJobsExtendedInfo, ObjectTypeOperationStatusProvisionILRExtendedInfo}
}

// ObjectTypeBasicRecoveryPoint enumerates the values for object type basic recovery point.
type ObjectTypeBasicRecoveryPoint string

const (
	// ObjectTypeRecoveryPoint ...
	ObjectTypeRecoveryPoint ObjectTypeBasicRecoveryPoint = "RecoveryPoint"
)

// PossibleObjectTypeBasicRecoveryPointValues returns an array of possible values for the ObjectTypeBasicRecoveryPoint const type.
func PossibleObjectTypeBasicRecoveryPointValues() []ObjectTypeBasicRecoveryPoint {
	return []ObjectTypeBasicRecoveryPoint{ObjectTypeRecoveryPoint}
}

// OperationStatusValues enumerates the values for operation status values.
type OperationStatusValues string

const (
	// OperationStatusValuesCanceled ...
	OperationStatusValuesCanceled OperationStatusValues = "Canceled"
	// OperationStatusValuesFailed ...
	OperationStatusValuesFailed OperationStatusValues = "Failed"
	// OperationStatusValuesInProgress ...
	OperationStatusValuesInProgress OperationStatusValues = "InProgress"
	// OperationStatusValuesInvalid ...
	OperationStatusValuesInvalid OperationStatusValues = "Invalid"
	// OperationStatusValuesSucceeded ...
	OperationStatusValuesSucceeded OperationStatusValues = "Succeeded"
)

// PossibleOperationStatusValuesValues returns an array of possible values for the OperationStatusValues const type.
func PossibleOperationStatusValuesValues() []OperationStatusValues {
	return []OperationStatusValues{OperationStatusValuesCanceled, OperationStatusValuesFailed, OperationStatusValuesInProgress, OperationStatusValuesInvalid, OperationStatusValuesSucceeded}
}

// OperationType enumerates the values for operation type.
type OperationType string

const (
	// OperationTypeInvalid ...
	OperationTypeInvalid OperationType = "Invalid"
	// OperationTypeRegister ...
	OperationTypeRegister OperationType = "Register"
	// OperationTypeReregister ...
	OperationTypeReregister OperationType = "Reregister"
)

// PossibleOperationTypeValues returns an array of possible values for the OperationType const type.
func PossibleOperationTypeValues() []OperationType {
	return []OperationType{OperationTypeInvalid, OperationTypeRegister, OperationTypeReregister}
}

// ProtectableContainerType enumerates the values for protectable container type.
type ProtectableContainerType string

const (
	// ProtectableContainerTypeProtectableContainer ...
	ProtectableContainerTypeProtectableContainer ProtectableContainerType = "ProtectableContainer"
	// ProtectableContainerTypeStorageContainer ...
	ProtectableContainerTypeStorageContainer ProtectableContainerType = "StorageContainer"
	// ProtectableContainerTypeVMAppContainer ...
	ProtectableContainerTypeVMAppContainer ProtectableContainerType = "VMAppContainer"
)

// PossibleProtectableContainerTypeValues returns an array of possible values for the ProtectableContainerType const type.
func PossibleProtectableContainerTypeValues() []ProtectableContainerType {
	return []ProtectableContainerType{ProtectableContainerTypeProtectableContainer, ProtectableContainerTypeStorageContainer, ProtectableContainerTypeVMAppContainer}
}

// ProtectableItemType enumerates the values for protectable item type.
type ProtectableItemType string

const (
	// ProtectableItemTypeAzureFileShare ...
	ProtectableItemTypeAzureFileShare ProtectableItemType = "AzureFileShare"
	// ProtectableItemTypeAzureVMWorkloadProtectableItem ...
	ProtectableItemTypeAzureVMWorkloadProtectableItem ProtectableItemType = "AzureVmWorkloadProtectableItem"
	// ProtectableItemTypeIaaSVMProtectableItem ...
	ProtectableItemTypeIaaSVMProtectableItem ProtectableItemType = "IaaSVMProtectableItem"
	// ProtectableItemTypeMicrosoftClassicComputevirtualMachines ...
	ProtectableItemTypeMicrosoftClassicComputevirtualMachines ProtectableItemType = "Microsoft.ClassicCompute/virtualMachines"
	// ProtectableItemTypeMicrosoftComputevirtualMachines ...
	ProtectableItemTypeMicrosoftComputevirtualMachines ProtectableItemType = "Microsoft.Compute/virtualMachines"
	// ProtectableItemTypeSAPAseSystem ...
	ProtectableItemTypeSAPAseSystem ProtectableItemType = "SAPAseSystem"
	// ProtectableItemTypeSAPHanaDatabase ...
	ProtectableItemTypeSAPHanaDatabase ProtectableItemType = "SAPHanaDatabase"
	// ProtectableItemTypeSAPHanaSystem ...
	ProtectableItemTypeSAPHanaSystem ProtectableItemType = "SAPHanaSystem"
	// ProtectableItemTypeSQLAvailabilityGroupContainer ...
	ProtectableItemTypeSQLAvailabilityGroupContainer ProtectableItemType = "SQLAvailabilityGroupContainer"
	// ProtectableItemTypeSQLDataBase ...
	ProtectableItemTypeSQLDataBase ProtectableItemType = "SQLDataBase"
	// ProtectableItemTypeSQLInstance ...
	ProtectableItemTypeSQLInstance ProtectableItemType = "SQLInstance"
	// ProtectableItemTypeWorkloadProtectableItem ...
	ProtectableItemTypeWorkloadProtectableItem ProtectableItemType = "WorkloadProtectableItem"
)

// PossibleProtectableItemTypeValues returns an array of possible values for the ProtectableItemType const type.
func PossibleProtectableItemTypeValues() []ProtectableItemType {
	return []ProtectableItemType{ProtectableItemTypeAzureFileShare, ProtectableItemTypeAzureVMWorkloadProtectableItem, ProtectableItemTypeIaaSVMProtectableItem, ProtectableItemTypeMicrosoftClassicComputevirtualMachines, ProtectableItemTypeMicrosoftComputevirtualMachines, ProtectableItemTypeSAPAseSystem, ProtectableItemTypeSAPHanaDatabase, ProtectableItemTypeSAPHanaSystem, ProtectableItemTypeSQLAvailabilityGroupContainer, ProtectableItemTypeSQLDataBase, ProtectableItemTypeSQLInstance, ProtectableItemTypeWorkloadProtectableItem}
}

// ProtectionIntentItemType enumerates the values for protection intent item type.
type ProtectionIntentItemType string

const (
	// ProtectionIntentItemTypeAzureResourceItem ...
	ProtectionIntentItemTypeAzureResourceItem ProtectionIntentItemType = "AzureResourceItem"
	// ProtectionIntentItemTypeAzureWorkloadAutoProtectionIntent ...
	ProtectionIntentItemTypeAzureWorkloadAutoProtectionIntent ProtectionIntentItemType = "AzureWorkloadAutoProtectionIntent"
	// ProtectionIntentItemTypeAzureWorkloadSQLAutoProtectionIntent ...
	ProtectionIntentItemTypeAzureWorkloadSQLAutoProtectionIntent ProtectionIntentItemType = "AzureWorkloadSQLAutoProtectionIntent"
	// ProtectionIntentItemTypeProtectionIntent ...
	ProtectionIntentItemTypeProtectionIntent ProtectionIntentItemType = "ProtectionIntent"
	// ProtectionIntentItemTypeRecoveryServiceVaultItem ...
	ProtectionIntentItemTypeRecoveryServiceVaultItem ProtectionIntentItemType = "RecoveryServiceVaultItem"
)

// PossibleProtectionIntentItemTypeValues returns an array of possible values for the ProtectionIntentItemType const type.
func PossibleProtectionIntentItemTypeValues() []ProtectionIntentItemType {
	return []ProtectionIntentItemType{ProtectionIntentItemTypeAzureResourceItem, ProtectionIntentItemTypeAzureWorkloadAutoProtectionIntent, ProtectionIntentItemTypeAzureWorkloadSQLAutoProtectionIntent, ProtectionIntentItemTypeProtectionIntent, ProtectionIntentItemTypeRecoveryServiceVaultItem}
}

// ProtectionStatus enumerates the values for protection status.
type ProtectionStatus string

const (
	// ProtectionStatusInvalid ...
	ProtectionStatusInvalid ProtectionStatus = "Invalid"
	// ProtectionStatusNotProtected ...
	ProtectionStatusNotProtected ProtectionStatus = "NotProtected"
	// ProtectionStatusProtected ...
	ProtectionStatusProtected ProtectionStatus = "Protected"
	// ProtectionStatusProtecting ...
	ProtectionStatusProtecting ProtectionStatus = "Protecting"
	// ProtectionStatusProtectionFailed ...
	ProtectionStatusProtectionFailed ProtectionStatus = "ProtectionFailed"
)

// PossibleProtectionStatusValues returns an array of possible values for the ProtectionStatus const type.
func PossibleProtectionStatusValues() []ProtectionStatus {
	return []ProtectionStatus{ProtectionStatusInvalid, ProtectionStatusNotProtected, ProtectionStatusProtected, ProtectionStatusProtecting, ProtectionStatusProtectionFailed}
}

// SQLDataDirectoryType enumerates the values for sql data directory type.
type SQLDataDirectoryType string

const (
	// SQLDataDirectoryTypeData ...
	SQLDataDirectoryTypeData SQLDataDirectoryType = "Data"
	// SQLDataDirectoryTypeInvalid ...
	SQLDataDirectoryTypeInvalid SQLDataDirectoryType = "Invalid"
	// SQLDataDirectoryTypeLog ...
	SQLDataDirectoryTypeLog SQLDataDirectoryType = "Log"
)

// PossibleSQLDataDirectoryTypeValues returns an array of possible values for the SQLDataDirectoryType const type.
func PossibleSQLDataDirectoryTypeValues() []SQLDataDirectoryType {
	return []SQLDataDirectoryType{SQLDataDirectoryTypeData, SQLDataDirectoryTypeInvalid, SQLDataDirectoryTypeLog}
}

// StorageType enumerates the values for storage type.
type StorageType string

const (
	// StorageTypeGeoRedundant ...
	StorageTypeGeoRedundant StorageType = "GeoRedundant"
	// StorageTypeInvalid ...
	StorageTypeInvalid StorageType = "Invalid"
	// StorageTypeLocallyRedundant ...
	StorageTypeLocallyRedundant StorageType = "LocallyRedundant"
)

// PossibleStorageTypeValues returns an array of possible values for the StorageType const type.
func PossibleStorageTypeValues() []StorageType {
	return []StorageType{StorageTypeGeoRedundant, StorageTypeInvalid, StorageTypeLocallyRedundant}
}

// StorageTypeState enumerates the values for storage type state.
type StorageTypeState string

const (
	// StorageTypeStateInvalid ...
	StorageTypeStateInvalid StorageTypeState = "Invalid"
	// StorageTypeStateLocked ...
	StorageTypeStateLocked StorageTypeState = "Locked"
	// StorageTypeStateUnlocked ...
	StorageTypeStateUnlocked StorageTypeState = "Unlocked"
)

// PossibleStorageTypeStateValues returns an array of possible values for the StorageTypeState const type.
func PossibleStorageTypeStateValues() []StorageTypeState {
	return []StorageTypeState{StorageTypeStateInvalid, StorageTypeStateLocked, StorageTypeStateUnlocked}
}

// SupportStatus enumerates the values for support status.
type SupportStatus string

const (
	// SupportStatusDefaultOFF ...
	SupportStatusDefaultOFF SupportStatus = "DefaultOFF"
	// SupportStatusDefaultON ...
	SupportStatusDefaultON SupportStatus = "DefaultON"
	// SupportStatusInvalid ...
	SupportStatusInvalid SupportStatus = "Invalid"
	// SupportStatusNotSupported ...
	SupportStatusNotSupported SupportStatus = "NotSupported"
	// SupportStatusSupported ...
	SupportStatusSupported SupportStatus = "Supported"
)

// PossibleSupportStatusValues returns an array of possible values for the SupportStatus const type.
func PossibleSupportStatusValues() []SupportStatus {
	return []SupportStatus{SupportStatusDefaultOFF, SupportStatusDefaultON, SupportStatusInvalid, SupportStatusNotSupported, SupportStatusSupported}
}

// Type enumerates the values for type.
type Type string

const (
	// TypeBackupProtectedItemCountSummary ...
	TypeBackupProtectedItemCountSummary Type = "BackupProtectedItemCountSummary"
	// TypeBackupProtectionContainerCountSummary ...
	TypeBackupProtectionContainerCountSummary Type = "BackupProtectionContainerCountSummary"
	// TypeInvalid ...
	TypeInvalid Type = "Invalid"
)

// PossibleTypeValues returns an array of possible values for the Type const type.
func PossibleTypeValues() []Type {
	return []Type{TypeBackupProtectedItemCountSummary, TypeBackupProtectionContainerCountSummary, TypeInvalid}
}

// TypeEnum enumerates the values for type enum.
type TypeEnum string

const (
	// TypeEnumCopyOnlyFull ...
	TypeEnumCopyOnlyFull TypeEnum = "CopyOnlyFull"
	// TypeEnumDifferential ...
	TypeEnumDifferential TypeEnum = "Differential"
	// TypeEnumFull ...
	TypeEnumFull TypeEnum = "Full"
	// TypeEnumInvalid ...
	TypeEnumInvalid TypeEnum = "Invalid"
	// TypeEnumLog ...
	TypeEnumLog TypeEnum = "Log"
)

// PossibleTypeEnumValues returns an array of possible values for the TypeEnum const type.
func PossibleTypeEnumValues() []TypeEnum {
	return []TypeEnum{TypeEnumCopyOnlyFull, TypeEnumDifferential, TypeEnumFull, TypeEnumInvalid, TypeEnumLog}
}

// UsagesUnit enumerates the values for usages unit.
type UsagesUnit string

const (
	// Bytes ...
	Bytes UsagesUnit = "Bytes"
	// BytesPerSecond ...
	BytesPerSecond UsagesUnit = "BytesPerSecond"
	// Count ...
	Count UsagesUnit = "Count"
	// CountPerSecond ...
	CountPerSecond UsagesUnit = "CountPerSecond"
	// Percent ...
	Percent UsagesUnit = "Percent"
	// Seconds ...
	Seconds UsagesUnit = "Seconds"
)

// PossibleUsagesUnitValues returns an array of possible values for the UsagesUnit const type.
func PossibleUsagesUnitValues() []UsagesUnit {
	return []UsagesUnit{Bytes, BytesPerSecond, Count, CountPerSecond, Percent, Seconds}
}

// ValidationStatus enumerates the values for validation status.
type ValidationStatus string

const (
	// ValidationStatusFailed ...
	ValidationStatusFailed ValidationStatus = "Failed"
	// ValidationStatusInvalid ...
	ValidationStatusInvalid ValidationStatus = "Invalid"
	// ValidationStatusSucceeded ...
	ValidationStatusSucceeded ValidationStatus = "Succeeded"
)

// PossibleValidationStatusValues returns an array of possible values for the ValidationStatus const type.
func PossibleValidationStatusValues() []ValidationStatus {
	return []ValidationStatus{ValidationStatusFailed, ValidationStatusInvalid, ValidationStatusSucceeded}
}

// WorkloadItemType enumerates the values for workload item type.
type WorkloadItemType string

const (
	// WorkloadItemTypeInvalid ...
	WorkloadItemTypeInvalid WorkloadItemType = "Invalid"
	// WorkloadItemTypeSAPAseDatabase ...
	WorkloadItemTypeSAPAseDatabase WorkloadItemType = "SAPAseDatabase"
	// WorkloadItemTypeSAPAseSystem ...
	WorkloadItemTypeSAPAseSystem WorkloadItemType = "SAPAseSystem"
	// WorkloadItemTypeSAPHanaDatabase ...
	WorkloadItemTypeSAPHanaDatabase WorkloadItemType = "SAPHanaDatabase"
	// WorkloadItemTypeSAPHanaSystem ...
	WorkloadItemTypeSAPHanaSystem WorkloadItemType = "SAPHanaSystem"
	// WorkloadItemTypeSQLDataBase ...
	WorkloadItemTypeSQLDataBase WorkloadItemType = "SQLDataBase"
	// WorkloadItemTypeSQLInstance ...
	WorkloadItemTypeSQLInstance WorkloadItemType = "SQLInstance"
)

// PossibleWorkloadItemTypeValues returns an array of possible values for the WorkloadItemType const type.
func PossibleWorkloadItemTypeValues() []WorkloadItemType {
	return []WorkloadItemType{WorkloadItemTypeInvalid, WorkloadItemTypeSAPAseDatabase, WorkloadItemTypeSAPAseSystem, WorkloadItemTypeSAPHanaDatabase, WorkloadItemTypeSAPHanaSystem, WorkloadItemTypeSQLDataBase, WorkloadItemTypeSQLInstance}
}

// WorkloadItemTypeBasicWorkloadItem enumerates the values for workload item type basic workload item.
type WorkloadItemTypeBasicWorkloadItem string

const (
	// WorkloadItemTypeAzureVMWorkloadItem ...
	WorkloadItemTypeAzureVMWorkloadItem WorkloadItemTypeBasicWorkloadItem = "AzureVmWorkloadItem"
	// WorkloadItemTypeSAPAseDatabase1 ...
	WorkloadItemTypeSAPAseDatabase1 WorkloadItemTypeBasicWorkloadItem = "SAPAseDatabase"
	// WorkloadItemTypeSAPAseSystem1 ...
	WorkloadItemTypeSAPAseSystem1 WorkloadItemTypeBasicWorkloadItem = "SAPAseSystem"
	// WorkloadItemTypeSAPHanaDatabase1 ...
	WorkloadItemTypeSAPHanaDatabase1 WorkloadItemTypeBasicWorkloadItem = "SAPHanaDatabase"
	// WorkloadItemTypeSAPHanaSystem1 ...
	WorkloadItemTypeSAPHanaSystem1 WorkloadItemTypeBasicWorkloadItem = "SAPHanaSystem"
	// WorkloadItemTypeSQLDataBase1 ...
	WorkloadItemTypeSQLDataBase1 WorkloadItemTypeBasicWorkloadItem = "SQLDataBase"
	// WorkloadItemTypeSQLInstance1 ...
	WorkloadItemTypeSQLInstance1 WorkloadItemTypeBasicWorkloadItem = "SQLInstance"
	// WorkloadItemTypeWorkloadItem ...
	WorkloadItemTypeWorkloadItem WorkloadItemTypeBasicWorkloadItem = "WorkloadItem"
)

// PossibleWorkloadItemTypeBasicWorkloadItemValues returns an array of possible values for the WorkloadItemTypeBasicWorkloadItem const type.
func PossibleWorkloadItemTypeBasicWorkloadItemValues() []WorkloadItemTypeBasicWorkloadItem {
	return []WorkloadItemTypeBasicWorkloadItem{WorkloadItemTypeAzureVMWorkloadItem, WorkloadItemTypeSAPAseDatabase1, WorkloadItemTypeSAPAseSystem1, WorkloadItemTypeSAPHanaDatabase1, WorkloadItemTypeSAPHanaSystem1, WorkloadItemTypeSQLDataBase1, WorkloadItemTypeSQLInstance1, WorkloadItemTypeWorkloadItem}
}

// WorkloadType enumerates the values for workload type.
type WorkloadType string

const (
	// WorkloadTypeAzureFileShare ...
	WorkloadTypeAzureFileShare WorkloadType = "AzureFileShare"
	// WorkloadTypeAzureSQLDb ...
	WorkloadTypeAzureSQLDb WorkloadType = "AzureSqlDb"
	// WorkloadTypeClient ...
	WorkloadTypeClient WorkloadType = "Client"
	// WorkloadTypeExchange ...
	WorkloadTypeExchange WorkloadType = "Exchange"
	// WorkloadTypeFileFolder ...
	WorkloadTypeFileFolder WorkloadType = "FileFolder"
	// WorkloadTypeGenericDataSource ...
	WorkloadTypeGenericDataSource WorkloadType = "GenericDataSource"
	// WorkloadTypeInvalid ...
	WorkloadTypeInvalid WorkloadType = "Invalid"
	// WorkloadTypeSAPAseDatabase ...
	WorkloadTypeSAPAseDatabase WorkloadType = "SAPAseDatabase"
	// WorkloadTypeSAPHanaDatabase ...
	WorkloadTypeSAPHanaDatabase WorkloadType = "SAPHanaDatabase"
	// WorkloadTypeSharepoint ...
	WorkloadTypeSharepoint WorkloadType = "Sharepoint"
	// WorkloadTypeSQLDataBase ...
	WorkloadTypeSQLDataBase WorkloadType = "SQLDataBase"
	// WorkloadTypeSQLDB ...
	WorkloadTypeSQLDB WorkloadType = "SQLDB"
	// WorkloadTypeSystemState ...
	WorkloadTypeSystemState WorkloadType = "SystemState"
	// WorkloadTypeVM ...
	WorkloadTypeVM WorkloadType = "VM"
	// WorkloadTypeVMwareVM ...
	WorkloadTypeVMwareVM WorkloadType = "VMwareVM"
)

// PossibleWorkloadTypeValues returns an array of possible values for the WorkloadType const type.
func PossibleWorkloadTypeValues() []WorkloadType {
	return []WorkloadType{WorkloadTypeAzureFileShare, WorkloadTypeAzureSQLDb, WorkloadTypeClient, WorkloadTypeExchange, WorkloadTypeFileFolder, WorkloadTypeGenericDataSource, WorkloadTypeInvalid, WorkloadTypeSAPAseDatabase, WorkloadTypeSAPHanaDatabase, WorkloadTypeSharepoint, WorkloadTypeSQLDataBase, WorkloadTypeSQLDB, WorkloadTypeSystemState, WorkloadTypeVM, WorkloadTypeVMwareVM}
}