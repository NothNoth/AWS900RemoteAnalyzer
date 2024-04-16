package main

import (
	"encoding/binary"
	"errors"
	"fmt"
	"strings"

	"github.com/quarkslab/wirego/wirego/wirego"
)

// Define here enum identifiers, used to refer to a specific field
const (
	//AWS900 remote common fields
	CmdCode      wirego.FieldId = 1
	DestCode     wirego.FieldId = 2
	DeskSerial   wirego.FieldId = 3
	RemoteSerial wirego.FieldId = 4

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
	CmdResponseChanNamesAndImages_ChannelName wirego.FieldId = 83

	CmdResponseGetProjectNameAndTitle_ProjectName wirego.FieldId = 90
	CmdResponseGetProjectNameAndTitle_ProjectInfo wirego.FieldId = 91
	CmdResponseGetProjectNameAndTitle_TitleName   wirego.FieldId = 92
	CmdResponseGetProjectNameAndTitle_TitleInfo   wirego.FieldId = 93

	CmdRequestFileBlock_BlockPosition wirego.FieldId = 101
	CmdRequestFileBlock_BlockLen      wirego.FieldId = 102
	CmdRequestFileBlock_Zero          wirego.FieldId = 103
	CmdRequestFileBlock_FileData      wirego.FieldId = 104

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
	fields = append(fields, wirego.WiresharkField{WiregoFieldId: CmdResponseChanNamesAndImages_ChannelName, Name: "Channel name", Filter: "aws900.chan_name", ValueType: wirego.ValueTypeCString, DisplayMode: wirego.DisplayModeNone})

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

func parseSendHeartbeat(packet []byte, res *wirego.DissectResult, offs int) string {
	commandFields := wirego.DissectField{WiregoFieldId: CmdFields, Offset: offs, Length: len(packet) - offs}

	before := offs
	productName, _ := getString(packet, &offs)
	productNameField := wirego.DissectField{WiregoFieldId: CmdResponseGetDeskReply_ProductName, Offset: before, Length: len(productName) + 1}
	commandFields.SubFields = append(commandFields.SubFields, productNameField)

	res.Fields = append(res.Fields, commandFields)
	return productName
}

func parseGetIsChanStereo(packet []byte, res *wirego.DissectResult, offs int) string {
	commandFields := wirego.DissectField{WiregoFieldId: CmdFields, Offset: offs, Length: len(packet) - offs}

	before := offs
	chanNumber, _ := getByte(packet, &offs)
	chanNumberField := wirego.DissectField{WiregoFieldId: SSLRemote_ChannelNumber, Offset: before, Length: 1}
	commandFields.SubFields = append(commandFields.SubFields, chanNumberField)

	res.Fields = append(res.Fields, commandFields)
	return fmt.Sprintf("on channel %d", chanNumber)
}

func parseGetIsChanStereoReply(packet []byte, res *wirego.DissectResult, offs int) string {
	commandFields := wirego.DissectField{WiregoFieldId: CmdFields, Offset: offs, Length: len(packet) - offs}

	before := offs
	chanNumber, _ := getByte(packet, &offs)
	chanNumberField := wirego.DissectField{WiregoFieldId: SSLRemote_ChannelNumber, Offset: before, Length: 1}
	commandFields.SubFields = append(commandFields.SubFields, chanNumberField)

	before = offs
	result, _ := getByte(packet, &offs)
	resultField := wirego.DissectField{WiregoFieldId: CmdResponseIsChanStereo_IsStereo, Offset: before, Length: 1}
	commandFields.SubFields = append(commandFields.SubFields, resultField)

	res.Fields = append(res.Fields, commandFields)
	return fmt.Sprintf("for channel %d: %d", chanNumber, result)
}

func parseGetDeskReply(packet []byte, res *wirego.DissectResult, offs int) string {
	commandFields := wirego.DissectField{WiregoFieldId: CmdFields, Offset: offs, Length: len(packet) - offs}

	before := offs
	productName, _ := getString(packet, &offs)
	productNameField := wirego.DissectField{WiregoFieldId: CmdResponseGetDeskReply_ProductName, Offset: before, Length: len(productName) + 1}
	commandFields.SubFields = append(commandFields.SubFields, productNameField)

	before = offs
	versionMaj, _ := getUInt32(packet, &offs)
	versionMajField := wirego.DissectField{WiregoFieldId: CmdResponseGetDeskReply_VersionMajor, Offset: before, Length: 4}
	commandFields.SubFields = append(commandFields.SubFields, versionMajField)

	before = offs
	versionMin, _ := getUInt32(packet, &offs)
	versionMinField := wirego.DissectField{WiregoFieldId: CmdResponseGetDeskReply_VersionMinor, Offset: before, Length: 4}
	commandFields.SubFields = append(commandFields.SubFields, versionMinField)

	before = offs
	issue, _ := getUInt32(packet, &offs)
	issueField := wirego.DissectField{WiregoFieldId: CmdResponseGetDeskReply_Issue, Offset: before, Length: 4}
	commandFields.SubFields = append(commandFields.SubFields, issueField)

	buildVersionStart := offs
	buildVersion, _ := getString(packet, &offs)
	buildVersionField := wirego.DissectField{WiregoFieldId: CmdResponseGetDeskReply_BuildVersion, Offset: buildVersionStart, Length: len(buildVersion) + 1}
	commandFields.SubFields = append(commandFields.SubFields, buildVersionField)

	buildDateStart := offs
	buildDate, _ := getString(packet, &offs)
	buildDateField := wirego.DissectField{WiregoFieldId: CmdResponseGetDeskReply_BuildDate, Offset: buildDateStart, Length: len(buildDate) + 1}
	commandFields.SubFields = append(commandFields.SubFields, buildDateField)

	res.Fields = append(res.Fields, commandFields)

	return fmt.Sprintf("%s v%d.%d-%d", productName, versionMaj, versionMin, issue)
}

func parseGetDirectoryList(packet []byte, res *wirego.DissectResult, offs int) string {
	commandFields := wirego.DissectField{WiregoFieldId: CmdFields, Offset: offs, Length: len(packet) - offs}

	before := offs
	getByte(packet, &offs)
	directoryCountField := wirego.DissectField{WiregoFieldId: CmdRequestDirectoryList_Mode, Offset: before, Length: 1}
	commandFields.SubFields = append(commandFields.SubFields, directoryCountField)

	before = offs
	directoryName, _ := getString(packet, &offs)
	directoryNameField := wirego.DissectField{WiregoFieldId: SSLRemote_DirectoryName, Offset: before, Length: len(directoryName) + 1}
	commandFields.SubFields = append(commandFields.SubFields, directoryNameField)

	res.Fields = append(res.Fields, commandFields)
	return directoryName
}

func parseGetDirectoryListReply(packet []byte, res *wirego.DissectResult, offs int) string {
	commandFields := wirego.DissectField{WiregoFieldId: CmdFields, Offset: offs, Length: len(packet) - offs}

	//Directory Name
	before := offs
	directoryName, _ := getString(packet, &offs)
	dirNameField := wirego.DissectField{WiregoFieldId: SSLRemote_DirectoryName, Offset: before, Length: len(directoryName) + 1}
	commandFields.SubFields = append(commandFields.SubFields, dirNameField)

	for {
		if offs >= len(packet) {
			break
		}

		//Index
		before = offs
		getUInt16(packet, &offs)
		indexField := wirego.DissectField{WiregoFieldId: CmdResponseDirectoryList_Index, Offset: before, Length: 2}

		//Name
		before = offs
		name, _ := getString(packet, &offs)
		nameField := wirego.DissectField{WiregoFieldId: CmdResponseDirectoryList_Name, Offset: before, Length: len(name) + 1}

		//Info
		before = offs
		info, _ := getString(packet, &offs)
		infoField := wirego.DissectField{WiregoFieldId: CmdResponseDirectoryList_Info, Offset: before, Length: len(info) + 1}

		//IsDir
		getByte(packet, &offs)
		isDirField := wirego.DissectField{WiregoFieldId: SSLRemote_IsDir, Offset: before, Length: 1}

		//Time
		before = offs
		time, _ := getString(packet, &offs)
		timeField := wirego.DissectField{WiregoFieldId: SSLRemote_Time, Offset: before, Length: len(time) + 1}

		//Date
		before = offs
		date, _ := getString(packet, &offs)
		dateField := wirego.DissectField{WiregoFieldId: SSLRemote_Date, Offset: before, Length: len(date) + 1}

		//Size
		before = offs
		getUInt32(packet, &offs)
		sizeField := wirego.DissectField{WiregoFieldId: SSLRemote_Size, Offset: before, Length: 4}

		commandFields.SubFields = append(commandFields.SubFields, []wirego.DissectField{indexField, nameField, infoField, isDirField, timeField, dateField, sizeField}...)
	}

	res.Fields = append(res.Fields, commandFields)
	return directoryName
}

func parseSendDiskInfo(packet []byte, res *wirego.DissectResult, offs int) string {
	commandFields := wirego.DissectField{WiregoFieldId: CmdFields, Offset: offs, Length: len(packet) - offs}

	before := offs
	free, _ := getUInt32(packet, &offs)
	freeField := wirego.DissectField{WiregoFieldId: CmdResponseSendDiskInfo_Free, Offset: before, Length: 4}
	commandFields.SubFields = append(commandFields.SubFields, freeField)

	before = offs
	size, _ := getUInt32(packet, &offs)
	sizeField := wirego.DissectField{WiregoFieldId: SSLRemote_Size, Offset: before, Length: 4}
	commandFields.SubFields = append(commandFields.SubFields, sizeField)

	before = offs
	getByte(packet, &offs)
	archiveDoneField := wirego.DissectField{WiregoFieldId: CmdResponseSendDiskInfo_ArchiveDone, Offset: before, Length: 1}
	commandFields.SubFields = append(commandFields.SubFields, archiveDoneField)

	res.Fields = append(res.Fields, commandFields)
	return fmt.Sprintf("Free space %d%%", free*100/size)
}

func parseGetExtNamesReply(packet []byte, res *wirego.DissectResult, offs int) string {
	commandFields := wirego.DissectField{WiregoFieldId: CmdFields, Offset: offs, Length: len(packet) - offs}

	allStr := ""
	for {
		ext, err := getByte(packet, &offs)
		if ext == 0x00 || err != nil {
			break
		}

		before := offs
		number, _ := getByte(packet, &offs)
		numberField := wirego.DissectField{WiregoFieldId: CmdResponseExtNames_Number, Offset: before, Length: 1}
		commandFields.SubFields = append(commandFields.SubFields, numberField)

		//Date
		before = offs
		name, _ := getString(packet, &offs)
		nameField := wirego.DissectField{WiregoFieldId: CmdResponseExtNames_Name, Offset: before, Length: len(name) + 1}
		commandFields.SubFields = append(commandFields.SubFields, nameField)

		allStr += fmt.Sprintf("%d:%s ", number, name)
	}
	res.Fields = append(res.Fields, commandFields)
	return allStr
}

func parseGetChanNamesAndImages(packet []byte, res *wirego.DissectResult, offs int) string {
	commandFields := wirego.DissectField{WiregoFieldId: CmdFields, Offset: offs, Length: len(packet) - offs}

	before := offs
	first, _ := getByte(packet, &offs)
	firstField := wirego.DissectField{WiregoFieldId: CmdRequestChanNamesAndImages_FirstChannel, Offset: before, Length: 1}
	commandFields.SubFields = append(commandFields.SubFields, firstField)

	before = offs
	last, _ := getByte(packet, &offs)
	lastField := wirego.DissectField{WiregoFieldId: CmdRequestChanNamesAndImages_LastChannel, Offset: before, Length: 1}
	commandFields.SubFields = append(commandFields.SubFields, lastField)

	res.Fields = append(res.Fields, commandFields)
	return fmt.Sprintf("from %d to %d", first, last)
}

func parseGetChanNamesAndImagesReply(packet []byte, res *wirego.DissectResult, offs int) string {
	commandFields := wirego.DissectField{WiregoFieldId: CmdFields, Offset: offs, Length: len(packet) - offs}

	nameAll := ""
	for {
		if offs >= len(packet) {
			break
		}

		before := offs
		channel, _ := getByte(packet, &offs)
		if channel == 0 {
			break
		}
		commandFields.SubFields = append(commandFields.SubFields, wirego.DissectField{WiregoFieldId: SSLRemote_ChannelNumber, Offset: before, Length: 1})

		before = offs
		chanName, _ := getString(packet, &offs)
		commandFields.SubFields = append(commandFields.SubFields, wirego.DissectField{WiregoFieldId: CmdResponseChanNamesAndImages_ChannelName, Offset: before, Length: len(chanName) + 1})
		nameAll += chanName + "|"
	}

	res.Fields = append(res.Fields, commandFields)
	return nameAll
}

func parseGetProjectNameAndTitleReply(packet []byte, res *wirego.DissectResult, offs int) string {
	commandFields := wirego.DissectField{WiregoFieldId: CmdFields, Offset: offs, Length: len(packet) - offs}

	before := offs
	projectName, _ := getString(packet, &offs)
	commandFields.SubFields = append(commandFields.SubFields, wirego.DissectField{WiregoFieldId: CmdResponseGetProjectNameAndTitle_ProjectName, Offset: before, Length: len(projectName) + 1})

	before = offs
	projectInfo, _ := getString(packet, &offs)
	commandFields.SubFields = append(commandFields.SubFields, wirego.DissectField{WiregoFieldId: CmdResponseGetProjectNameAndTitle_ProjectInfo, Offset: before, Length: len(projectInfo) + 1})

	before = offs
	titleName, _ := getString(packet, &offs)
	commandFields.SubFields = append(commandFields.SubFields, wirego.DissectField{WiregoFieldId: CmdResponseGetProjectNameAndTitle_TitleName, Offset: before, Length: len(titleName) + 1})

	before = offs
	titleInfo, _ := getString(packet, &offs)
	commandFields.SubFields = append(commandFields.SubFields, wirego.DissectField{WiregoFieldId: CmdResponseGetProjectNameAndTitle_TitleInfo, Offset: before, Length: len(titleInfo) + 1})

	res.Fields = append(res.Fields, commandFields)

	return strings.TrimSuffix(projectInfo, "\n") + " / " + strings.TrimSuffix(titleInfo, "\n")
}

func parseSendRequestFileBlock(packet []byte, res *wirego.DissectResult, offs int) string {
	commandFields := wirego.DissectField{WiregoFieldId: CmdFields, Offset: offs, Length: len(packet) - offs}

	before := offs
	path, _ := getString(packet, &offs)
	commandFields.SubFields = append(commandFields.SubFields, wirego.DissectField{WiregoFieldId: SSLRemote_Path, Offset: before, Length: len(path) + 1})

	before = offs
	blockPos, _ := getUInt32(packet, &offs)
	commandFields.SubFields = append(commandFields.SubFields, wirego.DissectField{WiregoFieldId: CmdRequestFileBlock_BlockPosition, Offset: before, Length: 4})

	before = offs
	blockLen, _ := getUInt32(packet, &offs)
	commandFields.SubFields = append(commandFields.SubFields, wirego.DissectField{WiregoFieldId: CmdRequestFileBlock_BlockLen, Offset: before, Length: 4})

	before = offs
	getUInt32(packet, &offs)
	commandFields.SubFields = append(commandFields.SubFields, wirego.DissectField{WiregoFieldId: CmdRequestFileBlock_Zero, Offset: before, Length: 4})

	res.Fields = append(res.Fields, commandFields)

	return fmt.Sprintf("read %s: %d bytes at %d", path, blockLen, blockPos)
}

func parseAckRequestFileBlock(packet []byte, res *wirego.DissectResult, offs int) string {
	commandFields := wirego.DissectField{WiregoFieldId: CmdFields, Offset: offs, Length: len(packet) - offs}

	before := offs
	path, _ := getString(packet, &offs)
	commandFields.SubFields = append(commandFields.SubFields, wirego.DissectField{WiregoFieldId: SSLRemote_Path, Offset: before, Length: len(path) + 1})

	before = offs
	getUInt32(packet, &offs)
	commandFields.SubFields = append(commandFields.SubFields, wirego.DissectField{WiregoFieldId: CmdRequestFileBlock_BlockPosition, Offset: before, Length: 4})

	before = offs
	getUInt32(packet, &offs)
	commandFields.SubFields = append(commandFields.SubFields, wirego.DissectField{WiregoFieldId: CmdRequestFileBlock_BlockLen, Offset: before, Length: 4})

	commandFields.SubFields = append(commandFields.SubFields, wirego.DissectField{WiregoFieldId: CmdRequestFileBlock_FileData, Offset: offs, Length: len(packet) - offs})

	res.Fields = append(res.Fields, commandFields)

	return fmt.Sprintf("%s (%d bytes of data)", path, len(packet)-offs)
}

func parseGetMixPassesListReply(packet []byte, res *wirego.DissectResult, offs int) string {
	commandFields := wirego.DissectField{WiregoFieldId: CmdFields, Offset: offs, Length: len(packet) - offs}

	before := offs
	dirPath, _ := getString(packet, &offs)
	commandFields.SubFields = append(commandFields.SubFields, wirego.DissectField{WiregoFieldId: SSLRemote_Path, Offset: before, Length: len(dirPath) + 1})

	count := 0
	for {
		before = offs
		fileIndex, _ := getUInt16(packet, &offs)
		commandFields.SubFields = append(commandFields.SubFields, wirego.DissectField{WiregoFieldId: SSLRemote_FileIndex, Offset: before, Length: 2})
		if fileIndex == 0 {
			break
		}

		count++
		before = offs
		name, _ := getString(packet, &offs)
		commandFields.SubFields = append(commandFields.SubFields, wirego.DissectField{WiregoFieldId: SSLRemote_MixPassName, Offset: before, Length: len(name) + 1})

		before = offs
		getUInt32(packet, &offs)
		commandFields.SubFields = append(commandFields.SubFields, wirego.DissectField{WiregoFieldId: SSLRemote_MixPassNumber, Offset: before, Length: 4})

		before = offs
		getUInt32(packet, &offs)
		commandFields.SubFields = append(commandFields.SubFields, wirego.DissectField{WiregoFieldId: SSLRemote_MixPassBasedOn, Offset: before, Length: 4})

		before = offs
		timeStr, _ := getString(packet, &offs)
		commandFields.SubFields = append(commandFields.SubFields, wirego.DissectField{WiregoFieldId: SSLRemote_Time, Offset: before, Length: len(timeStr) + 1})

		before = offs
		dateStr, _ := getString(packet, &offs)
		commandFields.SubFields = append(commandFields.SubFields, wirego.DissectField{WiregoFieldId: SSLRemote_Date, Offset: before, Length: len(dateStr) + 1})

		before = offs
		getUInt32(packet, &offs)
		commandFields.SubFields = append(commandFields.SubFields, wirego.DissectField{WiregoFieldId: SSLRemote_Size, Offset: before, Length: 4})
	}
	res.Fields = append(res.Fields, commandFields)

	return fmt.Sprintf("%s: %d mix passes", dirPath, count)
}

func parseGetTrListReply(packet []byte, res *wirego.DissectResult, offs int) string {
	commandFields := wirego.DissectField{WiregoFieldId: CmdFields, Offset: offs, Length: len(packet) - offs}

	before := offs
	dirPath, _ := getString(packet, &offs)
	commandFields.SubFields = append(commandFields.SubFields, wirego.DissectField{WiregoFieldId: SSLRemote_Path, Offset: before, Length: len(dirPath) + 1})

	count := 0
	for {
		before = offs
		fileIndex, _ := getUInt16(packet, &offs)
		commandFields.SubFields = append(commandFields.SubFields, wirego.DissectField{WiregoFieldId: SSLRemote_FileIndex, Offset: before, Length: 2})
		if fileIndex == 0 {
			break
		}
		count++

		before = offs
		name, _ := getString(packet, &offs)
		commandFields.SubFields = append(commandFields.SubFields, wirego.DissectField{WiregoFieldId: SSLRemote_TRName, Offset: before, Length: len(name) + 1})

		before = offs
		info, _ := getString(packet, &offs)
		commandFields.SubFields = append(commandFields.SubFields, wirego.DissectField{WiregoFieldId: SSLRemote_TRInfo, Offset: before, Length: len(info) + 1})

		before = offs
		getByte(packet, &offs)
		commandFields.SubFields = append(commandFields.SubFields, wirego.DissectField{WiregoFieldId: SSLRemote_IsDir, Offset: before, Length: 1})

		before = offs
		timeStr, _ := getString(packet, &offs)
		commandFields.SubFields = append(commandFields.SubFields, wirego.DissectField{WiregoFieldId: SSLRemote_Time, Offset: before, Length: len(timeStr) + 1})

		before = offs
		dateStr, _ := getString(packet, &offs)
		commandFields.SubFields = append(commandFields.SubFields, wirego.DissectField{WiregoFieldId: SSLRemote_Date, Offset: before, Length: len(dateStr) + 1})

		before = offs
		getUInt32(packet, &offs)
		commandFields.SubFields = append(commandFields.SubFields, wirego.DissectField{WiregoFieldId: SSLRemote_Size, Offset: before, Length: 4})

		before = offs
		getByte(packet, &offs)
		commandFields.SubFields = append(commandFields.SubFields, wirego.DissectField{WiregoFieldId: SSLRemote_TRIsSelected, Offset: before, Length: 1})
	}

	res.Fields = append(res.Fields, commandFields)
	return fmt.Sprintf("%s: %d TR", dirPath, count)
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

func getString(packet []byte, offs *int) (string, error) {
	var str string
	var err error

	ends := *offs
	for {
		if *offs >= len(packet) {
			return "", errors.New("out of bounds")
		}
		if packet[ends] == 0x00 {
			break
		}
		ends++
	}
	str = string(packet[*offs:ends])
	*offs = ends + 1
	return str, err
}

func getUInt32(packet []byte, offs *int) (uint32, error) {

	if (*offs)+3 >= len(packet) {
		return 0, errors.New("out of bounds")
	}
	size := binary.BigEndian.Uint32(packet[*offs : (*offs)+4])
	(*offs) += 4
	return size, nil
}

func getUInt16(packet []byte, offs *int) (uint16, error) {
	if *offs+1 >= len(packet) {
		return 0, errors.New("out of bounds")
	}
	size := binary.BigEndian.Uint16(packet[*offs : (*offs)+2])
	(*offs) += 2
	return size, nil
}

func getByte(packet []byte, offs *int) (byte, error) {
	if *offs >= len(packet) {
		return 0, errors.New("out of bounds")
	}
	b := packet[*offs]
	(*offs) += 1
	return b, nil
}
