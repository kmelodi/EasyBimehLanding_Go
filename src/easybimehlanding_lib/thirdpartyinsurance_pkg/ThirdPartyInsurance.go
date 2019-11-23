/*
 * easybimehlanding_lib
 *
 * This file was automatically generated by APIMATIC v2.0 ( https://apimatic.io ).
 */

package thirdpartyinsurance_pkg

import "easybimehlanding_lib/configuration_pkg"
import "easybimehlanding_lib/models_pkg"

/*
 * Interface for the THIRDPARTYINSURANCE_IMPL
 */
type THIRDPARTYINSURANCE interface {
    GetCarBrands (int64, string) (*models_pkg.CarBrands, error)

    GetRiskLevel (int64, string) (*models_pkg.RiskLevel, error)

    GetCarBrandTips (int64, string) (*models_pkg.CarBrandTips, error)

    GetCarUses (int64, string) (*models_pkg.CarUses, error)

    GetHasPlan (string, int64, string) (*models_pkg.HasPlan, error)
}

/*
 * Factory for the THIRDPARTYINSURANCE interaface returning THIRDPARTYINSURANCE_IMPL
 */
func NewTHIRDPARTYINSURANCE(config configuration_pkg.CONFIGURATION) *THIRDPARTYINSURANCE_IMPL {
    client := new(THIRDPARTYINSURANCE_IMPL)
    client.config = config
    return client
}