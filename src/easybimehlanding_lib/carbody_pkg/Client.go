/*
 * easybimehlanding_lib
 *
 * This file was automatically generated by APIMATIC v2.0 ( https://apimatic.io ).
 */

package carbody_pkg


import(
	"encoding/json"
	"github.com/apimatic/unirest-go"
	"easybimehlanding_lib/apihelper_pkg"
	"easybimehlanding_lib/configuration_pkg"
	"easybimehlanding_lib/models_pkg"
)
/*
 * Client structure as interface implementation
 */
type CARBODY_IMPL struct {
     config configuration_pkg.CONFIGURATION
}

/**
 * دریافت برند خودرو
 * @param    string        xApiKey       parameter: Required
 * @return	Returns the *models_pkg.CarBrand response from the API call
 */
func (me *CARBODY_IMPL) GetCarBrand (
            xApiKey string) (*models_pkg.CarBrand, error) {
    //the endpoint path uri
    _pathUrl := "/ComboData/CarBrands"

    //variable to hold errors
    var err error = nil
    //the base uri for api requests
    _queryBuilder := configuration_pkg.GetBaseURI(configuration_pkg.SERVER_1,me.config);

    //prepare query string for API call
   _queryBuilder = _queryBuilder + _pathUrl

    //validate and preprocess url
    _queryBuilder, err = apihelper_pkg.CleanUrl(_queryBuilder)
    if err != nil {
        //error in url validation or cleaning
        return nil, err
    }
    //prepare headers for the outgoing request
    headers := map[string]interface{} {
        "user-agent" : "APIMATIC 2.0",
        "accept" : "application/json",
        "x-api-key" : apihelper_pkg.ToString(xApiKey, ""),
    }

    //prepare API request
    _request := unirest.Get(_queryBuilder, headers)
    //and invoke the API call request to fetch the response
    _response, err := unirest.AsString(_request,false);
    if err != nil {
        //error in API invocation
        return nil, err
    }

    //error handling using HTTP status codes
    if (_response.Code < 200) || (_response.Code > 206) { //[200,206] = HTTP OK
        err = apihelper_pkg.NewAPIError("HTTP Response Not OK" , _response.Code, _response.RawBody)
    }
    if(err != nil) {
        //error detected in status code validation
        return nil, err
    }

    //returning the response
    var retVal *models_pkg.CarBrand = &models_pkg.CarBrand{}
    err = json.Unmarshal(_response.RawBody, &retVal)

    if err != nil {
        //error in parsing
        return nil, err
    }
    return retVal, nil

}

/**
 * دریافت لیست تیپ خودرو
 * @param    int64         carBrandId     parameter: Required
 * @param    string        xApiKey        parameter: Required
 * @return	Returns the *models_pkg.CarBrandTips response from the API call
 */
func (me *CARBODY_IMPL) GetCarBrandTips (
            carBrandId int64,
            xApiKey string) (*models_pkg.CarBrandTips, error) {
    //the endpoint path uri
    _pathUrl := "/ComboData/CarBrandTips"

    //variable to hold errors
    var err error = nil
    //the base uri for api requests
    _queryBuilder := configuration_pkg.GetBaseURI(configuration_pkg.SERVER_1,me.config);

    //prepare query string for API call
   _queryBuilder = _queryBuilder + _pathUrl

    //process optional query parameters
    _queryBuilder, err = apihelper_pkg.AppendUrlWithQueryParameters(_queryBuilder, map[string]interface{} {
        "carBrandId" : carBrandId,
    })
    if err != nil {
        //error in query param handling
        return nil, err
    }

    //validate and preprocess url
    _queryBuilder, err = apihelper_pkg.CleanUrl(_queryBuilder)
    if err != nil {
        //error in url validation or cleaning
        return nil, err
    }
    //prepare headers for the outgoing request
    headers := map[string]interface{} {
        "user-agent" : "APIMATIC 2.0",
        "accept" : "application/json",
        "x-api-key" : apihelper_pkg.ToString(xApiKey, ""),
    }

    //prepare API request
    _request := unirest.Get(_queryBuilder, headers)
    //and invoke the API call request to fetch the response
    _response, err := unirest.AsString(_request,false);
    if err != nil {
        //error in API invocation
        return nil, err
    }

    //error handling using HTTP status codes
    if (_response.Code < 200) || (_response.Code > 206) { //[200,206] = HTTP OK
        err = apihelper_pkg.NewAPIError("HTTP Response Not OK" , _response.Code, _response.RawBody)
    }
    if(err != nil) {
        //error detected in status code validation
        return nil, err
    }

    //returning the response
    var retVal *models_pkg.CarBrandTips = &models_pkg.CarBrandTips{}
    err = json.Unmarshal(_response.RawBody, &retVal)

    if err != nil {
        //error in parsing
        return nil, err
    }
    return retVal, nil

}

/**
 * آیا این نوع بیمه نامه، طرح بیمه ای دارد؟
 * @param    string        subDomain               parameter: Required
 * @param    int64         insurancePolicyType     parameter: Required
 * @param    string        xApiKey                 parameter: Required
 * @return	Returns the *models_pkg.HasPlan response from the API call
 */
func (me *CARBODY_IMPL) GetHasPlan (
            subDomain string,
            insurancePolicyType int64,
            xApiKey string) (*models_pkg.HasPlan, error) {
    //the endpoint path uri
    _pathUrl := "/InsurancePolicyPlan/HasPlan"

    //variable to hold errors
    var err error = nil
    //the base uri for api requests
    _queryBuilder := configuration_pkg.GetBaseURI(configuration_pkg.SERVER_1,me.config);

    //prepare query string for API call
   _queryBuilder = _queryBuilder + _pathUrl

    //process optional query parameters
    _queryBuilder, err = apihelper_pkg.AppendUrlWithQueryParameters(_queryBuilder, map[string]interface{} {
        "subDomain" : subDomain,
        "insurancePolicyType" : insurancePolicyType,
    })
    if err != nil {
        //error in query param handling
        return nil, err
    }

    //validate and preprocess url
    _queryBuilder, err = apihelper_pkg.CleanUrl(_queryBuilder)
    if err != nil {
        //error in url validation or cleaning
        return nil, err
    }
    //prepare headers for the outgoing request
    headers := map[string]interface{} {
        "user-agent" : "APIMATIC 2.0",
        "accept" : "application/json",
        "x-api-key" : apihelper_pkg.ToString(xApiKey, ""),
    }

    //prepare API request
    _request := unirest.Get(_queryBuilder, headers)
    //and invoke the API call request to fetch the response
    _response, err := unirest.AsString(_request,false);
    if err != nil {
        //error in API invocation
        return nil, err
    }

    //error handling using HTTP status codes
    if (_response.Code < 200) || (_response.Code > 206) { //[200,206] = HTTP OK
        err = apihelper_pkg.NewAPIError("HTTP Response Not OK" , _response.Code, _response.RawBody)
    }
    if(err != nil) {
        //error detected in status code validation
        return nil, err
    }

    //returning the response
    var retVal *models_pkg.HasPlan = &models_pkg.HasPlan{}
    err = json.Unmarshal(_response.RawBody, &retVal)

    if err != nil {
        //error in parsing
        return nil, err
    }
    return retVal, nil

}

