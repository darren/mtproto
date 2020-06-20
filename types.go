package mtproto

const (
	SESSION EventType = "session"
	MCONN   EventType = "mconn"
)

type EventType string
type Event interface {
	Type() EventType
}

// Session Events
type newsession struct {
	// If connID is zero, Manager makes new connection and assigns it the new session.
	// Otherwise, the new session is allocated to the connection of connID.
	connId int32
	//phonenumber string
	//addr        string
	//useIPv6     bool
	phone   string
	apiid   int32
	apihash string
	ip      string
	port    int
	resp    chan sessionResponse
}

type loadsession struct {
	// If connID is zero, Manager makes new connection and assigns it the loaded session.
	// Otherwise, the loaded session is allocated to the connection of connID.
	connId int32
	//phone  string
	//preferredAddr string
	policy retryPolicy
	resp   chan sessionResponse
}

type sessionResponse struct {
	connId  int32
	session *Session
	err     error
}

// Established = made + bound
type SessionEstablished struct {
	session *Session
}

type discardSession struct {
	connId    int32
	sessionId int64
	resp      chan sessionResponse
}

type SessionDiscarded struct {
	boundConnId                  int32
	discardedSessionId           int64
	discardedSessionUpdatesState *PredUpdatesState
}

// discardSession + newsession
type renewSession struct {
	sessionId int64
	phone     string
	apiID     int32
	apiHash   string
	ip        string
	port      int
	resp      chan sessionResponse
}

// discardSession + loadsession
type refreshSession struct {
	sessionId int64
	phone     string
	policy    retryPolicy
	resp      chan sessionResponse
}
type retryPolicy string

const noRetry retryPolicy = "noRetry"
const untilSuccess retryPolicy = "untilSuccess"

// Connection Events
type ConnectionOpened struct {
	mconn     *Conn
	sessionID int64
}
type sessionBound struct {
	mconn     *Conn
	sessionID int64
}
type sessionUnbound struct {
	mconn            *Conn
	unboundSessionID int64
}
type closeConnection struct {
	connId int32
	resp   chan error
}
type connectionClosed struct {
	closedConnId int32
}

// Update Event
type updateReceived struct {
	update Update
}

func (e newsession) Type() EventType         { return SESSION }
func (e loadsession) Type() EventType        { return SESSION }
func (e SessionEstablished) Type() EventType { return SESSION }
func (e renewSession) Type() EventType       { return SESSION }
func (e refreshSession) Type() EventType     { return SESSION }
func (e discardSession) Type() EventType     { return SESSION }
func (e SessionDiscarded) Type() EventType   { return SESSION }
func (e ConnectionOpened) Type() EventType   { return MCONN }
func (e sessionBound) Type() EventType       { return MCONN }
func (e sessionUnbound) Type() EventType     { return MCONN }
func (e closeConnection) Type() EventType    { return MCONN }
func (e connectionClosed) Type() EventType   { return MCONN }
func (e updateReceived) Type() EventType     { return SESSION }

//func (e newsession) SessionId() (int64)          {return 0}
//func (e loadsession) SessionId() (int64)         {return 0}
//func (e SessionEstablished) SessionId() (int64)  {return e.session.sessionID}
//func (e renewSession) SessionId() (int64)        {return e.sessionID}
//func (e refreshSession) SessionId() (int64)      {return e.sessionID}
//func (e discardSession) SessionId() (int64)   {return e.sessionID}
//func (e SessionDiscarded) SessionId() (int64) {return e.discardedSessionId}
//func (e ConnectionOpened) SessionId() (int64) {return 0}
//func (e sessionBound) SessionId() (int64) {
//	session, err := e.mconn.Session()
//	if err != nil {
//		return 0
//	}
//	return session.sessionID
//}
//func (e sessionUnbound) SessionId() (int64) 		{return e.unboundSessionID}
//func (e closeConnection) SessionId() (int64) 	{return 0}
//func (e connectionClosed) SessionId() (int64) 	{return 0}
//
//func (e newsession) ConnectionId() (int32)         {return 0}
//func (e loadsession) ConnectionId() (int32)        {return 0}
//func (e SessionEstablished) ConnectionId() (int32) {return e.connID}
//func (e renewSession) ConnectionId() (int32)       {return e.connID}
//func (e refreshSession) ConnectionId() (int32)     {return e.connID}
//func (e discardSession) ConnectionId() (int32)     {return e.connID}
//func (e SessionDiscarded) ConnectionId() (int32)   {return e.boundConnId}
//func (e ConnectionOpened) ConnectionId() (int32)   {return e.mconn.connID}
//func (e sessionBound) ConnectionId() (int32)       {return e.mconn.connID}
//func (e sessionUnbound) ConnectionId() (int32)     {return e.mconn.connID}
//func (e closeConnection) ConnectionId() (int32)    {return e.connID}
//func (e connectionClosed) ConnectionId() (int32)   {return e.closedConnId}

type Update interface {
	Predicate
	UpdateDate() int32
}

type UpdateCallback interface {
	OnUpdate(update Update)
}

//PredUpdatesState Value = 1;

//PredUpdatesTooLong UpdatesTooLong = 1;
//PredUpdateShortMessage UpdateShortMessage = 2;
//PredUpdateShortChatMessage UpdateShortChatMessage = 3;
//PredUpdateShort UpdateShort = 4;
//PredUpdatesCombined UpdatesCombined = 5;
//PredUpdates Updates = 6;
//PredUpdateShortSentMessage UpdateShortSentMessage = 7;

func (u *PredUpdatesState) UpdateDate() int32 { return u.Date }

func (u *PredUpdateShortMessage) UpdateDate() int32     { return u.Date }
func (u *PredUpdateShortChatMessage) UpdateDate() int32 { return u.Date }
func (u *PredUpdateShort) UpdateDate() int32            { return u.Date }
func (u *PredUpdates) UpdateDate() int32                { return u.Date }
func (u *PredUpdateShortSentMessage) UpdateDate() int32 { return u.Date }

func (u *PredUpdatesDifference) UpdateDate() int32      { return 0 }
func (u *PredUpdatesDifferenceSlice) UpdateDate() int32 { return 0 }

//func (u US_updates_difference) UpdateDate() int32         { return 0 }
func (u *PredUpdateNewMessage) UpdateDate() int32           { return 0 }
func (u *PredUpdateReadMessagesContents) UpdateDate() int32 { return 0 }
func (u *PredUpdateDeleteMessages) UpdateDate() int32       { return 0 }
func (u *PredUpdateNewEncryptedMessage) UpdateDate() int32  { return 0 }

func (u *PredUpdateChannel) UpdateDate() int32             { return 0 }
func (u *PredUpdateChannelMessageViews) UpdateDate() int32 { return 0 }
func (u *PredUpdateChannelTooLong) UpdateDate() int32      { return 0 }
func (u *PredUpdateReadChannelInbox) UpdateDate() int32    { return 0 }
func (u *PredUpdateReadChannelOutbox) UpdateDate() int32   { return 0 }
func (u *PredUpdateNewChannelMessage) UpdateDate() int32   { return 0 }
func (u *PredMessagesChannelMessages) UpdateDate() int32   { return 0 }
