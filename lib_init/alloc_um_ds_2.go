package lib_init

import (   
    "fmt"
//	"log"
	"sort"
    "encoding/json"
    "strconv"
//  "time"
//    "github.com/valyala/fastjson"
//    "os" 
//    "github.com/boltdb/bolt"
//    "errors"

//  // "io/ioutil"
//  //"bytes"
//  //"runtime"
//  //"encoding/gob"
//  
//  "math/rand"
//
//    L "cds_go_2/lib"
    S "cds_go_2/config"
)


//    c_time    time.Time,

//----------------------------------------------
func Alloc_Um_Ds2(
    json_go      S.AutoGenerated,
    byteValues   []byte, 
    data         map[string]map[string]string,
    TotalDict    map[string]map[string]float64,
    um           S.User_Media_STC, 
    Um_base_cost float64,
    )  (err error) {

    p := fmt.Println

    //var err              error
    var ud               S.Um_Ds_STC
    //var Um_Total_Cost    float64
    var Media_Cost    float64

    //Um_Total_Cost = 0
    ////---------------------------------------------------
    tbl_name   := "Um_Ds"
    var keys []string
    for k, _ := range data[tbl_name] {
        keys = append(keys, k)
    } // for k, v
    sort.Strings(keys)  // or sort.Ints(keys), sort.Sort(...), etc., per <K>

    n := 0
    //nnn := 0
    for _, k := range keys  {
        n++
        v := data[tbl_name][k]
        //fmt.Println(n,k, "=>", v)
       //.................................................
       byt_k := []byte(k)
       //err := json.Unmarshal(byt_ds, &k_Ds)
       err = json.Unmarshal(byt_k, &ud.Um_ds_key)
       if err != nil {
           fmt.Println("There was an error:", err)
       }
       
       byt_v := []byte(v)
       err = json.Unmarshal(byt_v, &ud.Slot_Price)
       if err != nil {
           fmt.Println("There was an error:", err)
       }
    
        if ud.Um_ds_key.UsMd != um.UsMd { continue }

        // Total_Slot,Total_Cost,err  := base_alloc(json_go, data, um, ud,);    __err_panic(err)
        // 
        // fmt.Println(n,"Total_Slot =", Total_Slot)
        // fmt.Println(n,"Total_Cost =", Total_Cost)
        // 
        // ud.Slot_Price = Total_Cost
        // 
        // data["Um_Ds"][k] = fmt.Sprintf("%.6f", ud.Slot_Price)  
        // 
        // Um_Total_Cost += Total_Cost

        fmt.Println(n,"Alloc_Um_Ds2 ud =", ud)

        Media_Cost,err = strconv.ParseFloat(um.UsMdVl.Media_Cost, 64)
        //p("Media_Cost =",Media_Cost)

        Cost_koef := Um_base_cost /Media_Cost
        //p("Cost_koef =",Cost_koef)

        Media_Ds_Cost := ud.Slot_Price / Cost_koef
        p("Media_Ds_Cost =",Media_Ds_Cost)
        err  := alloc_um_ds_day(json_go, data, um, ud, Cost_koef);    __err_panic(err)


        // if n >= 2 { break }

    } //for _, k := range keys  



    return  err

} // func alloc_ow
