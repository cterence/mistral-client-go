# AuditLogEventType

## Enum


* `USER_CREATE` (value: `"user.create"`)

* `USER_DELETE` (value: `"user.delete"`)

* `USER_LOG_IN` (value: `"user.log_in"`)

* `USER_INFO_UPDATE` (value: `"user.info.update"`)

* `USER_PASSWORD_UPDATE` (value: `"user.password.update"`)

* `USER_PHONE_NUMBER_VERIFY` (value: `"user.phone_number.verify"`)

* `USER_ORGANIZATION_LEAVE` (value: `"user.organization.leave"`)

* `USER_ORGANIZATION_ROLE_UPDATE` (value: `"user.organization.role.update"`)

* `USER_ORGANIZATION_DELETE` (value: `"user.organization.delete"`)

* `USER_ORGANIZATION_JOIN` (value: `"user.organization.join"`)

* `ORGANIZATION_CREATE` (value: `"organization.create"`)

* `ORGANIZATION_DELETE` (value: `"organization.delete"`)

* `ORGANIZATION_UPDATE` (value: `"organization.update"`)

* `ORGANIZATION_INVITE_SEND` (value: `"organization.invite.send"`)

* `ORGANIZATION_INVITE_RESEND` (value: `"organization.invite.resend"`)

* `ORGANIZATION_INVITE_ACCEPTED` (value: `"organization.invite.accepted"`)

* `ORGANIZATION_INVITE_REVOKED` (value: `"organization.invite.revoked"`)

* `ORGANIZATION_JOIN_BY_EMAIL_DOMAIN` (value: `"organization.join_by_email_domain"`)

* `ORGANIZATION_DOMAIN_VERIFICATION_DISABLE` (value: `"organization.domain_verification.disable"`)

* `ORGANIZATION_DOMAIN_VERIFICATION_ENABLE` (value: `"organization.domain_verification.enable"`)

* `ORGANIZATION_EMAIL_DOMAIN_AUTHENTICATION_ENABLE` (value: `"organization.email_domain_authentication.enable"`)

* `ORGANIZATION_EMAIL_DOMAIN_AUTHENTICATION_DISABLE` (value: `"organization.email_domain_authentication.disable"`)

* `ORGANIZATION_SAML_AUTHENTICATION_ENABLE` (value: `"organization.saml_authentication.enable"`)

* `ORGANIZATION_SAML_AUTHENTICATION_DISABLE` (value: `"organization.saml_authentication.disable"`)

* `ORGANIZATION_SAML_AUTHENTICATION_SSO_USER_PROVISIONING_UPDATE` (value: `"organization.saml_authentication.sso_user_provisioning.update"`)

* `ORGANIZATION_SCIM_SYNC_TRIGGER` (value: `"organization.scim_sync.trigger"`)

* `ORGANIZATION_SCIM_SECRET_ROTATE` (value: `"organization.scim.secret.rotate"`)

* `ORGANIZATION_SEAT_AUTO_ASSIGN_ENABLE` (value: `"organization.seat_auto_assign.enable"`)

* `ORGANIZATION_SEAT_AUTO_ASSIGN_DISABLE` (value: `"organization.seat_auto_assign.disable"`)

* `ORGANIZATION_KIND_UPDATE` (value: `"organization.kind.update"`)

* `ORGANIZATION_SSO_SEAT_AUTO_ASSIGN_ENABLE` (value: `"organization.sso_seat_auto_assign.enable"`)

* `ORGANIZATION_SSO_SEAT_AUTO_ASSIGN_DISABLE` (value: `"organization.sso_seat_auto_assign.disable"`)

* `WORKSPACE_CREATE` (value: `"workspace.create"`)

* `WORKSPACE_UPDATE` (value: `"workspace.update"`)

* `WORKSPACE_DELETE` (value: `"workspace.delete"`)

* `WORKSPACE_MEMBER_ADD` (value: `"workspace.member.add"`)

* `WORKSPACE_MEMBER_REMOVE` (value: `"workspace.member.remove"`)

* `WORKSPACE_MEMBER_ROLE_UPDATE` (value: `"workspace.member.role.update"`)

* `USER_GROUP_CREATE` (value: `"user_group.create"`)

* `USER_GROUP_UPDATE` (value: `"user_group.update"`)

* `USER_GROUP_DELETE` (value: `"user_group.delete"`)

* `USER_GROUP_MEMBER_ADD` (value: `"user_group.member.add"`)

