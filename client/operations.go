package client

type operation interface {
	Process(state *albionState)
}

//go:generate stringer -type=OperationType
type OperationType uint16

const (
	opUnused                                  = 0
	opPing                                    = 1
	opJoin                                    = 2
	opCreateAccount                           = 3
	opLogin                                   = 4
	opSendCrashLog                            = 5
	opSendGamePingInfo                        = 6
	opCreateCharacter                         = 7
	opDeleteCharacter                         = 8
	opSelectCharacter                         = 9
	opRedeemKeycode                           = 10
	opGetGameServerByCluster                  = 11
	opGetActiveSubscription                   = 12
	opGetShopPurchaseUrl                      = 13
	opGetBuyTrialDetails                      = 14
	opGetReferralSeasonDetails                = 15
	opGetReferralLink                         = 16
	opGetAvailableTrialKeys                   = 17
	opGetShopTilesForCategory                 = 18
	opMove                                    = 19
	opAttackStart                             = 20
	opCastStart                               = 21
	opCastCancel                              = 22
	opTerminateToggleSpell                    = 23
	opChannelingCancel                        = 24
	opAttackBuildingStart                     = 25
	opInventoryDestroyItem                    = 26
	opInventoryMoveItem                       = 27
	opInventorySplitStack                     = 28
	opGetClusterData                          = 29
	opChangeCluster                           = 30
	opConsoleCommand                          = 31
	opChatMessage                             = 32
	opReportClientError                       = 33
	opRegisterToObject                        = 34
	opUnRegisterFromObject                    = 35
	opCraftBuildingChangeSettings             = 36
	opCraftBuildingTakeMoney                  = 37
	opRepairBuildingChangeSettings            = 38
	opRepairBuildingTakeMoney                 = 39
	opActionBuildingChangeSettings            = 40
	opHarvestStart                            = 41
	opHarvestCancel                           = 42
	opTakeSilver                              = 43
	opActionOnBuildingStart                   = 44
	opActionOnBuildingCancel                  = 45
	opItemRerollQualityStart                  = 46
	opItemRerollQualityCancel                 = 47
	opInstallResourceStart                    = 48
	opInstallResourceCancel                   = 49
	opInstallSilver                           = 50
	opBuildingFillNutrition                   = 51
	opBuildingChangeRenovationState           = 52
	opBuildingBuySkin                         = 53
	opBuildingClaim                           = 54
	opBuildingGiveup                          = 55
	opBuildingNutritionSilverStorageDeposit   = 56
	opBuildingNutritionSilverStorageWithdraw  = 57
	opBuildingNutritionSilverRewardSet        = 58
	opConstructionSiteCreate                  = 59
	opPlaceableItemPlace                      = 60
	opPlaceableItemPlaceCancel                = 61
	opPlaceableObjectPickup                   = 62
	opFurnitureObjectUse                      = 63
	opFarmableHarvest                         = 64
	opFarmableFinishGrownItem                 = 65
	opFarmableDestroy                         = 66
	opFarmableGetProduct                      = 67
	opFarmableFill                            = 68
	opLaborerObjectPlace                      = 69
	opLaborerObjectPlaceCancel                = 70
	opTearDownConstructionSite                = 71
	opCastleGateUse                           = 72
	opAuctionCreateOffer                      = 73
	opAuctionCreateRequest                    = 74
	opAuctionGetOffers                        = 75
	opAuctionGetRequests                      = 76
	opAuctionBuyOffer                         = 77
	opAuctionAbortAuction                     = 78
	opAuctionModifyAuction                    = 79
	opAuctionAbortOffer                       = 80
	opAuctionAbortRequest                     = 81
	opAuctionSellRequest                      = 82
	opAuctionGetFinishedAuctions              = 83
	opAuctionFetchAuction                     = 84
	opAuctionGetMyOpenOffers                  = 85
	opAuctionGetMyOpenRequests                = 86
	opAuctionGetMyOpenAuctions                = 87
	opAuctionGetItemsAverage                  = 88
	opAuctionGetItemAverageStats              = 89
	opAuctionGetItemAverageValue              = 90
	opContainerOpen                           = 91
	opContainerClose                          = 92
	opContainerManageSubContainer             = 93
	opRespawn                                 = 94
	opSuicide                                 = 95
	opJoinGuild                               = 96
	opLeaveGuild                              = 97
	opCreateGuild                             = 98
	opInviteToGuild                           = 99
	opDeclineGuildInvitation                  = 100
	opKickFromGuild                           = 101
	opDuellingChallengePlayer                 = 102
	opDuellingAcceptChallenge                 = 103
	opDuellingDenyChallenge                   = 104
	opChangeClusterTax                        = 105
	opClaimTerritory                          = 106
	opGiveUpTerritory                         = 107
	opChangeTerritoryAccessRights             = 108
	opGetMonolithInfo                         = 109
	opGetClaimInfo                            = 110
	opGetAttackInfo                           = 111
	opGetTerritorySeasonPoints                = 112
	opGetAttackSchedule                       = 113
	opScheduleAttack                          = 114
	opGetMatches                              = 115
	opGetMatchDetails                         = 116
	opJoinMatch                               = 117
	opLeaveMatch                              = 118
	opChangeChatSettings                      = 119
	opLogoutStart                             = 120
	opLogoutCancel                            = 121
	opClaimOrbStart                           = 122
	opClaimOrbCancel                          = 123
	opMatchLootChestOpeningStart              = 124
	opMatchLootChestOpeningCancel             = 125
	opDepositToGuildAccount                   = 126
	opWithdrawalFromAccount                   = 127
	opChangeGuildPayUpkeepFlag                = 128
	opChangeGuildTax                          = 129
	opGetMyTerritories                        = 130
	opMorganaCommand                          = 131
	opGetServerInfo                           = 132
	opInviteMercenaryToMatch                  = 133
	opSubscribeToCluster                      = 134
	opAnswerMercenaryInvitation               = 135
	opGetCharacterEquipment                   = 136
	opGetCharacterSteamAchievements           = 137
	opGetCharacterStats                       = 138
	opGetKillHistoryDetails                   = 139
	opLearnMasteryLevel                       = 140
	opReSpecAchievement                       = 141
	opChangeAvatar                            = 142
	opGetRankings                             = 143
	opGetRank                                 = 144
	opGetGvgSeasonRankings                    = 145
	opGetGvgSeasonRank                        = 146
	opGetGvgSeasonHistoryRankings             = 147
	opGetGvgSeasonGuildMemberHistory          = 148
	opKickFromGvGMatch                        = 149
	opGetChestLogs                            = 150
	opGetAccessRightLogs                      = 151
	opGetGuildAccountLogs                     = 152
	opGetGuildAccountLogsLargeAmount          = 153
	opInviteToPlayerTrade                     = 154
	opPlayerTradeCancel                       = 155
	opPlayerTradeInvitationAccept             = 156
	opPlayerTradeAddItem                      = 157
	opPlayerTradeRemoveItem                   = 158
	opPlayerTradeAcceptTrade                  = 159
	opPlayerTradeSetSilverOrGold              = 160
	opSendMiniMapPing                         = 161
	opStuck                                   = 162
	opBuyRealEstate                           = 163
	opClaimRealEstate                         = 164
	opGiveUpRealEstate                        = 165
	opChangeRealEstateOutline                 = 166
	opGetMailInfos                            = 167
	opReadMail                                = 168
	opSendNewMail                             = 169
	opDeleteMail                              = 170
	opClaimAttachmentFromMail                 = 171
	opUpdateLfgInfo                           = 172
	opGetLfgInfos                             = 173
	opGetMyGuildLfgInfo                       = 174
	opGetLfgDescriptionText                   = 175
	opLfgApplyToGuild                         = 176
	opAnswerLfgGuildApplication               = 177
	opGetClusterInfo                          = 178
	opRegisterChatPeer                        = 179
	opSendChatMessage                         = 180
	opJoinChatChannel                         = 181
	opLeaveChatChannel                        = 182
	opSendWhisperMessage                      = 183
	opSay                                     = 184
	opPlayEmote                               = 185
	opStopEmote                               = 186
	opGetClusterMapInfo                       = 187
	opAccessRightsChangeSettings              = 188
	opMount                                   = 189
	opMountCancel                             = 190
	opBuyJourney                              = 191
	opSetSaleStatusForEstate                  = 192
	opResolveGuildOrPlayerName                = 193
	opGetRespawnInfos                         = 194
	opMakeHome                                = 195
	opLeaveHome                               = 196
	opResurrectionReply                       = 197
	opAllianceCreate                          = 198
	opAllianceDisband                         = 199
	opAllianceGetMemberInfos                  = 200
	opAllianceInvite                          = 201
	opAllianceAnswerInvitation                = 202
	opAllianceCancelInvitation                = 203
	opAllianceKickGuild                       = 204
	opAllianceLeave                           = 205
	opAllianceChangeGoldPaymentFlag           = 206
	opAllianceGetDetailInfo                   = 207
	opGetIslandInfos                          = 208
	opAbandonMyIsland                         = 209
	opBuyMyIsland                             = 210
	opBuyGuildIsland                          = 211
	opAbandonGuildIsland                      = 212
	opUpgradeMyIsland                         = 213
	opUpgradeGuildIsland                      = 214
	opMoveMyIsland                            = 215
	opMoveGuildIsland                         = 216
	opTerritoryFillNutrition                  = 217
	opTeleportBack                            = 218
	opPartyInvitePlayer                       = 219
	opPartyAnswerInvitation                   = 220
	opPartyLeave                              = 221
	opPartyKickPlayer                         = 222
	opPartyMakeLeader                         = 223
	opPartyChangeLootSetting                  = 224
	opPartyMarkObject                         = 225
	opPartySetRole                            = 226
	opGetGuildMOTD                            = 227
	opSetGuildMOTD                            = 228
	opExitEnterStart                          = 229
	opExitEnterCancel                         = 230
	opQuestGiverRequest                       = 231
	opGoldMarketGetBuyOffer                   = 232
	opGoldMarketGetBuyOfferFromSilver         = 233
	opGoldMarketGetSellOffer                  = 234
	opGoldMarketGetSellOfferFromSilver        = 235
	opGoldMarketBuyGold                       = 236
	opGoldMarketSellGold                      = 237
	opGoldMarketCreateSellOrder               = 238
	opGoldMarketCreateBuyOrder                = 239
	opGoldMarketGetInfos                      = 240
	opGoldMarketCancelOrder                   = 241
	opGoldMarketGetAverageInfo                = 242
	opSiegeCampClaimStart                     = 243
	opSiegeCampClaimCancel                    = 244
	opTreasureChestUsingStart                 = 245
	opTreasureChestUsingCancel                = 246
	opUseLootChest                            = 247
	opUseShrine                               = 248
	opLaborerStartJob                         = 249
	opLaborerTakeJobLoot                      = 250
	opLaborerDismiss                          = 251
	opLaborerMove                             = 252
	opLaborerBuyItem                          = 253
	opLaborerUpgrade                          = 254
	opBuyPremium                              = 255
	opBuyTrial                                = 256
	opRealEstateGetAuctionData                = 257
	opRealEstateBidOnAuction                  = 258
	opGetSiegeCampCooldown                    = 259
	opFriendInvite                            = 260
	opFriendAnswerInvitation                  = 261
	opFriendCancelnvitation                   = 262
	opFriendRemove                            = 263
	opInventoryStack                          = 264
	opInventorySort                           = 265
	opEquipmentItemChangeSpell                = 266
	opExpeditionRegister                      = 267
	opExpeditionRegisterCancel                = 268
	opJoinExpedition                          = 269
	opDeclineExpeditionInvitation             = 270
	opVoteStart                               = 271
	opVoteDoVote                              = 272
	opRatingDoRate                            = 273
	opEnteringExpeditionStart                 = 274
	opEnteringExpeditionCancel                = 275
	opActivateExpeditionCheckPoint            = 276
	opArenaRegister                           = 277
	opArenaRegisterCancel                     = 278
	opArenaLeave                              = 279
	opJoinArenaMatch                          = 280
	opDeclineArenaInvitation                  = 281
	opEnteringArenaStart                      = 282
	opEnteringArenaCancel                     = 283
	opArenaCustomMatch                        = 284
	opArenaCustomMatchCreate                  = 285
	opUpdateCharacterStatement                = 286
	opBoostFarmable                           = 287
	opGetStrikeHistory                        = 288
	opUseFunction                             = 289
	opUsePortalEntrance                       = 290
	opResetPortalBinding                      = 291
	opQueryPortalBinding                      = 292
	opClaimPaymentTransaction                 = 293
	opChangeUseFlag                           = 294
	opClientPerformanceStats                  = 295
	opExtendedHardwareStats                   = 296
	opClientLowMemoryWarning                  = 297
	opTerritoryClaimStart                     = 298
	opTerritoryClaimCancel                    = 299
	opRequestAppStoreProducts                 = 300
	opVerifyProductPurchase                   = 301
	opQueryGuildPlayerStats                   = 302
	opTrackAchievements                       = 303
	opSetAchievementsAutoLearn                = 304
	opDepositItemToGuildCurrency              = 305
	opWithdrawalItemFromGuildCurrency         = 306
	opAuctionSellSpecificItemRequest          = 307
	opFishingStart                            = 308
	opFishingCasting                          = 309
	opFishingCast                             = 310
	opFishingCatch                            = 311
	opFishingPull                             = 312
	opFishingGiveLine                         = 313
	opFishingFinish                           = 314
	opFishingCancel                           = 315
	opCreateGuildAccessTag                    = 316
	opDeleteGuildAccessTag                    = 317
	opRenameGuildAccessTag                    = 318
	opFlagGuildAccessTagGuildPermission       = 319
	opAssignGuildAccessTag                    = 320
	opRemoveGuildAccessTagFromPlayer          = 321
	opModifyGuildAccessTagEditors             = 322
	opRequestPublicAccessTags                 = 323
	opChangeAccessTagPublicFlag               = 324
	opUpdateGuildAccessTag                    = 325
	opSteamStartMicrotransaction              = 326
	opSteamFinishMicrotransaction             = 327
	opSteamIdHasActiveAccount                 = 328
	opCheckEmailAccountState                  = 329
	opLinkAccountToSteamId                    = 330
	opBuyGvgSeasonBooster                     = 331
	opChangeFlaggingPrepare                   = 332
	opOverCharge                              = 333
	opOverChargeEnd                           = 334
	opRequestTrusted                          = 335
	opChangeGuildLogo                         = 336
	opPartyFinderRegisterForUpdates           = 337
	opPartyFinderUnregisterForUpdates         = 338
	opPartyFinderEnlistNewPartySearch         = 339
	opPartyFinderDeletePartySearch            = 340
	opPartyFinderChangePartySearch            = 341
	opPartyFinderChangeRole                   = 342
	opPartyFinderApplyForGroup                = 343
	opPartyFinderAcceptOrDeclineApplyForGroup = 344
	opPartyFinderGetEquipmentSnapshot         = 345
	opPartyFinderRegisterApplicants           = 346
	opPartyFinderUnregisterApplicants         = 347
	opPartyFinderFulltextSearch               = 348
	opPartyFinderRequestEquipmentSnapshot     = 349
	opGetPersonalSeasonTrackerData            = 350
	opUseConsumableFromInventory              = 351
	opClaimPersonalSeasonReward               = 352
	opEasyAntiCheatMessageToServer            = 353
	opRetaliationAttackClaimTerritory         = 354
	opSetNextTutorialState                    = 355
	opAddPlayerToMuteList                     = 356
	opRemovePlayerFromMuteList                = 357
	opMakeTerritoryHome                       = 358
	opLeaveTerritoryHome                      = 359
	opProductShopUserEvent                    = 360
	opGetVanityUnlocks                        = 361
	opBuyVanityUnlock                         = 362
	opGetMountSkins                           = 363
	opSetMountSkin                            = 364
	opChangeCustomization                     = 365
	opSetFavoriteIsland                       = 366
	opGetGuildChallengePoints                 = 367
)
