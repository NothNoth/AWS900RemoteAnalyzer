package main

import (
	"fmt"

	"github.com/quarkslab/wirego/wirego/wirego"
)

// Define here enum identifiers, used to refer to a specific field
const (
	//AWS900 remote common fields
	CmdCode      wirego.FieldId = 1
	DestCode     wirego.FieldId = 2
	DeskSerial   wirego.FieldId = 3
	RemoteSerial wirego.FieldId = 4

	//Common fields
	SSLRemote_Path           wirego.FieldId = 1000
	SSLRemote_Time           wirego.FieldId = 1001
	SSLRemote_Date           wirego.FieldId = 1002
	SSLRemote_Size           wirego.FieldId = 1003
	SSLRemote_FileIndex      wirego.FieldId = 1004
	SSLRemote_IsDir          wirego.FieldId = 1005
	SSLRemote_ChannelNumber  wirego.FieldId = 1006
	SSLRemote_DirectoryName  wirego.FieldId = 1007
	SSLRemote_TRName         wirego.FieldId = 1008
	SSLRemote_TRInfo         wirego.FieldId = 1009
	SSLRemote_TRIsSelected   wirego.FieldId = 1010
	SSLRemote_MixPassName    wirego.FieldId = 1011
	SSLRemote_MixPassNumber  wirego.FieldId = 1012
	SSLRemote_MixPassBasedOn wirego.FieldId = 1013
	SSLRemote_ChannelName    wirego.FieldId = 1014

	//Command specific fields
	CmdFields wirego.FieldId = 10

	CmdResponseIsChanStereo_IsStereo wirego.FieldId = 22

	CmdResponseGetDeskReply_ProductName  wirego.FieldId = 30
	CmdResponseGetDeskReply_VersionMajor wirego.FieldId = 31
	CmdResponseGetDeskReply_VersionMinor wirego.FieldId = 32
	CmdResponseGetDeskReply_Issue        wirego.FieldId = 33
	CmdResponseGetDeskReply_BuildVersion wirego.FieldId = 34
	CmdResponseGetDeskReply_BuildDate    wirego.FieldId = 35

	CmdRequestDirectoryList_Mode wirego.FieldId = 40

	CmdResponseDirectoryList_Index wirego.FieldId = 51
	CmdResponseDirectoryList_Name  wirego.FieldId = 52
	CmdResponseDirectoryList_Info  wirego.FieldId = 53

	CmdResponseSendDiskInfo_Free        wirego.FieldId = 60
	CmdResponseSendDiskInfo_ArchiveDone wirego.FieldId = 62

	CmdResponseExtNames_Number wirego.FieldId = 70
	CmdResponseExtNames_Name   wirego.FieldId = 71

	CmdRequestChanNamesAndImages_FirstChannel wirego.FieldId = 80
	CmdRequestChanNamesAndImages_LastChannel  wirego.FieldId = 81

	CmdResponseGetProjectNameAndTitle_ProjectName wirego.FieldId = 90
	CmdResponseGetProjectNameAndTitle_ProjectInfo wirego.FieldId = 91
	CmdResponseGetProjectNameAndTitle_TitleName   wirego.FieldId = 92
	CmdResponseGetProjectNameAndTitle_TitleInfo   wirego.FieldId = 93

	CmdRequestFileBlock_BlockPosition wirego.FieldId = 101
	CmdRequestFileBlock_BlockLen      wirego.FieldId = 102
	CmdRequestFileBlock_Zero          wirego.FieldId = 103
	CmdRequestFileBlock_FileData      wirego.FieldId = 104
)

// Since we implement the wirego.WiregoInterface we need some structure to hold it.
type WiregoMinimalExample struct {
}

// Unused (but mandatory)
func main() {}

// Called at golang environment initialization (you should probably not touch this)
func init() {
	var wge WiregoMinimalExample

	//Register to the wirego package
	wirego.Register(wge)
	wirego.ResultsCacheEnable(false)
}

// This function is called when the plugin is loaded.
func (WiregoMinimalExample) Setup() error {

	return nil
}

// This function shall return the plugin name
func (WiregoMinimalExample) GetName() string {
	return "AWS900 remote"
}

// This function shall return the wireshark filter
func (WiregoMinimalExample) GetFilter() string {
	return "aws900"
}