* `USER_GROUP_MEMBER_REMOVE` (value: `"user_group.member.remove"`)

* `USER_GROUP_WORKSPACE_PROVISION` (value: `"user_group.workspace.provision"`)

* `USER_GROUP_WORKSPACE_DEPROVISION` (value: `"user_group.workspace.deprovision"`)

* `USER_GROUP_WORKSPACE_UPDATE` (value: `"user_group.workspace.update"`)

* `BILLING_INFORMATION_UPDATED` (value: `"billing.information.updated"`)

* `BILLING_PAYMENT_METHOD_ADDED` (value: `"billing.payment_method.added"`)

* `BILLING_PAYMENT_METHOD_REMOVED` (value: `"billing.payment_method.removed"`)

* `BILLING_PAYMENT_METHOD_DEFAULT_CHANGED` (value: `"billing.payment_method.default_changed"`)

* `BILLING_INVOICE_RETRIED` (value: `"billing.invoice.retried"`)

* `BILLING_SUBSCRIPTION_SUBSCRIBE` (value: `"billing.subscription.subscribe"`)

* `BILLING_SUBSCRIPTION_UNSUBSCRIBE` (value: `"billing.subscription.unsubscribe"`)

* `BILLING_SUBSCRIPTION_CANCEL_UNSUBSCRIBE` (value: `"billing.subscription.cancel_unsubscribe"`)

* `BILLING_SUBSCRIPTION_ADD_SEATS` (value: `"billing.subscription.add_seats"`)

* `BILLING_SUBSCRIPTION_REMOVE_SEATS` (value: `"billing.subscription.remove_seats"`)

* `BILLING_SUBSCRIPTION_TERMINATION_DATE_UPDATED` (value: `"billing.subscription.termination_date.updated"`)

* `BILLING_CREDITS_ADDED` (value: `"billing.credits.added"`)

* `BILLING_GIFT_CODE_USED` (value: `"billing.gift_code.used"`)

* `BILLING_CREDIT_COUPON_USED` (value: `"billing.credit_coupon.used"`)

* `BILLING_SUBSCRIPTION_COUPON_USED` (value: `"billing.subscription_coupon.used"`)

* `BILLING_MONTHLY_LIMIT_UPDATED` (value: `"billing.monthly_limit.updated"`)

* `BILLING_WORKSPACE_MONTHLY_LIMIT_UPDATED` (value: `"billing.workspace_monthly_limit.updated"`)

* `BILLING_SHARED_BUDGET_OVERRIDE_UPDATED` (value: `"billing.shared_budget_override.updated"`)

* `BILLING_AUTO_RECHARGE_UPDATED` (value: `"billing.auto_recharge.updated"`)

* `BILLING_SEAT_GRANT` (value: `"billing.seat.grant"`)

* `BILLING_SEAT_REVOKE` (value: `"billing.seat.revoke"`)

* `BILLING_PRIORITY_SERVICE_TIER_UPDATED` (value: `"billing.priority_service_tier.updated"`)

* `BILLING_PRIORITY_SERVICE_TIER_REMOVED` (value: `"billing.priority_service_tier.removed"`)

* `LE_CHAT_CONVERSATION_CREATED` (value: `"le_chat.conversation.created"`)

* `LE_CHAT_TEMPORARY_CONVERSATION_CREATED` (value: `"le_chat.temporary_conversation.created"`)

* `LE_CHAT_TEMPORARY_CONVERSATION_CONVERTED` (value: `"le_chat.temporary_conversation.converted"`)

* `LE_CHAT_CONVERSATION_DELETED` (value: `"le_chat.conversation.deleted"`)

* `LE_CHAT_CONVERSATION_BATCH_DELETED` (value: `"le_chat.conversation_batch.deleted"`)

* `LE_CHAT_CONVERSATION_PUBLIC_SHARING_ENABLED` (value: `"le_chat.conversation.public_sharing.enabled"`)

* `LE_CHAT_CONVERSATION_PUBLIC_SHARING_DISABLED` (value: `"le_chat.conversation.public_sharing.disabled"`)

* `LE_CHAT_FLASH_ANSWERS_ENABLED` (value: `"le_chat.flash_answers.enabled"`)

* `LE_CHAT_FLASH_ANSWERS_DISABLED` (value: `"le_chat.flash_answers.disabled"`)

* `LE_CHAT_LOCALISATION_SHARING_ENABLED` (value: `"le_chat.localisation_sharing.enabled"`)

