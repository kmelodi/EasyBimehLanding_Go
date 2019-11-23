# Getting started

EasyBimehConnect یک ساب برند از ایزی بیمه است که وظیفه ارائه خدمات B2B را برپایه API و B2B2C را بر پایه وایت لیبل بر عهده دارد. اگر اپلیکیشن و یا سایت غیر بیمه‌ای دارید و تمایل به فروش بیمه نامه دارید از امروز میتوانید با کمترین هزینه و کمترین زمان به زنجیره نوآوری در صنعت بیمه متصل شوید و تجارت جدیدی بسازید.

## How to Build


* In order to successfully build and run your SDK files, you are required to have the following setup in your system:
    * **Go**  (Visit [https://golang.org/doc/install](https://golang.org/doc/install) for more details on how to install Go)
    * **Java VM** Version 8 or later
    * **Eclipse 4.6 (Neon)** or later ([http://www.eclipse.org/neon/](http://www.eclipse.org/neon/))
    * **GoClipse** setup within above installed Eclipse (Follow the instructions at [this link](https://github.com/GoClipse/goclipse/blob/latest/documentation/Installation.md#instructions) to setup GoClipse)
* Ensure that ```GOPATH``` environment variable is set in the system variables. If not, set it to your workspace directory where you will be adding your Go projects.
* The generated code uses unirest-go http library. Therefore, you will need internet access to resolve this dependency. If Go is properly installed and configured, run the following command to pull the dependency:

```Go
go get github.com/apimatic/unirest-go
```

This will install unirest-go in the ```GOPATH``` you specified in the system variables.

Now follow the steps mentioned below to build your SDK:

1. Open eclipse in the Go language perspective and click on the ```Import``` option in ```File``` menu.

![Importing SDK into Eclipse - Step 1](https://apidocs.io/illustration/go?step=import0)

2. Select ```General -> Existing Projects into Workspace``` option from the tree list.

![Importing SDK into Eclipse - Step 2](https://apidocs.io/illustration/go?step=import1)

3. In ```Select root directory```, provide path to the unzipped archive for the generated code. Once the path is set and the Project becomes visible under ```Projects``` click ```Finish```

![Importing SDK into Eclipse - Step 3](https://apidocs.io/illustration/go?step=import2&workspaceFolder=Easy%20Bimeh%20Landing-GoLang&projectName=easybimehlanding_lib)

4. The Go library will be imported and its files will be visible in the Project Explorer

![Importing SDK into Eclipse - Step 4](https://apidocs.io/illustration/go?step=import3&projectName=easybimehlanding_lib)

## How to Use

The following section explains how to use the EasybimehlandingLib library in a new project.

### 1. Add a new Test Project

Create a new project in Eclipse by ```File``` -> ```New``` -> ```Go Project```

![Add a new project in Eclipse](https://apidocs.io/illustration/go?step=createNewProject0)

Name the Project as ```Test``` and click ```Finish```

![Create a new Maven Project - Step 1](https://apidocs.io/illustration/go?step=createNewProject1)

Create a new directory in the ```src``` directory of this project

![Create a new Maven Project - Step 2](https://apidocs.io/illustration/go?step=createNewProject2&projectName=easybimehlanding_lib)

Name it ```test.com```

![Create a new Maven Project - Step 3](https://apidocs.io/illustration/go?step=createNewProject3&projectName=easybimehlanding_lib)

Now create a new file inside ```src/test.com```

![Create a new Maven Project - Step 4](https://apidocs.io/illustration/go?step=createNewProject4&projectName=easybimehlanding_lib)

Name it ```testsdk.go```

![Create a new Maven Project - Step 5](https://apidocs.io/illustration/go?step=createNewProject5&projectName=easybimehlanding_lib)

In this Go file, you can start adding code to initialize the client library. Sample code to initialize the client library and using its methods is given in the subsequent sections.

### 2. Configure the Test Project

You need to import your generated library in this project in order to make use of its functions. In order to import the library, you can add its path in the ```GOPATH``` for this project. Follow the below steps:

Right click on the project name and click on ```Properties```

![Adding dependency to the client library - Step 1](https://apidocs.io/illustration/go?step=testProject0&projectName=easybimehlanding_lib)

Choose ```Go Compiler``` from the side menu. Check ```Use project specific settings``` and uncheck ```Use same value as the GOPATH environment variable.```. By default, the GOPATH value from the environment variables will be visible in ```Eclipse GOPATH```. Do not remove this as this points to the unirest dependency.

![Adding dependency to the client library - Step 2](https://apidocs.io/illustration/go?step=testProject1)

Append the library path to this GOPATH

![Adding dependency to the client library - Step 3](https://apidocs.io/illustration/go?step=testProject2&workspaceFolder=Easy%20Bimeh%20Landing-GoLang)

Once the path is appended, click on ```OK```

### 3. Build the Test Project

Right click on the project name and click on ```Build Project```

![Build Project](https://apidocs.io/illustration/go?step=buildProject&projectName=easybimehlanding_lib)

### 4. Run the Test Project

If the build is successful, right click on your Go file and click on ```Run As``` -> ```Go Application```

![Run Project](https://apidocs.io/illustration/go?step=runProject&projectName=easybimehlanding_lib)

# Class Reference

## <a name="list_of_controllers"></a>List of Controllers

* [filemanager_pkg](#filemanager_pkg)
* [liabilitydoctorinsurance_pkg](#liabilitydoctorinsurance_pkg)
* [carbody_pkg](#carbody_pkg)
* [electronicequipmentinsurance_pkg](#electronicequipmentinsurance_pkg)
* [otherinsurancetypes_pkg](#otherinsurancetypes_pkg)
* [thirdpartyinsurance_pkg](#thirdpartyinsurance_pkg)
* [motorcycleinsurance_pkg](#motorcycleinsurance_pkg)
* [fireinsurance_pkg](#fireinsurance_pkg)
* [earthquakeinsurance_pkg](#earthquakeinsurance_pkg)
* [travelinsurance_pkg](#travelinsurance_pkg)
* [elevatorinsurance_pkg](#elevatorinsurance_pkg)
* [main_pkg](#main_pkg)
* [combodata_pkg](#combodata_pkg)
* [trackingdamage_pkg](#trackingdamage_pkg)
* [footer_pkg](#footer_pkg)
* [insurancepolicyplan_pkg](#insurancepolicyplan_pkg)

## <a name="filemanager_pkg"></a>![Class: ](https://apidocs.io/img/class.png ".filemanager_pkg") filemanager_pkg

### Get instance

Factory for the ``` FILEMANAGER ``` interface can be accessed from the package filemanager_pkg.

```go
fileManager := filemanager_pkg.NewFILEMANAGER()
```

### <a name="upload"></a>![Method: ](https://apidocs.io/img/method.png ".filemanager_pkg.Upload") Upload

> آپلود فایل در ایزی بیمه
> بعد از آپلود، ادرس فایل باید در api های بعدی ارسال شود.


```go
func (me *FILEMANAGER_IMPL) Upload(
            subDomain string,
            xApiKey string,
            file string)(*models_pkg.BaseModelUpload,error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| subDomain |  ``` Required ```  | دامنه یا زیر دامنه ی مرکز بیمه |
| xApiKey |  ``` Required ```  | کلید اختصاصی ارتباط با سرور |
| file |  ``` Required ```  | فایل ارسالی |


#### Example Usage

```go
subDomain := "hfz1"
xApiKey := "d6dfd932-75d8-e911-811a-000c294ecf01"
file := "file"

var result *models_pkg.BaseModelUpload
result,_ = fileManager.Upload(subDomain, xApiKey, file)

```


[Back to List of Controllers](#list_of_controllers)

## <a name="liabilitydoctorinsurance_pkg"></a>![Class: ](https://apidocs.io/img/class.png ".liabilitydoctorinsurance_pkg") liabilitydoctorinsurance_pkg

### Get instance

Factory for the ``` LIABILITYDOCTORINSURANCE ``` interface can be accessed from the package liabilitydoctorinsurance_pkg.

```go
liabilityDoctorInsurance := liabilitydoctorinsurance_pkg.NewLIABILITYDOCTORINSURANCE()
```

### <a name="get_liability_doctor_insurance"></a>![Method: ](https://apidocs.io/img/method.png ".liabilitydoctorinsurance_pkg.GetLiabilityDoctorInsurance") GetLiabilityDoctorInsurance

> در یافت اطلاعات اولیه برای استعلام بیمه مسئولیت پزشکان


```go
func (me *LIABILITYDOCTORINSURANCE_IMPL) GetLiabilityDoctorInsurance(
            subDomain string,
            xApiKey string)(*models_pkg.BaseModelLiabilityDoctorInsurance,error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| subDomain |  ``` Required ```  | دامنه یا زیر دامنه ی مرکز بیمه |
| xApiKey |  ``` Required ```  | کلید اختصاصی ارتباط با سرور |


#### Example Usage

```go
subDomain := "hfz1"
xApiKey := "d6dfd932-75d8-e911-811a-000c294ecf01"

var result *models_pkg.BaseModelLiabilityDoctorInsurance
result,_ = liabilityDoctorInsurance.GetLiabilityDoctorInsurance(subDomain, xApiKey)

```


### <a name="get_medical_specialties"></a>![Method: ](https://apidocs.io/img/method.png ".liabilitydoctorinsurance_pkg.GetMedicalSpecialties") GetMedicalSpecialties

> دریافت لیست تخصص های پزشکی


```go
func (me *LIABILITYDOCTORINSURANCE_IMPL) GetMedicalSpecialties(
            id string,
            xApiKey string)(*models_pkg.BaseModelMedicalSpecialties,error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| id |  ``` Required ```  | نوع تخصص => ParamedicalExpertise => پیراپزشکی MedicalExpertise => پزشکی |
| xApiKey |  ``` Required ```  | کلید اختصاصی ارتباط با سرور |


#### Example Usage

```go
id := "ParamedicalExpertise"
xApiKey := "d6dfd932-75d8-e911-811a-000c294ecf01"

var result *models_pkg.BaseModelMedicalSpecialties
result,_ = liabilityDoctorInsurance.GetMedicalSpecialties(id, xApiKey)

```


[Back to List of Controllers](#list_of_controllers)

## <a name="carbody_pkg"></a>![Class: ](https://apidocs.io/img/class.png ".carbody_pkg") carbody_pkg

### Get instance

Factory for the ``` CARBODY ``` interface can be accessed from the package carbody_pkg.

```go
carBody := carbody_pkg.NewCARBODY()
```

### <a name="get_car_brand"></a>![Method: ](https://apidocs.io/img/method.png ".carbody_pkg.GetCarBrand") GetCarBrand

> دریافت برند خودرو


```go
func (me *CARBODY_IMPL) GetCarBrand(xApiKey string)(*models_pkg.CarBrand,error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| xApiKey |  ``` Required ```  | شناسه ی اختصاصی ارتباط با سرور |


#### Example Usage

```go
xApiKey := "d6dfd932-75d8-e911-811a-000c294ecf01"

var result *models_pkg.CarBrand
result,_ = carBody.GetCarBrand(xApiKey)

```


### <a name="get_car_brand_tips"></a>![Method: ](https://apidocs.io/img/method.png ".carbody_pkg.GetCarBrandTips") GetCarBrandTips

> دریافت لیست تیپ خودرو


```go
func (me *CARBODY_IMPL) GetCarBrandTips(
            carBrandId int64,
            xApiKey string)(*models_pkg.CarBrandTips,error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| carBrandId |  ``` Required ```  | شناسه ی برند خودرو |
| xApiKey |  ``` Required ```  | کلید اختصاصی ارتباط با سرور |


#### Example Usage

```go
carBrandId,_ := strconv.ParseInt("190", 10, 8)
xApiKey := "d6dfd932-75d8-e911-811a-000c294ecf01"

var result *models_pkg.CarBrandTips
result,_ = carBody.GetCarBrandTips(carBrandId, xApiKey)

```


### <a name="get_has_plan"></a>![Method: ](https://apidocs.io/img/method.png ".carbody_pkg.GetHasPlan") GetHasPlan

> آیا این نوع بیمه نامه، طرح بیمه ای دارد؟


```go
func (me *CARBODY_IMPL) GetHasPlan(
            subDomain string,
            insurancePolicyType int64,
            xApiKey string)(*models_pkg.HasPlan,error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| subDomain |  ``` Required ```  | دامنه یا زیر دامنه ی مرکز بیمه |
| insurancePolicyType |  ``` Required ```  | شناسه ی نوع بیمه نامه => بیمه بدنه=2 |
| xApiKey |  ``` Required ```  | کلید اختصاصی ارتباط با سرور |


#### Example Usage

```go
subDomain := "hfz1"
insurancePolicyType,_ := strconv.ParseInt("2", 10, 8)
xApiKey := "d6dfd932-75d8-e911-811a-000c294ecf01"

var result *models_pkg.HasPlan
result,_ = carBody.GetHasPlan(subDomain, insurancePolicyType, xApiKey)

```


[Back to List of Controllers](#list_of_controllers)

## <a name="electronicequipmentinsurance_pkg"></a>![Class: ](https://apidocs.io/img/class.png ".electronicequipmentinsurance_pkg") electronicequipmentinsurance_pkg

### Get instance

Factory for the ``` ELECTRONICEQUIPMENTINSURANCE ``` interface can be accessed from the package electronicequipmentinsurance_pkg.

```go
electronicEquipmentInsurance := electronicequipmentinsurance_pkg.NewELECTRONICEQUIPMENTINSURANCE()
```

### <a name="get_electronic_equipment_insurance"></a>![Method: ](https://apidocs.io/img/method.png ".electronicequipmentinsurance_pkg.GetElectronicEquipmentInsurance") GetElectronicEquipmentInsurance

> دریافت اطلاعات اولیه استعلام بیمه نامه ی تجهیزات الکترونیک


```go
func (me *ELECTRONICEQUIPMENTINSURANCE_IMPL) GetElectronicEquipmentInsurance(
            subDomain string,
            xApiKey string)(*models_pkg.BaseModelElectronicEquipmentInsurance,error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| subDomain |  ``` Required ```  | دامنه یا زیر دامنه ی مرکز بیمه |
| xApiKey |  ``` Required ```  | کلید اختصاصی ارتباط با سرور |


#### Example Usage

```go
subDomain := "subDomain"
xApiKey := "x-api-key"

var result *models_pkg.BaseModelElectronicEquipmentInsurance
result,_ = electronicEquipmentInsurance.GetElectronicEquipmentInsurance(subDomain, xApiKey)

```


### <a name="get_device_brand_types"></a>![Method: ](https://apidocs.io/img/method.png ".electronicequipmentinsurance_pkg.GetDeviceBrandTypes") GetDeviceBrandTypes

> دریافت لیست نوع برند دستگاه


```go
func (me *ELECTRONICEQUIPMENTINSURANCE_IMPL) GetDeviceBrandTypes(
            deviceGroup int64,
            deviceTypeId int64,
            xApiKey string)(*models_pkg.BaseModelDeviceBrandTypes,error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| deviceGroup |  ``` Required ```  | شناسه ی گروه دستگاه |
| deviceTypeId |  ``` Required ```  | شناسه ی نوع دستگاه |
| xApiKey |  ``` Required ```  | کلید اختصاصی ارتباط با سرور |


#### Example Usage

```go
deviceGroup,_ := strconv.ParseInt("1", 10, 8)
deviceTypeId,_ := strconv.ParseInt("1", 10, 8)
xApiKey := "d6dfd932-75d8-e911-811a-000c294ecf01"

var result *models_pkg.BaseModelDeviceBrandTypes
result,_ = electronicEquipmentInsurance.GetDeviceBrandTypes(deviceGroup, deviceTypeId, xApiKey)

```


### <a name="get_divice_franchisee"></a>![Method: ](https://apidocs.io/img/method.png ".electronicequipmentinsurance_pkg.GetDiviceFranchisee") GetDiviceFranchisee

> دریافت لیست فرانشیر استعلام بیمه نامه ی تجهیزات الکترونیک


```go
func (me *ELECTRONICEQUIPMENTINSURANCE_IMPL) GetDiviceFranchisee(
            deviceModelId int64,
            xApiKey string)(*models_pkg.BaseModelDiviceFranchisee,error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| deviceModelId |  ``` Required ```  | شناسه ی مدل دستگاه |
| xApiKey |  ``` Required ```  | کلید اختصاصی ارتباط با سرور |


#### Example Usage

```go
deviceModelId,_ := strconv.ParseInt("1340", 10, 8)
xApiKey := "d6dfd932-75d8-e911-811a-000c294ecf01"

var result *models_pkg.BaseModelDiviceFranchisee
result,_ = electronicEquipmentInsurance.GetDiviceFranchisee(deviceModelId, xApiKey)

```


[Back to List of Controllers](#list_of_controllers)

## <a name="otherinsurancetypes_pkg"></a>![Class: ](https://apidocs.io/img/class.png ".otherinsurancetypes_pkg") otherinsurancetypes_pkg

### Get instance

Factory for the ``` OTHERINSURANCETYPES ``` interface can be accessed from the package otherinsurancetypes_pkg.

```go
otherInsuranceTypes := otherinsurancetypes_pkg.NewOTHERINSURANCETYPES()
```

### <a name="get_other_insurance_types"></a>![Method: ](https://apidocs.io/img/method.png ".otherinsurancetypes_pkg.GetOtherInsuranceTypes") GetOtherInsuranceTypes

> دریافت لیست سایر بیمه نامه ها


```go
func (me *OTHERINSURANCETYPES_IMPL) GetOtherInsuranceTypes(
            subDomain string,
            xApiKey string)(*models_pkg.OtherInsuranceTypes,error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| subDomain |  ``` Required ```  | دامنه یا زیر دامنه ی مرکز بیمه |
| xApiKey |  ``` Required ```  | کلید اختصاصی ارتباط با سرور |


#### Example Usage

```go
subDomain := "hfz1"
xApiKey := "d6dfd932-75d8-e911-811a-000c294ecf01"

var result *models_pkg.OtherInsuranceTypes
result,_ = otherInsuranceTypes.GetOtherInsuranceTypes(subDomain, xApiKey)

```


### <a name="get_send_sms_token"></a>![Method: ](https://apidocs.io/img/method.png ".otherinsurancetypes_pkg.GetSendSmsToken") GetSendSmsToken

> ارسال توکن تایید شماره تماس، برای احراز هویت کاربر


```go
func (me *OTHERINSURANCETYPES_IMPL) GetSendSmsToken(
            mobile string,
            insuranceCentreSubDomain string,
            xApiKey string)(*models_pkg.SendSmsToken,error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| mobile |  ``` Required ```  | شماره موبایل |
| insuranceCentreSubDomain |  ``` Required ```  | دامنه یا زیردامنه ی مرکز بیمه |
| xApiKey |  ``` Required ```  | کلید اختصاصی ارتباط با سرور |


#### Example Usage

```go
mobile := "09018318086"
insuranceCentreSubDomain := "hfz1"
xApiKey := "d6dfd932-75d8-e911-811a-000c294ecf01"

var result *models_pkg.SendSmsToken
result,_ = otherInsuranceTypes.GetSendSmsToken(mobile, insuranceCentreSubDomain, xApiKey)

```


### <a name="get_verify_sms_token"></a>![Method: ](https://apidocs.io/img/method.png ".otherinsurancetypes_pkg.GetVerifySmsToken") GetVerifySmsToken

> تایید توکن پیامک شده به کاربر، برای احراز هویت


```go
func (me *OTHERINSURANCETYPES_IMPL) GetVerifySmsToken(
            mobile string,
            token int64,
            insuranceCentreSubDomain string,
            aliasName string,
            resource string,
            xApiKey string)(*models_pkg.Status200,error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| mobile |  ``` Required ```  | شماره موبایل |
| token |  ``` Required ```  | توکن دریافتی کاربر از پیامک |
| insuranceCentreSubDomain |  ``` Required ```  | دامنه یا زیر دامنه ی اختصاصی مرکز بیمه |
| aliasName |  ``` Required ```  | نام و نام خانوادگی کاربر |
| resource |  ``` Required ```  | دامنه ی درخواست دهنده |
| xApiKey |  ``` Required ```  | کلید اختصاصی ارتباط با سرور |


#### Example Usage

```go
mobile := "09018318086"
token,_ := strconv.ParseInt("27763", 10, 8)
insuranceCentreSubDomain := "hfz1"
aliasName := "علی موسوی"
resource := "https://hfz1.easybimeh.com"
xApiKey := "d6dfd932-75d8-e911-811a-000c294ecf01"

var result *models_pkg.Status200
result,_ = otherInsuranceTypes.GetVerifySmsToken(mobile, token, insuranceCentreSubDomain, aliasName, resource, xApiKey)

```

#### Errors
 
| Error Code | Error Description |
|------------|-------------------|
| 400 | Bad Request |



[Back to List of Controllers](#list_of_controllers)

## <a name="thirdpartyinsurance_pkg"></a>![Class: ](https://apidocs.io/img/class.png ".thirdpartyinsurance_pkg") thirdpartyinsurance_pkg

### Get instance

Factory for the ``` THIRDPARTYINSURANCE ``` interface can be accessed from the package thirdpartyinsurance_pkg.

```go
thirdPartyInsurance := thirdpartyinsurance_pkg.NewTHIRDPARTYINSURANCE()
```

### <a name="get_car_brands"></a>![Method: ](https://apidocs.io/img/method.png ".thirdpartyinsurance_pkg.GetCarBrands") GetCarBrands

> دریافت لیست برند خودرو ها


```go
func (me *THIRDPARTYINSURANCE_IMPL) GetCarBrands(
            carTypeGroup int64,
            xApiKey string)(*models_pkg.CarBrands,error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| carTypeGroup |  ``` Required ```  | شناسه ی گروه خودرو |
| xApiKey |  ``` Required ```  | کلید اختصاصی ارتباط با سرور |


#### Example Usage

```go
carTypeGroup,_ := strconv.ParseInt("1", 10, 8)
xApiKey := "d6dfd932-75d8-e911-811a-000c294ecf01"

var result *models_pkg.CarBrands
result,_ = thirdPartyInsurance.GetCarBrands(carTypeGroup, xApiKey)

```


### <a name="get_risk_level"></a>![Method: ](https://apidocs.io/img/method.png ".thirdpartyinsurance_pkg.GetRiskLevel") GetRiskLevel

> دریافت لیست تخفیف های بیمه


```go
func (me *THIRDPARTYINSURANCE_IMPL) GetRiskLevel(
            insurancePolicyType int64,
            xApiKey string)(*models_pkg.RiskLevel,error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| insurancePolicyType |  ``` Required ```  | شناسه ی نوع بیمه نامه |
| xApiKey |  ``` Required ```  | کلید اختصاصی ارتباط با سرور |


#### Example Usage

```go
insurancePolicyType,_ := strconv.ParseInt("0", 10, 8)
xApiKey := "d6dfd932-75d8-e911-811a-000c294ecf01"

var result *models_pkg.RiskLevel
result,_ = thirdPartyInsurance.GetRiskLevel(insurancePolicyType, xApiKey)

```


### <a name="get_car_brand_tips"></a>![Method: ](https://apidocs.io/img/method.png ".thirdpartyinsurance_pkg.GetCarBrandTips") GetCarBrandTips

> دریافت لیست تیپ خودرو


```go
func (me *THIRDPARTYINSURANCE_IMPL) GetCarBrandTips(
            carBrandId int64,
            xApiKey string)(*models_pkg.CarBrandTips,error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| carBrandId |  ``` Required ```  | شناسه ی برند خودرو |
| xApiKey |  ``` Required ```  | کلید اختصاصی ارتباط با سرور |


#### Example Usage

```go
carBrandId,_ := strconv.ParseInt("190", 10, 8)
xApiKey := "d6dfd932-75d8-e911-811a-000c294ecf01"

var result *models_pkg.CarBrandTips
result,_ = thirdPartyInsurance.GetCarBrandTips(carBrandId, xApiKey)

```


### <a name="get_car_uses"></a>![Method: ](https://apidocs.io/img/method.png ".thirdpartyinsurance_pkg.GetCarUses") GetCarUses

> دریافت لیست نوع کاربری خودرو


```go
func (me *THIRDPARTYINSURANCE_IMPL) GetCarUses(
            carTypeId int64,
            xApiKey string)(*models_pkg.CarUses,error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| carTypeId |  ``` Required ```  | شناسه ی نوع خودرو |
| xApiKey |  ``` Required ```  | کلید اختصاصی ارتباط با سرور |


#### Example Usage

```go
carTypeId,_ := strconv.ParseInt("103", 10, 8)
xApiKey := "d6dfd932-75d8-e911-811a-000c294ecf01"

var result *models_pkg.CarUses
result,_ = thirdPartyInsurance.GetCarUses(carTypeId, xApiKey)

```


### <a name="get_has_plan"></a>![Method: ](https://apidocs.io/img/method.png ".thirdpartyinsurance_pkg.GetHasPlan") GetHasPlan

> آیا این نوع بیمه نامه، طرح بیمه ای دارد؟


```go
func (me *THIRDPARTYINSURANCE_IMPL) GetHasPlan(
            subDomain string,
            insurancePolicyType int64,
            xApiKey string)(*models_pkg.HasPlan,error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| subDomain |  ``` Required ```  | دامنه یا زیر دامنه ی مرکز بیمه |
| insurancePolicyType |  ``` Required ```  | شناسه ی نوع بیمه |
| xApiKey |  ``` Required ```  | کلید اختصاصی ارتباط با سرور |


#### Example Usage

```go
subDomain := "hfz1"
insurancePolicyType,_ := strconv.ParseInt("0", 10, 8)
xApiKey := "d6dfd932-75d8-e911-811a-000c294ecf01"

var result *models_pkg.HasPlan
result,_ = thirdPartyInsurance.GetHasPlan(subDomain, insurancePolicyType, xApiKey)

```


[Back to List of Controllers](#list_of_controllers)

## <a name="motorcycleinsurance_pkg"></a>![Class: ](https://apidocs.io/img/class.png ".motorcycleinsurance_pkg") motorcycleinsurance_pkg

### Get instance

Factory for the ``` MOTORCYCLEINSURANCE ``` interface can be accessed from the package motorcycleinsurance_pkg.

```go
motorcycleInsurance := motorcycleinsurance_pkg.NewMOTORCYCLEINSURANCE()
```

### <a name="get_car_brands"></a>![Method: ](https://apidocs.io/img/method.png ".motorcycleinsurance_pkg.GetCarBrands") GetCarBrands

> دریافت لیست برند موتور سیکلت


```go
func (me *MOTORCYCLEINSURANCE_IMPL) GetCarBrands(
            carTypeGroup int64,
            xApiKey string)(*models_pkg.CarBrands,error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| carTypeGroup |  ``` Required ```  | شناسه ی گروه خودرویی، موتور سیکلت =>0 |
| xApiKey |  ``` Required ```  | کلید اختصاصی ارتباط با سرور |


#### Example Usage

```go
carTypeGroup,_ := strconv.ParseInt("0", 10, 8)
xApiKey := "d6dfd932-75d8-e911-811a-000c294ecf01"

var result *models_pkg.CarBrands
result,_ = motorcycleInsurance.GetCarBrands(carTypeGroup, xApiKey)

```


### <a name="get_car_brand_tips"></a>![Method: ](https://apidocs.io/img/method.png ".motorcycleinsurance_pkg.GetCarBrandTips") GetCarBrandTips

> دریافت لیست تیپ موتور سیکلت


```go
func (me *MOTORCYCLEINSURANCE_IMPL) GetCarBrandTips(
            carBrandId int64,
            xApiKey string)(*models_pkg.CarBrandTips,error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| carBrandId |  ``` Required ```  | شناسه ی برند موتور سیکلت |
| xApiKey |  ``` Required ```  | کلید اختصاصی ارتباط با سرور |


#### Example Usage

```go
carBrandId,_ := strconv.ParseInt("472", 10, 8)
xApiKey := "d6dfd932-75d8-e911-811a-000c294ecf01"

var result *models_pkg.CarBrandTips
result,_ = motorcycleInsurance.GetCarBrandTips(carBrandId, xApiKey)

```


### <a name="get_has_plan"></a>![Method: ](https://apidocs.io/img/method.png ".motorcycleinsurance_pkg.GetHasPlan") GetHasPlan

> آیا این نوع بیمه نامه، طرح بیمه ای دارد؟


```go
func (me *MOTORCYCLEINSURANCE_IMPL) GetHasPlan(
            subDomain string,
            insurancePolicyType int64,
            xApiKey string)(*models_pkg.HasPlan,error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| subDomain |  ``` Required ```  | دامنه یا زیر دامنه ی مرکز بیمه |
| insurancePolicyType |  ``` Required ```  | شناسه ی نوع بیمه نامه |
| xApiKey |  ``` Required ```  | کلید اختصاصی ارتباط با سرور |


#### Example Usage

```go
subDomain := "hfz1"
insurancePolicyType,_ := strconv.ParseInt("7", 10, 8)
xApiKey := "d6dfd932-75d8-e911-811a-000c294ecf01"

var result *models_pkg.HasPlan
result,_ = motorcycleInsurance.GetHasPlan(subDomain, insurancePolicyType, xApiKey)

```


[Back to List of Controllers](#list_of_controllers)

## <a name="fireinsurance_pkg"></a>![Class: ](https://apidocs.io/img/class.png ".fireinsurance_pkg") fireinsurance_pkg

### Get instance

Factory for the ``` FIREINSURANCE ``` interface can be accessed from the package fireinsurance_pkg.

```go
fireInsurance := fireinsurance_pkg.NewFIREINSURANCE()
```

### <a name="get_fire_insurance"></a>![Method: ](https://apidocs.io/img/method.png ".fireinsurance_pkg.GetFireInsurance") GetFireInsurance

> دریافت اطلاعات پایه بیمه ی آتش سوزی


```go
func (me *FIREINSURANCE_IMPL) GetFireInsurance(
            subDomain string,
            insurancePolicyId int64,
            xApiKey string)(*models_pkg.BaseModelFireInsurance,error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| subDomain |  ``` Required ```  | دامنه یا زیر دامنه ی مرکز بیمه |
| insurancePolicyId |  ``` Required ```  | شناسه ی بیمه نامه |
| xApiKey |  ``` Required ```  | کلید اختصاصی ارتباط با سرور |


#### Example Usage

```go
subDomain := "hfz1"
insurancePolicyId,_ := strconv.ParseInt("0", 10, 8)
xApiKey := "d6dfd932-75d8-e911-811a-000c294ecf01"

var result *models_pkg.BaseModelFireInsurance
result,_ = fireInsurance.GetFireInsurance(subDomain, insurancePolicyId, xApiKey)

```


[Back to List of Controllers](#list_of_controllers)

## <a name="earthquakeinsurance_pkg"></a>![Class: ](https://apidocs.io/img/class.png ".earthquakeinsurance_pkg") earthquakeinsurance_pkg

### Get instance

Factory for the ``` EARTHQUAKEINSURANCE ``` interface can be accessed from the package earthquakeinsurance_pkg.

```go
earthquakeInsurance := earthquakeinsurance_pkg.NewEARTHQUAKEINSURANCE()
```

### <a name="get_earthquake"></a>![Method: ](https://apidocs.io/img/method.png ".earthquakeinsurance_pkg.GetEarthquake") GetEarthquake

> دریافت اطلاعات پایه ی بیمه ی زلزله


```go
func (me *EARTHQUAKEINSURANCE_IMPL) GetEarthquake(
            subDomain string,
            insurancePolicyId int64,
            insurancePolicyType int64,
            xApiKey string)(*models_pkg.BaseModelEarthquake,error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| subDomain |  ``` Required ```  | دامنه یا زیر دامنه ی مرکز بیمه |
| insurancePolicyId |  ``` Required ```  | شناسه ی بیمه نامه |
| insurancePolicyType |  ``` Required ```  | شناسه ی نوع بیمه نامه |
| xApiKey |  ``` Required ```  | کلید اختصاصی ارتباط با سرور |


#### Example Usage

```go
subDomain := "hfz1"
insurancePolicyId,_ := strconv.ParseInt("0", 10, 8)
insurancePolicyType,_ := strconv.ParseInt("6", 10, 8)
xApiKey := "d6dfd932-75d8-e911-811a-000c294ecf01"

var result *models_pkg.BaseModelEarthquake
result,_ = earthquakeInsurance.GetEarthquake(subDomain, insurancePolicyId, insurancePolicyType, xApiKey)

```


[Back to List of Controllers](#list_of_controllers)

## <a name="travelinsurance_pkg"></a>![Class: ](https://apidocs.io/img/class.png ".travelinsurance_pkg") travelinsurance_pkg

### Get instance

Factory for the ``` TRAVELINSURANCE ``` interface can be accessed from the package travelinsurance_pkg.

```go
travelInsurance := travelinsurance_pkg.NewTRAVELINSURANCE()
```

### <a name="get_travel_insurance"></a>![Method: ](https://apidocs.io/img/method.png ".travelinsurance_pkg.GetTravelInsurance") GetTravelInsurance

> TODO: Add Description


```go
func (me *TRAVELINSURANCE_IMPL) GetTravelInsurance(
            subDomain string,
            insurancePolicyId int64,
            xApiKey string)(*models_pkg.BaseModelTravelInsurance,error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| subDomain |  ``` Required ```  | دامنه یا زیر دامنه ی مرکز بیمه |
| insurancePolicyId |  ``` Required ```  | شناسه ی بیمه نامه |
| xApiKey |  ``` Required ```  | کلید اختصاصی ارتباط با سرور |


#### Example Usage

```go
subDomain := "hfz1"
insurancePolicyId,_ := strconv.ParseInt("0", 10, 8)
xApiKey := "d6dfd932-75d8-e911-811a-000c294ecf01"

var result *models_pkg.BaseModelTravelInsurance
result,_ = travelInsurance.GetTravelInsurance(subDomain, insurancePolicyId, xApiKey)

```


[Back to List of Controllers](#list_of_controllers)

## <a name="elevatorinsurance_pkg"></a>![Class: ](https://apidocs.io/img/class.png ".elevatorinsurance_pkg") elevatorinsurance_pkg

### Get instance

Factory for the ``` ELEVATORINSURANCE ``` interface can be accessed from the package elevatorinsurance_pkg.

```go
elevatorInsurance := elevatorinsurance_pkg.NewELEVATORINSURANCE()
```

### <a name="get_elevator_insurance"></a>![Method: ](https://apidocs.io/img/method.png ".elevatorinsurance_pkg.GetElevatorInsurance") GetElevatorInsurance

> دریافت اطلاعات پایه ی بیمه نامه ی آسانسور


```go
func (me *ELEVATORINSURANCE_IMPL) GetElevatorInsurance(
            subDomain string,
            insurancePolicyId int64,
            xApiKey string)(*models_pkg.BaseModelElevatorInsurance,error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| subDomain |  ``` Required ```  | دامنه یا زیر دامنه ی مرکز بیمه |
| insurancePolicyId |  ``` Required ```  | شناسه ی بیمه نامه |
| xApiKey |  ``` Required ```  | کلید اختصاصی ارتباط با سرور |


#### Example Usage

```go
subDomain := "hfz1"
insurancePolicyId,_ := strconv.ParseInt("0", 10, 8)
xApiKey := "d6dfd932-75d8-e911-811a-000c294ecf01"

var result *models_pkg.BaseModelElevatorInsurance
result,_ = elevatorInsurance.GetElevatorInsurance(subDomain, insurancePolicyId, xApiKey)

```


[Back to List of Controllers](#list_of_controllers)

## <a name="main_pkg"></a>![Class: ](https://apidocs.io/img/class.png ".main_pkg") main_pkg

### Get instance

Factory for the ``` MAIN ``` interface can be accessed from the package main_pkg.

```go
main := main_pkg.NewMAIN()
```

### <a name="get_portal_landing_page"></a>![Method: ](https://apidocs.io/img/method.png ".main_pkg.GetPortalLandingPage") GetPortalLandingPage

> در یافت اطلاعات لندینگ مراکز بیمه


```go
func (me *MAIN_IMPL) GetPortalLandingPage(
            id string,
            xApiKey string)(*models_pkg.BaseModelPortalLandingPage,error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| id |  ``` Required ```  | دامنه یا زیر دامنه ی مرکز بیمه |
| xApiKey |  ``` Required ```  | کلید اختصاصی ارتباط با سرور |


#### Example Usage

```go
id := "hfz1"
xApiKey := "d6dfd932-75d8-e911-811a-000c294ecf01"

var result *models_pkg.BaseModelPortalLandingPage
result,_ = main.GetPortalLandingPage(id, xApiKey)

```


### <a name="get_insurance_centre_policy_types"></a>![Method: ](https://apidocs.io/img/method.png ".main_pkg.GetInsuranceCentrePolicyTypes") GetInsuranceCentrePolicyTypes

> دریافت لیست بیمه ی های ارائه شده توسط مرکز بیمه


```go
func (me *MAIN_IMPL) GetInsuranceCentrePolicyTypes(
            subDomain string,
            xApiKey string)(*models_pkg.BaseModelInsuranceCentrePolicyTypes,error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| subDomain |  ``` Required ```  | دامنه یا زیر دامنه ی مرکز بیمه |
| xApiKey |  ``` Required ```  | کلید اختصاصی ارتباط با سرور |


#### Example Usage

```go
subDomain := "hfz1"
xApiKey := "d6dfd932-75d8-e911-811a-000c294ecf01"

var result *models_pkg.BaseModelInsuranceCentrePolicyTypes
result,_ = main.GetInsuranceCentrePolicyTypes(subDomain, xApiKey)

```


[Back to List of Controllers](#list_of_controllers)

## <a name="combodata_pkg"></a>![Class: ](https://apidocs.io/img/class.png ".combodata_pkg") combodata_pkg

### Get instance

Factory for the ``` COMBODATA ``` interface can be accessed from the package combodata_pkg.

```go
comboData := combodata_pkg.NewCOMBODATA()
```

### <a name="get_damage_type"></a>![Method: ](https://apidocs.io/img/method.png ".combodata_pkg.GetDamageType") GetDamageType

> دریافت لیست نوع خسارت


```go
func (me *COMBODATA_IMPL) GetDamageType(xApiKey string)(*models_pkg.BaseModelDamageType,error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| xApiKey |  ``` Required ```  | کلید اختصاصی ارتباط با سرور |


#### Example Usage

```go
xApiKey := "d6dfd932-75d8-e911-811a-000c294ecf01"

var result *models_pkg.BaseModelDamageType
result,_ = comboData.GetDamageType(xApiKey)

```


### <a name="get_insurance_types"></a>![Method: ](https://apidocs.io/img/method.png ".combodata_pkg.GetInsuranceTypes") GetInsuranceTypes

> دریافت لیست نوع بیمه نامه


```go
func (me *COMBODATA_IMPL) GetInsuranceTypes(
            subDomain string,
            issueInsurance bool,
            xApiKey string)(*models_pkg.InsuranceTypes,error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| subDomain |  ``` Required ```  | دامنه یا زیر دامنه ی مرکز بیمه |
| issueInsurance |  ``` Required ```  | دریافت بیمه نامه های قابل صدور |
| xApiKey |  ``` Required ```  | کلید اختصاصی ارتباط با سرور |


#### Example Usage

```go
subDomain := "hfz1"
issueInsurance := false
xApiKey := "d6dfd932-75d8-e911-811a-000c294ecf01"

var result *models_pkg.InsuranceTypes
result,_ = comboData.GetInsuranceTypes(subDomain, issueInsurance, xApiKey)

```


### <a name="get_insurance_companies"></a>![Method: ](https://apidocs.io/img/method.png ".combodata_pkg.GetInsuranceCompanies") GetInsuranceCompanies

> دریافت لیست شرکت های بیمه


```go
func (me *COMBODATA_IMPL) GetInsuranceCompanies(
            subDomain string,
            insuranceTypeId int64,
            xApiKey string)(*models_pkg.InsuranceCompanies,error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| subDomain |  ``` Required ```  | دامنه یا زیر دامنه ی مرکز بیمه |
| insuranceTypeId |  ``` Required ```  | شناسه ی نوع بیمه نامه |
| xApiKey |  ``` Required ```  | کلید اختصاصی ارتباط با سرور |


#### Example Usage

```go
subDomain := "hfz1"
insuranceTypeId,_ := strconv.ParseInt("1", 10, 8)
xApiKey := "d6dfd932-75d8-e911-811a-000c294ecf01"

var result *models_pkg.InsuranceCompanies
result,_ = comboData.GetInsuranceCompanies(subDomain, insuranceTypeId, xApiKey)

```


[Back to List of Controllers](#list_of_controllers)

## <a name="trackingdamage_pkg"></a>![Class: ](https://apidocs.io/img/class.png ".trackingdamage_pkg") trackingdamage_pkg

### Get instance

Factory for the ``` TRACKINGDAMAGE ``` interface can be accessed from the package trackingdamage_pkg.

```go
trackingDamage := trackingdamage_pkg.NewTRACKINGDAMAGE()
```

### <a name="get_tracking_code"></a>![Method: ](https://apidocs.io/img/method.png ".trackingdamage_pkg.GetTrackingCode") GetTrackingCode

> استعلام وضعیت خسارت


```go
func (me *TRACKINGDAMAGE_IMPL) GetTrackingCode(
            mTrackingCode string,
            xApiKey string)(*models_pkg.BaseModelTrakingCode,error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| mTrackingCode |  ``` Required ```  | کد پیگیری خسارت |
| xApiKey |  ``` Required ```  | کلید اختصاصی ارتباط با سرور |


#### Example Usage

```go
mTrackingCode := "/{TrackingCode}"
xApiKey := "d6dfd932-75d8-e911-811a-000c294ecf01"

var result *models_pkg.BaseModelTrakingCode
result,_ = trackingDamage.GetTrackingCode(mTrackingCode, xApiKey)

```

#### Errors
 
| Error Code | Error Description |
|------------|-------------------|
| 400 | Bad Request |



### <a name="create_tracking_damage"></a>![Method: ](https://apidocs.io/img/method.png ".trackingdamage_pkg.CreateTrackingDamage") CreateTrackingDamage

> ثبت خسارت بیمه


```go
func (me *TRACKINGDAMAGE_IMPL) CreateTrackingDamage(
            body *models_pkg.TrackingDamageRequest,
            xApiKey string,
            contentType string)(*models_pkg.TrackingDamage,error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| body |  ``` Required ```  | اطلاعات خسارت |
| xApiKey |  ``` Required ```  | کلید اختصاصی ارتباط با سرور |
| contentType |  ``` Required ```  | TODO: Add a parameter description |


#### Example Usage

```go
bodyValue := []byte("{\r\n  \"personalityType\": 0,\r\n  \"trackingDamageStatus\": [\r\n    {\r\n      \"trackingDamageFile\": [\r\n        {\r\n          \"id\": 162747,\r\n          \"url\": \"https://media.easybimeh.com//Easybimeh/FileManager/InsuranceCentre/hfz1/637089119345134776.jpeg\",\r\n          \"title\": \"کارت شناسایی\"\r\n        }\r\n      ]\r\n    }\r\n  ],\r\n  \"description\": \"بدنه ی خودرو خسارت دیده\",\r\n  \"insuranceTypeId\": 1,\r\n  \"insuranceCompanyId\": 34,\r\n  \"insurancePolicyNumber\": \"123456\",\r\n  \"damageType\": \"مالی\",\r\n  \"name\": \"کاظم\",\r\n  \"nationalCode\": \"3080118383\",\r\n  \"mobile\": \"09018318086\",\r\n  \"insuredProfile\": \"پژو 405\",\r\n  \"subDomain\": \"hfz1\"\r\n}")
var body *models_pkg.TrackingDamageRequest
json.Unmarshal(bodyValue,&body)
xApiKey := "d6dfd932-75d8-e911-811a-000c294ecf01"
contentType := "application/json"

var result *models_pkg.TrackingDamage
result,_ = trackingDamage.CreateTrackingDamage(body, xApiKey, contentType)

```


### <a name="get_status_status_collections"></a>![Method: ](https://apidocs.io/img/method.png ".trackingdamage_pkg.GetStatusStatusCollections") GetStatusStatusCollections

> دریافت لیست وضعیت های خسارت


```go
func (me *TRACKINGDAMAGE_IMPL) GetStatusStatusCollections(
            statusType int64,
            xApiKey string)(*models_pkg.BaseModelStatusStatusCollections,error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| statusType |  ``` Required ```  | نوع وضعیت ها ی خسارت => 0 |
| xApiKey |  ``` Required ```  | کلید اختصاصی ارتباط با سرور |


#### Example Usage

```go
statusType,_ := strconv.ParseInt("0", 10, 8)
xApiKey := "d6dfd932-75d8-e911-811a-000c294ecf01"

var result *models_pkg.BaseModelStatusStatusCollections
result,_ = trackingDamage.GetStatusStatusCollections(statusType, xApiKey)

```


### <a name="get_status"></a>![Method: ](https://apidocs.io/img/method.png ".trackingdamage_pkg.GetStatus") GetStatus

> دریافت اطلاعات وضعیت


```go
func (me *TRACKINGDAMAGE_IMPL) GetStatus(
            entityId int64,
            xApiKey string)(*models_pkg.BaseModelStatus,error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| entityId |  ``` Required ```  | شناسه ی وضعیت |
| xApiKey |  ``` Required ```  | کلید اختصاصی ارتباط با سرور |


#### Example Usage

```go
entityId,_ := strconv.ParseInt("1129", 10, 8)
xApiKey := "d6dfd932-75d8-e911-811a-000c294ecf01"

var result *models_pkg.BaseModelStatus
result,_ = trackingDamage.GetStatus(entityId, xApiKey)

```


[Back to List of Controllers](#list_of_controllers)

## <a name="footer_pkg"></a>![Class: ](https://apidocs.io/img/class.png ".footer_pkg") footer_pkg

### Get instance

Factory for the ``` FOOTER ``` interface can be accessed from the package footer_pkg.

```go
footer := footer_pkg.NewFOOTER()
```

### <a name="get_portal_landing_contact_about"></a>![Method: ](https://apidocs.io/img/method.png ".footer_pkg.GetPortalLandingContactAbout") GetPortalLandingContactAbout

> دریافت اطلاعات درباره ی ما


```go
func (me *FOOTER_IMPL) GetPortalLandingContactAbout(xApiKey string)(*models_pkg.BaseModelPortalLandingContactAbout,error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| xApiKey |  ``` Required ```  | TODO: Add a parameter description |


#### Example Usage

```go
xApiKey := "d6dfd932-75d8-e911-811a-000c294ecf01"

var result *models_pkg.BaseModelPortalLandingContactAbout
result,_ = footer.GetPortalLandingContactAbout(xApiKey)

```


### <a name="get_faq_insurance_centre"></a>![Method: ](https://apidocs.io/img/method.png ".footer_pkg.GetFaqInsuranceCentre") GetFaqInsuranceCentre

> دریافت لیست سوالات متداول


```go
func (me *FOOTER_IMPL) GetFaqInsuranceCentre(xApiKey string)(*models_pkg.BaseModelFaqInsuranceCentre,error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| xApiKey |  ``` Required ```  | کلید اختصاصی ارتباط با سرور |


#### Example Usage

```go
xApiKey := "d6dfd932-75d8-e911-811a-000c294ecf01"

var result *models_pkg.BaseModelFaqInsuranceCentre
result,_ = footer.GetFaqInsuranceCentre(xApiKey)

```


### <a name="get_insurance_policy_tracking"></a>![Method: ](https://apidocs.io/img/method.png ".footer_pkg.GetInsurancePolicyTracking") GetInsurancePolicyTracking

> پیگیری وضعیت بیمه نامه


```go
func (me *FOOTER_IMPL) GetInsurancePolicyTracking(
            trackingCode int64,
            nationalCode int64,
            xApiKey string)(*models_pkg.BaseModelInsurancePolicyTracking,error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| trackingCode |  ``` Required ```  | شماره ی پیگیری بیمه نامه |
| nationalCode |  ``` Required ```  | کد ملی کاربر |
| xApiKey |  ``` Required ```  | کلید اختصاصی ارتباط با سرور |


#### Example Usage

```go
trackingCode,_ := strconv.ParseInt("213981083", 10, 8)
nationalCode,_ := strconv.ParseInt("3080115309", 10, 8)
xApiKey := "d6dfd932-75d8-e911-811a-000c294ecf01"

var result *models_pkg.BaseModelInsurancePolicyTracking
result,_ = footer.GetInsurancePolicyTracking(trackingCode, nationalCode, xApiKey)

```

#### Errors
 
| Error Code | Error Description |
|------------|-------------------|
| 500 | Internal Server Error |



[Back to List of Controllers](#list_of_controllers)

## <a name="insurancepolicyplan_pkg"></a>![Class: ](https://apidocs.io/img/class.png ".insurancepolicyplan_pkg") insurancepolicyplan_pkg

### Get instance

Factory for the ``` INSURANCEPOLICYPLAN ``` interface can be accessed from the package insurancepolicyplan_pkg.

```go
insurancePolicyPlan := insurancepolicyplan_pkg.NewINSURANCEPOLICYPLAN()
```

### <a name="get_special_plan"></a>![Method: ](https://apidocs.io/img/method.png ".insurancepolicyplan_pkg.GetSpecialPlan") GetSpecialPlan

> دریافت لیست طرح های بیمه ای


```go
func (me *INSURANCEPOLICYPLAN_IMPL) GetSpecialPlan(
            subDomain string,
            xApiKey string)(*models_pkg.BaseModelSpecialPlan,error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| subDomain |  ``` Required ```  | دامنه یا زیر دامنه ی مرکز بیمه |
| xApiKey |  ``` Required ```  | کلید اختصاصی ارتباط با سرور |


#### Example Usage

```go
subDomain := "hfz1"
xApiKey := "d6dfd932-75d8-e911-811a-000c294ecf01"

var result *models_pkg.BaseModelSpecialPlan
result,_ = insurancePolicyPlan.GetSpecialPlan(subDomain, xApiKey)

```


[Back to List of Controllers](#list_of_controllers)



