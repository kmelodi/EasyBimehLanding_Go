/*
 * easybimehlanding_lib
 *
 * This file was automatically generated by APIMATIC v2.0 ( https://apimatic.io ).
 */

package EasyBimehLandingClient

import(
	"easybimehlanding_lib/configuration_pkg"
	"easybimehlanding_lib/filemanager_pkg"
	"easybimehlanding_lib/liabilitydoctorinsurance_pkg"
	"easybimehlanding_lib/carbody_pkg"
	"easybimehlanding_lib/electronicequipmentinsurance_pkg"
	"easybimehlanding_lib/otherinsurancetypes_pkg"
	"easybimehlanding_lib/thirdpartyinsurance_pkg"
	"easybimehlanding_lib/motorcycleinsurance_pkg"
	"easybimehlanding_lib/fireinsurance_pkg"
	"easybimehlanding_lib/earthquakeinsurance_pkg"
	"easybimehlanding_lib/travelinsurance_pkg"
	"easybimehlanding_lib/elevatorinsurance_pkg"
	"easybimehlanding_lib/main_pkg"
	"easybimehlanding_lib/combodata_pkg"
	"easybimehlanding_lib/trackingdamage_pkg"
	"easybimehlanding_lib/footer_pkg"
	"easybimehlanding_lib/insurancepolicyplan_pkg"
)
/*
 * Client structure as interface implementation
 */
type EASYBIMEHLANDING_IMPL struct {
     filemanager filemanager_pkg.FILEMANAGER
     liabilitydoctorinsurance liabilitydoctorinsurance_pkg.LIABILITYDOCTORINSURANCE
     carbody carbody_pkg.CARBODY
     electronicequipmentinsurance electronicequipmentinsurance_pkg.ELECTRONICEQUIPMENTINSURANCE
     otherinsurancetypes otherinsurancetypes_pkg.OTHERINSURANCETYPES
     thirdpartyinsurance thirdpartyinsurance_pkg.THIRDPARTYINSURANCE
     motorcycleinsurance motorcycleinsurance_pkg.MOTORCYCLEINSURANCE
     fireinsurance fireinsurance_pkg.FIREINSURANCE
     earthquakeinsurance earthquakeinsurance_pkg.EARTHQUAKEINSURANCE
     travelinsurance travelinsurance_pkg.TRAVELINSURANCE
     elevatorinsurance elevatorinsurance_pkg.ELEVATORINSURANCE
     main main_pkg.MAIN
     combodata combodata_pkg.COMBODATA
     trackingdamage trackingdamage_pkg.TRACKINGDAMAGE
     footer footer_pkg.FOOTER
     insurancepolicyplan insurancepolicyplan_pkg.INSURANCEPOLICYPLAN
     config  configuration_pkg.CONFIGURATION
}

