package main

import (
	"fmt"

	"github.com/quarkslab/wirego/wirego_remote/go/wirego"
)

// Define here enum identifiers, used to refer to a specific field
const (
	//AWS900 remote common fields
	CmdCode      wirego.FieldId = iota
	DestCode     wirego.FieldId = iota
	DeskSerial   wirego.FieldId = iota
	RemoteSerial wirego.FieldId = iota

	//Command specific fields
	CommandDetails wirego.FieldId = iota

	//Common fields
	SSLRemote_Path           wirego.FieldId = iota
	SSLRemote_Time           wirego.FieldId = iota
	SSLRemote_Date           wirego.FieldId = iota
	SSLRemote_Size           wirego.FieldId = iota
	SSLRemote_FileInfo       wirego.FieldId = iota
	SSLRemote_FileIndex      wirego.FieldId = iota
	SSLRemote_IsDir          wirego.FieldId = iota
	SSLRemote_ChannelNumber  wirego.FieldId = iota
	SSLRemote_DirectoryName  wirego.FieldId = iota
	SSLRemote_TRName         wirego.FieldId = iota
	SSLRemote_TRInfo         wirego.FieldId = iota
	SSLRemote_TRIsSelected   wirego.FieldId = iota
	SSLRemote_MixPassName    wirego.FieldId = iota
	SSLRemote_MixPassNumber  wirego.FieldId = iota
	SSLRemote_MixPassBasedOn wirego.FieldId = iota
	SSLRemote_ChannelName    wirego.FieldId = iota
	SSLRemote_Zero           wirego.FieldId = iota

	// Command specific fields
	CmdResponseIsChanStereo_IsStereo wirego.FieldId = iota

	CmdResponseGetDeskReply_ProductName  wirego.FieldId = iota
	CmdResponseGetDeskReply_VersionMajor wirego.FieldId = iota
	CmdResponseGetDeskReply_VersionMinor wirego.FieldId = iota
	CmdResponseGetDeskReply_Issue        wirego.FieldId = iota
	CmdResponseGetDeskReply_BuildVersion wirego.FieldId = iota
	CmdResponseGetDeskReply_BuildDate    wirego.FieldId = iota

	CmdRequestDirectoryList_Mode wirego.FieldId = iota

	CmdResponseDirectoryList_Index wirego.FieldId = iota
	CmdResponseDirectoryList_Name  wirego.FieldId = iota
	CmdResponseDirectoryList_Info  wirego.FieldId = iota

	CmdResponseSendDiskInfo_Free        wirego.FieldId = iota
	CmdResponseSendDiskInfo_ArchiveDone wirego.FieldId = iota

	CmdResponseExtNames_Number wirego.FieldId = iota
	CmdResponseExtNames_Name   wirego.FieldId = iota

	CmdRequestChanNamesAndImages_FirstChannel wirego.FieldId = iota
	CmdRequestChanNamesAndImages_LastChannel  wirego.FieldId = iota

	CmdResponseGetProjectNameAndTitle_ProjectName wirego.FieldId = iota
	CmdResponseGetProjectNameAndTitle_ProjectInfo wirego.FieldId = iota
	CmdResponseGetProjectNameAndTitle_TitleName   wirego.FieldId = iota
	CmdResponseGetProjectNameAndTitle_TitleInfo   wirego.FieldId = iota

	CmdRequestFileBlock_BlockPosition wirego.FieldId = iota
	CmdRequestFileBlock_BlockLen      wirego.FieldId = iota
	CmdRequestFileBlock_FileData      wirego.FieldId = iota

	CmdAckWriteFileBlock_Position wirego.FieldId = iota
	CmdAckWriteFileBlock_Length   wirego.FieldId = iota
	SSLRemote_TODO                wirego.FieldId = iota
)

// Since we implement the wirego.WiregoInterface we need some structure to hold it.
type WiregoMinimalExample struct {
}

