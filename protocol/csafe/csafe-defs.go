package csafe

// Define CSAFE commands and their byte representations here

// Framing Constants
const (
	EXT_FRAME_START_BYTE = 0xF0 + iota
	FRAME_START_BYTE
	FRAME_END_BYTE
	FRAME_STUFF_BYTE
)

const FRAME_MAX_STUFF_OFFSET_BYTE = 0x03

const FRAME_FLG_LEN = 2
const EXT_FRAME_ADDR_LEN = 2
const FRAME_CHKSUM_LEN = 1

const SHORT_CMD_TYPE_MSK = 0x80
const LONG_CMD_HDR_LENGTH = 2
const LONG_CMD_BYTE_CNT_OFFSET = 1
const RSP_HDR_LENGTH = 2

const FRAME_STD_TYPE = 0
const FRAME_EXT_TYPE = 1

const DESTINATION_ADDR_HOST = 0x00
const DESTINATION_ADDR_ERG_MASTER = 0x01
const DESTINATION_ADDR_BROADCAST = 0xFF
const DESTINATION_ADDR_ERG_DEFAULT = 0xFD

const FRAME_MAXSIZE = 96
const INTERFRAMEGAP_MIN = 50 // msec
const CMDUPLIST_MAXSIZE = 10
const MEMORY_BLOCKSIZE = 64
const FORCEPLOT_BLOCKSIZE = 32
const HEARTBEAT_BLOCKSIZE = 32

/* Manufacturer Info */
const MANUFACTURE_ID = 22 // assigned by Fitlinxx for Concept2
const CLASS_ID = 2        // standard CSAFE equipment
const MODEL_NUM = 5       // PM4

const UNITS_TYPE = 0 // Metric
const SERIALNUM_DIGITS = 9

const HMS_FORMAT_CNT = 3
const YMD_FORMAT_CNT = 3
const ERRORCODE_FORMAT_CNT = 3

/* Command space partitioning for standard commands */
const CTRL_CMD_LONG_MIN = 0x01
const CFG_CMD_LONG_MIN = 0x10
const DATA_CMD_LONG_MIN = 0x20
const AUDIO_CMD_LONG_MIN = 0x40
const TEXTCFG_CMD_LONG_MIN = 0x60
const TEXTSTATUS_CMD_LONG_MIN = 0x65
const CAP_CMD_LONG_MIN = 0x70
const CSAFE_SETPMCFG_CMD = 0x76

const CTRL_CMD_SHORT_MIN = 0x80
const STATUS_CMD_SHORT_MIN = 0x91
const DATA_CMD_SHORT_MIN = 0xA0
const AUDIO_CMD_SHORT_MIN = 0xC0
const TEXTCFG_CMD_SHORT_MIN = 0xE0
const TEXTSTATUS_CMD_SHORT_MIN = 0xE5

/* Standard Short Control Commands */
type SHORT_CTRL_CMDS byte

const (
	GETSTATUS_CMD  SHORT_CTRL_CMDS = 0x80 + iota // CTRL_CMD_SHORT_MIN
	RESET_CMD                                    // 0x81
	GOIDLE_CMD                                   // 0x82
	GOHAVEID_CMD                                 // 0x83
	_                                            // 0x84
	GOINUSE_CMD                                  // 0x85
	GOFINISHED_CMD                               // 0x86
	GOREADY_CMD                                  // 0x87
	BADID_CMD                                    // 0x88
	CTRL_CMD_SHORT_MAX
)

/* Standard Short Status Commands */
type SHORT_STATUS_CMDS byte

const (
	GETVERSION_CMD      SHORT_STATUS_CMDS = 0x91 // STATUS_CMD_SHORT_MIN
	GETID_CMD                                    // 0x92
	GETUNITS_CMD                                 // 0x93
	GETSERIAL_CMD                                // 0x94
	_                                            // 0x95
	_                                            // 0x96
	_                                            // 0x97
	GETLIST_CMD                                  // 0x98
	GETUTILIZATION_CMD                           // 0x99
	GETMOTORCURRENT_CMD                          // 0x9A
	GETODOMETER_CMD                              // 0x9B
	GETERRORCODE_CMD                             // 0x9C
	GETSERVICECODE_CMD                           // 0x9D
	GETUSERCFG1_CMD                              // 0x9E
	GETUSERCFG2_CMD                              // 0x9F
	STATUS_CMD_SHORT_MAX
)

