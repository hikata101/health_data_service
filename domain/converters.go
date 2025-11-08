package domain

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"strings"

	pb "github.com/hikata101/health_data_service/gen/github.com/hikata101/health_data_service/v1"
	"google.golang.org/protobuf/encoding/protojson"
)

var dictionary_countries = map[pb.Country]string{
	pb.Country_COUNTRY_ALBANIA:                                         "ALB",
	pb.Country_COUNTRY_ANDORRA:                                         "AND",
	pb.Country_COUNTRY_ARMENIA:                                         "ARM",
	pb.Country_COUNTRY_AUSTRIA:                                         "AUT",
	pb.Country_COUNTRY_AZERBAIJAN:                                      "AZE",
	pb.Country_COUNTRY_BELARUS:                                         "BLR",
	pb.Country_COUNTRY_BELGIUM:                                         "BEL",
	pb.Country_COUNTRY_BELGIUM_BRUXELLES:                               "BE-BRU",
	pb.Country_COUNTRY_BELGIUM_FLANDERS:                                "BE-VLG",
	pb.Country_COUNTRY_BELGIUM_WALLONIA:                                "BE-WAL",
	pb.Country_COUNTRY_BOSNIA_AND_HERZEGOVINA:                          "BIH",
	pb.Country_COUNTRY_BULGARIA:                                        "BGR",
	pb.Country_COUNTRY_CANADA:                                          "CAN",
	pb.Country_COUNTRY_CROATIA:                                         "HRV",
	pb.Country_COUNTRY_CYPRUS:                                          "CYP",
	pb.Country_COUNTRY_CZECH_REPUBLIC:                                  "CZE",
	pb.Country_COUNTRY_DENMARK:                                         "DNK",
	pb.Country_COUNTRY_ESTONIA:                                         "EST",
	pb.Country_COUNTRY_FINLAND:                                         "FIN",
	pb.Country_COUNTRY_FRANCE:                                          "FRA",
	pb.Country_COUNTRY_GEORGIA:                                         "GEO",
	pb.Country_COUNTRY_GERMANY:                                         "DEU",
	pb.Country_COUNTRY_GREECE:                                          "GRC",
	pb.Country_COUNTRY_GREENLAND:                                       "GRL",
	pb.Country_COUNTRY_HUNGARY:                                         "HUN",
	pb.Country_COUNTRY_ICELAND:                                         "ISL",
	pb.Country_COUNTRY_IRELAND:                                         "IRL",
	pb.Country_COUNTRY_ISRAEL:                                          "ISR",
	pb.Country_COUNTRY_ITALY:                                           "ITA",
	pb.Country_COUNTRY_KAZAKHSTAN:                                      "KAZ",
	pb.Country_COUNTRY_KYRGYZSTAN:                                      "KGZ",
	pb.Country_COUNTRY_LATVIA:                                          "LVA",
	pb.Country_COUNTRY_LITHUANIA:                                       "LTU",
	pb.Country_COUNTRY_LUXEMBOURG:                                      "LUX",
	pb.Country_COUNTRY_MALTA:                                           "MLT",
	pb.Country_COUNTRY_MONACO:                                          "MCO",
	pb.Country_COUNTRY_MONTENEGRO:                                      "MNE",
	pb.Country_COUNTRY_NETHERLANDS:                                     "NLD",
	pb.Country_COUNTRY_NORTH_MACEDONIA:                                 "MKD",
	pb.Country_COUNTRY_NORWAY:                                          "NOR",
	pb.Country_COUNTRY_POLAND:                                          "POL",
	pb.Country_COUNTRY_PORTUGAL:                                        "PRT",
	pb.Country_COUNTRY_REPUBLIC_OF_MOLDOVA:                             "MDA",
	pb.Country_COUNTRY_ROMANIA:                                         "ROU",
	pb.Country_COUNTRY_RUSSIAN_FEDERATION:                              "RUS",
	pb.Country_COUNTRY_SAN_MARINO:                                      "SMR",
	pb.Country_COUNTRY_SERBIA:                                          "SRB",
	pb.Country_COUNTRY_SERBIA_VOJVODINA:                                "RS-SRB",
	pb.Country_COUNTRY_THE_FORMER_STATE_UNION_OF_SERBIA_AND_MONTENEGRO: "SCG",
	pb.Country_COUNTRY_SLOVAKIA:                                        "SVK",
	pb.Country_COUNTRY_SLOVENIA:                                        "SVN",
	pb.Country_COUNTRY_SPAIN:                                           "ESP",
	pb.Country_COUNTRY_SWEDEN:                                          "SWE",
	pb.Country_COUNTRY_SWITZERLAND:                                     "CHE",
	pb.Country_COUNTRY_TAJIKISTAN:                                      "TJK",
	pb.Country_COUNTRY_TURKIYE:                                         "TUR",
	pb.Country_COUNTRY_TURKMENISTAN:                                    "TKM",
	pb.Country_COUNTRY_UKRAINE:                                         "UKR",
	pb.Country_COUNTRY_UNITED_KINGDOM:                                  "GBR",
	pb.Country_COUNTRY_UNITED_KINGDOM_ENGLAND:                          "GB-ENG",
	pb.Country_COUNTRY_UNITED_KINGDOM_NORTHERN_IRELAND:                 "GB-NIR",
	pb.Country_COUNTRY_UNITED_KINGDOM_SCOTLAND:                         "GB-SCT",
	pb.Country_COUNTRY_UNITED_KINGDOM_WALES:                            "GB-WLS",
	pb.Country_COUNTRY_UNITED_STATES_OF_AMERICA:                        "USA",
	pb.Country_COUNTRY_REPUBLIC_OF_UZBEKISTAN:                          "UZB",
	pb.Country_COUNTRY_KOSOVO:                                          "RS-XKX",
}

