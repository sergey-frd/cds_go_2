package lib_init

import (   
    "fmt"

//        "math"
        "strconv"
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
    //"encoding/json"
//  //"encoding/gob"
//  
//  "math/rand"

    L "cds_go_1/lib"
    S "cds_go_1/config"

)

//----------------------------------------------
//func Alloc_Um(byteValues  []byte, 
func Init_Um_Day(byteValues  []byte, 
    data       map[string]map[string]string,
    TotalDict       map[string]map[string]float64,
    um         S.User_Media_STC,
    ) (err error) {


    p := fmt.Println
    // var err         error
    //var total       float64

    //var Ds   S.Digital_Signage_STC
    //var um   S.User_Media_STC
    // var iCounter int

    var keys []string
    for k, _ := range data["Digital_Signage"] {
        keys = append(keys, k)
    } // for k, v
    sort.Strings(keys)


     
     year, month, day, hour, min, sec := L.Diff(um.UsMdVl.Start_time, um.UsMdVl.End_time)
     fmt.Printf("diff = %d years, %d months, %d days, %d hours, %d mins and %d seconds\n",
         year, month, day, hour, min, sec)
 
     diff := um.UsMdVl.End_time.Sub(um.UsMdVl.Start_time)
     p("diff =",diff)
     p("diff.Hours())        =",diff.Hours())
     p("diff.Minutes())      =",diff.Minutes())
     p("diff.Seconds())      =",diff.Seconds())
 
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
 
     //e  = e.Add(time.Hour * 24)
     diff = e.Sub(t)
     //p("2 t =",t)
     //p("2 e =",e)
     //p("2 diff =",diff)
     diff_Hours := diff.Hours()
     p("diff_Hours =",diff_Hours)
     diff_days := diff_Hours/24
     p("diff_days  =",diff_days)
     

     Media_Cost,err := strconv.ParseFloat(um.UsMdVl.Media_Cost, 64)
     p("Media_Cost =",Media_Cost)
     Day_Media_Cost := Media_Cost /float64(diff_days)
     p("Day_Media_Cost =",Day_Media_Cost)



     err = Alloc_Um_Ds(byteValues,data,TotalDict,um,Day_Media_Cost);  __err_panic(err)

     // c := t
     // iCounter = int(diff_days)
     // //p("2 iCounter =",iCounter)
     // 
     // 
     // 
     // 
     // n := 0
     // for i := 0; i < iCounter; i++ {
     //     n++
     //     //c  = c.Add(time.Day * time.Duration(i))
     //     c  = c.Add(time.Hour * 24)
     // 
     //     p(i,c)
     //     if um.UsMdVl.End_time.Equal(c) {
     //         //p("break i =",i)
     //         break
     //     } 
     //     // err = Gen_Lvl_Ti(byteValues,data,ps,um,Ds,c,);  __err_panic(err)
     //     // err = Init_Um_Ds(byteValues,data,TotalDict,um,c);  __err_panic(err)
     // 
     //     //    //!!!!!!!!!!!!!!!!!!!!
     // 
     // } // for i
     // p("##### c =",c)


    return  err

} // func alloc_ow