/* Standard Short Data Commands */
type SHORT_DATA_CMDS byte

const (
	GETTWORK_CMD      SHORT_DATA_CMDS = 0xA0 + iota // DATA_CMD_SHORT_MIN
	GETHORIZONTAL_CMD                               // 0xA1
	GETVERTICAL_CMD                                 // 0xA2
	GETCALORIES_CMD                                 // 0xA3
	GETPROGRAM_CMD                                  // 0xA4
	GETSPEED_CMD                                    // 0xA5
	GETPACE_CMD                                     // 0xA6
	GETCADENCE_CMD                                  // 0xA7
	GETGRADE_CMD                                    // 0xA8
	GETGEAR_CMD                                     // 0xA9
	GETUPLIST_CMD                                   // 0xAA
	GETUSERINFO_CMD                                 // 0xAB
	GETTORQUE_CMD                                   // 0xAC
	_                                               // 0xAD
	_                                               // 0xAE
	_                                               // 0xAF
	GETHRCUR_CMD                                    // 0xB0
	_                                               // 0xB1
	GETHRTZONE_CMD                                  // 0xB2
	GETMETS_CMD                                     // 0xB3
	GETPOWER_CMD                                    // 0xB4
	GETHRAVG_CMD                                    // 0xB5
	GETHRMAX_CMD                                    // 0xB6
	_                                               // 0xB7
	_                                               // 0xB8
	_                                               // 0xB9
	_                                               // 0xBA
	_                                               // 0xBB
	_                                               // 0xBC
	_                                               // 0xBD
	GETUSERDATA1_CMD                                // 0xBE
	GETUSERDATA2_CMD                                // 0xBF
	DATA_CMD_SHORT_MAX
)

/* Standard Short Audio Commands */
type SHORT_AUDIO_CMDS byte

const (
	GETAUDIOCHANNEL_CMD SHORT_AUDIO_CMDS = 0xC0 // AUDIO_CMD_SHORT_MIN
	GETAUDIOVOLUME_CMD                          // 0xC1
	GETAUDIOMUTE_CMD                            // 0xC2
	AUDIO_CMD_SHORT_MAX
)

/* Standard Short Text Configuration Commands */
type SHORT_TEXTCFG_CMDS byte

const (
	ENDTEXT_CMD      SHORT_TEXTCFG_CMDS = 0xE0 + iota // TEXTCFG_CMD_SHORT_MIN
	DISPLAYPOPUP_CMD                                  // 0xE1
	TEXTCFG_CMD_SHORT_MAX
)

/* Standard Short Text Status Commands */
type SHORT_TEXTSTATUS_CMDS byte

const (
	GETPOPUPSTATUS_CMD SHORT_TEXTSTATUS_CMDS = 0xE5 + iota // TEXTSTATUS_CMD_SHORT_MIN
	TEXTSTATUS_CMD_SHORT_MAX
)

/* Standard Long commands */

/* Standard Long Control Commands */
type LONG_CTRL_CMDS byte

const (
	AUTOUPLOAD_CMD  LONG_CTRL_CMDS = 0x01 + iota // CTRL_CMD_LONG_MIN
	UPLIST_CMD                                   // 0x02
	_                                            // 0x03
	UPSTATUSSEC_CMD                              // 0x04
	UPLISTSEC_CMD                                // 0x05
	CTRL_CMD_LONG_MAX
)

/* Standard Long Configuration Commands */
type LONG_CFG_CMDS byte

const (
	IDDIGITS_CMD    LONG_CFG_CMDS = 0x10 + iota //  CFG_CMD_LONG_MIN
	SETTIME_CMD                                 // 0x11
	SETDATE_CMD                                 // 0x12
	SETTIMEOUT_CMD                              // 0x13
	_                                           // 0x14
	_                                           // 0x15
	_                                           // 0x16
	_                                           // 0x17
	_                                           // 0x18
	_                                           // 0x19
	SETUSERCFG1_CMD                             // 0x1A
	SETUSERCFG2_CMD                             // 0x1B
	CFG_CMD_LONG_MAX
)

/* Standard Long Data Commands */
type LONG_DATA_CMDS byte

