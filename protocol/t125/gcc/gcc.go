package gcc

import  (
	"../../../core"
	"../per"
	"errors"
)

var t124_02_98_oid []uint8{0, 0, 20, 124, 0, 1 }
var h221_cs_key string = "Duca"
var h221_sc_key string = "McDn"


/**
 * @see http://msdn.microsoft.com/en-us/library/cc240509.aspx
 */
type Message uint16
const (
//server -> client
	SC_CORE Message = 0x0C01
	SC_SECURITY = 0x0C02
	SC_NET = 0x0C03
	//client -> server
	CS_CORE = 0xC001
	CS_SECURITY = 0xC002
	CS_NET = 0xC003
	CS_CLUSTER = 0xC004
	CS_MONITOR = 0xC005
)

/**
 * @see http://msdn.microsoft.com/en-us/library/cc240510.aspx
 */
type ColorDepth uint16
const (
	RNS_UD_COLOR_8BPP ColorDepth = 0xCA01
	RNS_UD_COLOR_16BPP_555 = 0xCA02
	RNS_UD_COLOR_16BPP_565 = 0xCA03
	RNS_UD_COLOR_24BPP = 0xCA04
)

/**
 * @see http://msdn.microsoft.com/en-us/library/cc240510.aspx
 */
type HighColor uint16
const (
	HIGH_COLOR_4BPP HighColor = 0x0004
	HIGH_COLOR_8BPP = 0x0008
	HIGH_COLOR_15BPP = 0x000f
	HIGH_COLOR_16BPP = 0x0010
	HIGH_COLOR_24BPP = 0x0018
)

/**
 * @see http://msdn.microsoft.com/en-us/library/cc240510.aspx
*/
type Support uint16
const (
	RNS_UD_24BPP_SUPPORT uint16 = 0x0001
	RNS_UD_16BPP_SUPPORT = 0x0002
	RNS_UD_15BPP_SUPPORT = 0x0004
	RNS_UD_32BPP_SUPPORT = 0x0008
)

/**
 * @see http://msdn.microsoft.com/en-us/library/cc240510.aspx
*/
type CapabilityFlag uint16

const (
	RNS_UD_CS_SUPPORT_ERRINFO_PDU uint16 = 0x0001
	RNS_UD_CS_WANT_32BPP_SESSION = 0x0002
	RNS_UD_CS_SUPPORT_STATUSINFO_PDU = 0x0004
	RNS_UD_CS_STRONG_ASYMMETRIC_KEYS = 0x0008
	RNS_UD_CS_UNUSED = 0x0010
	RNS_UD_CS_VALID_CONNECTION_TYPE = 0x0020
	RNS_UD_CS_SUPPORT_MONITOR_LAYOUT_PDU = 0x0040
	RNS_UD_CS_SUPPORT_NETCHAR_AUTODETECT = 0x0080
	RNS_UD_CS_SUPPORT_DYNVC_GFX_PROTOCOL = 0x0100
	RNS_UD_CS_SUPPORT_DYNAMIC_TIME_ZONE = 0x0200
	RNS_UD_CS_SUPPORT_HEARTBEAT_PDU = 0x0400
)

/**
 * @see http://msdn.microsoft.com/en-us/library/cc240510.aspx
*/
type ConnectionType uint8
const (
CONNECTION_TYPE_MODEM ConnectionType =  0x01
CONNECTION_TYPE_BROADBAND_LOW = 0x02
CONNECTION_TYPE_SATELLITEV = 0x03
CONNECTION_TYPE_BROADBAND_HIGH = 0x04
CONNECTION_TYPE_WAN = 0x05
CONNECTION_TYPE_LAN = 0x06
CONNECTION_TYPE_AUTODETECT = 0x07
)

/**
 * @see http://msdn.microsoft.com/en-us/library/cc240510.aspx
 */
type VERSION uint32
const (
RDP_VERSION_4 VERSION = 0x00080001
RDP_VERSION_5_PLUS = 0x00080004
)
type Sequence uint16
const (
	RNS_UD_SAS_DEL Sequence =  0xAA03
)

/**
 * @see http://msdn.microsoft.com/en-us/library/cc240511.aspx
 */
