package domain

import (
	"encoding/csv"
	"fmt"
	"strconv"
	"strings"
	"time"

	health_data "github.com/hikata101/health_data/v1"
	pb "github.com/hikata101/health_data_service/gen/github.com/hikata101/health_data_service/v1"
)

var dictionary_countries = map[health_data.Country]string{
	health_data.Country_COUNTRY_ALBANIA:                                         "ALB",
	health_data.Country_COUNTRY_ANDORRA:                                         "AND",
	health_data.Country_COUNTRY_ARMENIA:                                         "ARM",
	health_data.Country_COUNTRY_AUSTRIA:                                         "AUT",
	health_data.Country_COUNTRY_AZERBAIJAN:                                      "AZE",
	health_data.Country_COUNTRY_BELARUS:                                         "BLR",
	health_data.Country_COUNTRY_BELGIUM:                                         "BEL",
	health_data.Country_COUNTRY_BELGIUM_BRUXELLES:                               "BE-BRU",
	health_data.Country_COUNTRY_BELGIUM_FLANDERS:                                "BE-VLG",
	health_data.Country_COUNTRY_BELGIUM_WALLONIA:                                "BE-WAL",
	health_data.Country_COUNTRY_BOSNIA_AND_HERZEGOVINA:                          "BIH",
	health_data.Country_COUNTRY_BULGARIA:                                        "BGR",
	health_data.Country_COUNTRY_CANADA:                                          "CAN",
	health_data.Country_COUNTRY_CROATIA:                                         "HRV",
	health_data.Country_COUNTRY_CYPRUS:                                          "CYP",
	health_data.Country_COUNTRY_CZECH_REPUBLIC:                                  "CZE",
	health_data.Country_COUNTRY_DENMARK:                                         "DNK",
	health_data.Country_COUNTRY_ESTONIA:                                         "EST",
	health_data.Country_COUNTRY_FINLAND:                                         "FIN",
	health_data.Country_COUNTRY_FRANCE:                                          "FRA",
	health_data.Country_COUNTRY_GEORGIA:                                         "GEO",
	health_data.Country_COUNTRY_GERMANY:                                         "DEU",
	health_data.Country_COUNTRY_GREECE:                                          "GRC",
	health_data.Country_COUNTRY_GREENLAND:                                       "GRL",
	health_data.Country_COUNTRY_HUNGARY:                                         "HUN",
	health_data.Country_COUNTRY_ICELAND:                                         "ISL",
	health_data.Country_COUNTRY_IRELAND:                                         "IRL",
	health_data.Country_COUNTRY_ISRAEL:                                          "ISR",
	health_data.Country_COUNTRY_ITALY:                                           "ITA",
	health_data.Country_COUNTRY_KAZAKHSTAN:                                      "KAZ",
	health_data.Country_COUNTRY_KYRGYZSTAN:                                      "KGZ",
	health_data.Country_COUNTRY_LATVIA:                                          "LVA",
	health_data.Country_COUNTRY_LITHUANIA:                                       "LTU",
	health_data.Country_COUNTRY_LUXEMBOURG:                                      "LUX",
	health_data.Country_COUNTRY_MALTA:                                           "MLT",
	health_data.Country_COUNTRY_MONACO:                                          "MCO",
	health_data.Country_COUNTRY_MONTENEGRO:                                      "MNE",
	health_data.Country_COUNTRY_NETHERLANDS:                                     "NLD",
	health_data.Country_COUNTRY_NORTH_MACEDONIA:                                 "MKD",
	health_data.Country_COUNTRY_NORWAY:                                          "NOR",
	health_data.Country_COUNTRY_POLAND:                                          "POL",
	health_data.Country_COUNTRY_PORTUGAL:                                        "PRT",
	health_data.Country_COUNTRY_REPUBLIC_OF_MOLDOVA:                             "MDA",
	health_data.Country_COUNTRY_ROMANIA:                                         "ROU",
	health_data.Country_COUNTRY_RUSSIAN_FEDERATION:                              "RUS",
	health_data.Country_COUNTRY_SAN_MARINO:                                      "SMR",
	health_data.Country_COUNTRY_SERBIA:                                          "SRB",
	health_data.Country_COUNTRY_SERBIA_VOJVODINA:                                "RS-SRB",
	health_data.Country_COUNTRY_THE_FORMER_STATE_UNION_OF_SERBIA_AND_MONTENEGRO: "SCG",
	health_data.Country_COUNTRY_SLOVAKIA:                                        "SVK",
	health_data.Country_COUNTRY_SLOVENIA:                                        "SVN",
	health_data.Country_COUNTRY_SPAIN:                                           "ESP",
	health_data.Country_COUNTRY_SWEDEN:                                          "SWE",
	health_data.Country_COUNTRY_SWITZERLAND:                                     "CHE",
	health_data.Country_COUNTRY_TAJIKISTAN:                                      "TJK",
	health_data.Country_COUNTRY_TURKIYE:                                         "TUR",
	health_data.Country_COUNTRY_TURKMENISTAN:                                    "TKM",
	health_data.Country_COUNTRY_UKRAINE:                                         "UKR",
	health_data.Country_COUNTRY_UNITED_KINGDOM:                                  "GBR",
	health_data.Country_COUNTRY_UNITED_KINGDOM_ENGLAND:                          "GB-ENG",
	health_data.Country_COUNTRY_UNITED_KINGDOM_NORTHERN_IRELAND:                 "GB-NIR",
	health_data.Country_COUNTRY_UNITED_KINGDOM_SCOTLAND:                         "GB-SCT",
	health_data.Country_COUNTRY_UNITED_KINGDOM_WALES:                            "GB-WLS",
	health_data.Country_COUNTRY_UNITED_STATES_OF_AMERICA:                        "USA",
	health_data.Country_COUNTRY_REPUBLIC_OF_UZBEKISTAN:                          "UZB",
	health_data.Country_COUNTRY_KOSOVO:                                          "RS-XKX",
}