const (
	SETTWORK_CMD      LONG_DATA_CMDS = 0x20 + iota // DATA_CMD_LONG_MIN
	SETHORIZONTAL_CMD                              // 0x21
	SETVERTICAL_CMD                                // 0x22
	SETCALORIES_CMD                                // 0x23
	SETPROGRAM_CMD                                 // 0x24
	SETSPEED_CMD                                   // 0x25
	_                                              // 0x26
	_                                              // 0x27
	SETGRADE_CMD                                   // 0x28
	SETGEAR_CMD                                    // 0x29
	_                                              // 0x2A
	SETUSERINFO_CMD                                // 0x2B
	SETTORQUE_CMD                                  // 0x2C
	SETLEVEL_CMD                                   // 0x2D
	_                                              // 0x2E
	_                                              // 0x2F
	SETTARGETHR_CMD                                // 0x30
	_                                              // 0x31
	SETGOAL_CMD                                    // 0x32
	SETMETS_CMD                                    // 0x33
	SETPOWER_CMD                                   // 0x34
	SETHRZONE_CMD                                  // 0x35
	SETHRMAX_CMD                                   // 0x36
	DATA_CMD_LONG_MAX
)

/* Standard Long Audio Commands */
type LONG_AUDIO_CMDS byte

const (
	SETCHANNELRANGE_CMD LONG_AUDIO_CMDS = 0x40 + iota // AUDIO_CMD_LONG_MIN
	SETVOLUMERANGE_CMD                                // 0x41
	SETAUDIOMUTE_CMD                                  // 0x42
	SETAUDIOCHANNEL_CMD                               // 0x43
	SETAUDIOVOLUME_CMD                                // 0x44
	AUDIO_CMD_LONG_MAX
)

/* Standard Long Text Configuration Commands */
type LONG_TEXTCFG_CMDS byte

const (
	STARTTEXT_CMD  LONG_TEXTCFG_CMDS = 0x60 + iota // TEXTCFG_CMD_LONG_MIN
	APPENDTEXT_CMD                                 // 0x61
	TEXTCFG_CMD_LONG_MAX
)

/* Standard Long Text Status Commands */
type LONG_TEXTSTATUS_CMDS byte

const (
	GETTEXTSTATUS_CMD LONG_TEXTSTATUS_CMDS = 0x65 + iota //  TEXTSTATUS_CMD_LONG_MIN
	TEXTSTATUS_CMD_LONG_MAX
)

/* Standard Long Capabilities Commands */
type LONG_CAP_CMDS byte

const GETCAPS_CMD LONG_CAP_CMDS = 0x70 //  CAP_CMD_LONG_MIN
const (
	GETUSERCAPS1_CMD LONG_CAP_CMDS = 0x7E + iota // 0x7E
	GETUSERCAPS2_CMD                             // 0x7F
	CAP_CMD_LONG_MAX
)

/*
   The currently defined CSAFE command space is augmented by adding 4 command
   wrappers to allow pushing and pulling of configuration/data from the
   host to the PM
   SETPMCFG_CMD    Push configuration from host to PM
   SETPMDATA_CMD   Push data from host to PM
   GETPMCFG_CMD    Pull configuration to host from PM
   GETPMDATA_CMD   PUll data to host from PM
   Note: These commands have been added for Concept 2 and do not comply
   with the existing CSAFE command set
*/
type LONG_PMPROPRIETARY_CMDS byte

const (
	SETPMCFG_CMD  LONG_PMPROPRIETARY_CMDS = 0x76 + iota //   PMPROPRIETARY_CMD_LONG_MIN
	SETPMDATA_CMD                                       // 0x77
)
const (
	GETPMCFG_CMD  LONG_PMPROPRIETARY_CMDS = 0x7E + iota // 0x7E
	GETPMDATA_CMD                                       // 0x7F
	PMPROPRIETARY_CMD_LONG_MAX
)

/* Command space partitioning for PM proprietary commands */
const GETPMCFG_CMD_SHORT_MIN = 0x80
const GETPMCFG_CMD_LONG_MIN = 0x50
const SETPMCFG_CMD_SHORT_MIN = 0xE0
const SETPMCFG_CMD_LONG_MIN = 0x00
const GETPMDATA_CMD_SHORT_MIN = 0xA0
const GETPMDATA_CMD_LONG_MIN = 0x68
const SETPMDATA_CMD_SHORT_MIN = 0xD0
const SETPMDATA_CMD_LONG_MIN = 0x30

/*
   Custom Short PULL Configuration Commands for PM
*/
type PM_SHORT_PULL_CFG_CMDS byte

