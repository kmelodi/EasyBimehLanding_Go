/*
 * easybimehlanding_lib
 *
 * This file was automatically generated by APIMATIC v2.0 ( https://apimatic.io ).
 */

package elevatorinsurance_pkg

import "easybimehlanding_lib/configuration_pkg"
import "easybimehlanding_lib/models_pkg"

/*
 * Interface for the ELEVATORINSURANCE_IMPL
 */
type ELEVATORINSURANCE interface {
    GetElevatorInsurance (string, int64, string) (*models_pkg.BaseModelElevatorInsurance, error)
}

/*
 * Factory for the ELEVATORINSURANCE interaface returning ELEVATORINSURANCE_IMPL
 */
func NewELEVATORINSURANCE(config configuration_pkg.CONFIGURATION) *ELEVATORINSURANCE_IMPL {
    client := new(ELEVATORINSURANCE_IMPL)
    client.config = config
    return client
}