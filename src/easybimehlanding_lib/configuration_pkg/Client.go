/*
 * easybimehlanding_lib
 *
 * This file was automatically generated by APIMATIC v2.0 ( https://apimatic.io ).
 */

package configuration_pkg

import(
	"easybimehlanding_lib/apihelper_pkg"
)
/* Setting up enums for Environments and Servers 
*/
type Environments int

type Servers int

// Environment Enums
const (
     INSURANCE_SERVER_NOTIFAANO_IR Environments = 1 + iota
)

// Servers Enums
const (
 	 SERVER_1 Servers = 1 + iota
)

type CONFIGURATION_IMPL struct {
}

// Setting up Default Environment
var Environment = INSURANCE_SERVER_NOTIFAANO_IR

// A map of environments and their corresponding servers/baseurls
var EnvironmentsMap = map[Environments](map[Servers]string){

    INSURANCE_SERVER_NOTIFAANO_IR : map[Servers]string{
        SERVER_1:"http://server.notifaano.ir/api",
    },
}
 
// Make and return the map of parameters
func GetBaseURIParameters(config CONFIGURATION) map[string]interface{} {
     kvpMap := map[string]interface{}{
    }
    return kvpMap;
}

// Gets the URL for a particular alias in the current environment and appends it with template parameters
// return the baseurl
func GetBaseURI(server Servers, config CONFIGURATION) string {
    url := EnvironmentsMap[Environment][server]
    appendedURL, _ := apihelper_pkg.AppendUrlWithTemplateParameters(url, GetBaseURIParameters(config))
    return appendedURL

}
