package main

type SSLCode uint32

const (
	SSLMessageToDesk                           SSLCode = 0x00000001
	SSLMessageToRemote                         SSLCode = 0x00000002
	SSLMessageGetDesk                          SSLCode = 0x00000005
	SSLMessageGetDeskReply                     SSLCode = 0x00000006
	SSLMessageSendHeartbeat                    SSLCode = 0x00000007
	SSLMessageSetOwnName                       SSLCode = 0x00000008
	SSLMessageSetOwnNameReply                  SSLCode = 0x00000009
	SSLMessageGetProjectNameAndTitle           SSLCode = 0x0000000a
	SSLMessageGetProjectNameAndTitleReply      SSLCode = 0x0000000b
	SSLMessageGetSync                          SSLCode = 0x0000000f
	SSLMessageXpatchIphoneGetDesk              SSLCode = 0x00000010
	SSLMessageGetChanNamesAndImages            SSLCode = 0x00000014
	SSLMessageGetChanNamesAndImagesReply       SSLCode = 0x00000015
	SSLMessageSetDefaultChanNames              SSLCode = 0x0000001c
	SSLMessageSetDefaultChanNamesReply         SSLCode = 0x0000001d
	SSLMessageSetChanNames                     SSLCode = 0x0000001e
	SSLMessageSetChanNamesReply                SSLCode = 0x00000020
	SSLMessageSetChanImages                    SSLCode = 0x0000001f
	SSLMessageSetChanImagesReply               SSLCode = 0x00000021
	SSLMessageGetIsChanStereo                  SSLCode = 0x00000022
	SSLMessageGetIsChanStereoReply             SSLCode = 0x00000023
	SSLMessageGetImageLibNames                 SSLCode = 0x00000028
	SSLMessageGetImageLibNamesReply            SSLCode = 0x00000029
	SSLMessageGetExtNames                      SSLCode = 0x00000032
	SSLMessageGetExtNamesReply                 SSLCode = 0x00000033
	SSLMessageSetExtNames                      SSLCode = 0x00000034
	SSLMessageSetExtNamesReply                 SSLCode = 0x00000035
	SSLMessageGetDirectoryList                 SSLCode = 0x0000003c
	SSLMessageGetDirectoryListReply            SSLCode = 0x0000003d
	SSLMessageGetMixPassesList                 SSLCode = 0x0000003e
	SSLMessageGetMixPassesListReply            SSLCode = 0x0000003f
	SSLMessageGetTrList                        SSLCode = 0x00000040
	SSLMessageGetTrListReply                   SSLCode = 0x00000041
	SSLMessageSetFileInfo                      SSLCode = 0x00000046
	SSLMessageSetFileInfoReply                 SSLCode = 0x00000047
	SSLMessageSendDiskInfo                     SSLCode = 0x00000048
	SSLMessageAckSendDiskInfo                  SSLCode = 0x00000049
	SSLMessageGetDiskInfo                      SSLCode = 0x0000004a
	SSLMessageAckGetDiskInfo                   SSLCode = 0x0000004b
	SSLMessageSendUpdateBlock                  SSLCode = 0x00000050
	SSLMessageAckUpdateBlock                   SSLCode = 0x00000051
	SSLMessageSendSaveToFlash                  SSLCode = 0x0000005a
	SSLMessageAckSaveToFlash                   SSLCode = 0x0000005b
	SSLMessageSendCreateZip                    SSLCode = 0x00000064
	SSLMessageAckCreateZip                     SSLCode = 0x00000065
	SSLMessageSendRequestFileBlock             SSLCode = 0x0000006e
	SSLMessageAckRequestFileBlock              SSLCode = 0x0000006f
	SSLMessageSendWriteFileBlock               SSLCode = 0x00000078
	SSLMessageAckWriteFileBlock                SSLCode = 0x00000079
	SSLMessageSendUnzipFile                    SSLCode = 0x00000082
	SSLMessageAckUnzipFile                     SSLCode = 0x00000083
	SSLMessageSendMoveFile                     SSLCode = 0x0000008c
	SSLMessageAckMoveFile                      SSLCode = 0x0000008d
	SSLMessageSendCopyFile                     SSLCode = 0x00000096
	SSLMessageAckCopyFile                      SSLCode = 0x00000097
	SSLMessageSendDeleteFile                   SSLCode = 0x000000a0
	SSLMessageAckDeleteFile                    SSLCode = 0x000000a1
	SSLMessageSendRtc                          SSLCode = 0x000000aa
	SSLMessageAckRtc                           SSLCode = 0x000000ab
	SSLMessageSendSmartCopy                    SSLCode = 0x000000b4
	SSLMessageAckSmartCopy                     SSLCode = 0x000000b5
	SSLMessageSendSmartCopyWithName            SSLCode = 0x000000b6
	SSLMessageAckSmartCopyWithName             SSLCode = 0x000000b7
	SSLMessageSendTitleDetailsChanged          SSLCode = 0x000000be
	SSLMessageAckTitleDetailsChanged           SSLCode = 0x000000bf
	SSLMessageSendMakeNewProject               SSLCode = 0x000000c8
	SSLMessageAckMakeNewProject                SSLCode = 0x000000c9
	SSLMessageSendMakeNewProjectTitle          SSLCode = 0x000000d2
	SSLMessageAckMakeNewProjectTitle           SSLCode = 0x000000d3
	SSLMessageSendMakeNewProjectTitleWithName  SSLCode = 0x000000d4
	SSLMessageAckMakeNewProjectTitleWithName   SSLCode = 0x000000d5
	SSLMessageSendSelectProjectTitle           SSLCode = 0x000000dc
	SSLMessageAckSelectProjectTitle            SSLCode = 0x000000dd
	SSLMessageSendDeleteProjectTitle           SSLCode = 0x000000e6
	SSLMessageAckDeleteProjectTitle            SSLCode = 0x000000e7
	SSLMessageSendDeleteProject                SSLCode = 0x000000f0
	SSLMessageAckDeleteProject                 SSLCode = 0x000000f1
	SSLMessageSendCopyProjectTitle             SSLCode = 0x000000fa
	SSLMessageAckCopyProjectTitle              SSLCode = 0x000000fb
	SSLMessageSendCopyProject                  SSLCode = 0x00000104
	SSLMessageAckCopyProject                   SSLCode = 0x00000105
	SSLMessageSendArchiveDone                  SSLCode = 0x00000106
	SSLMessageAckArchiveDone                   SSLCode = 0x00000107
	SSLMessageSendMakeNewProjectWithName       SSLCode = 0x00000108
	SSLMessageAckMakeNewProjectWithName        SSLCode = 0x00000109
	SSLMessageSendMakeNewProjectWithPresetOpts SSLCode = 0x0000010a
	SSLMessageAckMakeNewProjectWithPresetOpts  SSLCode = 0x0000010b
	SSLMessageSendSetTrEnable                  SSLCode = 0x0000012c
	SSLMessageAckSetTrEnable                   SSLCode = 0x0000012d
	SSLMessageSendGetTrState                   SSLCode = 0x0000012e
	SSLMessageAckGetTrState                    SSLCode = 0x0000012f
	SSLMessageSendTakeTrSnap                   SSLCode = 0x00000136
	SSLMessageAckTakeTrSnap                    SSLCode = 0x00000137
	SSLMessageSendSelectTrSnap                 SSLCode = 0x00000140
	SSLMessageAckSelectTrSnap                  SSLCode = 0x00000141
	SSLMessageSendDeleteTrSnap                 SSLCode = 0x0000014a
	SSLMessageAckDeleteTrSnap                  SSLCode = 0x0000014b
	SSLMessageSendCopyTrChanData               SSLCode = 0x00000154
	SSLMessageAckCopyTrChanData                SSLCode = 0x00000155
	SSLMessageSendSwopTrChanData               SSLCode = 0x0000015e
	SSLMessageAckSwopTrChanData                SSLCode = 0x0000015f
	SSLMessageSendSetTrAllChans                SSLCode = 0x00000168
	SSLMessageAckSetTrAllChans                 SSLCode = 0x00000169
	SSLMessageSendSetTrChan                    SSLCode = 0x00000172
	SSLMessageAckSetTrChan                     SSLCode = 0x00000173
	SSLMessageSendGetInsertNames               SSLCode = 0x00000190
	SSLMessageAckGetInsertNames                SSLCode = 0x00000191
	SSLMessageSendSetInsertNames               SSLCode = 0x0000019a
	SSLMessageAckSetInsertNames                SSLCode = 0x0000019b
	SSLMessageSendGetChainNames                SSLCode = 0x000001a4
	SSLMessageAckGetChainNames                 SSLCode = 0x000001a5
	SSLMessageSendSetInsertToChan              SSLCode = 0x000001ae
	SSLMessageAckSetInsertToChan               SSLCode = 0x000001af
	SSLMessageSendGetChanMatrixInfo            SSLCode = 0x000001b8
	SSLMessageAckGetChanMatrixInfo             SSLCode = 0x000001b9
	SSLMessageSendGetNumChains                 SSLCode = 0x000001c2
	SSLMessageAckGetNumChains                  SSLCode = 0x000001c3
	SSLMessageSendNewEditChain                 SSLCode = 0x000001cc
	SSLMessageAckNewEditChain                  SSLCode = 0x000001cd
	SSLMessageSendSaveEditChain                SSLCode = 0x000001d6
	SSLMessageAckSaveEditChain                 SSLCode = 0x000001d7
	SSLMessageSendGetEditChainData             SSLCode = 0x000001e0
	SSLMessageAckGetEditChainData              SSLCode = 0x000001e1
	SSLMessageSendAddDevToEditChain            SSLCode = 0x000001ea
	SSLMessageAckAddDevToEditChain             SSLCode = 0x000001eb
	SSLMessageSendCancelEditChain              SSLCode = 0x000001f4
	SSLMessageAckCancelEditChain               SSLCode = 0x000001f5
	SSLMessageSendAssignChainToChan            SSLCode = 0x000001fe
	SSLMessageAckAssignChainToChan             SSLCode = 0x000001ff
	SSLMessageSendDeassignChan                 SSLCode = 0x00000208
	SSLMessageAckDeassignChan                  SSLCode = 0x00000209
	SSLMessageSendRemFromEditChain             SSLCode = 0x00000212
	SSLMessageAckRemFromEditChain              SSLCode = 0x00000213
	SSLMessageSendEditExistingChain            SSLCode = 0x0000021c
	SSLMessageAckEditExistingChain             SSLCode = 0x0000021d
	SSLMessageSendDeleteChain                  SSLCode = 0x00000226
	SSLMessageAckDeleteChain                   SSLCode = 0x00000227
	SSLMessageSendRenameEditChain              SSLCode = 0x00000230
	SSLMessageAckRenameEditChain               SSLCode = 0x00000231
	SSLMessageSendInsDevToEditChain            SSLCode = 0x0000023a
	SSLMessageAckInsDevToEditChain             SSLCode = 0x0000023b
	SSLMessageSendGetValidEditChainDevs        SSLCode = 0x00000244
	SSLMessageAckGetValidEditChainDevs         SSLCode = 0x00000245
	SSLMessageSendGetInsertInfoV2              SSLCode = 0x000028a0
	SSLMessageAckGetInsertInfoV2               SSLCode = 0x000028a1
	SSLMessageSendSetInsertNamesV2             SSLCode = 0x000028aa
	SSLMessageAckSetInsertNamesV2              SSLCode = 0x000028ab
	SSLMessageSendGetChainInfoV2               SSLCode = 0x000028b4
	SSLMessageAckGetChainInfoV2                SSLCode = 0x000028b5
	SSLMessageSendSetInsertToChanV2            SSLCode = 0x000028be
	SSLMessageAckSetInsertToChanV2             SSLCode = 0x000028bf
	SSLMessageSendGetChanMatrixInfoV2          SSLCode = 0x000028c8
	SSLMessageAckGetChanMatrixInfoV2           SSLCode = 0x000028c9
	SSLMessageSendGetNumChainsV2               SSLCode = 0x000028d2
	SSLMessageAckGetNumChainsV2                SSLCode = 0x000028d3
	SSLMessageSendAssignChainToChanV2          SSLCode = 0x0000290e
	SSLMessageAckAssignChainToChanV2           SSLCode = 0x0000290f
	SSLMessageSendDeassignChanV2               SSLCode = 0x00002918
	SSLMessageAckDeassignChanV2                SSLCode = 0x00002919
	SSLMessageSendDeleteChainV2                SSLCode = 0x00002936
	SSLMessageAckDeleteChainV2                 SSLCode = 0x00002937
	SSLMessageSendRenameChain                  SSLCode = 0x00002940
	SSLMessageAckRenameChain                   SSLCode = 0x00002941
	SSLMessageSendSaveInsertsToChain           SSLCode = 0x0000294a
	SSLMessageAckSaveInsertsToChain            SSLCode = 0x0000294b
	SSLMessageSendSetChanInserts               SSLCode = 0x00002954
	SSLMessageAckSetChanInserts                SSLCode = 0x00002955
	SSLMessageSendDeleteChanInsert             SSLCode = 0x00002968
	SSLMessageAckDeleteChanInsert              SSLCode = 0x00002969
	SSLMessageSendReorderChanInserts           SSLCode = 0x00002972
	SSLMessageAckReorderChanInserts            SSLCode = 0x00002973
	SSLMessageSendSetChanStereoInsert          SSLCode = 0x0000297c
	SSLMessageAckSetChanStereoInsert           SSLCode = 0x0000297d
	SSLMessageSendGetMatrixPresetList          SSLCode = 0x00002986
	SSLMessageAckGetMatrixPresetList           SSLCode = 0x00002987
	SSLMessageSendLoadMatrixPreset             SSLCode = 0x00002990
	SSLMessageAckLoadMatrixPreset              SSLCode = 0x00002991
	SSLMessageSendSaveMatrixPreset             SSLCode = 0x0000299a
	SSLMessageAckSaveMatrixPreset              SSLCode = 0x0000299b
	SSLMessageSendDeleteMatrixPreset           SSLCode = 0x000029a4
	SSLMessageAckDeleteMatrixPreset            SSLCode = 0x000029a5
	SSLMessageSendRenameMatrixPreset           SSLCode = 0x000029ae
	SSLMessageAckRenameMatrixPreset            SSLCode = 0x000029af
	SSLMessageSendClearInserts                 SSLCode = 0x000029b8
	SSLMessageAckClearInserts                  SSLCode = 0x000029b9
	SSLMessageSendSetLowerScribMode            SSLCode = 0x000029d6
	SSLMessageAckSetLowerScribMode             SSLCode = 0x000029d7
	SSLMessageSendGetLowerScribMode            SSLCode = 0x000029e0
	SSLMessageAckGetLowerScribMode             SSLCode = 0x000029e1
	SSLMessageSendSetDisplay1732               SSLCode = 0x000029ea
	SSLMessageAckSetDisplay1732                SSLCode = 0x000029eb
	SSLMessageSendGetDisplay1732               SSLCode = 0x000029f4
	SSLMessageAckGetDisplay1732                SSLCode = 0x000029f5
	SSLMessageSendSetFlipScribStrip            SSLCode = 0x000029fe
	SSLMessageAckSetFlipScribStrip             SSLCode = 0x000029ff
	SSLMessageSendGetFlipScribStrip            SSLCode = 0x00002a08
	SSLMessageAckGetFlipScribStrip             SSLCode = 0x00002a09
	SSLMessageSendRenameChanNamesPreset        SSLCode = 0x00002a12
	SSLMessageAckRenameChanNamesPreset         SSLCode = 0x00002a13
	SSLMessageSendDeleteChanNamesPreset        SSLCode = 0x00002a1c
	SSLMessageAckDeleteChanNamesPreset         SSLCode = 0x00002a1d
	SSLMessageSendGetChanNamesPresetList       SSLCode = 0x00002a26
	SSLMessageAckGetChanNamesPresetList        SSLCode = 0x00002a27
	SSLMessageSendSaveChanNamesPreset          SSLCode = 0x00002a30
	SSLMessageAckSaveChanNamesPreset           SSLCode = 0x00002a31
	SSLMessageSendLoadChanNamesPreset          SSLCode = 0x00002a3a
	SSLMessageAckLoadChanNamesPreset           SSLCode = 0x00002a3b
	SSLMessageSendGetAutomationMode            SSLCode = 0x00002a94
	SSLMessageAckGetAutomationMode             SSLCode = 0x00002a95
	SSLMessageSendSetAutomationMode            SSLCode = 0x00002af8
	SSLMessageAckSetAutomationMode             SSLCode = 0x00002af9
	SSLMessageSendGetMotorsOffTouchEn          SSLCode = 0x00002b5c
	SSLMessageAckGetMotorsOffTouchEn           SSLCode = 0x00002b5d
	SSLMessageSendSetMotorsOffTouchEn          SSLCode = 0x00002bc0
	SSLMessageAckSetMotorsOffTouchEn           SSLCode = 0x00002bc1
	SSLMessageSendGetMdacMeterEn               SSLCode = 0x00002c24
	SSLMessageAckGetMdacMeterEn                SSLCode = 0x00002c25
	SSLMessageSendSetMdacMeterEn               SSLCode = 0x00002c88
	SSLMessageAckSetMdacMeterEn                SSLCode = 0x00002c89
	SSLMessageSendGetEditKeymapName            SSLCode = 0x00000258
	SSLMessageAckGetEditKeymapName             SSLCode = 0x00000259
	SSLMessageSendSetEditKeymapName            SSLCode = 0x00000262
	SSLMessageAckSetEditKeymapName             SSLCode = 0x00000263
	SSLMessageSendGetEditKeymapData            SSLCode = 0x0000026c
	SSLMessageAckGetEditKeymapData             SSLCode = 0x0000026d
	SSLMessageSendGetEditKeymapKeycaps         SSLCode = 0x00000276
	SSLMessageAckGetEditKeymapKeycap           SSLCode = 0x00000277
	SSLMessageSendGetEditKeymapSize            SSLCode = 0x00000280
	SSLMessageAckGetEditKeymapSize             SSLCode = 0x00000281
	SSLMessageSendSetUsbCmd                    SSLCode = 0x0000028a
	SSLMessageAckSetUsbCmd                     SSLCode = 0x0000028b
	SSLMessageSendSetKeycapName                SSLCode = 0x00000294
	SSLMessageAckSetKeycapName                 SSLCode = 0x00000295
	SSLMessageSendSetKeyBlank                  SSLCode = 0x0000029e
	SSLMessageAckSetKeyBlank                   SSLCode = 0x0000029f
	SSLMessageSendSetSaveEditKeymap            SSLCode = 0x000002a8
	SSLMessageAckSetSaveEditKeymap             SSLCode = 0x000002a9
	SSLMessageSendGetMidiFunctionList          SSLCode = 0x000002b2
	SSLMessageAckGetMidiFunctionList           SSLCode = 0x000002b3
	SSLMessageSendSetMidiCmd                   SSLCode = 0x000002bc
	SSLMessageAckSetMidiCmd                    SSLCode = 0x000002bd
	SSLMessageSendSetNewMenuCmd                SSLCode = 0x000002c6
	SSLMessageAckSetNewMenuCmd                 SSLCode = 0x000002c7
	SSLMessageSendSetMenuSubKeycapName         SSLCode = 0x000002d0
	SSLMessageAckSetMenuSubKeycapName          SSLCode = 0x000002d1
	SSLMessageSendSetMenuSubMidiCmd            SSLCode = 0x000002da
	SSLMessageAckSetMenuSubMidiCmd             SSLCode = 0x000002db
	SSLMessageSendSetMenuSubUsbCmd             SSLCode = 0x000002e4
	SSLMessageAckSetMenuSubUsbCmd              SSLCode = 0x000002e5
	SSLMessageSendSetMenuSubBlankCmd           SSLCode = 0x000002ee
	SSLMessageAckSetMenuSubBlankCmd            SSLCode = 0x000002ef
	SSLMessageSendRestartConsole               SSLCode = 0x000002f8
	SSLMessageSendFollowKeyState               SSLCode = 0x00000302
	SSLMessageAckFollowKeyState                SSLCode = 0x00000303
	SSLMessageSendGetProfileForDawLayer        SSLCode = 0x00000320
	SSLMessageAckGetProfileForDawLayer         SSLCode = 0x00000321
	SSLMessageSendCopyProfileToNew             SSLCode = 0x0000032a
	SSLMessageAckCopyProfileToNew              SSLCode = 0x0000032b
	SSLMessageSendSetProfileForDawLayer        SSLCode = 0x00000334
	SSLMessageAckSetProfileForDawLayer         SSLCode = 0x00000335
	SSLMessageSendClearProfileForDawLayer      SSLCode = 0x0000033e
	SSLMessageAckClearProfileForDawLayer       SSLCode = 0x0000033f
	SSLMessageSendGetProfiles                  SSLCode = 0x00000348
	SSLMessageAckGetProfiles                   SSLCode = 0x00000349
	SSLMessageSendRenameProfiles               SSLCode = 0x00000352
	SSLMessageAckRenameProfiles                SSLCode = 0x00000353
	SSLMessageSendDeleteProfiles               SSLCode = 0x0000035c
	SSLMessageAckDeleteProfiles                SSLCode = 0x0000035d
	SSLMessageSendGetTransportLockDawLayer     SSLCode = 0x00000366
	SSLMessageAckGetTransportLockDawLayer      SSLCode = 0x00000367
	SSLMessageSendSetTransportLockDawLayer     SSLCode = 0x00000370
	SSLMessageAckSetTransportLockDawLayer      SSLCode = 0x00000371
	SSLMessageSendProfileNameExists            SSLCode = 0x0000037a
	SSLMessageAckProfileNameExists             SSLCode = 0x0000037b
	SSLMessageSendProfileNameInUse             SSLCode = 0x00000384
	SSLMessageAckProfileNameInUse              SSLCode = 0x00000385
	SSLMessageSendGetDawLayerProtocol          SSLCode = 0x0000038e
	SSLMessageAckGetDawLayerProtocol           SSLCode = 0x0000038f
	SSLMessageSendSaveProfileAs                SSLCode = 0x00000398
	SSLMessageAckSaveProfileAs                 SSLCode = 0x00000399
	SSLMessageSendProfileIsReadOnly            SSLCode = 0x000003a2
	SSLMessageAckProfileIsReadOnly             SSLCode = 0x000003a3
	SSLMessageSendGetProfilePath               SSLCode = 0x000003ac
	SSLMessageAckGetProfilePath                SSLCode = 0x000003ad
	SSLMessageSendGetCcNamesList               SSLCode = 0x000003b6
	SSLMessageAckGetCcNamesList                SSLCode = 0x000003b7
	SSLMessageSendSetCcNamesList               SSLCode = 0x000003c0
	SSLMessageAckSetCcNamesList                SSLCode = 0x000003c1
	SSLMessageSendGetFlipStatus                SSLCode = 0x000003e8
	SSLMessageAckGetFlipStatus                 SSLCode = 0x000003e9
	SSLMessageSendSetFlipStatus                SSLCode = 0x000003f2
	SSLMessageAckSetFlipStatus                 SSLCode = 0x000003f3
	SSLMessageSendGetHandshakingStatus         SSLCode = 0x000003fc
	SSLMessageAckGetHandshakingStatus          SSLCode = 0x000003fd
	SSLMessageSendSetHandshakingStatus         SSLCode = 0x00000406
	SSLMessageAckSetHandshakingStatus          SSLCode = 0x00000407
	SSLMessageSendGetAutoModeOnScribsStatus    SSLCode = 0x00000410
	SSLMessageAckGetAutoModeOnScribsStatus     SSLCode = 0x00000411
	SSLMessageSendSetAutoModeOnScribsStatus    SSLCode = 0x0000041a
	SSLMessageAckSetAutoModeOnScribsStatus     SSLCode = 0x0000041b
	SSLMessageSendGetDefaultWheelModeStatus    SSLCode = 0x00000424
	SSLMessageAckGetDefaultWheelModeStatus     SSLCode = 0x00000425
	SSLMessageSendSetDefaultWheelModeStatus    SSLCode = 0x0000042e
	SSLMessageAckSetDefaultWheelModeStatus     SSLCode = 0x0000042f
	SSLMessageSendSetFaderDbReadoutStatus      SSLCode = 0x00000438
	SSLMessageAckSetFaderDbReadoutStatus       SSLCode = 0x00000439
	SSLMessageSendGetFaderDbReadoutStatus      SSLCode = 0x00000442
	SSLMessageAckGetFaderDbReadoutStatus       SSLCode = 0x00000443
	SSLMessageSendDeleteMix                    SSLCode = 0x0000044c
	SSLMessageAckDeleteMix                     SSLCode = 0x0000044d
	SSLMessageSendCopyMix                      SSLCode = 0x0000044e
	SSLMessageAckCopyMix                       SSLCode = 0x0000044f
	SSLMessageGetXpatchPresetsList             SSLCode = 0x000007d0
	SSLMessageGetXpatchPresetsListReply        SSLCode = 0x000007d1
	SSLMessageSendPresetData                   SSLCode = 0x000007d5
	SSLMessageSetXpatchPresetSelected          SSLCode = 0x000007da
	SSLMessageGetXpatchPresetSelected          SSLCode = 0x000007db
	SSLMessageSetXpatchPresetSelectedReply     SSLCode = 0x000007dc
	SSLMessageGetXpatchPresetEdited            SSLCode = 0x000007dd
	SSLMessageGetXpatchPresetEditedReply       SSLCode = 0x000007de
	SSLMessageSetXpatchPresetName              SSLCode = 0x000007e4
	SSLMessageDeleteXpatchPreset               SSLCode = 0x000007ee
	SSLMessagePasteXpatchPreset                SSLCode = 0x000007f8
	SSLMessageSwapXpatchPreset                 SSLCode = 0x00000802
	SSLMessageSaveXpatchPreset                 SSLCode = 0x00000803
	SSLMessageGetXpatchChanSetup               SSLCode = 0x0000080c
	SSLMessageGetXpatchChanSetupReply          SSLCode = 0x0000080d
	SSLMessageSetXpatchInputMinus10db          SSLCode = 0x00000816
	SSLMessageSetXpatchInputMinus10dbReply     SSLCode = 0x00000817
	SSLMessageSetXpatchOutputMinus10db         SSLCode = 0x00000820
	SSLMessageSetXpatchOutputMinus10dbReply    SSLCode = 0x00000821
	SSLMessageSetXpatchChanMode                SSLCode = 0x0000082a
	SSLMessageSetXpatchChanModeReply           SSLCode = 0x0000082b
	SSLMessageSetXpatchDeviceName              SSLCode = 0x00000bb8
	SSLMessageSetXpatchDeviceNameReply         SSLCode = 0x00000bb9
	SSLMessageSetXpatchDestName                SSLCode = 0x00000bc2
	SSLMessageSetXpatchDestNameReply           SSLCode = 0x00000bc3
	SSLMessageGetXpatchMidiSetup               SSLCode = 0x00000bc7
	SSLMessageGetXpatchMidiSetupReply          SSLCode = 0x00000bc8
	SSLMessageSetXpatchMidiEnable              SSLCode = 0x00000bcc
	SSLMessageSetXpatchMidiEnableReply         SSLCode = 0x00000bcd
	SSLMessageSetXpatchMidiChannel             SSLCode = 0x00000be0
	SSLMessageSetXpatchMidiChannelReply        SSLCode = 0x00000be1
	SSLMessageGetXpatchRoutingData             SSLCode = 0x00000bea
	SSLMessageGetXpatchRoutingDataReply        SSLCode = 0x00000beb
	SSLMessageSetXpatchRoute                   SSLCode = 0x00000bf4
	SSLMessageGetXpatchChainsList              SSLCode = 0x00000fa0
	SSLMessageGetXpatchChainsListReply         SSLCode = 0x00000fa1
	SSLMessageSendChainData                    SSLCode = 0x00000fa5
	SSLMessageSetXpatchChainName               SSLCode = 0x00000faa
	SSLMessageDeleteXpatchChain                SSLCode = 0x00000fb4
	SSLMessageSetXpatchEditChain               SSLCode = 0x00000fbe
	SSLMessageGetXpatchEditChain               SSLCode = 0x00000fc8
	SSLMessageGetXpatchEditChainReply          SSLCode = 0x00000fd2
	SSLMessageSetXpatchEditChainLinkSrc        SSLCode = 0x00000fdc
	SSLMessageGetXpatchEditChainTouched        SSLCode = 0x00000fe6
	SSLMessageGetXpatchEditChainTouchedReply   SSLCode = 0x00000fe7
	SSLMessageSaveXpatchEditChain              SSLCode = 0x00000ff0
	SSLMessageSetXpatchLinkReplaceMode         SSLCode = 0x00000ffa
	SSLMessageSetXpatchLinkReplaceModeReply    SSLCode = 0x00000ffb
	SSLMessageSendClearAll                     SSLCode = 0x00001388
	SSLMessageSetDhcp                          SSLCode = 0x000013ec
	SSLMessageSetIp                            SSLCode = 0x000013f6
	SSLMessageGetIpSettings                    SSLCode = 0x00001400
	SSLMessageSetIpSettingsReply               SSLCode = 0x00001401
	SSLMessageTestComms                        SSLCode = 0x00001402
	SSLMessageTestCommsReply                   SSLCode = 0x00001403
	SSLMessageSendSetDanteEnabled              SSLCode = 0x0000140a
	SSLMessageSendGetDanteEnabled              SSLCode = 0x0000140c
	SSLMessageAckGetDanteEnabled               SSLCode = 0x0000140d
	SSLMessageSendGetCpuVersion                SSLCode = 0x0000140e
	SSLMessageAckGetCpuVersion                 SSLCode = 0x0000140f
)