const (
	PM_GET_FW_VERSION                    PM_SHORT_PULL_CFG_CMDS = 0x80 + iota // GETPMCFG_CMD_SHORT_MIN
	PM_GET_HW_VERSION                                                         // 0x81
	PM_GET_HW_ADDRESS                                                         // 0x82
	PM_GET_TICK_TIMEBASE                                                      // 0x83
	PM_GET_HRM                                                                // 0x84
	_                                                                         // Unused 0x85
	PM_GET_SCREENSTATESTATUS                                                  // 0x86
	PM_GET_RACE_LANE_REQUEST                                                  // 0x87
	PM_GET_ERG_LOGICALADDR_REQUEST                                            // 0x88
	PM_GET_WORKOUTTYPE                                                        // 0x89
	PM_GET_DISPLAYTYPE                                                        // 0x8A
	PM_GET_DISPLAYUNITS                                                       // 0x8B
	PM_GET_LANGUAGETYPE                                                       // 0x8C
	PM_GET_WORKOUTSTATE                                                       // 0x8D
	PM_GET_INTERVALTYPE                                                       // 0x8E
	PM_GET_OPERATIONALSTATE                                                   // 0x8F
	PM_GET_LOGCARDSTATE                                                       // 0x90
	PM_GET_LOGCARDSTATUS                                                      // 0x91
	PM_GET_POWERUPSTATE                                                       // 0x92
	PM_GET_ROWINGSTATE                                                        // 0x93
	PM_GET_SCREENCONTENT_VERSION                                              // 0x94
	PM_GET_COMMUNICATIONSTATE                                                 // 0x95
	PM_GET_RACEPARTICIPANTCOUNT                                               // 0x96
	PM_GET_BATTERYLEVELPERCENT                                                // 0x97
	PM_GET_RACEMODESTATUS                                                     // 0x98
	PM_GET_INTERNALLOGPARAMS                                                  // 0x99
	PM_GET_PRODUCTCONFIGURATION                                               // 0x9A
	PM_GET_ERGSLAVEDISCOVERREQUESTSTATUS                                      // 0x9B
	PM_GET_WIFICONFIG                                                         // 0x9C
	PM_GET_CPUTICKRATE                                                        // 0x9D
	PM_GET_LOGCARDCENSUS                                                      // 0x9E
	PM_GET_WORKOUTINTERVALCOUNT                                               // 0x9F
	GETPMCFG_CMD_SHORT_MAX
)

/*
   Custom Short PULL Data Commands for PM
*/
type PM_SHORT_PULL_DATA_CMDS byte

const (
	PM_GET_WORKTIME                  PM_SHORT_PULL_DATA_CMDS = 0xA0 + iota // GETPMDATA_CMD_SHORT_MIN
	PM_GET_PROJECTED_WORKTIME                                              // 0xA1
	PM_GET_TOTAL_RESTTIME                                                  // 0xA2
	PM_GET_WORKDISTANCE                                                    // 0xA3
	PM_GET_TOTAL_WORKDISTANCE                                              // 0xA4
	PM_GET_PROJECTED_WORKDISTANCE                                          // 0xA5
	PM_GET_RESTDISTANCE                                                    // 0xA6
	PM_GET_TOTAL_RESTDISTANCE                                              // 0xA7
	PM_GET_STROKE_500MPACE                                                 // 0xA8
	PM_GET_STROKE_POWER                                                    // 0xA9
	PM_GET_STROKE_CALORICBURNRATE                                          // 0xAA
	PM_GET_SPLIT_AVG_500MPACE                                              // 0xAB
	PM_GET_SPLIT_AVG_POWER                                                 // 0xAC
	PM_GET_SPLIT_AVG_CALORICBURNRATE                                       // 0xAD
	PM_GET_SPLIT_AVG_CALORIES                                              // 0xAE
	PM_GET_TOTAL_AVG_500MPACE                                              // 0xAF
	PM_GET_TOTAL_AVG_POWER                                                 // 0xB0
	PM_GET_TOTAL_AVG_CALORICBURNRATE                                       // 0xB1
	PM_GET_TOTAL_AVG_CALORIES                                              // 0xB2
	PM_GET_STROKERATE                                                      // 0xB3
	PM_GET_SPLIT_AVG_STROKERATE                                            // 0xB4
	PM_GET_TOTAL_AVG_STROKERATE                                            // 0xB5
	PM_GET_AVG_HEARTRATE                                                   // 0xB6
	PM_GET_ENDING_AVG_HEARTRATE                                            // 0xB7
	PM_GET_REST_AVG_HEARTRATE                                              // 0xB8
	PM_GET_SPLITTIME                                                       // 0xB9
	PM_GET_LASTSPLITTIME                                                   // 0xBA
	PM_GET_SPLITDISTANCE                                                   // 0xBB
	PM_GET_LASTSPLITDISTANCE                                               // 0xBC
	PM_GET_LASTRESTDISTANCE                                                // 0xBD
	PM_GET_TARGETPACETIME                                                  // 0xBE
	PM_GET_STROKESTATE                                                     // 0xBF
	PM_GET_STROKERATESTATE                                                 // 0xC0
	PM_GET_DRAGFACTOR                                                      // 0xC1
	PM_GET_ENCODERPERIOD                                                   // 0xC2
	PM_GET_HEARTRATESTATE                                                  // 0xC3
	PM_GET_SYNCDATA                                                        // 0xC4
	PM_GET_SYNCDATAALL                                                     // 0xC5
	PM_GET_RACEDATA                                                        // 0xC6
	PM_GET_TICKTIME                                                        // 0xC7
	PM_GET_ERRORTYPE                                                       // 0xC8
	PM_GET_ERRORVALUE                                                      // 0xC9
	PM_GET_STATUSTYPE                                                      // 0xCA
	PM_GET_STATUSVALUE                                                     // 0xCB
	PM_GET_EPMSTATUS                                                       // 0xCC
	PM_GET_DISPLAYUPDATETIME                                               // 0xCD
	PM_GET_SYNCFRACTIONALTIME                                              // 0xCE
	PM_GET_RESTTIME                                                        // 0xCF
	GETPMDATA_CMD_SHORT_MAX
)

