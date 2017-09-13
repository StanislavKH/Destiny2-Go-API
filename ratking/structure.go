package ratking

import(
    "time"
)

type Milestones []Milestone
type Milestone struct {
	MilestoneHash   int64 `json:"milestoneHash"`
	AvailableQuests []AvailableQuest `json:"availableQuests"`
	StartDate time.Time `json:"startDate"`
	EndDate   time.Time `json:"endDate"`
}
type AvailableQuest struct {
	QuestItemHash int64 `json:"questItemHash"`
	Activity      struct {
	ActivityHash int64 `json:"activityHash"`
		Variants     []struct {
			ActivityHash int64 `json:"activityHash"`
		} `json:"variants"`
	} `json:"activity"`
}

type MilestoneContent struct {
	About  string   `json:"about"`
	Status string   `json:"status"`
	Tips   []string `json:"tips"`
}

type ByPlayerNameResponse []struct {
	DisplayName    string `json:"displayName"`
	IconPath       string `json:"iconPath"`
	MembershipID   string `json:"membershipId"`
	MembershipType int    `json:"membershipType"`
}

type CommonAccountStats struct {
	MergedDeletedCharacters struct {
		Results struct {
		} `json:"results"`
		Merged struct {
		} `json:"merged"`
	} `json:"mergedDeletedCharacters"`
	MergedAllCharacters struct {
		Results struct {
			AllPvE struct {
				AllTime struct {
					ActivitiesCleared struct {
						StatID string `json:"statId"`
						Basic  struct {
							Value        float64 `json:"value"`
							DisplayValue string  `json:"displayValue"`
						} `json:"basic"`
					} `json:"activitiesCleared"`
					WeaponKillsSuper struct {
						StatID string `json:"statId"`
						Basic  struct {
							Value        float64 `json:"value"`
							DisplayValue string  `json:"displayValue"`
						} `json:"basic"`
						Pga struct {
							Value        float64 `json:"value"`
							DisplayValue string  `json:"displayValue"`
						} `json:"pga"`
					} `json:"weaponKillsSuper"`
					ActivitiesEntered struct {
						StatID string `json:"statId"`
						Basic  struct {
							Value        float64 `json:"value"`
							DisplayValue string  `json:"displayValue"`
						} `json:"basic"`
					} `json:"activitiesEntered"`
					WeaponKillsMelee struct {
						StatID string `json:"statId"`
						Basic  struct {
							Value        float64 `json:"value"`
							DisplayValue string  `json:"displayValue"`
						} `json:"basic"`
						Pga struct {
							Value        float64 `json:"value"`
							DisplayValue string  `json:"displayValue"`
						} `json:"pga"`
					} `json:"weaponKillsMelee"`
					WeaponKillsGrenade struct {
						StatID string `json:"statId"`
						Basic  struct {
							Value        float64 `json:"value"`
							DisplayValue string  `json:"displayValue"`
						} `json:"basic"`
						Pga struct {
							Value        float64 `json:"value"`
							DisplayValue string  `json:"displayValue"`
						} `json:"pga"`
					} `json:"weaponKillsGrenade"`
					AbilityKills struct {
						StatID string `json:"statId"`
						Basic  struct {
							Value        float64 `json:"value"`
							DisplayValue string  `json:"displayValue"`
						} `json:"basic"`
						Pga struct {
							Value        float64 `json:"value"`
							DisplayValue string  `json:"displayValue"`
						} `json:"pga"`
					} `json:"abilityKills"`
					Assists struct {
						StatID string `json:"statId"`
						Basic  struct {
							Value        float64 `json:"value"`
							DisplayValue string  `json:"displayValue"`
						} `json:"basic"`
						Pga struct {
							Value        float64 `json:"value"`
							DisplayValue string  `json:"displayValue"`
						} `json:"pga"`
					} `json:"assists"`
					TotalDeathDistance struct {
						StatID string `json:"statId"`
						Basic  struct {
							Value        float64 `json:"value"`
							DisplayValue string  `json:"displayValue"`
						} `json:"basic"`
					} `json:"totalDeathDistance"`
					AverageDeathDistance struct {
						StatID string `json:"statId"`
						Basic  struct {
							Value        float64 `json:"value"`
							DisplayValue string  `json:"displayValue"`
						} `json:"basic"`
					} `json:"averageDeathDistance"`
					TotalKillDistance struct {
						StatID string `json:"statId"`
						Basic  struct {
							Value        float64 `json:"value"`
							DisplayValue string  `json:"displayValue"`
						} `json:"basic"`
					} `json:"totalKillDistance"`
					Kills struct {
						StatID string `json:"statId"`
						Basic  struct {
							Value        float64 `json:"value"`
							DisplayValue string  `json:"displayValue"`
						} `json:"basic"`
						Pga struct {
							Value        float64 `json:"value"`
							DisplayValue string  `json:"displayValue"`
						} `json:"pga"`
					} `json:"kills"`
					AverageKillDistance struct {
						StatID string `json:"statId"`
						Basic  struct {
							Value        float64 `json:"value"`
							DisplayValue string  `json:"displayValue"`
						} `json:"basic"`
					} `json:"averageKillDistance"`
					SecondsPlayed struct {
						StatID string `json:"statId"`
						Basic  struct {
							Value        float64 `json:"value"`
							DisplayValue string  `json:"displayValue"`
						} `json:"basic"`
						Pga struct {
							Value        float64 `json:"value"`
							DisplayValue string  `json:"displayValue"`
						} `json:"pga"`
					} `json:"secondsPlayed"`
					Deaths struct {
						StatID string `json:"statId"`
						Basic  struct {
							Value        float64 `json:"value"`
							DisplayValue string  `json:"displayValue"`
						} `json:"basic"`
						Pga struct {
							Value        float64 `json:"value"`
							DisplayValue string  `json:"displayValue"`
						} `json:"pga"`
					} `json:"deaths"`
					AverageLifespan struct {
						StatID string `json:"statId"`
						Basic  struct {
							Value        float64 `json:"value"`
							DisplayValue string  `json:"displayValue"`
						} `json:"basic"`
					} `json:"averageLifespan"`
					BestSingleGameKills struct {
						StatID string `json:"statId"`
						Basic  struct {
							Value        float64 `json:"value"`
							DisplayValue string  `json:"displayValue"`
						} `json:"basic"`
					} `json:"bestSingleGameKills"`
					KillsDeathsRatio struct {
						StatID string `json:"statId"`
						Basic  struct {
							Value        float64 `json:"value"`
							DisplayValue string  `json:"displayValue"`
						} `json:"basic"`
					} `json:"killsDeathsRatio"`
					KillsDeathsAssists struct {
						StatID string `json:"statId"`
						Basic  struct {
							Value        float64 `json:"value"`
							DisplayValue string  `json:"displayValue"`
						} `json:"basic"`
					} `json:"killsDeathsAssists"`
					ObjectivesCompleted struct {
						StatID string `json:"statId"`
						Basic  struct {
							Value        float64 `json:"value"`
							DisplayValue string  `json:"displayValue"`
						} `json:"basic"`
						Pga struct {
							Value        float64 `json:"value"`
							DisplayValue string  `json:"displayValue"`
						} `json:"pga"`
					} `json:"objectivesCompleted"`
					PrecisionKills struct {
						StatID string `json:"statId"`
						Basic  struct {
							Value        float64 `json:"value"`
							DisplayValue string  `json:"displayValue"`
						} `json:"basic"`
						Pga struct {
							Value        float64 `json:"value"`
							DisplayValue string  `json:"displayValue"`
						} `json:"pga"`
					} `json:"precisionKills"`
					ResurrectionsPerformed struct {
						StatID string `json:"statId"`
						Basic  struct {
							Value        float64 `json:"value"`
							DisplayValue string  `json:"displayValue"`
						} `json:"basic"`
						Pga struct {
							Value        float64 `json:"value"`
							DisplayValue string  `json:"displayValue"`
						} `json:"pga"`
					} `json:"resurrectionsPerformed"`
					ResurrectionsReceived struct {
						StatID string `json:"statId"`
						Basic  struct {
							Value        float64 `json:"value"`
							DisplayValue string  `json:"displayValue"`
						} `json:"basic"`
						Pga struct {
							Value        float64 `json:"value"`
							DisplayValue string  `json:"displayValue"`
						} `json:"pga"`
					} `json:"resurrectionsReceived"`
					Suicides struct {
						StatID string `json:"statId"`
						Basic  struct {
							Value        float64 `json:"value"`
							DisplayValue string  `json:"displayValue"`
						} `json:"basic"`
						Pga struct {
							Value        float64 `json:"value"`
							DisplayValue string  `json:"displayValue"`
						} `json:"pga"`
					} `json:"suicides"`
					WeaponKillsAutoRifle struct {
						StatID string `json:"statId"`
						Basic  struct {
							Value        float64 `json:"value"`
							DisplayValue string  `json:"displayValue"`
						} `json:"basic"`
						Pga struct {
							Value        float64 `json:"value"`
							DisplayValue string  `json:"displayValue"`
						} `json:"pga"`
					} `json:"weaponKillsAutoRifle"`
					WeaponKillsFusionRifle struct {
						StatID string `json:"statId"`
						Basic  struct {
							Value        float64 `json:"value"`
							DisplayValue string  `json:"displayValue"`
						} `json:"basic"`
						Pga struct {
							Value        float64 `json:"value"`
							DisplayValue string  `json:"displayValue"`
						} `json:"pga"`
					} `json:"weaponKillsFusionRifle"`
					WeaponKillsHandCannon struct {
						StatID string `json:"statId"`
						Basic  struct {
							Value        float64 `json:"value"`
							DisplayValue string  `json:"displayValue"`
						} `json:"basic"`
						Pga struct {
							Value        float64 `json:"value"`
							DisplayValue string  `json:"displayValue"`
						} `json:"pga"`
					} `json:"weaponKillsHandCannon"`
					WeaponKillsMachinegun struct {
						StatID string `json:"statId"`
						Basic  struct {
							Value        float64 `json:"value"`
							DisplayValue string  `json:"displayValue"`
						} `json:"basic"`
						Pga struct {
							Value        float64 `json:"value"`
							DisplayValue string  `json:"displayValue"`
						} `json:"pga"`
					} `json:"weaponKillsMachinegun"`
					WeaponKillsPulseRifle struct {
						StatID string `json:"statId"`
						Basic  struct {
							Value        float64 `json:"value"`
							DisplayValue string  `json:"displayValue"`
						} `json:"basic"`
						Pga struct {
							Value        float64 `json:"value"`
							DisplayValue string  `json:"displayValue"`
						} `json:"pga"`
					} `json:"weaponKillsPulseRifle"`
					WeaponKillsRocketLauncher struct {
						StatID string `json:"statId"`
						Basic  struct {
							Value        float64 `json:"value"`
							DisplayValue string  `json:"displayValue"`
						} `json:"basic"`
						Pga struct {
							Value        float64 `json:"value"`
							DisplayValue string  `json:"displayValue"`
						} `json:"pga"`
					} `json:"weaponKillsRocketLauncher"`
					WeaponKillsScoutRifle struct {
						StatID string `json:"statId"`
						Basic  struct {
							Value        float64 `json:"value"`
							DisplayValue string  `json:"displayValue"`
						} `json:"basic"`
						Pga struct {
							Value        float64 `json:"value"`
							DisplayValue string  `json:"displayValue"`
						} `json:"pga"`
					} `json:"weaponKillsScoutRifle"`
					WeaponKillsShotgun struct {
						StatID string `json:"statId"`
						Basic  struct {
							Value        float64 `json:"value"`
							DisplayValue string  `json:"displayValue"`
						} `json:"basic"`
						Pga struct {
							Value        float64 `json:"value"`
							DisplayValue string  `json:"displayValue"`
						} `json:"pga"`
					} `json:"weaponKillsShotgun"`
					WeaponKillsSniper struct {
						StatID string `json:"statId"`
						Basic  struct {
							Value        float64 `json:"value"`
							DisplayValue string  `json:"displayValue"`
						} `json:"basic"`
						Pga struct {
							Value        float64 `json:"value"`
							DisplayValue string  `json:"displayValue"`
						} `json:"pga"`
					} `json:"weaponKillsSniper"`
					WeaponKillsSubmachinegun struct {
						StatID string `json:"statId"`
						Basic  struct {
							Value        float64 `json:"value"`
							DisplayValue string  `json:"displayValue"`
						} `json:"basic"`
						Pga struct {
							Value        float64 `json:"value"`
							DisplayValue string  `json:"displayValue"`
						} `json:"pga"`
					} `json:"weaponKillsSubmachinegun"`
					WeaponKillsRelic struct {
						StatID string `json:"statId"`
						Basic  struct {
							Value        float64 `json:"value"`
							DisplayValue string  `json:"displayValue"`
						} `json:"basic"`
						Pga struct {
							Value        float64 `json:"value"`
							DisplayValue string  `json:"displayValue"`
						} `json:"pga"`
					} `json:"weaponKillsRelic"`
					WeaponKillsSideArm struct {
						StatID string `json:"statId"`
						Basic  struct {
							Value        float64 `json:"value"`
							DisplayValue string  `json:"displayValue"`
						} `json:"basic"`
						Pga struct {
							Value        float64 `json:"value"`
							DisplayValue string  `json:"displayValue"`
						} `json:"pga"`
					} `json:"weaponKillsSideArm"`
					WeaponKillsSword struct {
						StatID string `json:"statId"`
						Basic  struct {
							Value        float64 `json:"value"`
							DisplayValue string  `json:"displayValue"`
						} `json:"basic"`
						Pga struct {
							Value        float64 `json:"value"`
							DisplayValue string  `json:"displayValue"`
						} `json:"pga"`
					} `json:"weaponKillsSword"`
					WeaponKillsAbility struct {
						StatID string `json:"statId"`
						Basic  struct {
							Value        float64 `json:"value"`
							DisplayValue string  `json:"displayValue"`
						} `json:"basic"`
						Pga struct {
							Value        float64 `json:"value"`
							DisplayValue string  `json:"displayValue"`
						} `json:"pga"`
					} `json:"weaponKillsAbility"`
					WeaponBestType struct {
						StatID string `json:"statId"`
						Basic  struct {
							Value        float64 `json:"value"`
							DisplayValue string  `json:"displayValue"`
						} `json:"basic"`
					} `json:"weaponBestType"`
					AllParticipantsCount struct {
						StatID string `json:"statId"`
						Basic  struct {
							Value        float64 `json:"value"`
							DisplayValue string  `json:"displayValue"`
						} `json:"basic"`
					} `json:"allParticipantsCount"`
					AllParticipantsTimePlayed struct {
						StatID string `json:"statId"`
						Basic  struct {
							Value        float64 `json:"value"`
							DisplayValue string  `json:"displayValue"`
						} `json:"basic"`
					} `json:"allParticipantsTimePlayed"`
					LongestKillSpree struct {
						StatID string `json:"statId"`
						Basic  struct {
							Value        float64 `json:"value"`
							DisplayValue string  `json:"displayValue"`
						} `json:"basic"`
					} `json:"longestKillSpree"`
					LongestSingleLife struct {
						StatID string `json:"statId"`
						Basic  struct {
							Value        float64 `json:"value"`
							DisplayValue string  `json:"displayValue"`
						} `json:"basic"`
					} `json:"longestSingleLife"`
					MostPrecisionKills struct {
						StatID string `json:"statId"`
						Basic  struct {
							Value        float64 `json:"value"`
							DisplayValue string  `json:"displayValue"`
						} `json:"basic"`
					} `json:"mostPrecisionKills"`
					OrbsDropped struct {
						StatID string `json:"statId"`
						Basic  struct {
							Value        float64 `json:"value"`
							DisplayValue string  `json:"displayValue"`
						} `json:"basic"`
						Pga struct {
							Value        float64 `json:"value"`
							DisplayValue string  `json:"displayValue"`
						} `json:"pga"`
					} `json:"orbsDropped"`
					OrbsGathered struct {
						StatID string `json:"statId"`
						Basic  struct {
							Value        float64 `json:"value"`
							DisplayValue string  `json:"displayValue"`
						} `json:"basic"`
						Pga struct {
							Value        float64 `json:"value"`
							DisplayValue string  `json:"displayValue"`
						} `json:"pga"`
					} `json:"orbsGathered"`
					PublicEventsCompleted struct {
						StatID string `json:"statId"`
						Basic  struct {
							Value        float64 `json:"value"`
							DisplayValue string  `json:"displayValue"`
						} `json:"basic"`
						Pga struct {
							Value        float64 `json:"value"`
							DisplayValue string  `json:"displayValue"`
						} `json:"pga"`
					} `json:"publicEventsCompleted"`
					RemainingTimeAfterQuitSeconds struct {
						StatID string `json:"statId"`
						Basic  struct {
							Value        float64 `json:"value"`
							DisplayValue string  `json:"displayValue"`
						} `json:"basic"`
						Pga struct {
							Value        float64 `json:"value"`
							DisplayValue string  `json:"displayValue"`
						} `json:"pga"`
					} `json:"remainingTimeAfterQuitSeconds"`
					TotalActivityDurationSeconds struct {
						StatID string `json:"statId"`
						Basic  struct {
							Value        float64 `json:"value"`
							DisplayValue string  `json:"displayValue"`
						} `json:"basic"`
						Pga struct {
							Value        float64 `json:"value"`
							DisplayValue string  `json:"displayValue"`
						} `json:"pga"`
					} `json:"totalActivityDurationSeconds"`
					FastestCompletionMs struct {
						StatID string `json:"statId"`
						Basic  struct {
							Value        float64 `json:"value"`
							DisplayValue string  `json:"displayValue"`
						} `json:"basic"`
					} `json:"fastestCompletionMs"`
					LongestKillDistance struct {
						StatID string `json:"statId"`
						Basic  struct {
							Value        float64 `json:"value"`
							DisplayValue string  `json:"displayValue"`
						} `json:"basic"`
					} `json:"longestKillDistance"`
					HighestCharacterLevel struct {
						StatID string `json:"statId"`
						Basic  struct {
							Value        float64 `json:"value"`
							DisplayValue string  `json:"displayValue"`
						} `json:"basic"`
					} `json:"highestCharacterLevel"`
					HighestLightLevel struct {
						StatID string `json:"statId"`
						Basic  struct {
							Value        float64 `json:"value"`
							DisplayValue string  `json:"displayValue"`
						} `json:"basic"`
					} `json:"highestLightLevel"`
				} `json:"allTime"`
			} `json:"allPvE"`
			AllPvP struct {
				AllTime struct {
					WeaponKillsSuper struct {
						StatID string `json:"statId"`
						Basic  struct {
							Value        float64 `json:"value"`
							DisplayValue string  `json:"displayValue"`
						} `json:"basic"`
						Pga struct {
							Value        float64 `json:"value"`
							DisplayValue string  `json:"displayValue"`
						} `json:"pga"`
					} `json:"weaponKillsSuper"`
					ActivitiesEntered struct {
						StatID string `json:"statId"`
						Basic  struct {
							Value        float64 `json:"value"`
							DisplayValue string  `json:"displayValue"`
						} `json:"basic"`
					} `json:"activitiesEntered"`
					WeaponKillsMelee struct {
						StatID string `json:"statId"`
						Basic  struct {
							Value        float64 `json:"value"`
							DisplayValue string  `json:"displayValue"`
						} `json:"basic"`
						Pga struct {
							Value        float64 `json:"value"`
							DisplayValue string  `json:"displayValue"`
						} `json:"pga"`
					} `json:"weaponKillsMelee"`
					WeaponKillsGrenade struct {
						StatID string `json:"statId"`
						Basic  struct {
							Value        float64 `json:"value"`
							DisplayValue string  `json:"displayValue"`
						} `json:"basic"`
						Pga struct {
							Value        float64 `json:"value"`
							DisplayValue string  `json:"displayValue"`
						} `json:"pga"`
					} `json:"weaponKillsGrenade"`
					AbilityKills struct {
						StatID string `json:"statId"`
						Basic  struct {
							Value        float64 `json:"value"`
							DisplayValue string  `json:"displayValue"`
						} `json:"basic"`
						Pga struct {
							Value        float64 `json:"value"`
							DisplayValue string  `json:"displayValue"`
						} `json:"pga"`
					} `json:"abilityKills"`
					ActivitiesWon struct {
						StatID string `json:"statId"`
						Basic  struct {
							Value        float64 `json:"value"`
							DisplayValue string  `json:"displayValue"`
						} `json:"basic"`
					} `json:"activitiesWon"`
					Assists struct {
						StatID string `json:"statId"`
						Basic  struct {
							Value        float64 `json:"value"`
							DisplayValue string  `json:"displayValue"`
						} `json:"basic"`
						Pga struct {
							Value        float64 `json:"value"`
							DisplayValue string  `json:"displayValue"`
						} `json:"pga"`
					} `json:"assists"`
					TotalDeathDistance struct {
						StatID string `json:"statId"`
						Basic  struct {
							Value        float64 `json:"value"`
							DisplayValue string  `json:"displayValue"`
						} `json:"basic"`
					} `json:"totalDeathDistance"`
					AverageDeathDistance struct {
						StatID string `json:"statId"`
						Basic  struct {
							Value        float64 `json:"value"`
							DisplayValue string  `json:"displayValue"`
						} `json:"basic"`
					} `json:"averageDeathDistance"`
					TotalKillDistance struct {
						StatID string `json:"statId"`
						Basic  struct {
							Value        float64 `json:"value"`
							DisplayValue string  `json:"displayValue"`
						} `json:"basic"`
					} `json:"totalKillDistance"`
					Kills struct {
						StatID string `json:"statId"`
						Basic  struct {
							Value        float64 `json:"value"`
							DisplayValue string  `json:"displayValue"`
						} `json:"basic"`
						Pga struct {
							Value        float64 `json:"value"`
							DisplayValue string  `json:"displayValue"`
						} `json:"pga"`
					} `json:"kills"`
					AverageKillDistance struct {
						StatID string `json:"statId"`
						Basic  struct {
							Value        float64 `json:"value"`
							DisplayValue string  `json:"displayValue"`
						} `json:"basic"`
					} `json:"averageKillDistance"`
					SecondsPlayed struct {
						StatID string `json:"statId"`
						Basic  struct {
							Value        float64 `json:"value"`
							DisplayValue string  `json:"displayValue"`
						} `json:"basic"`
						Pga struct {
							Value        float64 `json:"value"`
							DisplayValue string  `json:"displayValue"`
						} `json:"pga"`
					} `json:"secondsPlayed"`
					Deaths struct {
						StatID string `json:"statId"`
						Basic  struct {
							Value        float64 `json:"value"`
							DisplayValue string  `json:"displayValue"`
						} `json:"basic"`
						Pga struct {
							Value        float64 `json:"value"`
							DisplayValue string  `json:"displayValue"`
						} `json:"pga"`
					} `json:"deaths"`
					AverageLifespan struct {
						StatID string `json:"statId"`
						Basic  struct {
							Value        float64 `json:"value"`
							DisplayValue string  `json:"displayValue"`
						} `json:"basic"`
					} `json:"averageLifespan"`
					Score struct {
						StatID string `json:"statId"`
						Basic  struct {
							Value        float64 `json:"value"`
							DisplayValue string  `json:"displayValue"`
						} `json:"basic"`
						Pga struct {
							Value        float64 `json:"value"`
							DisplayValue string  `json:"displayValue"`
						} `json:"pga"`
					} `json:"score"`
					AverageScorePerKill struct {
						StatID string `json:"statId"`
						Basic  struct {
							Value        float64 `json:"value"`
							DisplayValue string  `json:"displayValue"`
						} `json:"basic"`
					} `json:"averageScorePerKill"`
					AverageScorePerLife struct {
						StatID string `json:"statId"`
						Basic  struct {
							Value        float64 `json:"value"`
							DisplayValue string  `json:"displayValue"`
						} `json:"basic"`
					} `json:"averageScorePerLife"`
					BestSingleGameKills struct {
						StatID string `json:"statId"`
						Basic  struct {
							Value        float64 `json:"value"`
							DisplayValue string  `json:"displayValue"`
						} `json:"basic"`
					} `json:"bestSingleGameKills"`
					BestSingleGameScore struct {
						StatID string `json:"statId"`
						Basic  struct {
							Value        float64 `json:"value"`
							DisplayValue string  `json:"displayValue"`
						} `json:"basic"`
					} `json:"bestSingleGameScore"`
					KillsDeathsRatio struct {
						StatID string `json:"statId"`
						Basic  struct {
							Value        float64 `json:"value"`
							DisplayValue string  `json:"displayValue"`
						} `json:"basic"`
					} `json:"killsDeathsRatio"`
					KillsDeathsAssists struct {
						StatID string `json:"statId"`
						Basic  struct {
							Value        float64 `json:"value"`
							DisplayValue string  `json:"displayValue"`
						} `json:"basic"`
					} `json:"killsDeathsAssists"`
					ObjectivesCompleted struct {
						StatID string `json:"statId"`
						Basic  struct {
							Value        float64 `json:"value"`
							DisplayValue string  `json:"displayValue"`
						} `json:"basic"`
						Pga struct {
							Value        float64 `json:"value"`
							DisplayValue string  `json:"displayValue"`
						} `json:"pga"`
					} `json:"objectivesCompleted"`
					PrecisionKills struct {
						StatID string `json:"statId"`
						Basic  struct {
							Value        float64 `json:"value"`
							DisplayValue string  `json:"displayValue"`
						} `json:"basic"`
						Pga struct {
							Value        float64 `json:"value"`
							DisplayValue string  `json:"displayValue"`
						} `json:"pga"`
					} `json:"precisionKills"`
					ResurrectionsPerformed struct {
						StatID string `json:"statId"`
						Basic  struct {
							Value        float64 `json:"value"`
							DisplayValue string  `json:"displayValue"`
						} `json:"basic"`
						Pga struct {
							Value        float64 `json:"value"`
							DisplayValue string  `json:"displayValue"`
						} `json:"pga"`
					} `json:"resurrectionsPerformed"`
					ResurrectionsReceived struct {
						StatID string `json:"statId"`
						Basic  struct {
							Value        float64 `json:"value"`
							DisplayValue string  `json:"displayValue"`
						} `json:"basic"`
						Pga struct {
							Value        float64 `json:"value"`
							DisplayValue string  `json:"displayValue"`
						} `json:"pga"`
					} `json:"resurrectionsReceived"`
					Suicides struct {
						StatID string `json:"statId"`
						Basic  struct {
							Value        float64 `json:"value"`
							DisplayValue string  `json:"displayValue"`
						} `json:"basic"`
						Pga struct {
							Value        float64 `json:"value"`
							DisplayValue string  `json:"displayValue"`
						} `json:"pga"`
					} `json:"suicides"`
					WeaponKillsAutoRifle struct {
						StatID string `json:"statId"`
						Basic  struct {
							Value        float64 `json:"value"`
							DisplayValue string  `json:"displayValue"`
						} `json:"basic"`
						Pga struct {
							Value        float64 `json:"value"`
							DisplayValue string  `json:"displayValue"`
						} `json:"pga"`
					} `json:"weaponKillsAutoRifle"`
					WeaponKillsFusionRifle struct {
						StatID string `json:"statId"`
						Basic  struct {
							Value        float64 `json:"value"`
							DisplayValue string  `json:"displayValue"`
						} `json:"basic"`
						Pga struct {
							Value        float64 `json:"value"`
							DisplayValue string  `json:"displayValue"`
						} `json:"pga"`
					} `json:"weaponKillsFusionRifle"`
					WeaponKillsHandCannon struct {
						StatID string `json:"statId"`
						Basic  struct {
							Value        float64 `json:"value"`
							DisplayValue string  `json:"displayValue"`
						} `json:"basic"`
						Pga struct {
							Value        float64 `json:"value"`
							DisplayValue string  `json:"displayValue"`
						} `json:"pga"`
					} `json:"weaponKillsHandCannon"`
					WeaponKillsMachinegun struct {
						StatID string `json:"statId"`
						Basic  struct {
							Value        float64 `json:"value"`
							DisplayValue string  `json:"displayValue"`
						} `json:"basic"`
						Pga struct {
							Value        float64 `json:"value"`
							DisplayValue string  `json:"displayValue"`
						} `json:"pga"`
					} `json:"weaponKillsMachinegun"`
					WeaponKillsPulseRifle struct {
						StatID string `json:"statId"`
						Basic  struct {
							Value        float64 `json:"value"`
							DisplayValue string  `json:"displayValue"`
						} `json:"basic"`
						Pga struct {
							Value        float64 `json:"value"`
							DisplayValue string  `json:"displayValue"`
						} `json:"pga"`
					} `json:"weaponKillsPulseRifle"`
					WeaponKillsRocketLauncher struct {
						StatID string `json:"statId"`
						Basic  struct {
							Value        float64 `json:"value"`
							DisplayValue string  `json:"displayValue"`
						} `json:"basic"`
						Pga struct {
							Value        float64 `json:"value"`
							DisplayValue string  `json:"displayValue"`
						} `json:"pga"`
					} `json:"weaponKillsRocketLauncher"`
					WeaponKillsScoutRifle struct {
						StatID string `json:"statId"`
						Basic  struct {
							Value        float64 `json:"value"`
							DisplayValue string  `json:"displayValue"`
						} `json:"basic"`
						Pga struct {
							Value        float64 `json:"value"`
							DisplayValue string  `json:"displayValue"`
						} `json:"pga"`
					} `json:"weaponKillsScoutRifle"`
					WeaponKillsShotgun struct {
						StatID string `json:"statId"`
						Basic  struct {
							Value        float64 `json:"value"`
							DisplayValue string  `json:"displayValue"`
						} `json:"basic"`
						Pga struct {
							Value        float64 `json:"value"`
							DisplayValue string  `json:"displayValue"`
						} `json:"pga"`
					} `json:"weaponKillsShotgun"`
					WeaponKillsSniper struct {
						StatID string `json:"statId"`
						Basic  struct {
							Value        float64 `json:"value"`
							DisplayValue string  `json:"displayValue"`
						} `json:"basic"`
						Pga struct {
							Value        float64 `json:"value"`
							DisplayValue string  `json:"displayValue"`
						} `json:"pga"`
					} `json:"weaponKillsSniper"`
					WeaponKillsSubmachinegun struct {
						StatID string `json:"statId"`
						Basic  struct {
							Value        float64 `json:"value"`
							DisplayValue string  `json:"displayValue"`
						} `json:"basic"`
						Pga struct {
							Value        float64 `json:"value"`
							DisplayValue string  `json:"displayValue"`
						} `json:"pga"`
					} `json:"weaponKillsSubmachinegun"`
					WeaponKillsRelic struct {
						StatID string `json:"statId"`
						Basic  struct {
							Value        float64 `json:"value"`
							DisplayValue string  `json:"displayValue"`
						} `json:"basic"`
						Pga struct {
							Value        float64 `json:"value"`
							DisplayValue string  `json:"displayValue"`
						} `json:"pga"`
					} `json:"weaponKillsRelic"`
					WeaponKillsSideArm struct {
						StatID string `json:"statId"`
						Basic  struct {
							Value        float64 `json:"value"`
							DisplayValue string  `json:"displayValue"`
						} `json:"basic"`
						Pga struct {
							Value        float64 `json:"value"`
							DisplayValue string  `json:"displayValue"`
						} `json:"pga"`
					} `json:"weaponKillsSideArm"`
					WeaponKillsSword struct {
						StatID string `json:"statId"`
						Basic  struct {
							Value        float64 `json:"value"`
							DisplayValue string  `json:"displayValue"`
						} `json:"basic"`
						Pga struct {
							Value        float64 `json:"value"`
							DisplayValue string  `json:"displayValue"`
						} `json:"pga"`
					} `json:"weaponKillsSword"`
					WeaponKillsAbility struct {
						StatID string `json:"statId"`
						Basic  struct {
							Value        float64 `json:"value"`
							DisplayValue string  `json:"displayValue"`
						} `json:"basic"`
						Pga struct {
							Value        float64 `json:"value"`
							DisplayValue string  `json:"displayValue"`
						} `json:"pga"`
					} `json:"weaponKillsAbility"`
					WeaponBestType struct {
						StatID string `json:"statId"`
						Basic  struct {
							Value        float64 `json:"value"`
							DisplayValue string  `json:"displayValue"`
						} `json:"basic"`
					} `json:"weaponBestType"`
					WinLossRatio struct {
						StatID string `json:"statId"`
						Basic  struct {
							Value        float64 `json:"value"`
							DisplayValue string  `json:"displayValue"`
						} `json:"basic"`
					} `json:"winLossRatio"`
					AllParticipantsCount struct {
						StatID string `json:"statId"`
						Basic  struct {
							Value        float64 `json:"value"`
							DisplayValue string  `json:"displayValue"`
						} `json:"basic"`
					} `json:"allParticipantsCount"`
					AllParticipantsScore struct {
						StatID string `json:"statId"`
						Basic  struct {
							Value        float64 `json:"value"`
							DisplayValue string  `json:"displayValue"`
						} `json:"basic"`
					} `json:"allParticipantsScore"`
					AllParticipantsTimePlayed struct {
						StatID string `json:"statId"`
						Basic  struct {
							Value        float64 `json:"value"`
							DisplayValue string  `json:"displayValue"`
						} `json:"basic"`
					} `json:"allParticipantsTimePlayed"`
					LongestKillSpree struct {
						StatID string `json:"statId"`
						Basic  struct {
							Value        float64 `json:"value"`
							DisplayValue string  `json:"displayValue"`
						} `json:"basic"`
					} `json:"longestKillSpree"`
					LongestSingleLife struct {
						StatID string `json:"statId"`
						Basic  struct {
							Value        float64 `json:"value"`
							DisplayValue string  `json:"displayValue"`
						} `json:"basic"`
					} `json:"longestSingleLife"`
					MostPrecisionKills struct {
						StatID string `json:"statId"`
						Basic  struct {
							Value        float64 `json:"value"`
							DisplayValue string  `json:"displayValue"`
						} `json:"basic"`
					} `json:"mostPrecisionKills"`
					OrbsDropped struct {
						StatID string `json:"statId"`
						Basic  struct {
							Value        float64 `json:"value"`
							DisplayValue string  `json:"displayValue"`
						} `json:"basic"`
						Pga struct {
							Value        float64 `json:"value"`
							DisplayValue string  `json:"displayValue"`
						} `json:"pga"`
					} `json:"orbsDropped"`
					OrbsGathered struct {
						StatID string `json:"statId"`
						Basic  struct {
							Value        float64 `json:"value"`
							DisplayValue string  `json:"displayValue"`
						} `json:"basic"`
						Pga struct {
							Value        float64 `json:"value"`
							DisplayValue string  `json:"displayValue"`
						} `json:"pga"`
					} `json:"orbsGathered"`
					RemainingTimeAfterQuitSeconds struct {
						StatID string `json:"statId"`
						Basic  struct {
							Value        float64 `json:"value"`
							DisplayValue string  `json:"displayValue"`
						} `json:"basic"`
						Pga struct {
							Value        float64 `json:"value"`
							DisplayValue string  `json:"displayValue"`
						} `json:"pga"`
					} `json:"remainingTimeAfterQuitSeconds"`
					TeamScore struct {
						StatID string `json:"statId"`
						Basic  struct {
							Value        float64 `json:"value"`
							DisplayValue string  `json:"displayValue"`
						} `json:"basic"`
						Pga struct {
							Value        float64 `json:"value"`
							DisplayValue string  `json:"displayValue"`
						} `json:"pga"`
					} `json:"teamScore"`
					TotalActivityDurationSeconds struct {
						StatID string `json:"statId"`
						Basic  struct {
							Value        float64 `json:"value"`
							DisplayValue string  `json:"displayValue"`
						} `json:"basic"`
						Pga struct {
							Value        float64 `json:"value"`
							DisplayValue string  `json:"displayValue"`
						} `json:"pga"`
					} `json:"totalActivityDurationSeconds"`
					CombatRating struct {
						StatID string `json:"statId"`
						Basic  struct {
							Value        float64 `json:"value"`
							DisplayValue string  `json:"displayValue"`
						} `json:"basic"`
					} `json:"combatRating"`
					FastestCompletionMs struct {
						StatID string `json:"statId"`
						Basic  struct {
							Value        float64 `json:"value"`
							DisplayValue string  `json:"displayValue"`
						} `json:"basic"`
					} `json:"fastestCompletionMs"`
					LongestKillDistance struct {
						StatID string `json:"statId"`
						Basic  struct {
							Value        float64 `json:"value"`
							DisplayValue string  `json:"displayValue"`
						} `json:"basic"`
					} `json:"longestKillDistance"`
					HighestCharacterLevel struct {
						StatID string `json:"statId"`
						Basic  struct {
							Value        float64 `json:"value"`
							DisplayValue string  `json:"displayValue"`
						} `json:"basic"`
					} `json:"highestCharacterLevel"`
					HighestLightLevel struct {
						StatID string `json:"statId"`
						Basic  struct {
							Value        float64 `json:"value"`
							DisplayValue string  `json:"displayValue"`
						} `json:"basic"`
					} `json:"highestLightLevel"`
				} `json:"allTime"`
			} `json:"allPvP"`
		} `json:"results"`
		Merged struct {
			AllTime struct {
				ActivitiesCleared struct {
					StatID string `json:"statId"`
					Basic  struct {
						Value        float64 `json:"value"`
						DisplayValue string  `json:"displayValue"`
					} `json:"basic"`
				} `json:"activitiesCleared"`
				WeaponKillsSuper struct {
					StatID string `json:"statId"`
					Basic  struct {
						Value        float64 `json:"value"`
						DisplayValue string  `json:"displayValue"`
					} `json:"basic"`
					Pga struct {
						Value        float64 `json:"value"`
						DisplayValue string  `json:"displayValue"`
					} `json:"pga"`
				} `json:"weaponKillsSuper"`
				ActivitiesEntered struct {
					StatID string `json:"statId"`
					Basic  struct {
						Value        float64 `json:"value"`
						DisplayValue string  `json:"displayValue"`
					} `json:"basic"`
				} `json:"activitiesEntered"`
				WeaponKillsMelee struct {
					StatID string `json:"statId"`
					Basic  struct {
						Value        float64 `json:"value"`
						DisplayValue string  `json:"displayValue"`
					} `json:"basic"`
					Pga struct {
						Value        float64 `json:"value"`
						DisplayValue string  `json:"displayValue"`
					} `json:"pga"`
				} `json:"weaponKillsMelee"`
				WeaponKillsGrenade struct {
					StatID string `json:"statId"`
					Basic  struct {
						Value        float64 `json:"value"`
						DisplayValue string  `json:"displayValue"`
					} `json:"basic"`
					Pga struct {
						Value        float64 `json:"value"`
						DisplayValue string  `json:"displayValue"`
					} `json:"pga"`
				} `json:"weaponKillsGrenade"`
				AbilityKills struct {
					StatID string `json:"statId"`
					Basic  struct {
						Value        float64 `json:"value"`
						DisplayValue string  `json:"displayValue"`
					} `json:"basic"`
					Pga struct {
						Value        float64 `json:"value"`
						DisplayValue string  `json:"displayValue"`
					} `json:"pga"`
				} `json:"abilityKills"`
				Assists struct {
					StatID string `json:"statId"`
					Basic  struct {
						Value        float64 `json:"value"`
						DisplayValue string  `json:"displayValue"`
					} `json:"basic"`
					Pga struct {
						Value        float64 `json:"value"`
						DisplayValue string  `json:"displayValue"`
					} `json:"pga"`
				} `json:"assists"`
				TotalDeathDistance struct {
					StatID string `json:"statId"`
					Basic  struct {
						Value        float64 `json:"value"`
						DisplayValue string  `json:"displayValue"`
					} `json:"basic"`
				} `json:"totalDeathDistance"`
				AverageDeathDistance struct {
					StatID string `json:"statId"`
					Basic  struct {
						Value        float64 `json:"value"`
						DisplayValue string  `json:"displayValue"`
					} `json:"basic"`
				} `json:"averageDeathDistance"`
				TotalKillDistance struct {
					StatID string `json:"statId"`
					Basic  struct {
						Value        float64 `json:"value"`
						DisplayValue string  `json:"displayValue"`
					} `json:"basic"`
				} `json:"totalKillDistance"`
				Kills struct {
					StatID string `json:"statId"`
					Basic  struct {
						Value        float64 `json:"value"`
						DisplayValue string  `json:"displayValue"`
					} `json:"basic"`
					Pga struct {
						Value        float64 `json:"value"`
						DisplayValue string  `json:"displayValue"`
					} `json:"pga"`
				} `json:"kills"`
				AverageKillDistance struct {
					StatID string `json:"statId"`
					Basic  struct {
						Value        float64 `json:"value"`
						DisplayValue string  `json:"displayValue"`
					} `json:"basic"`
				} `json:"averageKillDistance"`
				SecondsPlayed struct {
					StatID string `json:"statId"`
					Basic  struct {
						Value        float64 `json:"value"`
						DisplayValue string  `json:"displayValue"`
					} `json:"basic"`
					Pga struct {
						Value        float64 `json:"value"`
						DisplayValue string  `json:"displayValue"`
					} `json:"pga"`
				} `json:"secondsPlayed"`
				Deaths struct {
					StatID string `json:"statId"`
					Basic  struct {
						Value        float64 `json:"value"`
						DisplayValue string  `json:"displayValue"`
					} `json:"basic"`
					Pga struct {
						Value        float64 `json:"value"`
						DisplayValue string  `json:"displayValue"`
					} `json:"pga"`
				} `json:"deaths"`
				AverageLifespan struct {
					StatID string `json:"statId"`
					Basic  struct {
						Value        float64 `json:"value"`
						DisplayValue string  `json:"displayValue"`
					} `json:"basic"`
				} `json:"averageLifespan"`
				BestSingleGameKills struct {
					StatID string `json:"statId"`
					Basic  struct {
						Value        float64 `json:"value"`
						DisplayValue string  `json:"displayValue"`
					} `json:"basic"`
				} `json:"bestSingleGameKills"`
				KillsDeathsRatio struct {
					StatID string `json:"statId"`
					Basic  struct {
						Value        float64 `json:"value"`
						DisplayValue string  `json:"displayValue"`
					} `json:"basic"`
				} `json:"killsDeathsRatio"`
				KillsDeathsAssists struct {
					StatID string `json:"statId"`
					Basic  struct {
						Value        float64 `json:"value"`
						DisplayValue string  `json:"displayValue"`
					} `json:"basic"`
				} `json:"killsDeathsAssists"`
				ObjectivesCompleted struct {
					StatID string `json:"statId"`
					Basic  struct {
						Value        float64 `json:"value"`
						DisplayValue string  `json:"displayValue"`
					} `json:"basic"`
					Pga struct {
						Value        float64 `json:"value"`
						DisplayValue string  `json:"displayValue"`
					} `json:"pga"`
				} `json:"objectivesCompleted"`
				PrecisionKills struct {
					StatID string `json:"statId"`
					Basic  struct {
						Value        float64 `json:"value"`
						DisplayValue string  `json:"displayValue"`
					} `json:"basic"`
					Pga struct {
						Value        float64 `json:"value"`
						DisplayValue string  `json:"displayValue"`
					} `json:"pga"`
				} `json:"precisionKills"`
				ResurrectionsPerformed struct {
					StatID string `json:"statId"`
					Basic  struct {
						Value        float64 `json:"value"`
						DisplayValue string  `json:"displayValue"`
					} `json:"basic"`
					Pga struct {
						Value        float64 `json:"value"`
						DisplayValue string  `json:"displayValue"`
					} `json:"pga"`
				} `json:"resurrectionsPerformed"`
				ResurrectionsReceived struct {
					StatID string `json:"statId"`
					Basic  struct {
						Value        float64 `json:"value"`
						DisplayValue string  `json:"displayValue"`
					} `json:"basic"`
					Pga struct {
						Value        float64 `json:"value"`
						DisplayValue string  `json:"displayValue"`
					} `json:"pga"`
				} `json:"resurrectionsReceived"`
				Suicides struct {
					StatID string `json:"statId"`
					Basic  struct {
						Value        float64 `json:"value"`
						DisplayValue string  `json:"displayValue"`
					} `json:"basic"`
					Pga struct {
						Value        float64 `json:"value"`
						DisplayValue string  `json:"displayValue"`
					} `json:"pga"`
				} `json:"suicides"`
				WeaponKillsAutoRifle struct {
					StatID string `json:"statId"`
					Basic  struct {
						Value        float64 `json:"value"`
						DisplayValue string  `json:"displayValue"`
					} `json:"basic"`
					Pga struct {
						Value        float64 `json:"value"`
						DisplayValue string  `json:"displayValue"`
					} `json:"pga"`
				} `json:"weaponKillsAutoRifle"`
				WeaponKillsFusionRifle struct {
					StatID string `json:"statId"`
					Basic  struct {
						Value        float64 `json:"value"`
						DisplayValue string  `json:"displayValue"`
					} `json:"basic"`
					Pga struct {
						Value        float64 `json:"value"`
						DisplayValue string  `json:"displayValue"`
					} `json:"pga"`
				} `json:"weaponKillsFusionRifle"`
				WeaponKillsHandCannon struct {
					StatID string `json:"statId"`
					Basic  struct {
						Value        float64 `json:"value"`
						DisplayValue string  `json:"displayValue"`
					} `json:"basic"`
					Pga struct {
						Value        float64 `json:"value"`
						DisplayValue string  `json:"displayValue"`
					} `json:"pga"`
				} `json:"weaponKillsHandCannon"`
				WeaponKillsMachinegun struct {
					StatID string `json:"statId"`
					Basic  struct {
						Value        float64 `json:"value"`
						DisplayValue string  `json:"displayValue"`
					} `json:"basic"`
					Pga struct {
						Value        float64 `json:"value"`
						DisplayValue string  `json:"displayValue"`
					} `json:"pga"`
				} `json:"weaponKillsMachinegun"`
				WeaponKillsPulseRifle struct {
					StatID string `json:"statId"`
					Basic  struct {
						Value        float64 `json:"value"`
						DisplayValue string  `json:"displayValue"`
					} `json:"basic"`
					Pga struct {
						Value        float64 `json:"value"`
						DisplayValue string  `json:"displayValue"`
					} `json:"pga"`
				} `json:"weaponKillsPulseRifle"`
				WeaponKillsRocketLauncher struct {
					StatID string `json:"statId"`
					Basic  struct {
						Value        float64 `json:"value"`
						DisplayValue string  `json:"displayValue"`
					} `json:"basic"`
					Pga struct {
						Value        float64 `json:"value"`
						DisplayValue string  `json:"displayValue"`
					} `json:"pga"`
				} `json:"weaponKillsRocketLauncher"`
				WeaponKillsScoutRifle struct {
					StatID string `json:"statId"`
					Basic  struct {
						Value        float64 `json:"value"`
						DisplayValue string  `json:"displayValue"`
					} `json:"basic"`
					Pga struct {
						Value        float64 `json:"value"`
						DisplayValue string  `json:"displayValue"`
					} `json:"pga"`
				} `json:"weaponKillsScoutRifle"`
				WeaponKillsShotgun struct {
					StatID string `json:"statId"`
					Basic  struct {
						Value        float64 `json:"value"`
						DisplayValue string  `json:"displayValue"`
					} `json:"basic"`
					Pga struct {
						Value        float64 `json:"value"`
						DisplayValue string  `json:"displayValue"`
					} `json:"pga"`
				} `json:"weaponKillsShotgun"`
				WeaponKillsSniper struct {
					StatID string `json:"statId"`
					Basic  struct {
						Value        float64 `json:"value"`
						DisplayValue string  `json:"displayValue"`
					} `json:"basic"`
					Pga struct {
						Value        float64 `json:"value"`
						DisplayValue string  `json:"displayValue"`
					} `json:"pga"`
				} `json:"weaponKillsSniper"`
				WeaponKillsSubmachinegun struct {
					StatID string `json:"statId"`
					Basic  struct {
						Value        float64 `json:"value"`
						DisplayValue string  `json:"displayValue"`
					} `json:"basic"`
					Pga struct {
						Value        float64 `json:"value"`
						DisplayValue string  `json:"displayValue"`
					} `json:"pga"`
				} `json:"weaponKillsSubmachinegun"`
				WeaponKillsRelic struct {
					StatID string `json:"statId"`
					Basic  struct {
						Value        float64 `json:"value"`
						DisplayValue string  `json:"displayValue"`
					} `json:"basic"`
					Pga struct {
						Value        float64 `json:"value"`
						DisplayValue string  `json:"displayValue"`
					} `json:"pga"`
				} `json:"weaponKillsRelic"`
				WeaponKillsSideArm struct {
					StatID string `json:"statId"`
					Basic  struct {
						Value        float64 `json:"value"`
						DisplayValue string  `json:"displayValue"`
					} `json:"basic"`
					Pga struct {
						Value        float64 `json:"value"`
						DisplayValue string  `json:"displayValue"`
					} `json:"pga"`
				} `json:"weaponKillsSideArm"`
				WeaponKillsSword struct {
					StatID string `json:"statId"`
					Basic  struct {
						Value        float64 `json:"value"`
						DisplayValue string  `json:"displayValue"`
					} `json:"basic"`
					Pga struct {
						Value        float64 `json:"value"`
						DisplayValue string  `json:"displayValue"`
					} `json:"pga"`
				} `json:"weaponKillsSword"`
				WeaponKillsAbility struct {
					StatID string `json:"statId"`
					Basic  struct {
						Value        float64 `json:"value"`
						DisplayValue string  `json:"displayValue"`
					} `json:"basic"`
					Pga struct {
						Value        float64 `json:"value"`
						DisplayValue string  `json:"displayValue"`
					} `json:"pga"`
				} `json:"weaponKillsAbility"`
				WeaponBestType struct {
					StatID string `json:"statId"`
					Basic  struct {
						Value        float64 `json:"value"`
						DisplayValue string  `json:"displayValue"`
					} `json:"basic"`
				} `json:"weaponBestType"`
				AllParticipantsCount struct {
					StatID string `json:"statId"`
					Basic  struct {
						Value        float64 `json:"value"`
						DisplayValue string  `json:"displayValue"`
					} `json:"basic"`
				} `json:"allParticipantsCount"`
				AllParticipantsTimePlayed struct {
					StatID string `json:"statId"`
					Basic  struct {
						Value        float64 `json:"value"`
						DisplayValue string  `json:"displayValue"`
					} `json:"basic"`
				} `json:"allParticipantsTimePlayed"`
				LongestKillSpree struct {
					StatID string `json:"statId"`
					Basic  struct {
						Value        float64 `json:"value"`
						DisplayValue string  `json:"displayValue"`
					} `json:"basic"`
				} `json:"longestKillSpree"`
				LongestSingleLife struct {
					StatID string `json:"statId"`
					Basic  struct {
						Value        float64 `json:"value"`
						DisplayValue string  `json:"displayValue"`
					} `json:"basic"`
				} `json:"longestSingleLife"`
				MostPrecisionKills struct {
					StatID string `json:"statId"`
					Basic  struct {
						Value        float64 `json:"value"`
						DisplayValue string  `json:"displayValue"`
					} `json:"basic"`
				} `json:"mostPrecisionKills"`
				OrbsDropped struct {
					StatID string `json:"statId"`
					Basic  struct {
						Value        float64 `json:"value"`
						DisplayValue string  `json:"displayValue"`
					} `json:"basic"`
					Pga struct {
						Value        float64 `json:"value"`
						DisplayValue string  `json:"displayValue"`
					} `json:"pga"`
				} `json:"orbsDropped"`
				OrbsGathered struct {
					StatID string `json:"statId"`
					Basic  struct {
						Value        float64 `json:"value"`
						DisplayValue string  `json:"displayValue"`
					} `json:"basic"`
					Pga struct {
						Value        float64 `json:"value"`
						DisplayValue string  `json:"displayValue"`
					} `json:"pga"`
				} `json:"orbsGathered"`
				PublicEventsCompleted struct {
					StatID string `json:"statId"`
					Basic  struct {
						Value        float64 `json:"value"`
						DisplayValue string  `json:"displayValue"`
					} `json:"basic"`
					Pga struct {
						Value        float64 `json:"value"`
						DisplayValue string  `json:"displayValue"`
					} `json:"pga"`
				} `json:"publicEventsCompleted"`
				RemainingTimeAfterQuitSeconds struct {
					StatID string `json:"statId"`
					Basic  struct {
						Value        float64 `json:"value"`
						DisplayValue string  `json:"displayValue"`
					} `json:"basic"`
					Pga struct {
						Value        float64 `json:"value"`
						DisplayValue string  `json:"displayValue"`
					} `json:"pga"`
				} `json:"remainingTimeAfterQuitSeconds"`
				TotalActivityDurationSeconds struct {
					StatID string `json:"statId"`
					Basic  struct {
						Value        float64 `json:"value"`
						DisplayValue string  `json:"displayValue"`
					} `json:"basic"`
					Pga struct {
						Value        float64 `json:"value"`
						DisplayValue string  `json:"displayValue"`
					} `json:"pga"`
				} `json:"totalActivityDurationSeconds"`
				FastestCompletionMs struct {
					StatID string `json:"statId"`
					Basic  struct {
						Value        float64 `json:"value"`
						DisplayValue string  `json:"displayValue"`
					} `json:"basic"`
				} `json:"fastestCompletionMs"`
				LongestKillDistance struct {
					StatID string `json:"statId"`
					Basic  struct {
						Value        float64 `json:"value"`
						DisplayValue string  `json:"displayValue"`
					} `json:"basic"`
				} `json:"longestKillDistance"`
				HighestCharacterLevel struct {
					StatID string `json:"statId"`
					Basic  struct {
						Value        float64 `json:"value"`
						DisplayValue string  `json:"displayValue"`
					} `json:"basic"`
				} `json:"highestCharacterLevel"`
				HighestLightLevel struct {
					StatID string `json:"statId"`
					Basic  struct {
						Value        float64 `json:"value"`
						DisplayValue string  `json:"displayValue"`
					} `json:"basic"`
				} `json:"highestLightLevel"`
				ActivitiesWon struct {
					StatID string `json:"statId"`
					Basic  struct {
						Value        float64 `json:"value"`
						DisplayValue string  `json:"displayValue"`
					} `json:"basic"`
				} `json:"activitiesWon"`
				Score struct {
					StatID string `json:"statId"`
					Basic  struct {
						Value        float64 `json:"value"`
						DisplayValue string  `json:"displayValue"`
					} `json:"basic"`
					Pga struct {
						Value        float64 `json:"value"`
						DisplayValue string  `json:"displayValue"`
					} `json:"pga"`
				} `json:"score"`
				AverageScorePerKill struct {
					StatID string `json:"statId"`
					Basic  struct {
						Value        float64 `json:"value"`
						DisplayValue string  `json:"displayValue"`
					} `json:"basic"`
				} `json:"averageScorePerKill"`
				AverageScorePerLife struct {
					StatID string `json:"statId"`
					Basic  struct {
						Value        float64 `json:"value"`
						DisplayValue string  `json:"displayValue"`
					} `json:"basic"`
				} `json:"averageScorePerLife"`
				BestSingleGameScore struct {
					StatID string `json:"statId"`
					Basic  struct {
						Value        float64 `json:"value"`
						DisplayValue string  `json:"displayValue"`
					} `json:"basic"`
				} `json:"bestSingleGameScore"`
				WinLossRatio struct {
					StatID string `json:"statId"`
					Basic  struct {
						Value        float64 `json:"value"`
						DisplayValue string  `json:"displayValue"`
					} `json:"basic"`
				} `json:"winLossRatio"`
				AllParticipantsScore struct {
					StatID string `json:"statId"`
					Basic  struct {
						Value        float64 `json:"value"`
						DisplayValue string  `json:"displayValue"`
					} `json:"basic"`
				} `json:"allParticipantsScore"`
				TeamScore struct {
					StatID string `json:"statId"`
					Basic  struct {
						Value        float64 `json:"value"`
						DisplayValue string  `json:"displayValue"`
					} `json:"basic"`
					Pga struct {
						Value        float64 `json:"value"`
						DisplayValue string  `json:"displayValue"`
					} `json:"pga"`
				} `json:"teamScore"`
				CombatRating struct {
					StatID string `json:"statId"`
					Basic  struct {
						Value        float64 `json:"value"`
						DisplayValue string  `json:"displayValue"`
					} `json:"basic"`
				} `json:"combatRating"`
			} `json:"allTime"`
		} `json:"merged"`
	} `json:"mergedAllCharacters"`
	Characters []struct {
		CharacterID string `json:"characterId"`
		Deleted     bool   `json:"deleted"`
		Results     struct {
			AllPvP struct {
				AllTime struct {
					WeaponKillsSuper struct {
						StatID string `json:"statId"`
						Basic  struct {
							Value        float64 `json:"value"`
							DisplayValue string  `json:"displayValue"`
						} `json:"basic"`
						Pga struct {
							Value        float64 `json:"value"`
							DisplayValue string  `json:"displayValue"`
						} `json:"pga"`
					} `json:"weaponKillsSuper"`
					ActivitiesEntered struct {
						StatID string `json:"statId"`
						Basic  struct {
							Value        float64 `json:"value"`
							DisplayValue string  `json:"displayValue"`
						} `json:"basic"`
					} `json:"activitiesEntered"`
					WeaponKillsMelee struct {
						StatID string `json:"statId"`
						Basic  struct {
							Value        float64 `json:"value"`
							DisplayValue string  `json:"displayValue"`
						} `json:"basic"`
						Pga struct {
							Value        float64 `json:"value"`
							DisplayValue string  `json:"displayValue"`
						} `json:"pga"`
					} `json:"weaponKillsMelee"`
					WeaponKillsGrenade struct {
						StatID string `json:"statId"`
						Basic  struct {
							Value        float64 `json:"value"`
							DisplayValue string  `json:"displayValue"`
						} `json:"basic"`
						Pga struct {
							Value        float64 `json:"value"`
							DisplayValue string  `json:"displayValue"`
						} `json:"pga"`
					} `json:"weaponKillsGrenade"`
					AbilityKills struct {
						StatID string `json:"statId"`
						Basic  struct {
							Value        float64 `json:"value"`
							DisplayValue string  `json:"displayValue"`
						} `json:"basic"`
						Pga struct {
							Value        float64 `json:"value"`
							DisplayValue string  `json:"displayValue"`
						} `json:"pga"`
					} `json:"abilityKills"`
					ActivitiesWon struct {
						StatID string `json:"statId"`
						Basic  struct {
							Value        float64 `json:"value"`
							DisplayValue string  `json:"displayValue"`
						} `json:"basic"`
					} `json:"activitiesWon"`
					Assists struct {
						StatID string `json:"statId"`
						Basic  struct {
							Value        float64 `json:"value"`
							DisplayValue string  `json:"displayValue"`
						} `json:"basic"`
						Pga struct {
							Value        float64 `json:"value"`
							DisplayValue string  `json:"displayValue"`
						} `json:"pga"`
					} `json:"assists"`
					TotalDeathDistance struct {
						StatID string `json:"statId"`
						Basic  struct {
							Value        float64 `json:"value"`
							DisplayValue string  `json:"displayValue"`
						} `json:"basic"`
					} `json:"totalDeathDistance"`
					AverageDeathDistance struct {
						StatID string `json:"statId"`
						Basic  struct {
							Value        float64 `json:"value"`
							DisplayValue string  `json:"displayValue"`
						} `json:"basic"`
					} `json:"averageDeathDistance"`
					TotalKillDistance struct {
						StatID string `json:"statId"`
						Basic  struct {
							Value        float64 `json:"value"`
							DisplayValue string  `json:"displayValue"`
						} `json:"basic"`
					} `json:"totalKillDistance"`
					Kills struct {
						StatID string `json:"statId"`
						Basic  struct {
							Value        float64 `json:"value"`
							DisplayValue string  `json:"displayValue"`
						} `json:"basic"`
						Pga struct {
							Value        float64 `json:"value"`
							DisplayValue string  `json:"displayValue"`
						} `json:"pga"`
					} `json:"kills"`
					AverageKillDistance struct {
						StatID string `json:"statId"`
						Basic  struct {
							Value        float64 `json:"value"`
							DisplayValue string  `json:"displayValue"`
						} `json:"basic"`
					} `json:"averageKillDistance"`
					SecondsPlayed struct {
						StatID string `json:"statId"`
						Basic  struct {
							Value        float64 `json:"value"`
							DisplayValue string  `json:"displayValue"`
						} `json:"basic"`
						Pga struct {
							Value        float64 `json:"value"`
							DisplayValue string  `json:"displayValue"`
						} `json:"pga"`
					} `json:"secondsPlayed"`
					Deaths struct {
						StatID string `json:"statId"`
						Basic  struct {
							Value        float64 `json:"value"`
							DisplayValue string  `json:"displayValue"`
						} `json:"basic"`
						Pga struct {
							Value        float64 `json:"value"`
							DisplayValue string  `json:"displayValue"`
						} `json:"pga"`
					} `json:"deaths"`
					AverageLifespan struct {
						StatID string `json:"statId"`
						Basic  struct {
							Value        float64 `json:"value"`
							DisplayValue string  `json:"displayValue"`
						} `json:"basic"`
					} `json:"averageLifespan"`
					Score struct {
						StatID string `json:"statId"`
						Basic  struct {
							Value        float64 `json:"value"`
							DisplayValue string  `json:"displayValue"`
						} `json:"basic"`
						Pga struct {
							Value        float64 `json:"value"`
							DisplayValue string  `json:"displayValue"`
						} `json:"pga"`
					} `json:"score"`
					AverageScorePerKill struct {
						StatID string `json:"statId"`
						Basic  struct {
							Value        float64 `json:"value"`
							DisplayValue string  `json:"displayValue"`
						} `json:"basic"`
					} `json:"averageScorePerKill"`
					AverageScorePerLife struct {
						StatID string `json:"statId"`
						Basic  struct {
							Value        float64 `json:"value"`
							DisplayValue string  `json:"displayValue"`
						} `json:"basic"`
					} `json:"averageScorePerLife"`
					BestSingleGameKills struct {
						StatID string `json:"statId"`
						Basic  struct {
							Value        float64 `json:"value"`
							DisplayValue string  `json:"displayValue"`
						} `json:"basic"`
					} `json:"bestSingleGameKills"`
					BestSingleGameScore struct {
						StatID string `json:"statId"`
						Basic  struct {
							Value        float64 `json:"value"`
							DisplayValue string  `json:"displayValue"`
						} `json:"basic"`
					} `json:"bestSingleGameScore"`
					KillsDeathsRatio struct {
						StatID string `json:"statId"`
						Basic  struct {
							Value        float64 `json:"value"`
							DisplayValue string  `json:"displayValue"`
						} `json:"basic"`
					} `json:"killsDeathsRatio"`
					KillsDeathsAssists struct {
						StatID string `json:"statId"`
						Basic  struct {
							Value        float64 `json:"value"`
							DisplayValue string  `json:"displayValue"`
						} `json:"basic"`
					} `json:"killsDeathsAssists"`
					ObjectivesCompleted struct {
						StatID string `json:"statId"`
						Basic  struct {
							Value        float64 `json:"value"`
							DisplayValue string  `json:"displayValue"`
						} `json:"basic"`
						Pga struct {
							Value        float64 `json:"value"`
							DisplayValue string  `json:"displayValue"`
						} `json:"pga"`
					} `json:"objectivesCompleted"`
					PrecisionKills struct {
						StatID string `json:"statId"`
						Basic  struct {
							Value        float64 `json:"value"`
							DisplayValue string  `json:"displayValue"`
						} `json:"basic"`
						Pga struct {
							Value        float64 `json:"value"`
							DisplayValue string  `json:"displayValue"`
						} `json:"pga"`
					} `json:"precisionKills"`
					ResurrectionsPerformed struct {
						StatID string `json:"statId"`
						Basic  struct {
							Value        float64 `json:"value"`
							DisplayValue string  `json:"displayValue"`
						} `json:"basic"`
						Pga struct {
							Value        float64 `json:"value"`
							DisplayValue string  `json:"displayValue"`
						} `json:"pga"`
					} `json:"resurrectionsPerformed"`
					ResurrectionsReceived struct {
						StatID string `json:"statId"`
						Basic  struct {
							Value        float64 `json:"value"`
							DisplayValue string  `json:"displayValue"`
						} `json:"basic"`
						Pga struct {
							Value        float64 `json:"value"`
							DisplayValue string  `json:"displayValue"`
						} `json:"pga"`
					} `json:"resurrectionsReceived"`
					Suicides struct {
						StatID string `json:"statId"`
						Basic  struct {
							Value        float64 `json:"value"`
							DisplayValue string  `json:"displayValue"`
						} `json:"basic"`
						Pga struct {
							Value        float64 `json:"value"`
							DisplayValue string  `json:"displayValue"`
						} `json:"pga"`
					} `json:"suicides"`
					WeaponKillsAutoRifle struct {
						StatID string `json:"statId"`
						Basic  struct {
							Value        float64 `json:"value"`
							DisplayValue string  `json:"displayValue"`
						} `json:"basic"`
						Pga struct {
							Value        float64 `json:"value"`
							DisplayValue string  `json:"displayValue"`
						} `json:"pga"`
					} `json:"weaponKillsAutoRifle"`
					WeaponKillsFusionRifle struct {
						StatID string `json:"statId"`
						Basic  struct {
							Value        float64 `json:"value"`
							DisplayValue string  `json:"displayValue"`
						} `json:"basic"`
						Pga struct {
							Value        float64 `json:"value"`
							DisplayValue string  `json:"displayValue"`
						} `json:"pga"`
					} `json:"weaponKillsFusionRifle"`
					WeaponKillsHandCannon struct {
						StatID string `json:"statId"`
						Basic  struct {
							Value        float64 `json:"value"`
							DisplayValue string  `json:"displayValue"`
						} `json:"basic"`
						Pga struct {
							Value        float64 `json:"value"`
							DisplayValue string  `json:"displayValue"`
						} `json:"pga"`
					} `json:"weaponKillsHandCannon"`
					WeaponKillsMachinegun struct {
						StatID string `json:"statId"`
						Basic  struct {
							Value        float64 `json:"value"`
							DisplayValue string  `json:"displayValue"`
						} `json:"basic"`
						Pga struct {
							Value        float64 `json:"value"`
							DisplayValue string  `json:"displayValue"`
						} `json:"pga"`
					} `json:"weaponKillsMachinegun"`
					WeaponKillsPulseRifle struct {
						StatID string `json:"statId"`
						Basic  struct {
							Value        float64 `json:"value"`
							DisplayValue string  `json:"displayValue"`
						} `json:"basic"`
						Pga struct {
							Value        float64 `json:"value"`
							DisplayValue string  `json:"displayValue"`
						} `json:"pga"`
					} `json:"weaponKillsPulseRifle"`
					WeaponKillsRocketLauncher struct {
						StatID string `json:"statId"`
						Basic  struct {
							Value        float64 `json:"value"`
							DisplayValue string  `json:"displayValue"`
						} `json:"basic"`
						Pga struct {
							Value        float64 `json:"value"`
							DisplayValue string  `json:"displayValue"`
						} `json:"pga"`
					} `json:"weaponKillsRocketLauncher"`
					WeaponKillsScoutRifle struct {
						StatID string `json:"statId"`
						Basic  struct {
							Value        float64 `json:"value"`
							DisplayValue string  `json:"displayValue"`
						} `json:"basic"`
						Pga struct {
							Value        float64 `json:"value"`
							DisplayValue string  `json:"displayValue"`
						} `json:"pga"`
					} `json:"weaponKillsScoutRifle"`
					WeaponKillsShotgun struct {
						StatID string `json:"statId"`
						Basic  struct {
							Value        float64 `json:"value"`
							DisplayValue string  `json:"displayValue"`
						} `json:"basic"`
						Pga struct {
							Value        float64 `json:"value"`
							DisplayValue string  `json:"displayValue"`
						} `json:"pga"`
					} `json:"weaponKillsShotgun"`
					WeaponKillsSniper struct {
						StatID string `json:"statId"`
						Basic  struct {
							Value        float64 `json:"value"`
							DisplayValue string  `json:"displayValue"`
						} `json:"basic"`
						Pga struct {
							Value        float64 `json:"value"`
							DisplayValue string  `json:"displayValue"`
						} `json:"pga"`
					} `json:"weaponKillsSniper"`
					WeaponKillsSubmachinegun struct {
						StatID string `json:"statId"`
						Basic  struct {
							Value        float64 `json:"value"`
							DisplayValue string  `json:"displayValue"`
						} `json:"basic"`
						Pga struct {
							Value        float64 `json:"value"`
							DisplayValue string  `json:"displayValue"`
						} `json:"pga"`
					} `json:"weaponKillsSubmachinegun"`
					WeaponKillsRelic struct {
						StatID string `json:"statId"`
						Basic  struct {
							Value        float64 `json:"value"`
							DisplayValue string  `json:"displayValue"`
						} `json:"basic"`
						Pga struct {
							Value        float64 `json:"value"`
							DisplayValue string  `json:"displayValue"`
						} `json:"pga"`
					} `json:"weaponKillsRelic"`
					WeaponKillsSideArm struct {
						StatID string `json:"statId"`
						Basic  struct {
							Value        float64 `json:"value"`
							DisplayValue string  `json:"displayValue"`
						} `json:"basic"`
						Pga struct {
							Value        float64 `json:"value"`
							DisplayValue string  `json:"displayValue"`
						} `json:"pga"`
					} `json:"weaponKillsSideArm"`
					WeaponKillsSword struct {
						StatID string `json:"statId"`
						Basic  struct {
							Value        float64 `json:"value"`
							DisplayValue string  `json:"displayValue"`
						} `json:"basic"`
						Pga struct {
							Value        float64 `json:"value"`
							DisplayValue string  `json:"displayValue"`
						} `json:"pga"`
					} `json:"weaponKillsSword"`
					WeaponKillsAbility struct {
						StatID string `json:"statId"`
						Basic  struct {
							Value        float64 `json:"value"`
							DisplayValue string  `json:"displayValue"`
						} `json:"basic"`
						Pga struct {
							Value        float64 `json:"value"`
							DisplayValue string  `json:"displayValue"`
						} `json:"pga"`
					} `json:"weaponKillsAbility"`
					WeaponBestType struct {
						StatID string `json:"statId"`
						Basic  struct {
							Value        float64 `json:"value"`
							DisplayValue string  `json:"displayValue"`
						} `json:"basic"`
					} `json:"weaponBestType"`
					WinLossRatio struct {
						StatID string `json:"statId"`
						Basic  struct {
							Value        float64 `json:"value"`
							DisplayValue string  `json:"displayValue"`
						} `json:"basic"`
					} `json:"winLossRatio"`
					AllParticipantsCount struct {
						StatID string `json:"statId"`
						Basic  struct {
							Value        float64 `json:"value"`
							DisplayValue string  `json:"displayValue"`
						} `json:"basic"`
					} `json:"allParticipantsCount"`
					AllParticipantsScore struct {
						StatID string `json:"statId"`
						Basic  struct {
							Value        float64 `json:"value"`
							DisplayValue string  `json:"displayValue"`
						} `json:"basic"`
					} `json:"allParticipantsScore"`
					AllParticipantsTimePlayed struct {
						StatID string `json:"statId"`
						Basic  struct {
							Value        float64 `json:"value"`
							DisplayValue string  `json:"displayValue"`
						} `json:"basic"`
					} `json:"allParticipantsTimePlayed"`
					LongestKillSpree struct {
						StatID string `json:"statId"`
						Basic  struct {
							Value        float64 `json:"value"`
							DisplayValue string  `json:"displayValue"`
						} `json:"basic"`
					} `json:"longestKillSpree"`
					LongestSingleLife struct {
						StatID string `json:"statId"`
						Basic  struct {
							Value        float64 `json:"value"`
							DisplayValue string  `json:"displayValue"`
						} `json:"basic"`
					} `json:"longestSingleLife"`
					MostPrecisionKills struct {
						StatID string `json:"statId"`
						Basic  struct {
							Value        float64 `json:"value"`
							DisplayValue string  `json:"displayValue"`
						} `json:"basic"`
					} `json:"mostPrecisionKills"`
					OrbsDropped struct {
						StatID string `json:"statId"`
						Basic  struct {
							Value        float64 `json:"value"`
							DisplayValue string  `json:"displayValue"`
						} `json:"basic"`
						Pga struct {
							Value        float64 `json:"value"`
							DisplayValue string  `json:"displayValue"`
						} `json:"pga"`
					} `json:"orbsDropped"`
					OrbsGathered struct {
						StatID string `json:"statId"`
						Basic  struct {
							Value        float64 `json:"value"`
							DisplayValue string  `json:"displayValue"`
						} `json:"basic"`
						Pga struct {
							Value        float64 `json:"value"`
							DisplayValue string  `json:"displayValue"`
						} `json:"pga"`
					} `json:"orbsGathered"`
					RemainingTimeAfterQuitSeconds struct {
						StatID string `json:"statId"`
						Basic  struct {
							Value        float64 `json:"value"`
							DisplayValue string  `json:"displayValue"`
						} `json:"basic"`
						Pga struct {
							Value        float64 `json:"value"`
							DisplayValue string  `json:"displayValue"`
						} `json:"pga"`
					} `json:"remainingTimeAfterQuitSeconds"`
					TeamScore struct {
						StatID string `json:"statId"`
						Basic  struct {
							Value        float64 `json:"value"`
							DisplayValue string  `json:"displayValue"`
						} `json:"basic"`
						Pga struct {
							Value        float64 `json:"value"`
							DisplayValue string  `json:"displayValue"`
						} `json:"pga"`
					} `json:"teamScore"`
					TotalActivityDurationSeconds struct {
						StatID string `json:"statId"`
						Basic  struct {
							Value        float64 `json:"value"`
							DisplayValue string  `json:"displayValue"`
						} `json:"basic"`
						Pga struct {
							Value        float64 `json:"value"`
							DisplayValue string  `json:"displayValue"`
						} `json:"pga"`
					} `json:"totalActivityDurationSeconds"`
					CombatRating struct {
						StatID string `json:"statId"`
						Basic  struct {
							Value        float64 `json:"value"`
							DisplayValue string  `json:"displayValue"`
						} `json:"basic"`
					} `json:"combatRating"`
					FastestCompletionMs struct {
						StatID string `json:"statId"`
						Basic  struct {
							Value        float64 `json:"value"`
							DisplayValue string  `json:"displayValue"`
						} `json:"basic"`
					} `json:"fastestCompletionMs"`
					LongestKillDistance struct {
						StatID string `json:"statId"`
						Basic  struct {
							Value        float64 `json:"value"`
							DisplayValue string  `json:"displayValue"`
						} `json:"basic"`
					} `json:"longestKillDistance"`
					HighestCharacterLevel struct {
						StatID string `json:"statId"`
						Basic  struct {
							Value        float64 `json:"value"`
							DisplayValue string  `json:"displayValue"`
						} `json:"basic"`
					} `json:"highestCharacterLevel"`
					HighestLightLevel struct {
						StatID string `json:"statId"`
						Basic  struct {
							Value        float64 `json:"value"`
							DisplayValue string  `json:"displayValue"`
						} `json:"basic"`
					} `json:"highestLightLevel"`
				} `json:"allTime"`
			} `json:"allPvP"`
			AllPvE struct {
				AllTime struct {
					ActivitiesCleared struct {
						StatID string `json:"statId"`
						Basic  struct {
							Value        float64 `json:"value"`
							DisplayValue string  `json:"displayValue"`
						} `json:"basic"`
					} `json:"activitiesCleared"`
					WeaponKillsSuper struct {
						StatID string `json:"statId"`
						Basic  struct {
							Value        float64 `json:"value"`
							DisplayValue string  `json:"displayValue"`
						} `json:"basic"`
						Pga struct {
							Value        float64 `json:"value"`
							DisplayValue string  `json:"displayValue"`
						} `json:"pga"`
					} `json:"weaponKillsSuper"`
					ActivitiesEntered struct {
						StatID string `json:"statId"`
						Basic  struct {
							Value        float64 `json:"value"`
							DisplayValue string  `json:"displayValue"`
						} `json:"basic"`
					} `json:"activitiesEntered"`
					WeaponKillsMelee struct {
						StatID string `json:"statId"`
						Basic  struct {
							Value        float64 `json:"value"`
							DisplayValue string  `json:"displayValue"`
						} `json:"basic"`
						Pga struct {
							Value        float64 `json:"value"`
							DisplayValue string  `json:"displayValue"`
						} `json:"pga"`
					} `json:"weaponKillsMelee"`
					WeaponKillsGrenade struct {
						StatID string `json:"statId"`
						Basic  struct {
							Value        float64 `json:"value"`
							DisplayValue string  `json:"displayValue"`
						} `json:"basic"`
						Pga struct {
							Value        float64 `json:"value"`
							DisplayValue string  `json:"displayValue"`
						} `json:"pga"`
					} `json:"weaponKillsGrenade"`
					AbilityKills struct {
						StatID string `json:"statId"`
						Basic  struct {
							Value        float64 `json:"value"`
							DisplayValue string  `json:"displayValue"`
						} `json:"basic"`
						Pga struct {
							Value        float64 `json:"value"`
							DisplayValue string  `json:"displayValue"`
						} `json:"pga"`
					} `json:"abilityKills"`
					Assists struct {
						StatID string `json:"statId"`
						Basic  struct {
							Value        float64 `json:"value"`
							DisplayValue string  `json:"displayValue"`
						} `json:"basic"`
						Pga struct {
							Value        float64 `json:"value"`
							DisplayValue string  `json:"displayValue"`
						} `json:"pga"`
					} `json:"assists"`
					TotalDeathDistance struct {
						StatID string `json:"statId"`
						Basic  struct {
							Value        float64 `json:"value"`
							DisplayValue string  `json:"displayValue"`
						} `json:"basic"`
					} `json:"totalDeathDistance"`
					AverageDeathDistance struct {
						StatID string `json:"statId"`
						Basic  struct {
							Value        float64 `json:"value"`
							DisplayValue string  `json:"displayValue"`
						} `json:"basic"`
					} `json:"averageDeathDistance"`
					TotalKillDistance struct {
						StatID string `json:"statId"`
						Basic  struct {
							Value        float64 `json:"value"`
							DisplayValue string  `json:"displayValue"`
						} `json:"basic"`
					} `json:"totalKillDistance"`
					Kills struct {
						StatID string `json:"statId"`
						Basic  struct {
							Value        float64 `json:"value"`
							DisplayValue string  `json:"displayValue"`
						} `json:"basic"`
						Pga struct {
							Value        float64 `json:"value"`
							DisplayValue string  `json:"displayValue"`
						} `json:"pga"`
					} `json:"kills"`
					AverageKillDistance struct {
						StatID string `json:"statId"`
						Basic  struct {
							Value        float64 `json:"value"`
							DisplayValue string  `json:"displayValue"`
						} `json:"basic"`
					} `json:"averageKillDistance"`
					SecondsPlayed struct {
						StatID string `json:"statId"`
						Basic  struct {
							Value        float64 `json:"value"`
							DisplayValue string  `json:"displayValue"`
						} `json:"basic"`
						Pga struct {
							Value        float64 `json:"value"`
							DisplayValue string  `json:"displayValue"`
						} `json:"pga"`
					} `json:"secondsPlayed"`
					Deaths struct {
						StatID string `json:"statId"`
						Basic  struct {
							Value        float64 `json:"value"`
							DisplayValue string  `json:"displayValue"`
						} `json:"basic"`
						Pga struct {
							Value        float64 `json:"value"`
							DisplayValue string  `json:"displayValue"`
						} `json:"pga"`
					} `json:"deaths"`
					AverageLifespan struct {
						StatID string `json:"statId"`
						Basic  struct {
							Value        float64 `json:"value"`
							DisplayValue string  `json:"displayValue"`
						} `json:"basic"`
					} `json:"averageLifespan"`
					BestSingleGameKills struct {
						StatID string `json:"statId"`
						Basic  struct {
							Value        float64 `json:"value"`
							DisplayValue string  `json:"displayValue"`
						} `json:"basic"`
					} `json:"bestSingleGameKills"`
					KillsDeathsRatio struct {
						StatID string `json:"statId"`
						Basic  struct {
							Value        float64 `json:"value"`
							DisplayValue string  `json:"displayValue"`
						} `json:"basic"`
					} `json:"killsDeathsRatio"`
					KillsDeathsAssists struct {
						StatID string `json:"statId"`
						Basic  struct {
							Value        float64 `json:"value"`
							DisplayValue string  `json:"displayValue"`
						} `json:"basic"`
					} `json:"killsDeathsAssists"`
					ObjectivesCompleted struct {
						StatID string `json:"statId"`
						Basic  struct {
							Value        float64 `json:"value"`
							DisplayValue string  `json:"displayValue"`
						} `json:"basic"`
						Pga struct {
							Value        float64 `json:"value"`
							DisplayValue string  `json:"displayValue"`
						} `json:"pga"`
					} `json:"objectivesCompleted"`
					PrecisionKills struct {
						StatID string `json:"statId"`
						Basic  struct {
							Value        float64 `json:"value"`
							DisplayValue string  `json:"displayValue"`
						} `json:"basic"`
						Pga struct {
							Value        float64 `json:"value"`
							DisplayValue string  `json:"displayValue"`
						} `json:"pga"`
					} `json:"precisionKills"`
					ResurrectionsPerformed struct {
						StatID string `json:"statId"`
						Basic  struct {
							Value        float64 `json:"value"`
							DisplayValue string  `json:"displayValue"`
						} `json:"basic"`
						Pga struct {
							Value        float64 `json:"value"`
							DisplayValue string  `json:"displayValue"`
						} `json:"pga"`
					} `json:"resurrectionsPerformed"`
					ResurrectionsReceived struct {
						StatID string `json:"statId"`
						Basic  struct {
							Value        float64 `json:"value"`
							DisplayValue string  `json:"displayValue"`
						} `json:"basic"`
						Pga struct {
							Value        float64 `json:"value"`
							DisplayValue string  `json:"displayValue"`
						} `json:"pga"`
					} `json:"resurrectionsReceived"`
					Suicides struct {
						StatID string `json:"statId"`
						Basic  struct {
							Value        float64 `json:"value"`
							DisplayValue string  `json:"displayValue"`
						} `json:"basic"`
						Pga struct {
							Value        float64 `json:"value"`
							DisplayValue string  `json:"displayValue"`
						} `json:"pga"`
					} `json:"suicides"`
					WeaponKillsAutoRifle struct {
						StatID string `json:"statId"`
						Basic  struct {
							Value        float64 `json:"value"`
							DisplayValue string  `json:"displayValue"`
						} `json:"basic"`
						Pga struct {
							Value        float64 `json:"value"`
							DisplayValue string  `json:"displayValue"`
						} `json:"pga"`
					} `json:"weaponKillsAutoRifle"`
					WeaponKillsFusionRifle struct {
						StatID string `json:"statId"`
						Basic  struct {
							Value        float64 `json:"value"`
							DisplayValue string  `json:"displayValue"`
						} `json:"basic"`
						Pga struct {
							Value        float64 `json:"value"`
							DisplayValue string  `json:"displayValue"`
						} `json:"pga"`
					} `json:"weaponKillsFusionRifle"`
					WeaponKillsHandCannon struct {
						StatID string `json:"statId"`
						Basic  struct {
							Value        float64 `json:"value"`
							DisplayValue string  `json:"displayValue"`
						} `json:"basic"`
						Pga struct {
							Value        float64 `json:"value"`
							DisplayValue string  `json:"displayValue"`
						} `json:"pga"`
					} `json:"weaponKillsHandCannon"`
					WeaponKillsMachinegun struct {
						StatID string `json:"statId"`
						Basic  struct {
							Value        float64 `json:"value"`
							DisplayValue string  `json:"displayValue"`
						} `json:"basic"`
						Pga struct {
							Value        float64 `json:"value"`
							DisplayValue string  `json:"displayValue"`
						} `json:"pga"`
					} `json:"weaponKillsMachinegun"`
					WeaponKillsPulseRifle struct {
						StatID string `json:"statId"`
						Basic  struct {
							Value        float64 `json:"value"`
							DisplayValue string  `json:"displayValue"`
						} `json:"basic"`
						Pga struct {
							Value        float64 `json:"value"`
							DisplayValue string  `json:"displayValue"`
						} `json:"pga"`
					} `json:"weaponKillsPulseRifle"`
					WeaponKillsRocketLauncher struct {
						StatID string `json:"statId"`
						Basic  struct {
							Value        float64 `json:"value"`
							DisplayValue string  `json:"displayValue"`
						} `json:"basic"`
						Pga struct {
							Value        float64 `json:"value"`
							DisplayValue string  `json:"displayValue"`
						} `json:"pga"`
					} `json:"weaponKillsRocketLauncher"`
					WeaponKillsScoutRifle struct {
						StatID string `json:"statId"`
						Basic  struct {
							Value        float64 `json:"value"`
							DisplayValue string  `json:"displayValue"`
						} `json:"basic"`
						Pga struct {
							Value        float64 `json:"value"`
							DisplayValue string  `json:"displayValue"`
						} `json:"pga"`
					} `json:"weaponKillsScoutRifle"`
					WeaponKillsShotgun struct {
						StatID string `json:"statId"`
						Basic  struct {
							Value        float64 `json:"value"`
							DisplayValue string  `json:"displayValue"`
						} `json:"basic"`
						Pga struct {
							Value        float64 `json:"value"`
							DisplayValue string  `json:"displayValue"`
						} `json:"pga"`
					} `json:"weaponKillsShotgun"`
					WeaponKillsSniper struct {
						StatID string `json:"statId"`
						Basic  struct {
							Value        float64 `json:"value"`
							DisplayValue string  `json:"displayValue"`
						} `json:"basic"`
						Pga struct {
							Value        float64 `json:"value"`
							DisplayValue string  `json:"displayValue"`
						} `json:"pga"`
					} `json:"weaponKillsSniper"`
					WeaponKillsSubmachinegun struct {
						StatID string `json:"statId"`
						Basic  struct {
							Value        float64 `json:"value"`
							DisplayValue string  `json:"displayValue"`
						} `json:"basic"`
						Pga struct {
							Value        float64 `json:"value"`
							DisplayValue string  `json:"displayValue"`
						} `json:"pga"`
					} `json:"weaponKillsSubmachinegun"`
					WeaponKillsRelic struct {
						StatID string `json:"statId"`
						Basic  struct {
							Value        float64 `json:"value"`
							DisplayValue string  `json:"displayValue"`
						} `json:"basic"`
						Pga struct {
							Value        float64 `json:"value"`
							DisplayValue string  `json:"displayValue"`
						} `json:"pga"`
					} `json:"weaponKillsRelic"`
					WeaponKillsSideArm struct {
						StatID string `json:"statId"`
						Basic  struct {
							Value        float64 `json:"value"`
							DisplayValue string  `json:"displayValue"`
						} `json:"basic"`
						Pga struct {
							Value        float64 `json:"value"`
							DisplayValue string  `json:"displayValue"`
						} `json:"pga"`
					} `json:"weaponKillsSideArm"`
					WeaponKillsSword struct {
						StatID string `json:"statId"`
						Basic  struct {
							Value        float64 `json:"value"`
							DisplayValue string  `json:"displayValue"`
						} `json:"basic"`
						Pga struct {
							Value        float64 `json:"value"`
							DisplayValue string  `json:"displayValue"`
						} `json:"pga"`
					} `json:"weaponKillsSword"`
					WeaponKillsAbility struct {
						StatID string `json:"statId"`
						Basic  struct {
							Value        float64 `json:"value"`
							DisplayValue string  `json:"displayValue"`
						} `json:"basic"`
						Pga struct {
							Value        float64 `json:"value"`
							DisplayValue string  `json:"displayValue"`
						} `json:"pga"`
					} `json:"weaponKillsAbility"`
					WeaponBestType struct {
						StatID string `json:"statId"`
						Basic  struct {
							Value        float64 `json:"value"`
							DisplayValue string  `json:"displayValue"`
						} `json:"basic"`
					} `json:"weaponBestType"`
					AllParticipantsCount struct {
						StatID string `json:"statId"`
						Basic  struct {
							Value        float64 `json:"value"`
							DisplayValue string  `json:"displayValue"`
						} `json:"basic"`
					} `json:"allParticipantsCount"`
					AllParticipantsTimePlayed struct {
						StatID string `json:"statId"`
						Basic  struct {
							Value        float64 `json:"value"`
							DisplayValue string  `json:"displayValue"`
						} `json:"basic"`
					} `json:"allParticipantsTimePlayed"`
					LongestKillSpree struct {
						StatID string `json:"statId"`
						Basic  struct {
							Value        float64 `json:"value"`
							DisplayValue string  `json:"displayValue"`
						} `json:"basic"`
					} `json:"longestKillSpree"`
					LongestSingleLife struct {
						StatID string `json:"statId"`
						Basic  struct {
							Value        float64 `json:"value"`
							DisplayValue string  `json:"displayValue"`
						} `json:"basic"`
					} `json:"longestSingleLife"`
					MostPrecisionKills struct {
						StatID string `json:"statId"`
						Basic  struct {
							Value        float64 `json:"value"`
							DisplayValue string  `json:"displayValue"`
						} `json:"basic"`
					} `json:"mostPrecisionKills"`
					OrbsDropped struct {
						StatID string `json:"statId"`
						Basic  struct {
							Value        float64 `json:"value"`
							DisplayValue string  `json:"displayValue"`
						} `json:"basic"`
						Pga struct {
							Value        float64 `json:"value"`
							DisplayValue string  `json:"displayValue"`
						} `json:"pga"`
					} `json:"orbsDropped"`
					OrbsGathered struct {
						StatID string `json:"statId"`
						Basic  struct {
							Value        float64 `json:"value"`
							DisplayValue string  `json:"displayValue"`
						} `json:"basic"`
						Pga struct {
							Value        float64 `json:"value"`
							DisplayValue string  `json:"displayValue"`
						} `json:"pga"`
					} `json:"orbsGathered"`
					PublicEventsCompleted struct {
						StatID string `json:"statId"`
						Basic  struct {
							Value        float64 `json:"value"`
							DisplayValue string  `json:"displayValue"`
						} `json:"basic"`
						Pga struct {
							Value        float64 `json:"value"`
							DisplayValue string  `json:"displayValue"`
						} `json:"pga"`
					} `json:"publicEventsCompleted"`
					RemainingTimeAfterQuitSeconds struct {
						StatID string `json:"statId"`
						Basic  struct {
							Value        float64 `json:"value"`
							DisplayValue string  `json:"displayValue"`
						} `json:"basic"`
						Pga struct {
							Value        float64 `json:"value"`
							DisplayValue string  `json:"displayValue"`
						} `json:"pga"`
					} `json:"remainingTimeAfterQuitSeconds"`
					TotalActivityDurationSeconds struct {
						StatID string `json:"statId"`
						Basic  struct {
							Value        float64 `json:"value"`
							DisplayValue string  `json:"displayValue"`
						} `json:"basic"`
						Pga struct {
							Value        float64 `json:"value"`
							DisplayValue string  `json:"displayValue"`
						} `json:"pga"`
					} `json:"totalActivityDurationSeconds"`
					FastestCompletionMs struct {
						StatID string `json:"statId"`
						Basic  struct {
							Value        float64 `json:"value"`
							DisplayValue string  `json:"displayValue"`
						} `json:"basic"`
					} `json:"fastestCompletionMs"`
					LongestKillDistance struct {
						StatID string `json:"statId"`
						Basic  struct {
							Value        float64 `json:"value"`
							DisplayValue string  `json:"displayValue"`
						} `json:"basic"`
					} `json:"longestKillDistance"`
					HighestCharacterLevel struct {
						StatID string `json:"statId"`
						Basic  struct {
							Value        float64 `json:"value"`
							DisplayValue string  `json:"displayValue"`
						} `json:"basic"`
					} `json:"highestCharacterLevel"`
					HighestLightLevel struct {
						StatID string `json:"statId"`
						Basic  struct {
							Value        float64 `json:"value"`
							DisplayValue string  `json:"displayValue"`
						} `json:"basic"`
					} `json:"highestLightLevel"`
				} `json:"allTime"`
			} `json:"allPvE"`
		} `json:"results"`
		Merged struct {
			AllTime struct {
				ActivitiesCleared struct {
					StatID string `json:"statId"`
					Basic  struct {
						Value        float64 `json:"value"`
						DisplayValue string  `json:"displayValue"`
					} `json:"basic"`
				} `json:"activitiesCleared"`
				WeaponKillsSuper struct {
					StatID string `json:"statId"`
					Basic  struct {
						Value        float64 `json:"value"`
						DisplayValue string  `json:"displayValue"`
					} `json:"basic"`
					Pga struct {
						Value        float64 `json:"value"`
						DisplayValue string  `json:"displayValue"`
					} `json:"pga"`
				} `json:"weaponKillsSuper"`
				ActivitiesEntered struct {
					StatID string `json:"statId"`
					Basic  struct {
						Value        float64 `json:"value"`
						DisplayValue string  `json:"displayValue"`
					} `json:"basic"`
				} `json:"activitiesEntered"`
				WeaponKillsMelee struct {
					StatID string `json:"statId"`
					Basic  struct {
						Value        float64 `json:"value"`
						DisplayValue string  `json:"displayValue"`
					} `json:"basic"`
					Pga struct {
						Value        float64 `json:"value"`
						DisplayValue string  `json:"displayValue"`
					} `json:"pga"`
				} `json:"weaponKillsMelee"`
				WeaponKillsGrenade struct {
					StatID string `json:"statId"`
					Basic  struct {
						Value        float64 `json:"value"`
						DisplayValue string  `json:"displayValue"`
					} `json:"basic"`
					Pga struct {
						Value        float64 `json:"value"`
						DisplayValue string  `json:"displayValue"`
					} `json:"pga"`
				} `json:"weaponKillsGrenade"`
				AbilityKills struct {
					StatID string `json:"statId"`
					Basic  struct {
						Value        float64 `json:"value"`
						DisplayValue string  `json:"displayValue"`
					} `json:"basic"`
					Pga struct {
						Value        float64 `json:"value"`
						DisplayValue string  `json:"displayValue"`
					} `json:"pga"`
				} `json:"abilityKills"`
				Assists struct {
					StatID string `json:"statId"`
					Basic  struct {
						Value        float64 `json:"value"`
						DisplayValue string  `json:"displayValue"`
					} `json:"basic"`
					Pga struct {
						Value        float64 `json:"value"`
						DisplayValue string  `json:"displayValue"`
					} `json:"pga"`
				} `json:"assists"`
				TotalDeathDistance struct {
					StatID string `json:"statId"`
					Basic  struct {
						Value        float64 `json:"value"`
						DisplayValue string  `json:"displayValue"`
					} `json:"basic"`
				} `json:"totalDeathDistance"`
				AverageDeathDistance struct {
					StatID string `json:"statId"`
					Basic  struct {
						Value        float64 `json:"value"`
						DisplayValue string  `json:"displayValue"`
					} `json:"basic"`
				} `json:"averageDeathDistance"`
				TotalKillDistance struct {
					StatID string `json:"statId"`
					Basic  struct {
						Value        float64 `json:"value"`
						DisplayValue string  `json:"displayValue"`
					} `json:"basic"`
				} `json:"totalKillDistance"`
				Kills struct {
					StatID string `json:"statId"`
					Basic  struct {
						Value        float64 `json:"value"`
						DisplayValue string  `json:"displayValue"`
					} `json:"basic"`
					Pga struct {
						Value        float64 `json:"value"`
						DisplayValue string  `json:"displayValue"`
					} `json:"pga"`
				} `json:"kills"`
				AverageKillDistance struct {
					StatID string `json:"statId"`
					Basic  struct {
						Value        float64 `json:"value"`
						DisplayValue string  `json:"displayValue"`
					} `json:"basic"`
				} `json:"averageKillDistance"`
				SecondsPlayed struct {
					StatID string `json:"statId"`
					Basic  struct {
						Value        float64 `json:"value"`
						DisplayValue string  `json:"displayValue"`
					} `json:"basic"`
					Pga struct {
						Value        float64 `json:"value"`
						DisplayValue string  `json:"displayValue"`
					} `json:"pga"`
				} `json:"secondsPlayed"`
				Deaths struct {
					StatID string `json:"statId"`
					Basic  struct {
						Value        float64 `json:"value"`
						DisplayValue string  `json:"displayValue"`
					} `json:"basic"`
					Pga struct {
						Value        float64 `json:"value"`
						DisplayValue string  `json:"displayValue"`
					} `json:"pga"`
				} `json:"deaths"`
				AverageLifespan struct {
					StatID string `json:"statId"`
					Basic  struct {
						Value        float64 `json:"value"`
						DisplayValue string  `json:"displayValue"`
					} `json:"basic"`
				} `json:"averageLifespan"`
				BestSingleGameKills struct {
					StatID string `json:"statId"`
					Basic  struct {
						Value        float64 `json:"value"`
						DisplayValue string  `json:"displayValue"`
					} `json:"basic"`
				} `json:"bestSingleGameKills"`
				KillsDeathsRatio struct {
					StatID string `json:"statId"`
					Basic  struct {
						Value        float64 `json:"value"`
						DisplayValue string  `json:"displayValue"`
					} `json:"basic"`
				} `json:"killsDeathsRatio"`
				KillsDeathsAssists struct {
					StatID string `json:"statId"`
					Basic  struct {
						Value        float64 `json:"value"`
						DisplayValue string  `json:"displayValue"`
					} `json:"basic"`
				} `json:"killsDeathsAssists"`
				ObjectivesCompleted struct {
					StatID string `json:"statId"`
					Basic  struct {
						Value        float64 `json:"value"`
						DisplayValue string  `json:"displayValue"`
					} `json:"basic"`
					Pga struct {
						Value        float64 `json:"value"`
						DisplayValue string  `json:"displayValue"`
					} `json:"pga"`
				} `json:"objectivesCompleted"`
				PrecisionKills struct {
					StatID string `json:"statId"`
					Basic  struct {
						Value        float64 `json:"value"`
						DisplayValue string  `json:"displayValue"`
					} `json:"basic"`
					Pga struct {
						Value        float64 `json:"value"`
						DisplayValue string  `json:"displayValue"`
					} `json:"pga"`
				} `json:"precisionKills"`
				ResurrectionsPerformed struct {
					StatID string `json:"statId"`
					Basic  struct {
						Value        float64 `json:"value"`
						DisplayValue string  `json:"displayValue"`
					} `json:"basic"`
					Pga struct {
						Value        float64 `json:"value"`
						DisplayValue string  `json:"displayValue"`
					} `json:"pga"`
				} `json:"resurrectionsPerformed"`
				ResurrectionsReceived struct {
					StatID string `json:"statId"`
					Basic  struct {
						Value        float64 `json:"value"`
						DisplayValue string  `json:"displayValue"`
					} `json:"basic"`
					Pga struct {
						Value        float64 `json:"value"`
						DisplayValue string  `json:"displayValue"`
					} `json:"pga"`
				} `json:"resurrectionsReceived"`
				Suicides struct {
					StatID string `json:"statId"`
					Basic  struct {
						Value        float64 `json:"value"`
						DisplayValue string  `json:"displayValue"`
					} `json:"basic"`
					Pga struct {
						Value        float64 `json:"value"`
						DisplayValue string  `json:"displayValue"`
					} `json:"pga"`
				} `json:"suicides"`
				WeaponKillsAutoRifle struct {
					StatID string `json:"statId"`
					Basic  struct {
						Value        float64 `json:"value"`
						DisplayValue string  `json:"displayValue"`
					} `json:"basic"`
					Pga struct {
						Value        float64 `json:"value"`
						DisplayValue string  `json:"displayValue"`
					} `json:"pga"`
				} `json:"weaponKillsAutoRifle"`
				WeaponKillsFusionRifle struct {
					StatID string `json:"statId"`
					Basic  struct {
						Value        float64 `json:"value"`
						DisplayValue string  `json:"displayValue"`
					} `json:"basic"`
					Pga struct {
						Value        float64 `json:"value"`
						DisplayValue string  `json:"displayValue"`
					} `json:"pga"`
				} `json:"weaponKillsFusionRifle"`
				WeaponKillsHandCannon struct {
					StatID string `json:"statId"`
					Basic  struct {
						Value        float64 `json:"value"`
						DisplayValue string  `json:"displayValue"`
					} `json:"basic"`
					Pga struct {
						Value        float64 `json:"value"`
						DisplayValue string  `json:"displayValue"`
					} `json:"pga"`
				} `json:"weaponKillsHandCannon"`
				WeaponKillsMachinegun struct {
					StatID string `json:"statId"`
					Basic  struct {
						Value        float64 `json:"value"`
						DisplayValue string  `json:"displayValue"`
					} `json:"basic"`
					Pga struct {
						Value        float64 `json:"value"`
						DisplayValue string  `json:"displayValue"`
					} `json:"pga"`
				} `json:"weaponKillsMachinegun"`
				WeaponKillsPulseRifle struct {
					StatID string `json:"statId"`
					Basic  struct {
						Value        float64 `json:"value"`
						DisplayValue string  `json:"displayValue"`
					} `json:"basic"`
					Pga struct {
						Value        float64 `json:"value"`
						DisplayValue string  `json:"displayValue"`
					} `json:"pga"`
				} `json:"weaponKillsPulseRifle"`
				WeaponKillsRocketLauncher struct {
					StatID string `json:"statId"`
					Basic  struct {
						Value        float64 `json:"value"`
						DisplayValue string  `json:"displayValue"`
					} `json:"basic"`
					Pga struct {
						Value        float64 `json:"value"`
						DisplayValue string  `json:"displayValue"`
					} `json:"pga"`
				} `json:"weaponKillsRocketLauncher"`
				WeaponKillsScoutRifle struct {
					StatID string `json:"statId"`
					Basic  struct {
						Value        float64 `json:"value"`
						DisplayValue string  `json:"displayValue"`
					} `json:"basic"`
					Pga struct {
						Value        float64 `json:"value"`
						DisplayValue string  `json:"displayValue"`
					} `json:"pga"`
				} `json:"weaponKillsScoutRifle"`
				WeaponKillsShotgun struct {
					StatID string `json:"statId"`
					Basic  struct {
						Value        float64 `json:"value"`
						DisplayValue string  `json:"displayValue"`
					} `json:"basic"`
					Pga struct {
						Value        float64 `json:"value"`
						DisplayValue string  `json:"displayValue"`
					} `json:"pga"`
				} `json:"weaponKillsShotgun"`
				WeaponKillsSniper struct {
					StatID string `json:"statId"`
					Basic  struct {
						Value        float64 `json:"value"`
						DisplayValue string  `json:"displayValue"`
					} `json:"basic"`
					Pga struct {
						Value        float64 `json:"value"`
						DisplayValue string  `json:"displayValue"`
					} `json:"pga"`
				} `json:"weaponKillsSniper"`
				WeaponKillsSubmachinegun struct {
					StatID string `json:"statId"`
					Basic  struct {
						Value        float64 `json:"value"`
						DisplayValue string  `json:"displayValue"`
					} `json:"basic"`
					Pga struct {
						Value        float64 `json:"value"`
						DisplayValue string  `json:"displayValue"`
					} `json:"pga"`
				} `json:"weaponKillsSubmachinegun"`
				WeaponKillsRelic struct {
					StatID string `json:"statId"`
					Basic  struct {
						Value        float64 `json:"value"`
						DisplayValue string  `json:"displayValue"`
					} `json:"basic"`
					Pga struct {
						Value        float64 `json:"value"`
						DisplayValue string  `json:"displayValue"`
					} `json:"pga"`
				} `json:"weaponKillsRelic"`
				WeaponKillsSideArm struct {
					StatID string `json:"statId"`
					Basic  struct {
						Value        float64 `json:"value"`
						DisplayValue string  `json:"displayValue"`
					} `json:"basic"`
					Pga struct {
						Value        float64 `json:"value"`
						DisplayValue string  `json:"displayValue"`
					} `json:"pga"`
				} `json:"weaponKillsSideArm"`
				WeaponKillsSword struct {
					StatID string `json:"statId"`
					Basic  struct {
						Value        float64 `json:"value"`
						DisplayValue string  `json:"displayValue"`
					} `json:"basic"`
					Pga struct {
						Value        float64 `json:"value"`
						DisplayValue string  `json:"displayValue"`
					} `json:"pga"`
				} `json:"weaponKillsSword"`
				WeaponKillsAbility struct {
					StatID string `json:"statId"`
					Basic  struct {
						Value        float64 `json:"value"`
						DisplayValue string  `json:"displayValue"`
					} `json:"basic"`
					Pga struct {
						Value        float64 `json:"value"`
						DisplayValue string  `json:"displayValue"`
					} `json:"pga"`
				} `json:"weaponKillsAbility"`
				WeaponBestType struct {
					StatID string `json:"statId"`
					Basic  struct {
						Value        float64 `json:"value"`
						DisplayValue string  `json:"displayValue"`
					} `json:"basic"`
				} `json:"weaponBestType"`
				AllParticipantsCount struct {
					StatID string `json:"statId"`
					Basic  struct {
						Value        float64 `json:"value"`
						DisplayValue string  `json:"displayValue"`
					} `json:"basic"`
				} `json:"allParticipantsCount"`
				AllParticipantsTimePlayed struct {
					StatID string `json:"statId"`
					Basic  struct {
						Value        float64 `json:"value"`
						DisplayValue string  `json:"displayValue"`
					} `json:"basic"`
				} `json:"allParticipantsTimePlayed"`
				LongestKillSpree struct {
					StatID string `json:"statId"`
					Basic  struct {
						Value        float64 `json:"value"`
						DisplayValue string  `json:"displayValue"`
					} `json:"basic"`
				} `json:"longestKillSpree"`
				LongestSingleLife struct {
					StatID string `json:"statId"`
					Basic  struct {
						Value        float64 `json:"value"`
						DisplayValue string  `json:"displayValue"`
					} `json:"basic"`
				} `json:"longestSingleLife"`
				MostPrecisionKills struct {
					StatID string `json:"statId"`
					Basic  struct {
						Value        float64 `json:"value"`
						DisplayValue string  `json:"displayValue"`
					} `json:"basic"`
				} `json:"mostPrecisionKills"`
				OrbsDropped struct {
					StatID string `json:"statId"`
					Basic  struct {
						Value        float64 `json:"value"`
						DisplayValue string  `json:"displayValue"`
					} `json:"basic"`
					Pga struct {
						Value        float64 `json:"value"`
						DisplayValue string  `json:"displayValue"`
					} `json:"pga"`
				} `json:"orbsDropped"`
				OrbsGathered struct {
					StatID string `json:"statId"`
					Basic  struct {
						Value        float64 `json:"value"`
						DisplayValue string  `json:"displayValue"`
					} `json:"basic"`
					Pga struct {
						Value        float64 `json:"value"`
						DisplayValue string  `json:"displayValue"`
					} `json:"pga"`
				} `json:"orbsGathered"`
				PublicEventsCompleted struct {
					StatID string `json:"statId"`
					Basic  struct {
						Value        float64 `json:"value"`
						DisplayValue string  `json:"displayValue"`
					} `json:"basic"`
					Pga struct {
						Value        float64 `json:"value"`
						DisplayValue string  `json:"displayValue"`
					} `json:"pga"`
				} `json:"publicEventsCompleted"`
				RemainingTimeAfterQuitSeconds struct {
					StatID string `json:"statId"`
					Basic  struct {
						Value        float64 `json:"value"`
						DisplayValue string  `json:"displayValue"`
					} `json:"basic"`
					Pga struct {
						Value        float64 `json:"value"`
						DisplayValue string  `json:"displayValue"`
					} `json:"pga"`
				} `json:"remainingTimeAfterQuitSeconds"`
				TotalActivityDurationSeconds struct {
					StatID string `json:"statId"`
					Basic  struct {
						Value        float64 `json:"value"`
						DisplayValue string  `json:"displayValue"`
					} `json:"basic"`
					Pga struct {
						Value        float64 `json:"value"`
						DisplayValue string  `json:"displayValue"`
					} `json:"pga"`
				} `json:"totalActivityDurationSeconds"`
				FastestCompletionMs struct {
					StatID string `json:"statId"`
					Basic  struct {
						Value        float64 `json:"value"`
						DisplayValue string  `json:"displayValue"`
					} `json:"basic"`
				} `json:"fastestCompletionMs"`
				LongestKillDistance struct {
					StatID string `json:"statId"`
					Basic  struct {
						Value        float64 `json:"value"`
						DisplayValue string  `json:"displayValue"`
					} `json:"basic"`
				} `json:"longestKillDistance"`
				HighestCharacterLevel struct {
					StatID string `json:"statId"`
					Basic  struct {
						Value        float64 `json:"value"`
						DisplayValue string  `json:"displayValue"`
					} `json:"basic"`
				} `json:"highestCharacterLevel"`
				HighestLightLevel struct {
					StatID string `json:"statId"`
					Basic  struct {
						Value        float64 `json:"value"`
						DisplayValue string  `json:"displayValue"`
					} `json:"basic"`
				} `json:"highestLightLevel"`
				ActivitiesWon struct {
					StatID string `json:"statId"`
					Basic  struct {
						Value        float64 `json:"value"`
						DisplayValue string  `json:"displayValue"`
					} `json:"basic"`
				} `json:"activitiesWon"`
				Score struct {
					StatID string `json:"statId"`
					Basic  struct {
						Value        float64 `json:"value"`
						DisplayValue string  `json:"displayValue"`
					} `json:"basic"`
					Pga struct {
						Value        float64 `json:"value"`
						DisplayValue string  `json:"displayValue"`
					} `json:"pga"`
				} `json:"score"`
				AverageScorePerKill struct {
					StatID string `json:"statId"`
					Basic  struct {
						Value        float64 `json:"value"`
						DisplayValue string  `json:"displayValue"`
					} `json:"basic"`
				} `json:"averageScorePerKill"`
				AverageScorePerLife struct {
					StatID string `json:"statId"`
					Basic  struct {
						Value        float64 `json:"value"`
						DisplayValue string  `json:"displayValue"`
					} `json:"basic"`
				} `json:"averageScorePerLife"`
				BestSingleGameScore struct {
					StatID string `json:"statId"`
					Basic  struct {
						Value        float64 `json:"value"`
						DisplayValue string  `json:"displayValue"`
					} `json:"basic"`
				} `json:"bestSingleGameScore"`
				WinLossRatio struct {
					StatID string `json:"statId"`
					Basic  struct {
						Value        float64 `json:"value"`
						DisplayValue string  `json:"displayValue"`
					} `json:"basic"`
				} `json:"winLossRatio"`
				AllParticipantsScore struct {
					StatID string `json:"statId"`
					Basic  struct {
						Value        float64 `json:"value"`
						DisplayValue string  `json:"displayValue"`
					} `json:"basic"`
				} `json:"allParticipantsScore"`
				TeamScore struct {
					StatID string `json:"statId"`
					Basic  struct {
						Value        float64 `json:"value"`
						DisplayValue string  `json:"displayValue"`
					} `json:"basic"`
					Pga struct {
						Value        float64 `json:"value"`
						DisplayValue string  `json:"displayValue"`
					} `json:"pga"`
				} `json:"teamScore"`
				CombatRating struct {
					StatID string `json:"statId"`
					Basic  struct {
						Value        float64 `json:"value"`
						DisplayValue string  `json:"displayValue"`
					} `json:"basic"`
				} `json:"combatRating"`
			} `json:"allTime"`
		} `json:"merged"`
	} `json:"characters"`
}