// GetFields returns the list of fields descriptor that we may eventually return
// when dissecting a packet payload
func (WiregoMinimalExample) GetFields() []wirego.WiresharkField {
	var fields []wirego.WiresharkField

	//Common fields for all commands
	fields = append(fields, wirego.WiresharkField{WiregoFieldId: CmdCode, Name: "Command", Filter: "aws900.command", ValueType: wirego.ValueTypeUInt32, DisplayMode: wirego.DisplayModeHexadecimal})
	fields = append(fields, wirego.WiresharkField{WiregoFieldId: DestCode, Name: "Dest", Filter: "aws900.dest", ValueType: wirego.ValueTypeUInt32, DisplayMode: wirego.DisplayModeHexadecimal})
	fields = append(fields, wirego.WiresharkField{WiregoFieldId: DeskSerial, Name: "Desk Serial", Filter: "aws900.desk_serial", ValueType: wirego.ValueTypeUInt32, DisplayMode: wirego.DisplayModeDecimal})
	fields = append(fields, wirego.WiresharkField{WiregoFieldId: RemoteSerial, Name: "Remote Serial", Filter: "aws900.remote_serial", ValueType: wirego.ValueTypeUInt32, DisplayMode: wirego.DisplayModeDecimal})

	//Specific command fields
	fields = append(fields, wirego.WiresharkField{WiregoFieldId: CmdFields, Name: "Command arguments", Filter: "aws900.command_args", ValueType: wirego.ValueTypeNone, DisplayMode: wirego.DisplayModeNone})

	//Stereo
	fields = append(fields, wirego.WiresharkField{WiregoFieldId: CmdResponseIsChanStereo_IsStereo, Name: "Is stereo", Filter: "aws900.is_stereo", ValueType: wirego.ValueTypeBool, DisplayMode: wirego.DisplayModeDecimal})

	//Desk
	fields = append(fields, wirego.WiresharkField{WiregoFieldId: CmdResponseGetDeskReply_ProductName, Name: "Product name", Filter: "aws900.deskinfo_name", ValueType: wirego.ValueTypeCString, DisplayMode: wirego.DisplayModeNone})
	fields = append(fields, wirego.WiresharkField{WiregoFieldId: CmdResponseGetDeskReply_VersionMajor, Name: "Version major", Filter: "aws900.deskinfo_version_major", ValueType: wirego.ValueTypeUInt32, DisplayMode: wirego.DisplayModeDecimal})
	fields = append(fields, wirego.WiresharkField{WiregoFieldId: CmdResponseGetDeskReply_VersionMinor, Name: "Version minor", Filter: "aws900.deskinfo_version_minor", ValueType: wirego.ValueTypeUInt32, DisplayMode: wirego.DisplayModeDecimal})
	fields = append(fields, wirego.WiresharkField{WiregoFieldId: CmdResponseGetDeskReply_Issue, Name: "Issue", Filter: "aws900.deskinfo_issue", ValueType: wirego.ValueTypeUInt32, DisplayMode: wirego.DisplayModeDecimal})
	fields = append(fields, wirego.WiresharkField{WiregoFieldId: CmdResponseGetDeskReply_BuildVersion, Name: "Build version", Filter: "aws900.deskinfo_build_version", ValueType: wirego.ValueTypeCString, DisplayMode: wirego.DisplayModeNone})
	fields = append(fields, wirego.WiresharkField{WiregoFieldId: CmdResponseGetDeskReply_BuildDate, Name: "Build date", Filter: "aws900.deskinfo_build_date", ValueType: wirego.ValueTypeCString, DisplayMode: wirego.DisplayModeNone})

	//GetDirectoryListing
	fields = append(fields, wirego.WiresharkField{WiregoFieldId: CmdRequestDirectoryList_Mode, Name: "Mode", Filter: "aws900.listing_flag", ValueType: wirego.ValueTypeBool, DisplayMode: wirego.DisplayModeHexadecimal})

	//Directory Listing response
	fields = append(fields, wirego.WiresharkField{WiregoFieldId: CmdResponseDirectoryList_Index, Name: "File index", Filter: "aws900.listing_file_index", ValueType: wirego.ValueTypeUInt16, DisplayMode: wirego.DisplayModeDecimal})
	fields = append(fields, wirego.WiresharkField{WiregoFieldId: CmdResponseDirectoryList_Name, Name: "Name", Filter: "aws900.listing_name", ValueType: wirego.ValueTypeCString, DisplayMode: wirego.DisplayModeNone})
	fields = append(fields, wirego.WiresharkField{WiregoFieldId: CmdResponseDirectoryList_Info, Name: "Info", Filter: "aws900.listing_info", ValueType: wirego.ValueTypeCString, DisplayMode: wirego.DisplayModeNone})

	//Disk Info
	fields = append(fields, wirego.WiresharkField{WiregoFieldId: CmdResponseSendDiskInfo_Free, Name: "Free", Filter: "aws900.disk_free", ValueType: wirego.ValueTypeUInt32, DisplayMode: wirego.DisplayModeDecimal})
	fields = append(fields, wirego.WiresharkField{WiregoFieldId: CmdResponseSendDiskInfo_ArchiveDone, Name: "Archive done", Filter: "aws900.disk_archive_done", ValueType: wirego.ValueTypeBool, DisplayMode: wirego.DisplayModeDecimal})

	//Ext names
	fields = append(fields, wirego.WiresharkField{WiregoFieldId: CmdResponseExtNames_Number, Name: "Number", Filter: "aws900.ext_number", ValueType: wirego.ValueTypeUInt8, DisplayMode: wirego.DisplayModeDecimal})
	fields = append(fields, wirego.WiresharkField{WiregoFieldId: CmdResponseExtNames_Name, Name: "Name", Filter: "aws900.ext_name", ValueType: wirego.ValueTypeCString, DisplayMode: wirego.DisplayModeNone})

	//Getchan names and images
	fields = append(fields, wirego.WiresharkField{WiregoFieldId: CmdRequestChanNamesAndImages_FirstChannel, Name: "First", Filter: "aws900.get_chan_names_first", ValueType: wirego.ValueTypeUInt8, DisplayMode: wirego.DisplayModeDecimal})
	fields = append(fields, wirego.WiresharkField{WiregoFieldId: CmdRequestChanNamesAndImages_LastChannel, Name: "Last", Filter: "aws900.get_chan_names_last", ValueType: wirego.ValueTypeUInt8, DisplayMode: wirego.DisplayModeDecimal})

	//Project name and title
	fields = append(fields, wirego.WiresharkField{WiregoFieldId: CmdResponseGetProjectNameAndTitle_ProjectName, Name: "Project name", Filter: "aws900.project_name", ValueType: wirego.ValueTypeCString, DisplayMode: wirego.DisplayModeNone})
	fields = append(fields, wirego.WiresharkField{WiregoFieldId: CmdResponseGetProjectNameAndTitle_ProjectInfo, Name: "Project info", Filter: "aws900.project_info", ValueType: wirego.ValueTypeCString, DisplayMode: wirego.DisplayModeNone})
	fields = append(fields, wirego.WiresharkField{WiregoFieldId: CmdResponseGetProjectNameAndTitle_TitleName, Name: "Title name", Filter: "aws900.title_name", ValueType: wirego.ValueTypeCString, DisplayMode: wirego.DisplayModeNone})
	fields = append(fields, wirego.WiresharkField{WiregoFieldId: CmdResponseGetProjectNameAndTitle_TitleInfo, Name: "Title info", Filter: "aws900.title_info", ValueType: wirego.ValueTypeCString, DisplayMode: wirego.DisplayModeNone})

	//Requets file block
	fields = append(fields, wirego.WiresharkField{WiregoFieldId: CmdRequestFileBlock_BlockPosition, Name: "Block position", Filter: "aws900.block_position", ValueType: wirego.ValueTypeUInt32, DisplayMode: wirego.DisplayModeDecimal})
	fields = append(fields, wirego.WiresharkField{WiregoFieldId: CmdRequestFileBlock_BlockLen, Name: "Block length", Filter: "aws900.block_len", ValueType: wirego.ValueTypeUInt32, DisplayMode: wirego.DisplayModeDecimal})
	fields = append(fields, wirego.WiresharkField{WiregoFieldId: CmdRequestFileBlock_Zero, Name: "Zero", Filter: "aws900.zero", ValueType: wirego.ValueTypeUInt32, DisplayMode: wirego.DisplayModeDecimal})
	fields = append(fields, wirego.WiresharkField{WiregoFieldId: CmdRequestFileBlock_FileData, Name: "File contents", Filter: "aws900.file", ValueType: wirego.ValueTypeNone, DisplayMode: wirego.DisplayModeNone})

	//Common fields
	fields = append(fields, wirego.WiresharkField{WiregoFieldId: SSLRemote_Path, Name: "Path", Filter: "aws900.path", ValueType: wirego.ValueTypeCString, DisplayMode: wirego.DisplayModeNone})
	fields = append(fields, wirego.WiresharkField{WiregoFieldId: SSLRemote_Time, Name: "Time", Filter: "aws900.time", ValueType: wirego.ValueTypeCString, DisplayMode: wirego.DisplayModeNone})
	fields = append(fields, wirego.WiresharkField{WiregoFieldId: SSLRemote_Date, Name: "Date", Filter: "aws900.date", ValueType: wirego.ValueTypeCString, DisplayMode: wirego.DisplayModeNone})
	fields = append(fields, wirego.WiresharkField{WiregoFieldId: SSLRemote_Size, Name: "Size", Filter: "aws900.size", ValueType: wirego.ValueTypeUInt32, DisplayMode: wirego.DisplayModeDecimal})
	fields = append(fields, wirego.WiresharkField{WiregoFieldId: SSLRemote_FileIndex, Name: "File index", Filter: "aws900.file_index", ValueType: wirego.ValueTypeUInt16, DisplayMode: wirego.DisplayModeDecimal})
	fields = append(fields, wirego.WiresharkField{WiregoFieldId: SSLRemote_IsDir, Name: "IsDir", Filter: "aws900.is_dir", ValueType: wirego.ValueTypeBool, DisplayMode: wirego.DisplayModeDecimal})

	fields = append(fields, wirego.WiresharkField{WiregoFieldId: SSLRemote_MixPassName, Name: "Mix pass name", Filter: "aws900.mixpass_name", ValueType: wirego.ValueTypeCString, DisplayMode: wirego.DisplayModeNone})
	fields = append(fields, wirego.WiresharkField{WiregoFieldId: SSLRemote_MixPassNumber, Name: "Mix pass number", Filter: "aws900.mixpass_number", ValueType: wirego.ValueTypeUInt32, DisplayMode: wirego.DisplayModeDecimal})
	fields = append(fields, wirego.WiresharkField{WiregoFieldId: SSLRemote_MixPassBasedOn, Name: "Mix pass based on", Filter: "aws900.mixpass_based_on", ValueType: wirego.ValueTypeUInt32, DisplayMode: wirego.DisplayModeDecimal})
	fields = append(fields, wirego.WiresharkField{WiregoFieldId: SSLRemote_TRName, Name: "TR name", Filter: "aws900.tr_name", ValueType: wirego.ValueTypeCString, DisplayMode: wirego.DisplayModeNone})
	fields = append(fields, wirego.WiresharkField{WiregoFieldId: SSLRemote_TRInfo, Name: "TR info", Filter: "aws900.tr_info", ValueType: wirego.ValueTypeCString, DisplayMode: wirego.DisplayModeNone})
	fields = append(fields, wirego.WiresharkField{WiregoFieldId: SSLRemote_TRIsSelected, Name: "TR is selected", Filter: "aws900.tr_is_selected", ValueType: wirego.ValueTypeBool, DisplayMode: wirego.DisplayModeDecimal})
	fields = append(fields, wirego.WiresharkField{WiregoFieldId: SSLRemote_ChannelNumber, Name: "Channel number", Filter: "aws900.channel_number", ValueType: wirego.ValueTypeUInt8, DisplayMode: wirego.DisplayModeDecimal})
	fields = append(fields, wirego.WiresharkField{WiregoFieldId: SSLRemote_ChannelName, Name: "Channel name", Filter: "aws900.channel_name", ValueType: wirego.ValueTypeCString, DisplayMode: wirego.DisplayModeNone})
	fields = append(fields, wirego.WiresharkField{WiregoFieldId: SSLRemote_DirectoryName, Name: "Directory name", Filter: "aws900.directory_name", ValueType: wirego.ValueTypeCString, DisplayMode: wirego.DisplayModeNone})

	return fields
}

