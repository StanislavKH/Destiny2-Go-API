## Destiny2-Go-API wrapper

```
import(
    rk "github.com/StanislavKH/Destiny2-Go-API/ratking"
)

//Developer API Key 
const(
    APIKey = "XXXXXXXXXXXXXXXXXXXXXXXXXXXXX"
)

func main() {
    api := rk.NewClient(nil, APIKey)
    data, err := api.Player.SearchDestinyPlayer("psn", "SearchByThisName")
    if err != nil {
        return err
    }
    //DO SOMETHING with data
}
```

## Active endpoints
Method | Name | Params
------ | ---- | --------
GET | GetDestinyManifest | Nothing
GET | SearchDestinyPlayer| membershipType int, Name string
GET | GetHistoricalStatsForAccount| membershipType int, destinyMembershipId string