/*
   Custom Short PUSH Data Commands for PM
*/
type PM_SHORT_PUSH_DATA_CMDS byte

const (
	PM_SET_SYNC_DISTANCE       PM_SHORT_PUSH_DATA_CMDS = 0xD0 + iota // SETPMDATA_CMD_SHORT_MIN
	PM_SET_SYNC_STROKEPACE                                           // 0xD1
	PM_SET_SYNC_AVG_HEARTRATE                                        // 0xD2
	PM_SET_SYNC_TIME                                                 // 0xD3
	PM_SET_SYNC_SPLIT_DATA                                           // 0xD4
	PM_SET_SYNC_ENCODER_PERIOD                                       // 0xD5
	PM_SET_SYNC_VERSION_INFO                                         // 0xD6
	PM_SET_SYNC_RACETICKTIME                                         // 0xD7
	PM_SET_SYNC_DATAALL                                              // 0xD8
	_                                                                // Unused 0xD9
	_                                                                // Unused 0xDA
	_                                                                // Unused 0xDB
	_                                                                // Unused 0xDC
	_                                                                // Unused 0xDD
	_                                                                // Unused 0xDE
	_                                                                // Unused 0xDF
	SETPMDATA_CMD_SHORT_MAX
)

/*
   Custom Short PUSH Configuration Commands for PM
*/
type PM_SHORT_PUSH_CFG_CMDS byte

const (
	PM_SET_RESET_ALL       PM_SHORT_PUSH_CFG_CMDS = 0xE0 + iota // SETPMCFG_CMD_SHORT_MIN
	PM_SET_RESET_ERGNUMBER                                      // 0xE1
	_                                                           // Unused 0xE2
	_                                                           // Unused 0xE3
	_                                                           // Unused 0xE4
	_                                                           // Unused 0xE5
	_                                                           // Unused 0xE6
	_                                                           // Unused 0xE7
	_                                                           // Unused 0xE8
	_                                                           // Unused 0xE9
	_                                                           // Unused 0xEA
	_                                                           // Unused 0xEB
	_                                                           // Unused 0xEC
	_                                                           // Unused 0xED
	_                                                           // Unused 0xEE
	_                                                           // Unused 0xEF
	SETPMCFG_CMD_SHORT_MAX
)

/*
   Custom Long PUSH Configuration Commands for PM
*/
type PM_LONG_PUSH_CFG_CMDS byte