type EncryptionMethod uint32
const (
ENCRYPTION_FLAG_40BIT uint32 = 0x00000001
ENCRYPTION_FLAG_128BIT = 0x00000002
ENCRYPTION_FLAG_56BIT = 0x00000008
FIPS_ENCRYPTION_FLAG = 0x00000010
)

/**
 * @see http://msdn.microsoft.com/en-us/library/cc240518.aspx
 */
type EncryptionLevel uint32
const (
	ENCRYPTION_LEVEL_NONE EncryptionLevel = 0x00000000
	ENCRYPTION_LEVEL_LOW = 0x00000001
	ENCRYPTION_LEVEL_CLIENT_COMPATIBLE = 0x00000002
	ENCRYPTION_LEVEL_HIGH = 0x00000003
	ENCRYPTION_LEVEL_FIPS = 0x00000004
)

/**
 * @see http://msdn.microsoft.com/en-us/library/cc240513.aspx
 */
type ChannelOptions uint32
const (
	CHANNEL_OPTION_INITIALIZED ChannelOptions =  0x80000000
	CHANNEL_OPTION_ENCRYPT_RDP = 0x40000000
	CHANNEL_OPTION_ENCRYPT_SC = 0x20000000
	CHANNEL_OPTION_ENCRYPT_CS = 0x10000000
	CHANNEL_OPTION_PRI_HIGH = 0x08000000
	CHANNEL_OPTION_PRI_MED = 0x04000000
	CHANNEL_OPTION_PRI_LOW = 0x02000000
	CHANNEL_OPTION_COMPRESS_RDP = 0x00800000
	CHANNEL_OPTION_COMPRESS = 0x00400000
	CHANNEL_OPTION_SHOW_PROTOCOL = 0x00200000
	REMOTE_CONTROL_PERSISTENT = 0x00100000
)

/**
 * IBM_101_102_KEYS is the most common keyboard type
 */
type KeyboardType uint32
const (
	KT_IBM_PC_XT_83_KEY KeyboardType = 0x00000001
	KT_OLIVETTI = 0x00000002
	KT_IBM_PC_AT_84_KEY = 0x00000003
	KT_IBM_101_102_KEYS = 0x00000004
	KT_NOKIA_1050 = 0x00000005
	KT_NOKIA_9140 = 0x00000006
	KT_JAPANESE = 0x00000007
)

/**
 * @see http://technet.microsoft.com/en-us/library/cc766503%28WS.10%29.aspx
 */
type KeyboardLayout uint32
const (
	ARABIC KeyboardLayout = 0x00000401
	BULGARIAN = 0x00000402
	CHINESE_US_KEYBOARD = 0x00000404
	CZECH = 0x00000405
	DANISH = 0x00000406
	GERMAN = 0x00000407
	GREEK = 0x00000408
	US = 0x00000409
	SPANISH = 0x0000040a
	FINNISH = 0x0000040b
	FRENCH = 0x0000040c
	HEBREW = 0x0000040d
	HUNGARIAN = 0x0000040e
	ICELANDIC = 0x0000040f
	ITALIAN = 0x00000410
	JAPANESE = 0x00000411
	KOREAN = 0x00000412
	DUTCH = 0x00000413
	NORWEGIAN = 0x00000414
)

/**
 * @see http://msdn.microsoft.com/en-us/library/cc240521.aspx
 */
type CertificateType uint32
const (
	CERT_CHAIN_VERSION_1 CertificateType = 0x00000001
	CERT_CHAIN_VERSION_2 = 0x00000002
)

