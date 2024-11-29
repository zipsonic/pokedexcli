package api

type LocationAreaResponse struct {
	Count    int
	Next     *string
	Previous *string
	Results  []struct {
		Name string
		Url  string
	}
}
