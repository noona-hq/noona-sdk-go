package hq

const (
	DeviceTypeCheckin DeviceType = "checkin"

	DeviceStatusUnpaired DeviceStatus = "unpaired"
	DeviceStatusPairing  DeviceStatus = "pairing"
	DeviceStatusPaired   DeviceStatus = "paired"
	DeviceStatusRevoked  DeviceStatus = "revoked"

	DeviceFieldName                        DeviceField = "name"
	DeviceFieldConfigurationTerminal       DeviceField = "configuration.terminal"
	DeviceFieldConfigurationSuccessMessage DeviceField = "configuration.success_message"
)
