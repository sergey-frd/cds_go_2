package strct

import (
        "time"
)
//.............................
type Ymd_KEY   struct {
    Year       int           
    Month      int    
    Day        int 
}

//.............................
type Ymdh_KEY   struct {
    Year       int           
    Month      int    
    Day        int 
    Hour       int           
}

//.............................
type Base_Ti_Val   struct {
    Index[]int
    Total_Slot    int
    Total_Cost    float64
}

//.............................
type Base_Ti_KEY   struct {
    CnCtNbDs  Digital_Signage_KEY
    Ymd_key  Ymd_KEY 
}
 
//.............................
type Base_TI_STC   struct {
    Base_ti_key   Base_Ti_KEY
    Base_ti_val   Base_Ti_Val
}

//.............................
type Ow_Ds_Used_KEY   struct {
    UsMd       User_Media_KEY
    CnCtNbDs   Digital_Signage_KEY
    Index      int  
    Start_Ymdh Ymdh_KEY 
    End_Ymdh   Ymdh_KEY 
}

type Ow_Ds_Used_STC   struct {
    Ow_day_ds_key Ow_Day_Ds_KEY
    Total_Cost    float64
}

//.............................
type Ds_Ti_KEY   struct {
    CnCtNbDs  Digital_Signage_KEY
    Ymdh_key  Ymdh_KEY 
}
 
//.............................
type DS_TI_STC   struct {
    Ds_ti_key  Ds_Ti_KEY
    Index      int  
}

//.............................
type Ow_Day_Ds_KEY   struct {
    Ymd_key   Ymd_KEY 
    CnCtNbDs  Digital_Signage_KEY
}

type Ow_Day_Ds_STC   struct {
    Ow_day_ds_key Ow_Day_Ds_KEY
    Total_Cost    float64
}


//.............................
type Ow_Day_KEY struct {
    Ymd_key     Ymd_KEY 
    ID_Owner    string
}

type Ow_Day_STC   struct {
    Ow_day_key          Ow_Day_KEY
    Total_Cost          float64
}

//.............................
//.............................
//.............................
type Ow_Um_KEY   struct {
    UsMd                User_Media_KEY
    ID_Owner            string
}

type Ow_Um_STC   struct {
    OwUm_Key            Ow_Um_KEY
    Total_Cost          float64
}

type Ow_UmNbDs_KEY   struct {
    UsMd                User_Media_KEY
    CnCtNbDs            Digital_Signage_KEY
}

type Ow_UmNbDs_STC   struct {
    OwUmNbDs_Key       Ow_UmNbDs_KEY
    Total_Cost         float64
}

//.............................
type Ow_UmNbDsTi_KEY   struct {
    UsMd                User_Media_KEY
    NbDsTi_key          NbDsTi_KEY
}

type Ow_UmNbDsTi_STC   struct {
    OwUmNbDsTi_Key     Ow_UmNbDsTi_KEY
    Total_Cost         float64
}
//.............................
type NbDsTi_KEY   struct {
    CnCtNbDs          Digital_Signage_KEY
    ID_Time_Interval  string
}

//.............................
//.............................
type NbDsTiSl_KEY   struct {
    NbDsTi_key  NbDsTi_KEY
    Year  int           
    Month int    
    Day   int           
    Hour  int           
    Index int           
}

type NbDsTiSl_STC   struct {
     NbDsTiSl_Key   NbDsTiSl_KEY
     Slot_Price     float64
}

//.............................
type Um_NbDsTiSl_KEY   struct {
    UsMd          User_Media_KEY
    NbDsTiSl_key  NbDsTiSl_KEY
}

type Um_NbDsTiSl_STC   struct {
    UmNbDsTiSl_Key     Um_NbDsTiSl_KEY
    Slot_Price         float64
}


//.............................
type User_Media_KEY   struct {
    ID_User     string
    ID_Media    string
}

type User_Media_VAL   struct {
    Media_Name  string
    Media_Cost  string
    Media_Slots string
    Start_time  time.Time
    End_time    time.Time

}

type User_Media_STC   struct {
    UsMd        User_Media_KEY
    UsMdVl      User_Media_VAL
}

//.............................
type Media_STC   struct {
    ID_Media    string
    MdV         Md_VAL
}

type Md_VAL   struct {
    Type_Media  string
    Slots       string
}

//.............................
type Time_Interval_VAL   struct {
    Price            string
    D_Sign_People    string
    Slots            string
}

//.............................
type Time_Interval_STC   struct {
    ID_Time_Interval string
    TiVl             Time_Interval_VAL
}

//.............................
type Neighborhoods_STC   struct {
    CnCtNb           Neighborhoods_KEY
    Neighborhoods    string
}

//.............................
type Neighborhoods_KEY   struct {
    CnCt             City_KEY
    ID_Neighborhoods string
}
 
//.............................
type Digital_Signage_KEY   struct {
    CnCtNb              Neighborhoods_KEY
    ID_Digital_Signage  string
    ID_Owner            string
}
 
