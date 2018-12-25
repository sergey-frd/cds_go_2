package lib_gen

import (   
    "fmt"
    "log"
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

//---------------------------------------------------------------
func __err_panic(err error) {
	if err != nil {
		log.Fatal(err)
		panic(err)
	}
}

//----------------------------------------------
//func Alloc_Um(byteValues  []byte, 
func Gen_Lvl_Um2(byteValues  []byte, 
    data            map[string]map[string]string,
    )  (err error) {


    p := fmt.Println
    //var total       float64
    var indx int
    var um   S.User_Media_STC
    //var ps   S.Um_NbDsTiSl_STC  // Payd_Slots
    var Um_nbds  S.Um_NbDs_STC
    Um_nbds.Um_nbds_key.Ds_keys = make([]S.Digital_Signage_KEY,10)

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
        // err = json.Unmarshal(byt_k, &ps.UmNbDsTiSl_Key.UsMd);  __err_panic(err)
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
        Um_nbds.Um_nbds_key.UsMd = um.UsMd
        fmt.Println("um Um_nbds =",Um_nbds)
        p("Um_nbds.Um_nbds_key.Ds_keys[0] =", Um_nbds.Um_nbds_key.Ds_keys[0])
        //p("Um_nbds.Um_nbds_key.Ds_keys[0].ID_Digital_Signage =", Um_nbds.Um_nbds_key.Ds_keys[0].ID_Digital_Signage)
        //if Um_nbds.Um_nbds_key.Ds_keys[0].ID_Digital_Signage == "" { p("ID_Digital_Signage = empty")}

       indx = 0
       err = Gen_Lvl_Ct2(byteValues,data,Um_nbds,um,indx,);  __err_panic(err)


        //!!!!!!!!!!!!!!!!!!!!
        if n >=1 { break }

    }

    return  err

} // func alloc_ow