/**
 * @param {type.Type} data
 * @returns {type.Component}
*/
type Block struct {
	Type uint16
	Length uint16
	Data core.Data
}
/*
func NewGccBlock(data []byte) *GccBlock{
	block := GccBlock{}
var self = {
// type of data block
type : new type.UInt16Le(function() {
return self.data.obj.__TYPE__;
}),
// length of entire packet
length : new type.UInt16Le(function() {
return new type.Component(self).size();
}),
// data block
data : data || new type.Factory(function(s){
var options = {
readLength : new type.CallableValue( function () {
return self.length.value - 4;
})
};
switch(self.type.value) {
case MessageType.SC_CORE:
self.data = serverCoreData(options).read(s);
break;
case MessageType.SC_SECURITY:
self.data = serverSecurityData(options).read(s);
break;
case MessageType.SC_NET:
self.data = serverNetworkData(null, options).read(s);
break;
case MessageType.CS_CORE:
self.data = clientCoreData(options).read(s);
break;
case MessageType.CS_SECURITY:
self.data = clientSecurityData(options).read(s);
break;
case MessageType.CS_NET:
self.data = clientNetworkData(null, options).read(s);
break;
default:
log.debug("unknown gcc block type " + self.type.value);
self.data = new type.BinaryString(null, options).read(s);
}
})
};

return new type.Component(self);
}*/

/**
 * Main client informations
 * 	keyboard
 * 	screen definition
 * 	color depth
 * @see http://msdn.microsoft.com/en-us/library/cc240510.aspx
 * @param opt {object} Classic type options
 * @returns {type.Component}
*/

/* JS
rdpVersion : new type.UInt32Le(VERSION.RDP_VERSION_5_PLUS),
desktopWidth : new type.UInt16Le(1280),
desktopHeight : new type.UInt16Le(800),
colorDepth : new type.UInt16Le(ColorDepth.RNS_UD_COLOR_8BPP),
sasSequence : new type.UInt16Le(Sequence.RNS_UD_SAS_DEL),
kbdLayout : new type.UInt32Le(KeyboardLayout.FRENCH),
clientBuild : new type.UInt32Le(3790),
clientName : new type.BinaryString(new Buffer('node-rdpjs\x00\x00\x00\x00\x00\x00', 'ucs2'), { readLength : new type.CallableValue(32) }),
keyboardType : new type.UInt32Le(KeyboardType.IBM_101_102_KEYS),
keyboardSubType : new type.UInt32Le(0),
keyboardFnKeys : new type.UInt32Le(12),
imeFileName : new type.BinaryString(new Buffer(Array(64 + 1).join('\x00')), { readLength : new type.CallableValue(64), optional : true }),
postBeta2ColorDepth : new type.UInt16Le(ColorDepth.RNS_UD_COLOR_8BPP, { optional : true }),
clientProductId : new type.UInt16Le(1, { optional : true }),
serialNumber : new type.UInt32Le(0, { optional : true }),
highColorDepth : new type.UInt16Le(HighColor.HIGH_COLOR_24BPP, { optional : true }),
supportedColorDepths : new type.UInt16Le(Support.RNS_UD_15BPP_SUPPORT | Support.RNS_UD_16BPP_SUPPORT | Support.RNS_UD_24BPP_SUPPORT | Support.RNS_UD_32BPP_SUPPORT, { optional : true }),
earlyCapabilityFlags : new type.UInt16Le(CapabilityFlag.RNS_UD_CS_SUPPORT_ERRINFO_PDU, { optional : true }),
clientDigProductId : new type.BinaryString(new Buffer(Array(64 + 1).join('\x00')), { optional : true, readLength : new type.CallableValue(64) }),
connectionType : new type.UInt8(0, { optional : true }),
pad1octet : new type.UInt8(0, { optional : true }),
serverSelectedProtocol : new type.UInt32Le(0, { optional : true })
 */
type ClientCoreData struct {
	RdpVersion VERSION
	DesktopWidth uint16
	DesktopHeight uint16
	ColorDepth ColorDepth
	SasSequence Sequence
	KbdLayout KeyboardLayout
	ClientBuild uint32
	ClientName [32]byte
	KeyboardType uint32
	KeyboardSubType uint32
	KeyboardFnKeys uint32
	ImeFileName [64]byte
	PostBeta2ColorDepth ColorDepth //optional
	ClientProductId uint16  //optional
	SerialNumber uint32  //optional
	HighColorDepth HighColor  //optional
	SupportedColorDepths uint16 //optional
	EarlyCapabilityFlags uint16 //optional
	ClientDigProductId [64]byte //optional
	ConnectionType uint8 //optional
	Pad1octet uint8 //optional
	ServerSelectedProtocol uint32 //optional
}

