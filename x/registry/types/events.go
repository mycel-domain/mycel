package types

// Register domain event
const (
	EventTypeRegsterDomain = "register-domain"

	AttributeRegisterDomainEventName           = "name"
	AttributeRegisterDomainEventParent         = "parent"
	AttributeRegisterDomainEventExpirationDate = "expiration-date"
)

// Register top-level-domain event
const (
	EventTypeRegsterTopLevelDomain = "register-top-level-domain"

	AttributeRegisterTopLevelDomainEventName           = "name"
	AttributeRegisterTopLevelDomainEventExpirationDate = "expiration-date"
)

// Update wallet record event
const (
	EventTypeUpdateWalletRecord = "update-wallet-record"

	AttributeUpdateWalletRecordEventDomainName       = "name"
	AttributeUpdateWalletRecordEventDomainParent     = "parent"
	AttributeUpdateWalletRecordEventWalletRecordType = "wallet-record-type"
	AttributeUpdateWalletRecordEventValue            = "value"
)

// Update dns record event
const (
	EventTypeUpdateDnsRecord = "update-dns-record"

	AttributeUpdateDnsRecordEventDomainName    = "name"
	AttributeUpdateDnsRecordEventDomainParent  = "parent"
	AttributeUpdateDnsRecordEventDnsRecordType = "dns-record-type"
	AttributeUpdateDnsRecordEventValue         = "value"
)

// Set registration fees event
const (
	EventTypeSetRegistrationFees = "set-registration-fees"

	AttributeSetRegistrationFeesDomain = "domain"
)