// Unused (but mandatory)
func main() {
	var wge WiregoMinimalExample

	wg, err := wirego.New("ipc:///tmp/wirego0", false, wge)
	if err != nil {
		fmt.Println(err)
		return
	}
	wg.ResultsCacheEnable(false)

	wg.Listen()
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
	fields = append(fields, wirego.WiresharkField{WiregoFieldId: CommandDetails, Name: "Command arguments", Filter: "aws900.command_args", ValueType: wirego.ValueTypeNone, DisplayMode: wirego.DisplayModeNone})
	fields = append(fields, wirego.WiresharkField{WiregoFieldId: SSLRemote_TODO, Name: "UNSUPPORTED PACKET", Filter: "aws900.todo", ValueType: wirego.ValueTypeNone, DisplayMode: wirego.DisplayModeNone})

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
	fields = append(fields, wirego.WiresharkField{WiregoFieldId: CmdRequestFileBlock_FileData, Name: "File contents", Filter: "aws900.file", ValueType: wirego.ValueTypeNone, DisplayMode: wirego.DisplayModeNone})

	//Common fields
	fields = append(fields, wirego.WiresharkField{WiregoFieldId: SSLRemote_Path, Name: "Path", Filter: "aws900.path", ValueType: wirego.ValueTypeCString, DisplayMode: wirego.DisplayModeNone})
	fields = append(fields, wirego.WiresharkField{WiregoFieldId: SSLRemote_Time, Name: "Time", Filter: "aws900.time", ValueType: wirego.ValueTypeCString, DisplayMode: wirego.DisplayModeNone})
	fields = append(fields, wirego.WiresharkField{WiregoFieldId: SSLRemote_Date, Name: "Date", Filter: "aws900.date", ValueType: wirego.ValueTypeCString, DisplayMode: wirego.DisplayModeNone})
	fields = append(fields, wirego.WiresharkField{WiregoFieldId: SSLRemote_Size, Name: "Size", Filter: "aws900.size", ValueType: wirego.ValueTypeUInt32, DisplayMode: wirego.DisplayModeDecimal})
	fields = append(fields, wirego.WiresharkField{WiregoFieldId: SSLRemote_FileInfo, Name: "File info", Filter: "aws900.file_info", ValueType: wirego.ValueTypeCString, DisplayMode: wirego.DisplayModeNone})
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
	fields = append(fields, wirego.WiresharkField{WiregoFieldId: SSLRemote_Zero, Name: "Zero", Filter: "aws900.zero", ValueType: wirego.ValueTypeUInt32, DisplayMode: wirego.DisplayModeDecimal})

	fields = append(fields, wirego.WiresharkField{WiregoFieldId: CmdAckWriteFileBlock_Position, Name: "File block position", Filter: "aws900.file_block_position", ValueType: wirego.ValueTypeUInt32, DisplayMode: wirego.DisplayModeDecimal})
	fields = append(fields, wirego.WiresharkField{WiregoFieldId: CmdAckWriteFileBlock_Length, Name: "File block length", Filter: "aws900.file_block_length", ValueType: wirego.ValueTypeUInt32, DisplayMode: wirego.DisplayModeDecimal})

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
	case SSLMessageSetFileInfo:
		info = parseSetFileInfo(packet, &res, offs)
	case SSLMessageAckWriteFileBlock:
		info = parseAckWriteFileBlock(packet, &res, offs)
	case SSLMessageSendWriteFileBlock:
		info = parseSendWriteFileBlock(packet, &res, offs)
	case SSLMessageSendDeleteFile:
		info = parseSendDeleteFile(packet, &res, offs)
	case SSLMessageSendTitleDetailsChanged:
		info = parseSendTitleDetailsChanged(packet, &res, offs)
	default:
		if offs != len(packet) {
			commandFields := wirego.DissectField{WiregoFieldId: CommandDetails, Offset: offs, Length: len(packet) - offs}
			todo := wirego.DissectField{WiregoFieldId: SSLRemote_TODO, Offset: offs, Length: len(packet) - offs}
			res.Fields = append(res.Fields, commandFields)
			res.Fields = append(res.Fields, todo)
			info = fmt.Sprintf("TODO (%d/%d parsed)", offs, len(packet))
		}

	}

	res.Info = fmt.Sprintf("Message %s %s", SSLCodeToString(SSLCode(cmdCode)), info)

	return &res
}
