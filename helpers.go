package hq

const (
	EventStatusCancelled  string = "cancelled"
	EventStatusCheckedOut string = "checkedOut"
	EventStatusNoshow     string = "noshow"
	EventStatusShowedUp   string = "showedUp"

	// Deprecated: use CustomerGroupSystemTypeBlacklist.
	Blacklist CustomerGroupSystemType = CustomerGroupSystemTypeBlacklist
)

func (p CompanyPOSSettingsCheckoutFirstTab) Ptr() *CompanyPOSSettingsCheckoutFirstTab {
	return &p
}
