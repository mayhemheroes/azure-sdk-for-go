//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package generated

const (
	ModuleName    = "generated"
	ModuleVersion = "v0.2.0"
)

// ActionType - The type of the action.
type ActionType string

const (
	ActionTypeEmailContacts ActionType = "EmailContacts"
	ActionTypeAutoRenew     ActionType = "AutoRenew"
)

// PossibleActionTypeValues returns the possible values for the ActionType const type.
func PossibleActionTypeValues() []ActionType {
	return []ActionType{
		ActionTypeEmailContacts,
		ActionTypeAutoRenew,
	}
}

// ToPtr returns a *ActionType pointing to the current value.
func (c ActionType) ToPtr() *ActionType {
	return &c
}

// DataAction - Supported permissions for data actions.
type DataAction string

const (
	// DataActionBackupHsmKeys - Backup HSM keys.
	DataActionBackupHsmKeys DataAction = "Microsoft.KeyVault/managedHsm/keys/backup/action"
	// DataActionCreateHsmKey - Create an HSM key.
	DataActionCreateHsmKey DataAction = "Microsoft.KeyVault/managedHsm/keys/create"
	// DataActionDecryptHsmKey - Decrypt using an HSM key.
	DataActionDecryptHsmKey DataAction = "Microsoft.KeyVault/managedHsm/keys/decrypt/action"
	// DataActionDeleteHsmKey - Delete an HSM key.
	DataActionDeleteHsmKey DataAction = "Microsoft.KeyVault/managedHsm/keys/delete"
	// DataActionDeleteRoleAssignment - Delete role assignment.
	DataActionDeleteRoleAssignment DataAction = "Microsoft.KeyVault/managedHsm/roleAssignments/delete/action"
	// DataActionDeleteRoleDefinition - Delete role definition.
	DataActionDeleteRoleDefinition DataAction = "Microsoft.KeyVault/managedHsm/roleDefinitions/delete/action"
	// DataActionDownloadHsmSecurityDomain - Download an HSM security domain.
	DataActionDownloadHsmSecurityDomain DataAction = "Microsoft.KeyVault/managedHsm/securitydomain/download/action"
	// DataActionDownloadHsmSecurityDomainStatus - Check status of HSM security domain download.
	DataActionDownloadHsmSecurityDomainStatus DataAction = "Microsoft.KeyVault/managedHsm/securitydomain/download/read"
	// DataActionEncryptHsmKey - Encrypt using an HSM key.
	DataActionEncryptHsmKey DataAction = "Microsoft.KeyVault/managedHsm/keys/encrypt/action"
	// DataActionExportHsmKey - Export an HSM key.
	DataActionExportHsmKey DataAction = "Microsoft.KeyVault/managedHsm/keys/export/action"
	// DataActionGetRoleAssignment - Get role assignment.
	DataActionGetRoleAssignment DataAction = "Microsoft.KeyVault/managedHsm/roleAssignments/read/action"
	// DataActionImportHsmKey - Import an HSM key.
	DataActionImportHsmKey DataAction = "Microsoft.KeyVault/managedHsm/keys/import/action"
	// DataActionPurgeDeletedHsmKey - Purge a deleted HSM key.
	DataActionPurgeDeletedHsmKey DataAction = "Microsoft.KeyVault/managedHsm/keys/deletedKeys/delete"
	// DataActionRandomNumbersGenerate - Generate random numbers.
	DataActionRandomNumbersGenerate DataAction = "Microsoft.KeyVault/managedHsm/rng/action"
	// DataActionReadDeletedHsmKey - Read deleted HSM key.
	DataActionReadDeletedHsmKey DataAction = "Microsoft.KeyVault/managedHsm/keys/deletedKeys/read/action"
	// DataActionReadHsmBackupStatus - Read an HSM backup status.
	DataActionReadHsmBackupStatus DataAction = "Microsoft.KeyVault/managedHsm/backup/status/action"
	// DataActionReadHsmKey - Read HSM key metadata.
	DataActionReadHsmKey DataAction = "Microsoft.KeyVault/managedHsm/keys/read/action"
	// DataActionReadHsmRestoreStatus - Read an HSM restore status.
	DataActionReadHsmRestoreStatus DataAction = "Microsoft.KeyVault/managedHsm/restore/status/action"
	// DataActionReadHsmSecurityDomainStatus - Check the status of the HSM security domain exchange file.
	DataActionReadHsmSecurityDomainStatus DataAction = "Microsoft.KeyVault/managedHsm/securitydomain/upload/read"
	// DataActionReadHsmSecurityDomainTransferKey - Download an HSM security domain transfer key.
	DataActionReadHsmSecurityDomainTransferKey DataAction = "Microsoft.KeyVault/managedHsm/securitydomain/transferkey/read"
	// DataActionReadRoleDefinition - Get role definition.
	DataActionReadRoleDefinition DataAction = "Microsoft.KeyVault/managedHsm/roleDefinitions/read/action"
	// DataActionRecoverDeletedHsmKey - Recover deleted HSM key.
	DataActionRecoverDeletedHsmKey DataAction = "Microsoft.KeyVault/managedHsm/keys/deletedKeys/recover/action"
	// DataActionReleaseKey - Release an HSM key using Secure Key Release.
	DataActionReleaseKey DataAction = "Microsoft.KeyVault/managedHsm/keys/release/action"
	// DataActionRestoreHsmKeys - Restore HSM keys.
	DataActionRestoreHsmKeys DataAction = "Microsoft.KeyVault/managedHsm/keys/restore/action"
	// DataActionSignHsmKey - Sign using an HSM key.
	DataActionSignHsmKey DataAction = "Microsoft.KeyVault/managedHsm/keys/sign/action"
	// DataActionStartHsmBackup - Start an HSM backup.
	DataActionStartHsmBackup DataAction = "Microsoft.KeyVault/managedHsm/backup/start/action"
	// DataActionStartHsmRestore - Start an HSM restore.
	DataActionStartHsmRestore DataAction = "Microsoft.KeyVault/managedHsm/restore/start/action"
	// DataActionUnwrapHsmKey - Unwrap using an HSM key.
	DataActionUnwrapHsmKey DataAction = "Microsoft.KeyVault/managedHsm/keys/unwrap/action"
	// DataActionUploadHsmSecurityDomain - Upload an HSM security domain.
	DataActionUploadHsmSecurityDomain DataAction = "Microsoft.KeyVault/managedHsm/securitydomain/upload/action"
	// DataActionVerifyHsmKey - Verify using an HSM key.
	DataActionVerifyHsmKey DataAction = "Microsoft.KeyVault/managedHsm/keys/verify/action"
	// DataActionWrapHsmKey - Wrap using an HSM key.
	DataActionWrapHsmKey DataAction = "Microsoft.KeyVault/managedHsm/keys/wrap/action"
	// DataActionWriteHsmKey - Update an HSM key.
	DataActionWriteHsmKey DataAction = "Microsoft.KeyVault/managedHsm/keys/write/action"
	// DataActionWriteRoleAssignment - Create or update role assignment.
	DataActionWriteRoleAssignment DataAction = "Microsoft.KeyVault/managedHsm/roleAssignments/write/action"
	// DataActionWriteRoleDefinition - Create or update role definition.
	DataActionWriteRoleDefinition DataAction = "Microsoft.KeyVault/managedHsm/roleDefinitions/write/action"
)

