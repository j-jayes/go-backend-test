package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	g "github.com/serpapi/google-search-results-golang"
)

type SearchParams struct {
	APIKey       string
	Engine       string
	Location     string
	GoogleDomain string
	GL           string
	HL           string
}

// SearchResult represents the search result returned by the SERP API.
type SearchResult map[string]interface{}

// PerformSearch carries out a search using the SERP API based on given parameters
func PerformSearch(query string) (map[string]interface{}, error) {
	params := SearchParams{
		// Add your configurations here:
		APIKey:       "e1e7050eb9085d438c9f322b2348b38a06889b20ab667106653dbc84a8135a02",
		Engine:       "google",
		Location:     "Greater London, England, United Kingdom",
		GoogleDomain: "google.co.uk",
		GL:           "uk",
		HL:           "en",
	}

	parameter := map[string]string{
		"api_key":       params.APIKey,
		"engine":        params.Engine,
		"q":             query, // set the query
		"location":      params.Location,
		"google_domain": params.GoogleDomain,
		"gl":            params.GL,
		"hl":            params.HL,
	}

	search := g.NewGoogleSearch(parameter, params.APIKey)
	return search.GetJSON()
}

// @Summary Perform a SERP search
// @Description Carries out a search using the SERP API based on the given parameters
// @Accept  json
// @Produce  json
// @Param query query string true "Search Query"
// @Success 200 {object} SearchResult
// @Failure 500 {object} map[string]string{"error":"error message"}
// @Router /search/{query} [get]
func SearchResults(c *gin.Context) {
	query := c.Param("query")

	results, err := PerformSearch(query)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Here you can format results with the ShowSearchResults function or directly send them
	c.JSON(http.StatusOK, results)
}

// ShowSearchResults formats and displays the search results.
/* func ShowSearchResults(organicResults []g.OrganicResult) {
	for _, result := range organicResults {
		fmt.Println("Title:", result.Title)
		fmt.Println("Link:", result.Link)
		fmt.Println("Snippet:", result.Snippet)
		fmt.Println("Source:", result.Source)
		fmt.Println("------------------------------------------------------")
	}
} */