// GetDetectionFilters returns a wireshark filter that will select which packets
// will be sent to your dissector for parsing.
// Two types of filters can be defined: Integers or Strings
func (WiregoMinimalExample) GetDetectionFilters() []wirego.DetectionFilter {
	var filters []wirego.DetectionFilter

	filters = append(filters, wirego.DetectionFilter{FilterType: wirego.DetectionFilterTypeInt, Name: "udp.port", ValueInt: 50081})

	return filters
}

// GetDissectorFilterHeuristics returns a list of protocols on top of which detection heuristic
// should be called.
func (WiregoMinimalExample) GetDetectionHeuristicsParents() []string {
	//We want to apply our detection heuristic on all tcp payloads
	return []string{}
}

func (WiregoMinimalExample) DetectionHeuristic(packetNumber int, src string, dst string, layer string, packet []byte) bool {
	return false
}

// DissectPacket provides the packet payload to be parsed.
func (WiregoMinimalExample) DissectPacket(packetNumber int, src string, dst string, layer string, packet []byte) *wirego.DissectResult {
	var res wirego.DissectResult
	var info string

	offs := 0

	before := offs
	cmdCode, _ := getUInt32(packet, &offs)
	res.Fields = append(res.Fields, wirego.DissectField{WiregoFieldId: CmdCode, Offset: before, Length: 4})

	before = offs
	getUInt32(packet, &offs)
	res.Fields = append(res.Fields, wirego.DissectField{WiregoFieldId: DestCode, Offset: before, Length: 4})

	before = offs
	getUInt32(packet, &offs)
	res.Fields = append(res.Fields, wirego.DissectField{WiregoFieldId: DeskSerial, Offset: before, Length: 4})

	before = offs
	getUInt32(packet, &offs)

	res.Fields = append(res.Fields, wirego.DissectField{WiregoFieldId: RemoteSerial, Offset: before, Length: 4})

	res.Protocol = "AWS900 remote protocol"

	switch SSLCode(cmdCode) {
	case SSLMessageSendHeartbeat:
		info = parseSendHeartbeat(packet, &res, offs)
	case SSLMessageGetIsChanStereo:
		info = parseGetIsChanStereo(packet, &res, offs)
	case SSLMessageGetIsChanStereoReply:
		info = parseGetIsChanStereoReply(packet, &res, offs)
	case SSLMessageGetDeskReply:
		info = parseGetDeskReply(packet, &res, offs)
	case SSLMessageGetDirectoryList:
		info = parseGetDirectoryList(packet, &res, offs)
	case SSLMessageGetDirectoryListReply:
		info = parseGetDirectoryListReply(packet, &res, offs)
	case SSLMessageSendDiskInfo:
		info = parseSendDiskInfo(packet, &res, offs)
	case SSLMessageGetExtNamesReply:
		info = parseGetExtNamesReply(packet, &res, offs)
	case SSLMessageGetChanNamesAndImages:
		info = parseGetChanNamesAndImages(packet, &res, offs)
	case SSLMessageGetChanNamesAndImagesReply:
		info = parseGetChanNamesAndImagesReply(packet, &res, offs)
	case SSLMessageGetProjectNameAndTitleReply:
		info = parseGetProjectNameAndTitleReply(packet, &res, offs)
	case SSLMessageSendRequestFileBlock:
		info = parseSendRequestFileBlock(packet, &res, offs)
	case SSLMessageAckRequestFileBlock:
		info = parseAckRequestFileBlock(packet, &res, offs)
	case SSLMessageGetMixPassesListReply:
		info = parseGetMixPassesListReply(packet, &res, offs)
	case SSLMessageGetTrListReply:
		info = parseGetTrListReply(packet, &res, offs)
	case SSLMessageSetChanNames:
		info = parseSetChanNames(packet, &res, offs)
	case SSLMessageSetChanNamesReply:
		info = parseSetChanNames(packet, &res, offs)
	default:
		if offs != len(packet) {
			commandFields := wirego.DissectField{WiregoFieldId: CmdFields, Offset: offs, Length: len(packet) - offs}
			res.Fields = append(res.Fields, commandFields)
			info = fmt.Sprintf("TODO (%d/%d parsed)", offs, len(packet))
		}

	}

	res.Info = fmt.Sprintf("Message %s %s", SSLCodeToString(SSLCode(cmdCode)), info)

	return &res
}