const (
	PM_SET_BAUDRATE                 PM_LONG_PUSH_CFG_CMDS = 0x00 + iota // SETPMCFG_CMD_LONG_MIN
	PM_SET_WORKOUTTYPE                                                  // 0x01
	PM_SET_STARTTYPE                                                    // 0x02
	PM_SET_WORKOUTDURATION                                              // 0x03
	PM_SET_RESTDURATION                                                 // 0x04
	PM_SET_SPLITDURATION                                                // 0x05
	PM_SET_TARGETPACETIME                                               // 0x06
	PM_SET_INTERVALIDENTIFIER                                           // 0x07
	PM_SET_OPERATIONALSTATE                                             // 0x08
	PM_SET_RACETYPE                                                     // 0x09
	PM_SET_WARMUPDURATION                                               // 0x0A
	PM_SET_RACELANESETUP                                                // 0x0B
	PM_SET_RACELANEVERIFY                                               // 0x0C
	PM_SET_RACESTARTPARAMS                                              // 0x0D
	PM_SET_ERGSLAVEDISCOVERYREQUEST                                     // 0x0E
	PM_SET_BOATNUMBER                                                   // 0x0F
	PM_SET_ERGNUMBER                                                    // 0x10
	PM_SET_COMMUNICATIONSTATE                                           // 0x11
	PM_SET_CMDUPLIST                                                    // 0x12
	PM_SET_SCREENSTATE                                                  // 0x13
	PM_CONFIGURE_WORKOUT                                                // 0x14
	PM_SET_TARGETAVGWATTS                                               // 0x15
	PM_SET_TARGETCALSPERHR                                              // 0x16
	PM_SET_INTERVALTYPE                                                 // 0x17
	PM_SET_WORKOUTINTERVALCOUNT                                         // 0x18
	PM_SET_DISPLAYUPDATERATE                                            // 0x19
	PM_SET_AUTHENPASSWORD                                               // 0x1A
	PM_SET_TICKTIME                                                     // 0x1B
	PM_SET_TICKTIMEOFFSET                                               // 0x1C
	PM_SET_RACEDATASAMPLETICKS                                          // 0x1D
	PM_SET_RACEOPERATIONTYPE                                            // 0x1E
	PM_SET_RACESTATUSDISPLAYTICKS                                       // 0x1F
	PM_SET_RACESTATUSWARNINGTICKS                                       // 0x20
	PM_SET_RACEIDLEMODEPARAMS                                           // 0x21
	PM_SET_DATETIME                                                     // 0x22
	PM_SET_LANGUAGETYPE                                                 // 0x23
	PM_SET_WIFICONFIG                                                   // 0x24
	PM_SET_CPUTICKRATE                                                  // 0x25
	PM_SET_LOGCARDUSER                                                  // 0x26
	PM_SET_SCREENERRORMODE                                              // 0x27
	PM_SET_CABLETEST                                                    // 0x28
	PM_SET_USER_ID                                                      // 0x29
	PM_SET_USER_PROFILE                                                 // 0x2A
	PM_SET_HRM                                                          // 0x2B
	_                                                                   // Unused 0x2C
	_                                                                   // Unused 0x2D
	_                                                                   // Unused 0x2E
	PM_SET_SENSOR_CHANNEL                                               // 0x2F sensor channel
	SETPMCFG_CMD_LONG_MAX
)

/*
   Custom Long PUSH Data Commands for PM
*/
type PM_LONG_PUSH_DATA_CMDS byte

const (
	PM_SET_TEAM_DISTANCE        PM_LONG_PUSH_DATA_CMDS = 0x30 + iota // SETPMDATA_CMD_LONG_MIN
	PM_SET_TEAM_FINISH_TIME                                          // 0x31
	PM_SET_RACEPARTICIPANT                                           // 0x32
	PM_SET_RACESTATUS                                                // 0x33
	PM_SET_LOGCARDMEMORY                                             // 0x34
	PM_SET_DISPLAYSTRING                                             // 0x35
	PM_SET_DISPLAYBITMAP                                             // 0x36
	PM_SET_LOCALRACEPARTICIPANT                                      // 0x37
	_                                                                // Unused 0x38
	_                                                                // Unused 0x39
	_                                                                // Unused 0x3A
	_                                                                // Unused 0x3B
	_                                                                // Unused 0x3C
	_                                                                // Unused 0x3D
	_                                                                // Unused 0x3E
	_                                                                // Unused 0x3F
	_                                                                // Unused 0x40
	_                                                                // Unused 0x41
	_                                                                // Unused 0x42
	_                                                                // Unused 0x43
	_                                                                // Unused 0x44
	_                                                                // Unused 0x45
	_                                                                // Unused 0x46
	_                                                                // Unused 0x47
	_                                                                // Unused 0x48
	_                                                                // Unused 0x49
	_                                                                // Unused 0x4A
	_                                                                // Unused 0x4B
	_                                                                // Unused 0x4C
	_                                                                // Unused 0x4D
	PM_SET_ANTRFMODE                                                 // 0x4E mfg support only
	PM_SET_MEMORY                                                    // 0x4F debug only
	SETPMDATA_CMD_LONG_MAX
)