func NewClientCoreData() *ClientCoreData {
	return &ClientCoreData{
		 RDP_VERSION_5_PLUS, 1280, 800,RNS_UD_COLOR_8BPP,
		RNS_UD_SAS_DEL, US, 3790, [32]byte{'g', 'o', 'r', 'd', 'p'},KT_IBM_101_102_KEYS,
		0,12, [64] byte{}, RNS_UD_COLOR_8BPP,1,0, HIGH_COLOR_24BPP,
		RNS_UD_15BPP_SUPPORT | RNS_UD_16BPP_SUPPORT | RNS_UD_24BPP_SUPPORT | RNS_UD_32BPP_SUPPORT,
		RNS_UD_CS_SUPPORT_ERRINFO_PDU, [64] byte{}, 0, 0, 0}
	}
}

func (c *ClientCoreData) GetType() Message {
	return CS_CORE
}

func (data * ClientCoreData)Write(writer core.Writer) error {
	core.WriteUInt32LE(uint32(data.RdpVersion), writer)
	core.WriteUInt16LE(data.DesktopWidth, writer)
	core.WriteUInt16LE(data.DesktopHeight, writer)
	core.WriteUInt16LE(uint16(data.ColorDepth), writer)
	core.WriteUInt16LE(uint16(data.SasSequence), writer)
	core.WriteUInt32LE(uint32(data.KbdLayout), writer)
	core.WriteUInt32LE(data.ClientBuild, writer)
	core.WriteBytes(data.ClientName[:], writer)
	core.WriteUInt32LE(data.KeyboardType, writer)
	core.WriteUInt32LE(data.KeyboardFnKeys, writer)
	core.WriteBytes(data.ImeFileName[:], writer)
	core.WriteUInt16LE(uint16(data.PostBeta2ColorDepth), writer)
	core.WriteUInt16LE(data.ClientProductId, writer)
	core.WriteUInt32LE(data.SerialNumber, writer)
	core.WriteUInt16LE(uint16(data.HighColorDepth), writer)
	core.WriteUInt16LE(data.SupportedColorDepths, writer)
	core.WriteUInt16LE(data.EarlyCapabilityFlags, writer)
	core.WriteBytes(data.ClientDigProductId[:], writer)
	core.WriteUInt8(data.ConnectionType, writer)
	core.WriteUInt8(data.Pad1octet, writer)
	core.WriteUInt32LE(data.ServerSelectedProtocol, writer)
	return nil
}

func (data * ClientCoreData) Read(reader core.Reader) error {
	data.RdpVersion = VERSION(core.ReadUInt32LE(reader))
	data.DesktopWidth = core.ReadUInt16LE(reader)
	data.DesktopHeight = core.ReadUInt16LE(reader)
	data.ColorDepth = ColorDepth(core.ReadUInt16LE(reader))
	data.SasSequence = Sequence(core.ReadUInt16LE(reader))
	data.KbdLayout = KeyboardLayout(core.ReadUInt32LE(reader))
	data.ClientBuild = core.ReadUInt32LE(reader)
	core.ReadBytes(data.ClientName[:], reader)
	data.KeyboardType = core.ReadUInt32LE(reader)
	data.KeyboardFnKeys = core.ReadUInt32LE(reader)
	core.ReadBytes(data.ImeFileName[:], reader)
	//optional
	data.PostBeta2ColorDepth = ColorDepth(core.ReadUInt16LE(reader))
	data.ClientProductId = core.ReadUInt16LE(reader)
	data.SerialNumber = core.ReadUInt32LE(reader)
	data.HighColorDepth = HighColor(core.ReadUInt16LE(reader))
	data.SupportedColorDepths = core.ReadUInt16LE(reader)
	data.EarlyCapabilityFlags = core.ReadUInt16LE(reader)
	core.ReadBytes(data.ClientDigProductId[:], reader)
	data.ConnectionType = core.ReadUInt8(reader)
	data.Pad1octet = core.ReadUInt8(reader)
	data.ServerSelectedProtocol = core.ReadUInt32LE(reader)
	return nil
}

