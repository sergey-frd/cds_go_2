package lib_gen

import (   
//    "fmt"
//    "log"
	"sort"
//    "github.com/valyala/fastjson"
//    "os" 
//    "github.com/boltdb/bolt"
//    "errors"

//    "strconv"
//  // "io/ioutil"
//  //"bytes"
//  //"runtime"
    "encoding/json"
//  //"encoding/gob"
//  
//  "time"
//  "math/rand"
//
//    L "cds_go_2/lib"
    S "cds_go_2/config"

)

//----------------------------------------------
//func Alloc_Um(byteValues  []byte, 
func Gen_Lvl_Um3(byteValues  []byte, 
    data            map[string]map[string]string,
    )  (err error) {


    //var total       float64

    var um   S.User_Media_STC
    //var ps   S.Um_NbDsTiSl_STC
    var ud   S.Um_Ds_STC

    var keys []string
    for k, _ := range data["User_Media"] {
        keys = append(keys, k)
    } // for k, v
    sort.Strings(keys)  // or sort.Ints(keys), sort.Sort(...), etc., per <K>
    n := 0
    for _, k := range keys  {
        n++
        v := data["User_Media"][k]


        //.................................................
        byt_k := []byte(k)
        //err := json.Unmarshal(byt_ds, &k_Ds)
        //err = json.Unmarshal(byt_k, &ps.UmNbDsTiSl_Key.UsMd);  __err_panic(err)
        err = json.Unmarshal(byt_k, &um.UsMd);  __err_panic(err)
        // if err != nil {
        //     fmt.Println("There was an error:", err)
        // }

        byt_v := []byte(v)
        err = json.Unmarshal(byt_v, &um.UsMdVl);  __err_panic(err)
        // if err != nil {
        //     fmt.Println("There was an error:", err)
        // }
        //fmt.Println("um =", um)


        //fmt.Println("um =", um)
        //fmt.Println("um ps =", ps)
        //ps.UmNbDsTiSl_Key.UsMd = um.UsMd
        ud.Um_ds_key.UsMd = um.UsMd
        err = Gen_Lvl_Ct3(byteValues,data,ud,um,);  __err_panic(err)


        //!!!!!!!!!!!!!!!!!!!!
        //if n >=1 { break }

    }

    return  err

} // func alloc_ow

