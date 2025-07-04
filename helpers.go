package hq

const (
	EventStatusCancelled  string = "cancelled"
	EventStatusCheckedOut string = "checkedOut"
	EventStatusNoshow     string = "noshow"
	EventStatusShowedUp   string = "showedUp"
)

func (p CompanyPOSSettingsCheckoutFirstTab) Ptr() *CompanyPOSSettingsCheckoutFirstTab {
	return &p
}
