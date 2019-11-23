/*
 * easybimehlanding_lib
 *
 * This file was automatically generated by APIMATIC v2.0 ( https://apimatic.io ).
 */

package main_pkg

import "easybimehlanding_lib/configuration_pkg"
import "easybimehlanding_lib/models_pkg"

/*
 * Interface for the MAIN_IMPL
 */
type MAIN interface {
    GetPortalLandingPage (string, string) (*models_pkg.BaseModelPortalLandingPage, error)

    GetInsuranceCentrePolicyTypes (string, string) (*models_pkg.BaseModelInsuranceCentrePolicyTypes, error)
}

/*
 * Factory for the MAIN interaface returning MAIN_IMPL
 */
func NewMAIN(config configuration_pkg.CONFIGURATION) *MAIN_IMPL {
    client := new(MAIN_IMPL)
    client.config = config
    return client
}