* `LE_CHAT_LOCALISATION_SHARING_DISABLED` (value: `"le_chat.localisation_sharing.disabled"`)

* `LE_CHAT_MEMORIES_ENABLED` (value: `"le_chat.memories.enabled"`)

* `LE_CHAT_MEMORIES_DISABLED` (value: `"le_chat.memories.disabled"`)

* `LE_CHAT_DATA_TRAINING_ENABLED` (value: `"le_chat.data.training_enabled"`)

* `LE_CHAT_DATA_TRAINING_DISABLED` (value: `"le_chat.data.training_disabled"`)

* `LE_CHAT_ACTIONS_EXTERNAL_LINK` (value: `"le_chat.actions.external_link"`)

* `ADMIN_API_KEY_CREATED` (value: `"admin_api_key.created"`)

* `ADMIN_API_KEY_DELETE` (value: `"admin_api_key.delete"`)

* `API_KEY_CREATE` (value: `"api_key.create"`)

* `API_KEY_ROTATE` (value: `"api_key.rotate"`)

* `API_KEY_DELETE` (value: `"api_key.delete"`)

* `API_KEY_POLICY_UPDATE` (value: `"api_key_policy.update"`)

* `SECRET_STORE_ENTRY_CREATE` (value: `"secret_store.entry.create"`)

* `SECRET_STORE_ENTRY_UPDATE` (value: `"secret_store.entry.update"`)

* `SECRET_STORE_ENTRY_DELETE` (value: `"secret_store.entry.delete"`)

* `SERVICE_ACCOUNT_CREATE` (value: `"service_account.create"`)

* `SERVICE_ACCOUNT_UPDATE` (value: `"service_account.update"`)

* `SERVICE_ACCOUNT_DELETE` (value: `"service_account.delete"`)

* `SERVICE_ACCOUNT_CLIENT_SECRET_CREATE` (value: `"service_account.client_secret.create"`)

* `SERVICE_ACCOUNT_CLIENT_SECRET_DELETE` (value: `"service_account.client_secret.delete"`)

* `SERVICE_ACCOUNT_ROLES_SET` (value: `"service_account.roles.set"`)

* `WORKLOAD_IDENTITY_CREDENTIAL_REGISTRATION_CREATE` (value: `"workload_identity.credential_registration.create"`)

* `WORKLOAD_IDENTITY_CREDENTIAL_CREATE` (value: `"workload_identity.credential.create"`)

* `TRUSTED_ISSUER_CREATE` (value: `"trusted_issuer.create"`)

* `TRUSTED_ISSUER_UPDATE` (value: `"trusted_issuer.update"`)

* `TRUSTED_ISSUER_DELETE` (value: `"trusted_issuer.delete"`)

* `AGENT_CREATE` (value: `"agent.create"`)

* `AGENT_DELETE` (value: `"agent.delete"`)

* `AGENT_UPDATE` (value: `"agent.update"`)

* `SKILL_CREATE` (value: `"skill.create"`)

* `SKILL_DELETE` (value: `"skill.delete"`)

* `SKILL_ENABLE` (value: `"skill.enable"`)

* `SKILL_DISABLE` (value: `"skill.disable"`)

* `SKILL_UPDATE` (value: `"skill.update"`)

* `SKILL_SHARE` (value: `"skill.share"`)

* `SKILL_UNSHARE` (value: `"skill.unshare"`)

* `SKILL_LOAD` (value: `"skill.load"`)

* `SKILL_FORCE_LOAD` (value: `"skill.force_load"`)

* `SKILL_VERSION_CREATE` (value: `"skill.version.create"`)

* `PLUGIN_CREATE` (value: `"plugin.create"`)

* `PLUGIN_UPDATE` (value: `"plugin.update"`)

* `PLUGIN_DELETE` (value: `"plugin.delete"`)

* `PROMPT_CREATE` (value: `"prompt.create"`)

* `PROMPT_DELETE` (value: `"prompt.delete"`)

* `PROMPT_UPDATE` (value: `"prompt.update"`)

* `PROMPT_VERSION_CREATE` (value: `"prompt.version.create"`)

* `KNOWLEDGE_BASE_CREATE` (value: `"knowledge_base.create"`)

* `KNOWLEDGE_BASE_DELETE` (value: `"knowledge_base.delete"`)

* `KNOWLEDGE_BASE_UPDATE` (value: `"knowledge_base.update"`)

* `KNOWLEDGE_BASE_VERSION_CREATE` (value: `"knowledge_base.version.create"`)

