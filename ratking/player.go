package ratking

import (
	"encoding/json"
	"fmt"
)

type PlayerService struct {
	client *Client
}

//GetDestinyManifest Returns the current version of the manifest as object
func (ps *PlayerService) GetDestinyManifest() (Manifest, error) {
	u := fmt.Sprintf("Destiny2/Manifest/")

	retData := Manifest{}
	response, err := ps.client.Platform.PlatformRequest("GET", u)
	if err != nil {
		return retData, err
	}
	err = json.Unmarshal(response, &retData)
	if err != nil {
		return retData, err
	}
	return retData, nil
}

//SearchDestinyPlayer Returns a list of Destiny memberships given a full Gamertag or PSN ID.
func (ps *PlayerService) SearchDestinyPlayer(membershipType string, playerName string) (ByPlayerNameResponse, error) {
	pfmType := BungieMembershipType[membershipType]
	u := fmt.Sprintf("Destiny2/SearchDestinyPlayer/%d/%s/", pfmType, playerName)

	retData := ByPlayerNameResponse{}
	response, err := ps.client.Platform.PlatformRequest("GET", u)
	if err != nil {
		return retData, err
	}
	err = json.Unmarshal(response, &retData)
	if err != nil {
		return retData, err
	}
	return retData, nil
}

//GetProfile Returns Destiny Profile information for the supplied membership.
func (ps *PlayerService) GetProfile(membershipType string, membershipID string) error {
	//pfmType := BungieMembershipType[membershipType]
	//u := fmt.Sprintf("Destiny2/%d/Profile/%s?components=200", pfmType, membershipID)
        u := fmt.Sprintf("Destiny2/Milestones")
	fmt.Println(u)

	//retData := &{}interface{}

	response, err := ps.client.Platform.PlatformRequest("GET", u)
	if err != nil {
		fmt.Println(err)

	}
	fmt.Println(string(response))
	return nil
}

func (ps *PlayerService) GetPublicMilestoneContent(milestoneHash string) (MilestoneContent, error) {
	u := fmt.Sprintf("Destiny2/Milestones/%s/Content/", milestoneHash)
	retData := MilestoneContent{}

	response, err := ps.client.Platform.PlatformRequest("GET", u)
	if err != nil {
		return retData, err
	}
	err = json.Unmarshal(response, &retData)
	if err != nil {
		return retData, err
	}
	return retData, nil
}


//GetPublicMilestones Gets public information about currently available Milestones.
func (ps *PlayerService) GetPublicMilestones() (Milestones, error) {
	u := fmt.Sprintf("Destiny2/Milestones")

	var incomingData map[string]*json.RawMessage
	var milestone Milestone
	var milestonesList Milestones

	response, err := ps.client.Platform.PlatformRequest("GET", u)
	if err != nil {
		return milestonesList, err
	}
	err = json.Unmarshal(response, &incomingData)
	if err != nil {
		return milestonesList, err
	}
	if len(incomingData) > 0 {
		for _,v := range incomingData {
			err = json.Unmarshal(*v, &milestone)
			if err != nil {
				return milestonesList, err
			}
			milestonesList = append(milestonesList, milestone)
		}
	}
	return milestonesList, nil
}

//GetClanWeeklyRewardState Returns information on the weekly clan rewards and if the clan has earned them or not.
//Note that this will always report rewards as not redeemed.
func (ps *PlayerService) GetClanWeeklyRewardState(clanID int) (ClanWeeklyRewardState, error) {
	u := fmt.Sprintf("Destiny2/Clan/%d/WeeklyRewardState/", clanID)

	retData := ClanWeeklyRewardState{}
	response, err := ps.client.Platform.PlatformRequest("GET", u)
	if err != nil {
		return retData,err
	}
	err = json.Unmarshal(response, &retData)
	if err != nil {
		return retData, err
	}
	return retData, nil
}

//GetHistoricalStatsDefinition Returns historical stats definitions.
func (ps *PlayerService) HistoricalStatsDefinition() (HistoricalStatsDefinition, error) {
	u := fmt.Sprintf("Destiny2/Stats/Definition/")

	retData := HistoricalStatsDefinition{}
	response, err := ps.client.Platform.PlatformRequest("GET", u)
	if err != nil {
		return retData, err
	}
	err = json.Unmarshal(response, &retData)
	if err != nil {
		return retData, err
	}
	return retData, nil
}

//GetHistoricalStatsForAccount Gets aggregate historical stats organized around each character for a given account.
//PREVIEW: This endpoint is still in beta, and may experience rough edges.
//The schema is in final form, but there may be bugs that prevent desirable operation.
func (ps *PlayerService) GetHistoricalStatsForAccount(membershipType string, membershipID string) (CommonAccountStats, error) {
	pfmType := BungieMembershipType[membershipType]
	u := fmt.Sprintf("Destiny2/%d/Account/%s/Stats/", pfmType, membershipID)

	retData := CommonAccountStats{}
	response, err := ps.client.Platform.PlatformRequest("GET", u)
	if err != nil {
		return retData, err
	}
	err = json.Unmarshal(response, &retData)
	if err != nil {
		return retData, err
	}
	return retData, nil
}