// PossibleDataActionValues returns the possible values for the DataAction const type.
func PossibleDataActionValues() []DataAction {
	return []DataAction{
		DataActionBackupHsmKeys,
		DataActionCreateHsmKey,
		DataActionDecryptHsmKey,
		DataActionDeleteHsmKey,
		DataActionDeleteRoleAssignment,
		DataActionDeleteRoleDefinition,
		DataActionDownloadHsmSecurityDomain,
		DataActionDownloadHsmSecurityDomainStatus,
		DataActionEncryptHsmKey,
		DataActionExportHsmKey,
		DataActionGetRoleAssignment,
		DataActionImportHsmKey,
		DataActionPurgeDeletedHsmKey,
		DataActionRandomNumbersGenerate,
		DataActionReadDeletedHsmKey,
		DataActionReadHsmBackupStatus,
		DataActionReadHsmKey,
		DataActionReadHsmRestoreStatus,
		DataActionReadHsmSecurityDomainStatus,
		DataActionReadHsmSecurityDomainTransferKey,
		DataActionReadRoleDefinition,
		DataActionRecoverDeletedHsmKey,
		DataActionReleaseKey,
		DataActionRestoreHsmKeys,
		DataActionSignHsmKey,
		DataActionStartHsmBackup,
		DataActionStartHsmRestore,
		DataActionUnwrapHsmKey,
		DataActionUploadHsmSecurityDomain,
		DataActionVerifyHsmKey,
		DataActionWrapHsmKey,
		DataActionWriteHsmKey,
		DataActionWriteRoleAssignment,
		DataActionWriteRoleDefinition,
	}
}

