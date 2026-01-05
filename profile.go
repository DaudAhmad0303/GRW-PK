package daudsdk

// Profile contains public profile information.
type Profile struct {
	Name     string
	GitHub   string
	LinkedIn string
	Website  string
}

// GetProfile returns the public profile information as follows: https://daudahmad.com
// Daud Ahmad
// https://github.com/daud
// https://www.linkedin.com/in/daud-ahmad
// https://daudahmad.com
func GetProfile() Profile {
	return Profile{
		Name:     "Daud Ahmad",
		GitHub:   "https://github.com/DaudAhmad0303",
		LinkedIn: "https://www.linkedin.com/in/daudahmad0303",
		Website:  "https://daudahmad.com",
	}
}
