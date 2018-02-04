package protocol

/**
 * Message type present in X224 packet header
 */

type messageType struct {
X224_TPDU_CONNECTION_REQUEST byte
X224_TPDU_CONNECTION_CONFIRM byte
X224_TPDU_DISCONNECT_REQUEST byte
X224_TPDU_DATA byte
X224_TPDU_ERROR byte
}

var MessageType messageType = messageType { 0xE0, 0xD0,0x80, 0xF0, 0x70 }


/**
 * Type of negotiation present in negotiation packet
 */
var NegotiationType = {
TYPE_RDP_NEG_REQ : 0x01,
TYPE_RDP_NEG_RSP : 0x02,
TYPE_RDP_NEG_FAILURE : 0x03
};

/**
 * Protocols available for x224 layer
 */
var Protocols = {
PROTOCOL_RDP : 0x00000000,
PROTOCOL_SSL : 0x00000001,
PROTOCOL_HYBRID : 0x00000002,
PROTOCOL_HYBRID_EX : 0x00000008
};

/**
 * Use to negotiate security layer of RDP stack
 * In node-rdpjs only ssl is available
 * @param opt {object} component type options
 * @see request -> http://msdn.microsoft.com/en-us/library/cc240500.aspx
 * @see response -> http://msdn.microsoft.com/en-us/library/cc240506.aspx
 * @see failure ->http://msdn.microsoft.com/en-us/library/cc240507.aspx
 */
function negotiation(opt) {
var self = {
type : new type.UInt8(),
flag : new type.UInt8(),
length : new type.UInt16Le(0x0008, { constant : true }),
result : new type.UInt32Le()
};
return new type.Component(self, opt);
}

/**
 * X224 client connection request
 * @param opt {object} component type options
 * @see	http://msdn.microsoft.com/en-us/library/cc240470.aspx
 */
function clientConnectionRequestPDU(opt, cookie) {
var self = {
len : new type.UInt8(function() {
return new type.Component(self).size() - 1;
}),
code : new type.UInt8(MessageType.X224_TPDU_CONNECTION_REQUEST, { constant : true }),
padding : new type.Component([new type.UInt16Le(), new type.UInt16Le(), new type.UInt8()]),
cookie : cookie || new type.Factory( function (s) {
var offset = 0;
while (true) {
var token = s.buffer.readUInt16LE(s.offset + offset);
if (token === 0x0a0d) {
self.cookie = new type.BinaryString(null, { readLength : new type.CallableValue(offset + 2) }).read(s);
return;
}
else {
offset += 1;
}
}
}, { conditional : function () {
return self.len.value > 14;
}}),
protocolNeg : negotiation({ optional : true })
};

return new type.Component(self, opt);
}

/**
 * X224 Server connection confirm
 * @param opt {object} component type options
 * @see	http://msdn.microsoft.com/en-us/library/cc240506.aspx
 */
function serverConnectionConfirm(opt) {
var self = {
len : new type.UInt8(function() {
return new type.Component(self).size() - 1;
}),
code : new type.UInt8(MessageType.X224_TPDU_CONNECTION_CONFIRM, { constant : true }),
padding : new type.Component([new type.UInt16Le(), new type.UInt16Le(), new type.UInt8()]),
protocolNeg : negotiation({ optional : true })
};

return new type.Component(self, opt);
}

/**
 * Header of each data message from x224 layer
 * @returns {type.Component}
 */
function x224DataHeader() {
var self = {
header : new type.UInt8(2),
messageType : new type.UInt8(MessageType.X224_TPDU_DATA, { constant : true }),
separator : new type.UInt8(0x80, { constant : true })
};
return new type.Component(self);
}

/**
 * Common X224 Automata
 * @param presentation {Layer} presentation layer
 */
function X224(transport) {
this.transport = transport;
this.requestedProtocol = Protocols.PROTOCOL_SSL;
this.selectedProtocol = Protocols.PROTOCOL_SSL;

var self = this;
this.transport.on('close', function() {
self.emit('close');
}).on('error', function (err) {
self.emit('error', err);
});
}

//inherit from Layer
inherits(X224, events.EventEmitter);

/**
 * Main data received function
 * after connection sequence
 * @param s {type.Stream} stream formated from transport layer
 */
X224.prototype.recvData = function(s) {
// check header
x224DataHeader().read(s);
this.emit('data', s);
};

/**
 * Format message from x224 layer to transport layer
 * @param message {type}
 * @returns {type.Component} x224 formated message
 */