//.............................
type Digital_Signage_VAL   struct {
    Dig_Sign         string
    DS_Cost          string
    DS_Perc_Quality  string
}
 
//.............................
type Digital_Signage_STC   struct {
    CnCtNbDs  Digital_Signage_KEY
    DsVal     Digital_Signage_VAL
}

//............................
type City_STC   struct {
    CnCt   City_KEY
    City   string
}


//.............................
type City_KEY   struct {
    ID_Country       string
    ID_City          string
}

//.............................
type Price_STC   struct {
    Index        string
    Price        string
}

type AutoGenerated struct {
	Comment string `json:"comment"`
	Base    struct {
		Comment                     string  `json:"comment"`
		ProjectName                 string  `json:"project_name"`
		ReleaseNumber               string  `json:"release_number"`
		ExcelName                   string  `json:"excel_Name"`
		DbName                      string  `json:"db_Name"`
		IDTestCountry               string  `json:"ID_Test_Country"`
		IDTestCity                  string  `json:"ID_Test_City"`
		CASEXCLS2DB                 string  `json:"CASE_XCLS_2_DB"`
		CASELOADDICT                string  `json:"CASE_LOAD_DICT"`
		CASEGENALLFILES             string  `json:"CASE_GEN_ALL_FILES"`
		CASEDBPRINTTBL              string  `json:"CASE_DB_PRINT_TBL"`
		CASEINIT                    string  `json:"CASE_INIT"`
		CASEUMALLOCATION            string  `json:"CASE_UM_ALLOCATION"`
		CASEDBPRINTALL              string  `json:"CASE_DB_PRINT_ALL"`
		FLAGDBPRINTLENONLY          int     `json:"FLAG_DB_PRINT_LEN_ONLY"`
		CASEMAPPRINTALL             string  `json:"CASE_MAP_PRINT_ALL"`
		FLAGMAPPRINTLENONLY         int     `json:"FLAG_MAP_PRINT_LEN_ONLY"`
		FLAGMAPPRINT                int     `json:"FLAG_MAP_PRINT"`
		Clip4ALLCountry             int     `json:"Clip_4_ALL_Country"`
		ClipCodeCountry             string  `json:"Clip_Code_Country"`
		ClipCodeCity                string  `json:"Clip_Code_City"`
		ClipBudgetMin               int     `json:"Clip_Budget_Min"`
		ClipBudgetMax               int     `json:"Clip_Budget_Max"`
		ClipBudgetStart             int     `json:"Clip_Budget_Start"`
		ClipBudgetEnd               int     `json:"Clip_Budget_End"`
		NotUsedMaxSlotCounter       int     `json:"Not_Used_Max_Slot_Counter"`
		DSignCounter                int     `json:"D_Sign_Counter"`
		DSignUsedCounter            int     `json:"D_Sign_Used_Counter"`
		TimeIntervalCounter         int     `json:"Time_Interval_Counter"`
		UserClipCounter             int     `json:"User_Clip_Counter"`
		UserCounter                 int     `json:"User_Counter"`
		NeighborhoodsCounter        int     `json:"Neighborhoods_Counter"`
		MaxDigitalSignage           int     `json:"Max_Digital_Signage"`
		MaxNeighborhoods            int     `json:"Max_Neighborhoods"`
		UserClipNeighborhoodCounter int     `json:"User_Clip_Neighborhood_Counter"`
		OwnerCounter                int     `json:"Owner_Counter"`
		CodeFreeUsers               int     `json:"Code_Free_Users"`
		CodeFreeClips               int     `json:"Code_Free_Clips"`
		CodeOtherUsers              int     `json:"Code_Other_Users"`
		CodeOtherClips              int     `json:"Code_Other_Clips"`
		TimeAddHourMin              int     `json:"Time_Add_Hour_Min"`
		TimeAddHourMax              int     `json:"Time_Add_Hour_Max"`
		TimeAddDayMin               int     `json:"Time_Add_Day_Min"`
		TimeAddDayMax               int     `json:"Time_Add_Day_Max"`
		DigSignMaxTimeIntervPrice   int     `json:"Dig_Sign_Max_Time_Interv_Price"`
		DigSignMaxTotalPrice        int     `json:"Dig_Sign_Max_Total_Price"`
		DigSignTimeIntSQRMin        float64 `json:"Dig_Sign_Time_Int_SQR_min"`
		DigSignTimeIntSQRMax        float64 `json:"Dig_Sign_Time_Int_SQR_max"`
		Min4PioplePrice             string  `json:"Min_4_Piople_Price"`
		DigitalSignageCostMax       int     `json:"Digital_Signage_Cost_Max"`
		DigitalSignageCostMin       int     `json:"Digital_Signage_Cost_Min"`
		TimeIntervalPeople          int     `json:"Time_Interval_People"`
		TimeSlotCost                int     `json:"Time_Slot_Cost"`
		TimeSlotBusyPers            int     `json:"Time_Slot_Busy_Pers"`
	} `json:"Base"`
	DBTableList []string `json:"DB_Table_List"`
} // AutoGenerated