// ToPtr returns a *DataAction pointing to the current value.
func (c DataAction) ToPtr() *DataAction {
	return &c
}

// DeletionRecoveryLevel - Reflects the deletion recovery level currently in effect for certificates in the current vault.
// If it contains 'Purgeable', the certificate can be permanently deleted by a privileged user; otherwise,
// only the system can purge the certificate, at the end of the retention interval.
type DeletionRecoveryLevel string

const (
	// DeletionRecoveryLevelCustomizedRecoverable - Denotes a vault state in which deletion is recoverable without the possibility
	// for immediate and permanent deletion (i.e. purge when 7<= SoftDeleteRetentionInDays < 90).This level guarantees the recoverability
	// of the deleted entity during the retention interval and while the subscription is still available.
	DeletionRecoveryLevelCustomizedRecoverable DeletionRecoveryLevel = "CustomizedRecoverable"
	// DeletionRecoveryLevelCustomizedRecoverableProtectedSubscription - Denotes a vault and subscription state in which deletion
	// is recoverable, immediate and permanent deletion (i.e. purge) is not permitted, and in which the subscription itself cannot
	// be permanently canceled when 7<= SoftDeleteRetentionInDays < 90. This level guarantees the recoverability of the deleted
	// entity during the retention interval, and also reflects the fact that the subscription itself cannot be cancelled.
	DeletionRecoveryLevelCustomizedRecoverableProtectedSubscription DeletionRecoveryLevel = "CustomizedRecoverable+ProtectedSubscription"
	// DeletionRecoveryLevelCustomizedRecoverablePurgeable - Denotes a vault state in which deletion is recoverable, and which
	// also permits immediate and permanent deletion (i.e. purge when 7<= SoftDeleteRetentionInDays < 90). This level guarantees
	// the recoverability of the deleted entity during the retention interval, unless a Purge operation is requested, or the subscription
	// is cancelled.
	DeletionRecoveryLevelCustomizedRecoverablePurgeable DeletionRecoveryLevel = "CustomizedRecoverable+Purgeable"
	// DeletionRecoveryLevelPurgeable - Denotes a vault state in which deletion is an irreversible operation, without the possibility
	// for recovery. This level corresponds to no protection being available against a Delete operation; the data is irretrievably
	// lost upon accepting a Delete operation at the entity level or higher (vault, resource group, subscription etc.)
	DeletionRecoveryLevelPurgeable DeletionRecoveryLevel = "Purgeable"
	// DeletionRecoveryLevelRecoverable - Denotes a vault state in which deletion is recoverable without the possibility for immediate
	// and permanent deletion (i.e. purge). This level guarantees the recoverability of the deleted entity during the retention
	// interval(90 days) and while the subscription is still available. System wil permanently delete it after 90 days, if not
	// recovered
	DeletionRecoveryLevelRecoverable DeletionRecoveryLevel = "Recoverable"
	// DeletionRecoveryLevelRecoverableProtectedSubscription - Denotes a vault and subscription state in which deletion is recoverable
	// within retention interval (90 days), immediate and permanent deletion (i.e. purge) is not permitted, and in which the subscription
	// itself cannot be permanently canceled. System wil permanently delete it after 90 days, if not recovered
	DeletionRecoveryLevelRecoverableProtectedSubscription DeletionRecoveryLevel = "Recoverable+ProtectedSubscription"
	// DeletionRecoveryLevelRecoverablePurgeable - Denotes a vault state in which deletion is recoverable, and which also permits
	// immediate and permanent deletion (i.e. purge). This level guarantees the recoverability of the deleted entity during the
	// retention interval (90 days), unless a Purge operation is requested, or the subscription is cancelled. System wil permanently
	// delete it after 90 days, if not recovered
	DeletionRecoveryLevelRecoverablePurgeable DeletionRecoveryLevel = "Recoverable+Purgeable"
)

