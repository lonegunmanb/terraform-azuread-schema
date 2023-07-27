package generated

import (
	tfjson "github.com/hashicorp/terraform-json"
	"github.com/lonegunmanb/terraform-azuread-schema/v2/generated/data"
	"github.com/lonegunmanb/terraform-azuread-schema/v2/generated/resource"
)

var Resources map[string]*tfjson.Schema
var DataSources map[string]*tfjson.Schema

func init() {
	resources := make(map[string]*tfjson.Schema, 0)
	dataSources := make(map[string]*tfjson.Schema, 0)  
	resources["azuread_access_package"] = resource.AzureadAccessPackageSchema()  
	resources["azuread_access_package_assignment_policy"] = resource.AzureadAccessPackageAssignmentPolicySchema()  
	resources["azuread_access_package_catalog"] = resource.AzureadAccessPackageCatalogSchema()  
	resources["azuread_access_package_catalog_role_assignment"] = resource.AzureadAccessPackageCatalogRoleAssignmentSchema()  
	resources["azuread_access_package_resource_catalog_association"] = resource.AzureadAccessPackageResourceCatalogAssociationSchema()  
	resources["azuread_access_package_resource_package_association"] = resource.AzureadAccessPackageResourcePackageAssociationSchema()  
	resources["azuread_administrative_unit"] = resource.AzureadAdministrativeUnitSchema()  
	resources["azuread_administrative_unit_member"] = resource.AzureadAdministrativeUnitMemberSchema()  
	resources["azuread_administrative_unit_role_member"] = resource.AzureadAdministrativeUnitRoleMemberSchema()  
	resources["azuread_app_role_assignment"] = resource.AzureadAppRoleAssignmentSchema()  
	resources["azuread_application"] = resource.AzureadApplicationSchema()  
	resources["azuread_application_certificate"] = resource.AzureadApplicationCertificateSchema()  
	resources["azuread_application_federated_identity_credential"] = resource.AzureadApplicationFederatedIdentityCredentialSchema()  
	resources["azuread_application_password"] = resource.AzureadApplicationPasswordSchema()  
	resources["azuread_application_pre_authorized"] = resource.AzureadApplicationPreAuthorizedSchema()  
	resources["azuread_claims_mapping_policy"] = resource.AzureadClaimsMappingPolicySchema()  
	resources["azuread_conditional_access_policy"] = resource.AzureadConditionalAccessPolicySchema()  
	resources["azuread_custom_directory_role"] = resource.AzureadCustomDirectoryRoleSchema()  
	resources["azuread_directory_role"] = resource.AzureadDirectoryRoleSchema()  
	resources["azuread_directory_role_assignment"] = resource.AzureadDirectoryRoleAssignmentSchema()  
	resources["azuread_directory_role_member"] = resource.AzureadDirectoryRoleMemberSchema()  
	resources["azuread_group"] = resource.AzureadGroupSchema()  
	resources["azuread_group_member"] = resource.AzureadGroupMemberSchema()  
	resources["azuread_invitation"] = resource.AzureadInvitationSchema()  
	resources["azuread_named_location"] = resource.AzureadNamedLocationSchema()  
	resources["azuread_service_principal"] = resource.AzureadServicePrincipalSchema()  
	resources["azuread_service_principal_certificate"] = resource.AzureadServicePrincipalCertificateSchema()  
	resources["azuread_service_principal_claims_mapping_policy_assignment"] = resource.AzureadServicePrincipalClaimsMappingPolicyAssignmentSchema()  
	resources["azuread_service_principal_delegated_permission_grant"] = resource.AzureadServicePrincipalDelegatedPermissionGrantSchema()  
	resources["azuread_service_principal_password"] = resource.AzureadServicePrincipalPasswordSchema()  
	resources["azuread_service_principal_token_signing_certificate"] = resource.AzureadServicePrincipalTokenSigningCertificateSchema()  
	resources["azuread_synchronization_job"] = resource.AzureadSynchronizationJobSchema()  
	resources["azuread_synchronization_secret"] = resource.AzureadSynchronizationSecretSchema()  
	resources["azuread_user"] = resource.AzureadUserSchema()  
	resources["azuread_user_flow_attribute"] = resource.AzureadUserFlowAttributeSchema()  
	dataSources["azuread_access_package"] = data.AzureadAccessPackageSchema()  
	dataSources["azuread_access_package_catalog"] = data.AzureadAccessPackageCatalogSchema()  
	dataSources["azuread_access_package_catalog_role"] = data.AzureadAccessPackageCatalogRoleSchema()  
	dataSources["azuread_administrative_unit"] = data.AzureadAdministrativeUnitSchema()  
	dataSources["azuread_application"] = data.AzureadApplicationSchema()  
	dataSources["azuread_application_published_app_ids"] = data.AzureadApplicationPublishedAppIdsSchema()  
	dataSources["azuread_application_template"] = data.AzureadApplicationTemplateSchema()  
	dataSources["azuread_client_config"] = data.AzureadClientConfigSchema()  
	dataSources["azuread_directory_object"] = data.AzureadDirectoryObjectSchema()  
	dataSources["azuread_directory_role_templates"] = data.AzureadDirectoryRoleTemplatesSchema()  
	dataSources["azuread_directory_roles"] = data.AzureadDirectoryRolesSchema()  
	dataSources["azuread_domains"] = data.AzureadDomainsSchema()  
	dataSources["azuread_group"] = data.AzureadGroupSchema()  
	dataSources["azuread_groups"] = data.AzureadGroupsSchema()  
	dataSources["azuread_named_location"] = data.AzureadNamedLocationSchema()  
	dataSources["azuread_service_principal"] = data.AzureadServicePrincipalSchema()  
	dataSources["azuread_service_principals"] = data.AzureadServicePrincipalsSchema()  
	dataSources["azuread_user"] = data.AzureadUserSchema()  
	dataSources["azuread_users"] = data.AzureadUsersSchema()  
	Resources = resources
	DataSources = dataSources
}