/**
 * @see http://msdn.microsoft.com/en-us/library/cc240517.aspx
 * @param opt {object} Classic type options
 * @returns {type.Component}
*/
/*
function serverCoreData(opt) {
var self = {
__TYPE__ : MessageType.SC_CORE,
rdpVersion : new type.UInt32Le(VERSION.RDP_VERSION_5_PLUS),
clientRequestedProtocol : new type.UInt32Le(null, { optional : true }),
earlyCapabilityFlags : new type.UInt32Le(null, { optional : true })
};

return new type.Component(self, opt);
}*/

type ServerCoreData struct {
	RdpVersion VERSION
	ClientRequestedProtocol uint32 //optional
	EarlyCapabilityFlags uint32 //optional
}

func NewServerCoreData() *ServerCoreData {
	return &ServerCoreData{
		RDP_VERSION_5_PLUS, 0, 0}
}

func (c *ServerCoreData) GetType() Message {
	return SC_CORE
}


func (data * ServerCoreData)Write(writer core.Writer) error {
	core.WriteUInt32LE(uint32(data.RdpVersion), writer)
	core.WriteUInt32LE(data.ClientRequestedProtocol, writer)
	core.WriteUInt32LE(data.EarlyCapabilityFlags, writer)
	return nil
}

func (data * ServerCoreData) Read(reader core.Reader) error {
	data.RdpVersion = VERSION(core.ReadUInt32LE(reader))
	data.ClientRequestedProtocol = core.ReadUInt32LE(reader)
	data.EarlyCapabilityFlags = core.ReadUInt32LE(reader)
	return nil
}

/**
 * @see http://msdn.microsoft.com/en-us/library/cc240511.aspx
 * @param opt {object} Classic type options
 * @returns {type.Component}
*/
/*function clientSecurityData(opt) {
var self = {
__TYPE__ : MessageType.CS_SECURITY,
encryptionMethods : new type.UInt32Le(EncryptionMethod.ENCRYPTION_FLAG_40BIT | EncryptionMethod.ENCRYPTION_FLAG_56BIT | EncryptionMethod.ENCRYPTION_FLAG_128BIT),
extEncryptionMethods : new type.UInt32Le()
};

return new type.Component(self, opt);
}*/

type ClientSecurityData struct {
	EncryptionMethods uint32
	ExtEncryptionMethods uint32
}

func NewClientSecurityData() *ClientSecurityData {
	return &ClientSecurityData{
		ENCRYPTION_FLAG_40BIT | ENCRYPTION_FLAG_56BIT | ENCRYPTION_FLAG_128BIT,
		00}
}

func (c *ClientSecurityData) GetType() Message {
	return CS_SECURITY
}


func (data * ClientSecurityData)Write(writer core.Writer) error {
	core.WriteUInt32LE(data.EncryptionMethods, writer)
	core.WriteUInt32LE(data.ExtEncryptionMethods, writer)
	return nil
}

func (data * ClientSecurityData) Read(reader core.Reader) error {
	data.EncryptionMethods = core.ReadUInt32LE(reader)
	data.ExtEncryptionMethods = core.ReadUInt32LE(reader)
	return nil
}

/**
 * Only use for SSL (RDP security layer TODO)
 * @see http://msdn.microsoft.com/en-us/library/cc240518.aspx
 * @param opt {object} Classic type options
 * @returns {type.Component}
*/
/*function serverSecurityData(opt) {
var self = {
__TYPE__ : MessageType.SC_SECURITY,
encryptionMethod : new type.UInt32Le(),
encryptionLevel : new type.UInt32Le()
};

return new type.Component(self, opt);
}*/

type ServerSecurityData struct {
	EncryptionMethod uint32
	EncryptionLevel uint32
}

func NewServerSecurityData() *ServerSecurityData {
	return &ServerSecurityData{
		0, 0}
}

func (c *ServerSecurityData) GetType() Message {
	return SC_SECURITY
}


func (data * ServerSecurityData)Write(writer core.Writer) error {
	core.WriteUInt32LE(data.EncryptionMethod, writer)
	core.WriteUInt32LE(data.EncryptionLevel, writer)
	return nil
}

