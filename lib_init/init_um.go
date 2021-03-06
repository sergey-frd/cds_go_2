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
//  "time"
//  "math/rand"
//
//    L "cds_go_2/lib"
    S "cds_go_2/config"
)
// //---------------------------------------------------------------
// func __err_panic(err error) {
// 	if err != nil {
// 		log.Fatal(err)
// 		panic(err)
// 	}
// }
//----------------------------------------------
func Init_Um(byteValues  []byte, 
    data            map[string]map[string]string,
    TotalDict       map[string]map[string]float64,
    )  (err error) {

    var       tbl_name   string
    //var       ps         S.Um_NbDsTiSl_STC
    //var       od         S.Ow_Day_STC
    //var       ods        S.Ow_Day_Ds_STC

    //var total       float64

    //var um          S.User_Media_STC

    // var dsfs        S.NbDsTiSl_STC     // Free Ds Slots
    // var err         error
    // 
    // var total       float64
    // var count       int

    //var dsfs  S.NbDsTiSl_STC     // Free Ds Slots
    //var Ds          S.Digital_Signage_STC ,

    //fmt.Println("Alloc_OwUmNbDsTiSl oundt =",oundt)
    //fmt.Println("Alloc_OwUmNbDs ou.OwUm_Key.ID_Owner =",ou.OwUm_Key.ID_Owner)
    //fmt.Println("Alloc_Ow um.UsMd =",um.UsMd)



    var um          S.User_Media_STC

    var u_keys []string
    tbl_name   = "User_Media"
    for u_k, _ := range data[tbl_name] {
        u_keys = append(u_keys, u_k)
    } // for k, v
    sort.Strings(u_keys)  // or sort.Ints(keys), sort.Sort(...), etc., per <K>

    for _, u_k := range u_keys  {
        u_v := data[tbl_name][u_k]

        //.................................................
        byt_u_k := []byte(u_k)
        //err := json.Unmarshal(byt_ds, &k_Ds)
        err = json.Unmarshal(byt_u_k, &um.UsMd)
        if err != nil {
            fmt.Println("There was an error:", err)
        }

        byt_u_v := []byte(u_v)
        err = json.Unmarshal(byt_u_v, &um.UsMdVl)
        if err != nil {
            fmt.Println("There was an error:", err)
        }
        fmt.Println("um =", um)
        err = Init_Um_Ds(byteValues,data,TotalDict,um,);  __err_panic(err)

        //---------------------------------------------------

        //!!!!!!!!!!!!!!!!!!!!
        break

    } // for _, k := range keys  {

    return  err

} // func 
