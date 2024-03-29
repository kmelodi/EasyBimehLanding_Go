/*
 * easybimehlanding_lib
 *
 * This file was automatically generated by APIMATIC v2.0 ( https://apimatic.io ).
 */

package liabilitydoctorinsurance_pkg

import "easybimehlanding_lib/configuration_pkg"
import "easybimehlanding_lib/models_pkg"

/*
 * Interface for the LIABILITYDOCTORINSURANCE_IMPL
 */
type LIABILITYDOCTORINSURANCE interface {
    GetLiabilityDoctorInsurance (string, string) (*models_pkg.BaseModelLiabilityDoctorInsurance, error)

    GetMedicalSpecialties (string, string) (*models_pkg.BaseModelMedicalSpecialties, error)
}

/*
 * Factory for the LIABILITYDOCTORINSURANCE interaface returning LIABILITYDOCTORINSURANCE_IMPL
 */
func NewLIABILITYDOCTORINSURANCE(config configuration_pkg.CONFIGURATION) *LIABILITYDOCTORINSURANCE_IMPL {
    client := new(LIABILITYDOCTORINSURANCE_IMPL)
    client.config = config
    return client
}