var dictionary_codes map[health_data.WHOEuropeCodes]string = map[health_data.WHOEuropeCodes]string{
	health_data.WHOEuropeCodes_WHO_EUROPE_CODE_INFANT_MORTALITY_INDICATOR:         "HFA_73",
	health_data.WHOEuropeCodes_WHO_EUROPE_CODE_EARLY_NEONATAL_MORTALITY_INDICATOR: "HFA_78",
	health_data.WHOEuropeCodes_WHO_EUROPE_CODE_CIRCULATORY_DISEASES_INDICATOR:     "HFA_98",
	health_data.WHOEuropeCodes_WHO_EUROPE_CODE_ISCHAEMIC_HEART_DISEASE_INDICATOR:  "HFA_107",
	health_data.WHOEuropeCodes_WHO_EUROPE_CODE_CEREBROVASCULAR_DISEASES_INDICATOR: "HFA_119",
}

// HFA_73: Estimated infant mortality per 1000 live births (world health report)
var infant_mortality_indicator string = "HFA_73"

// HFA_78: Early neonatal deaths per 1000 live births
var early_neonatal_mortality_indicator string = "HFA_78"

// HFA_98: Diseases of circulatory system, 0–64, per 100 000, by sex (age-standardized death rate)
var circulatory_diseases_indicator string = "HFA_98"

// HFA_107: Ischaemic heart disease, 0–64, per 100 000, by sex (age-standardized death rate)
var ischaemic_heart_disease_indicator string = "HFA_107"

// HFA_119: Cerebrovascular diseases, all ages, per 100 000, by sex (age-standardized death rate)
var cerebrovascular_diseases_indicator string = "HFA_119"

func GetCountryCode(country health_data.Country) (string, error) {
	code, exists := dictionary_countries[country]
	if !exists {
		return "", fmt.Errorf("country code not found for country: %v", country)
	}
	return code, nil
}

func GetCountry(code string) (health_data.Country, error) {
	for country, c := range dictionary_countries {
		if c == code {
			return country, nil
		}
	}
	return health_data.Country_COUNTRY_UNKNOWN, fmt.Errorf("country not found for code: %s", code)
}

func GetIndicatorCode(indicator health_data.WHOEuropeCodes) (string, error) {
	code, exists := dictionary_codes[indicator]
	if !exists {
		return "", fmt.Errorf("indicator code not found for indicator: %v", indicator)
	}
	return code, nil
}