* `CUSTOM_VOICE_CREATE` (value: `"custom_voice.create"`)

* `CUSTOM_VOICE_UPDATE` (value: `"custom_voice.update"`)

* `CUSTOM_VOICE_DELETE` (value: `"custom_voice.delete"`)

* `FEATURE_PERMISSION_OVERRIDE_CREATED` (value: `"feature_permission.override.created"`)

* `FEATURE_PERMISSION_OVERRIDE_DELETED` (value: `"feature_permission.override.deleted"`)

* `RESOURCE_SHARE` (value: `"resource.share"`)

* `RESOURCE_UNSHARE` (value: `"resource.unshare"`)

* `LA_PLATEFORME_TRAINING_ENABLED` (value: `"la_plateforme.training.enabled"`)

* `LA_PLATEFORME_TRAINING_DISABLED` (value: `"la_plateforme.training.disabled"`)

* `CLOUD_RUNTIME_APP_CREATED` (value: `"cloud_runtime.app.created"`)

* `CLOUD_RUNTIME_APP_UPDATED` (value: `"cloud_runtime.app.updated"`)

* `CLOUD_RUNTIME_APP_DELETED` (value: `"cloud_runtime.app.deleted"`)

* `CLOUD_RUNTIME_APP_PAUSED` (value: `"cloud_runtime.app.paused"`)

* `CLOUD_RUNTIME_APP_RESUMED` (value: `"cloud_runtime.app.resumed"`)

* `CLOUD_RUNTIME_DEPLOYMENT_CREATED` (value: `"cloud_runtime.deployment.created"`)

* `CLOUD_RUNTIME_DEPLOYMENT_SUCCEEDED` (value: `"cloud_runtime.deployment.succeeded"`)

* `CLOUD_RUNTIME_DEPLOYMENT_CANCELED` (value: `"cloud_runtime.deployment.canceled"`)

* `CLOUD_RUNTIME_DEPLOYMENT_STOPPED` (value: `"cloud_runtime.deployment.stopped"`)

* `CLOUD_RUNTIME_DEPLOYMENT_FAILED` (value: `"cloud_runtime.deployment.failed"`)

* `CLOUD_RUNTIME_DEPLOYMENT_AUTOSCALED` (value: `"cloud_runtime.deployment.autoscaled"`)

* `CLOUD_RUNTIME_DOMAIN_CREATED` (value: `"cloud_runtime.domain.created"`)

* `CLOUD_RUNTIME_DOMAIN_UPDATED` (value: `"cloud_runtime.domain.updated"`)

* `CLOUD_RUNTIME_DOMAIN_DELETED` (value: `"cloud_runtime.domain.deleted"`)

* `CLOUD_RUNTIME_SECRET_CREATED` (value: `"cloud_runtime.secret.created"`)

* `CLOUD_RUNTIME_SECRET_UPDATED` (value: `"cloud_runtime.secret.updated"`)

* `CLOUD_RUNTIME_SECRET_DELETED` (value: `"cloud_runtime.secret.deleted"`)

* `CLOUD_RUNTIME_SERVICE_CREATED` (value: `"cloud_runtime.service.created"`)

* `CLOUD_RUNTIME_SERVICE_PAUSED` (value: `"cloud_runtime.service.paused"`)

* `CLOUD_RUNTIME_SERVICE_RESUMED` (value: `"cloud_runtime.service.resumed"`)

* `CLOUD_RUNTIME_SERVICE_DELETED` (value: `"cloud_runtime.service.deleted"`)

* `CLOUD_RUNTIME_SERVICE_MANUALLY_SCALED` (value: `"cloud_runtime.service.manually_scaled"`)

* `CLOUD_RUNTIME_SERVICE_MANUAL_SCALING_DELETED` (value: `"cloud_runtime.service.manual_scaling_deleted"`)

* `CLOUD_RUNTIME_PERSISTENT_VOLUME_CREATED` (value: `"cloud_runtime.persistent_volume.created"`)

* `CLOUD_RUNTIME_PERSISTENT_VOLUME_DELETED` (value: `"cloud_runtime.persistent_volume.deleted"`)

* `CLOUD_RUNTIME_PERSISTENT_VOLUME_ATTACHED` (value: `"cloud_runtime.persistent_volume.attached"`)

* `CLOUD_RUNTIME_PERSISTENT_VOLUME_DETACHED` (value: `"cloud_runtime.persistent_volume.detached"`)

* `FINE_TUNING_JOB_CREATE` (value: `"fine_tuning_job.create"`)