func (data * ServerSecurityData) Read(reader core.Reader) error {
	data.EncryptionMethod = core.ReadUInt32LE(reader)
	data.EncryptionLevel = core.ReadUInt32LE(reader)
	return nil
}

/**
 * Channel definition
 * @param opt {object} Classic type options
 * @returns {type.Component}
*/
/*function channelDef (opt) {
var self = {
name : new type.BinaryString(null, { readLength : new type.CallableValue(8) }),
options : new type.UInt32Le()
};

return new type.Component(self, opt);
}*/

type ChannelDef struct {
	Name [8] byte
	Options uint32
}

/**
 * Optional channel requests (sound, clipboard ...)
 * @param opt {object} Classic type options
 * @returns {type.Component}
*/
/*
function clientNetworkData(channelDefArray, opt) {
var self = {
__TYPE__ : MessageType.CS_NET,
channelCount : new type.UInt32Le( function () {
return self.channelDefArray.obj.length;
}),
channelDefArray : channelDefArray || new type.Factory( function (s) {
self.channelDefArray = new type.Component([]);

for (var i = 0; i < self.channelCount.value; i++) {
self.channelDefArray.obj.push(channelDef().read(s));
}
})
};

return new type.Component(self, opt);
}*/
type ClientNetworkData struct {
	ChannelDefArray []ChannelDef
}

func NewClientNetworkData() *ClientNetworkData {
	return &ClientNetworkData{}
}

func (c *ClientNetworkData) GetType() Message {
	return CS_NET
}


func (data * ClientNetworkData)Write(writer core.Writer) error {
	core.WriteUInt32LE(uint32(len(data.ChannelDefArray)), writer)
	for _, cd := range data.ChannelDefArray {
		core.WriteBytes(cd.Name[:], writer)
		core.WriteUInt32LE(cd.Options, writer)
	}

	return nil
}

func (data * ClientNetworkData) Read(reader core.Reader) error {
	length := int(core.ReadUInt32LE(reader))
	for i := 0; i < length; i++ {
		var cd ChannelDef
		core.ReadBytes(cd.Name[:], reader)
		cd.Options = core.ReadUInt32LE(reader)
		data.ChannelDefArray = append(data.ChannelDefArray, cd)
	}

	return nil
}
/**
 * @param channelIds {type.Component} list of available channels
 * @param opt {object} Classic type options
 * @returns {type.Component}
*/
/*function serverNetworkData (channelIds, opt) {
var self = {
__TYPE__ : MessageType.SC_NET,
MCSChannelId : new type.UInt16Le(1003, { constant : true }),
channelCount : new type.UInt16Le(function () {
return self.channelIdArray.obj.length;
}),
channelIdArray : channelIds || new type.Factory( function (s) {
self.channelIdArray = new type.Component([]);
for (var i = 0; i < self.channelCount.value; i++) {
self.channelIdArray.obj.push(new type.UInt16Le().read(s));
}
}),
pad : new type.UInt16Le(null, { conditional : function () {
return (self.channelCount.value % 2) === 1;
}})
};

return new type.Component(self, opt);
}*/

type ServerNetworkData struct {
	ChannelIdArray []uint16
}

func NewServerNetworkData() *ServerNetworkData {
	return &ServerNetworkData{}
}

func (c *ServerNetworkData) GetType() Message {
	return SC_NET
}


func (data * ServerNetworkData)Write(writer core.Writer) error {
	core.WriteUInt16LE(uint16(1003), writer)
	core.WriteUInt32LE(uint32(len(data.ChannelIdArray)), writer)
	for _, cd := range data.ChannelIdArray {
		core.WriteUInt16LE(cd, writer)
	}

	return nil
}

func (data * ServerNetworkData) Read(reader core.Reader) error {
	id := core.ReadUInt16LE(reader)
	if id != 1003 {
		return errors.New("must be 1003")
	}
	length := int(core.ReadUInt32LE(reader))
	for i := 0; i < length; i++ {
		var cd uint16 = core.ReadUInt16LE(reader)
		data.ChannelIdArray = append(data.ChannelIdArray, cd)
	}

	return nil
}

