package main

import (
	"fmt"
	"log"
	"strconv"
)

type OptionInt struct {
	set bool
	val int
}

const dashes = "\"--\""

func (o *OptionInt) UnmarshalJSON(b []byte) error {
	bs := string(b)

	if bs == dashes {
		*o = OptionInt{
			set: false,
			val: 0,
		}

		return nil
	}

	ans, err := strconv.Atoi(bs)
	if err != nil {
		return err
	}
	*o = OptionInt{
		set: true,
		val: ans,
	}

	return nil
}

func (o OptionInt) MarshalJSON() ([]byte, error) {
	if !o.set {
		return []byte(dashes), nil
	}

	return []byte(fmt.Sprintf("%d", o.val)), nil
}

type AreaSnow struct {
	BaseCm             string `json:"BaseCm"`
	BaseIn             string `json:"BaseIn"`
	Last24HoursCm      string `json:"Last24HoursCm"`
	Last24HoursIn      string `json:"Last24HoursIn"`
	Last48HoursCm      string `json:"Last48HoursCm"`
	Last48HoursIn      string `json:"Last48HoursIn"`
	Last72HoursCm      string `json:"Last72HoursCm"`
	Last72HoursIn      string `json:"Last72HoursIn"`
	Last7DaysCm        string `json:"Last7DaysCm"`
	Last7DaysIn        string `json:"Last7DaysIn"`
	SinceLiftsClosedCm string `json:"SinceLiftsClosedCm"`
	SinceLiftsClosedIn string `json:"SinceLiftsClosedIn"`
}

type ClosedOpen struct {
	Close string `json:"Close"`
	Open  string `json:"Open"`
}

type Condition struct {
	Conditions       string `json:"Conditions"`
	Default          string `json:"Default"`
	DewPointC        string `json:"DewPointC"`
	DewPointF        string `json:"DewPointF"`
	FeedSavedTime    string `json:"FeedSavedTime"`
	Humidity         int64  `json:"Humidity"`
	Icon             string `json:"Icon"`
	PressureIN       string `json:"PressureIN"`
	PressureMB       string `json:"PressureMB"`
	Skies            string `json:"Skies"`
	TemperatureC     string `json:"TemperatureC"`
	TemperatureF     string `json:"TemperatureF"`
	TemperatureHighC string `json:"TemperatureHighC"`
	TemperatureHighF string `json:"TemperatureHighF"`
	TemperatureLowC  string `json:"TemperatureLowC"`
	TemperatureLowF  string `json:"TemperatureLowF"`
	UvIndex          string `json:"UvIndex"`
	WindChillC       string `json:"WindChillC"`
	WindChillF       string `json:"WindChillF"`
	WindDirection    string `json:"WindDirection"`
	WindGustsKph     int64  `json:"WindGustsKph"`
	WindGustsMph     int64  `json:"WindGustsMph"`
	WindStrengthKph  int64  `json:"WindStrengthKph"`
	WindStrengthMph  int64  `json:"WindStrengthMph"`
}

type Conditions struct {
	Base        Condition `json:"Base"`
	MidMountain Condition `json:"MidMountain"`
	Summit      Condition `json:"Summit"`
}

type AveWind struct {
	Dir string `json:"dir"`
	Kph string `json:"kph"`
	Mph string `json:"mph"`
}

type Forecast struct {
	Avewind               *AveWind `json:"avewind,omitempty"`
	Conditions            *string  `json:"conditions,omitempty"`
	Date                  string   `json:"date"`
	ForecastedSnowCm      int64    `json:"forecasted_snow_cm"`
	ForecastedSnowDayCm   *string  `json:"forecasted_snow_day_cm,omitempty"`
	ForecastedSnowDayIn   *string  `json:"forecasted_snow_day_in,omitempty"`
	ForecastedSnowIn      int64    `json:"forecasted_snow_in"`
	ForecastedSnowNightCm string   `json:"forecasted_snow_night_cm"`
	ForecastedSnowNightIn string   `json:"forecasted_snow_night_in"`
	Icon                  *string  `json:"icon,omitempty"`
	Skies                 *string  `json:"skies,omitempty"`
	TempHighC             *string  `json:"temp_high_c,omitempty"`
	TempHighF             *string  `json:"temp_high_f,omitempty"`
	TempLowC              string   `json:"temp_low_c"`
	TempLowF              string   `json:"temp_low_f"`
}

