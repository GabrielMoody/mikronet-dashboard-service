package dto

type (
	GetDriverQuery struct {
		Verified *bool `query:"verified"`
	}

	AddRoute struct {
		RouteName string `json:"route_name"`
	}
)