X224.prototype.send = function(message) {
this.transport.send(new type.Component([x224DataHeader(), message]));
};

/**
 * Client x224 automata
 * @param transport {events.EventEmitter} (bind data events)
 */
function Client(transport) {
X224.call(this, transport);

}

//inherit from X224 automata
inherits(Client, X224);

/**
 * Client automata connect event
 */
Client.prototype.connect = function() {
var message = clientConnectionRequestPDU(null, new type.BinaryString());
message.obj.protocolNeg.obj.type.value = NegotiationType.TYPE_RDP_NEG_REQ;
message.obj.protocolNeg.obj.result.value = this.requestedProtocol;
this.transport.send(message);

// next state wait connection confirm packet
var self = this;
this.transport.once('data', function(s) {
self.recvConnectionConfirm(s);
});
};

/**
 * close stack
 */
Client.prototype.close = function() {
this.transport.close();
};

/**
 * Receive connection from server
 * @param s {Stream}
 */
Client.prototype.recvConnectionConfirm = function(s) {
var message = serverConnectionConfirm().read(s);

if(message.obj.protocolNeg.obj.type.value == NegotiationType.TYPE_RDP_NEG_FAILURE) {
throw new error.ProtocolError('NODE_RDP_PROTOCOL_X224_NEG_FAILURE');
}

if(message.obj.protocolNeg.obj.type.value == NegotiationType.TYPE_RDP_NEG_RSP) {
this.selectedProtocol = message.obj.protocolNeg.obj.result.value;
}

if([Protocols.PROTOCOL_HYBRID, Protocols.PROTOCOL_HYBRID_EX].indexOf(this.selectedProtocol) !== -1) {
throw new error.ProtocolError('NODE_RDP_PROTOCOL_X224_NLA_NOT_SUPPORTED');
}

if(this.selectedProtocol == Protocols.PROTOCOL_RDP) {
log.warn("RDP standard security selected");
return;
}

// finish connection sequence
var self = this;
this.transport.on('data', function(s) {
self.recvData(s);
});

if(this.selectedProtocol == Protocols.PROTOCOL_SSL) {
log.info("SSL standard security selected");
this.transport.transport.startTLS(function() {
self.emit('connect', self.selectedProtocol);
});
return;
}
};

/**
 * Server x224 automata
 */
function Server(transport, keyFilePath, crtFilePath) {
X224.call(this, transport);
this.keyFilePath = keyFilePath;
this.crtFilePath = crtFilePath;
var self = this;
this.transport.once('data', function (s) {
self.recvConnectionRequest(s);
});
}

//inherit from X224 automata
inherits(Server, X224);

/**
 * @see http://msdn.microsoft.com/en-us/library/cc240470.aspx
 * @param s {type.Stream}
 */
Server.prototype.recvConnectionRequest = function (s) {
var request = clientConnectionRequestPDU().read(s);
if (!request.obj.protocolNeg.isReaded) {
throw new Error('NODE_RDP_PROTOCOL_X224_NO_BASIC_SECURITY_LAYER');
}

this.requestedProtocol = request.obj.protocolNeg.obj.result.value;
this.selectedProtocol = this.requestedProtocol & Protocols.PROTOCOL_SSL;

if (!(this.selectedProtocol & Protocols.PROTOCOL_SSL)) {
var confirm = serverConnectionConfirm();
confirm.obj.protocolNeg.obj.type.value = NegociationType.TYPE_RDP_NEG_FAILURE;
confirm.obj.protocolNeg.obj.result.value = NegotiationFailureCode.SSL_REQUIRED_BY_SERVER;
this.transport.send(confirm);
this.close();
}
else {
this.sendConnectionConfirm();
}
};

/**
 * Start SSL connection if needed
 * @see http://msdn.microsoft.com/en-us/library/cc240501.aspx
 */
Server.prototype.sendConnectionConfirm = function () {
var confirm = serverConnectionConfirm();
confirm.obj.protocolNeg.obj.type.value = NegotiationType.TYPE_RDP_NEG_RSP;
confirm.obj.protocolNeg.obj.result.value = this.selectedProtocol;
this.transport.send(confirm);

// finish connection sequence
var self = this;
this.transport.on('data', function(s) {
self.recvData(s);
});

this.transport.transport.listenTLS(this.keyFilePath, this.crtFilePath, function() {
log.info('start SSL connection');
self.emit('connect', self.requestedProtocol);
});
};

/**
 * Module exports
 */
module.exports = {
Client : Client,
Server : Server
};