type Manifest struct {
	Version                  string `json:"version"`
	MobileAssetContentPath   string `json:"mobileAssetContentPath"`
	MobileGearAssetDataBases []struct {
		Version int    `json:"version"`
		Path    string `json:"path"`
	} `json:"mobileGearAssetDataBases"`
	MobileWorldContentPaths struct {
		En    string `json:"en"`
		Fr    string `json:"fr"`
		Es    string `json:"es"`
		De    string `json:"de"`
		It    string `json:"it"`
		Ja    string `json:"ja"`
		PtBr  string `json:"pt-br"`
		EsMx  string `json:"es-mx"`
		Ru    string `json:"ru"`
		Pl    string `json:"pl"`
		ZhCht string `json:"zh-cht"`
	} `json:"mobileWorldContentPaths"`
	MobileClanBannerDatabasePath string `json:"mobileClanBannerDatabasePath"`
	MobileGearCDN                struct {
		Geometry    string `json:"Geometry"`
		Texture     string `json:"Texture"`
		PlateRegion string `json:"PlateRegion"`
		Gear        string `json:"Gear"`
		Shader      string `json:"Shader"`
	} `json:"mobileGearCDN"`
}

type HistoricalStatsDefinition struct {
	AbilityKills struct {
		Category        int    `json:"category"`
		Group           int    `json:"group"`
		Modes           []int  `json:"modes"`
		PeriodTypes     []int  `json:"periodTypes"`
		StatDescription string `json:"statDescription"`
		StatID          string `json:"statId"`
		StatName        string `json:"statName"`
		UnitLabel       string `json:"unitLabel"`
		UnitType        int    `json:"unitType"`
		Weight          int    `json:"weight"`
	} `json:"abilityKills"`
	ActivitiesCleared struct {
		Category        int    `json:"category"`
		Group           int    `json:"group"`
		Modes           []int  `json:"modes"`
		PeriodTypes     []int  `json:"periodTypes"`
		StatDescription string `json:"statDescription"`
		StatID          string `json:"statId"`
		StatName        string `json:"statName"`
		UnitLabel       string `json:"unitLabel"`
		UnitType        int    `json:"unitType"`
		Weight          int    `json:"weight"`
	} `json:"activitiesCleared"`
	ActivitiesEntered struct {
		Category        int    `json:"category"`
		Group           int    `json:"group"`
		Modes           []int  `json:"modes"`
		PeriodTypes     []int  `json:"periodTypes"`
		StatDescription string `json:"statDescription"`
		StatID          string `json:"statId"`
		StatName        string `json:"statName"`
		UnitLabel       string `json:"unitLabel"`
		UnitType        int    `json:"unitType"`
		Weight          int    `json:"weight"`
	} `json:"activitiesEntered"`
	ActivitiesWon struct {
		Category    int    `json:"category"`
		Group       int    `json:"group"`
		Modes       []int  `json:"modes"`
		PeriodTypes []int  `json:"periodTypes"`
		StatID      string `json:"statId"`
		StatName    string `json:"statName"`
		UnitLabel   string `json:"unitLabel"`
		UnitType    int    `json:"unitType"`
		Weight      int    `json:"weight"`
	} `json:"activitiesWon"`
	ActivityAssists struct {
		Category    int    `json:"category"`
		Group       int    `json:"group"`
		Modes       []int  `json:"modes"`
		PeriodTypes []int  `json:"periodTypes"`
		StatID      string `json:"statId"`
		StatName    string `json:"statName"`
		UnitLabel   string `json:"unitLabel"`
		UnitType    int    `json:"unitType"`
		Weight      int    `json:"weight"`
	} `json:"activityAssists"`
	ActivityBestGoalsHit struct {
		Category    int    `json:"category"`
		Group       int    `json:"group"`
		Modes       []int  `json:"modes"`
		PeriodTypes []int  `json:"periodTypes"`
		StatID      string `json:"statId"`
		StatName    string `json:"statName"`
		UnitLabel   string `json:"unitLabel"`
		UnitType    int    `json:"unitType"`
		Weight      int    `json:"weight"`
	} `json:"activityBestGoalsHit"`
	ActivityBestSingleGameScore struct {
		Category    int    `json:"category"`
		Group       int    `json:"group"`
		Modes       []int  `json:"modes"`
		PeriodTypes []int  `json:"periodTypes"`
		StatID      string `json:"statId"`
		StatName    string `json:"statName"`
		UnitLabel   string `json:"unitLabel"`
		UnitType    int    `json:"unitType"`
		Weight      int    `json:"weight"`
	} `json:"activityBestSingleGameScore"`
	ActivityCompletedFailures struct {
		Category    int    `json:"category"`
		Group       int    `json:"group"`
		Modes       []int  `json:"modes"`
		PeriodTypes []int  `json:"periodTypes"`
		StatID      string `json:"statId"`
		StatName    string `json:"statName"`
		UnitLabel   string `json:"unitLabel"`
		UnitType    int    `json:"unitType"`
		Weight      int    `json:"weight"`
	} `json:"activityCompletedFailures"`
	ActivityCompletions struct {
		Category    int    `json:"category"`
		Group       int    `json:"group"`
		Modes       []int  `json:"modes"`
		PeriodTypes []int  `json:"periodTypes"`
		StatID      string `json:"statId"`
		StatName    string `json:"statName"`
		UnitLabel   string `json:"unitLabel"`
		UnitType    int    `json:"unitType"`
		Weight      int    `json:"weight"`
	} `json:"activityCompletions"`
	ActivityDeaths struct {
		Category    int    `json:"category"`
		Group       int    `json:"group"`
		Modes       []int  `json:"modes"`
		PeriodTypes []int  `json:"periodTypes"`
		StatID      string `json:"statId"`
		StatName    string `json:"statName"`
		UnitLabel   string `json:"unitLabel"`
		UnitType    int    `json:"unitType"`
		Weight      int    `json:"weight"`
	} `json:"activityDeaths"`
	ActivityDurationSeconds struct {
		Category    int    `json:"category"`
		Group       int    `json:"group"`
		Modes       []int  `json:"modes"`
		PeriodTypes []int  `json:"periodTypes"`
		StatID      string `json:"statId"`
		StatName    string `json:"statName"`
		UnitLabel   string `json:"unitLabel"`
		UnitType    int    `json:"unitType"`
		Weight      int    `json:"weight"`
	} `json:"activityDurationSeconds"`
	ActivityFastestObjectiveCompletionMs struct {
		Category    int    `json:"category"`
		Group       int    `json:"group"`
		Modes       []int  `json:"modes"`
		PeriodTypes []int  `json:"periodTypes"`
		StatID      string `json:"statId"`
		StatName    string `json:"statName"`
		UnitLabel   string `json:"unitLabel"`
		UnitType    int    `json:"unitType"`
		Weight      int    `json:"weight"`
	} `json:"activityFastestObjectiveCompletionMs"`
	ActivityGatesHit struct {
		Category    int           `json:"category"`
		Group       int           `json:"group"`
		Modes       []interface{} `json:"modes"`
		PeriodTypes []int         `json:"periodTypes"`
		StatID      string        `json:"statId"`
		StatName    string        `json:"statName"`
		UnitLabel   string        `json:"unitLabel"`
		UnitType    int           `json:"unitType"`
		Weight      int           `json:"weight"`
	} `json:"activityGatesHit"`
	ActivityGoalsMissed struct {
		Category    int    `json:"category"`
		Group       int    `json:"group"`
		Modes       []int  `json:"modes"`
		PeriodTypes []int  `json:"periodTypes"`
		StatID      string `json:"statId"`
		StatName    string `json:"statName"`
		UnitLabel   string `json:"unitLabel"`
		UnitType    int    `json:"unitType"`
		Weight      int    `json:"weight"`
	} `json:"activityGoalsMissed"`
	ActivityKills struct {
		Category    int    `json:"category"`
		Group       int    `json:"group"`
		Modes       []int  `json:"modes"`
		PeriodTypes []int  `json:"periodTypes"`
		StatID      string `json:"statId"`
		StatName    string `json:"statName"`
		UnitLabel   string `json:"unitLabel"`
		UnitType    int    `json:"unitType"`
		Weight      int    `json:"weight"`
	} `json:"activityKills"`
	ActivityKillsDeathsAssists struct {
		Category    int    `json:"category"`
		Group       int    `json:"group"`
		Modes       []int  `json:"modes"`
		PeriodTypes []int  `json:"periodTypes"`
		StatID      string `json:"statId"`
		StatName    string `json:"statName"`
		UnitLabel   string `json:"unitLabel"`
		UnitType    int    `json:"unitType"`
		Weight      int    `json:"weight"`
	} `json:"activityKillsDeathsAssists"`
	ActivityKillsDeathsRatio struct {
		Category    int    `json:"category"`
		Group       int    `json:"group"`
		Modes       []int  `json:"modes"`
		PeriodTypes []int  `json:"periodTypes"`
		StatID      string `json:"statId"`
		StatName    string `json:"statName"`
		UnitLabel   string `json:"unitLabel"`
		UnitType    int    `json:"unitType"`
		Weight      int    `json:"weight"`
	} `json:"activityKillsDeathsRatio"`
	ActivityPrecisionKills struct {
		Category    int    `json:"category"`
		Group       int    `json:"group"`
		Modes       []int  `json:"modes"`
		PeriodTypes []int  `json:"periodTypes"`
		StatID      string `json:"statId"`
		StatName    string `json:"statName"`
		UnitLabel   string `json:"unitLabel"`
		UnitType    int    `json:"unitType"`
		Weight      int    `json:"weight"`
	} `json:"activityPrecisionKills"`
	ActivitySecondsPlayed struct {
		Category    int    `json:"category"`
		Group       int    `json:"group"`
		Modes       []int  `json:"modes"`
		PeriodTypes []int  `json:"periodTypes"`
		StatID      string `json:"statId"`
		StatName    string `json:"statName"`
		UnitLabel   string `json:"unitLabel"`
		UnitType    int    `json:"unitType"`
		Weight      int    `json:"weight"`
	} `json:"activitySecondsPlayed"`
	ActivitySpecialActions struct {
		Category    int    `json:"category"`
		Group       int    `json:"group"`
		Modes       []int  `json:"modes"`
		PeriodTypes []int  `json:"periodTypes"`
		StatID      string `json:"statId"`
		StatName    string `json:"statName"`
		UnitLabel   string `json:"unitLabel"`
		UnitType    int    `json:"unitType"`
		Weight      int    `json:"weight"`
	} `json:"activitySpecialActions"`
	ActivitySpecialScore struct {
		Category    int    `json:"category"`
		Group       int    `json:"group"`
		Modes       []int  `json:"modes"`
		PeriodTypes []int  `json:"periodTypes"`
		StatID      string `json:"statId"`
		StatName    string `json:"statName"`
		UnitLabel   string `json:"unitLabel"`
		UnitType    int    `json:"unitType"`
		Weight      int    `json:"weight"`
	} `json:"activitySpecialScore"`
	ActivityWins struct {
		Category    int    `json:"category"`
		Group       int    `json:"group"`
		Modes       []int  `json:"modes"`
		PeriodTypes []int  `json:"periodTypes"`
		StatID      string `json:"statId"`
		StatName    string `json:"statName"`
		UnitLabel   string `json:"unitLabel"`
		UnitType    int    `json:"unitType"`
		Weight      int    `json:"weight"`
	} `json:"activityWins"`
	AllMedalsEarned struct {
		Category    int    `json:"category"`
		Group       int    `json:"group"`
		Modes       []int  `json:"modes"`
		PeriodTypes []int  `json:"periodTypes"`
		StatID      string `json:"statId"`
		StatName    string `json:"statName"`
		UnitLabel   string `json:"unitLabel"`
		UnitType    int    `json:"unitType"`
		Weight      int    `json:"weight"`
	} `json:"allMedalsEarned"`
	AllMedalsScore struct {
		Category    int    `json:"category"`
		Group       int    `json:"group"`
		Modes       []int  `json:"modes"`
		PeriodTypes []int  `json:"periodTypes"`
		StatID      string `json:"statId"`
		StatName    string `json:"statName"`
		UnitLabel   string `json:"unitLabel"`
		UnitType    int    `json:"unitType"`
		Weight      int    `json:"weight"`
	} `json:"allMedalsScore"`
	AllParticipantsCount struct {
		Category    int    `json:"category"`
		Group       int    `json:"group"`
		Modes       []int  `json:"modes"`
		PeriodTypes []int  `json:"periodTypes"`
		StatID      string `json:"statId"`
		StatName    string `json:"statName"`
		UnitLabel   string `json:"unitLabel"`
		UnitType    int    `json:"unitType"`
		Weight      int    `json:"weight"`
	} `json:"allParticipantsCount"`
	AllParticipantsScore struct {
		Category    int    `json:"category"`
		Group       int    `json:"group"`
		Modes       []int  `json:"modes"`
		PeriodTypes []int  `json:"periodTypes"`
		StatID      string `json:"statId"`
		StatName    string `json:"statName"`
		UnitLabel   string `json:"unitLabel"`
		UnitType    int    `json:"unitType"`
		Weight      int    `json:"weight"`
	} `json:"allParticipantsScore"`
	AllParticipantsTimePlayed struct {
		Category    int    `json:"category"`
		Group       int    `json:"group"`
		Modes       []int  `json:"modes"`
		PeriodTypes []int  `json:"periodTypes"`
		StatID      string `json:"statId"`
		StatName    string `json:"statName"`
		UnitLabel   string `json:"unitLabel"`
		UnitType    int    `json:"unitType"`
		Weight      int    `json:"weight"`
	} `json:"allParticipantsTimePlayed"`
	Assists struct {
		Category        int    `json:"category"`
		Group           int    `json:"group"`
		Modes           []int  `json:"modes"`
		PeriodTypes     []int  `json:"periodTypes"`
		StatDescription string `json:"statDescription"`
		StatID          string `json:"statId"`
		StatName        string `json:"statName"`
		UnitLabel       string `json:"unitLabel"`
		UnitType        int    `json:"unitType"`
		Weight          int    `json:"weight"`
	} `json:"assists"`
	AverageDeathDistance struct {
		Category    int    `json:"category"`
		Group       int    `json:"group"`
		Modes       []int  `json:"modes"`
		PeriodTypes []int  `json:"periodTypes"`
		StatID      string `json:"statId"`
		StatName    string `json:"statName"`
		UnitLabel   string `json:"unitLabel"`
		UnitType    int    `json:"unitType"`
		Weight      int    `json:"weight"`
	} `json:"averageDeathDistance"`
	AverageKillDistance struct {
		Category    int    `json:"category"`
		Group       int    `json:"group"`
		Modes       []int  `json:"modes"`
		PeriodTypes []int  `json:"periodTypes"`
		StatID      string `json:"statId"`
		StatName    string `json:"statName"`
		UnitLabel   string `json:"unitLabel"`
		UnitType    int    `json:"unitType"`
		Weight      int    `json:"weight"`
	} `json:"averageKillDistance"`
	AverageLifespan struct {
		Category    int    `json:"category"`
		Group       int    `json:"group"`
		Modes       []int  `json:"modes"`
		PeriodTypes []int  `json:"periodTypes"`
		StatID      string `json:"statId"`
		StatName    string `json:"statName"`
		UnitLabel   string `json:"unitLabel"`
		UnitType    int    `json:"unitType"`
		Weight      int    `json:"weight"`
	} `json:"averageLifespan"`
	AverageScorePerKill struct {
		Category    int    `json:"category"`
		Group       int    `json:"group"`
		Modes       []int  `json:"modes"`
		PeriodTypes []int  `json:"periodTypes"`
		StatID      string `json:"statId"`
		StatName    string `json:"statName"`
		UnitLabel   string `json:"unitLabel"`
		UnitType    int    `json:"unitType"`
		Weight      int    `json:"weight"`
	} `json:"averageScorePerKill"`
	AverageScorePerLife struct {
		Category    int    `json:"category"`
		Group       int    `json:"group"`
		Modes       []int  `json:"modes"`
		PeriodTypes []int  `json:"periodTypes"`
		StatID      string `json:"statId"`
		StatName    string `json:"statName"`
		UnitLabel   string `json:"unitLabel"`
		UnitType    int    `json:"unitType"`
		Weight      int    `json:"weight"`
	} `json:"averageScorePerLife"`
	BestSingleGameKills struct {
		Category        int    `json:"category"`
		Group           int    `json:"group"`
		Modes           []int  `json:"modes"`
		PeriodTypes     []int  `json:"periodTypes"`
		StatDescription string `json:"statDescription"`
		StatID          string `json:"statId"`
		StatName        string `json:"statName"`
		UnitLabel       string `json:"unitLabel"`
		UnitType        int    `json:"unitType"`
		Weight          int    `json:"weight"`
	} `json:"bestSingleGameKills"`
	BestSingleGameScore struct {
		Category    int    `json:"category"`
		Group       int    `json:"group"`
		Modes       []int  `json:"modes"`
		PeriodTypes []int  `json:"periodTypes"`
		StatID      string `json:"statId"`
		StatName    string `json:"statName"`
		UnitLabel   string `json:"unitLabel"`
		UnitType    int    `json:"unitType"`
		Weight      int    `json:"weight"`
	} `json:"bestSingleGameScore"`
	CarrierKills struct {
		Category    int           `json:"category"`
		Group       int           `json:"group"`
		Modes       []interface{} `json:"modes"`
		PeriodTypes []int         `json:"periodTypes"`
		StatID      string        `json:"statId"`
		StatName    string        `json:"statName"`
		UnitLabel   string        `json:"unitLabel"`
		UnitType    int           `json:"unitType"`
		Weight      int           `json:"weight"`
	} `json:"carrierKills"`
	CombatRating struct {
		Category        int    `json:"category"`
		Group           int    `json:"group"`
		Modes           []int  `json:"modes"`
		PeriodTypes     []int  `json:"periodTypes"`
		StatDescription string `json:"statDescription"`
		StatID          string `json:"statId"`
		StatName        string `json:"statName"`
		UnitLabel       string `json:"unitLabel"`
		UnitType        int    `json:"unitType"`
		Weight          int    `json:"weight"`
	} `json:"combatRating"`
	Completed struct {
		Category    int    `json:"category"`
		Group       int    `json:"group"`
		Modes       []int  `json:"modes"`
		PeriodTypes []int  `json:"periodTypes"`
		StatID      string `json:"statId"`
		StatName    string `json:"statName"`
		UnitLabel   string `json:"unitLabel"`
		UnitType    int    `json:"unitType"`
		Weight      int    `json:"weight"`
	} `json:"completed"`
	CompletionReason struct {
		Category    int    `json:"category"`
		Group       int    `json:"group"`
		Modes       []int  `json:"modes"`
		PeriodTypes []int  `json:"periodTypes"`
		StatID      string `json:"statId"`
		StatName    string `json:"statName"`
		UnitLabel   string `json:"unitLabel"`
		UnitType    int    `json:"unitType"`
		Weight      int    `json:"weight"`
	} `json:"completionReason"`
	DailyMedalsEarned struct {
		Category    int    `json:"category"`
		Group       int    `json:"group"`
		Modes       []int  `json:"modes"`
		PeriodTypes []int  `json:"periodTypes"`
		StatID      string `json:"statId"`
		StatName    string `json:"statName"`
		UnitLabel   string `json:"unitLabel"`
		UnitType    int    `json:"unitType"`
		Weight      int    `json:"weight"`
	} `json:"dailyMedalsEarned"`
	Deaths struct {
		Category        int    `json:"category"`
		Group           int    `json:"group"`
		Modes           []int  `json:"modes"`
		PeriodTypes     []int  `json:"periodTypes"`
		StatDescription string `json:"statDescription"`
		StatID          string `json:"statId"`
		StatName        string `json:"statName"`
		UnitLabel       string `json:"unitLabel"`
		UnitType        int    `json:"unitType"`
		Weight          int    `json:"weight"`
	} `json:"deaths"`
	DunkKills struct {
		Category    int           `json:"category"`
		Group       int           `json:"group"`
		Modes       []interface{} `json:"modes"`
		PeriodTypes []int         `json:"periodTypes"`
		StatID      string        `json:"statId"`
		StatName    string        `json:"statName"`
		UnitLabel   string        `json:"unitLabel"`
		UnitType    int           `json:"unitType"`
		Weight      int           `json:"weight"`
	} `json:"dunkKills"`
	FastestCompletionMs struct {
		Category    int    `json:"category"`
		Group       int    `json:"group"`
		Modes       []int  `json:"modes"`
		PeriodTypes []int  `json:"periodTypes"`
		StatID      string `json:"statId"`
		StatName    string `json:"statName"`
		UnitLabel   string `json:"unitLabel"`
		UnitType    int    `json:"unitType"`
		Weight      int    `json:"weight"`
	} `json:"fastestCompletionMs"`
	FastestCompletionMsForActivity struct {
		Category    int    `json:"category"`
		Group       int    `json:"group"`
		Modes       []int  `json:"modes"`
		PeriodTypes []int  `json:"periodTypes"`
		StatID      string `json:"statId"`
		StatName    string `json:"statName"`
		UnitLabel   string `json:"unitLabel"`
		UnitType    int    `json:"unitType"`
		Weight      int    `json:"weight"`
	} `json:"fastestCompletionMsForActivity"`
	FireteamID struct {
		Category    int    `json:"category"`
		Group       int    `json:"group"`
		Modes       []int  `json:"modes"`
		PeriodTypes []int  `json:"periodTypes"`
		StatID      string `json:"statId"`
		StatName    string `json:"statName"`
		UnitLabel   string `json:"unitLabel"`
		UnitType    int    `json:"unitType"`
		Weight      int    `json:"weight"`
	} `json:"fireteamId"`
	GatesHit struct {
		Category    int           `json:"category"`
		Group       int           `json:"group"`
		Modes       []interface{} `json:"modes"`
		PeriodTypes []int         `json:"periodTypes"`
		StatID      string        `json:"statId"`
		StatName    string        `json:"statName"`
		UnitLabel   string        `json:"unitLabel"`
		UnitType    int           `json:"unitType"`
		Weight      int           `json:"weight"`
	} `json:"gatesHit"`
	HighestCharacterLevel struct {
		Category    int    `json:"category"`
		Group       int    `json:"group"`
		Modes       []int  `json:"modes"`
		PeriodTypes []int  `json:"periodTypes"`
		StatID      string `json:"statId"`
		StatName    string `json:"statName"`
		UnitLabel   string `json:"unitLabel"`
		UnitType    int    `json:"unitType"`
		Weight      int    `json:"weight"`
	} `json:"highestCharacterLevel"`
	HighestLightLevel struct {
		Category    int    `json:"category"`
		Group       int    `json:"group"`
		Modes       []int  `json:"modes"`
		PeriodTypes []int  `json:"periodTypes"`
		StatID      string `json:"statId"`
		StatName    string `json:"statName"`
		UnitLabel   string `json:"unitLabel"`
		UnitType    int    `json:"unitType"`
		Weight      int    `json:"weight"`
	} `json:"highestLightLevel"`
	HighestSandboxLevel struct {
		Category    int    `json:"category"`
		Group       int    `json:"group"`
		Modes       []int  `json:"modes"`
		PeriodTypes []int  `json:"periodTypes"`
		StatID      string `json:"statId"`
		StatName    string `json:"statName"`
		UnitLabel   string `json:"unitLabel"`
		UnitType    int    `json:"unitType"`
		Weight      int    `json:"weight"`
	} `json:"highestSandboxLevel"`
	Kills struct {
		Category        int    `json:"category"`
		Group           int    `json:"group"`
		Modes           []int  `json:"modes"`
		PeriodTypes     []int  `json:"periodTypes"`
		StatDescription string `json:"statDescription"`
		StatID          string `json:"statId"`
		StatName        string `json:"statName"`
		UnitLabel       string `json:"unitLabel"`
		UnitType        int    `json:"unitType"`
		Weight          int    `json:"weight"`
	} `json:"kills"`
	KillsDeathsAssists struct {
		Category    int    `json:"category"`
		Group       int    `json:"group"`
		Modes       []int  `json:"modes"`
		PeriodTypes []int  `json:"periodTypes"`
		StatID      string `json:"statId"`
		StatName    string `json:"statName"`
		UnitLabel   string `json:"unitLabel"`
		UnitType    int    `json:"unitType"`
		Weight      int    `json:"weight"`
	} `json:"killsDeathsAssists"`
	KillsDeathsRatio struct {
		Category        int    `json:"category"`
		Group           int    `json:"group"`
		Modes           []int  `json:"modes"`
		PeriodTypes     []int  `json:"periodTypes"`
		StatDescription string `json:"statDescription"`
		StatID          string `json:"statId"`
		StatName        string `json:"statName"`
		UnitLabel       string `json:"unitLabel"`
		UnitType        int    `json:"unitType"`
		Weight          int    `json:"weight"`
	} `json:"killsDeathsRatio"`
	LbAssists struct {
		Category    int    `json:"category"`
		Group       int    `json:"group"`
		Modes       []int  `json:"modes"`
		PeriodTypes []int  `json:"periodTypes"`
		StatID      string `json:"statId"`
		StatName    string `json:"statName"`
		UnitLabel   string `json:"unitLabel"`
		UnitType    int    `json:"unitType"`
		Weight      int    `json:"weight"`
	} `json:"lbAssists"`
	LbDeaths struct {
		Category    int    `json:"category"`
		Group       int    `json:"group"`
		Modes       []int  `json:"modes"`
		PeriodTypes []int  `json:"periodTypes"`
		StatID      string `json:"statId"`
		StatName    string `json:"statName"`
		UnitLabel   string `json:"unitLabel"`
		UnitType    int    `json:"unitType"`
		Weight      int    `json:"weight"`
	} `json:"lbDeaths"`
	LbFastestCompletionMs struct {
		Category    int    `json:"category"`
		Group       int    `json:"group"`
		Modes       []int  `json:"modes"`
		PeriodTypes []int  `json:"periodTypes"`
		StatID      string `json:"statId"`
		StatName    string `json:"statName"`
		UnitLabel   string `json:"unitLabel"`
		UnitType    int    `json:"unitType"`
		Weight      int    `json:"weight"`
	} `json:"lbFastestCompletionMs"`
	LbKills struct {
		Category    int    `json:"category"`
		Group       int    `json:"group"`
		Modes       []int  `json:"modes"`
		PeriodTypes []int  `json:"periodTypes"`
		StatID      string `json:"statId"`
		StatName    string `json:"statName"`
		UnitLabel   string `json:"unitLabel"`
		UnitType    int    `json:"unitType"`
		Weight      int    `json:"weight"`
	} `json:"lbKills"`
	LbLongestKillDistance struct {
		Category    int    `json:"category"`
		Group       int    `json:"group"`
		Modes       []int  `json:"modes"`
		PeriodTypes []int  `json:"periodTypes"`
		StatID      string `json:"statId"`
		StatName    string `json:"statName"`
		UnitLabel   string `json:"unitLabel"`
		UnitType    int    `json:"unitType"`
		Weight      int    `json:"weight"`
	} `json:"lbLongestKillDistance"`
	LbLongestKillSpree struct {
		Category    int    `json:"category"`
		Group       int    `json:"group"`
		Modes       []int  `json:"modes"`
		PeriodTypes []int  `json:"periodTypes"`
		StatID      string `json:"statId"`
		StatName    string `json:"statName"`
		UnitLabel   string `json:"unitLabel"`
		UnitType    int    `json:"unitType"`
		Weight      int    `json:"weight"`
	} `json:"lbLongestKillSpree"`
	LbLongestSingleLife struct {
		Category    int    `json:"category"`
		Group       int    `json:"group"`
		Modes       []int  `json:"modes"`
		PeriodTypes []int  `json:"periodTypes"`
		StatID      string `json:"statId"`
		StatName    string `json:"statName"`
		UnitLabel   string `json:"unitLabel"`
		UnitType    int    `json:"unitType"`
		Weight      int    `json:"weight"`
	} `json:"lbLongestSingleLife"`
	LbMostPrecisionKills struct {
		Category    int    `json:"category"`
		Group       int    `json:"group"`
		Modes       []int  `json:"modes"`
		PeriodTypes []int  `json:"periodTypes"`
		StatID      string `json:"statId"`
		StatName    string `json:"statName"`
		UnitLabel   string `json:"unitLabel"`
		UnitType    int    `json:"unitType"`
		Weight      int    `json:"weight"`
	} `json:"lbMostPrecisionKills"`
	LbObjectivesCompleted struct {
		Category    int    `json:"category"`
		Group       int    `json:"group"`
		Modes       []int  `json:"modes"`
		PeriodTypes []int  `json:"periodTypes"`
		StatID      string `json:"statId"`
		StatName    string `json:"statName"`
		UnitLabel   string `json:"unitLabel"`
		UnitType    int    `json:"unitType"`
		Weight      int    `json:"weight"`
	} `json:"lbObjectivesCompleted"`
	LbPrecisionKills struct {
		Category    int    `json:"category"`
		Group       int    `json:"group"`
		Modes       []int  `json:"modes"`
		PeriodTypes []int  `json:"periodTypes"`
		StatID      string `json:"statId"`
		StatName    string `json:"statName"`
		UnitLabel   string `json:"unitLabel"`
		UnitType    int    `json:"unitType"`
		Weight      int    `json:"weight"`
	} `json:"lbPrecisionKills"`
	LbSingleGameKills struct {
		Category    int    `json:"category"`
		Group       int    `json:"group"`
		Modes       []int  `json:"modes"`
		PeriodTypes []int  `json:"periodTypes"`
		StatID      string `json:"statId"`
		StatName    string `json:"statName"`
		UnitLabel   string `json:"unitLabel"`
		UnitType    int    `json:"unitType"`
		Weight      int    `json:"weight"`
	} `json:"lbSingleGameKills"`
	LbSingleGameScore struct {
		Category    int    `json:"category"`
		Group       int    `json:"group"`
		Modes       []int  `json:"modes"`
		PeriodTypes []int  `json:"periodTypes"`
		StatID      string `json:"statId"`
		StatName    string `json:"statName"`
		UnitLabel   string `json:"unitLabel"`
		UnitType    int    `json:"unitType"`
		Weight      int    `json:"weight"`
	} `json:"lbSingleGameScore"`
	LongestKillDistance struct {
		Category    int    `json:"category"`
		Group       int    `json:"group"`
		Modes       []int  `json:"modes"`
		PeriodTypes []int  `json:"periodTypes"`
		StatID      string `json:"statId"`
		StatName    string `json:"statName"`
		UnitLabel   string `json:"unitLabel"`
		UnitType    int    `json:"unitType"`
		Weight      int    `json:"weight"`
	} `json:"longestKillDistance"`
	LongestKillSpree struct {
		Category    int    `json:"category"`
		Group       int    `json:"group"`
		Modes       []int  `json:"modes"`
		PeriodTypes []int  `json:"periodTypes"`
		StatID      string `json:"statId"`
		StatName    string `json:"statName"`
		UnitLabel   string `json:"unitLabel"`
		UnitType    int    `json:"unitType"`
		Weight      int    `json:"weight"`
	} `json:"longestKillSpree"`
	LongestSingleLife struct {
		Category    int    `json:"category"`
		Group       int    `json:"group"`
		Modes       []int  `json:"modes"`
		PeriodTypes []int  `json:"periodTypes"`
		StatID      string `json:"statId"`
		StatName    string `json:"statName"`
		UnitLabel   string `json:"unitLabel"`
		UnitType    int    `json:"unitType"`
		Weight      int    `json:"weight"`
	} `json:"longestSingleLife"`
	MaximumPowerLevel struct {
		Category    int    `json:"category"`
		Group       int    `json:"group"`
		Modes       []int  `json:"modes"`
		PeriodTypes []int  `json:"periodTypes"`
		StatID      string `json:"statId"`
		StatName    string `json:"statName"`
		UnitLabel   string `json:"unitLabel"`
		UnitType    int    `json:"unitType"`
		Weight      int    `json:"weight"`
	} `json:"maximumPowerLevel"`
	MedalAbilityDawnbladeAerial struct {
		Category            int    `json:"category"`
		Group               int    `json:"group"`
		MedalTierHash       int    `json:"medalTierHash"`
		MedalTierIdentifier string `json:"medalTierIdentifier"`
		Modes               []int  `json:"modes"`
		PeriodTypes         []int  `json:"periodTypes"`
		StatDescription     string `json:"statDescription"`
		StatID              string `json:"statId"`
		StatName            string `json:"statName"`
		UnitLabel           string `json:"unitLabel"`
		UnitType            int    `json:"unitType"`
		Weight              int    `json:"weight"`
	} `json:"medalAbilityDawnbladeAerial"`
	MedalAbilityDawnbladeSlam struct {
		Category            int    `json:"category"`
		Group               int    `json:"group"`
		MedalTierHash       int    `json:"medalTierHash"`
		MedalTierIdentifier string `json:"medalTierIdentifier"`
		Modes               []int  `json:"modes"`
		PeriodTypes         []int  `json:"periodTypes"`
		StatDescription     string `json:"statDescription"`
		StatID              string `json:"statId"`
		StatName            string `json:"statName"`
		UnitLabel           string `json:"unitLabel"`
		UnitType            int    `json:"unitType"`
		Weight              int    `json:"weight"`
	} `json:"medalAbilityDawnbladeSlam"`
	MedalAbilityFlowwalkerMulti struct {
		Category            int    `json:"category"`
		Group               int    `json:"group"`
		MedalTierHash       int    `json:"medalTierHash"`
		MedalTierIdentifier string `json:"medalTierIdentifier"`
		Modes               []int  `json:"modes"`
		PeriodTypes         []int  `json:"periodTypes"`
		StatDescription     string `json:"statDescription"`
		StatID              string `json:"statId"`
		StatName            string `json:"statName"`
		UnitLabel           string `json:"unitLabel"`
		UnitType            int    `json:"unitType"`
		Weight              int    `json:"weight"`
	} `json:"medalAbilityFlowwalkerMulti"`
	MedalAbilityFlowwalkerQuick struct {
		Category            int    `json:"category"`
		Group               int    `json:"group"`
		MedalTierHash       int    `json:"medalTierHash"`
		MedalTierIdentifier string `json:"medalTierIdentifier"`
		Modes               []int  `json:"modes"`
		PeriodTypes         []int  `json:"periodTypes"`
		StatDescription     string `json:"statDescription"`
		StatID              string `json:"statId"`
		StatName            string `json:"statName"`
		UnitLabel           string `json:"unitLabel"`
		UnitType            int    `json:"unitType"`
		Weight              int    `json:"weight"`
	} `json:"medalAbilityFlowwalkerQuick"`
	MedalAbilityGunslingerMulti struct {
		Category            int    `json:"category"`
		Group               int    `json:"group"`
		MedalTierHash       int    `json:"medalTierHash"`
		MedalTierIdentifier string `json:"medalTierIdentifier"`
		Modes               []int  `json:"modes"`
		PeriodTypes         []int  `json:"periodTypes"`
		StatDescription     string `json:"statDescription"`
		StatID              string `json:"statId"`
		StatName            string `json:"statName"`
		UnitLabel           string `json:"unitLabel"`
		UnitType            int    `json:"unitType"`
		Weight              int    `json:"weight"`
	} `json:"medalAbilityGunslingerMulti"`
	MedalAbilityGunslingerQuick struct {
		Category            int    `json:"category"`
		Group               int    `json:"group"`
		MedalTierHash       int    `json:"medalTierHash"`
		MedalTierIdentifier string `json:"medalTierIdentifier"`
		Modes               []int  `json:"modes"`
		PeriodTypes         []int  `json:"periodTypes"`
		StatDescription     string `json:"statDescription"`
		StatID              string `json:"statId"`
		StatName            string `json:"statName"`
		UnitLabel           string `json:"unitLabel"`
		UnitType            int    `json:"unitType"`
		Weight              int    `json:"weight"`
	} `json:"medalAbilityGunslingerQuick"`
	MedalAbilityJuggernautCombo struct {
		Category            int    `json:"category"`
		Group               int    `json:"group"`
		MedalTierHash       int    `json:"medalTierHash"`
		MedalTierIdentifier string `json:"medalTierIdentifier"`
		Modes               []int  `json:"modes"`
		PeriodTypes         []int  `json:"periodTypes"`
		StatDescription     string `json:"statDescription"`
		StatID              string `json:"statId"`
		StatName            string `json:"statName"`
		UnitLabel           string `json:"unitLabel"`
		UnitType            int    `json:"unitType"`
		Weight              int    `json:"weight"`
	} `json:"medalAbilityJuggernautCombo"`
	MedalAbilityJuggernautSlam struct {
		Category            int    `json:"category"`
		Group               int    `json:"group"`
		MedalTierHash       int    `json:"medalTierHash"`
		MedalTierIdentifier string `json:"medalTierIdentifier"`
		Modes               []int  `json:"modes"`
		PeriodTypes         []int  `json:"periodTypes"`
		StatDescription     string `json:"statDescription"`
		StatID              string `json:"statId"`
		StatName            string `json:"statName"`
		UnitLabel           string `json:"unitLabel"`
		UnitType            int    `json:"unitType"`
		Weight              int    `json:"weight"`
	} `json:"medalAbilityJuggernautSlam"`
	MedalAbilityNightstalkerLongRange struct {
		Category            int    `json:"category"`
		Group               int    `json:"group"`
		MedalTierHash       int    `json:"medalTierHash"`
		MedalTierIdentifier string `json:"medalTierIdentifier"`
		Modes               []int  `json:"modes"`
		PeriodTypes         []int  `json:"periodTypes"`
		StatDescription     string `json:"statDescription"`
		StatID              string `json:"statId"`
		StatName            string `json:"statName"`
		UnitLabel           string `json:"unitLabel"`
		UnitType            int    `json:"unitType"`
		Weight              int    `json:"weight"`
	} `json:"medalAbilityNightstalkerLongRange"`
	MedalAbilityNightstalkerTetherQuick struct {
		Category            int    `json:"category"`
		Group               int    `json:"group"`
		MedalTierHash       int    `json:"medalTierHash"`
		MedalTierIdentifier string `json:"medalTierIdentifier"`
		Modes               []int  `json:"modes"`
		PeriodTypes         []int  `json:"periodTypes"`
		StatDescription     string `json:"statDescription"`
		StatID              string `json:"statId"`
		StatName            string `json:"statName"`
		UnitLabel           string `json:"unitLabel"`
		UnitType            int    `json:"unitType"`
		Weight              int    `json:"weight"`
	} `json:"medalAbilityNightstalkerTetherQuick"`
	MedalAbilitySentinelCombo struct {
		Category            int    `json:"category"`
		Group               int    `json:"group"`
		MedalTierHash       int    `json:"medalTierHash"`
		MedalTierIdentifier string `json:"medalTierIdentifier"`
		Modes               []int  `json:"modes"`
		PeriodTypes         []int  `json:"periodTypes"`
		StatDescription     string `json:"statDescription"`
		StatID              string `json:"statId"`
		StatName            string `json:"statName"`
		UnitLabel           string `json:"unitLabel"`
		UnitType            int    `json:"unitType"`
		Weight              int    `json:"weight"`
	} `json:"medalAbilitySentinelCombo"`
	MedalAbilitySentinelWard struct {
		Category            int    `json:"category"`
		Group               int    `json:"group"`
		MedalTierHash       int    `json:"medalTierHash"`
		MedalTierIdentifier string `json:"medalTierIdentifier"`
		Modes               []int  `json:"modes"`
		PeriodTypes         []int  `json:"periodTypes"`
		StatDescription     string `json:"statDescription"`
		StatID              string `json:"statId"`
		StatName            string `json:"statName"`
		UnitLabel           string `json:"unitLabel"`
		UnitType            int    `json:"unitType"`
		Weight              int    `json:"weight"`
	} `json:"medalAbilitySentinelWard"`
	MedalAbilityStormcallerLandfall struct {
		Category            int    `json:"category"`
		Group               int    `json:"group"`
		MedalTierHash       int    `json:"medalTierHash"`
		MedalTierIdentifier string `json:"medalTierIdentifier"`
		Modes               []int  `json:"modes"`
		PeriodTypes         []int  `json:"periodTypes"`
		StatDescription     string `json:"statDescription"`
		StatID              string `json:"statId"`
		StatName            string `json:"statName"`
		UnitLabel           string `json:"unitLabel"`
		UnitType            int    `json:"unitType"`
		Weight              int    `json:"weight"`
	} `json:"medalAbilityStormcallerLandfall"`
	MedalAbilityStormcallerMulti struct {
		Category            int    `json:"category"`
		Group               int    `json:"group"`
		MedalTierHash       int    `json:"medalTierHash"`
		MedalTierIdentifier string `json:"medalTierIdentifier"`
		Modes               []int  `json:"modes"`
		PeriodTypes         []int  `json:"periodTypes"`
		StatDescription     string `json:"statDescription"`
		StatID              string `json:"statId"`
		StatName            string `json:"statName"`
		UnitLabel           string `json:"unitLabel"`
		UnitType            int    `json:"unitType"`
		Weight              int    `json:"weight"`
	} `json:"medalAbilityStormcallerMulti"`
	MedalAbilitySunbreakerLongRange struct {
		Category            int    `json:"category"`
		Group               int    `json:"group"`
		MedalTierHash       int    `json:"medalTierHash"`
		MedalTierIdentifier string `json:"medalTierIdentifier"`
		Modes               []int  `json:"modes"`
		PeriodTypes         []int  `json:"periodTypes"`
		StatDescription     string `json:"statDescription"`
		StatID              string `json:"statId"`
		StatName            string `json:"statName"`
		UnitLabel           string `json:"unitLabel"`
		UnitType            int    `json:"unitType"`
		Weight              int    `json:"weight"`
	} `json:"medalAbilitySunbreakerLongRange"`
	MedalAbilitySunbreakerMulti struct {
		Category            int    `json:"category"`
		Group               int    `json:"group"`
		MedalTierHash       int    `json:"medalTierHash"`
		MedalTierIdentifier string `json:"medalTierIdentifier"`
		Modes               []int  `json:"modes"`
		PeriodTypes         []int  `json:"periodTypes"`
		StatDescription     string `json:"statDescription"`
		StatID              string `json:"statId"`
		StatName            string `json:"statName"`
		UnitLabel           string `json:"unitLabel"`
		UnitType            int    `json:"unitType"`
		Weight              int    `json:"weight"`
	} `json:"medalAbilitySunbreakerMulti"`
	MedalAbilityVoidwalkerDistance struct {
		Category            int    `json:"category"`
		Group               int    `json:"group"`
		MedalTierHash       int    `json:"medalTierHash"`
		MedalTierIdentifier string `json:"medalTierIdentifier"`
		Modes               []int  `json:"modes"`
		PeriodTypes         []int  `json:"periodTypes"`
		StatDescription     string `json:"statDescription"`
		StatID              string `json:"statId"`
		StatName            string `json:"statName"`
		UnitLabel           string `json:"unitLabel"`
		UnitType            int    `json:"unitType"`
		Weight              int    `json:"weight"`
	} `json:"medalAbilityVoidwalkerDistance"`
	MedalAbilityVoidwalkerVortex struct {
		Category            int    `json:"category"`
		Group               int    `json:"group"`
		MedalTierHash       int    `json:"medalTierHash"`
		MedalTierIdentifier string `json:"medalTierIdentifier"`
		Modes               []int  `json:"modes"`
		PeriodTypes         []int  `json:"periodTypes"`
		StatDescription     string `json:"statDescription"`
		StatID              string `json:"statId"`
		StatName            string `json:"statName"`
		UnitLabel           string `json:"unitLabel"`
		UnitType            int    `json:"unitType"`
		Weight              int    `json:"weight"`
	} `json:"medalAbilityVoidwalkerVortex"`
	MedalAvenger struct {
		Category            int    `json:"category"`
		Group               int    `json:"group"`
		MedalTierHash       int    `json:"medalTierHash"`
		MedalTierIdentifier string `json:"medalTierIdentifier"`
		Modes               []int  `json:"modes"`
		PeriodTypes         []int  `json:"periodTypes"`
		StatDescription     string `json:"statDescription"`
		StatID              string `json:"statId"`
		StatName            string `json:"statName"`
		UnitLabel           string `json:"unitLabel"`
		UnitType            int    `json:"unitType"`
		Weight              int    `json:"weight"`
	} `json:"medalAvenger"`
	MedalControlAdvantageHold struct {
		Category            int    `json:"category"`
		Group               int    `json:"group"`
		MedalTierHash       int    `json:"medalTierHash"`
		MedalTierIdentifier string `json:"medalTierIdentifier"`
		Modes               []int  `json:"modes"`
		PeriodTypes         []int  `json:"periodTypes"`
		StatDescription     string `json:"statDescription"`
		StatID              string `json:"statId"`
		StatName            string `json:"statName"`
		UnitLabel           string `json:"unitLabel"`
		UnitType            int    `json:"unitType"`
		Weight              int    `json:"weight"`
	} `json:"medalControlAdvantageHold"`
	MedalControlAdvantageStreak struct {
		Category            int    `json:"category"`
		Group               int    `json:"group"`
		MedalTierHash       int    `json:"medalTierHash"`
		MedalTierIdentifier string `json:"medalTierIdentifier"`
		Modes               []int  `json:"modes"`
		PeriodTypes         []int  `json:"periodTypes"`
		StatDescription     string `json:"statDescription"`
		StatID              string `json:"statId"`
		StatName            string `json:"statName"`
		UnitLabel           string `json:"unitLabel"`
		UnitType            int    `json:"unitType"`
		Weight              int    `json:"weight"`
	} `json:"medalControlAdvantageStreak"`
	MedalControlCaptureAllZones struct {
		Category            int    `json:"category"`
		Group               int    `json:"group"`
		MedalTierHash       int    `json:"medalTierHash"`
		MedalTierIdentifier string `json:"medalTierIdentifier"`
		Modes               []int  `json:"modes"`
		PeriodTypes         []int  `json:"periodTypes"`
		StatDescription     string `json:"statDescription"`
		StatID              string `json:"statId"`
		StatName            string `json:"statName"`
		UnitLabel           string `json:"unitLabel"`
		UnitType            int    `json:"unitType"`
		Weight              int    `json:"weight"`
	} `json:"medalControlCaptureAllZones"`
	MedalControlMostAdvantage struct {
		Category            int    `json:"category"`
		Group               int    `json:"group"`
		MedalTierHash       int    `json:"medalTierHash"`
		MedalTierIdentifier string `json:"medalTierIdentifier"`
		Modes               []int  `json:"modes"`
		PeriodTypes         []int  `json:"periodTypes"`
		StatDescription     string `json:"statDescription"`
		StatID              string `json:"statId"`
		StatName            string `json:"statName"`
		UnitLabel           string `json:"unitLabel"`
		UnitType            int    `json:"unitType"`
		Weight              int    `json:"weight"`
	} `json:"medalControlMostAdvantage"`
	MedalControlPerimeterKill struct {
		Category            int    `json:"category"`
		Group               int    `json:"group"`
		MedalTierHash       int    `json:"medalTierHash"`
		MedalTierIdentifier string `json:"medalTierIdentifier"`
		Modes               []int  `json:"modes"`
		PeriodTypes         []int  `json:"periodTypes"`
		StatDescription     string `json:"statDescription"`
		StatID              string `json:"statId"`
		StatName            string `json:"statName"`
		UnitLabel           string `json:"unitLabel"`
		UnitType            int    `json:"unitType"`
		Weight              int    `json:"weight"`
	} `json:"medalControlPerimeterKill"`
	MedalControlPowerPlayWipe struct {
		Category            int    `json:"category"`
		Group               int    `json:"group"`
		MedalTierHash       int    `json:"medalTierHash"`
		MedalTierIdentifier string `json:"medalTierIdentifier"`
		Modes               []int  `json:"modes"`
		PeriodTypes         []int  `json:"periodTypes"`
		StatDescription     string `json:"statDescription"`
		StatID              string `json:"statId"`
		StatName            string `json:"statName"`
		UnitLabel           string `json:"unitLabel"`
		UnitType            int    `json:"unitType"`
		Weight              int    `json:"weight"`
	} `json:"medalControlPowerPlayWipe"`
	MedalCountdownDefense struct {
		Category            int    `json:"category"`
		Group               int    `json:"group"`
		MedalTierHash       int    `json:"medalTierHash"`
		MedalTierIdentifier string `json:"medalTierIdentifier"`
		Modes               []int  `json:"modes"`
		PeriodTypes         []int  `json:"periodTypes"`
		StatDescription     string `json:"statDescription"`
		StatID              string `json:"statId"`
		StatName            string `json:"statName"`
		UnitLabel           string `json:"unitLabel"`
		UnitType            int    `json:"unitType"`
		Weight              int    `json:"weight"`
	} `json:"medalCountdownDefense"`
	MedalCountdownDefusedLastStand struct {
		Category            int    `json:"category"`
		Group               int    `json:"group"`
		MedalTierHash       int    `json:"medalTierHash"`
		MedalTierIdentifier string `json:"medalTierIdentifier"`
		Modes               []int  `json:"modes"`
		PeriodTypes         []int  `json:"periodTypes"`
		StatDescription     string `json:"statDescription"`
		StatID              string `json:"statId"`
		StatName            string `json:"statName"`
		UnitLabel           string `json:"unitLabel"`
		UnitType            int    `json:"unitType"`
		Weight              int    `json:"weight"`
	} `json:"medalCountdownDefusedLastStand"`
	MedalCountdownDefusedMulti struct {
		Category            int    `json:"category"`
		Group               int    `json:"group"`
		MedalTierHash       int    `json:"medalTierHash"`
		MedalTierIdentifier string `json:"medalTierIdentifier"`
		Modes               []int  `json:"modes"`
		PeriodTypes         []int  `json:"periodTypes"`
		StatDescription     string `json:"statDescription"`
		StatID              string `json:"statId"`
		StatName            string `json:"statName"`
		UnitLabel           string `json:"unitLabel"`
		UnitType            int    `json:"unitType"`
		Weight              int    `json:"weight"`
	} `json:"medalCountdownDefusedMulti"`
	MedalCountdownDetonated struct {
		Category            int    `json:"category"`
		Group               int    `json:"group"`
		MedalTierHash       int    `json:"medalTierHash"`
		MedalTierIdentifier string `json:"medalTierIdentifier"`
		Modes               []int  `json:"modes"`
		PeriodTypes         []int  `json:"periodTypes"`
		StatDescription     string `json:"statDescription"`
		StatID              string `json:"statId"`
		StatName            string `json:"statName"`
		UnitLabel           string `json:"unitLabel"`
		UnitType            int    `json:"unitType"`
		Weight              int    `json:"weight"`
	} `json:"medalCountdownDetonated"`
	MedalCountdownPerfect struct {
		Category            int    `json:"category"`
		Group               int    `json:"group"`
		MedalTierHash       int    `json:"medalTierHash"`
		MedalTierIdentifier string `json:"medalTierIdentifier"`
		Modes               []int  `json:"modes"`
		PeriodTypes         []int  `json:"periodTypes"`
		StatDescription     string `json:"statDescription"`
		StatID              string `json:"statId"`
		StatName            string `json:"statName"`
		UnitLabel           string `json:"unitLabel"`
		UnitType            int    `json:"unitType"`
		Weight              int    `json:"weight"`
	} `json:"medalCountdownPerfect"`
	MedalCountdownRoundAllAlive struct {
		Category            int    `json:"category"`
		Group               int    `json:"group"`
		MedalTierHash       int    `json:"medalTierHash"`
		MedalTierIdentifier string `json:"medalTierIdentifier"`
		Modes               []int  `json:"modes"`
		PeriodTypes         []int  `json:"periodTypes"`
		StatDescription     string `json:"statDescription"`
		StatID              string `json:"statId"`
		StatName            string `json:"statName"`
		UnitLabel           string `json:"unitLabel"`
		UnitType            int    `json:"unitType"`
		Weight              int    `json:"weight"`
	} `json:"medalCountdownRoundAllAlive"`
	MedalCycle struct {
		Category            int    `json:"category"`
		Group               int    `json:"group"`
		MedalTierHash       int    `json:"medalTierHash"`
		MedalTierIdentifier string `json:"medalTierIdentifier"`
		Modes               []int  `json:"modes"`
		PeriodTypes         []int  `json:"periodTypes"`
		StatDescription     string `json:"statDescription"`
		StatID              string `json:"statId"`
		StatName            string `json:"statName"`
		UnitLabel           string `json:"unitLabel"`
		UnitType            int    `json:"unitType"`
		Weight              int    `json:"weight"`
	} `json:"medalCycle"`
	MedalDefeatHunterDodge struct {
		Category            int    `json:"category"`
		Group               int    `json:"group"`
		MedalTierHash       int    `json:"medalTierHash"`
		MedalTierIdentifier string `json:"medalTierIdentifier"`
		Modes               []int  `json:"modes"`
		PeriodTypes         []int  `json:"periodTypes"`
		StatDescription     string `json:"statDescription"`
		StatID              string `json:"statId"`
		StatName            string `json:"statName"`
		UnitLabel           string `json:"unitLabel"`
		UnitType            int    `json:"unitType"`
		Weight              int    `json:"weight"`
	} `json:"medalDefeatHunterDodge"`
	MedalDefeatTitanBrace struct {
		Category            int    `json:"category"`
		Group               int    `json:"group"`
		MedalTierHash       int    `json:"medalTierHash"`
		MedalTierIdentifier string `json:"medalTierIdentifier"`
		Modes               []int  `json:"modes"`
		PeriodTypes         []int  `json:"periodTypes"`
		StatDescription     string `json:"statDescription"`
		StatID              string `json:"statId"`
		StatName            string `json:"statName"`
		UnitLabel           string `json:"unitLabel"`
		UnitType            int    `json:"unitType"`
		Weight              int    `json:"weight"`
	} `json:"medalDefeatTitanBrace"`
	MedalDefeatWarlockSigil struct {
		Category            int    `json:"category"`
		Group               int    `json:"group"`
		MedalTierHash       int    `json:"medalTierHash"`
		MedalTierIdentifier string `json:"medalTierIdentifier"`
		Modes               []int  `json:"modes"`
		PeriodTypes         []int  `json:"periodTypes"`
		StatDescription     string `json:"statDescription"`
		StatID              string `json:"statId"`
		StatName            string `json:"statName"`
		UnitLabel           string `json:"unitLabel"`
		UnitType            int    `json:"unitType"`
		Weight              int    `json:"weight"`
	} `json:"medalDefeatWarlockSigil"`
	MedalDefense struct {
		Category            int    `json:"category"`
		Group               int    `json:"group"`
		MedalTierHash       int    `json:"medalTierHash"`
		MedalTierIdentifier string `json:"medalTierIdentifier"`
		Modes               []int  `json:"modes"`
		PeriodTypes         []int  `json:"periodTypes"`
		StatDescription     string `json:"statDescription"`
		StatID              string `json:"statId"`
		StatName            string `json:"statName"`
		UnitLabel           string `json:"unitLabel"`
		UnitType            int    `json:"unitType"`
		Weight              int    `json:"weight"`
	} `json:"medalDefense"`
	MedalMatchBlowout struct {
		Category            int    `json:"category"`
		Group               int    `json:"group"`
		MedalTierHash       int    `json:"medalTierHash"`
		MedalTierIdentifier string `json:"medalTierIdentifier"`
		Modes               []int  `json:"modes"`
		PeriodTypes         []int  `json:"periodTypes"`
		StatDescription     string `json:"statDescription"`
		StatID              string `json:"statId"`
		StatName            string `json:"statName"`
		UnitLabel           string `json:"unitLabel"`
		UnitType            int    `json:"unitType"`
		Weight              int    `json:"weight"`
	} `json:"medalMatchBlowout"`
	MedalMatchComeback struct {
		Category            int    `json:"category"`
		Group               int    `json:"group"`
		MedalTierHash       int    `json:"medalTierHash"`
		MedalTierIdentifier string `json:"medalTierIdentifier"`
		Modes               []int  `json:"modes"`
		PeriodTypes         []int  `json:"periodTypes"`
		StatDescription     string `json:"statDescription"`
		StatID              string `json:"statId"`
		StatName            string `json:"statName"`
		UnitLabel           string `json:"unitLabel"`
		UnitType            int    `json:"unitType"`
		Weight              int    `json:"weight"`
	} `json:"medalMatchComeback"`
	MedalMatchMostDamage struct {
		Category            int    `json:"category"`
		Group               int    `json:"group"`
		MedalTierHash       int    `json:"medalTierHash"`
		MedalTierIdentifier string `json:"medalTierIdentifier"`
		Modes               []int  `json:"modes"`
		PeriodTypes         []int  `json:"periodTypes"`
		StatDescription     string `json:"statDescription"`
		StatID              string `json:"statId"`
		StatName            string `json:"statName"`
		UnitLabel           string `json:"unitLabel"`
		UnitType            int    `json:"unitType"`
		Weight              int    `json:"weight"`
	} `json:"medalMatchMostDamage"`
	MedalMatchNeverTrailed struct {
		Category            int    `json:"category"`
		Group               int    `json:"group"`
		MedalTierHash       int    `json:"medalTierHash"`
		MedalTierIdentifier string `json:"medalTierIdentifier"`
		Modes               []int  `json:"modes"`
		PeriodTypes         []int  `json:"periodTypes"`
		StatDescription     string `json:"statDescription"`
		StatID              string `json:"statId"`
		StatName            string `json:"statName"`
		UnitLabel           string `json:"unitLabel"`
		UnitType            int    `json:"unitType"`
		Weight              int    `json:"weight"`
	} `json:"medalMatchNeverTrailed"`
	MedalMatchOvertime struct {
		Category            int    `json:"category"`
		Group               int    `json:"group"`
		MedalTierHash       int    `json:"medalTierHash"`
		MedalTierIdentifier string `json:"medalTierIdentifier"`
		Modes               []int  `json:"modes"`
		PeriodTypes         []int  `json:"periodTypes"`
		StatDescription     string `json:"statDescription"`
		StatID              string `json:"statId"`
		StatName            string `json:"statName"`
		UnitLabel           string `json:"unitLabel"`
		UnitType            int    `json:"unitType"`
		Weight              int    `json:"weight"`
	} `json:"medalMatchOvertime"`
	MedalMatchUndefeated struct {
		Category            int    `json:"category"`
		Group               int    `json:"group"`
		MedalTierHash       int    `json:"medalTierHash"`
		MedalTierIdentifier string `json:"medalTierIdentifier"`
		Modes               []int  `json:"modes"`
		PeriodTypes         []int  `json:"periodTypes"`
		StatDescription     string `json:"statDescription"`
		StatID              string `json:"statId"`
		StatName            string `json:"statName"`
		UnitLabel           string `json:"unitLabel"`
		UnitType            int    `json:"unitType"`
		Weight              int    `json:"weight"`
	} `json:"medalMatchUndefeated"`
	MedalMulti2x struct {
		Category            int    `json:"category"`
		Group               int    `json:"group"`
		MedalTierHash       int    `json:"medalTierHash"`
		MedalTierIdentifier string `json:"medalTierIdentifier"`
		Modes               []int  `json:"modes"`
		PeriodTypes         []int  `json:"periodTypes"`
		StatDescription     string `json:"statDescription"`
		StatID              string `json:"statId"`
		StatName            string `json:"statName"`
		UnitLabel           string `json:"unitLabel"`
		UnitType            int    `json:"unitType"`
		Weight              int    `json:"weight"`
	} `json:"medalMulti2x"`
	MedalMulti3x struct {
		Category            int    `json:"category"`
		Group               int    `json:"group"`
		MedalTierHash       int    `json:"medalTierHash"`
		MedalTierIdentifier string `json:"medalTierIdentifier"`
		Modes               []int  `json:"modes"`
		PeriodTypes         []int  `json:"periodTypes"`
		StatDescription     string `json:"statDescription"`
		StatID              string `json:"statId"`
		StatName            string `json:"statName"`
		UnitLabel           string `json:"unitLabel"`
		UnitType            int    `json:"unitType"`
		Weight              int    `json:"weight"`
	} `json:"medalMulti3x"`
	MedalMulti4x struct {
		Category            int    `json:"category"`
		Group               int    `json:"group"`
		MedalTierHash       int    `json:"medalTierHash"`
		MedalTierIdentifier string `json:"medalTierIdentifier"`
		Modes               []int  `json:"modes"`
		PeriodTypes         []int  `json:"periodTypes"`
		StatDescription     string `json:"statDescription"`
		StatID              string `json:"statId"`
		StatName            string `json:"statName"`
		UnitLabel           string `json:"unitLabel"`
		UnitType            int    `json:"unitType"`
		Weight              int    `json:"weight"`
	} `json:"medalMulti4x"`
	MedalMultiEntireTeam struct {
		Category            int    `json:"category"`
		Group               int    `json:"group"`
		MedalTierHash       int    `json:"medalTierHash"`
		MedalTierIdentifier string `json:"medalTierIdentifier"`
		Modes               []int  `json:"modes"`
		PeriodTypes         []int  `json:"periodTypes"`
		StatDescription     string `json:"statDescription"`
		StatID              string `json:"statId"`
		StatName            string `json:"statName"`
		UnitLabel           string `json:"unitLabel"`
		UnitType            int    `json:"unitType"`
		Weight              int    `json:"weight"`
	} `json:"medalMultiEntireTeam"`
	MedalPayback struct {
		Category            int    `json:"category"`
		Group               int    `json:"group"`
		MedalTierHash       int    `json:"medalTierHash"`
		MedalTierIdentifier string `json:"medalTierIdentifier"`
		Modes               []int  `json:"modes"`
		PeriodTypes         []int  `json:"periodTypes"`
		StatDescription     string `json:"statDescription"`
		StatID              string `json:"statId"`
		StatName            string `json:"statName"`
		UnitLabel           string `json:"unitLabel"`
		UnitType            int    `json:"unitType"`
		Weight              int    `json:"weight"`
	} `json:"medalPayback"`
	MedalQuickStrike struct {
		Category            int    `json:"category"`
		Group               int    `json:"group"`
		MedalTierHash       int    `json:"medalTierHash"`
		MedalTierIdentifier string `json:"medalTierIdentifier"`
		Modes               []int  `json:"modes"`
		PeriodTypes         []int  `json:"periodTypes"`
		StatDescription     string `json:"statDescription"`
		StatID              string `json:"statId"`
		StatName            string `json:"statName"`
		UnitLabel           string `json:"unitLabel"`
		UnitType            int    `json:"unitType"`
		Weight              int    `json:"weight"`
	} `json:"medalQuickStrike"`
	MedalStreak10x struct {
		Category            int    `json:"category"`
		Group               int    `json:"group"`
		MedalTierHash       int    `json:"medalTierHash"`
		MedalTierIdentifier string `json:"medalTierIdentifier"`
		Modes               []int  `json:"modes"`
		PeriodTypes         []int  `json:"periodTypes"`
		StatDescription     string `json:"statDescription"`
		StatID              string `json:"statId"`
		StatName            string `json:"statName"`
		UnitLabel           string `json:"unitLabel"`
		UnitType            int    `json:"unitType"`
		Weight              int    `json:"weight"`
	} `json:"medalStreak10x"`
	MedalStreak5x struct {
		Category            int    `json:"category"`
		Group               int    `json:"group"`
		MedalTierHash       int    `json:"medalTierHash"`
		MedalTierIdentifier string `json:"medalTierIdentifier"`
		Modes               []int  `json:"modes"`
		PeriodTypes         []int  `json:"periodTypes"`
		StatDescription     string `json:"statDescription"`
		StatID              string `json:"statId"`
		StatName            string `json:"statName"`
		UnitLabel           string `json:"unitLabel"`
		UnitType            int    `json:"unitType"`
		Weight              int    `json:"weight"`
	} `json:"medalStreak5x"`
	MedalStreakAbsurd struct {
		Category            int    `json:"category"`
		Group               int    `json:"group"`
		MedalTierHash       int    `json:"medalTierHash"`
		MedalTierIdentifier string `json:"medalTierIdentifier"`
		Modes               []int  `json:"modes"`
		PeriodTypes         []int  `json:"periodTypes"`
		StatDescription     string `json:"statDescription"`
		StatID              string `json:"statId"`
		StatName            string `json:"statName"`
		UnitLabel           string `json:"unitLabel"`
		UnitType            int    `json:"unitType"`
		Weight              int    `json:"weight"`
	} `json:"medalStreakAbsurd"`
	MedalStreakCombined struct {
		Category            int    `json:"category"`
		Group               int    `json:"group"`
		MedalTierHash       int    `json:"medalTierHash"`
		MedalTierIdentifier string `json:"medalTierIdentifier"`
		Modes               []int  `json:"modes"`
		PeriodTypes         []int  `json:"periodTypes"`
		StatDescription     string `json:"statDescription"`
		StatID              string `json:"statId"`
		StatName            string `json:"statName"`
		UnitLabel           string `json:"unitLabel"`
		UnitType            int    `json:"unitType"`
		Weight              int    `json:"weight"`
	} `json:"medalStreakCombined"`
	MedalStreakShutdown struct {
		Category            int    `json:"category"`
		Group               int    `json:"group"`
		MedalTierHash       int    `json:"medalTierHash"`
		MedalTierIdentifier string `json:"medalTierIdentifier"`
		Modes               []int  `json:"modes"`
		PeriodTypes         []int  `json:"periodTypes"`
		StatDescription     string `json:"statDescription"`
		StatID              string `json:"statId"`
		StatName            string `json:"statName"`
		UnitLabel           string `json:"unitLabel"`
		UnitType            int    `json:"unitType"`
		Weight              int    `json:"weight"`
	} `json:"medalStreakShutdown"`
	MedalStreakTeam struct {
		Category            int    `json:"category"`
		Group               int    `json:"group"`
		MedalTierHash       int    `json:"medalTierHash"`
		MedalTierIdentifier string `json:"medalTierIdentifier"`
		Modes               []int  `json:"modes"`
		PeriodTypes         []int  `json:"periodTypes"`
		StatDescription     string `json:"statDescription"`
		StatID              string `json:"statId"`
		StatName            string `json:"statName"`
		UnitLabel           string `json:"unitLabel"`
		UnitType            int    `json:"unitType"`
		Weight              int    `json:"weight"`
	} `json:"medalStreakTeam"`
	MedalSuperShutdown struct {
		Category            int    `json:"category"`
		Group               int    `json:"group"`
		MedalTierHash       int    `json:"medalTierHash"`
		MedalTierIdentifier string `json:"medalTierIdentifier"`
		Modes               []int  `json:"modes"`
		PeriodTypes         []int  `json:"periodTypes"`
		StatDescription     string `json:"statDescription"`
		StatID              string `json:"statId"`
		StatName            string `json:"statName"`
		UnitLabel           string `json:"unitLabel"`
		UnitType            int    `json:"unitType"`
		Weight              int    `json:"weight"`
	} `json:"medalSuperShutdown"`
	MedalSupremacyCrestCreditStreak struct {
		Category            int    `json:"category"`
		Group               int    `json:"group"`
		MedalTierHash       int    `json:"medalTierHash"`
		MedalTierIdentifier string `json:"medalTierIdentifier"`
		Modes               []int  `json:"modes"`
		PeriodTypes         []int  `json:"periodTypes"`
		StatDescription     string `json:"statDescription"`
		StatID              string `json:"statId"`
		StatName            string `json:"statName"`
		UnitLabel           string `json:"unitLabel"`
		UnitType            int    `json:"unitType"`
		Weight              int    `json:"weight"`
	} `json:"medalSupremacyCrestCreditStreak"`
	MedalSupremacyFirstCrest struct {
		Category            int    `json:"category"`
		Group               int    `json:"group"`
		MedalTierHash       int    `json:"medalTierHash"`
		MedalTierIdentifier string `json:"medalTierIdentifier"`
		Modes               []int  `json:"modes"`
		PeriodTypes         []int  `json:"periodTypes"`
		StatDescription     string `json:"statDescription"`
		StatID              string `json:"statId"`
		StatName            string `json:"statName"`
		UnitLabel           string `json:"unitLabel"`
		UnitType            int    `json:"unitType"`
		Weight              int    `json:"weight"`
	} `json:"medalSupremacyFirstCrest"`
	MedalSupremacyNeverCollected struct {
		Category            int    `json:"category"`
		Group               int    `json:"group"`
		MedalTierHash       int    `json:"medalTierHash"`
		MedalTierIdentifier string `json:"medalTierIdentifier"`
		Modes               []int  `json:"modes"`
		PeriodTypes         []int  `json:"periodTypes"`
		StatDescription     string `json:"statDescription"`
		StatID              string `json:"statId"`
		StatName            string `json:"statName"`
		UnitLabel           string `json:"unitLabel"`
		UnitType            int    `json:"unitType"`
		Weight              int    `json:"weight"`
	} `json:"medalSupremacyNeverCollected"`
	MedalSupremacyPerfectSecureRate struct {
		Category            int    `json:"category"`
		Group               int    `json:"group"`
		MedalTierHash       int    `json:"medalTierHash"`
		MedalTierIdentifier string `json:"medalTierIdentifier"`
		Modes               []int  `json:"modes"`
		PeriodTypes         []int  `json:"periodTypes"`
		StatDescription     string `json:"statDescription"`
		StatID              string `json:"statId"`
		StatName            string `json:"statName"`
		UnitLabel           string `json:"unitLabel"`
		UnitType            int    `json:"unitType"`
		Weight              int    `json:"weight"`
	} `json:"medalSupremacyPerfectSecureRate"`
	MedalSupremacyRecoverStreak struct {
		Category            int    `json:"category"`
		Group               int    `json:"group"`
		MedalTierHash       int    `json:"medalTierHash"`
		MedalTierIdentifier string `json:"medalTierIdentifier"`
		Modes               []int  `json:"modes"`
		PeriodTypes         []int  `json:"periodTypes"`
		StatDescription     string `json:"statDescription"`
		StatID              string `json:"statId"`
		StatName            string `json:"statName"`
		UnitLabel           string `json:"unitLabel"`
		UnitType            int    `json:"unitType"`
		Weight              int    `json:"weight"`
	} `json:"medalSupremacyRecoverStreak"`
	MedalSupremacySecureStreak struct {
		Category            int    `json:"category"`
		Group               int    `json:"group"`
		MedalTierHash       int    `json:"medalTierHash"`
		MedalTierIdentifier string `json:"medalTierIdentifier"`
		Modes               []int  `json:"modes"`
		PeriodTypes         []int  `json:"periodTypes"`
		StatDescription     string `json:"statDescription"`
		StatID              string `json:"statId"`
		StatName            string `json:"statName"`
		UnitLabel           string `json:"unitLabel"`
		UnitType            int    `json:"unitType"`
		Weight              int    `json:"weight"`
	} `json:"medalSupremacySecureStreak"`
	MedalSurvivalComeback struct {
		Category            int    `json:"category"`
		Group               int    `json:"group"`
		MedalTierHash       int    `json:"medalTierHash"`
		MedalTierIdentifier string `json:"medalTierIdentifier"`
		Modes               []int  `json:"modes"`
		PeriodTypes         []int  `json:"periodTypes"`
		StatDescription     string `json:"statDescription"`
		StatID              string `json:"statId"`
		StatName            string `json:"statName"`
		UnitLabel           string `json:"unitLabel"`
		UnitType            int    `json:"unitType"`
		Weight              int    `json:"weight"`
	} `json:"medalSurvivalComeback"`
	MedalSurvivalKnockout struct {
		Category            int    `json:"category"`
		Group               int    `json:"group"`
		MedalTierHash       int    `json:"medalTierHash"`
		MedalTierIdentifier string `json:"medalTierIdentifier"`
		Modes               []int  `json:"modes"`
		PeriodTypes         []int  `json:"periodTypes"`
		StatDescription     string `json:"statDescription"`
		StatID              string `json:"statId"`
		StatName            string `json:"statName"`
		UnitLabel           string `json:"unitLabel"`
		UnitType            int    `json:"unitType"`
		Weight              int    `json:"weight"`
	} `json:"medalSurvivalKnockout"`
	MedalSurvivalQuickWipe struct {
		Category            int    `json:"category"`
		Group               int    `json:"group"`
		MedalTierHash       int    `json:"medalTierHash"`
		MedalTierIdentifier string `json:"medalTierIdentifier"`
		Modes               []int  `json:"modes"`
		PeriodTypes         []int  `json:"periodTypes"`
		StatDescription     string `json:"statDescription"`
		StatID              string `json:"statId"`
		StatName            string `json:"statName"`
		UnitLabel           string `json:"unitLabel"`
		UnitType            int    `json:"unitType"`
		Weight              int    `json:"weight"`
	} `json:"medalSurvivalQuickWipe"`
	MedalSurvivalTeamUndefeated struct {
		Category            int    `json:"category"`
		Group               int    `json:"group"`
		MedalTierHash       int    `json:"medalTierHash"`
		MedalTierIdentifier string `json:"medalTierIdentifier"`
		Modes               []int  `json:"modes"`
		PeriodTypes         []int  `json:"periodTypes"`
		StatDescription     string `json:"statDescription"`
		StatID              string `json:"statId"`
		StatName            string `json:"statName"`
		UnitLabel           string `json:"unitLabel"`
		UnitType            int    `json:"unitType"`
		Weight              int    `json:"weight"`
	} `json:"medalSurvivalTeamUndefeated"`
	MedalSurvivalUndefeated struct {
		Category            int    `json:"category"`
		Group               int    `json:"group"`
		MedalTierHash       int    `json:"medalTierHash"`
		MedalTierIdentifier string `json:"medalTierIdentifier"`
		Modes               []int  `json:"modes"`
		PeriodTypes         []int  `json:"periodTypes"`
		StatDescription     string `json:"statDescription"`
		StatID              string `json:"statId"`
		StatName            string `json:"statName"`
		UnitLabel           string `json:"unitLabel"`
		UnitType            int    `json:"unitType"`
		Weight              int    `json:"weight"`
	} `json:"medalSurvivalUndefeated"`
	MedalSurvivalWinLastStand struct {
		Category            int    `json:"category"`
		Group               int    `json:"group"`
		MedalTierHash       int    `json:"medalTierHash"`
		MedalTierIdentifier string `json:"medalTierIdentifier"`
		Modes               []int  `json:"modes"`
		PeriodTypes         []int  `json:"periodTypes"`
		StatDescription     string `json:"statDescription"`
		StatID              string `json:"statId"`
		StatName            string `json:"statName"`
		UnitLabel           string `json:"unitLabel"`
		UnitType            int    `json:"unitType"`
		Weight              int    `json:"weight"`
	} `json:"medalSurvivalWinLastStand"`
	MedalWeaponAuto struct {
		Category            int    `json:"category"`
		Group               int    `json:"group"`
		MedalTierHash       int    `json:"medalTierHash"`
		MedalTierIdentifier string `json:"medalTierIdentifier"`
		Modes               []int  `json:"modes"`
		PeriodTypes         []int  `json:"periodTypes"`
		StatDescription     string `json:"statDescription"`
		StatID              string `json:"statId"`
		StatName            string `json:"statName"`
		UnitLabel           string `json:"unitLabel"`
		UnitType            int    `json:"unitType"`
		Weight              int    `json:"weight"`
	} `json:"medalWeaponAuto"`
	MedalWeaponFusion struct {
		Category            int    `json:"category"`
		Group               int    `json:"group"`
		MedalTierHash       int    `json:"medalTierHash"`
		MedalTierIdentifier string `json:"medalTierIdentifier"`
		Modes               []int  `json:"modes"`
		PeriodTypes         []int  `json:"periodTypes"`
		StatDescription     string `json:"statDescription"`
		StatID              string `json:"statId"`
		StatName            string `json:"statName"`
		UnitLabel           string `json:"unitLabel"`
		UnitType            int    `json:"unitType"`
		Weight              int    `json:"weight"`
	} `json:"medalWeaponFusion"`
	MedalWeaponGrenade struct {
		Category            int    `json:"category"`
		Group               int    `json:"group"`
		MedalTierHash       int    `json:"medalTierHash"`
		MedalTierIdentifier string `json:"medalTierIdentifier"`
		Modes               []int  `json:"modes"`
		PeriodTypes         []int  `json:"periodTypes"`
		StatDescription     string `json:"statDescription"`
		StatID              string `json:"statId"`
		StatName            string `json:"statName"`
		UnitLabel           string `json:"unitLabel"`
		UnitType            int    `json:"unitType"`
		Weight              int    `json:"weight"`
	} `json:"medalWeaponGrenade"`
	MedalWeaponHandCannon struct {
		Category            int    `json:"category"`
		Group               int    `json:"group"`
		MedalTierHash       int    `json:"medalTierHash"`
		MedalTierIdentifier string `json:"medalTierIdentifier"`
		Modes               []int  `json:"modes"`
		PeriodTypes         []int  `json:"periodTypes"`
		StatDescription     string `json:"statDescription"`
		StatID              string `json:"statId"`
		StatName            string `json:"statName"`
		UnitLabel           string `json:"unitLabel"`
		UnitType            int    `json:"unitType"`
		Weight              int    `json:"weight"`
	} `json:"medalWeaponHandCannon"`
	MedalWeaponPulse struct {
		Category            int    `json:"category"`
		Group               int    `json:"group"`
		MedalTierHash       int    `json:"medalTierHash"`
		MedalTierIdentifier string `json:"medalTierIdentifier"`
		Modes               []int  `json:"modes"`
		PeriodTypes         []int  `json:"periodTypes"`
		StatDescription     string `json:"statDescription"`
		StatID              string `json:"statId"`
		StatName            string `json:"statName"`
		UnitLabel           string `json:"unitLabel"`
		UnitType            int    `json:"unitType"`
		Weight              int    `json:"weight"`
	} `json:"medalWeaponPulse"`
	MedalWeaponRocket struct {
		Category            int    `json:"category"`
		Group               int    `json:"group"`
		MedalTierHash       int    `json:"medalTierHash"`
		MedalTierIdentifier string `json:"medalTierIdentifier"`
		Modes               []int  `json:"modes"`
		PeriodTypes         []int  `json:"periodTypes"`
		StatDescription     string `json:"statDescription"`
		StatID              string `json:"statId"`
		StatName            string `json:"statName"`
		UnitLabel           string `json:"unitLabel"`
		UnitType            int    `json:"unitType"`
		Weight              int    `json:"weight"`
	} `json:"medalWeaponRocket"`
	MedalWeaponScout struct {
		Category            int    `json:"category"`
		Group               int    `json:"group"`
		MedalTierHash       int    `json:"medalTierHash"`
		MedalTierIdentifier string `json:"medalTierIdentifier"`
		Modes               []int  `json:"modes"`
		PeriodTypes         []int  `json:"periodTypes"`
		StatDescription     string `json:"statDescription"`
		StatID              string `json:"statId"`
		StatName            string `json:"statName"`
		UnitLabel           string `json:"unitLabel"`
		UnitType            int    `json:"unitType"`
		Weight              int    `json:"weight"`
	} `json:"medalWeaponScout"`
	MedalWeaponShotgun struct {
		Category            int    `json:"category"`
		Group               int    `json:"group"`
		MedalTierHash       int    `json:"medalTierHash"`
		MedalTierIdentifier string `json:"medalTierIdentifier"`
		Modes               []int  `json:"modes"`
		PeriodTypes         []int  `json:"periodTypes"`
		StatDescription     string `json:"statDescription"`
		StatID              string `json:"statId"`
		StatName            string `json:"statName"`
		UnitLabel           string `json:"unitLabel"`
		UnitType            int    `json:"unitType"`
		Weight              int    `json:"weight"`
	} `json:"medalWeaponShotgun"`
	MedalWeaponSidearm struct {
		Category            int    `json:"category"`
		Group               int    `json:"group"`
		MedalTierHash       int    `json:"medalTierHash"`
		MedalTierIdentifier string `json:"medalTierIdentifier"`
		Modes               []int  `json:"modes"`
		PeriodTypes         []int  `json:"periodTypes"`
		StatDescription     string `json:"statDescription"`
		StatID              string `json:"statId"`
		StatName            string `json:"statName"`
		UnitLabel           string `json:"unitLabel"`
		UnitType            int    `json:"unitType"`
		Weight              int    `json:"weight"`
	} `json:"medalWeaponSidearm"`
	MedalWeaponSmg struct {
		Category            int    `json:"category"`
		Group               int    `json:"group"`
		MedalTierHash       int    `json:"medalTierHash"`
		MedalTierIdentifier string `json:"medalTierIdentifier"`
		Modes               []int  `json:"modes"`
		PeriodTypes         []int  `json:"periodTypes"`
		StatDescription     string `json:"statDescription"`
		StatID              string `json:"statId"`
		StatName            string `json:"statName"`
		UnitLabel           string `json:"unitLabel"`
		UnitType            int    `json:"unitType"`
		Weight              int    `json:"weight"`
	} `json:"medalWeaponSmg"`
	MedalWeaponSniper struct {
		Category            int    `json:"category"`
		Group               int    `json:"group"`
		MedalTierHash       int    `json:"medalTierHash"`
		MedalTierIdentifier string `json:"medalTierIdentifier"`
		Modes               []int  `json:"modes"`
		PeriodTypes         []int  `json:"periodTypes"`
		StatDescription     string `json:"statDescription"`
		StatID              string `json:"statId"`
		StatName            string `json:"statName"`
		UnitLabel           string `json:"unitLabel"`
		UnitType            int    `json:"unitType"`
		Weight              int    `json:"weight"`
	} `json:"medalWeaponSniper"`
	MedalWeaponSword struct {
		Category            int    `json:"category"`
		Group               int    `json:"group"`
		MedalTierHash       int    `json:"medalTierHash"`
		MedalTierIdentifier string `json:"medalTierIdentifier"`
		Modes               []int  `json:"modes"`
		PeriodTypes         []int  `json:"periodTypes"`
		StatDescription     string `json:"statDescription"`
		StatID              string `json:"statId"`
		StatName            string `json:"statName"`
		UnitLabel           string `json:"unitLabel"`
		UnitType            int    `json:"unitType"`
		Weight              int    `json:"weight"`
	} `json:"medalWeaponSword"`
	MedalsUnknown struct {
		Category    int    `json:"category"`
		Group       int    `json:"group"`
		Modes       []int  `json:"modes"`
		PeriodTypes []int  `json:"periodTypes"`
		StatID      string `json:"statId"`
		StatName    string `json:"statName"`
		UnitLabel   string `json:"unitLabel"`
		UnitType    int    `json:"unitType"`
		Weight      int    `json:"weight"`
	} `json:"medalsUnknown"`
	MostPrecisionKills struct {
		Category    int    `json:"category"`
		Group       int    `json:"group"`
		Modes       []int  `json:"modes"`
		PeriodTypes []int  `json:"periodTypes"`
		StatID      string `json:"statId"`
		StatName    string `json:"statName"`
		UnitLabel   string `json:"unitLabel"`
		UnitType    int    `json:"unitType"`
		Weight      int    `json:"weight"`
	} `json:"mostPrecisionKills"`
	ObjectivesCompleted struct {
		Category    int    `json:"category"`
		Group       int    `json:"group"`
		Modes       []int  `json:"modes"`
		PeriodTypes []int  `json:"periodTypes"`
		StatID      string `json:"statId"`
		StatName    string `json:"statName"`
		UnitLabel   string `json:"unitLabel"`
		UnitType    int    `json:"unitType"`
		Weight      int    `json:"weight"`
	} `json:"objectivesCompleted"`
	OrbsDropped struct {
		Category        int    `json:"category"`
		Group           int    `json:"group"`
		Modes           []int  `json:"modes"`
		PeriodTypes     []int  `json:"periodTypes"`
		StatDescription string `json:"statDescription"`
		StatID          string `json:"statId"`
		StatName        string `json:"statName"`
		UnitLabel       string `json:"unitLabel"`
		UnitType        int    `json:"unitType"`
		Weight          int    `json:"weight"`
	} `json:"orbsDropped"`
	OrbsGathered struct {
		Category    int    `json:"category"`
		Group       int    `json:"group"`
		Modes       []int  `json:"modes"`
		PeriodTypes []int  `json:"periodTypes"`
		StatID      string `json:"statId"`
		StatName    string `json:"statName"`
		UnitLabel   string `json:"unitLabel"`
		UnitType    int    `json:"unitType"`
		Weight      int    `json:"weight"`
	} `json:"orbsGathered"`
	PlayerCount struct {
		Category    int    `json:"category"`
		Group       int    `json:"group"`
		Modes       []int  `json:"modes"`
		PeriodTypes []int  `json:"periodTypes"`
		StatID      string `json:"statId"`
		StatName    string `json:"statName"`
		UnitLabel   string `json:"unitLabel"`
		UnitType    int    `json:"unitType"`
		Weight      int    `json:"weight"`
	} `json:"playerCount"`
	PrecisionKills struct {
		Category        int    `json:"category"`
		Group           int    `json:"group"`
		Modes           []int  `json:"modes"`
		PeriodTypes     []int  `json:"periodTypes"`
		StatDescription string `json:"statDescription"`
		StatID          string `json:"statId"`
		StatName        string `json:"statName"`
		UnitLabel       string `json:"unitLabel"`
		UnitType        int    `json:"unitType"`
		Weight          int    `json:"weight"`
	} `json:"precisionKills"`
	PublicEventsCompleted struct {
		Category        int    `json:"category"`
		Group           int    `json:"group"`
		Modes           []int  `json:"modes"`
		PeriodTypes     []int  `json:"periodTypes"`
		StatDescription string `json:"statDescription"`
		StatID          string `json:"statId"`
		StatName        string `json:"statName"`
		UnitLabel       string `json:"unitLabel"`
		UnitType        int    `json:"unitType"`
		Weight          int    `json:"weight"`
	} `json:"publicEventsCompleted"`
	PublicEventsJoined struct {
		Category    int    `json:"category"`
		Group       int    `json:"group"`
		Modes       []int  `json:"modes"`
		PeriodTypes []int  `json:"periodTypes"`
		StatID      string `json:"statId"`
		StatName    string `json:"statName"`
		UnitLabel   string `json:"unitLabel"`
		UnitType    int    `json:"unitType"`
		Weight      int    `json:"weight"`
	} `json:"publicEventsJoined"`
	RaceCompletionMilliseconds struct {
		Category    int           `json:"category"`
		Group       int           `json:"group"`
		Modes       []interface{} `json:"modes"`
		PeriodTypes []int         `json:"periodTypes"`
		StatID      string        `json:"statId"`
		StatName    string        `json:"statName"`
		UnitLabel   string        `json:"unitLabel"`
		UnitType    int           `json:"unitType"`
		Weight      int           `json:"weight"`
	} `json:"raceCompletionMilliseconds"`
	RaceCompletionSeconds struct {
		Category    int           `json:"category"`
		Group       int           `json:"group"`
		Modes       []interface{} `json:"modes"`
		PeriodTypes []int         `json:"periodTypes"`
		StatID      string        `json:"statId"`
		StatName    string        `json:"statName"`
		UnitLabel   string        `json:"unitLabel"`
		UnitType    int           `json:"unitType"`
		Weight      int           `json:"weight"`
	} `json:"raceCompletionSeconds"`
	RemainingTimeAfterQuitSeconds struct {
		Category    int    `json:"category"`
		Group       int    `json:"group"`
		Modes       []int  `json:"modes"`
		PeriodTypes []int  `json:"periodTypes"`
		StatID      string `json:"statId"`
		StatName    string `json:"statName"`
		UnitLabel   string `json:"unitLabel"`
		UnitType    int    `json:"unitType"`
		Weight      int    `json:"weight"`
	} `json:"remainingTimeAfterQuitSeconds"`
	ResurrectionsPerformed struct {
		Category        int    `json:"category"`
		Group           int    `json:"group"`
		Modes           []int  `json:"modes"`
		PeriodTypes     []int  `json:"periodTypes"`
		StatDescription string `json:"statDescription"`
		StatID          string `json:"statId"`
		StatName        string `json:"statName"`
		UnitLabel       string `json:"unitLabel"`
		UnitType        int    `json:"unitType"`
		Weight          int    `json:"weight"`
	} `json:"resurrectionsPerformed"`
	ResurrectionsReceived struct {
		Category    int    `json:"category"`
		Group       int    `json:"group"`
		Modes       []int  `json:"modes"`
		PeriodTypes []int  `json:"periodTypes"`
		StatID      string `json:"statId"`
		StatName    string `json:"statName"`
		UnitLabel   string `json:"unitLabel"`
		UnitType    int    `json:"unitType"`
		Weight      int    `json:"weight"`
	} `json:"resurrectionsReceived"`
	Score struct {
		Category        int    `json:"category"`
		Group           int    `json:"group"`
		Modes           []int  `json:"modes"`
		PeriodTypes     []int  `json:"periodTypes"`
		StatDescription string `json:"statDescription"`
		StatID          string `json:"statId"`
		StatName        string `json:"statName"`
		UnitLabel       string `json:"unitLabel"`
		UnitType        int    `json:"unitType"`
		Weight          int    `json:"weight"`
	} `json:"score"`
	SecondsPlayed struct {
		Category    int    `json:"category"`
		Group       int    `json:"group"`
		Modes       []int  `json:"modes"`
		PeriodTypes []int  `json:"periodTypes"`
		StatID      string `json:"statId"`
		StatName    string `json:"statName"`
		UnitLabel   string `json:"unitLabel"`
		UnitType    int    `json:"unitType"`
		Weight      int    `json:"weight"`
	} `json:"secondsPlayed"`
	SlamDunks struct {
		Category    int           `json:"category"`
		Group       int           `json:"group"`
		Modes       []interface{} `json:"modes"`
		PeriodTypes []int         `json:"periodTypes"`
		StatID      string        `json:"statId"`
		StatName    string        `json:"statName"`
		UnitLabel   string        `json:"unitLabel"`
		UnitType    int           `json:"unitType"`
		Weight      int           `json:"weight"`
	} `json:"slamDunks"`
	SparksCaptured struct {
		Category    int           `json:"category"`
		Group       int           `json:"group"`
		Modes       []interface{} `json:"modes"`
		PeriodTypes []int         `json:"periodTypes"`
		StatID      string        `json:"statId"`
		StatName    string        `json:"statName"`
		UnitLabel   string        `json:"unitLabel"`
		UnitType    int           `json:"unitType"`
		Weight      int           `json:"weight"`
	} `json:"sparksCaptured"`
	Standing struct {
		Category    int    `json:"category"`
		Group       int    `json:"group"`
		Modes       []int  `json:"modes"`
		PeriodTypes []int  `json:"periodTypes"`
		StatID      string `json:"statId"`
		StatName    string `json:"statName"`
		UnitLabel   string `json:"unitLabel"`
		UnitType    int    `json:"unitType"`
		Weight      int    `json:"weight"`
	} `json:"standing"`
	StartSeconds struct {
		Category    int    `json:"category"`
		Group       int    `json:"group"`
		Modes       []int  `json:"modes"`
		PeriodTypes []int  `json:"periodTypes"`
		StatID      string `json:"statId"`
		StatName    string `json:"statName"`
		UnitLabel   string `json:"unitLabel"`
		UnitType    int    `json:"unitType"`
		Weight      int    `json:"weight"`
	} `json:"startSeconds"`
	StyleDunks struct {
		Category    int           `json:"category"`
		Group       int           `json:"group"`
		Modes       []interface{} `json:"modes"`
		PeriodTypes []int         `json:"periodTypes"`
		StatID      string        `json:"statId"`
		StatName    string        `json:"statName"`
		UnitLabel   string        `json:"unitLabel"`
		UnitType    int           `json:"unitType"`
		Weight      int           `json:"weight"`
	} `json:"styleDunks"`
	Suicides struct {
		Category    int    `json:"category"`
		Group       int    `json:"group"`
		Modes       []int  `json:"modes"`
		PeriodTypes []int  `json:"periodTypes"`
		StatID      string `json:"statId"`
		StatName    string `json:"statName"`
		UnitLabel   string `json:"unitLabel"`
		UnitType    int    `json:"unitType"`
		Weight      int    `json:"weight"`
	} `json:"suicides"`
	Team struct {
		Category    int    `json:"category"`
		Group       int    `json:"group"`
		Modes       []int  `json:"modes"`
		PeriodTypes []int  `json:"periodTypes"`
		StatID      string `json:"statId"`
		StatName    string `json:"statName"`
		UnitLabel   string `json:"unitLabel"`
		UnitType    int    `json:"unitType"`
		Weight      int    `json:"weight"`
	} `json:"team"`
	TeamScore struct {
		Category    int    `json:"category"`
		Group       int    `json:"group"`
		Modes       []int  `json:"modes"`
		PeriodTypes []int  `json:"periodTypes"`
		StatID      string `json:"statId"`
		StatName    string `json:"statName"`
		UnitLabel   string `json:"unitLabel"`
		UnitType    int    `json:"unitType"`
		Weight      int    `json:"weight"`
	} `json:"teamScore"`
	TimePlayedSeconds struct {
		Category    int    `json:"category"`
		Group       int    `json:"group"`
		Modes       []int  `json:"modes"`
		PeriodTypes []int  `json:"periodTypes"`
		StatID      string `json:"statId"`
		StatName    string `json:"statName"`
		UnitLabel   string `json:"unitLabel"`
		UnitType    int    `json:"unitType"`
		Weight      int    `json:"weight"`
	} `json:"timePlayedSeconds"`
	TotalActivityDurationSeconds struct {
		Category    int    `json:"category"`
		Group       int    `json:"group"`
		Modes       []int  `json:"modes"`
		PeriodTypes []int  `json:"periodTypes"`
		StatID      string `json:"statId"`
		StatName    string `json:"statName"`
		UnitLabel   string `json:"unitLabel"`
		UnitType    int    `json:"unitType"`
		Weight      int    `json:"weight"`
	} `json:"totalActivityDurationSeconds"`
	TotalDeathDistance struct {
		Category    int    `json:"category"`
		Group       int    `json:"group"`
		Modes       []int  `json:"modes"`
		PeriodTypes []int  `json:"periodTypes"`
		StatID      string `json:"statId"`
		StatName    string `json:"statName"`
		UnitLabel   string `json:"unitLabel"`
		UnitType    int    `json:"unitType"`
		Weight      int    `json:"weight"`
	} `json:"totalDeathDistance"`
	TotalKillDistance struct {
		Category    int    `json:"category"`
		Group       int    `json:"group"`
		Modes       []int  `json:"modes"`
		PeriodTypes []int  `json:"periodTypes"`
		StatID      string `json:"statId"`
		StatName    string `json:"statName"`
		UnitLabel   string `json:"unitLabel"`
		UnitType    int    `json:"unitType"`
		Weight      int    `json:"weight"`
	} `json:"totalKillDistance"`
	UniqueWeaponAssistDamage struct {
		Category    int    `json:"category"`
		Group       int    `json:"group"`
		Modes       []int  `json:"modes"`
		PeriodTypes []int  `json:"periodTypes"`
		StatID      string `json:"statId"`
		StatName    string `json:"statName"`
		UnitLabel   string `json:"unitLabel"`
		UnitType    int    `json:"unitType"`
		Weight      int    `json:"weight"`
	} `json:"uniqueWeaponAssistDamage"`
	UniqueWeaponAssists struct {
		Category    int    `json:"category"`
		Group       int    `json:"group"`
		Modes       []int  `json:"modes"`
		PeriodTypes []int  `json:"periodTypes"`
		StatID      string `json:"statId"`
		StatName    string `json:"statName"`
		UnitLabel   string `json:"unitLabel"`
		UnitType    int    `json:"unitType"`
		Weight      int    `json:"weight"`
	} `json:"uniqueWeaponAssists"`
	UniqueWeaponKills struct {
		Category    int    `json:"category"`
		Group       int    `json:"group"`
		Modes       []int  `json:"modes"`
		PeriodTypes []int  `json:"periodTypes"`
		StatID      string `json:"statId"`
		StatName    string `json:"statName"`
		UnitLabel   string `json:"unitLabel"`
		UnitType    int    `json:"unitType"`
		Weight      int    `json:"weight"`
	} `json:"uniqueWeaponKills"`
	UniqueWeaponKillsPrecisionKills struct {
		Category    int    `json:"category"`
		Group       int    `json:"group"`
		Modes       []int  `json:"modes"`
		PeriodTypes []int  `json:"periodTypes"`
		StatID      string `json:"statId"`
		StatName    string `json:"statName"`
		UnitLabel   string `json:"unitLabel"`
		UnitType    int    `json:"unitType"`
		Weight      int    `json:"weight"`
	} `json:"uniqueWeaponKillsPrecisionKills"`
	UniqueWeaponPrecisionKills struct {
		Category    int    `json:"category"`
		Group       int    `json:"group"`
		Modes       []int  `json:"modes"`
		PeriodTypes []int  `json:"periodTypes"`
		StatID      string `json:"statId"`
		StatName    string `json:"statName"`
		UnitLabel   string `json:"unitLabel"`
		UnitType    int    `json:"unitType"`
		Weight      int    `json:"weight"`
	} `json:"uniqueWeaponPrecisionKills"`
	WeaponBestType struct {
		Category        int    `json:"category"`
		Group           int    `json:"group"`
		Modes           []int  `json:"modes"`
		PeriodTypes     []int  `json:"periodTypes"`
		StatDescription string `json:"statDescription"`
		StatID          string `json:"statId"`
		StatName        string `json:"statName"`
		UnitLabel       string `json:"unitLabel"`
		UnitType        int    `json:"unitType"`
		Weight          int    `json:"weight"`
	} `json:"weaponBestType"`
	WeaponKillsAbility struct {
		Category    int    `json:"category"`
		Group       int    `json:"group"`
		Modes       []int  `json:"modes"`
		PeriodTypes []int  `json:"periodTypes"`
		StatID      string `json:"statId"`
		StatName    string `json:"statName"`
		UnitLabel   string `json:"unitLabel"`
		UnitType    int    `json:"unitType"`
		Weight      int    `json:"weight"`
	} `json:"weaponKillsAbility"`
	WeaponKillsAutoRifle struct {
		Category    int    `json:"category"`
		Group       int    `json:"group"`
		Modes       []int  `json:"modes"`
		PeriodTypes []int  `json:"periodTypes"`
		StatID      string `json:"statId"`
		StatName    string `json:"statName"`
		UnitLabel   string `json:"unitLabel"`
		UnitType    int    `json:"unitType"`
		Weight      int    `json:"weight"`
	} `json:"weaponKillsAutoRifle"`
	WeaponKillsFusionRifle struct {
		Category    int    `json:"category"`
		Group       int    `json:"group"`
		Modes       []int  `json:"modes"`
		PeriodTypes []int  `json:"periodTypes"`
		StatID      string `json:"statId"`
		StatName    string `json:"statName"`
		UnitLabel   string `json:"unitLabel"`
		UnitType    int    `json:"unitType"`
		Weight      int    `json:"weight"`
	} `json:"weaponKillsFusionRifle"`
	WeaponKillsGrenade struct {
		Category    int    `json:"category"`
		Group       int    `json:"group"`
		Modes       []int  `json:"modes"`
		PeriodTypes []int  `json:"periodTypes"`
		StatID      string `json:"statId"`
		StatName    string `json:"statName"`
		UnitLabel   string `json:"unitLabel"`
		UnitType    int    `json:"unitType"`
		Weight      int    `json:"weight"`
	} `json:"weaponKillsGrenade"`
	WeaponKillsHandCannon struct {
		Category    int    `json:"category"`
		Group       int    `json:"group"`
		Modes       []int  `json:"modes"`
		PeriodTypes []int  `json:"periodTypes"`
		StatID      string `json:"statId"`
		StatName    string `json:"statName"`
		UnitLabel   string `json:"unitLabel"`
		UnitType    int    `json:"unitType"`
		Weight      int    `json:"weight"`
	} `json:"weaponKillsHandCannon"`
	WeaponKillsMachinegun struct {
		Category    int    `json:"category"`
		Group       int    `json:"group"`
		Modes       []int  `json:"modes"`
		PeriodTypes []int  `json:"periodTypes"`
		StatID      string `json:"statId"`
		StatName    string `json:"statName"`
		UnitLabel   string `json:"unitLabel"`
		UnitType    int    `json:"unitType"`
		Weight      int    `json:"weight"`
	} `json:"weaponKillsMachinegun"`
	WeaponKillsMelee struct {
		Category    int    `json:"category"`
		Group       int    `json:"group"`
		Modes       []int  `json:"modes"`
		PeriodTypes []int  `json:"periodTypes"`
		StatID      string `json:"statId"`
		StatName    string `json:"statName"`
		UnitLabel   string `json:"unitLabel"`
		UnitType    int    `json:"unitType"`
		Weight      int    `json:"weight"`
	} `json:"weaponKillsMelee"`
	WeaponKillsPrecisionKillsAutoRifle struct {
		Category    int    `json:"category"`
		Group       int    `json:"group"`
		Modes       []int  `json:"modes"`
		PeriodTypes []int  `json:"periodTypes"`
		StatID      string `json:"statId"`
		StatName    string `json:"statName"`
		UnitLabel   string `json:"unitLabel"`
		UnitType    int    `json:"unitType"`
		Weight      int    `json:"weight"`
	} `json:"weaponKillsPrecisionKillsAutoRifle"`
	WeaponKillsPrecisionKillsFusionRifle struct {
		Category    int    `json:"category"`
		Group       int    `json:"group"`
		Modes       []int  `json:"modes"`
		PeriodTypes []int  `json:"periodTypes"`
		StatID      string `json:"statId"`
		StatName    string `json:"statName"`
		UnitLabel   string `json:"unitLabel"`
		UnitType    int    `json:"unitType"`
		Weight      int    `json:"weight"`
	} `json:"weaponKillsPrecisionKillsFusionRifle"`
	WeaponKillsPrecisionKillsGrenade struct {
		Category    int    `json:"category"`
		Group       int    `json:"group"`
		Modes       []int  `json:"modes"`
		PeriodTypes []int  `json:"periodTypes"`
		StatID      string `json:"statId"`
		StatName    string `json:"statName"`
		UnitLabel   string `json:"unitLabel"`
		UnitType    int    `json:"unitType"`
		Weight      int    `json:"weight"`
	} `json:"weaponKillsPrecisionKillsGrenade"`
	WeaponKillsPrecisionKillsHandCannon struct {
		Category    int    `json:"category"`
		Group       int    `json:"group"`
		Modes       []int  `json:"modes"`
		PeriodTypes []int  `json:"periodTypes"`
		StatID      string `json:"statId"`
		StatName    string `json:"statName"`
		UnitLabel   string `json:"unitLabel"`
		UnitType    int    `json:"unitType"`
		Weight      int    `json:"weight"`
	} `json:"weaponKillsPrecisionKillsHandCannon"`
	WeaponKillsPrecisionKillsMachinegun struct {
		Category    int    `json:"category"`
		Group       int    `json:"group"`
		Modes       []int  `json:"modes"`
		PeriodTypes []int  `json:"periodTypes"`
		StatID      string `json:"statId"`
		StatName    string `json:"statName"`
		UnitLabel   string `json:"unitLabel"`
		UnitType    int    `json:"unitType"`
		Weight      int    `json:"weight"`
	} `json:"weaponKillsPrecisionKillsMachinegun"`
	WeaponKillsPrecisionKillsMelee struct {
		Category    int    `json:"category"`
		Group       int    `json:"group"`
		Modes       []int  `json:"modes"`
		PeriodTypes []int  `json:"periodTypes"`
		StatID      string `json:"statId"`
		StatName    string `json:"statName"`
		UnitLabel   string `json:"unitLabel"`
		UnitType    int    `json:"unitType"`
		Weight      int    `json:"weight"`
	} `json:"weaponKillsPrecisionKillsMelee"`
	WeaponKillsPrecisionKillsPulseRifle struct {
		Category    int    `json:"category"`
		Group       int    `json:"group"`
		Modes       []int  `json:"modes"`
		PeriodTypes []int  `json:"periodTypes"`
		StatID      string `json:"statId"`
		StatName    string `json:"statName"`
		UnitLabel   string `json:"unitLabel"`
		UnitType    int    `json:"unitType"`
		Weight      int    `json:"weight"`
	} `json:"weaponKillsPrecisionKillsPulseRifle"`
	WeaponKillsPrecisionKillsRelic struct {
		Category    int    `json:"category"`
		Group       int    `json:"group"`
		Modes       []int  `json:"modes"`
		PeriodTypes []int  `json:"periodTypes"`
		StatID      string `json:"statId"`
		StatName    string `json:"statName"`
		UnitLabel   string `json:"unitLabel"`
		UnitType    int    `json:"unitType"`
		Weight      int    `json:"weight"`
	} `json:"weaponKillsPrecisionKillsRelic"`
	WeaponKillsPrecisionKillsRocketLauncher struct {
		Category    int    `json:"category"`
		Group       int    `json:"group"`
		Modes       []int  `json:"modes"`
		PeriodTypes []int  `json:"periodTypes"`
		StatID      string `json:"statId"`
		StatName    string `json:"statName"`
		UnitLabel   string `json:"unitLabel"`
		UnitType    int    `json:"unitType"`
		Weight      int    `json:"weight"`
	} `json:"weaponKillsPrecisionKillsRocketLauncher"`
	WeaponKillsPrecisionKillsScoutRifle struct {
		Category    int    `json:"category"`
		Group       int    `json:"group"`
		Modes       []int  `json:"modes"`
		PeriodTypes []int  `json:"periodTypes"`
		StatID      string `json:"statId"`
		StatName    string `json:"statName"`
		UnitLabel   string `json:"unitLabel"`
		UnitType    int    `json:"unitType"`
		Weight      int    `json:"weight"`
	} `json:"weaponKillsPrecisionKillsScoutRifle"`
	WeaponKillsPrecisionKillsShotgun struct {
		Category    int    `json:"category"`
		Group       int    `json:"group"`
		Modes       []int  `json:"modes"`
		PeriodTypes []int  `json:"periodTypes"`
		StatID      string `json:"statId"`
		StatName    string `json:"statName"`
		UnitLabel   string `json:"unitLabel"`
		UnitType    int    `json:"unitType"`
		Weight      int    `json:"weight"`
	} `json:"weaponKillsPrecisionKillsShotgun"`
	WeaponKillsPrecisionKillsSideArm struct {
		Category    int    `json:"category"`
		Group       int    `json:"group"`
		Modes       []int  `json:"modes"`
		PeriodTypes []int  `json:"periodTypes"`
		StatID      string `json:"statId"`
		StatName    string `json:"statName"`
		UnitLabel   string `json:"unitLabel"`
		UnitType    int    `json:"unitType"`
		Weight      int    `json:"weight"`
	} `json:"weaponKillsPrecisionKillsSideArm"`
	WeaponKillsPrecisionKillsSniper struct {
		Category    int    `json:"category"`
		Group       int    `json:"group"`
		Modes       []int  `json:"modes"`
		PeriodTypes []int  `json:"periodTypes"`
		StatID      string `json:"statId"`
		StatName    string `json:"statName"`
		UnitLabel   string `json:"unitLabel"`
		UnitType    int    `json:"unitType"`
		Weight      int    `json:"weight"`
	} `json:"weaponKillsPrecisionKillsSniper"`
	WeaponKillsPrecisionKillsSubmachinegun struct {
		Category    int    `json:"category"`
		Group       int    `json:"group"`
		Modes       []int  `json:"modes"`
		PeriodTypes []int  `json:"periodTypes"`
		StatID      string `json:"statId"`
		StatName    string `json:"statName"`
		UnitLabel   string `json:"unitLabel"`
		UnitType    int    `json:"unitType"`
		Weight      int    `json:"weight"`
	} `json:"weaponKillsPrecisionKillsSubmachinegun"`
	WeaponKillsPrecisionKillsSuper struct {
		Category    int    `json:"category"`
		Group       int    `json:"group"`
		Modes       []int  `json:"modes"`
		PeriodTypes []int  `json:"periodTypes"`
		StatID      string `json:"statId"`
		StatName    string `json:"statName"`
		UnitLabel   string `json:"unitLabel"`
		UnitType    int    `json:"unitType"`
		Weight      int    `json:"weight"`
	} `json:"weaponKillsPrecisionKillsSuper"`
	WeaponKillsPulseRifle struct {
		Category    int    `json:"category"`
		Group       int    `json:"group"`
		Modes       []int  `json:"modes"`
		PeriodTypes []int  `json:"periodTypes"`
		StatID      string `json:"statId"`
		StatName    string `json:"statName"`
		UnitLabel   string `json:"unitLabel"`
		UnitType    int    `json:"unitType"`
		Weight      int    `json:"weight"`
	} `json:"weaponKillsPulseRifle"`
	WeaponKillsRelic struct {
		Category    int    `json:"category"`
		Group       int    `json:"group"`
		Modes       []int  `json:"modes"`
		PeriodTypes []int  `json:"periodTypes"`
		StatID      string `json:"statId"`
		StatName    string `json:"statName"`
		UnitLabel   string `json:"unitLabel"`
		UnitType    int    `json:"unitType"`
		Weight      int    `json:"weight"`
	} `json:"weaponKillsRelic"`
	WeaponKillsRocketLauncher struct {
		Category    int    `json:"category"`
		Group       int    `json:"group"`
		Modes       []int  `json:"modes"`
		PeriodTypes []int  `json:"periodTypes"`
		StatID      string `json:"statId"`
		StatName    string `json:"statName"`
		UnitLabel   string `json:"unitLabel"`
		UnitType    int    `json:"unitType"`
		Weight      int    `json:"weight"`
	} `json:"weaponKillsRocketLauncher"`
	WeaponKillsScoutRifle struct {
		Category    int    `json:"category"`
		Group       int    `json:"group"`
		Modes       []int  `json:"modes"`
		PeriodTypes []int  `json:"periodTypes"`
		StatID      string `json:"statId"`
		StatName    string `json:"statName"`
		UnitLabel   string `json:"unitLabel"`
		UnitType    int    `json:"unitType"`
		Weight      int    `json:"weight"`
	} `json:"weaponKillsScoutRifle"`
	WeaponKillsShotgun struct {
		Category    int    `json:"category"`
		Group       int    `json:"group"`
		Modes       []int  `json:"modes"`
		PeriodTypes []int  `json:"periodTypes"`
		StatID      string `json:"statId"`
		StatName    string `json:"statName"`
		UnitLabel   string `json:"unitLabel"`
		UnitType    int    `json:"unitType"`
		Weight      int    `json:"weight"`
	} `json:"weaponKillsShotgun"`
	WeaponKillsSideArm struct {
		Category    int    `json:"category"`
		Group       int    `json:"group"`
		Modes       []int  `json:"modes"`
		PeriodTypes []int  `json:"periodTypes"`
		StatID      string `json:"statId"`
		StatName    string `json:"statName"`
		UnitLabel   string `json:"unitLabel"`
		UnitType    int    `json:"unitType"`
		Weight      int    `json:"weight"`
	} `json:"weaponKillsSideArm"`
	WeaponKillsSniper struct {
		Category    int    `json:"category"`
		Group       int    `json:"group"`
		Modes       []int  `json:"modes"`
		PeriodTypes []int  `json:"periodTypes"`
		StatID      string `json:"statId"`
		StatName    string `json:"statName"`
		UnitLabel   string `json:"unitLabel"`
		UnitType    int    `json:"unitType"`
		Weight      int    `json:"weight"`
	} `json:"weaponKillsSniper"`
	WeaponKillsSubmachinegun struct {
		Category    int    `json:"category"`
		Group       int    `json:"group"`
		Modes       []int  `json:"modes"`
		PeriodTypes []int  `json:"periodTypes"`
		StatID      string `json:"statId"`
		StatName    string `json:"statName"`
		UnitLabel   string `json:"unitLabel"`
		UnitType    int    `json:"unitType"`
		Weight      int    `json:"weight"`
	} `json:"weaponKillsSubmachinegun"`
	WeaponKillsSuper struct {
		Category    int    `json:"category"`
		Group       int    `json:"group"`
		Modes       []int  `json:"modes"`
		PeriodTypes []int  `json:"periodTypes"`
		StatID      string `json:"statId"`
		StatName    string `json:"statName"`
		UnitLabel   string `json:"unitLabel"`
		UnitType    int    `json:"unitType"`
		Weight      int    `json:"weight"`
	} `json:"weaponKillsSuper"`
	WeaponKillsSword struct {
		Category    int    `json:"category"`
		Group       int    `json:"group"`
		Modes       []int  `json:"modes"`
		PeriodTypes []int  `json:"periodTypes"`
		StatID      string `json:"statId"`
		StatName    string `json:"statName"`
		UnitLabel   string `json:"unitLabel"`
		UnitType    int    `json:"unitType"`
		Weight      int    `json:"weight"`
	} `json:"weaponKillsSword"`
	WeaponPrecisionKillsAutoRifle struct {
		Category    int    `json:"category"`
		Group       int    `json:"group"`
		Modes       []int  `json:"modes"`
		PeriodTypes []int  `json:"periodTypes"`
		StatID      string `json:"statId"`
		StatName    string `json:"statName"`
		UnitLabel   string `json:"unitLabel"`
		UnitType    int    `json:"unitType"`
		Weight      int    `json:"weight"`
	} `json:"weaponPrecisionKillsAutoRifle"`
	WeaponPrecisionKillsFusionRifle struct {
		Category    int    `json:"category"`
		Group       int    `json:"group"`
		Modes       []int  `json:"modes"`
		PeriodTypes []int  `json:"periodTypes"`
		StatID      string `json:"statId"`
		StatName    string `json:"statName"`
		UnitLabel   string `json:"unitLabel"`
		UnitType    int    `json:"unitType"`
		Weight      int    `json:"weight"`
	} `json:"weaponPrecisionKillsFusionRifle"`
	WeaponPrecisionKillsGrenade struct {
		Category    int    `json:"category"`
		Group       int    `json:"group"`
		Modes       []int  `json:"modes"`
		PeriodTypes []int  `json:"periodTypes"`
		StatID      string `json:"statId"`
		StatName    string `json:"statName"`
		UnitLabel   string `json:"unitLabel"`
		UnitType    int    `json:"unitType"`
		Weight      int    `json:"weight"`
	} `json:"weaponPrecisionKillsGrenade"`
	WeaponPrecisionKillsHandCannon struct {
		Category    int    `json:"category"`
		Group       int    `json:"group"`
		Modes       []int  `json:"modes"`
		PeriodTypes []int  `json:"periodTypes"`
		StatID      string `json:"statId"`
		StatName    string `json:"statName"`
		UnitLabel   string `json:"unitLabel"`
		UnitType    int    `json:"unitType"`
		Weight      int    `json:"weight"`
	} `json:"weaponPrecisionKillsHandCannon"`
	WeaponPrecisionKillsMachinegun struct {
		Category    int    `json:"category"`
		Group       int    `json:"group"`
		Modes       []int  `json:"modes"`
		PeriodTypes []int  `json:"periodTypes"`
		StatID      string `json:"statId"`
		StatName    string `json:"statName"`
		UnitLabel   string `json:"unitLabel"`
		UnitType    int    `json:"unitType"`
		Weight      int    `json:"weight"`
	} `json:"weaponPrecisionKillsMachinegun"`
	WeaponPrecisionKillsMelee struct {
		Category    int    `json:"category"`
		Group       int    `json:"group"`
		Modes       []int  `json:"modes"`
		PeriodTypes []int  `json:"periodTypes"`
		StatID      string `json:"statId"`
		StatName    string `json:"statName"`
		UnitLabel   string `json:"unitLabel"`
		UnitType    int    `json:"unitType"`
		Weight      int    `json:"weight"`
	} `json:"weaponPrecisionKillsMelee"`
	WeaponPrecisionKillsPulseRifle struct {
		Category    int    `json:"category"`
		Group       int    `json:"group"`
		Modes       []int  `json:"modes"`
		PeriodTypes []int  `json:"periodTypes"`
		StatID      string `json:"statId"`
		StatName    string `json:"statName"`
		UnitLabel   string `json:"unitLabel"`
		UnitType    int    `json:"unitType"`
		Weight      int    `json:"weight"`
	} `json:"weaponPrecisionKillsPulseRifle"`
	WeaponPrecisionKillsRelic struct {
		Category    int    `json:"category"`
		Group       int    `json:"group"`
		Modes       []int  `json:"modes"`
		PeriodTypes []int  `json:"periodTypes"`
		StatID      string `json:"statId"`
		StatName    string `json:"statName"`
		UnitLabel   string `json:"unitLabel"`
		UnitType    int    `json:"unitType"`
		Weight      int    `json:"weight"`
	} `json:"weaponPrecisionKillsRelic"`
	WeaponPrecisionKillsRocketLauncher struct {
		Category    int    `json:"category"`
		Group       int    `json:"group"`
		Modes       []int  `json:"modes"`
		PeriodTypes []int  `json:"periodTypes"`
		StatID      string `json:"statId"`
		StatName    string `json:"statName"`
		UnitLabel   string `json:"unitLabel"`
		UnitType    int    `json:"unitType"`
		Weight      int    `json:"weight"`
	} `json:"weaponPrecisionKillsRocketLauncher"`
	WeaponPrecisionKillsScoutRifle struct {
		Category    int    `json:"category"`
		Group       int    `json:"group"`
		Modes       []int  `json:"modes"`
		PeriodTypes []int  `json:"periodTypes"`
		StatID      string `json:"statId"`
		StatName    string `json:"statName"`
		UnitLabel   string `json:"unitLabel"`
		UnitType    int    `json:"unitType"`
		Weight      int    `json:"weight"`
	} `json:"weaponPrecisionKillsScoutRifle"`
	WeaponPrecisionKillsShotgun struct {
		Category    int    `json:"category"`
		Group       int    `json:"group"`
		Modes       []int  `json:"modes"`
		PeriodTypes []int  `json:"periodTypes"`
		StatID      string `json:"statId"`
		StatName    string `json:"statName"`
		UnitLabel   string `json:"unitLabel"`
		UnitType    int    `json:"unitType"`
		Weight      int    `json:"weight"`
	} `json:"weaponPrecisionKillsShotgun"`
	WeaponPrecisionKillsSideArm struct {
		Category    int    `json:"category"`
		Group       int    `json:"group"`
		Modes       []int  `json:"modes"`
		PeriodTypes []int  `json:"periodTypes"`
		StatID      string `json:"statId"`
		StatName    string `json:"statName"`
		UnitLabel   string `json:"unitLabel"`
		UnitType    int    `json:"unitType"`
		Weight      int    `json:"weight"`
	} `json:"weaponPrecisionKillsSideArm"`
	WeaponPrecisionKillsSniper struct {
		Category    int    `json:"category"`
		Group       int    `json:"group"`
		Modes       []int  `json:"modes"`
		PeriodTypes []int  `json:"periodTypes"`
		StatID      string `json:"statId"`
		StatName    string `json:"statName"`
		UnitLabel   string `json:"unitLabel"`
		UnitType    int    `json:"unitType"`
		Weight      int    `json:"weight"`
	} `json:"weaponPrecisionKillsSniper"`
	WeaponPrecisionKillsSubmachinegun struct {
		Category    int    `json:"category"`
		Group       int    `json:"group"`
		Modes       []int  `json:"modes"`
		PeriodTypes []int  `json:"periodTypes"`
		StatID      string `json:"statId"`
		StatName    string `json:"statName"`
		UnitLabel   string `json:"unitLabel"`
		UnitType    int    `json:"unitType"`
		Weight      int    `json:"weight"`
	} `json:"weaponPrecisionKillsSubmachinegun"`
	WeaponPrecisionKillsSuper struct {
		Category    int    `json:"category"`
		Group       int    `json:"group"`
		Modes       []int  `json:"modes"`
		PeriodTypes []int  `json:"periodTypes"`
		StatID      string `json:"statId"`
		StatName    string `json:"statName"`
		UnitLabel   string `json:"unitLabel"`
		UnitType    int    `json:"unitType"`
		Weight      int    `json:"weight"`
	} `json:"weaponPrecisionKillsSuper"`
	WinLossRatio struct {
		Category    int    `json:"category"`
		Group       int    `json:"group"`
		Modes       []int  `json:"modes"`
		PeriodTypes []int  `json:"periodTypes"`
		StatID      string `json:"statId"`
		StatName    string `json:"statName"`
		UnitLabel   string `json:"unitLabel"`
		UnitType    int    `json:"unitType"`
		Weight      int    `json:"weight"`
	} `json:"winLossRatio"`
}

type ClanWeeklyRewardState struct {
	EndDate       string `json:"endDate"`
	MilestoneHash int    `json:"milestoneHash"`
	Rewards       []struct {
		Entries []struct {
			Earned          bool `json:"earned"`
			Redeemed        bool `json:"redeemed"`
			RewardEntryHash int  `json:"rewardEntryHash"`
		} `json:"entries"`
		RewardCategoryHash int `json:"rewardCategoryHash"`
	} `json:"rewards"`
	StartDate string `json:"startDate"`
}