func SSLCodeToString(c SSLCode) string {
	switch c {
	case SSLMessageToDesk:
		return "ToDesk"
	case SSLMessageToRemote:
		return "ToRemote"
	case SSLMessageGetDesk:
		return "GetDesk"
	case SSLMessageGetDeskReply:
		return "GetDeskReply"
	case SSLMessageSendHeartbeat:
		return "SendHeartbeat"
	case SSLMessageSetOwnName:
		return "SetOwnName"
	case SSLMessageSetOwnNameReply:
		return "SetOwnNameReply"
	case SSLMessageGetProjectNameAndTitle:
		return "GetProjectNameAndTitle"
	case SSLMessageGetProjectNameAndTitleReply:
		return "GetProjectNameAndTitleReply"
	case SSLMessageGetSync:
		return "GetSync"
	case SSLMessageXpatchIphoneGetDesk:
		return "XpatchIphoneGetDesk"
	case SSLMessageGetChanNamesAndImages:
		return "GetChanNamesAndImages"
	case SSLMessageGetChanNamesAndImagesReply:
		return "GetChanNamesAndImagesReply"
	case SSLMessageSetDefaultChanNames:
		return "SetDefaultChanNames"
	case SSLMessageSetDefaultChanNamesReply:
		return "SetDefaultChanNamesReply"
	case SSLMessageSetChanNames:
		return "SetChanNames"
	case SSLMessageSetChanNamesReply:
		return "SetChanNamesReply"
	case SSLMessageSetChanImages:
		return "SetChanImages"
	case SSLMessageSetChanImagesReply:
		return "SetChanImagesReply"
	case SSLMessageGetIsChanStereo:
		return "GetIsChanStereo"
	case SSLMessageGetIsChanStereoReply:
		return "GetIsChanStereoReply"
	case SSLMessageGetImageLibNames:
		return "GetImageLibNames"
	case SSLMessageGetImageLibNamesReply:
		return "GetImageLibNamesReply"
	case SSLMessageGetExtNames:
		return "GetExtNames"
	case SSLMessageGetExtNamesReply:
		return "GetExtNamesReply"
	case SSLMessageSetExtNames:
		return "SetExtNames"
	case SSLMessageSetExtNamesReply:
		return "SetExtNamesReply"
	case SSLMessageGetDirectoryList:
		return "GetDirectoryList"
	case SSLMessageGetDirectoryListReply:
		return "GetDirectoryListReply"
	case SSLMessageGetMixPassesList:
		return "GetMixPassesList"
	case SSLMessageGetMixPassesListReply:
		return "GetMixPassesListReply"
	case SSLMessageGetTrList:
		return "GetTrList"
	case SSLMessageGetTrListReply:
		return "GetTrListReply"
	case SSLMessageSetFileInfo:
		return "SetFileInfo"
	case SSLMessageSetFileInfoReply:
		return "SetFileInfoReply"
	case SSLMessageSendDiskInfo:
		return "SendDiskInfo"
	case SSLMessageAckSendDiskInfo:
		return "AckSendDiskInfo"
	case SSLMessageGetDiskInfo:
		return "GetDiskInfo"
	case SSLMessageAckGetDiskInfo:
		return "AckGetDiskInfo"
	case SSLMessageSendUpdateBlock:
		return "SendUpdateBlock"
	case SSLMessageAckUpdateBlock:
		return "AckUpdateBlock"
	case SSLMessageSendSaveToFlash:
		return "SendSaveToFlash"
	case SSLMessageAckSaveToFlash:
		return "AckSaveToFlash"
	case SSLMessageSendCreateZip:
		return "SendCreateZip"
	case SSLMessageAckCreateZip:
		return "AckCreateZip"
	case SSLMessageSendRequestFileBlock:
		return "SendRequestFileBlock"
	case SSLMessageAckRequestFileBlock:
		return "AckRequestFileBlock"
	case SSLMessageSendWriteFileBlock:
		return "SendWriteFileBlock"
	case SSLMessageAckWriteFileBlock:
		return "AckWriteFileBlock"
	case SSLMessageSendUnzipFile:
		return "SendUnzipFile"
	case SSLMessageAckUnzipFile:
		return "AckUnzipFile"
	case SSLMessageSendMoveFile:
		return "SendMoveFile"
	case SSLMessageAckMoveFile:
		return "AckMoveFile"
	case SSLMessageSendCopyFile:
		return "SendCopyFile"
	case SSLMessageAckCopyFile:
		return "AckCopyFile"
	case SSLMessageSendDeleteFile:
		return "SendDeleteFile"
	case SSLMessageAckDeleteFile:
		return "AckDeleteFile"
	case SSLMessageSendRtc:
		return "SendRtc"
	case SSLMessageAckRtc:
		return "AckRtc"
	case SSLMessageSendSmartCopy:
		return "SendSmartCopy"
	case SSLMessageAckSmartCopy:
		return "AckSmartCopy"
	case SSLMessageSendSmartCopyWithName:
		return "SendSmartCopyWithName"
	case SSLMessageAckSmartCopyWithName:
		return "AckSmartCopyWithName"
	case SSLMessageSendTitleDetailsChanged:
		return "SendTitleDetailsChanged"
	case SSLMessageAckTitleDetailsChanged:
		return "AckTitleDetailsChanged"
	case SSLMessageSendMakeNewProject:
		return "SendMakeNewProject"
	case SSLMessageAckMakeNewProject:
		return "AckMakeNewProject"
	case SSLMessageSendMakeNewProjectTitle:
		return "SendMakeNewProjectTitle"
	case SSLMessageAckMakeNewProjectTitle:
		return "AckMakeNewProjectTitle"
	case SSLMessageSendMakeNewProjectTitleWithName:
		return "SendMakeNewProjectTitleWithName"
	case SSLMessageAckMakeNewProjectTitleWithName:
		return "AckMakeNewProjectTitleWithName"
	case SSLMessageSendSelectProjectTitle:
		return "SendSelectProjectTitle"
	case SSLMessageAckSelectProjectTitle:
		return "AckSelectProjectTitle"
	case SSLMessageSendDeleteProjectTitle:
		return "SendDeleteProjectTitle"
	case SSLMessageAckDeleteProjectTitle:
		return "AckDeleteProjectTitle"
	case SSLMessageSendDeleteProject:
		return "SendDeleteProject"
	case SSLMessageAckDeleteProject:
		return "AckDeleteProject"
	case SSLMessageSendCopyProjectTitle:
		return "SendCopyProjectTitle"
	case SSLMessageAckCopyProjectTitle:
		return "AckCopyProjectTitle"
	case SSLMessageSendCopyProject:
		return "SendCopyProject"
	case SSLMessageAckCopyProject:
		return "AckCopyProject"
	case SSLMessageSendArchiveDone:
		return "SendArchiveDone"
	case SSLMessageAckArchiveDone:
		return "AckArchiveDone"
	case SSLMessageSendMakeNewProjectWithName:
		return "SendMakeNewProjectWithName"
	case SSLMessageAckMakeNewProjectWithName:
		return "AckMakeNewProjectWithName"
	case SSLMessageSendMakeNewProjectWithPresetOpts:
		return "SendMakeNewProjectWithPresetOpts"
	case SSLMessageAckMakeNewProjectWithPresetOpts:
		return "AckMakeNewProjectWithPresetOpts"
	case SSLMessageSendSetTrEnable:
		return "SendSetTrEnable"
	case SSLMessageAckSetTrEnable:
		return "AckSetTrEnable"
	case SSLMessageSendGetTrState:
		return "SendGetTrState"
	case SSLMessageAckGetTrState:
		return "AckGetTrState"
	case SSLMessageSendTakeTrSnap:
		return "SendTakeTrSnap"
	case SSLMessageAckTakeTrSnap:
		return "AckTakeTrSnap"
	case SSLMessageSendSelectTrSnap:
		return "SendSelectTrSnap"
	case SSLMessageAckSelectTrSnap:
		return "AckSelectTrSnap"
	case SSLMessageSendDeleteTrSnap:
		return "SendDeleteTrSnap"
	case SSLMessageAckDeleteTrSnap:
		return "AckDeleteTrSnap"
	case SSLMessageSendCopyTrChanData:
		return "SendCopyTrChanData"
	case SSLMessageAckCopyTrChanData:
		return "AckCopyTrChanData"
	case SSLMessageSendSwopTrChanData:
		return "SendSwopTrChanData"
	case SSLMessageAckSwopTrChanData:
		return "AckSwopTrChanData"
	case SSLMessageSendSetTrAllChans:
		return "SendSetTrAllChans"
	case SSLMessageAckSetTrAllChans:
		return "AckSetTrAllChans"
	case SSLMessageSendSetTrChan:
		return "SendSetTrChan"
	case SSLMessageAckSetTrChan:
		return "AckSetTrChan"
	case SSLMessageSendGetInsertNames:
		return "SendGetInsertNames"
	case SSLMessageAckGetInsertNames:
		return "AckGetInsertNames"
	case SSLMessageSendSetInsertNames:
		return "SendSetInsertNames"
	case SSLMessageAckSetInsertNames:
		return "AckSetInsertNames"
	case SSLMessageSendGetChainNames:
		return "SendGetChainNames"
	case SSLMessageAckGetChainNames:
		return "AckGetChainNames"
	case SSLMessageSendSetInsertToChan:
		return "SendSetInsertToChan"
	case SSLMessageAckSetInsertToChan:
		return "AckSetInsertToChan"
	case SSLMessageSendGetChanMatrixInfo:
		return "SendGetChanMatrixInfo"
	case SSLMessageAckGetChanMatrixInfo:
		return "AckGetChanMatrixInfo"
	case SSLMessageSendGetNumChains:
		return "SendGetNumChains"
	case SSLMessageAckGetNumChains:
		return "AckGetNumChains"
	case SSLMessageSendNewEditChain:
		return "SendNewEditChain"
	case SSLMessageAckNewEditChain:
		return "AckNewEditChain"
	case SSLMessageSendSaveEditChain:
		return "SendSaveEditChain"
	case SSLMessageAckSaveEditChain:
		return "AckSaveEditChain"
	case SSLMessageSendGetEditChainData:
		return "SendGetEditChainData"
	case SSLMessageAckGetEditChainData:
		return "AckGetEditChainData"
	case SSLMessageSendAddDevToEditChain:
		return "SendAddDevToEditChain"
	case SSLMessageAckAddDevToEditChain:
		return "AckAddDevToEditChain"
	case SSLMessageSendCancelEditChain:
		return "SendCancelEditChain"
	case SSLMessageAckCancelEditChain:
		return "AckCancelEditChain"
	case SSLMessageSendAssignChainToChan:
		return "SendAssignChainToChan"
	case SSLMessageAckAssignChainToChan:
		return "AckAssignChainToChan"
	case SSLMessageSendDeassignChan:
		return "SendDeassignChan"
	case SSLMessageAckDeassignChan:
		return "AckDeassignChan"
	case SSLMessageSendRemFromEditChain:
		return "SendRemFromEditChain"
	case SSLMessageAckRemFromEditChain:
		return "AckRemFromEditChain"
	case SSLMessageSendEditExistingChain:
		return "SendEditExistingChain"
	case SSLMessageAckEditExistingChain:
		return "AckEditExistingChain"
	case SSLMessageSendDeleteChain:
		return "SendDeleteChain"
	case SSLMessageAckDeleteChain:
		return "AckDeleteChain"
	case SSLMessageSendRenameEditChain:
		return "SendRenameEditChain"
	case SSLMessageAckRenameEditChain:
		return "AckRenameEditChain"
	case SSLMessageSendInsDevToEditChain:
		return "SendInsDevToEditChain"
	case SSLMessageAckInsDevToEditChain:
		return "AckInsDevToEditChain"
	case SSLMessageSendGetValidEditChainDevs:
		return "SendGetValidEditChainDevs"
	case SSLMessageAckGetValidEditChainDevs:
		return "AckGetValidEditChainDevs"
	case SSLMessageSendGetInsertInfoV2:
		return "SendGetInsertInfoV2"
	case SSLMessageAckGetInsertInfoV2:
		return "AckGetInsertInfoV2"
	case SSLMessageSendSetInsertNamesV2:
		return "SendSetInsertNamesV2"
	case SSLMessageAckSetInsertNamesV2:
		return "AckSetInsertNamesV2"
	case SSLMessageSendGetChainInfoV2:
		return "SendGetChainInfoV2"
	case SSLMessageAckGetChainInfoV2:
		return "AckGetChainInfoV2"
	case SSLMessageSendSetInsertToChanV2:
		return "SendSetInsertToChanV2"
	case SSLMessageAckSetInsertToChanV2:
		return "AckSetInsertToChanV2"
	case SSLMessageSendGetChanMatrixInfoV2:
		return "SendGetChanMatrixInfoV2"
	case SSLMessageAckGetChanMatrixInfoV2:
		return "AckGetChanMatrixInfoV2"
	case SSLMessageSendGetNumChainsV2:
		return "SendGetNumChainsV2"
	case SSLMessageAckGetNumChainsV2:
		return "AckGetNumChainsV2"
	case SSLMessageSendAssignChainToChanV2:
		return "SendAssignChainToChanV2"
	case SSLMessageAckAssignChainToChanV2:
		return "AckAssignChainToChanV2"
	case SSLMessageSendDeassignChanV2:
		return "SendDeassignChanV2"
	case SSLMessageAckDeassignChanV2:
		return "AckDeassignChanV2"
	case SSLMessageSendDeleteChainV2:
		return "SendDeleteChainV2"
	case SSLMessageAckDeleteChainV2:
		return "AckDeleteChainV2"
	case SSLMessageSendRenameChain:
		return "SendRenameChain"
	case SSLMessageAckRenameChain:
		return "AckRenameChain"
	case SSLMessageSendSaveInsertsToChain:
		return "SendSaveInsertsToChain"
	case SSLMessageAckSaveInsertsToChain:
		return "AckSaveInsertsToChain"
	case SSLMessageSendSetChanInserts:
		return "SendSetChanInserts"
	case SSLMessageAckSetChanInserts:
		return "AckSetChanInserts"
	case SSLMessageSendDeleteChanInsert:
		return "SendDeleteChanInsert"
	case SSLMessageAckDeleteChanInsert:
		return "AckDeleteChanInsert"
	case SSLMessageSendReorderChanInserts:
		return "SendReorderChanInserts"
	case SSLMessageAckReorderChanInserts:
		return "AckReorderChanInserts"
	case SSLMessageSendSetChanStereoInsert:
		return "SendSetChanStereoInsert"
	case SSLMessageAckSetChanStereoInsert:
		return "AckSetChanStereoInsert"
	case SSLMessageSendGetMatrixPresetList:
		return "SendGetMatrixPresetList"
	case SSLMessageAckGetMatrixPresetList:
		return "AckGetMatrixPresetList"
	case SSLMessageSendLoadMatrixPreset:
		return "SendLoadMatrixPreset"
	case SSLMessageAckLoadMatrixPreset:
		return "AckLoadMatrixPreset"
	case SSLMessageSendSaveMatrixPreset:
		return "SendSaveMatrixPreset"
	case SSLMessageAckSaveMatrixPreset:
		return "AckSaveMatrixPreset"
	case SSLMessageSendDeleteMatrixPreset:
		return "SendDeleteMatrixPreset"
	case SSLMessageAckDeleteMatrixPreset:
		return "AckDeleteMatrixPreset"
	case SSLMessageSendRenameMatrixPreset:
		return "SendRenameMatrixPreset"
	case SSLMessageAckRenameMatrixPreset:
		return "AckRenameMatrixPreset"
	case SSLMessageSendClearInserts:
		return "SendClearInserts"
	case SSLMessageAckClearInserts:
		return "AckClearInserts"
	case SSLMessageSendSetLowerScribMode:
		return "SendSetLowerScribMode"
	case SSLMessageAckSetLowerScribMode:
		return "AckSetLowerScribMode"
	case SSLMessageSendGetLowerScribMode:
		return "SendGetLowerScribMode"
	case SSLMessageAckGetLowerScribMode:
		return "AckGetLowerScribMode"
	case SSLMessageSendSetDisplay1732:
		return "SendSetDisplay1732"
	case SSLMessageAckSetDisplay1732:
		return "AckSetDisplay1732"
	case SSLMessageSendGetDisplay1732:
		return "SendGetDisplay1732"
	case SSLMessageAckGetDisplay1732:
		return "AckGetDisplay1732"
	case SSLMessageSendSetFlipScribStrip:
		return "SendSetFlipScribStrip"
	case SSLMessageAckSetFlipScribStrip:
		return "AckSetFlipScribStrip"
	case SSLMessageSendGetFlipScribStrip:
		return "SendGetFlipScribStrip"
	case SSLMessageAckGetFlipScribStrip:
		return "AckGetFlipScribStrip"
	case SSLMessageSendRenameChanNamesPreset:
		return "SendRenameChanNamesPreset"
	case SSLMessageAckRenameChanNamesPreset:
		return "AckRenameChanNamesPreset"
	case SSLMessageSendDeleteChanNamesPreset:
		return "SendDeleteChanNamesPreset"
	case SSLMessageAckDeleteChanNamesPreset:
		return "AckDeleteChanNamesPreset"
	case SSLMessageSendGetChanNamesPresetList:
		return "SendGetChanNamesPresetList"
	case SSLMessageAckGetChanNamesPresetList:
		return "AckGetChanNamesPresetList"
	case SSLMessageSendSaveChanNamesPreset:
		return "SendSaveChanNamesPreset"
	case SSLMessageAckSaveChanNamesPreset:
		return "AckSaveChanNamesPreset"
	case SSLMessageSendLoadChanNamesPreset:
		return "SendLoadChanNamesPreset"
	case SSLMessageAckLoadChanNamesPreset:
		return "AckLoadChanNamesPreset"
	case SSLMessageSendGetAutomationMode:
		return "SendGetAutomationMode"
	case SSLMessageAckGetAutomationMode:
		return "AckGetAutomationMode"
	case SSLMessageSendSetAutomationMode:
		return "SendSetAutomationMode"
	case SSLMessageAckSetAutomationMode:
		return "AckSetAutomationMode"
	case SSLMessageSendGetMotorsOffTouchEn:
		return "SendGetMotorsOffTouchEn"
	case SSLMessageAckGetMotorsOffTouchEn:
		return "AckGetMotorsOffTouchEn"
	case SSLMessageSendSetMotorsOffTouchEn:
		return "SendSetMotorsOffTouchEn"
	case SSLMessageAckSetMotorsOffTouchEn:
		return "AckSetMotorsOffTouchEn"
	case SSLMessageSendGetMdacMeterEn:
		return "SendGetMdacMeterEn"
	case SSLMessageAckGetMdacMeterEn:
		return "AckGetMdacMeterEn"
	case SSLMessageSendSetMdacMeterEn:
		return "SendSetMdacMeterEn"
	case SSLMessageAckSetMdacMeterEn:
		return "AckSetMdacMeterEn"
	case SSLMessageSendGetEditKeymapName:
		return "SendGetEditKeymapName"
	case SSLMessageAckGetEditKeymapName:
		return "AckGetEditKeymapName"
	case SSLMessageSendSetEditKeymapName:
		return "SendSetEditKeymapName"
	case SSLMessageAckSetEditKeymapName:
		return "AckSetEditKeymapName"
	case SSLMessageSendGetEditKeymapData:
		return "SendGetEditKeymapData"
	case SSLMessageAckGetEditKeymapData:
		return "AckGetEditKeymapData"
	case SSLMessageSendGetEditKeymapKeycaps:
		return "SendGetEditKeymapKeycaps"
	case SSLMessageAckGetEditKeymapKeycap:
		return "AckGetEditKeymapKeycap"
	case SSLMessageSendGetEditKeymapSize:
		return "SendGetEditKeymapSize"
	case SSLMessageAckGetEditKeymapSize:
		return "AckGetEditKeymapSize"
	case SSLMessageSendSetUsbCmd:
		return "SendSetUsbCmd"
	case SSLMessageAckSetUsbCmd:
		return "AckSetUsbCmd"
	case SSLMessageSendSetKeycapName:
		return "SendSetKeycapName"
	case SSLMessageAckSetKeycapName:
		return "AckSetKeycapName"
	case SSLMessageSendSetKeyBlank:
		return "SendSetKeyBlank"
	case SSLMessageAckSetKeyBlank:
		return "AckSetKeyBlank"
	case SSLMessageSendSetSaveEditKeymap:
		return "SendSetSaveEditKeymap"
	case SSLMessageAckSetSaveEditKeymap:
		return "AckSetSaveEditKeymap"
	case SSLMessageSendGetMidiFunctionList:
		return "SendGetMidiFunctionList"
	case SSLMessageAckGetMidiFunctionList:
		return "AckGetMidiFunctionList"
	case SSLMessageSendSetMidiCmd:
		return "SendSetMidiCmd"
	case SSLMessageAckSetMidiCmd:
		return "AckSetMidiCmd"
	case SSLMessageSendSetNewMenuCmd:
		return "SendSetNewMenuCmd"
	case SSLMessageAckSetNewMenuCmd:
		return "AckSetNewMenuCmd"
	case SSLMessageSendSetMenuSubKeycapName:
		return "SendSetMenuSubKeycapName"
	case SSLMessageAckSetMenuSubKeycapName:
		return "AckSetMenuSubKeycapName"
	case SSLMessageSendSetMenuSubMidiCmd:
		return "SendSetMenuSubMidiCmd"
	case SSLMessageAckSetMenuSubMidiCmd:
		return "AckSetMenuSubMidiCmd"
	case SSLMessageSendSetMenuSubUsbCmd:
		return "SendSetMenuSubUsbCmd"
	case SSLMessageAckSetMenuSubUsbCmd:
		return "AckSetMenuSubUsbCmd"
	case SSLMessageSendSetMenuSubBlankCmd:
		return "SendSetMenuSubBlankCmd"
	case SSLMessageAckSetMenuSubBlankCmd:
		return "AckSetMenuSubBlankCmd"
	case SSLMessageSendRestartConsole:
		return "SendRestartConsole"
	case SSLMessageSendFollowKeyState:
		return "SendFollowKeyState"
	case SSLMessageAckFollowKeyState:
		return "AckFollowKeyState"
	case SSLMessageSendGetProfileForDawLayer:
		return "SendGetProfileForDawLayer"
	case SSLMessageAckGetProfileForDawLayer:
		return "AckGetProfileForDawLayer"
	case SSLMessageSendCopyProfileToNew:
		return "SendCopyProfileToNew"
	case SSLMessageAckCopyProfileToNew:
		return "AckCopyProfileToNew"
	case SSLMessageSendSetProfileForDawLayer:
		return "SendSetProfileForDawLayer"
	case SSLMessageAckSetProfileForDawLayer:
		return "AckSetProfileForDawLayer"
	case SSLMessageSendClearProfileForDawLayer:
		return "SendClearProfileForDawLayer"
	case SSLMessageAckClearProfileForDawLayer:
		return "AckClearProfileForDawLayer"
	case SSLMessageSendGetProfiles:
		return "SendGetProfiles"
	case SSLMessageAckGetProfiles:
		return "AckGetProfiles"
	case SSLMessageSendRenameProfiles:
		return "SendRenameProfiles"
	case SSLMessageAckRenameProfiles:
		return "AckRenameProfiles"
	case SSLMessageSendDeleteProfiles:
		return "SendDeleteProfiles"
	case SSLMessageAckDeleteProfiles:
		return "AckDeleteProfiles"
	case SSLMessageSendGetTransportLockDawLayer:
		return "SendGetTransportLockDawLayer"
	case SSLMessageAckGetTransportLockDawLayer:
		return "AckGetTransportLockDawLayer"
	case SSLMessageSendSetTransportLockDawLayer:
		return "SendSetTransportLockDawLayer"
	case SSLMessageAckSetTransportLockDawLayer:
		return "AckSetTransportLockDawLayer"
	case SSLMessageSendProfileNameExists:
		return "SendProfileNameExists"
	case SSLMessageAckProfileNameExists:
		return "AckProfileNameExists"
	case SSLMessageSendProfileNameInUse:
		return "SendProfileNameInUse"
	case SSLMessageAckProfileNameInUse:
		return "AckProfileNameInUse"
	case SSLMessageSendGetDawLayerProtocol:
		return "SendGetDawLayerProtocol"
	case SSLMessageAckGetDawLayerProtocol:
		return "AckGetDawLayerProtocol"
	case SSLMessageSendSaveProfileAs:
		return "SendSaveProfileAs"
	case SSLMessageAckSaveProfileAs:
		return "AckSaveProfileAs"
	case SSLMessageSendProfileIsReadOnly:
		return "SendProfileIsReadOnly"
	case SSLMessageAckProfileIsReadOnly:
		return "AckProfileIsReadOnly"
	case SSLMessageSendGetProfilePath:
		return "SendGetProfilePath"
	case SSLMessageAckGetProfilePath:
		return "AckGetProfilePath"
	case SSLMessageSendGetCcNamesList:
		return "SendGetCcNamesList"
	case SSLMessageAckGetCcNamesList:
		return "AckGetCcNamesList"
	case SSLMessageSendSetCcNamesList:
		return "SendSetCcNamesList"
	case SSLMessageAckSetCcNamesList:
		return "AckSetCcNamesList"
	case SSLMessageSendGetFlipStatus:
		return "SendGetFlipStatus"
	case SSLMessageAckGetFlipStatus:
		return "AckGetFlipStatus"
	case SSLMessageSendSetFlipStatus:
		return "SendSetFlipStatus"
	case SSLMessageAckSetFlipStatus:
		return "AckSetFlipStatus"
	case SSLMessageSendGetHandshakingStatus:
		return "SendGetHandshakingStatus"
	case SSLMessageAckGetHandshakingStatus:
		return "AckGetHandshakingStatus"
	case SSLMessageSendSetHandshakingStatus:
		return "SendSetHandshakingStatus"
	case SSLMessageAckSetHandshakingStatus:
		return "AckSetHandshakingStatus"
	case SSLMessageSendGetAutoModeOnScribsStatus:
		return "SendGetAutoModeOnScribsStatus"
	case SSLMessageAckGetAutoModeOnScribsStatus:
		return "AckGetAutoModeOnScribsStatus"
	case SSLMessageSendSetAutoModeOnScribsStatus:
		return "SendSetAutoModeOnScribsStatus"
	case SSLMessageAckSetAutoModeOnScribsStatus:
		return "AckSetAutoModeOnScribsStatus"
	case SSLMessageSendGetDefaultWheelModeStatus:
		return "SendGetDefaultWheelModeStatus"
	case SSLMessageAckGetDefaultWheelModeStatus:
		return "AckGetDefaultWheelModeStatus"
	case SSLMessageSendSetDefaultWheelModeStatus:
		return "SendSetDefaultWheelModeStatus"
	case SSLMessageAckSetDefaultWheelModeStatus:
		return "AckSetDefaultWheelModeStatus"
	case SSLMessageSendSetFaderDbReadoutStatus:
		return "SendSetFaderDbReadoutStatus"
	case SSLMessageAckSetFaderDbReadoutStatus:
		return "AckSetFaderDbReadoutStatus"
	case SSLMessageSendGetFaderDbReadoutStatus:
		return "SendGetFaderDbReadoutStatus"
	case SSLMessageAckGetFaderDbReadoutStatus:
		return "AckGetFaderDbReadoutStatus"
	case SSLMessageSendDeleteMix:
		return "SendDeleteMix"
	case SSLMessageAckDeleteMix:
		return "AckDeleteMix"
	case SSLMessageSendCopyMix:
		return "SendCopyMix"
	case SSLMessageAckCopyMix:
		return "AckCopyMix"
	case SSLMessageGetXpatchPresetsList:
		return "GetXpatchPresetsList"
	case SSLMessageGetXpatchPresetsListReply:
		return "GetXpatchPresetsListReply"
	case SSLMessageSendPresetData:
		return "SendPresetData"
	case SSLMessageSetXpatchPresetSelected:
		return "SetXpatchPresetSelected"
	case SSLMessageGetXpatchPresetSelected:
		return "GetXpatchPresetSelected"
	case SSLMessageSetXpatchPresetSelectedReply:
		return "SetXpatchPresetSelectedReply"
	case SSLMessageGetXpatchPresetEdited:
		return "GetXpatchPresetEdited"
	case SSLMessageGetXpatchPresetEditedReply:
		return "GetXpatchPresetEditedReply"
	case SSLMessageSetXpatchPresetName:
		return "SetXpatchPresetName"
	case SSLMessageDeleteXpatchPreset:
		return "DeleteXpatchPreset"
	case SSLMessagePasteXpatchPreset:
		return "PasteXpatchPreset"
	case SSLMessageSwapXpatchPreset:
		return "SwapXpatchPreset"
	case SSLMessageSaveXpatchPreset:
		return "SaveXpatchPreset"
	case SSLMessageGetXpatchChanSetup:
		return "GetXpatchChanSetup"
	case SSLMessageGetXpatchChanSetupReply:
		return "GetXpatchChanSetupReply"
	case SSLMessageSetXpatchInputMinus10db:
		return "SetXpatchInputMinus10db"
	case SSLMessageSetXpatchInputMinus10dbReply:
		return "SetXpatchInputMinus10dbReply"
	case SSLMessageSetXpatchOutputMinus10db:
		return "SetXpatchOutputMinus10db"
	case SSLMessageSetXpatchOutputMinus10dbReply:
		return "SetXpatchOutputMinus10dbReply"
	case SSLMessageSetXpatchChanMode:
		return "SetXpatchChanMode"
	case SSLMessageSetXpatchChanModeReply:
		return "SetXpatchChanModeReply"
	case SSLMessageSetXpatchDeviceName:
		return "SetXpatchDeviceName"
	case SSLMessageSetXpatchDeviceNameReply:
		return "SetXpatchDeviceNameReply"
	case SSLMessageSetXpatchDestName:
		return "SetXpatchDestName"
	case SSLMessageSetXpatchDestNameReply:
		return "SetXpatchDestNameReply"
	case SSLMessageGetXpatchMidiSetup:
		return "GetXpatchMidiSetup"
	case SSLMessageGetXpatchMidiSetupReply:
		return "GetXpatchMidiSetupReply"
	case SSLMessageSetXpatchMidiEnable:
		return "SetXpatchMidiEnable"
	case SSLMessageSetXpatchMidiEnableReply:
		return "SetXpatchMidiEnableReply"
	case SSLMessageSetXpatchMidiChannel:
		return "SetXpatchMidiChannel"
	case SSLMessageSetXpatchMidiChannelReply:
		return "SetXpatchMidiChannelReply"
	case SSLMessageGetXpatchRoutingData:
		return "GetXpatchRoutingData"
	case SSLMessageGetXpatchRoutingDataReply:
		return "GetXpatchRoutingDataReply"
	case SSLMessageSetXpatchRoute:
		return "SetXpatchRoute"
	case SSLMessageGetXpatchChainsList:
		return "GetXpatchChainsList"
	case SSLMessageGetXpatchChainsListReply:
		return "GetXpatchChainsListReply"
	case SSLMessageSendChainData:
		return "SendChainData"
	case SSLMessageSetXpatchChainName:
		return "SetXpatchChainName"
	case SSLMessageDeleteXpatchChain:
		return "DeleteXpatchChain"
	case SSLMessageSetXpatchEditChain:
		return "SetXpatchEditChain"
	case SSLMessageGetXpatchEditChain:
		return "GetXpatchEditChain"
	case SSLMessageGetXpatchEditChainReply:
		return "GetXpatchEditChainReply"
	case SSLMessageSetXpatchEditChainLinkSrc:
		return "SetXpatchEditChainLinkSrc"
	case SSLMessageGetXpatchEditChainTouched:
		return "GetXpatchEditChainTouched"
	case SSLMessageGetXpatchEditChainTouchedReply:
		return "GetXpatchEditChainTouchedReply"
	case SSLMessageSaveXpatchEditChain:
		return "SaveXpatchEditChain"
	case SSLMessageSetXpatchLinkReplaceMode:
		return "SetXpatchLinkReplaceMode"
	case SSLMessageSetXpatchLinkReplaceModeReply:
		return "SetXpatchLinkReplaceModeReply"
	case SSLMessageSendClearAll:
		return "SendClearAll"
	case SSLMessageSetDhcp:
		return "SetDhcp"
	case SSLMessageSetIp:
		return "SetIp"
	case SSLMessageGetIpSettings:
		return "GetIpSettings"
	case SSLMessageSetIpSettingsReply:
		return "SetIpSettingsReply"
	case SSLMessageTestComms:
		return "TestComms"
	case SSLMessageTestCommsReply:
		return "TestCommsReply"
	case SSLMessageSendSetDanteEnabled:
		return "SendSetDanteEnabled"
	case SSLMessageSendGetDanteEnabled:
		return "SendGetDanteEnabled"
	case SSLMessageAckGetDanteEnabled:
		return "AckGetDanteEnabled"
	case SSLMessageSendGetCpuVersion:
		return "SendGetCpuVersion"
	case SSLMessageAckGetCpuVersion:
		return "AckGetCpuVersion"
	}
	return "unknown"
}
