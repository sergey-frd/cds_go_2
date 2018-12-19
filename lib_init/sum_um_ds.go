package lib_init

import (   
    "fmt"
//	"log"
	"sort"
    "encoding/json"
//    "github.com/valyala/fastjson"
//    "os" 
//    "github.com/boltdb/bolt"
//    "errors"

//    "strconv"
//  // "io/ioutil"
//  //"bytes"
//  //"runtime"
//  //"encoding/gob"
//  
  "time"
//  "math/rand"
//
    L "cds_go_2/lib"
    S "cds_go_2/config"
)


//    c_time    time.Time,

//----------------------------------------------
func sum_Um_Ds(byteValues  []byte, 
    data      map[string]map[string]string,
    TotalDict map[string]map[string]float64,
    um        S.User_Media_STC,
    )  (err error) {

    var       tbl_name   string
    var       ps         S.Um_NbDsTiSl_STC
    var       od         S.Ow_Day_STC
    var       ods        S.Ow_Day_Ds_STC
    var       total       float64

    //var total       float64

    //var um          S.User_Media_STC

    // var dsfs        S.NbDsTiSl_STC     // Free Ds Slots
    // var err         error
    // 
    // var count       int

    //var dsfs  S.NbDsTiSl_STC     // Free Ds Slots
    //var Ds          S.Digital_Signage_STC ,

    //fmt.Println("Alloc_OwUmNbDsTiSl oundt =",oundt)
    //fmt.Println("Alloc_OwUmNbDs ou.OwUm_Key.ID_Owner =",ou.OwUm_Key.ID_Owner)
    //fmt.Println("Alloc_Ow um.UsMd =",um.UsMd)

    //fmt.Println("um =", um)


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
 
     //e  = e.Add(time.Hour * 24)

    fmt.Println("t = ",t)
    fmt.Println("e = ",e)
    //---------------------------------------------------
    tbl_name   = "Free_Slots"


    fmt.Println("Free_Slots start len =", len(data[tbl_name]))

    var keys []string
    for k, _ := range data[tbl_name] {
        keys = append(keys, k)
    } // for k, v
    sort.Strings(keys)  // or sort.Ints(keys), sort.Sort(...), etc., per <K>

    //total = 0
    //count = 0
    n := 0
    nnn := 0
    for _, k := range keys  {
        v := data[tbl_name][k]

        //fmt.Println(k, "=>", v)

        //.................................................
        byt_k := []byte(k)
        //err := json.Unmarshal(byt_ds, &k_Ds)
        err = json.Unmarshal(byt_k, &ps.UmNbDsTiSl_Key)
        if err != nil {
            fmt.Println("There was an error:", err)
        }
        
        byt_v := []byte(v)
        err = json.Unmarshal(byt_v, &ps.Slot_Price)
        if err != nil {
            fmt.Println("There was an error:", err)
        }

        if ps.UmNbDsTiSl_Key.UsMd != um.UsMd { continue }

         c_time := time.Date(
            ps.UmNbDsTiSl_Key.NbDsTiSl_key.Year  ,
            time.Month(ps.UmNbDsTiSl_Key.NbDsTiSl_key.Month) ,
            ps.UmNbDsTiSl_Key.NbDsTiSl_key.Day   ,
            0   ,
            0   ,
            0   ,
            0   , 
            time.UTC,
            )

        if c_time.Before(t) { continue }
        if c_time.After(e)  { continue }

        if ps.UmNbDsTiSl_Key.NbDsTiSl_key.Index != 1 { continue }




        nnn  += 1
        od.Ow_day_key.Ymd_key.Year  = ps.UmNbDsTiSl_Key.NbDsTiSl_key.Year
        od.Ow_day_key.Ymd_key.Month = ps.UmNbDsTiSl_Key.NbDsTiSl_key.Month    
        od.Ow_day_key.Ymd_key.Day   = ps.UmNbDsTiSl_Key.NbDsTiSl_key.Day   
        od.Ow_day_key.ID_Owner      = ps.UmNbDsTiSl_Key.NbDsTiSl_key.NbDsTi_key.CnCtNbDs.ID_Owner

        ods.Ow_day_ds_key.Ymd_key   = od.Ow_day_key.Ymd_key
        ods.Ow_day_ds_key.CnCtNbDs  = ps.UmNbDsTiSl_Key.NbDsTiSl_key.NbDsTi_key.CnCtNbDs

        // ow_day_key_enc,    err := json.Marshal(od.Ow_day_key);     __err_panic(err)
        ow_day_ds_key_enc, err := json.Marshal(ods.Ow_day_ds_key); __err_panic(err)
        

        key := string(ow_day_ds_key_enc)

        sum := ps.Slot_Price

        //if sum < TotalDict["Sum_Ow_Ds"][key] {
        if  TotalDict["Paid_Ow_Ds"][key] + sum < TotalDict["Sum_Ow_Ds"][key] {

                // if nnn < 5 {
                //     fmt.Println(nnn,"ps = ", ps)
                //     fmt.Println(nnn,"ods.Ow_day_ds_key = ", ods.Ow_day_ds_key)
                // 
                //     fmt.Println("ps.UmNbDsTiSl_Key.UsMd = ",ps.UmNbDsTiSl_Key.UsMd)
                //     fmt.Println("ps.Slot_Price = ",ps.Slot_Price)
                //     fmt.Println("ps.UmNbDsTiSl_Key.NbDsTiSl_key = ",ps.UmNbDsTiSl_Key.NbDsTiSl_key)
                //     fmt.Println("ps.UmNbDsTiSl_Key.NbDsTiSl_key.Year  = ",ps.UmNbDsTiSl_Key.NbDsTiSl_key.Year )
                //     fmt.Println("ps.UmNbDsTiSl_Key.NbDsTiSl_key.Month = ",ps.UmNbDsTiSl_Key.NbDsTiSl_key.Month)
                //     fmt.Println("ps.UmNbDsTiSl_Key.NbDsTiSl_key.Day   = ",ps.UmNbDsTiSl_Key.NbDsTiSl_key.Day  )
                //     fmt.Println("ps.UmNbDsTiSl_Key.NbDsTiSl_key.Hour  = ",ps.UmNbDsTiSl_Key.NbDsTiSl_key.Hour )
                //     fmt.Println("ps.UmNbDsTiSl_Key.NbDsTiSl_key.Index = ",ps.UmNbDsTiSl_Key.NbDsTiSl_key.Index)
                //     fmt.Println("ps.UmNbDsTiSl_Key.NbDsTiSl_key.NbDsTi_key = ",ps.UmNbDsTiSl_Key.NbDsTiSl_key.NbDsTi_key)
                //     fmt.Println("ps.UmNbDsTiSl_Key.NbDsTiSl_key.NbDsTi_key.ID_Time_Interval = ",ps.UmNbDsTiSl_Key.NbDsTiSl_key.NbDsTi_key.ID_Time_Interval)
                //     fmt.Println("ps.UmNbDsTiSl_Key.NbDsTiSl_key.NbDsTi_key.CnCtNbDs = ",ps.UmNbDsTiSl_Key.NbDsTiSl_key.NbDsTi_key.CnCtNbDs)
                //     fmt.Println("ps.UmNbDsTiSl_Key.NbDsTiSl_key.NbDsTi_key.CnCtNbDs.ID_Digital_Signage = ",ps.UmNbDsTiSl_Key.NbDsTiSl_key.NbDsTi_key.CnCtNbDs.ID_Digital_Signage)
                //     fmt.Println("ps.UmNbDsTiSl_Key.NbDsTiSl_key.NbDsTi_key.CnCtNbDs.ID_Owner = ",ps.UmNbDsTiSl_Key.NbDsTiSl_key.NbDsTi_key.CnCtNbDs.ID_Owner)
                //     fmt.Println("ps.UmNbDsTiSl_Key.NbDsTiSl_key.NbDsTi_key.CnCtNbDs.CnCtNb = ",ps.UmNbDsTiSl_Key.NbDsTiSl_key.NbDsTi_key.CnCtNbDs.CnCtNb)
                //     fmt.Println("ps.UmNbDsTiSl_Key.NbDsTiSl_key.NbDsTi_key.CnCtNbDs.CnCtNb.CnCt = ",ps.UmNbDsTiSl_Key.NbDsTiSl_key.NbDsTi_key.CnCtNbDs.CnCtNb.CnCt            )
                //     fmt.Println("ps.UmNbDsTiSl_Key.NbDsTiSl_key.NbDsTi_key.CnCtNbDs.CnCtNb.CnCt.ID_Country  = ",ps.UmNbDsTiSl_Key.NbDsTiSl_key.NbDsTi_key.CnCtNbDs.CnCtNb.CnCt.ID_Country            )
                //     fmt.Println("ps.UmNbDsTiSl_Key.NbDsTiSl_key.NbDsTi_key.CnCtNbDs.CnCtNb.CnCt.ID_City     = ",ps.UmNbDsTiSl_Key.NbDsTiSl_key.NbDsTi_key.CnCtNbDs.CnCtNb.CnCt.ID_City               )
                //     fmt.Println("ps.UmNbDsTiSl_Key.NbDsTiSl_key.NbDsTi_key.CnCtNbDs.CnCtNb.ID_Neighborhoods = ",ps.UmNbDsTiSl_Key.NbDsTiSl_key.NbDsTi_key.CnCtNbDs.CnCtNb.ID_Neighborhoods)
                // }



                //fmt.Println(nnn,"TotalDict[Paid_Ow_Ds][key] = ", TotalDict["Paid_Ow_Ds"][key])
                //fmt.Println("sum = ",sum)
                TotalDict["Paid_Ow_Ds"][key] += sum
                total += sum
                data["Paid_Slots"][string(k)] = fmt.Sprintf("%f", v)     
                delete(data["Free_Slots"], string(k));
 
                //fmt.Println("+++++",nnn,ods.Ow_day_ds_key, sum, TotalDict["Sum_Ow_Ds"][key], TotalDict["Paid_Ow_Ds"][key])
                n++
        } else {
            //fmt.Println(">>>>>>",nnn,ods.Ow_day_ds_key, sum, TotalDict["Sum_Ow_Ds"][key], TotalDict["Paid_Ow_Ds"][key])
            continue
        }



        
        //!!!!!!!!!!!!!!
        //if nnn >= 300 { break }


    } // for _, k := range keys  {

    // fmt.Println(nnn,"Ow_Day   = ", TotalDict["Ow_Day"])
    //fmt.Println(nnn,"Ow_Day_Ds= ", TotalDict["Ow_Day_Ds"])
    // fmt.Println(nnn,"Ow_Day_Ds=> ", n)
    // 
    // err = Init_Um_Day(byteValues,data,TotalDict,um,);  __err_panic(err)

    fmt.Println(nnn,"total = ", total)

    fmt.Println("Free_Slots END len =", len(data[tbl_name]))

    err = L.Del_DB_Bucket(byteValues, "Free_Slots");          __err_panic(err)
    err = L.Save_Data_Map(byteValues, "Free_Slots"  , data ); __err_panic(err)
    err = L.LoadDict2(byteValues, data, "Free_Slots" );       __err_panic(err) 

    fmt.Println("Free_Slots END 2 len =", len(data[tbl_name]))

    return  err

} // func alloc_ow