/*
   Custom Long PULL Configuration Commands for PM
*/
type PM_LONG_PULL_CFG_CMDS byte

const (
	PM_GET_ERGNUMBER            PM_LONG_PULL_CFG_CMDS = 0x50 + iota // GETPMCFG_CMD_LONG_MIN
	PM_GET_ERGNUMBERREQUEST                                         // 0x51
	PM_GET_USERIDSTRING                                             // 0x52
	PM_GET_LOCALRACEPARTICIPANT                                     // 0x53
	PM_GET_USER_ID                                                  // 0x54
	PM_GET_USER_PROFILE                                             // 0x55
	_                                                               // PM_GET_WORKOUTPARAMETERS  // 0x56
	_                                                               // Unused 0x57
	_                                                               // Unused 0x58
	_                                                               // Unused 0x59
	_                                                               // Unused 0x5A
	_                                                               // Unused 0x5B
	_                                                               // Unused 0x5C
	_                                                               // Unused 0x5D
	_                                                               // Unused 0x5E
	_                                                               // Unused 0x5F
	_                                                               // Unused 0x60
	_                                                               // Unused 0x61
	_                                                               // Unused 0x62
	_                                                               // Unused 0x63
	_                                                               // Unused 0x64
	_                                                               // Unused 0x65
	_                                                               // Unused 0x66
	_                                                               // Unused 0x67
	GETPMCFG_CMD_LONG_MAX
)

/*
   Custom Long PULL Data Commands for PM
*/
type PM_LONG_PULL_DATA_CMDS byte

const (
	PM_GET_MEMORY            PM_LONG_PULL_DATA_CMDS = 0x68 + iota // GETPMDATA_CMD_LONG_MIN
	PM_GET_LOGCARDMEMORY                                          // 0x69
	PM_GET_INTERNALLOGMEMORY                                      // 0x6A
	PM_GET_FORCEPLOTDATA                                          // 0x6B
	PM_GET_HEARTBEATDATA                                          // 0x6C
	PM_GET_UI_EVENTS                                              // 0x6D
	_                                                             // Unused 0x6E
	_                                                             // Unused 0x6F
	_                                                             // Unused 0x70
	_                                                             // Unused 0x71
	_                                                             // Unused 0x72
	_                                                             // Unused 0x73
	_                                                             // Unused 0x74
	_                                                             // Unused 0x75
	_                                                             // Command Wrapper 0x76
	_                                                             // Command Wrapper 0x77
	_                                                             // Unused 0x78
	_                                                             // Unused 0x79
	_                                                             // Unused 0x7A
	_                                                             // Unused 0x7B
	_                                                             // Unused 0x7C
	_                                                             // Unused 0x7D
	_                                                             // Command Wrapper 0x7E
	_                                                             // Command Wrapper 0x7F
	GETPMDATA_CMD_LONG_MAX
)

/* Status byte flag and mask definitions */
const PREVOK_FLG = 0x00
const PREVREJECT_FLG = 0x10
const PREVBAD_FLG = 0x20
const PREVNOTRDY_FLG = 0x30
const PREVFRAMESTATUS_MSK = 0x30

const SLAVESTATE_ERR_FLG = 0x00
const SLAVESTATE_RDY_FLG = 0x01
const SLAVESTATE_IDLE_FLG = 0x02
const SLAVESTATE_HAVEID_FLG = 0x03
const SLAVESTATE_INUSE_FLG = 0x05
const SLAVESTATE_PAUSE_FLG = 0x06
const SLAVESTATE_FINISH_FLG = 0x07
const SLAVESTATE_MANUAL_FLG = 0x08
const SLAVESTATE_OFFLINE_FLG = 0x09

const FRAMECNT_FLG = 0x80

const SLAVESTATE_MSK = 0x0F

/* AUTOUPLOAD_CMD flag definitions */
const AUTOSTATUS_FLG = 0x01
const UPSTATUS_FLG = 0x02
const UPLIST_FLG = 0x04
const ACK_FLG = 0x10
const EXTERNCONTROL_FLG = 0x40