// PossibleDeletionRecoveryLevelValues returns the possible values for the DeletionRecoveryLevel const type.
func PossibleDeletionRecoveryLevelValues() []DeletionRecoveryLevel {
	return []DeletionRecoveryLevel{
		DeletionRecoveryLevelCustomizedRecoverable,
		DeletionRecoveryLevelCustomizedRecoverableProtectedSubscription,
		DeletionRecoveryLevelCustomizedRecoverablePurgeable,
		DeletionRecoveryLevelPurgeable,
		DeletionRecoveryLevelRecoverable,
		DeletionRecoveryLevelRecoverableProtectedSubscription,
		DeletionRecoveryLevelRecoverablePurgeable,
	}
}

// ToPtr returns a *DeletionRecoveryLevel pointing to the current value.
func (c DeletionRecoveryLevel) ToPtr() *DeletionRecoveryLevel {
	return &c
}

// JSONWebKeyCurveName - Elliptic curve name. For valid values, see JsonWebKeyCurveName.
type JSONWebKeyCurveName string

const (
	JSONWebKeyCurveNameP256  JSONWebKeyCurveName = "P-256"
	JSONWebKeyCurveNameP256K JSONWebKeyCurveName = "P-256K"
	JSONWebKeyCurveNameP384  JSONWebKeyCurveName = "P-384"
	JSONWebKeyCurveNameP521  JSONWebKeyCurveName = "P-521"
)

// PossibleJSONWebKeyCurveNameValues returns the possible values for the JSONWebKeyCurveName const type.
func PossibleJSONWebKeyCurveNameValues() []JSONWebKeyCurveName {
	return []JSONWebKeyCurveName{
		JSONWebKeyCurveNameP256,
		JSONWebKeyCurveNameP256K,
		JSONWebKeyCurveNameP384,
		JSONWebKeyCurveNameP521,
	}
}

// ToPtr returns a *JSONWebKeyCurveName pointing to the current value.
func (c JSONWebKeyCurveName) ToPtr() *JSONWebKeyCurveName {
	return &c
}

// JSONWebKeyType - The type of key pair to be used for the certificate.
type JSONWebKeyType string

const (
	JSONWebKeyTypeEC     JSONWebKeyType = "EC"
	JSONWebKeyTypeECHSM  JSONWebKeyType = "EC-HSM"
	JSONWebKeyTypeOct    JSONWebKeyType = "oct"
	JSONWebKeyTypeOctHSM JSONWebKeyType = "oct-HSM"
	JSONWebKeyTypeRSA    JSONWebKeyType = "RSA"
	JSONWebKeyTypeRSAHSM JSONWebKeyType = "RSA-HSM"
)

// PossibleJSONWebKeyTypeValues returns the possible values for the JSONWebKeyType const type.
func PossibleJSONWebKeyTypeValues() []JSONWebKeyType {
	return []JSONWebKeyType{
		JSONWebKeyTypeEC,
		JSONWebKeyTypeECHSM,
		JSONWebKeyTypeOct,
		JSONWebKeyTypeOctHSM,
		JSONWebKeyTypeRSA,
		JSONWebKeyTypeRSAHSM,
	}
}

