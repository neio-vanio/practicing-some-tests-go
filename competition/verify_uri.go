package competition

// CheckerURI check uri
type CheckerURI func(string) bool

// Result for chan
type Result struct {
	string
	bool
}

// VerifyURI check uri
func VerifyURI(c CheckerURI, uris []string) map[string]bool {
	results := make(map[string]bool)
	chanResult := make(chan Result)

	for _, uri := range uris {
		go func(u string) {
			chanResult <- Result{u, c(u)}
		}(uri)
	}

	for i := 0; i < len(uris); i++ {
		result := <-chanResult
		results[result.string] = result.bool
	}

	return results
}
