package vail

type Response struct {
	Data struct {
		Name               string `json:"name"`
		ShowOpenPercentage bool   `json:"showOpenPercentage"`
		SlopeLevelStats    []struct {
			Level string `json:"level"`
			Open  int64  `json:"open"`
			Total int64  `json:"total"`
		} `json:"slopeLevelStats"`
		SlopeType string `json:"slopeType"`
		Stations  []struct {
			Dateinfo string `json:"dateinfo"`
			Info     struct {
				Categories []struct {
					Active bool   `json:"active"`
					Name   string `json:"name"`
					Order  int64  `json:"order"`
				} `json:"categories"`
				MessageDiffusion struct {
					Active    bool   `json:"active"`
					End       int64  `json:"end"`
					FontColor string `json:"fontColor"`
					Start     int64  `json:"start"`
				} `json:"messageDiffusion"`
			} `json:"info"`
			Name            string `json:"name"`
			SlopeLevelStats []struct {
				Level string `json:"level"`
				Open  int64  `json:"open"`
				Total int64  `json:"total"`
			} `json:"slopeLevelStats"`
			SlopeMapID     int64  `json:"slopeMapId"`
			SlopeType      string `json:"slopeType"`
			SlopeTypeCount int64  `json:"slopeTypeCount"`
			States         []struct {
				Exposition      string        `json:"exposition"`
				High            string        `json:"high"`
				Links           []interface{} `json:"links"`
				Low             string        `json:"low"`
				Name            string        `json:"name"`
				SectorID        string        `json:"sectorId"`
				Skilifts        []Lift        `json:"skilifts"`
				SlopeLevelStats []struct {
					Level string `json:"level"`
					Open  int64  `json:"open"`
					Total int64  `json:"total"`
				} `json:"slopeLevelStats"`
				SlopeType      string `json:"slopeType"`
				SlopeTypeCount int64  `json:"slopeTypeCount"`
				Slopes         []struct {
					Elevation        string  `json:"elevation"`
					EndAltitude      string  `json:"endAltitude"`
					EndHour          string  `json:"endHour"`
					Exposition       string  `json:"exposition"`
					Length           string  `json:"length"`
					Level            string  `json:"level"`
					LiftGroupID      string  `json:"liftGroupId"`
					LiftIDInGroup    string  `json:"liftIdInGroup"`
					Maintenance      string  `json:"maintenance"`
					MaintenanceID    string  `json:"maintenanceId"`
					Message          *string `json:"message,omitempty"`
					Name             string  `json:"name"`
					OffbeatEndHour   string  `json:"offbeatEndHour"`
					OffbeatStartHour string  `json:"offbeatStartHour"`
					SectorID         string  `json:"sectorId"`
					Snowquality      string  `json:"snowquality"`
					StartAltitude    string  `json:"startAltitude"`
					StartHour        string  `json:"startHour"`
					State            string  `json:"state"`
					StateDate        string  `json:"stateDate"`
					StateType        string  `json:"stateType"`
					StationID        string  `json:"stationId"`
					Surface          string  `json:"surface"`
					Type             string  `json:"type"`
					UpdatedDate      string  `json:"updatedDate"`
					UpdatedEndHour   string  `json:"updatedEndHour"`
					UpdatedStartHour string  `json:"updatedStartHour"`
					UUID             int64   `json:"uuid"`
				} `json:"slopes"`
				StationID string `json:"stationId"`
				Stats     []struct {
					Open  int64  `json:"open"`
					Total int64  `json:"total"`
					Type  string `json:"type"`
				} `json:"stats"`
				UUID int64 `json:"uuid"`
			} `json:"states"`
			Stats []struct {
				Open  int64  `json:"open"`
				Total int64  `json:"total"`
				Type  string `json:"type"`
			} `json:"stats"`
		} `json:"stations"`
		Stats []struct {
			Open  int64  `json:"open"`
			Total int64  `json:"total"`
			Type  string `json:"type"`
		} `json:"stats"`
	} `json:"data"`
	LastUpdate string `json:"lastUpdate"`
	Status     int64  `json:"status"`
}

type Lift struct {
	Capacity         string  `json:"capacity"`
	Debit            string  `json:"debit"`
	Duration         string  `json:"duration"`
	Elevation        string  `json:"elevation"`
	EndAltitude      string  `json:"endAltitude"`
	EndHour          string  `json:"endHour"`
	Exposition       string  `json:"exposition"`
	Length           string  `json:"length"`
	LiftGroupID      string  `json:"liftGroupId"`
	LiftIDInGroup    string  `json:"liftIdInGroup"`
	Message          *string `json:"message,omitempty"`
	Name             string  `json:"name"`
	OffbeatEndHour   string  `json:"offbeatEndHour"`
	OffbeatStartHour string  `json:"offbeatStartHour"`
	SectorID         string  `json:"sectorId"`
	Speed            string  `json:"speed"`
	StartAltitude    string  `json:"startAltitude"`
	StartHour        string  `json:"startHour"`
	State            string  `json:"state"`
	StateDate        string  `json:"stateDate"`
	StateType        string  `json:"stateType"`
	StationID        string  `json:"stationId"`
	Type             string  `json:"type"`
	UpdatedDate      string  `json:"updatedDate"`
	UpdatedEndHour   string  `json:"updatedEndHour"`
	UpdatedStartHour string  `json:"updatedStartHour"`
	UUID             *int64  `json:"uuid,omitempty"`
	Wait             string  `json:"wait"`
}