// ToPtr returns a *JSONWebKeyType pointing to the current value.
func (c JSONWebKeyType) ToPtr() *JSONWebKeyType {
	return &c
}

type KeyUsageType string

const (
	KeyUsageTypeCRLSign          KeyUsageType = "cRLSign"
	KeyUsageTypeDataEncipherment KeyUsageType = "dataEncipherment"
	KeyUsageTypeDecipherOnly     KeyUsageType = "decipherOnly"
	KeyUsageTypeDigitalSignature KeyUsageType = "digitalSignature"
	KeyUsageTypeEncipherOnly     KeyUsageType = "encipherOnly"
	KeyUsageTypeKeyAgreement     KeyUsageType = "keyAgreement"
	KeyUsageTypeKeyCertSign      KeyUsageType = "keyCertSign"
	KeyUsageTypeKeyEncipherment  KeyUsageType = "keyEncipherment"
	KeyUsageTypeNonRepudiation   KeyUsageType = "nonRepudiation"
)

// PossibleKeyUsageTypeValues returns the possible values for the KeyUsageType const type.
func PossibleKeyUsageTypeValues() []KeyUsageType {
	return []KeyUsageType{
		KeyUsageTypeCRLSign,
		KeyUsageTypeDataEncipherment,
		KeyUsageTypeDecipherOnly,
		KeyUsageTypeDigitalSignature,
		KeyUsageTypeEncipherOnly,
		KeyUsageTypeKeyAgreement,
		KeyUsageTypeKeyCertSign,
		KeyUsageTypeKeyEncipherment,
		KeyUsageTypeNonRepudiation,
	}
}

// ToPtr returns a *KeyUsageType pointing to the current value.
func (c KeyUsageType) ToPtr() *KeyUsageType {
	return &c
}

// RoleDefinitionType - The role definition type.
type RoleDefinitionType string

const (
	RoleDefinitionTypeMicrosoftAuthorizationRoleDefinitions RoleDefinitionType = "Microsoft.Authorization/roleDefinitions"
)

// PossibleRoleDefinitionTypeValues returns the possible values for the RoleDefinitionType const type.
func PossibleRoleDefinitionTypeValues() []RoleDefinitionType {
	return []RoleDefinitionType{
		RoleDefinitionTypeMicrosoftAuthorizationRoleDefinitions,
	}
}

// ToPtr returns a *RoleDefinitionType pointing to the current value.
func (c RoleDefinitionType) ToPtr() *RoleDefinitionType {
	return &c
}

// RoleScope - The role scope.
type RoleScope string

const (
	// RoleScopeGlobal - Global scope
	RoleScopeGlobal RoleScope = "/"
	// RoleScopeKeys - Keys scope
	RoleScopeKeys RoleScope = "/keys"
)

// PossibleRoleScopeValues returns the possible values for the RoleScope const type.
func PossibleRoleScopeValues() []RoleScope {
	return []RoleScope{
		RoleScopeGlobal,
		RoleScopeKeys,
	}
}

// ToPtr returns a *RoleScope pointing to the current value.
func (c RoleScope) ToPtr() *RoleScope {
	return &c
}

// RoleType - The role type.
type RoleType string

const (
	// RoleTypeBuiltInRole - Built in role.
	RoleTypeBuiltInRole RoleType = "AKVBuiltInRole"
	// RoleTypeCustomRole - Custom role.
	RoleTypeCustomRole RoleType = "CustomRole"
)

// PossibleRoleTypeValues returns the possible values for the RoleType const type.
func PossibleRoleTypeValues() []RoleType {
	return []RoleType{
		RoleTypeBuiltInRole,
		RoleTypeCustomRole,
	}
}

// ToPtr returns a *RoleType pointing to the current value.
func (c RoleType) ToPtr() *RoleType {
	return &c
}