/**
 * Client or server GCC settings block
 * @param blocks {type.Component} array of gcc blocks
 * @param opt {object} options to component type
 * @returns {type.Component}
*/
/*
function settings(blocks, opt) {
var self = {
blocks : blocks || new type.Factory(function(s) {
self.blocks = new type.Component([]);
// read until end of stream
while(s.availableLength() > 0) {
self.blocks.obj.push(block().read(s));
}
}),
};

return new type.Component(self, opt);
}
*/
/**
 * Read GCC response from server
 * @param s {type.Stream} current stream
 * @returns {Array(type.Component)} list of server block
 */
/*function readConferenceCreateResponse(s) {
per.readChoice(s);

if(!per.readObjectIdentifier(s, t124_02_98_oid)) {
throw new error.ProtocolError('NODE_RDP_PROTOCOL_T125_GCC_BAD_OBJECT_IDENTIFIER_T124');
}

per.readLength(s);
per.readChoice(s);
per.readInteger16(s, 1001);
per.readInteger(s);
per.readEnumerates(s);
per.readNumberOfSet(s);
per.readChoice(s);

if (!per.readOctetStream(s, h221_sc_key, 4)) {
throw new error.ProtocolError('NODE_RDP_PROTOCOL_T125_GCC_BAD_H221_SC_KEY');
}

length = per.readLength(s);
serverSettings = settings(null, { readLength : new type.CallableValue(length) });

// Object magic
return serverSettings.read(s).obj.blocks.obj.map(function(e) {
return e.obj.data;
});
}
*/
/**
 * Read GCC request
 * @param s {type.Stream}
 * @returns {Array(type.Component)} list of client block
 */
/*function readConferenceCreateRequest (s) {
per.readChoice(s);
if (!per.readObjectIdentifier(s, t124_02_98_oid)) {
throw new error.ProtocolError('NODE_RDP_PROTOCOL_T125_GCC_BAD_H221_SC_KEY');
}
per.readLength(s);
per.readChoice(s);
per.readSelection(s);
per.readNumericString(s, 1);
per.readPadding(s, 1);

if (per.readNumberOfSet(s) !== 1) {
throw new error.ProtocolError('NODE_RDP_PROTOCOL_T125_GCC_BAD_SET');
}

if (per.readChoice(s) !== 0xc0) {
throw new error.ProtocolError('NODE_RDP_PROTOCOL_T125_GCC_BAD_CHOICE');
}

per.readOctetStream(s, h221_cs_key, 4);

length = per.readLength(s);
var clientSettings = settings(null, { readLength : new type.CallableValue(length) });

// Object magic
return clientSettings.read(s).obj.blocks.obj.map(function(e) {
return e.obj.data;
});
}
*/
/**
 * Built {type.Componen} from gcc user data
 * @param userData {type.Component} GCC data from client
 * @returns {type.Component} GCC encoded client user data
 */
/*function writeConferenceCreateRequest (userData) {
var userDataStream = new type.Stream(userData.size());
userData.write(userDataStream);

return new type.Component([
per.writeChoice(0), per.writeObjectIdentifier(t124_02_98_oid),
per.writeLength(userData.size() + 14), per.writeChoice(0),
per.writeSelection(0x08), per.writeNumericString("1", 1), per.writePadding(1),
per.writeNumberOfSet(1), per.writeChoice(0xc0),
per.writeOctetStream(new Buffer(h221_cs_key), 4), per.writeOctetStream(userDataStream.getValue())
]);
}

function writeConferenceCreateResponse (userData) {
var userDataStream = new type.Stream(userData.size());
userData.write(userDataStream);

return new type.Component([
per.writeChoice(0), per.writeObjectIdentifier(t124_02_98_oid),
per.writeLength(userData.size() + 14), per.writeChoice(0x14),
per.writeInteger16(0x79F3, 1001), per.writeInteger(1), per.writeEnumerates(0),
per.writeNumberOfSet(1), per.writeChoice(0xc0),
per.writeOctetStream(new Buffer(h221_sc_key), 4), per.writeOctetStream(userDataStream.getValue())
]);
}*/