/**
     * Access to Configuration
     * @return Returns the Configuration instance
*/
func (me *EASYBIMEHLANDING_IMPL) Configuration() configuration_pkg.CONFIGURATION {
    return me.config
}
/**
     * Access to FileManager controller
     * @return Returns the FileManager() instance
*/
func (me *EASYBIMEHLANDING_IMPL) FileManager() filemanager_pkg.FILEMANAGER {
    if(me.filemanager) == nil {
        me.filemanager = filemanager_pkg.NewFILEMANAGER(me.config)
    }
    return me.filemanager
}
/**
     * Access to LiabilityDoctorInsurance controller
     * @return Returns the LiabilityDoctorInsurance() instance
*/
func (me *EASYBIMEHLANDING_IMPL) LiabilityDoctorInsurance() liabilitydoctorinsurance_pkg.LIABILITYDOCTORINSURANCE {
    if(me.liabilitydoctorinsurance) == nil {
        me.liabilitydoctorinsurance = liabilitydoctorinsurance_pkg.NewLIABILITYDOCTORINSURANCE(me.config)
    }
    return me.liabilitydoctorinsurance
}
/**
     * Access to CarBody controller
     * @return Returns the CarBody() instance
*/
func (me *EASYBIMEHLANDING_IMPL) CarBody() carbody_pkg.CARBODY {
    if(me.carbody) == nil {
        me.carbody = carbody_pkg.NewCARBODY(me.config)
    }
    return me.carbody
}
/**
     * Access to ElectronicEquipmentInsurance controller
     * @return Returns the ElectronicEquipmentInsurance() instance
*/
func (me *EASYBIMEHLANDING_IMPL) ElectronicEquipmentInsurance() electronicequipmentinsurance_pkg.ELECTRONICEQUIPMENTINSURANCE {
    if(me.electronicequipmentinsurance) == nil {
        me.electronicequipmentinsurance = electronicequipmentinsurance_pkg.NewELECTRONICEQUIPMENTINSURANCE(me.config)
    }
    return me.electronicequipmentinsurance
}
/**
     * Access to OtherInsuranceTypes controller
     * @return Returns the OtherInsuranceTypes() instance
*/
func (me *EASYBIMEHLANDING_IMPL) OtherInsuranceTypes() otherinsurancetypes_pkg.OTHERINSURANCETYPES {
    if(me.otherinsurancetypes) == nil {
        me.otherinsurancetypes = otherinsurancetypes_pkg.NewOTHERINSURANCETYPES(me.config)
    }
    return me.otherinsurancetypes
}
/**
     * Access to ThirdPartyInsurance controller
     * @return Returns the ThirdPartyInsurance() instance
*/
func (me *EASYBIMEHLANDING_IMPL) ThirdPartyInsurance() thirdpartyinsurance_pkg.THIRDPARTYINSURANCE {
    if(me.thirdpartyinsurance) == nil {
        me.thirdpartyinsurance = thirdpartyinsurance_pkg.NewTHIRDPARTYINSURANCE(me.config)
    }
    return me.thirdpartyinsurance
}
/**
     * Access to MotorcycleInsurance controller
     * @return Returns the MotorcycleInsurance() instance
*/
func (me *EASYBIMEHLANDING_IMPL) MotorcycleInsurance() motorcycleinsurance_pkg.MOTORCYCLEINSURANCE {
    if(me.motorcycleinsurance) == nil {
        me.motorcycleinsurance = motorcycleinsurance_pkg.NewMOTORCYCLEINSURANCE(me.config)
    }
    return me.motorcycleinsurance
}
/**
     * Access to FireInsurance controller
     * @return Returns the FireInsurance() instance
*/
func (me *EASYBIMEHLANDING_IMPL) FireInsurance() fireinsurance_pkg.FIREINSURANCE {
    if(me.fireinsurance) == nil {
        me.fireinsurance = fireinsurance_pkg.NewFIREINSURANCE(me.config)
    }
    return me.fireinsurance
}
/**
     * Access to EarthquakeInsurance controller
     * @return Returns the EarthquakeInsurance() instance
*/
func (me *EASYBIMEHLANDING_IMPL) EarthquakeInsurance() earthquakeinsurance_pkg.EARTHQUAKEINSURANCE {
    if(me.earthquakeinsurance) == nil {
        me.earthquakeinsurance = earthquakeinsurance_pkg.NewEARTHQUAKEINSURANCE(me.config)
    }
    return me.earthquakeinsurance
}
/**
     * Access to TravelInsurance controller
     * @return Returns the TravelInsurance() instance
*/
func (me *EASYBIMEHLANDING_IMPL) TravelInsurance() travelinsurance_pkg.TRAVELINSURANCE {
    if(me.travelinsurance) == nil {
        me.travelinsurance = travelinsurance_pkg.NewTRAVELINSURANCE(me.config)
    }
    return me.travelinsurance
}
/**
     * Access to ElevatorInsurance controller
     * @return Returns the ElevatorInsurance() instance
*/
func (me *EASYBIMEHLANDING_IMPL) ElevatorInsurance() elevatorinsurance_pkg.ELEVATORINSURANCE {
    if(me.elevatorinsurance) == nil {
        me.elevatorinsurance = elevatorinsurance_pkg.NewELEVATORINSURANCE(me.config)
    }
    return me.elevatorinsurance
}
/**
     * Access to Main controller
     * @return Returns the Main() instance
*/
func (me *EASYBIMEHLANDING_IMPL) Main() main_pkg.MAIN {
    if(me.main) == nil {
        me.main = main_pkg.NewMAIN(me.config)
    }
    return me.main
}
/**
     * Access to ComboData controller
     * @return Returns the ComboData() instance
*/
func (me *EASYBIMEHLANDING_IMPL) ComboData() combodata_pkg.COMBODATA {
    if(me.combodata) == nil {
        me.combodata = combodata_pkg.NewCOMBODATA(me.config)
    }
    return me.combodata
}
/**
     * Access to TrackingDamage controller
     * @return Returns the TrackingDamage() instance
*/
func (me *EASYBIMEHLANDING_IMPL) TrackingDamage() trackingdamage_pkg.TRACKINGDAMAGE {
    if(me.trackingdamage) == nil {
        me.trackingdamage = trackingdamage_pkg.NewTRACKINGDAMAGE(me.config)
    }
    return me.trackingdamage
}
/**
     * Access to Footer controller
     * @return Returns the Footer() instance
*/
func (me *EASYBIMEHLANDING_IMPL) Footer() footer_pkg.FOOTER {
    if(me.footer) == nil {
        me.footer = footer_pkg.NewFOOTER(me.config)
    }
    return me.footer
}
/**
     * Access to InsurancePolicyPlan controller
     * @return Returns the InsurancePolicyPlan() instance
*/
func (me *EASYBIMEHLANDING_IMPL) InsurancePolicyPlan() insurancepolicyplan_pkg.INSURANCEPOLICYPLAN {
    if(me.insurancepolicyplan) == nil {
        me.insurancepolicyplan = insurancepolicyplan_pkg.NewINSURANCEPOLICYPLAN(me.config)
    }
    return me.insurancepolicyplan
}