type Forecasts struct {
	FiveDay          Forecast `json:"FiveDay"`
	ForecastedSnowCm string   `json:"ForecastedSnowCm"`
	ForecastedSnowIn string   `json:"ForecastedSnowIn"`
	FourDay          Forecast `json:"FourDay"`
	OneDay           Forecast `json:"OneDay"`
	TempHighC        string   `json:"TempHighC"`
	TempHighF        string   `json:"TempHighF"`
	TempLowC         string   `json:"TempLowC"`
	TempLowF         string   `json:"TempLowF"`
	ThreeDay         Forecast `json:"ThreeDay"`
	TwoDay           Forecast `json:"TwoDay"`
}

type Hours struct {
	Friday    ClosedOpen `json:"Friday"`
	Monday    ClosedOpen `json:"Monday"`
	Saturday  ClosedOpen `json:"Saturday"`
	Sunday    ClosedOpen `json:"Sunday"`
	Thursday  ClosedOpen `json:"Thursday"`
	Tuesday   ClosedOpen `json:"Tuesday"`
	Wednesday ClosedOpen `json:"Wednesday"`
}

type LiftStatus int

func (s *LiftStatus) UnmarshalJSON(b []byte) error {
	for i := LiftStatus(0); i < LiftStatusMax; i++ {
		bs, err := i.MarshalJSON()
		if err != nil {
			return err
		}

		if string(bs) == string(b) {
			*s = i
			return nil
		}
	}

	return fmt.Errorf("can't unmarshal %s", string(b))
}

func (s LiftStatus) MarshalJSON() ([]byte, error) {
	return []byte(strconv.Quote(s.String())), nil
}

const (
	LiftStatusScheduled LiftStatus = iota
	LiftStatusClosed
	LiftStatusWindHold
	LiftStatusOpen
	LiftStatusMechanicalHold
	LiftStatusAnticipatedWeatherImpact
	LiftStatusMax
)

func (s LiftStatus) String() string {
	switch s {
	case LiftStatusScheduled:
		return "Scheduled"
	case LiftStatusClosed:
		return "Closed"
	case LiftStatusWindHold:
		return "Wind Hold"
	case LiftStatusOpen:
		return "Open"
	case LiftStatusMechanicalHold:
		return "Mechanical Hold"
	case LiftStatusAnticipatedWeatherImpact:
		return "Anticipated Weather Impact"
	}

	log.Fatalf("how string %d", s)
	return ""
}

func (s LiftStatus) OneHot() map[string]interface{} {
	ans := map[string]interface{}{}
	for i := LiftStatus(0); i < LiftStatusMax; i++ {
		ans[i.String()] = false
		if i == s {
			ans[i.String()] = true
		}
	}
	return ans
}

type Lift struct {
	FirstTracks      string     `json:"FirstTracks"`
	Hours            Hours      `json:"Hours"`
	LiftType         string     `json:"LiftType"`
	MountainAreaName string     `json:"MountainAreaName"`
	Name             string     `json:"Name"`
	Status           LiftStatus `json:"Status"`
	StatusEnglish    string     `json:"StatusEnglish"`
	UpdateDate       string     `json:"UpdateDate"`
	WaitTime         OptionInt  `json:"WaitTime"`
	WaitTimeStatus   OptionInt  `json:"WaitTimeStatus"`
	WaitTimeString   string     `json:"WaitTimeString"`
}

type Trail struct {
	Difficulty       string `json:"Difficulty"`
	Glades           string `json:"Glades"`
	Grooming         string `json:"Grooming"`
	Moguls           string `json:"Moguls"`
	MountainAreaName string `json:"MountainAreaName"`
	Name             string `json:"Name"`
	NightSkiing      string `json:"NightSkiing"`
	Nordic           string `json:"Nordic"`
	SnowMaking       string `json:"SnowMaking"`
	Status           string `json:"Status"`
	StatusEnglish    string `json:"StatusEnglish"`
	Touring          string `json:"Touring"`
	TrailIcon        string `json:"TrailIcon"`
	UpdateDate       string `json:"UpdateDate"`
}

type Activity struct {
	Hours         Hours  `json:"Hours"`
	LinkURL       string `json:"LinkUrl"`
	Name          string `json:"Name"`
	Status        string `json:"Status"`
	StatusEnglish string `json:"StatusEnglish"`
	UpdateDate    string `json:"UpdateDate"`
}

