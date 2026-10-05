package hq

const (
	EventStatusCancelled  string = "cancelled"
	EventStatusCheckedOut string = "checkedOut"
	EventStatusNoshow     string = "noshow"
	EventStatusShowedUp   string = "showedUp"

	// Deprecated: use CustomerGroupSystemTypeBlacklist.
	Blacklist CustomerGroupSystemType = CustomerGroupSystemTypeBlacklist

	// Keep the established names when generated enum naming changes as schemas evolve.
	VoucherTemplateUpdateTypeAmount  VoucherTemplateUpdateType = Amount
	VoucherTemplateUpdateTypeService VoucherTemplateUpdateType = Service
)

func (p CompanyPOSSettingsCheckoutFirstTab) Ptr() *CompanyPOSSettingsCheckoutFirstTab {
	return &p
}
