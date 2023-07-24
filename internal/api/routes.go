package api

// Route is the model for the router setup
type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc apiFunc
	RequireAuth bool
}

// Routes are the main setup for our Router
type Routes []Route

func (s *APIServer) RoutesHandler() Routes {
	var routes = Routes{
		Route{"HealthCheck", "GET", "/health", s.HealthCheck, false},
		Route{"Mailbox", "POST", "/mailbox", s.Mailbox, true},
		// Route{"InsertLeaderboards", "POST", "/leaderboard", InsertPlayer},
		// Route{"GetLeaderboards", "GET", "/leaderboard/{which}", GetLeaderboards},
		// Route{"LeaderboardLookup", "POST", "/lb-lookup", LeaderboardLookup},

		// Route{"Auction", "POST", "/auction", Auctions},

		// Route{"PetsExistance", "POST", "/pets-exist", PetsExistance},
	}

	return routes
}
