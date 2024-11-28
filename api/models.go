package api

type LocationAreaResponse struct {
	Count        int
	NextPage     *string
	PreviousPage *string
	Results      []struct {
		Name string
		Url  string
	}
}