type MountainArea struct {
	Activities       []Activity `json:"Activities"`
	LastUpdate       string     `json:"LastUpdate"`
	Lifts            []Lift     `json:"Lifts"`
	Name             string     `json:"Name"`
	OpenTrailsCount  int64      `json:"OpenTrailsCount"`
	TotalTrailsCount int64      `json:"TotalTrailsCount"`
	Trails           []Trail    `json:"Trails"`
}

type SnowReport struct {
	AdditionalText          string   `json:"AdditionalText"`
	Alert                   string   `json:"Alert"`
	AnnualAverageSnowfallCm string   `json:"AnnualAverageSnowfallCm"`
	AnnualAverageSnowfallIn string   `json:"AnnualAverageSnowfallIn"`
	BaseArea                AreaSnow `json:"BaseArea"`
	BaseConditions          string   `json:"BaseConditions"`
	GroomedTrails           int64    `json:"GroomedTrails"`
	LastUpdate              string   `json:"LastUpdate"`
	LastUpdatedLift         map[string]string
	LiftNotification        string   `json:"LiftNotification"`
	MidMountainArea         AreaSnow `json:"MidMountainArea"`
	News                    string   `json:"News"`
	OpenNightParks          int64    `json:"OpenNightParks"`
	OpenNightTrails         int64    `json:"OpenNightTrails"`
	OpenTerrainAcres        string   `json:"OpenTerrainAcres"`
	OpenTerrainHectares     string   `json:"OpenTerrainHectares"`
	Report                  string   `json:"Report"`
	SafetyReport            string   `json:"SafetyReport"`
	SeasonTotalCm           string   `json:"SeasonTotalCm"`
	SeasonTotalIn           string   `json:"SeasonTotalIn"`
	SecondarySeasonTotalCm  string   `json:"SecondarySeasonTotalCm"`
	SecondarySeasonTotalIn  string   `json:"SecondarySeasonTotalIn"`
	SnowBaseRangeCM         string   `json:"SnowBaseRangeCM"`
	SnowBaseRangeIn         string   `json:"SnowBaseRangeIn"`
	StormRadar              string   `json:"StormRadar"`
	StormRadarButtonText    string   `json:"StormRadarButtonText"`
	StormTotalCM            string   `json:"StormTotalCM"`
	StormTotalIn            string   `json:"StormTotalIn"`
	SummitArea              AreaSnow `json:"SummitArea"`
	TotalActivities         int64    `json:"TotalActivities"`
	TotalLifts              int64    `json:"TotalLifts"`
	TotalNightParks         int64    `json:"TotalNightParks"`
	TotalNightTrails        int64    `json:"TotalNightTrails"`
	TotalOpenActivities     int64    `json:"TotalOpenActivities"`
	TotalOpenLifts          int64    `json:"TotalOpenLifts"`
	TotalOpenParks          int64    `json:"TotalOpenParks"`
	TotalOpenTrails         int64    `json:"TotalOpenTrails"`
	TotalParks              int64    `json:"TotalParks"`
	TotalTerrainAcres       string   `json:"TotalTerrainAcres"`
	TotalTerrainHectares    string   `json:"TotalTerrainHectares"`
	TotalTrails             int64    `json:"TotalTrails"`
}

type Response struct {
	CurrentConditions Conditions `json:"CurrentConditions"`
	Forecast          Forecasts  `json:"Forecast"`
	LastUpdate        string     `json:"LastUpdate"`
	LayoutOptions     struct {
		Disclaimer       string        `json:"Disclaimer"`
		PrimaryWeather   string        `json:"PrimaryWeather"`
		SecondaryWeather []interface{} `json:"SecondaryWeather"`
		Weather          []interface{} `json:"Weather"`
		SoldOut          string        `json:"soldOut"`
	} `json:"LayoutOptions"`
	MountainAreas   []MountainArea `json:"MountainAreas"`
	Name            string         `json:"Name"`
	OperatingStatus string         `json:"OperatingStatus"`
	ParkingLots     []interface{}  `json:"ParkingLots"`
	RoadOptions     []interface{}  `json:"RoadOptions"`
	SnowReport      SnowReport     `json:"SnowReport"`
}
