package util

import (
	"math"

	centralConstants "github.com/FrosTiK-SD/models/constant"
	miscModels "github.com/FrosTiK-SD/models/misc"
	opportunityModels "github.com/FrosTiK-SD/models/opportunity"
	jsoniter "github.com/json-iterator/go"
)

var json = jsoniter.ConfigCompatibleWithStandardLibrary

type BranchCompensation map[centralConstants.Branch]*opportunityModels.CompensationBreakup

type CourseCompensation map[centralConstants.Course]*BranchCompensation

type CompanyProfileMini struct {
	CompensationDetails CourseCompensation `bson:"compensationDetails" json:"compensationDetails"`
}

func GetCompensationRange(companyProfiles *[]opportunityModels.CompanyProfile) *[]opportunityModels.CompensationRange {
	var companyProfilesList []CompanyProfileMini
	currencyCompensationRange := make(map[miscModels.CurrencyType]*opportunityModels.CompensationRange)

	compensationBytes, _ := json.Marshal(companyProfiles)
	_ = json.Unmarshal(compensationBytes, &companyProfilesList)

	// Iterate over every company profile
	for _, compensationMap := range companyProfilesList {

		// Iterate over every course
		for _, course := range compensationMap.CompensationDetails {
			if course == nil {
				continue
			}

			// Iterate over branches
			for _, branch := range *course {
				// If TotalCTC is not set
				if branch == nil || (branch.TotalCTC == nil && branch.Fixed == nil) {
					continue
				}

				// Default to TotalCTC
				branchCompensation := branch.TotalCTC

				// If TotalCTC is not set then choose Fixed
				if branchCompensation == nil {
					branchCompensation = branch.Fixed
				}

				// For MIN
				if branchCompensation.Min != nil {
					if currencyCompensationRange[branchCompensation.Min.Currency] == nil {
						currencyCompensationRange[branchCompensation.Min.Currency] = &opportunityModels.CompensationRange{
							Min: branchCompensation.Min,
						}
					} else {
						if currencyCompensationRange[branchCompensation.Min.Currency].Min == nil {
							currencyCompensationRange[branchCompensation.Min.Currency].Min = branchCompensation.Min
						} else {
							currencyCompensationRange[branchCompensation.Min.Currency].Min.Amount = math.Min(currencyCompensationRange[branchCompensation.Min.Currency].Min.Amount, branchCompensation.Min.Amount)
						}
					}
				}

				// For MAX
				if branchCompensation.Max != nil {
					if currencyCompensationRange[branchCompensation.Max.Currency] == nil {
						currencyCompensationRange[branchCompensation.Max.Currency] = &opportunityModels.CompensationRange{
							Max: branchCompensation.Max,
						}
					} else {
						if currencyCompensationRange[branchCompensation.Max.Currency].Max == nil {
							currencyCompensationRange[branchCompensation.Max.Currency].Max = branchCompensation.Max
						} else {
							currencyCompensationRange[branchCompensation.Max.Currency].Max.Amount = math.Max(currencyCompensationRange[branchCompensation.Max.Currency].Max.Amount, branchCompensation.Max.Amount)
						}
					}
				}

			}
		}
	}

	compensationRange := make([]opportunityModels.CompensationRange, 0)
	for _, currency := range currencyCompensationRange {
		if currency != nil {
			compensationRange = append(compensationRange, *currency)
		}
	}

	return &compensationRange
}
