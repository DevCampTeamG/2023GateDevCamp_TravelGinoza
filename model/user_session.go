package model

const MENUE_MAX_STATE = 2
const STAMP_MAX_STATE = 4

type UserSessionState struct {
	Status map[string]LineBotSession `json:"user_session_state"`
}

func (uss *UserSessionState) InitUserSessionState() {
	uss.Status = map[string]LineBotSession{}
}

func (uss *UserSessionState) IsUserMenuSessionValid(userID string) bool {
	return uss.Status[userID].Ms.SessionState != 0
}

func (uss *UserSessionState) IsUserStampSessionValid(userID string) bool {
	return uss.Status[userID].Ss.SessionState != 0
}

func (uss *UserSessionState) UserMenuSessionValidToTrue(userID string) {
	uss.Status[userID] = LineBotSession{MenuSession{1}, StampSession{0}}
}

func (uss *UserSessionState) UserStampRallySessionValidToTrue(userID string) {
	uss.Status[userID] = LineBotSession{MenuSession{0}, StampSession{1}}
}
func (uss *UserSessionState) UserSessionClear(userID string) {
	uss.Status[userID] = LineBotSession{MenuSession{0}, StampSession{0}}
}

type LineBotSession struct {
	Ms MenuSession  `json:"menusession"`
	Ss StampSession `json:"stampsession"`
}

type MenuSession struct {
	SessionState int `json:"session_state"`
}

type StampSession struct {
	SessionState int `json:"session_state"`
}

func (Ms *MenuSession) SetSessionState() {
	Ms.SessionState += 1

	if Ms.SessionState > MENUE_MAX_STATE {
		Ms.SessionState = 1
	}

}

func (Ss *StampSession) SetSessionState() {
	Ss.SessionState += 1

	if Ss.SessionState > STAMP_MAX_STATE {
		Ss.SessionState = 1
	}

}

func (ms *MenuSession) GetSessionState() int {
	return ms.SessionState
}

func (ss *StampSession) GetSessionState() int {
	return ss.SessionState
}