var dictionary_codes map[pb.WHOEuropeCodes]string = map[pb.WHOEuropeCodes]string{
	pb.WHOEuropeCodes_WHO_EUROPE_CODE_INFANT_MORTALITY_INDICATOR:         "HFA_73",
	pb.WHOEuropeCodes_WHO_EUROPE_CODE_EARLY_NEONATAL_MORTALITY_INDICATOR: "HFA_78",
	pb.WHOEuropeCodes_WHO_EUROPE_CODE_CIRCULATORY_DISEASES_INDICATOR:     "HFA_98",
	pb.WHOEuropeCodes_WHO_EUROPE_CODE_ISCHAEMIC_HEART_DISEASE_INDICATOR:  "HFA_107",
	pb.WHOEuropeCodes_WHO_EUROPE_CODE_CEREBROVASCULAR_DISEASES_INDICATOR: "HFA_119",
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

func GetCountryCode(country pb.Country) (string, error) {
	code, exists := dictionary_countries[country]
	if !exists {
		return "", fmt.Errorf("country code not found for country: %v", country)
	}
	return code, nil
}

func GetCountry(code string) (pb.Country, error) {
	for country, c := range dictionary_countries {
		if c == code {
			return country, nil
		}
	}
	return pb.Country_COUNTRY_UNKNOWN, fmt.Errorf("country not found for code: %s", code)
}

func GetIndicatorCode(indicator pb.WHOEuropeCodes) (string, error) {
	code, exists := dictionary_codes[indicator]
	if !exists {
		return "", fmt.Errorf("indicator code not found for indicator: %v", indicator)
	}
	return code, nil
}

func GetIndicator(code string) (pb.WHOEuropeCodes, error) {
	for indicator, c := range dictionary_codes {
		if c == code {
			return indicator, nil
		}
	}
	return pb.WHOEuropeCodes_WHO_EUROPE_CODE_UNKNOWN, fmt.Errorf("indicator not found for code: %s", code)
}

// ParseWHOEuropeCSVToReply parses a WHO-Europe CSV-style string (like the sample you provided)
// into a pb.WHOEuropeReply. The function is tolerant: it first extracts simple metadata
// key/value pairs until it finds the data header row ("COUNTRY","COUNTRY_GRP","SEX","YEAR","VALUE"),
// then collects data rows into objects and attempts to unmarshal a JSON representation into
// pb.WHOEuropeReply using protojson with unknown fields discarded.
//
// Note: this implementation assumes the proto message uses common JSON field names such as
// "indicator", "lastUpdate", "description", "referenceLink" and a repeated "data" field whose
// elements map keys "COUNTRY","COUNTRY_GRP","SEX","YEAR","VALUE" (lowercased in JSON).
// If your generated proto uses different field names adjust the jsonMap construction accordingly.
func ParseWHOEuropeCSVToReply(csvStr string) (*pb.WHOEuropeReply, error) {
	r := csv.NewReader(strings.NewReader(csvStr))
	r.FieldsPerRecord = -1
	records, err := r.ReadAll()
	if err != nil {
		return nil, fmt.Errorf("csv read error: %w", err)
	}

	metadata := map[string]string{}
	var dataRows []map[string]string
	headerFound := false
	var header []string

	// Walk rows: metadata until header row is found, then data rows according to header
	for _, rec := range records {
		// skip empty records
		if len(rec) == 0 {
			continue
		}

		// trim spaces on all fields
		for i := range rec {
			rec[i] = strings.TrimSpace(rec[i])
		}

		// detect header row
		if len(rec) >= 5 && rec[0] == "COUNTRY" && rec[1] == "COUNTRY_GRP" && rec[2] == "SEX" {
			headerFound = true
			header = rec
			continue
		}

		if !headerFound {
			// treat as metadata key,value pairs when there are at least two columns
			if len(rec) >= 2 {
				key := rec[0]
				val := rec[1]
				metadata[key] = val
			}
			continue
		}

		// headerFound == true -> data rows
		// Map fields by header names
		rowMap := map[string]string{}
		for i, h := range header {
			if i < len(rec) {
				rowMap[h] = rec[i]
			} else {
				rowMap[h] = ""
			}
		}
		dataRows = append(dataRows, rowMap)
	}

	// Build a JSON-friendly map matching expected proto JSON shape
	jsonMap := map[string]interface{}{}

	// common metadata keys from your sample
	if v, ok := metadata["INDICATOR"]; ok {
		jsonMap["indicator"] = v
	}
	if v, ok := metadata["Last update"]; ok {
		jsonMap["lastUpdate"] = v
	}
	if v, ok := metadata["Description"]; ok {
		jsonMap["description"] = v
	}
	if v, ok := metadata["Reference link"]; ok {
		jsonMap["referenceLink"] = v
	}
	if v, ok := metadata["UNIT_TYPE"]; ok {
		jsonMap["unitType"] = v
	}
	// include the parsed data rows under "data"
	jsonMap["data"] = dataRows

	b, err := json.Marshal(jsonMap)
	if err != nil {
		return nil, fmt.Errorf("json marshal error: %w", err)
	}

	reply := &pb.WHOEuropeReply{}
	opts := protojson.UnmarshalOptions{DiscardUnknown: true}
	if err := opts.Unmarshal(b, reply); err != nil {
		return nil, fmt.Errorf("failed to unmarshal into pb.WHOEuropeReply: %w", err)
	}
	return reply, nil
}
