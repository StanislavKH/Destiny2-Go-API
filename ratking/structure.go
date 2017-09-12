package ratking

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