* `FINE_TUNING_JOB_CANCEL` (value: `"fine_tuning_job.cancel"`)

* `BATCH_JOB_CREATE` (value: `"batch_job.create"`)

* `BATCH_JOB_CANCEL` (value: `"batch_job.cancel"`)

* `BATCH_JOB_DELETE` (value: `"batch_job.delete"`)

* `DATA_CAPTURE_EXTRACT_JOB_CREATE` (value: `"data_capture.extract_job.create"`)

* `DATA_CAPTURE_EXTRACT_JOB_CANCEL` (value: `"data_capture.extract_job.cancel"`)

* `DATASET_CREATE` (value: `"dataset.create"`)

* `DATASET_DELETE` (value: `"dataset.delete"`)

* `LIBRARY_CREATE` (value: `"library.create"`)

* `LIBRARY_DELETE` (value: `"library.delete"`)

* `LIBRARY_UPDATE` (value: `"library.update"`)

* `LIBRARY_SHARE` (value: `"library.share"`)

* `LIBRARY_UNSHARE` (value: `"library.unshare"`)

* `LIBRARY_DOCUMENT_CREATE` (value: `"library.document.create"`)

* `LIBRARY_DOCUMENT_DELETE` (value: `"library.document.delete"`)

* `LIBRARY_DOCUMENT_BULK_DELETE` (value: `"library.document.bulk_delete"`)

* `LIBRARY_DOCUMENT_UPDATE` (value: `"library.document.update"`)

* `LIBRARY_DOCUMENT_REPROCESS` (value: `"library.document.reprocess"`)

* `INTEGRATION_CONNECTED` (value: `"integration.connected"`)

* `INTEGRATION_DISCONNECTED` (value: `"integration.disconnected"`)

* `INDEXING_WORKFLOW_COMPLETED` (value: `"indexing.workflow.completed"`)

* `INDEXING_DELETED` (value: `"indexing.deleted"`)

* `CONNECTION_ADMIN_SETUP_INDEX` (value: `"connection.admin.setup_index"`)

* `CONNECTION_ADMIN_DELETED` (value: `"connection.admin.deleted"`)

* `INTEGRATION_ACTIVATED_FOR_ORG` (value: `"integration.activated_for_org"`)

* `INTEGRATION_DEACTIVATED_FOR_ORG` (value: `"integration.deactivated_for_org"`)

* `INTEGRATION_ACTIVATED_FOR_WORKSPACE` (value: `"integration.activated_for_workspace"`)

* `INTEGRATION_DEACTIVATED_FOR_WORKSPACE` (value: `"integration.deactivated_for_workspace"`)

* `INTEGRATION_ACTIVATED_FOR_USER` (value: `"integration.activated_for_user"`)

* `INTEGRATION_DEACTIVATED_FOR_USER` (value: `"integration.deactivated_for_user"`)

* `INTEGRATION_CREATED` (value: `"integration.created"`)

* `INTEGRATION_UPDATED` (value: `"integration.updated"`)

* `INTEGRATION_DELETED` (value: `"integration.deleted"`)

* `INTEGRATION_TOOL_CALLED` (value: `"integration.tool_called"`)

* `INTEGRATION_TOOL_ACTIVATED` (value: `"integration.tool_activated"`)

* `INTEGRATION_TOOL_DEACTIVATED` (value: `"integration.tool_deactivated"`)

* `INTEGRATION_CREDENTIALS_CREATED_OR_UPDATED` (value: `"integration.credentials.created_or_updated"`)

* `INTEGRATION_CREDENTIALS_DELETED` (value: `"integration.credentials.deleted"`)

* `INTEGRATION_CREDENTIALS_REVOKED` (value: `"integration.credentials.revoked"`)

* `INTEGRATION_CREDENTIALS_REVOCATION_FAILED` (value: `"integration.credentials.revocation_failed"`)

* `INTEGRATION_PREFERENCES_CREATED_OR_UPDATED` (value: `"integration.preferences.created_or_updated"`)

* `INTEGRATION_PREFERENCES_DELETED` (value: `"integration.preferences.deleted"`)

* `INTEGRATION_AUTHENTICATION_METHOD_CREATED_OR_UPDATED` (value: `"integration.authentication_method.created_or_updated"`)

* `INTEGRATION_CONNECTION_CREATED` (value: `"integration.connection.created"`)

* `INTEGRATION_SHARED` (value: `"integration.shared"`)

* `INTEGRATION_UNSHARED` (value: `"integration.unshared"`)

