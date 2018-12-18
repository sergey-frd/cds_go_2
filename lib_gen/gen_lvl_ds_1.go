package lib_gen

import (   
//    "fmt"

//        "math"
//        "strconv"
        "time"
//
//    "log"
	"sort"
//    "github.com/valyala/fastjson"
//    "os" 
//    "github.com/boltdb/bolt"
//    "errors"

//  // "io/ioutil"
//  //"bytes"
//  //"runtime"
    "encoding/json"
//  //"encoding/gob"
//  
//  "math/rand"

//    L "cds_go_2/lib"
    S "cds_go_2/config"

)


//----------------------------------------------
//func Alloc_Um(byteValues  []byte, 
func Gen_Lvl_Ds(byteValues  []byte, 
    data       map[string]map[string]string,
    ps         S.Um_NbDsTiSl_STC,
    um         S.User_Media_STC,
    ) (err error) {


    //p := fmt.Println
    // var err         error
    //var total       float64

    var Ds   S.Digital_Signage_STC
    //var um   S.User_Media_STC
    var iCounter int

    var keys []string
    for k, _ := range data["Digital_Signage"] {
        keys = append(keys, k)
    } // for k, v
    sort.Strings(keys)
    nnn := 0
    for _, k := range keys  {

        byt_k := []byte(k)
        err = json.Unmarshal(byt_k, &Ds.CnCtNbDs);  __err_panic(err)


        if Ds.CnCtNbDs.CnCtNb != ps.UmNbDsTiSl_Key.NbDsTiSl_key.NbDsTi_key.CnCtNbDs.CnCtNb {
            continue
        }

        ds_value:= data["Digital_Signage"][k]
        byt_dsv := []byte(ds_value)
        err = json.Unmarshal(byt_dsv, &Ds.DsVal);  __err_panic(err)

        nnn  += 1


        ps.UmNbDsTiSl_Key.NbDsTiSl_key.NbDsTi_key.CnCtNbDs = Ds.CnCtNbDs

        //fmt.Println("      Ds =", Ds)
        //fmt.Println("      ds ps =", ps)
        //fmt.Println("      um.UsMdVl            =", um.UsMdVl)
        //fmt.Println("      um.UsMdVl.Start_time =", um.UsMdVl.Start_time)
        //fmt.Println("      um.UsMdVl.End_time   =", um.UsMdVl.End_time  )

        //year, month, day, hour, min, sec := L.Diff(um.UsMdVl.Start_time, um.UsMdVl.End_time)
        //fmt.Printf("      diff = %d years, %d months, %d days, %d hours, %d mins and %d seconds\n",
        //    year, month, day, hour, min, sec)
        // err = Gen_Lvl_Ti(byteValues,data,ps,);  __err_panic(err)

        diff := um.UsMdVl.End_time.Sub(um.UsMdVl.Start_time)
        //p("diff =",diff)
        //p("diff.Hours())        =",diff.Hours())
        //p("diff.Minutes())      =",diff.Minutes())
        //p("diff.Seconds())      =",diff.Seconds())

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

        e  = e.Add(time.Hour * 24)
        diff = e.Sub(t)
        //p("2 t =",t)
        //p("2 e =",e)
        //p("2 diff =",diff)
        //p("2 diff.Hours())        =",diff.Hours())
        
        c := t
        iCounter = int(diff.Hours())
        //p("2 iCounter =",iCounter)

        n := 0
        for i := 0; i < iCounter; i++ {
            n++
            c  = c.Add(time.Hour)

            //p(i,c)
            if um.UsMdVl.End_time.Equal(c) {
                //p("break i =",i)
                break
            } 
            //p("c =",c)
            //err = Gen_Lvl_Sl(byteValues,data,ps,um,c,);  __err_panic(err)
            err = Gen_Lvl_Ti(byteValues,data,ps,um,Ds,c,);  __err_panic(err)

            //    //!!!!!!!!!!!!!!!!!!!!

        } // for i

        //!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!
        //if nnn >=5 { break }
        //p("3 c =",c)


        //!!!!!!!!!!!!!!!!!!!!
        //break

    }

    return  err

} // func alloc_ow