func GetIndicator(code string) (health_data.WHOEuropeCodes, error) {
	for indicator, c := range dictionary_codes {
		if c == code {
			return indicator, nil
		}
	}
	return health_data.WHOEuropeCodes_WHO_EUROPE_CODE_UNKNOWN, fmt.Errorf("indicator not found for code: %s", code)
}

// ParseWHOEuropeCSVToReply parses a WHO-Europe CSV-style string (like the sample you provided)
// into a health_data.WhoEuropeResponse. The function is tolerant: it first extracts simple metadata
// key/value pairs until it finds the data header row ("COUNTRY","COUNTRY_GRP","SEX","YEAR","VALUE"),
// then collects data rows into objects and attempts to unmarshal a JSON representation into
// health_data.WhoEuropeResponse using protojson with unknown fields discarded.
//
// Note: this implementation assumes the proto message uses common JSON field names such as
// "indicator", "lastUpdate", "description", "referenceLink" and a repeated "data" field whose
// elements map keys "COUNTRY","COUNTRY_GRP","SEX","YEAR","VALUE" (lowercased in JSON).
// If your generated proto uses different field names adjust the jsonMap construction accordingly.
func ParseWHOEuropeCSVToReply(csvStr string) (*pb.WHOEuropeResponse, error) {
	csvStr = strings.ReplaceAll(csvStr, "\ufeff", "") // normalize line endings
	r := csv.NewReader(strings.NewReader(csvStr))
	r.FieldsPerRecord = -1
	records, err := r.ReadAll()
	if err != nil {
		return nil, fmt.Errorf("csv read error: %w", err)
	}

	datasets := []string{}
	response := &pb.WHOEuropeResponse{
		Csv: csvStr,
	}
	datas := []*health_data.WHOData{}

	response.DataSource = records[4][1]
	response.Unit = records[6][1]
	response.Representation = records[8][1]
	breakIndex := 0
	for i, rec := range records {
		if i < 10 {
			continue
		}
		if rec[0] == "DATA_MASK" {
			response.Masks = records[i][1]
			breakIndex = i + 6
			break
		}
		datasets = append(datasets, rec[1])
	}
	response.DataSet = datasets
	// Walk rows: metadata until header row is found, then data rows according to header
	for i, rec := range records[breakIndex:] {
		// skip empty records
		if len(rec) == 0 {
			continue
		}
		if rec[0] == "Last update" {
			response.LastUpdate = rec[1]
			response.Description = records[i+breakIndex+1][1]
			response.ReferenceLink = records[i+breakIndex+2][1]
			response.Copyright = records[i+breakIndex+4][0]
			break
		}
		// trim spaces on all fields
		for i := range rec {
			rec[i] = strings.TrimSpace(rec[i])
		}

		// parse year
		yearVal, err := strconv.Atoi(rec[3])
		if err != nil {
			return nil, fmt.Errorf("invalid year value %q at row %d: %w", rec[3], i+breakIndex, err)
		}
		// parse value
		val64, err := strconv.ParseFloat(rec[4], 32)
		if err != nil {
			return nil, fmt.Errorf("invalid numeric value %q at row %d: %w", rec[4], i+breakIndex, err)
		}

		data := &health_data.WHOData{
			Country:    rec[0],
			CountryGrp: rec[1],
			Sex:        rec[2],
			Year:       int32(yearVal),
			Value:      float32(val64),
		}
		datas = append(datas, data)
	}

	response.Data = datas
	return response, nil
}

func ConvertDateParamToQuery(start_date *pb.Date, end_date *pb.Date) string {
	years := ";YEAR:"
	if start_date != nil {
		if end_date != nil {
			years += fmt.Sprintf("%d-%d", start_date.Year, end_date.Year)
		} else {
			years += fmt.Sprintf("%d-%d", start_date.Year, time.Now().Year())
		}
	} else if end_date != nil {
		years += fmt.Sprintf("%d-%d", end_date.Year-100, end_date.Year)
	} else {
		years = ""
	}
	return years
}
