/*
 * easybimehlanding_lib
 *
 * This file was automatically generated by APIMATIC v2.0 ( https://apimatic.io ).
 */

package fireinsurance_pkg

import "easybimehlanding_lib/configuration_pkg"
import "easybimehlanding_lib/models_pkg"

/*
 * Interface for the FIREINSURANCE_IMPL
 */
type FIREINSURANCE interface {
    GetFireInsurance (string, int64, string) (*models_pkg.BaseModelFireInsurance, error)
}

/*
 * Factory for the FIREINSURANCE interaface returning FIREINSURANCE_IMPL
 */
func NewFIREINSURANCE(config configuration_pkg.CONFIGURATION) *FIREINSURANCE_IMPL {
    client := new(FIREINSURANCE_IMPL)
    client.config = config
    return client
}