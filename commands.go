package chip

// known CHiP commands, thanks to
// https://github.com/adamgreen/CHiP/blob/master/CHiP-BLE-Protocol.md
const (
	Action             = 0x07
	Drive              = 0x78
	GetAlarmDateTime   = 0x4A
	GetBatteryLevel    = 0x1C
	GetCurrentDateTime = 0x3A
	GetEyeBrightness   = 0x49
	GetSpeed           = 0x46
	GetVolume          = 0x16
	PlaySound          = 0x06
	SetAlarmDateTime   = 0x44
	SetCurrentDateTime = 0x43
	SetEyeBrightness   = 0x48
	SetSpeed           = 0x45
	SetVolume          = 0x18
	Sleep              = 0xFA
	Version            = 0x14

	// commands for MiP might they work for CHiP?
	// ClapDelayTime          = 0x20
	// ClapEnabled            = 0x1E
	// ClapTimes              = 0x1D
	// ContinousDrive         = 0x78
	// Detected               = 0x04
	// DetectionMode          = 0x0E
	// DisconnectApp          = 0xFE
	// DistanceDrive          = 0x70
	// DriveBackwardTime      = 0x72
	// DriveForwardTime       = 0x71
	// FlashChestLED          = 0x89
	// ForceBLEDisconnect     = 0xFC
	// GestureDetect          = 0x0A
	// GetGameMode            = 0x82
	// GetHardwareInfo        = 0x19
	// GetRadarMode           = 0x0D
	// GetSoftwareVersion     = 0x14
	// GetUp                  = 0x23
	// GetUserData            = 0x13
	// IRRemoteEnabled        = 0x10
	// RadarResponse          = 0x0C
	// ReadOdometer           = 0x85
	// ReceiveIRDongleCode    = 0x03
	// RequestChestLED        = 0x83
	// RequestClapEnabled     = 0x1F
	// RequestDetectionMode   = 0x0F
	// RequestHeadLED         = 0x8B
	// RequestIRRemoteEnabled = 0x11
	// RequestStatus          = 0x79
	// RequestWeightUpdate    = 0x81
	// ResetOdometer          = 0x86
	// SendIRDongleCode       = 0x8C
	// SetChestLED            = 0x84
	// SetGameMode            = 0x76
	// SetGestureRadarMode    = 0x0C
	// SetHeadLED             = 0x8a
	// SetPosition            = 0x08
	// SetUserData            = 0x12
	// ShakeDetected          = 0x1A
	// Stop                   = 0x77
	// TurnLeftAngle          = 0x73
	// TurnRightAngle         = 0x74
)

type ActionType uint8

const (
	ActionReset                  ActionType = 0x01
	ActionTypeit                 ActionType = 0x02
	ActionLieDown                ActionType = 0x03
	ActionAllIdleMode            ActionType = 0x04
	ActionDance                  ActionType = 0x05
	ActionVRTraining1            ActionType = 0x06
	ActionVRTraining2            ActionType = 0x07
	ActionReset2                 ActionType = 0x08
	ActionJump                   ActionType = 0x09
	ActionYoga                   ActionType = 0x0a
	ActionWatchCome              ActionType = 0x0b
	ActionWatchFollow            ActionType = 0x0c
	ActionWatchFetch             ActionType = 0x0d
	ActionBallTracking           ActionType = 0x0e
	ActionBallSoccer             ActionType = 0x0f
	ActionBase                   ActionType = 0x10
	ActionDanceBase              ActionType = 0x11
	ActionTypetopOrStandFromBase ActionType = 0x12
	ActionGuardMode              ActionType = 0x13
	ActionFreeRoam               ActionType = 0x14
	ActionFaceDown               ActionType = 0x15
)