/* CSAFE Slave Capabilities Codes */
const CAPCODE_PROTOCOL = 0x00
const CAPCODE_POWER = 0x01
const CAPCODE_TEXT = 0x02

/* CSAFE units format definitions: <type>_<unit>_<tens>_<decimals> */
const DISTANCE_MILE_0_0 = 0x01
const DISTANCE_MILE_0_1 = 0x02
const DISTANCE_MILE_0_2 = 0x03
const DISTANCE_MILE_0_3 = 0x04
const DISTANCE_FEET_0_0 = 0x05
const DISTANCE_INCH_0_0 = 0x06
const WEIGHT_LBS_0_0 = 0x07
const WEIGHT_LBS_0_1 = 0x08
const DISTANCE_FEET_1_0 = 0x0A
const SPEED_MILEPERHOUR_0_0 = 0x10
const SPEED_MILEPERHOUR_0_1 = 0x11
const SPEED_MILEPERHOUR_0_2 = 0x12
const SPEED_FEETPERMINUTE_0_0 = 0x13
const DISTANCE_KM_0_0 = 0x21
const DISTANCE_KM_0_1 = 0x22
const DISTANCE_KM_0_2 = 0x23
const DISTANCE_METER_0_0 = 0x24
const DISTANCE_METER_0_1 = 0x25
const DISTANCE_CM_0_0 = 0x26
const WEIGHT_KG_0_0 = 0x27
const WEIGHT_KG_0_1 = 0x28
const SPEED_KMPERHOUR_0_0 = 0x30
const SPEED_KMPERHOUR_0_1 = 0x31
const SPEED_KMPERHOUR_0_2 = 0x32
const SPEED_METERPERMINUTE_0_0 = 0x33
const PACE_MINUTEPERMILE_0_0 = 0x37
const PACE_MINUTEPERKM_0_0 = 0x38
const PACE_SECONDSPERKM_0_0 = 0x39
const PACE_SECONDSPERMILE_0_0 = 0x3A
const DISTANCE_FLOORS_0_0 = 0x41
const DISTANCE_FLOORS_0_1 = 0x42
const DISTANCE_STEPS_0_0 = 0x43
const DISTANCE_REVS_0_0 = 0x44
const DISTANCE_STRIDES_0_0 = 0x45
const DISTANCE_STROKES_0_0 = 0x46
const MISC_BEATS_0_0 = 0x47
const ENERGY_CALORIES_0_0 = 0x48
const GRADE_PERCENT_0_0 = 0x4A
const GRADE_PERCENT_0_2 = 0x4B
const GRADE_PERCENT_0_1 = 0x4C
const CADENCE_FLOORSPERMINUTE_0_1 = 0x4F
const CADENCE_FLOORSPERMINUTE_0_0 = 0x50
const CADENCE_STEPSPERMINUTE_0_0 = 0x51
const CADENCE_REVSPERMINUTE_0_0 = 0x52
const CADENCE_STRIDESPERMINUTE_0_0 = 0x53
const CADENCE_STROKESPERMINUTE_0_0 = 0x54
const MISC_BEATSPERMINUTE_0_0 = 0x55
const BURN_CALORIESPERMINUTE_0_0 = 0x56
const BURN_CALORIESPERHOUR_0_0 = 0x57
const POWER_WATTS_0_0 = 0x58
const ENERGY_INCHLB_0_0 = 0x5A
const ENERGY_FOOTLB_0_0 = 0x5B
const ENERGY_NM_0_0 = 0x5C

/* Conversion constants */
const KG_TO_LBS = 2.2046
const LBS_TO_KG = (1. / KG_TO_LBS)

/* ID Digits */
const IDDIGITS_MIN = 2
const IDDIGITS_MAX = 5
const DEFAULT_IDDIGITS = 5
const DEFAULT_ID = 0
const MANUAL_ID = 999999999

/* Slave State Tiimeout Parameters */
const DEFAULT_SLAVESTATE_TIMEOUT = 20 // seconds
const PAUSED_SLAVESTATE_TIMEOUT = 220 // seconds
const INUSE_SLAVESTATE_TIMEOUT = 6    // seconds
const IDLE_SLAVESTATE_TIMEOUT = 30    // seconds

/* Base Year */
const BASE_YEAR = 1900

/* Default time intervals */
const DEFAULT_STATUSUPDATE_INTERVAL = 256 // seconds
const DEFAULT_CMDUPLIST_INTERVAL = 256    // seconds
