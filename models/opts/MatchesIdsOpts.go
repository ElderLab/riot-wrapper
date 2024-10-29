package opts

// MatchesIdsOpts is a struct that contains the options for the GetMatchesIdsWithOpts function.
type MatchesIdsOpts struct {
	StartTime int `query:"startTime"`
	EndTime   int `query:"endTime"`
	Queue     int `query:"queue"`
	Type      int `query:"type"`
	Start     int `query:"start"`
	Count     int `query:"count"`
}