* `CONNECTORS_GATEWAY_TOOL_CALLED` (value: `"connectors_gateway.tool_called"`)

* `CONNECTORS_DEBUGGER_TOOL_CALLED` (value: `"connectors_debugger.tool_called"`)

* `CRAWLER_CONFIG_CREATE` (value: `"crawler.config.create"`)

* `CRAWLER_CONFIG_UPDATE` (value: `"crawler.config.update"`)

* `CRAWLER_CONFIG_DELETE` (value: `"crawler.config.delete"`)

* `CRAWLER_RUN_CREATE` (value: `"crawler.run.create"`)

* `CRAWLER_RUN_CANCEL` (value: `"crawler.run.cancel"`)

* `WEB_SEARCH_FILTER_CONFIG_UPDATE` (value: `"web_search.filter_config.update"`)

* `WEB_SEARCH_FILTER_DOMAIN_ADD` (value: `"web_search.filter_domain.add"`)

* `WEB_SEARCH_FILTER_DOMAIN_REMOVE` (value: `"web_search.filter_domain.remove"`)

* `WEB_FILTERS_CONFIG_UPDATE` (value: `"web_filters.config.update"`)

* `WEB_FILTERS_DOMAIN_ADD` (value: `"web_filters.domain.add"`)

* `WEB_FILTERS_DOMAIN_REMOVE` (value: `"web_filters.domain.remove"`)

* `SECRETS_MANAGER_SECRET_CREATED` (value: `"secrets_manager.secret.created"`)

* `SECRETS_MANAGER_SECRET_UPDATED` (value: `"secrets_manager.secret.updated"`)

* `SECRETS_MANAGER_SECRET_REVEALED` (value: `"secrets_manager.secret.revealed"`)

* `SECRETS_MANAGER_SECRET_METADATA_UPDATED` (value: `"secrets_manager.secret.metadata_updated"`)

* `SECRETS_MANAGER_SECRET_DELETED` (value: `"secrets_manager.secret.deleted"`)

* `SECRETS_MANAGER_SECRET_VERSION_DELETED` (value: `"secrets_manager.secret.version.deleted"`)

* `SECRETS_MANAGER_SECRET_VERSION_RESTORED` (value: `"secrets_manager.secret.version.restored"`)

* `RATE_LIMIT_RULE_CREATE` (value: `"rate_limit.rule.create"`)

* `RATE_LIMIT_RULE_UPDATE` (value: `"rate_limit.rule.update"`)

* `RATE_LIMIT_RULE_DELETE` (value: `"rate_limit.rule.delete"`)

* `WORKFLOW_EXECUTION_STARTED` (value: `"workflow.execution.started"`)

* `WORKFLOW_EXECUTION_CANCELED` (value: `"workflow.execution.canceled"`)

* `WORKFLOW_EXECUTION_TERMINATED` (value: `"workflow.execution.terminated"`)

* `WORKFLOW_EXECUTION_SIGNAL` (value: `"workflow.execution.signal"`)

* `WORKFLOW_EXECUTION_QUERY` (value: `"workflow.execution.query"`)

* `WORKFLOW_EXECUTION_UPDATE` (value: `"workflow.execution.update"`)

* `WORKFLOW_WORKER_REGISTERED` (value: `"workflow.worker.registered"`)

* `WORKFLOW_SCHEDULE_CREATED` (value: `"workflow.schedule.created"`)

* `WORKFLOW_SCHEDULE_UPDATED` (value: `"workflow.schedule.updated"`)

* `WORKFLOW_SCHEDULE_PAUSED` (value: `"workflow.schedule.paused"`)

* `WORKFLOW_SCHEDULE_RESUMED` (value: `"workflow.schedule.resumed"`)

* `WORKFLOW_SCHEDULE_DELETED` (value: `"workflow.schedule.deleted"`)

* `WORKFLOW_SCHEDULE_TRIGGERED` (value: `"workflow.schedule.triggered"`)

* `WORKFLOW_ARCHIVED` (value: `"workflow.archived"`)

* `WORKFLOW_UNARCHIVED` (value: `"workflow.unarchived"`)

* `WORKFLOW_DEPLOYMENT_AUTHORIZED_CREDENTIAL_GRANTED` (value: `"workflow.deployment.authorized_credential.granted"`)

* `WORKFLOW_DEPLOYMENT_AUTHORIZED_CREDENTIAL_REVOKED` (value: `"workflow.deployment.authorized_credential.revoked"`)


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


