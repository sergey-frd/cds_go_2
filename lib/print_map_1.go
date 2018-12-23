package lib

import (   
    "fmt"
//    "log"
    "sort"
//    "github.com/valyala/fastjson"
    //"github.com/boltdb/bolt"
    //"os" 
//    "errors"

    // "github.com/tealeg/xlsx"
    // "io/ioutil"

    //"encoding/gob"

    //L "cds_go_2/lib"

    //"encoding/json"

    // "time"
    // "math/rand"

    //    "path/filepath"
    //    L "cds_go_2/lib"
    S "cds_go_2/config"

)

//   TotalDict  map[string]map[string]float64,

//----------------------------------------------
func Print_Map_Table(
    json_go  S.AutoGenerated,
    data       map[string]map[string]string,
    t          string,
    )  (err error) {


    p := fmt.Println

    //p("json_go.DBTableList=", json_go.DBTableList)

    var keys []string
    //FLAG_MAP_PRINT_LEN_ONLY  := fastjson.GetInt(byteValues, "Base", "FLAG_MAP_PRINT_LEN_ONLY")
    FLAG_MAP_PRINT  := json_go.Base.FLAGMAPPRINT 
    //p("result[Base][project_name] =", result["Base"]["project_name"])


    for k, _ := range data[t] {
        keys = append(keys, string(k))
    } // for k, v
    sort.Strings(keys)  // or sort.Ints(keys), sort.Sort(...), etc., per <K>
    for i, k := range keys  {
        v := data[t][k]

        if FLAG_MAP_PRINT == 1 { p(t,i,k,"=>", v) }

    } // for 
    p(t,"len =",len(data[t])) 


    return err

} // func                       
//----------------------------------------------
func Print_Map(
    json_go  S.AutoGenerated,
    data     map[string]map[string]string,
    )  (err error) {


    p := fmt.Println

    //p("json_go.DBTableList=", json_go.DBTableList)

    var keys []string
    //FLAG_MAP_PRINT_LEN_ONLY  := fastjson.GetInt(byteValues, "Base", "FLAG_MAP_PRINT_LEN_ONLY")
    FLAG_MAP_PRINT_LEN_ONLY  := json_go.Base.FLAGMAPPRINTLENONLY
    //p("result[Base][project_name] =", result["Base"]["project_name"])


    for _, t := range json_go.DBTableList{
        //p("t=", t)

        for k, _ := range data[t] {
            keys = append(keys, string(k))
        } // for k, v
        sort.Strings(keys)  // or sort.Ints(keys), sort.Sort(...), etc., per <K>
        for i, k := range keys  {
            v := data[t][k]

            if FLAG_MAP_PRINT_LEN_ONLY != 1 { p(t,i,k,"=>", v) }

        }
        p("len",t,"=",len(data[t])) 

    } // for 

    return err

} // func                       

