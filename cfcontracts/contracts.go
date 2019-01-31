package cfcontracts

// IndividualMetaData is metadata for individuals.
type IndividualMetaData struct {
	DisplayName string `json:"displayName"` // The display name of the individual
	URL         string `json:"url"`         // The url to use to get more information about this individual
}

// OrganizationMetadata is metadata for organizations.
type OrganizationMetadata struct {
	DisplayName string `json:"displayName"` // The display name of the organization.
	URL         string `json:"url"`         // The url to use to get more information about the organization.
}

// Individual represents an individual
type Individual struct {
	DisplayName  string `json:"displayName"`  // The display name of the individual
	EmailAddress string `json:"emailAddress"` // The email address for the individual
	FirstName    string `json:"firstName"`    // The first name of the individual
	LastName     string `json:"lastName"`     // The last name of the individual
	Age          int    `json:"age"`          // The age of the individual
	Gender       string `json:"gender"`       // The gender of the individual
	Information  string `json:"information"`  // Information about the individual.
	Avatar       string `json:"avatar"`       // A display image for the individual.
}
