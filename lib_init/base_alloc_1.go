package lib_init

import (   
//    "fmt"
    "sort"
    "encoding/json"
    "time"

// //        "math"
//         "strconv"
// //        "time"
// //
// //    "log"
// //    "github.com/valyala/fastjson"
// //    "os" 
// //    "github.com/boltdb/bolt"
// //    "errors"
// 
// //  // "io/ioutil"
// //  //"bytes"
// //  //"runtime"
// //  //"encoding/gob"
// //  
// //  "math/rand"

// 
//     L "cds_go_2/lib"
     S "cds_go_2/config"

)


//  err = G.Gen_Ds_Ti(json_go,data,);  __err_panic(err)
//----------------------------------------------
//func Alloc_Um(byteValues  []byte, 
func base_alloc(
    json_go    S.AutoGenerated,
    data       map[string]map[string]string,
    um         S.User_Media_STC,
    ud         S.Um_Ds_STC,
    ) (int, float64, error) {


    var      err error

    var Total_Slot    int
    var Total_Cost    float64

    //p := fmt.Println
    // var err         error
    //var total       float64

    //var ct   S.City_STC
    //var Ds   S.Digital_Signage_STC



    var bti  S.Base_TI_STC
    //var odu  S.Ow_Ds_Used_STC
    //var dst  S.DS_TI_STC
    //var ct   S.User_Media_STC
    //var Ti   S.Time_Interval_STC

    //var iCounter int

    var      keys []string
    // var      Slot       int
    // var      Total_Slot int
    // var      Total_Cost float64
    // 
    // var    All_Used      int
    // var    New_Free_Slot int
    // var    Start_Indx    int
    // var    Interv_Count  int
    // var    Need_Used     int
    // var    High_Slot     int


    // TimeAddDayMin  := json_go.Base.TimeAddDayMin
    // TimeAddDayMax  := json_go.Base.TimeAddDayMax
    // 
    // 
    // 
    // //var um   S.User_Media_STC
    // 
    // //............................................................
    // //Clip_Code_Country  := json_go.Base.ClipCodeCountry 
    // //Clip_Code_City     := json_go.Base.ClipCodeCity    
    // //
    // //ct.CnCt.ID_Country  = Clip_Code_Country
    // //ct.CnCt.ID_City     = Clip_Code_City
    // 
    // Time_Interval_Counter     := json_go.Base.TimeIntervalCounter    
    // 
    // 
    // 
    // CodeFreeUsers := jsongo.Base.CodeFreeUsers 
    // CodeFreeClips := jsongo.Base.CodeFreeClips 
    // 
    // //um.UsMd.IDUser  = strconv.Itoa(CodeFreeUsers)
    // //um.UsMd.IDMedia = strconv.Itoa(CodeFreeClips)
    // 
    // odu.Owdaydskey.UsMd.IDUser  = strconv.Itoa(CodeFreeUsers)  
    // odu.Owdaydskey.UsMd.IDMedia = strconv.Itoa(CodeFreeClips)  
    // //.............................


    //Time_Interval_Counter  := json_go.Base.TimeIntervalCounter
    //inxArr := make([]int, Time_Interval_Counter)

    // var sYmd  S.Ymd_KEY
    // var eYmd  S.Ymd_KEY
    // 
    // 
    // sYmd.Year   =     um.UsMdVl.Start_time.Year() 
    // sYmd.Month  = int(um.UsMdVl.Start_time.Month())
    // sYmd.Day    =     um.UsMdVl.Start_time.Day()  
    // 
    // eYmd.Year   =     um.UsMdVl.End_time.Year() 
    // eYmd.Month  = int(um.UsMdVl.End_time.Month())
    // eYmd.Day    =     um.UsMdVl.End_time.Day()  
    // 
    // fmt.Println("sYmd =", sYmd)
    // fmt.Println("eYmd =", eYmd)



     t := time.Date(
         um.UsMdVl.Start_time.Year() ,
         um.UsMdVl.Start_time.Month(),
         um.UsMdVl.Start_time.Day()  ,
         0   ,
         0   ,
         0   ,
         0   , 
         time.UTC,
         )
 
     //t  = t.Add(time.Hour * -24)

     e := time.Date(
         um.UsMdVl.End_time.Year() ,
         um.UsMdVl.End_time.Month(),
         um.UsMdVl.End_time.Day()  ,
         0   ,
         0   ,
         0   ,
         0   , 
         time.UTC,
         )

    Total_Slot    = 0
    Total_Cost    = 0


    tblname   := "Base_Ti"
    for k, _ := range data[tblname] {
        keys = append(keys, k)
    } // for k, v
    sort.Strings(keys)
    nnn := 0
    //nn := 0
    for _, k := range keys  {
      bytk := []byte(k)
      err = json.Unmarshal(bytk, &bti.Base_ti_key);  __err_panic(err)

      // if Ds.CnCtNbDs.CnCtNb != ps.UmNbDsTiSlKey.NbDsTiSlkey.NbDsTikey.CnCtNbDs.CnCtNb {
      //     continue
      // }

      dsvalue:= data[tblname][k]
      bytdsv := []byte(dsvalue)
      err = json.Unmarshal(bytdsv, &bti.Base_ti_val);  __err_panic(err)

      nnn  += 1
      if bti.Base_ti_key.CnCtNbDs != ud.Um_ds_key.Ds_key { continue }

        c := time.Date(
             bti.Base_ti_key.Ymd_key.Year ,
             time.Month(bti.Base_ti_key.Ymd_key.Month),
             bti.Base_ti_key.Ymd_key.Day  ,
             0   ,
             0   ,
             0   ,
             0   , 
             time.UTC,
         )

        if c.Before(t) { continue }
        if c.After(e)  { continue }



        // fmt.Println(nnn,"bti =", bti)
        Total_Slot    += bti.Base_ti_val.Total_Slot
        Total_Cost    += bti.Base_ti_val.Total_Cost


      //if nnn >= 3 { break }

      //odu.Owdaydskey.CnCtNbDs  = dst.Dstikey.CnCtNbDs

       //  All_Used      = L.Random(1,10)
       //  Need_Used     = L.Random(0,5)
       // 
       //  High_Slot     = L.Random(0,3)
       //  New_Free_Slot = L.Random(0,6)
       //  All_Used      = L.Random(0,5)
       //  Start_Indx    = L.Random(0,11)
       //  Interv_Count  = L.Random(2,12)
       //  Need_Used     = L.Random(0,5)
       // 
       //  if Need_Used == 0 { continue }
       // 
       //  nn  += 1
       // 
       //  //p(nn,nnn,All_Used,Start_Indx,Interv_Count,"bti =", New_Free_Slot, bti)
       // 
       //  Ds.CnCtNbDs = bti.Base_ti_key.CnCtNbDs
       // 
       //  ds_encoded, err := json.Marshal(Ds.CnCtNbDs)
       //  __err_panic(err)
       // 
       //  //ds_encoded, err := json.Marshal(Ds.CnCtNbDs);  
       //  //p(nnn,"ds_encoded =", ds_encoded)
       //  //p(nnn,"string ds_encoded =", string(ds_encoded))
       //  ds_value:= data["Digital_Signage"][string(ds_encoded)]
       //  //p(nnn,"ds_value =", ds_value)
       //  byt_dsv := []byte(ds_value)
       //  err = json.Unmarshal(byt_dsv, &Ds.DsVal)
       //  __err_panic(err)
       //  //p(nnn,"Ds.DsVal =", Ds.DsVal)
       // 
       //  //p(nnn,"Ds =", Ds)
       // 
       //  bti.Base_ti_val.Total_Cost     = 0 
       //  Total_Slot = 0
       //  Total_Cost = 0
       // 
       //  inxArr = bti.Base_ti_val.Index
       //  for h := 0; h < Time_Interval_Counter; h++ {
       // 
       //      if All_Used == 1 {
       //          inxArr[h] = 0
       //      } else { // if All_Used
       //          if High_Slot == 1 {
       //              inxArr[h] = 6
       //          } else { // if High_Slot
       // 
       //              if Start_Indx <= h && Start_Indx+Interv_Count >= h {
       //              //if Start_Indx <= h {
       //                  inxArr[h] = New_Free_Slot
       //                  //p(h,Start_Indx,Start_Indx+Interv_Count,"inxArr[h] =", inxArr[h])
       //              } //if Start_Indx 
       //          } // if High_Slot
       // 
       //          Slot = inxArr[h]
       //          slots,slot_price,err := Get_Price_Slots(h,Slot,Ds,json_go,data,);  __err_panic(err)
       //          Total_Slot +=  slots
       //          Total_Cost +=  slot_price
       // 
       //      } // if All_Used
       //  } // for h := 0; h < diff_days; h++ 
       // 
       //  bti.Base_ti_val.Index      = inxArr 
       //  //bti.Base_ti_val.Total_Slot = Time_Interval_Counter * 6
       //  bti.Base_ti_val.Total_Slot = Total_Slot
       //  bti.Base_ti_val.Total_Cost = Total_Cost
       // 
       // 
       //  enc_Base_ti_key, err := json.Marshal(bti.Base_ti_key); __err_panic(err)
       //  enc_Base_ti_val, err := json.Marshal(bti.Base_ti_val); __err_panic(err)
       // 
       //  //p("dst =", dst)
       //  //dst_Index  := strconv.Itoa(dst.Index)
       //  data["Base_Ti"][string(enc_Base_ti_key)]  = string(enc_Base_ti_val)
       // 
       // 
       //  //p(nn,nnn,All_Used,Start_Indx,Interv_Count,"upd =", New_Free_Slot, bti)
       // 
       // 
       //  //!!!!!!!!!!!!!!!!!!!!
       //  //if nnn >= 5 { break }

    } // for , k := range keys 



    return  Total_Slot,Total_Cost,err

} // func allocow

