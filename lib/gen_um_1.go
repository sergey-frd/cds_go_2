package lib

import (   
    "fmt"
    "os" 
    "github.com/boltdb/bolt"
    // "errors"
    "strconv"
    // "io/ioutil"
    "github.com/valyala/fastjson"

    //"bytes"

    //"runtime"
    "encoding/json"
    //"encoding/gob"
    
    "time"
    "math/rand"

    S "cds_go_1/config"
    //L "cds_go_1/lib"

)

//----------------------------------------------
func Gen_Um_Bucket(byteValues  []byte, 
                       bucket_Name string, 
                       data        map[string]map[string]string) error {


    var  nn        int
    var ID_User    string 
    var ID_Media    string 
    
    var User_Name  string 
    var Media_Cost string 
    var Media_Key  string 
    var MdVal     S.Md_VAL
    var um        S.User_Media_STC

    Clip_Budget_Min := fastjson.GetInt(byteValues, "Base", "Clip_Budget_Min")
    Clip_Budget_Max := fastjson.GetInt(byteValues, "Base", "Clip_Budget_Max")

    // Time_Add_Hour_Min := fastjson.GetInt(byteValues, "Base", "Time_Add_Hour_Min")
    // Time_Add_Hour_Max := fastjson.GetInt(byteValues, "Base", "Time_Add_Hour_Max")


    Time_Add_Day_Min := fastjson.GetInt(byteValues, "Base", "Time_Add_Day_Min")
    Time_Add_Day_Max := fastjson.GetInt(byteValues, "Base", "Time_Add_Day_Max")

    User_Counter := fastjson.GetInt(byteValues, "Base", "User_Counter")
    //fmt.Printf("User_Counter = %d\n", User_Counter)

    len_data_User := len(data["User"])
    //fmt.Printf("len_data_User = %d\n", len_data_User)


    // !!!!!!!!!!!!!!!!!!!!!!!!!
    len_data_User =  User_Counter

    keys_User := []string{}
    for k := range data["User"] {
        keys_User = append(keys_User, k)
    }
    //fmt.Println("keys_User =",keys_User)

    rand.Seed(int64(time.Now().Nanosecond()))
    rand_User_Key := rand.Intn(len_data_User)
    //fmt.Println("rand_User_Key =",rand_User_Key)
    ID_User = keys_User[rand_User_Key]
    //fmt.Println("ID_User =",ID_User)

    User_Name = data["User"][ID_User]
    //fmt.Println("User_Name =",User_Name)




    len_data_Media := len(data["Media"])
    //fmt.Printf("len_data_Media = %d\n", len_data_Media)
    keys_Media := []string{}
    for k := range data["Media"] {
        keys_Media = append(keys_Media, k)
    }
    //fmt.Println("keys_Media =",keys_Media)



    dbFileName, err := GetDbName(byteValues);  __err_panic(err) 
    if _, err := os.Stat(dbFileName); os.IsNotExist(err) {
        __err_panic(err)

    }
    //fmt.Println(" ----------  L.Buckets Stat IsNotExist ---------- ")
    
    db, err := bolt.Open(dbFileName, 0600, nil)
    if err != nil {
        __err_panic(err)
    }
    defer db.Close()

    err = db.Update(func(tx *bolt.Tx) error {

        b, err := tx.CreateBucketIfNotExists([]byte(bucket_Name))
        __err_panic(err)
        //fmt.Println(  "bucket_Name b =",bucket_Name,b)

        nn = 0
        for key, _ := range data["User"] {
            nn  += 1
            //fmt.Println(nn, "Key =", key, "Value =", value)
            //fmt.Println(nn, "Key =", key)

            ID_User   = key
            User_Name = data["User"][key]
            //fmt.Println("User_Name =",User_Name)


            rand.Seed(int64(time.Now().Nanosecond()))
            rand_Media_Key := rand.Intn(len_data_Media) 
            
            Media_Key = keys_Media[rand_Media_Key]
            //ID_Media = strconv.Itoa(Media_Key)
            ID_Media = Media_Key
            //fmt.Println("ID_Media =",ID_Media)


            Media_Dict_Val := data["Media"][Media_Key]
            //fmt.Println("Media_Dict_Val =",Media_Dict_Val)


            byt := []byte(Media_Dict_Val)
            err = json.Unmarshal(byt, &MdVal)
            if err != nil {
                fmt.Println("There was an error:", err)
            }

            Media_Cost = strconv.Itoa(Random(Clip_Budget_Min, Clip_Budget_Max)*1000)

            um.UsMd.ID_User  = ID_User
            um.UsMd.ID_Media = ID_Media

            um.UsMdVl.Media_Name  = User_Name +"_"+ MdVal.Type_Media +"_"+ ID_Media
            um.UsMdVl.Media_Cost  = Media_Cost
            um.UsMdVl.Media_Slots = ID_Media


            currentTime := time.Now()   

            // t := time.Now()
            t := time.Date(
                currentTime.Year() ,
                currentTime.Month(),
                currentTime.Day()  ,
                0   ,
                0   ,
                0   ,
                0   , 
                time.UTC,
                )

            //um.UsMdVl.Start_time  = t.Add(time.Hour * 24 * Random(1,5))
            //um.UsMdVl.End_time    = t.Add(time.Hour * 24 * Random(6,10))


            //t_Min := Random(48,Time_Add_Hour_Min)
            //t_Max := Random(Time_Add_Hour_Max - t_Min + 24, Time_Add_Hour_Max)

            t_Min := Random(5,Time_Add_Day_Min)
            t_Max := Random(Time_Add_Day_Max - t_Min + 1, Time_Add_Day_Max)

            um.UsMdVl.Start_time  = t.Add(time.Hour * time.Duration(t_Min)*24)
            um.UsMdVl.End_time    = t.Add(time.Hour * time.Duration(t_Max)*24)

            // fmt.Println("um =",um)

            encoded, err := json.Marshal(um.UsMd)
            __err_panic(err)

            enc_val, err := json.Marshal(um.UsMdVl)
            __err_panic(err)

            err = b.Put([]byte(encoded), []byte(enc_val))
            __err_panic(err)


            if nn >= User_Counter {
                break	
            }
        }


        return err

    })  // err = db.Update(func(tx *bolt.Tx) error 

    return err

}


