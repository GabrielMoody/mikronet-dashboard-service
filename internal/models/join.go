package models

type BlockDriver struct {
	ID             string `json:"id"`
	Email          string `json:"email"`
	Name           string `json:"name"`
	ProfilePicture string `json:"profile_picture"`
}

type Drivers struct {
	ID             string `json:"id"`
	Email          string `json:"email"`
	Name           string `json:"name"`
	PhoneNumber    string `json:"phone_number"`
	LicenseNumber  string `json:"license_number"`
	SIM            string `json:"sim"`
	Verified       bool   `json:"verified"`
	ProfilePicture string `json:"profile_picture"`
	Status         string `json:"status"`
}

type Passengers struct {
	ID    string `json:"id"`
	Email string `json:"email"`
	Name  string `json:"name"`
}

type Reviews struct {
	ID            int    `json:"id"`
	PassengerName string `json:"passenger_name"`
	DriverName    string `json:"driver_name"`
	Comment       string `json:"comment"`
	Star          int    `json:"star"`
}
