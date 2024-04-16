package main

import (
	"fmt"
	"strings"

	"github.com/quarkslab/wirego/wirego/wirego"
)

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
		commandFields.SubFields = append(commandFields.SubFields, wirego.DissectField{WiregoFieldId: SSLRemote_ChannelName, Offset: before, Length: len(chanName) + 1})
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

	return strings.ReplaceAll(projectInfo, "\n", "-") + " / " + strings.ReplaceAll(titleInfo, "\n", "-")
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

func parseSetChanNames(packet []byte, res *wirego.DissectResult, offs int) string {
	commandFields := wirego.DissectField{WiregoFieldId: CmdFields, Offset: offs, Length: len(packet) - offs}
	allNames := ""

	for {
		before := offs
		chanNumber, _ := getByte(packet, &offs)
		commandFields.SubFields = append(commandFields.SubFields, wirego.DissectField{WiregoFieldId: SSLRemote_ChannelNumber, Offset: before, Length: 1})
		if chanNumber == 0 {
			break
		}

		before = offs
		name, _ := getString(packet, &offs)
		commandFields.SubFields = append(commandFields.SubFields, wirego.DissectField{WiregoFieldId: SSLRemote_ChannelName, Offset: before, Length: len(name) + 1})

		allNames += fmt.Sprintf("%d:%s ", chanNumber, name)
	}
	res.Fields = append(res.Fields, commandFields)
	return allNames
}
