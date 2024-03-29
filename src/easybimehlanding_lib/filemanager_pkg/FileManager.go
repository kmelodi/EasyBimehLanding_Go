/*
 * easybimehlanding_lib
 *
 * This file was automatically generated by APIMATIC v2.0 ( https://apimatic.io ).
 */

package filemanager_pkg

import "easybimehlanding_lib/configuration_pkg"
import "easybimehlanding_lib/models_pkg"

/*
 * Interface for the FILEMANAGER_IMPL
 */
type FILEMANAGER interface {
    Upload (string, string, string) (*models_pkg.BaseModelUpload, error)
}

/*
 * Factory for the FILEMANAGER interaface returning FILEMANAGER_IMPL
 */
func NewFILEMANAGER(config configuration_pkg.CONFIGURATION) *FILEMANAGER_IMPL {
    client := new(FILEMANAGER_IMPL)
    client.config = config
    return client
}