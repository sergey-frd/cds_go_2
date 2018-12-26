package lib_gen

import (   
//    "fmt"
//    "log"
	"sort"
    "github.com/valyala/fastjson"
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

    L "cds_go_2/lib"
    S "cds_go_2/config"

)


//----------------------------------------------
//func Alloc_Um(byteValues  []byte, 
func Gen_Lvl_Nb3(byteValues  []byte, 
    data       map[string]map[string]string,
    ud         S.Um_Ds_STC,
    um         S.User_Media_STC,
    ) error {


    var err         error
    //var total       float64

    var Nb       S.Neighborhoods_STC
    //var um   S.User_Media_STC
    var iCounter int

    Clip_4_ALL_Nb := L.Random(0, 2)
    Neighborhoods_Counter := fastjson.GetInt(byteValues, "Base", "Neighborhoods_Counter")
    Clip_Nb_Count := L.Random(1, Neighborhoods_Counter)

    var keys []string
    for k, _ := range data["Neighborhoods"] {
        keys = append(keys, k)
    } // for k, v
    sort.Strings(keys)
    nnn := 0
    for _, k := range keys  {

        byt_k := []byte(k)
        err = json.Unmarshal(byt_k,  &Nb.CnCtNb);  __err_panic(err)
        // if err != nil {
        //     fmt.Println("There was an error:", err)
        // }

        if Nb.CnCtNb.CnCt != ud.Um_ds_key.Ds_key.CnCtNb.CnCt {
            continue
        }

        nnn  += 1
        //Um_nbds.Um_nbds_key.Ds_keys[indx].CnCtNb = Nb.CnCtNb
        ud.Um_ds_key.Ds_key.CnCtNb = Nb.CnCtNb

        //fmt.Println("um =", um)
        //fmt.Println("    nb ud =", ud)
        //err = Gen_Lvl_Ds2(byteValues,data,Um_nbds,um,indx+nnn-1,);  __err_panic(err)
        err = Gen_Lvl_Ds3(byteValues,data,ud,um,);  __err_panic(err)

        if Clip_4_ALL_Nb == 1 {
            iCounter = Neighborhoods_Counter
        } else {
            iCounter = Clip_Nb_Count
        }

        if nnn >= iCounter {
            break	
        } 


        //!!!!!!!!!!!!!!!!!!!!
        break

    }

    return  err

} // func alloc_ow

