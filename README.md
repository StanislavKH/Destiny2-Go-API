## Destiny2-Go-API wrapper

```
import(
    "fmt"
    rk "github.com/StanislavKH/Destiny2-Go-API/ratking"
)

//Developer API Key 
const(
    APIKey = "XXXXXXXXXXXXXXXXXXXXXXXXXXXXX"
)

func main() {
    api := rk.NewClient(nil, APIKey)
    data, err := api.Player.SearchDestinyPlayer("TigerPsn", "SearchByThisName")
    if err != nil {
        fmt.Println(err)
    }
    //DO SOMETHING with data
}
```

## Available membershipType list from Bungie.net

Type  | Value
----- | -----
All   | -1
None  |  0
TigerXbox | 1
TigerPsn | 2
TigerBlizzard | 3
TigerDemon | 10
BungieNext | 254

## Active endpoints
Method | Name | Params
------ | ---- | --------
GET | GetDestinyManifest | Nothing
GET | SearchDestinyPlayer| membershipType string, Name string
GET | GetHistoricalStatsForAccount| membershipType string, destinyMembershipId string
