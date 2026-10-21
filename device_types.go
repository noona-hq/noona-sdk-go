package hq

const (
	DeviceTypeCheckin DeviceType = "checkin"

	DeviceStatusUnpaired DeviceStatus = "unpaired"
	DeviceStatusPairing  DeviceStatus = "pairing"
	DeviceStatusPaired   DeviceStatus = "paired"
	DeviceStatusRevoked  DeviceStatus = "revoked"

	DeviceFieldName                                    DeviceField = "name"
	DeviceFieldConfigurationTerminal                   DeviceField = "configuration.terminal"
	DeviceFieldConfigurationSuccessMessage             DeviceField = "configuration.success_message"
	DeviceFieldConfigurationSuccessMessageTranslations DeviceField = "configuration.success_message_translations"
	DeviceFieldConfigurationLanguages                  DeviceField = "configuration.languages"
	DeviceFieldConfigurationDefaultLanguage            DeviceField = "configuration.default_language"
	DeviceFieldConfigurationBrandColor                 DeviceField = "configuration.brand_color"
)
