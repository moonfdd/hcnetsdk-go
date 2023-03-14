package hcnetsdk

import (
	"unsafe"

	"github.com/moonfdd/ffmpeg-go/ffcommon"
	"github.com/moonfdd/hcnetsdk-go/hcnetsdkcommon"
)

// #ifndef _PLAYM4_H_
// const _PLAYM4_H_

// #if defined( _WINDLL)
//     const PLAYM4_API  extern "C" __declspec(dllexport)
// #else
//     const PLAYM4_API  extern "C" __declspec(dllimport)
// #endif

// Max channel numbers
const PLAYM4_MAX_SUPPORTS = 500

// Wave coef range;
const MIN_WAVE_COEF = -100
const MAX_WAVE_COEF = 100

// Timer type
const TIMER_1 = 1 //Only 16 timers for every process.Default TIMER;
const TIMER_2 = 2 //Not limit;But the precision less than TIMER_1;

// BUFFER AND DATA TYPE
const BUF_VIDEO_SRC = (1)     //mixed input,total src buffer size;splited input,video src buffer size
const BUF_AUDIO_SRC = (2)     //mixed input,not defined;splited input,audio src buffer size
const BUF_VIDEO_RENDER = (3)  //video render node count
const BUF_AUDIO_RENDER = (4)  //audio render node count
const BUF_VIDEO_DECODED = (5) //video decoded node count to render
const BUF_AUDIO_DECODED = (6) //audio decoded node count to render
const BUF_DISPLAY_NODE = (7)  //display node

// Error code
const PLAYM4_NOERROR = 0                 //no error
const PLAYM4_PARA_OVER = 1               //input parameter is invalid;
const PLAYM4_ORDER_ERROR = 2             //The order of the function to be called is error.
const PLAYM4_TIMER_ERROR = 3             //Create multimedia clock failed;
const PLAYM4_DEC_VIDEO_ERROR = 4         //Decode video data failed.
const PLAYM4_DEC_AUDIO_ERROR = 5         //Decode audio data failed.
const PLAYM4_ALLOC_MEMORY_ERROR = 6      //Allocate memory failed.
const PLAYM4_OPEN_FILE_ERROR = 7         //Open the file failed.
const PLAYM4_CREATE_OBJ_ERROR = 8        //Create thread or event failed
const PLAYM4_CREATE_DDRAW_ERROR = 9      //Create DirectDraw object failed.
const PLAYM4_CREATE_OFFSCREEN_ERROR = 10 //failed when creating off-screen surface.
const PLAYM4_BUF_OVER = 11               //buffer is overflow
const PLAYM4_CREATE_SOUND_ERROR = 12     //failed when creating audio device.
const PLAYM4_SET_VOLUME_ERROR = 13       //Set volume failed
const PLAYM4_SUPPORT_FILE_ONLY = 14      //The function only support play file.
const PLAYM4_SUPPORT_STREAM_ONLY = 15    //The function only support play stream.
const PLAYM4_SYS_NOT_SUPPORT = 16        //System not support.
const PLAYM4_FILEHEADER_UNKNOWN = 17     //No file header.
const PLAYM4_VERSION_INCORRECT = 18      //The version of decoder and encoder is not adapted.
const PLAYM4_INIT_DECODER_ERROR = 19     //Initialize decoder failed.
const PLAYM4_CHECK_FILE_ERROR = 20       //The file data is unknown.
const PLAYM4_INIT_TIMER_ERROR = 21       //Initialize multimedia clock failed.
const PLAYM4_BLT_ERROR = 22              //Blt failed.
const PLAYM4_UPDATE_ERROR = 23           //Update failed.
const PLAYM4_OPEN_FILE_ERROR_MULTI = 24  //openfile error, streamtype is multi
const PLAYM4_OPEN_FILE_ERROR_VIDEO = 25  //openfile error, streamtype is video
const PLAYM4_JPEG_COMPRESS_ERROR = 26    //JPEG compress error
const PLAYM4_EXTRACT_NOT_SUPPORT = 27    //Don't support the version of this file.
const PLAYM4_EXTRACT_DATA_ERROR = 28     //extract video data failed.
const PLAYM4_SECRET_KEY_ERROR = 29       //Secret key is error //add 20071218
const PLAYM4_DECODE_KEYFRAME_ERROR = 30  //add by hy 20090318
const PLAYM4_NEED_MORE_DATA = 31         //add by hy 20100617
const PLAYM4_INVALID_PORT = 32           //add by cj 20100913
const PLAYM4_NOT_FIND = 33               //add by cj 20110428
const PLAYM4_NEED_LARGER_BUFFER = 34     //add by pzj 20130528
const PLAYM4_FAIL_UNKNOWN = 99           //Fail, but the reason is unknown;

// 鱼眼功能错误码
const PLAYM4_FEC_ERR_ENABLEFAIL = 100       // 鱼眼模块加载失败
const PLAYM4_FEC_ERR_NOTENABLE = 101        // 鱼眼模块没有加载
const PLAYM4_FEC_ERR_NOSUBPORT = 102        // 子端口没有分配
const PLAYM4_FEC_ERR_PARAMNOTINIT = 103     // 没有初始化对应端口的参数
const PLAYM4_FEC_ERR_SUBPORTOVER = 104      // 子端口已经用完
const PLAYM4_FEC_ERR_EFFECTNOTSUPPORT = 105 // 该安装方式下这种效果不支持
const PLAYM4_FEC_ERR_INVALIDWND = 106       // 非法的窗口
const PLAYM4_FEC_ERR_PTZOVERFLOW = 107      // PTZ位置越界
const PLAYM4_FEC_ERR_RADIUSINVALID = 108    // 圆心参数非法
const PLAYM4_FEC_ERR_UPDATENOTSUPPORT = 109 // 指定的安装方式和矫正效果，该参数更新不支持
const PLAYM4_FEC_ERR_NOPLAYPORT = 110       // 播放库端口没有启用
const PLAYM4_FEC_ERR_PARAMVALID = 111       // 参数为空
const PLAYM4_FEC_ERR_INVALIDPORT = 112      // 非法子端口
const PLAYM4_FEC_ERR_PTZZOOMOVER = 113      // PTZ矫正范围越界
const PLAYM4_FEC_ERR_OVERMAXPORT = 114      // 矫正通道饱和，最大支持的矫正通道为四个
const PLAYM4_FEC_ERR_ENABLED = 115          //该端口已经启用了鱼眼模块
const PLAYM4_FEC_ERR_D3DACCENOTENABLE = 116 // D3D加速没有开启

// Max display regions.
const MAX_DISPLAY_WND = 4

// Display type
const DISPLAY_NORMAL = 0x00000001
const DISPLAY_QUARTER = 0x00000002
const DISPLAY_YC_SCALE = 0x00000004 //add by gb 20091116
const DISPLAY_NOTEARING = 0x00000008

// Display buffers
const MAX_DIS_FRAMES = 50
const MIN_DIS_FRAMES = 1

// Locate by
const BY_FRAMENUM = 1
const BY_FRAMETIME = 2

// Source buffer
const SOURCE_BUF_MAX = 1024 * 100000
const SOURCE_BUF_MIN = 1024 * 50

// Stream type
const STREAME_REALTIME = 0
const STREAME_FILE = 1

// frame type
const T_AUDIO16 = 101
const T_AUDIO8 = 100
const T_UYVY = 1
const T_YV12 = 3
const T_RGB32 = 7

// capability
const SUPPORT_DDRAW = 1
const SUPPORT_BLT = 2
const SUPPORT_BLTFOURCC = 4
const SUPPORT_BLTSHRINKX = 8
const SUPPORT_BLTSHRINKY = 16
const SUPPORT_BLTSTRETCHX = 32
const SUPPORT_BLTSTRETCHY = 64
const SUPPORT_SSE = 128
const SUPPORT_MMX = 256

// 以下宏定义用于HIK_MEDIAINFO结构
const FOURCC_HKMI = 0x484B4D49 // "HKMI" HIK_MEDIAINFO结构标记
// 系统封装格式
const SYSTEM_NULL = 0x0     // 没有系统层，纯音频流或视频流
const SYSTEM_HIK = 0x1      // 海康文件层
const SYSTEM_MPEG2_PS = 0x2 // PS封装
const SYSTEM_MPEG2_TS = 0x3 // TS封装
const SYSTEM_RTP = 0x4      // rtp封装
const SYSTEM_RTPHIK = 0x401 // rtp封装

// 视频编码类型
const VIDEO_NULL = 0x0  // 没有视频
const VIDEO_H264 = 0x1  // 海康H.264
const VIDEO_MPEG4 = 0x3 // 标准MPEG4
const VIDEO_MJPEG = 0x4
const VIDEO_AVC264 = 0x0100

// 音频编码类型
const AUDIO_NULL = 0x0000  // 没有音频
const AUDIO_ADPCM = 0x1000 // ADPCM
const AUDIO_MPEG = 0x2000  // MPEG 系列音频，解码器能自适应各种MPEG音频
const AUDIO_AAC = 0x2001   // AAC 编码
// G系列音频
const AUDIO_RAW_DATA8 = 0x7000   //采样率为8k的原始数据
const AUDIO_RAW_UDATA16 = 0x7001 //采样率为16k的原始数据，即L16
const AUDIO_G711_U = 0x7110
const AUDIO_G711_A = 0x7111
const AUDIO_G722_1 = 0x7221
const AUDIO_G723_1 = 0x7231
const AUDIO_G726_U = 0x7260
const AUDIO_G726_A = 0x7261
const AUDIO_G726_16 = 0x7262
const AUDIO_G729 = 0x7290
const AUDIO_AMR_NB = 0x3000

const SYNCDATA_VEH = 1 //同步数据:车载信息
const SYNCDATA_IVS = 2 //同步数据:智能信息

// motion flow type
const MOTION_FLOW_NONE = 0
const MOTION_FLOW_CPU = 1
const MOTION_FLOW_GPU = 2

// 音视频加密类型
const ENCRYPT_AES_3R_VIDEO = 1
const ENCRYPT_AES_10R_VIDEO = 2
const ENCRYPT_AES_3R_AUDIO = 1
const ENCRYPT_AES_10R_AUDIO = 2

type SYSTEMTIME struct {
	WYear         hcnetsdkcommon.FWord // 年
	WMonth        hcnetsdkcommon.FWord // 月
	WDayOfWeek    hcnetsdkcommon.FWord // 星期几，从星期日开始，0代表星期日、1代表星期一，以此类推
	WDay          hcnetsdkcommon.FWord // 日
	WHour         hcnetsdkcommon.FWord // 时
	WMinute       hcnetsdkcommon.FWord // 分
	WSecond       hcnetsdkcommon.FWord // 秒
	WMilliseconds hcnetsdkcommon.FWord // 毫秒
}

// Frame position
type FRAME_POS struct {
	NFilePos           ffcommon.FLong
	NFrameNum          ffcommon.FLong
	NFrameTime         ffcommon.FLong
	NErrorFrameNum     ffcommon.FLong
	PErrorTime         *SYSTEMTIME
	NErrorLostFrameNum ffcommon.FLong
	NErrorFrameSize    ffcommon.FLong
}

// Frame Info
type FRAME_INFO struct {
	NWidth     ffcommon.FLong
	NHeight    ffcommon.FLong
	NStamp     ffcommon.FLong
	NType      ffcommon.FLong
	NFrameRate ffcommon.FLong
	DwFrameNum hcnetsdkcommon.FDword
}

type DISPLAY_INFO_YUV struct {
	NPort    ffcommon.FLong           //通道号
	PBuf     ffcommon.FCharPStruct    //返回的第一路图像数据指针
	NBufLen  ffcommon.FUnsignedInt    //返回的第一路图像数据大小
	PBuf1    ffcommon.FCharPStruct    //返回的第二路图像数据指针
	NBufLen1 ffcommon.FUnsignedInt    //返回的第二路图像数据大小
	PBuf2    ffcommon.FCharPStruct    //返回的第三路图像数据指针
	NBufLen2 ffcommon.FUnsignedInt    //返回的第三路图像数据大小
	NWidth   ffcommon.FUnsignedInt    //画面宽
	NHeight  ffcommon.FUnsignedInt    //画面高
	NStamp   ffcommon.FUnsignedInt    //时标信息，单位毫秒
	NType    ffcommon.FUnsignedInt    //数据类型
	PUser    ffcommon.FVoidP          //用户数据
	Reserved [4]ffcommon.FUnsignedInt //保留
}

// Frame
type FRAME_TYPE struct {
	PDataBuf  ffcommon.FCharPStruct
	NSize     ffcommon.FLong
	NFrameNum ffcommon.FLong
	BIsAudio  bool
	NReserved ffcommon.FLong
}

// Watermark Info    //add by gb 080119
type WATERMARK_INFO struct {
	PDataBuf  ffcommon.FCharPStruct
	NSize     ffcommon.FLong
	NFrameNum ffcommon.FLong
	BRsaRight bool
	NReserved ffcommon.FLong
}

type SYNCDATA_INFO struct {
	DwDataType hcnetsdkcommon.FDword //和码流数据同步的附属信息类型，目前有：智能信息，车载信息
	DwDataLen  hcnetsdkcommon.FDword //附属信息数据长度
	PData      *byte                 //指向附属信息数据结构的指针,比如IVS_INFO结构
}

// #ifndef _HIK_MEDIAINFO_FLAG_
// const _HIK_MEDIAINFO_FLAG_
type HIK_MEDIAINFO struct // modified by gb 080425
{
	MediaFourcc  ffcommon.FUnsignedInt         // "HKMI": 0x484B4D49 Hikvision Media Information
	MediaVersion hcnetsdkcommon.FUnsignedShort // 版本号：指本信息结构版本号，目前为0x0101,即1.01版本，01：主版本号；01：子版本号。
	DeviceId     hcnetsdkcommon.FUnsignedShort // 设备ID，便于跟踪/分析

	SystemFormat hcnetsdkcommon.FUnsignedShort // 系统封装层
	VideoFormat  hcnetsdkcommon.FUnsignedShort // 视频编码类型

	AudioFormat        hcnetsdkcommon.FUnsignedShort // 音频编码类型
	AudioChannels      hcnetsdkcommon.FUnsignedChar  // 通道数
	AudioBitsPerSample hcnetsdkcommon.FUnsignedChar  // 样位率
	AudioSamplesrate   ffcommon.FUnsignedInt         // 采样率
	AudioBitrate       ffcommon.FUnsignedInt         // 压缩音频码率,单位：bit

	Reserved [4]ffcommon.FUnsignedInt // 保留
}

// #endif

type DISPLAY_INFO struct {
	NPort   ffcommon.FLong
	PBuf    ffcommon.FCharP
	NBufLen ffcommon.FLong
	NWidth  ffcommon.FLong
	NHeight ffcommon.FLong
	NStamp  ffcommon.FLong
	NType   ffcommon.FLong
	NUser   ffcommon.FLong
}

type DISPLAY_INFOEX struct {
	NPort        ffcommon.FLong
	PVideoBuf    ffcommon.FCharP
	NVideoBufLen ffcommon.FLong
	PPriBuf      ffcommon.FCharP
	NPriBufLen   ffcommon.FLong
	NWidth       ffcommon.FLong
	NHeight      ffcommon.FLong
	NStamp       ffcommon.FLong
	NType        ffcommon.FLong
	NUser        ffcommon.FLong
}

type PLAYM4_SYSTEM_TIME struct { //绝对时间

	DwYear hcnetsdkcommon.FDword //年
	DwMon  hcnetsdkcommon.FDword //月
	DwDay  hcnetsdkcommon.FDword //日
	DwHour hcnetsdkcommon.FDword //时
	DwMin  hcnetsdkcommon.FDword //分
	DwSec  hcnetsdkcommon.FDword //秒
	DwMs   hcnetsdkcommon.FDword //毫秒
}

// ENCRYPT Info
type ENCRYPT_INFO struct {
	NVideoEncryptType ffcommon.FLong //视频加密类型
	NAudioEncryptType ffcommon.FLong //音频加密类型
	NSetSecretKey     ffcommon.FLong //是否设置，1表示设置密钥，0表示没有设置密钥
}

//////////////////////////////////////////////////////////////////////////////
//API
//////////////////////////////////////////////////////////////////////////////

// //////////////ver 1.0///////////////////////////////////////
// Initialize DirecDraw.Now invalid.
// PLAYM4_API BOOL __stdcall  PlayM4_InitDDraw(HWND hWnd);
func PlayM4_InitDDraw(hWnd hcnetsdkcommon.FHwnd) (res bool) {
	t, _, _ := hcnetsdkcommon.GetHcnetsdkDll().NewProc("PlayM4_InitDDraw").Call(
		hWnd,
	)
	res = ffcommon.GoBool(t)
	return
}

// Release directDraw; Now invalid.
// PLAYM4_API BOOL __stdcall PlayM4_RealeseDDraw();
func PlayM4_RealeseDDraw() (res bool) {
	t, _, _ := hcnetsdkcommon.GetHcnetsdkDll().NewProc("PlayM4_RealeseDDraw").Call()
	res = ffcommon.GoBool(t)
	return
}

// PLAYM4_API BOOL __stdcall PlayM4_OpenFile(LONG nPort,LPSTR sFileName);
func PlayM4_OpenFile(nPort ffcommon.FLong, sFileName hcnetsdkcommon.FLpstr) (res bool) {
	t, _, _ := hcnetsdkcommon.GetHcnetsdkDll().NewProc("PlayM4_OpenFile").Call(
		uintptr(nPort),
		ffcommon.UintPtrFromString(sFileName),
	)
	res = ffcommon.GoBool(t)
	return
}

// PLAYM4_API BOOL __stdcall PlayM4_CloseFile(LONG nPort);
func PlayM4_CloseFile(nPort ffcommon.FLong) (res bool) {
	t, _, _ := hcnetsdkcommon.GetHcnetsdkDll().NewProc("PlayM4_CloseFile").Call(
		uintptr(nPort),
	)
	res = ffcommon.GoBool(t)
	return
}

// PLAYM4_API BOOL __stdcall PlayM4_Play(LONG nPort, HWND hWnd);
func PlayM4_Play(nPort ffcommon.FLong, hWnd hcnetsdkcommon.FHwnd) (res bool) {
	t, _, _ := hcnetsdkcommon.GetHcnetsdkDll().NewProc("PlayM4_Play").Call(
		uintptr(nPort),
		hWnd,
	)
	res = ffcommon.GoBool(t)
	return
}

// PLAYM4_API BOOL __stdcall PlayM4_Stop(LONG nPort);
func PlayM4_Stop(nPort ffcommon.FLong) (res bool) {
	t, _, _ := hcnetsdkcommon.GetHcnetsdkDll().NewProc("PlayM4_Stop").Call(
		uintptr(nPort),
	)
	res = ffcommon.GoBool(t)
	return
}

// PLAYM4_API BOOL __stdcall PlayM4_Pause(LONG nPort,DWORD nPause);
func PlayM4_Pause(nPort ffcommon.FLong, nPause hcnetsdkcommon.FDword) (res bool) {
	t, _, _ := hcnetsdkcommon.GetHcnetsdkDll().NewProc("PlayM4_Pause").Call(
		uintptr(nPort),
		uintptr(nPause),
	)
	res = ffcommon.GoBool(t)
	return
}

// PLAYM4_API BOOL __stdcall PlayM4_Fast(LONG nPort);
func PlayM4_Fast(nPort ffcommon.FLong) (res bool) {
	t, _, _ := hcnetsdkcommon.GetHcnetsdkDll().NewProc("PlayM4_Fast").Call(
		uintptr(nPort),
	)
	res = ffcommon.GoBool(t)
	return
}

// PLAYM4_API BOOL __stdcall PlayM4_Slow(LONG nPort);
func PlayM4_Slow(nPort ffcommon.FLong) (res bool) {
	t, _, _ := hcnetsdkcommon.GetHcnetsdkDll().NewProc("PlayM4_Slow").Call(
		uintptr(nPort),
	)
	res = ffcommon.GoBool(t)
	return
}

// PLAYM4_API BOOL __stdcall PlayM4_OneByOne(LONG nPort);
func PlayM4_OneByOne(nPort ffcommon.FLong) (res bool) {
	t, _, _ := hcnetsdkcommon.GetHcnetsdkDll().NewProc("PlayM4_OneByOne").Call(
		uintptr(nPort),
	)
	res = ffcommon.GoBool(t)
	return
}

// PLAYM4_API BOOL __stdcall PlayM4_SetPlayPos(LONG nPort,float fRelativePos);
func PlayM4_SetPlayPos(nPort ffcommon.FLong, fRelativePos ffcommon.FFloat) (res bool) {
	t, _, _ := hcnetsdkcommon.GetHcnetsdkDll().NewProc("PlayM4_SetPlayPos").Call(
		uintptr(nPort),
		*(*uintptr)(unsafe.Pointer(&fRelativePos)),
	)
	res = ffcommon.GoBool(t)
	return
}

// PLAYM4_API float __stdcall PlayM4_GetPlayPos(LONG nPort);
func PlayM4_GetPlayPos(nPort ffcommon.FLong) (res bool) {
	t, _, _ := hcnetsdkcommon.GetHcnetsdkDll().NewProc("PlayM4_GetPlayPos").Call(
		uintptr(nPort),
	)
	res = ffcommon.GoBool(t)
	return
}

// PLAYM4_API BOOL __stdcall PlayM4_SetFileEndMsg(LONG nPort,HWND hWnd,UINT nMsg);
func PlayM4_SetFileEndMsg(nPort ffcommon.FLong, hWnd hcnetsdkcommon.FHwnd, nMsg ffcommon.FUint) (res bool) {
	t, _, _ := hcnetsdkcommon.GetHcnetsdkDll().NewProc("PlayM4_SetFileEndMsg").Call(
		uintptr(nPort),
		hWnd,
		uintptr(nMsg),
	)
	res = ffcommon.GoBool(t)
	return
}

// PLAYM4_API BOOL __stdcall PlayM4_SetVolume(LONG nPort,WORD nVolume);
func PlayM4_SetVolume(nPort ffcommon.FLong, nVolume hcnetsdkcommon.FWord) (res bool) {
	t, _, _ := hcnetsdkcommon.GetHcnetsdkDll().NewProc("PlayM4_SetVolume").Call(
		uintptr(nPort),
		uintptr(nVolume),
	)
	res = ffcommon.GoBool(t)
	return
}

// PLAYM4_API BOOL __stdcall PlayM4_StopSound();
func PlayM4_StopSound() (res bool) {
	t, _, _ := hcnetsdkcommon.GetHcnetsdkDll().NewProc("PlayM4_StopSound").Call()
	res = ffcommon.GoBool(t)
	return
}

// PLAYM4_API BOOL __stdcall PlayM4_PlaySound(LONG nPort);
func PlayM4_PlaySound(nPort ffcommon.FLong) (res bool) {
	t, _, _ := hcnetsdkcommon.GetHcnetsdkDll().NewProc("PlayM4_PlaySound").Call(
		uintptr(nPort),
	)
	res = ffcommon.GoBool(t)
	return
}

// PLAYM4_API BOOL __stdcall PlayM4_OpenStream(LONG nPort,PBYTE pFileHeadBuf,DWORD nSize,DWORD nBufPoolSize);
func PlayM4_OpenStream(nPort ffcommon.FLong, pFileHeadBuf ffcommon.FBuf, nSize, nBufPoolSize hcnetsdkcommon.FDword) (res bool) {
	t, _, _ := hcnetsdkcommon.GetHcnetsdkDll().NewProc("PlayM4_OpenStream").Call(
		uintptr(nPort),
		uintptr(unsafe.Pointer(pFileHeadBuf)),
		uintptr(nSize),
		uintptr(nBufPoolSize),
	)
	res = ffcommon.GoBool(t)
	return
}

// PLAYM4_API BOOL __stdcall PlayM4_InputData(LONG nPort,PBYTE pBuf,DWORD nSize);
func PlayM4_InputData(nPort ffcommon.FLong, pBuf ffcommon.FBuf, nSize hcnetsdkcommon.FDword) (res bool) {
	t, _, _ := hcnetsdkcommon.GetHcnetsdkDll().NewProc("PlayM4_InputData").Call(
		uintptr(nPort),
		uintptr(unsafe.Pointer(pBuf)),
		uintptr(nSize),
	)
	res = ffcommon.GoBool(t)
	return
}

// PLAYM4_API BOOL __stdcall PlayM4_CloseStream(LONG nPort);
func PlayM4_CloseStream(nPort ffcommon.FLong) (res bool) {
	t, _, _ := hcnetsdkcommon.GetHcnetsdkDll().NewProc("PlayM4_CloseStream").Call(
		uintptr(nPort),
	)
	res = ffcommon.GoBool(t)
	return
}

// PLAYM4_API int  __stdcall  PlayM4_GetCaps();
func PlayM4_GetCaps() (res ffcommon.FInt) {
	t, _, _ := hcnetsdkcommon.GetHcnetsdkDll().NewProc("PlayM4_GetCaps").Call()
	res = ffcommon.FInt(t)
	return
}

// PLAYM4_API DWORD __stdcall PlayM4_GetFileTime(LONG nPort);
func PlayM4_GetFileTime(nPort ffcommon.FLong) (res hcnetsdkcommon.FDword) {
	t, _, _ := hcnetsdkcommon.GetHcnetsdkDll().NewProc("PlayM4_GetFileTime").Call(
		uintptr(nPort),
	)
	res = hcnetsdkcommon.FDword(t)
	return
}

// PLAYM4_API DWORD __stdcall PlayM4_GetPlayedTime(LONG nPort);
func PlayM4_GetPlayedTime(nPort ffcommon.FLong) (res hcnetsdkcommon.FDword) {
	t, _, _ := hcnetsdkcommon.GetHcnetsdkDll().NewProc("PlayM4_GetPlayedTime").Call(
		uintptr(nPort),
	)
	res = hcnetsdkcommon.FDword(t)
	return
}

// PLAYM4_API DWORD __stdcall PlayM4_GetPlayedFrames(LONG nPort);
func PlayM4_GetPlayedFrames(nPort ffcommon.FLong) (res hcnetsdkcommon.FDword) {
	t, _, _ := hcnetsdkcommon.GetHcnetsdkDll().NewProc("PlayM4_GetPlayedFrames").Call(
		uintptr(nPort),
	)
	res = hcnetsdkcommon.FDword(t)
	return
}

// 23
// //////////////ver 2.0 added///////////////////////////////////////
// PLAYM4_API BOOL __stdcall    PlayM4_SetDecCallBack(LONG nPort,void (CALLBACK* DecCBFun)(long nPort,char * pBuf,long nSize,FRAME_INFO * pFrameInfo, long nReserved1,long nReserved2));
func PlayM4_SetDecCallBack(nPort ffcommon.FLong, DisplayCBFun func(nPort ffcommon.FLong, pBuf ffcommon.FCharPStruct, nSize ffcommon.FLong, pFrameInfo *FRAME_INFO, nReserved1, nReserved2 ffcommon.FLong) uintptr) (res bool) {
	t, _, _ := hcnetsdkcommon.GetHcnetsdkDll().NewProc("PlayM4_SetDecCallBack").Call(
		uintptr(nPort),
		ffcommon.NewCallback(DisplayCBFun),
	)
	res = ffcommon.GoBool(t)
	return
}

// PLAYM4_API BOOL __stdcall    PlayM4_SetDisplayCallBackYUV(LONG nPort, void (CALLBACK* DisplayCBFun)(DISPLAY_INFO_YUV *pstDisplayInfo), BOOL bTrue, void* pUser);
func PlayM4_SetDisplayCallBackYUV(nPort ffcommon.FLong, DisplayCBFun func(pstDisplayInfo *DISPLAY_INFO_YUV) uintptr, bTrue bool, pUser ffcommon.FVoidP) (res bool) {
	t, _, _ := hcnetsdkcommon.GetHcnetsdkDll().NewProc("PlayM4_SetDisplayCallBackYUV").Call(
		uintptr(nPort),
		ffcommon.NewCallback(DisplayCBFun),
		ffcommon.CBool(bTrue),
		pUser,
	)
	res = ffcommon.GoBool(t)
	return
}

// PLAYM4_API BOOL __stdcall    PlayM4_SetDisplayCallBack(LONG nPort,void (CALLBACK* DisplayCBFun)(long nPort,char * pBuf,long nSize,long nWidth,long nHeight,long nStamp,long nType,long nReserved));
func PlayM4_SetDisplayCallBack(nPort ffcommon.FLong, DisplayCBFun func(nPort ffcommon.FLong, pBuf ffcommon.FCharPStruct, nSize, nWidth, nHeight, nStamp, nType, nReserved ffcommon.FLong) uintptr) (res bool) {
	t, _, _ := hcnetsdkcommon.GetHcnetsdkDll().NewProc("PlayM4_SetDisplayCallBack").Call(
		uintptr(nPort),
		ffcommon.NewCallback(DisplayCBFun),
	)
	res = ffcommon.GoBool(t)
	return
}

// PLAYM4_API BOOL __stdcall    PlayM4_ConvertToBmpFile(char * pBuf,long nSize,long nWidth,long nHeight,long nType,char *sFileName);
func PlayM4_ConvertToBmpFile(pBuf ffcommon.FBuf, nSize, nWidth, nHeight, nType ffcommon.FLong, sFileName ffcommon.FCharP) (res bool) {
	t, _, _ := hcnetsdkcommon.GetHcnetsdkDll().NewProc("PlayM4_ConvertToBmpFile").Call(
		uintptr(unsafe.Pointer(pBuf)),
		uintptr(nSize),
		uintptr(nWidth),
		uintptr(nHeight),
		uintptr(nType),
		ffcommon.UintPtrFromString(sFileName),
	)
	res = ffcommon.GoBool(t)
	return
}

// PLAYM4_API DWORD __stdcall    PlayM4_GetFileTotalFrames(LONG nPort);
func PlayM4_GetFileTotalFrames(nPort ffcommon.FLong) (res hcnetsdkcommon.FDword) {
	t, _, _ := hcnetsdkcommon.GetHcnetsdkDll().NewProc("PlayM4_GetFileTotalFrames").Call(
		uintptr(nPort),
	)
	res = hcnetsdkcommon.FDword(t)
	return
}

// PLAYM4_API DWORD __stdcall    PlayM4_GetCurrentFrameRate(LONG nPort);
func PlayM4_GetCurrentFrameRate(nPort ffcommon.FLong) (res hcnetsdkcommon.FDword) {
	t, _, _ := hcnetsdkcommon.GetHcnetsdkDll().NewProc("PlayM4_GetCurrentFrameRate").Call(
		uintptr(nPort),
	)
	res = hcnetsdkcommon.FDword(t)
	return
}

// PLAYM4_API DWORD __stdcall    PlayM4_GetPlayedTimeEx(LONG nPort);
func PlayM4_GetPlayedTimeEx(nPort ffcommon.FLong) (res hcnetsdkcommon.FDword) {
	t, _, _ := hcnetsdkcommon.GetHcnetsdkDll().NewProc("PlayM4_GetPlayedTimeEx").Call(
		uintptr(nPort),
	)
	res = hcnetsdkcommon.FDword(t)
	return
}

// PLAYM4_API BOOL __stdcall    PlayM4_SetPlayedTimeEx(LONG nPort,DWORD nTime);
func PlayM4_SetPlayedTimeEx(nPort ffcommon.FLong, nTime hcnetsdkcommon.FDword) (res bool) {
	t, _, _ := hcnetsdkcommon.GetHcnetsdkDll().NewProc("PlayM4_SetPlayedTimeEx").Call(
		uintptr(nPort),
		uintptr(nTime),
	)
	res = ffcommon.GoBool(t)
	return
}

// PLAYM4_API DWORD __stdcall    PlayM4_GetCurrentFrameNum(LONG nPort);
func PlayM4_GetCurrentFrameNum(nPort ffcommon.FLong) (res hcnetsdkcommon.FDword) {
	t, _, _ := hcnetsdkcommon.GetHcnetsdkDll().NewProc("PlayM4_GetCurrentFrameNum").Call(
		uintptr(nPort),
	)
	res = hcnetsdkcommon.FDword(t)
	return
}

// PLAYM4_API BOOL __stdcall    PlayM4_SetStreamOpenMode(LONG nPort,DWORD nMode);
func PlayM4_SetStreamOpenMode(nPort ffcommon.FLong, nMode hcnetsdkcommon.FDword) (res bool) {
	t, _, _ := hcnetsdkcommon.GetHcnetsdkDll().NewProc("PlayM4_SetStreamOpenMode").Call(
		uintptr(nPort),
		uintptr(nMode),
	)
	res = ffcommon.GoBool(t)
	return
}

// PLAYM4_API DWORD __stdcall    PlayM4_GetFileHeadLength();
func PlayM4_GetFileHeadLength() (res hcnetsdkcommon.FDword) {
	t, _, _ := hcnetsdkcommon.GetHcnetsdkDll().NewProc("PlayM4_GetFileHeadLength").Call()
	res = hcnetsdkcommon.FDword(t)
	return
}

// PLAYM4_API DWORD __stdcall    PlayM4_GetSdkVersion();
func PlayM4_GetSdkVersion() (res hcnetsdkcommon.FDword) {
	t, _, _ := hcnetsdkcommon.GetHcnetsdkDll().NewProc("PlayM4_GetSdkVersion").Call()
	res = hcnetsdkcommon.FDword(t)
	return
}

// 11
// //////////////ver 2.2 added///////////////////////////////////////
// PLAYM4_API DWORD __stdcall  PlayM4_GetLastError(LONG nPort);
func PlayM4_GetLastError(nPort ffcommon.FLong) (res hcnetsdkcommon.FDword) {
	t, _, _ := hcnetsdkcommon.GetHcnetsdkDll().NewProc("PlayM4_GetLastError").Call(
		uintptr(nPort),
	)
	res = hcnetsdkcommon.FDword(t)
	return
}

// PLAYM4_API BOOL __stdcall PlayM4_RefreshPlay(LONG nPort);
func PlayM4_RefreshPlay(nPort ffcommon.FLong) (res bool) {
	t, _, _ := hcnetsdkcommon.GetHcnetsdkDll().NewProc("PlayM4_RefreshPlay").Call(
		uintptr(nPort),
	)
	res = ffcommon.GoBool(t)
	return
}

// PLAYM4_API BOOL __stdcall PlayM4_SetOverlayMode(LONG nPort,BOOL bOverlay,COLORREF colorKey);
func PlayM4_SetOverlayMode(nPort ffcommon.FLong, bOverlay bool, colorKey hcnetsdkcommon.FColorref) (res bool) {
	t, _, _ := hcnetsdkcommon.GetHcnetsdkDll().NewProc("PlayM4_SetOverlayMode").Call(
		uintptr(nPort),
		ffcommon.CBool(bOverlay),
		uintptr(colorKey),
	)
	res = ffcommon.GoBool(t)
	return
}

// PLAYM4_API BOOL __stdcall PlayM4_GetPictureSize(LONG nPort,LONG *pWidth,LONG *pHeight);
func PlayM4_GetPictureSize(nPort ffcommon.FLong, pWidth, pHeight *ffcommon.FLong) (res bool) {
	t, _, _ := hcnetsdkcommon.GetHcnetsdkDll().NewProc("PlayM4_GetPictureSize").Call(
		uintptr(nPort),
		uintptr(unsafe.Pointer(pWidth)),
		uintptr(unsafe.Pointer(pHeight)),
	)
	res = ffcommon.GoBool(t)
	return
}

// PLAYM4_API BOOL __stdcall PlayM4_SetPicQuality(LONG nPort,BOOL bHighQuality);
func PlayM4_SetPicQuality(nPort ffcommon.FLong, bHighQuality bool) (res bool) {
	t, _, _ := hcnetsdkcommon.GetHcnetsdkDll().NewProc("PlayM4_SetPicQuality").Call(
		uintptr(nPort),
		ffcommon.CBool(bHighQuality),
	)
	res = ffcommon.GoBool(t)
	return
}

// PLAYM4_API BOOL __stdcall PlayM4_PlaySoundShare(LONG nPort);
func PlayM4_PlaySoundShare(nPort ffcommon.FLong) (res bool) {
	t, _, _ := hcnetsdkcommon.GetHcnetsdkDll().NewProc("PlayM4_PlaySoundShare").Call(
		uintptr(nPort),
	)
	res = ffcommon.GoBool(t)
	return
}

// PLAYM4_API BOOL __stdcall PlayM4_StopSoundShare(LONG nPort);
func PlayM4_StopSoundShare(nPort ffcommon.FLong) (res bool) {
	t, _, _ := hcnetsdkcommon.GetHcnetsdkDll().NewProc("PlayM4_StopSoundShare").Call(
		uintptr(nPort),
	)
	res = ffcommon.GoBool(t)
	return
}

// //7
// ////////////////ver 2.4 added///////////////////////////////////////
// PLAYM4_API LONG __stdcall PlayM4_GetStreamOpenMode(LONG nPort);
// PLAYM4_API LONG __stdcall PlayM4_GetOverlayMode(LONG nPort);
// PLAYM4_API COLORREF __stdcall PlayM4_GetColorKey(LONG nPort);
// PLAYM4_API WORD __stdcall PlayM4_GetVolume(LONG nPort);
// PLAYM4_API BOOL __stdcall PlayM4_GetPictureQuality(LONG nPort,BOOL *bHighQuality);
// PLAYM4_API DWORD __stdcall PlayM4_GetSourceBufferRemain(LONG nPort);
// PLAYM4_API BOOL __stdcall PlayM4_ResetSourceBuffer(LONG nPort);
// PLAYM4_API BOOL __stdcall PlayM4_SetSourceBufCallBack(LONG nPort,DWORD nThreShold,void (CALLBACK * SourceBufCallBack)(long nPort,DWORD nBufSize,DWORD dwUser,void*pResvered),DWORD dwUser,void *pReserved);
// PLAYM4_API BOOL __stdcall PlayM4_ResetSourceBufFlag(LONG nPort);
// PLAYM4_API BOOL __stdcall PlayM4_SetDisplayBuf(LONG nPort,DWORD nNum);
// PLAYM4_API DWORD __stdcall PlayM4_GetDisplayBuf(LONG nPort);
// PLAYM4_API BOOL __stdcall PlayM4_OneByOneBack(LONG nPort);
// PLAYM4_API BOOL __stdcall PlayM4_SetFileRefCallBack(LONG nPort, void (__stdcall *pFileRefDone)(DWORD nPort,DWORD nUser),DWORD nUser);
// PLAYM4_API BOOL __stdcall PlayM4_SetCurrentFrameNum(LONG nPort,DWORD nFrameNum);
// PLAYM4_API BOOL __stdcall PlayM4_GetKeyFramePos(LONG nPort,DWORD nValue, DWORD nType, PFRAME_POS pFramePos);
// PLAYM4_API BOOL __stdcall PlayM4_GetNextKeyFramePos(LONG nPort,DWORD nValue, DWORD nType, PFRAME_POS pFramePos);
// #if (WINVER >= 0x0400)
// //Note: These funtion must be builded under win2000 or above with Microsoft Platform sdk.
// //        You can download the sdk from "http://www.microsoft.com/msdownload/platformsdk/sdkupdate/";
// PLAYM4_API BOOL __stdcall PlayM4_InitDDrawDevice();
// PLAYM4_API void __stdcall PlayM4_ReleaseDDrawDevice();
// PLAYM4_API DWORD __stdcall PlayM4_GetDDrawDeviceTotalNums();
// PLAYM4_API BOOL __stdcall PlayM4_SetDDrawDevice(LONG nPort,DWORD nDeviceNum);
// //PLAYM4_API BOOL __stdcall PlayM4_GetDDrawDeviceInfo(DWORD nDeviceNum,LPSTR  lpDriverDescription,DWORD nDespLen,LPSTR lpDriverName ,DWORD nNameLen,HMONITOR *hhMonitor);
// PLAYM4_API int  __stdcall  PlayM4_GetCapsEx(DWORD nDDrawDeviceNum);
// #endif
// PLAYM4_API BOOL __stdcall PlayM4_ThrowBFrameNum(LONG nPort,DWORD nNum);
// //23
// ////////////////ver 2.5 added///////////////////////////////////////
// PLAYM4_API BOOL __stdcall PlayM4_SetDisplayType(LONG nPort,LONG nType);
// PLAYM4_API long __stdcall PlayM4_GetDisplayType(LONG nPort);
// //2
// ////////////////ver 3.0 added///////////////////////////////////////
// PLAYM4_API BOOL __stdcall PlayM4_SetDecCBStream(LONG nPort,DWORD nStream);
// PLAYM4_API BOOL __stdcall PlayM4_SetDisplayRegion(LONG nPort,DWORD nRegionNum, RECT *pSrcRect, HWND hDestWnd, BOOL bEnable);
// PLAYM4_API BOOL __stdcall PlayM4_RefreshPlayEx(LONG nPort,DWORD nRegionNum);
// #if (WINVER >= 0x0400)
// //Note: The funtion must be builded under win2000 or above with Microsoft Platform sdk.
// //        You can download the sdk from http://www.microsoft.com/msdownload/platformsdk/sdkupdate/;
// PLAYM4_API BOOL __stdcall PlayM4_SetDDrawDeviceEx(LONG nPort,DWORD nRegionNum,DWORD nDeviceNum);
// #endif
// //4
// /////////////////v3.2 added/////////////////////////////////////////

// PLAYM4_API BOOL __stdcall PlayM4_GetRefValue(LONG nPort,BYTE *pBuffer, DWORD *pSize);
// PLAYM4_API BOOL __stdcall PlayM4_SetRefValue(LONG nPort,BYTE *pBuffer, DWORD nSize);
// PLAYM4_API BOOL __stdcall PlayM4_OpenStreamEx(LONG nPort,PBYTE pFileHeadBuf,DWORD nSize,DWORD nBufPoolSize);
// PLAYM4_API BOOL __stdcall PlayM4_CloseStreamEx(LONG nPort);
// PLAYM4_API BOOL __stdcall PlayM4_InputVideoData(LONG nPort,PBYTE pBuf,DWORD nSize);
// PLAYM4_API BOOL __stdcall PlayM4_InputAudioData(LONG nPort,PBYTE pBuf,DWORD nSize);
// PLAYM4_API BOOL __stdcall PlayM4_RegisterDrawFun(LONG nPort,void (CALLBACK* DrawFun)(long nPort,HDC hDc,LONG nUser),LONG nUser);
// PLAYM4_API BOOL __stdcall PlayM4_RigisterDrawFun(LONG nPort,void (CALLBACK* DrawFun)(long nPort,HDC hDc,LONG nUser),LONG nUser);
// //8
// //////////////////v3.4/////////////////////////////////////////////////////
// PLAYM4_API BOOL __stdcall PlayM4_SetTimerType(LONG nPort,DWORD nTimerType,DWORD nReserved);
// PLAYM4_API BOOL __stdcall PlayM4_GetTimerType(LONG nPort,DWORD *pTimerType,DWORD *pReserved);
// PLAYM4_API BOOL __stdcall PlayM4_ResetBuffer(LONG nPort,DWORD nBufType);
// PLAYM4_API DWORD __stdcall PlayM4_GetBufferValue(LONG nPort,DWORD nBufType);

// //////////////////V3.6/////////////////////////////////////////////////////////
// PLAYM4_API BOOL __stdcall PlayM4_AdjustWaveAudio(LONG nPort,LONG nCoefficient);
// PLAYM4_API BOOL __stdcall PlayM4_SetVerifyCallBack(LONG nPort, DWORD nBeginTime, DWORD nEndTime, void (__stdcall * funVerify)(long nPort, FRAME_POS * pFilePos, DWORD bIsVideo, DWORD nUser),  DWORD  nUser);
// PLAYM4_API BOOL __stdcall PlayM4_SetAudioCallBack(LONG nPort, void (__stdcall * funAudio)(long nPort, char * pAudioBuf, long nSize, long nStamp, long nType, long nUser), long nUser);
// PLAYM4_API BOOL __stdcall PlayM4_SetEncTypeChangeCallBack(LONG nPort,void(CALLBACK *funEncChange)(long nPort,long nUser),long nUser);
// PLAYM4_API BOOL __stdcall PlayM4_SetColor(LONG nPort, DWORD nRegionNum, int nBrightness, int nContrast, int nSaturation, int nHue);
// PLAYM4_API BOOL __stdcall PlayM4_GetColor(LONG nPort, DWORD nRegionNum, int *pBrightness, int *pContrast, int *pSaturation, int *pHue);
// PLAYM4_API BOOL __stdcall PlayM4_SetEncChangeMsg(LONG nPort,HWND hWnd,UINT nMsg);
// PLAYM4_API BOOL __stdcall PlayM4_GetOriginalFrameCallBack(LONG nPort, BOOL bIsChange,BOOL bNormalSpeed,long nStartFrameNum,long nStartStamp,long nFileHeader,void(CALLBACK *funGetOrignalFrame)(long nPort,FRAME_TYPE *frameType, long nUser),long nUser);
// PLAYM4_API BOOL __stdcall PlayM4_GetFileSpecialAttr(LONG nPort, DWORD *pTimeStamp,DWORD *pFileNum ,DWORD *pReserved);
// PLAYM4_API DWORD __stdcall PlayM4_GetSpecialData(LONG nPort);
// PLAYM4_API BOOL __stdcall PlayM4_SetCheckWatermarkCallBack(LONG nPort,void(CALLBACK* funCheckWatermark)(long nPort,WATERMARK_INFO* pWatermarkInfo,DWORD nUser),DWORD nUser);
// PLAYM4_API BOOL __stdcall PlayM4_SetImageSharpen(LONG nPort,DWORD nLevel);
// PLAYM4_API BOOL __stdcall PlayM4_SetDecodeFrameType(LONG nPort,DWORD nFrameType);
// PLAYM4_API BOOL __stdcall PlayM4_SetPlayMode(LONG nPort,BOOL bNormal);
// PLAYM4_API BOOL __stdcall PlayM4_SetOverlayFlipMode(LONG nPort,BOOL bTrue);
// PLAYM4_API BOOL __stdcall PlayM4_SetOverlayPriInfoFlag(LONG nPort, DWORD nIntelType, BOOL bTrue,const char *pFontPath);

// //PLAYM4_API DWORD __stdcall PlayM4_GetAbsFrameNum(LONG nPort);

// //////////////////V4.7.0.0//////////////////////////////////////////////////////
// ////convert yuv to jpeg
// PLAYM4_API BOOL __stdcall PlayM4_ConvertToJpegFile(char * pBuf,long nSize,long nWidth,long nHeight,long nType,char *sFileName);
// PLAYM4_API BOOL __stdcall PlayM4_SetJpegQuality(long nQuality);
// //set deflash
// PLAYM4_API BOOL __stdcall PlayM4_SetDeflash(LONG nPort,BOOL bDefalsh);
// //PLAYM4_API BOOL __stdcall PlayM4_SetDecCallBackEx(LONG nPort,void (CALLBACK* DecCBFun)(long nPort,char * pBuf,long nSize,FRAME_INFO * pFrameInfo, long nReserved1,long nReserved2), char* pDest, long nDestSize);
// //////////////////V4.8.0.0/////////////////////////////////////////////////////////
// //check discontinuous frame number as error data?
// PLAYM4_API BOOL __stdcall PlayM4_CheckDiscontinuousFrameNum(LONG nPort, BOOL bCheck);
// //get bmp or jpeg
// PLAYM4_API BOOL __stdcall PlayM4_GetBMP(LONG nPort,PBYTE pBitmap,DWORD nBufSize,DWORD* pBmpSize);
// PLAYM4_API BOOL __stdcall PlayM4_GetJPEG(LONG nPort,PBYTE pJpeg,DWORD nBufSize,DWORD* pJpegSize);
// //dec call back mend
// PLAYM4_API BOOL __stdcall PlayM4_SetDecCallBackMend(LONG nPort,void (CALLBACK* DecCBFun)(long nPort,char * pBuf,long nSize,FRAME_INFO * pFrameInfo, long nUser,long nReserved2), long nUser);
// PLAYM4_API BOOL __stdcall PlayM4_SetSecretKey(LONG nPort, LONG lKeyType, char *pSecretKey, LONG lKeyLen);

// // add by gb 2007-12-23
// PLAYM4_API BOOL __stdcall PlayM4_SetFileEndCallback(LONG nPort, void(CALLBACK*FileEndCallback)(long nPort, void *pUser), void *pUser);

// // add by gb 080131 version 4.9.0.1
// PLAYM4_API BOOL __stdcall PlayM4_GetPort(LONG* nPort);
// PLAYM4_API BOOL __stdcall PlayM4_FreePort(LONG nPort);
// PLAYM4_API BOOL __stdcall PlayM4_SetDisplayCallBackEx(LONG nPort,void (CALLBACK* DisplayCBFun)(DISPLAY_INFO *pstDisplayInfo), long nUser);
// PLAYM4_API BOOL __stdcall PlayM4_SkipErrorData(LONG nPort, BOOL bSkip);
// PLAYM4_API BOOL __stdcall PlayM4_SetDecCallBackExMend(LONG nPort, void (CALLBACK* DecCBFun)(long nPort, char* pBuf, long nSize, FRAME_INFO* pFrameInfo,
//                                                       long nUser, long nReserved2), char* pDest, long nDestSize, long nUser);
// //reverse play add by chenjie 110609
// PLAYM4_API BOOL __stdcall PlayM4_ReversePlay(LONG nPort);
// PLAYM4_API BOOL __stdcall PlayM4_GetSystemTime(LONG nPort, PLAYM4_SYSTEM_TIME *pstSystemTime);

// //PLAYM4_API BOOL __stdcall PlayM4_SetDecodeERC(long nPort, unsigned int nLevel);

// #ifndef PLAYM4_SESSION_INFO_TAG
// const PLAYM4_SESSION_INFO_TAG
// //nProtocolType
// const PLAYM4_PROTOCOL_RTSP    1
// //nSessionInfoType
// const PLAYM4_SESSION_INFO_SDP 1

// typedef struct _PLAYM4_SESSION_INFO_     //交互信息结构
// {
//       int            nSessionInfoType;   //交互信息类型，比如SDP，比如海康私有信息头
//       int            nSessionInfoLen;    //交互信息长度
//       unsigned char* pSessionInfoData;   //交互信息数据

// } PLAYM4_SESSION_INFO;
// #endif

// PLAYM4_API BOOL __stdcall PlayM4_OpenStreamAdvanced(LONG nPort, int nProtocolType, PLAYM4_SESSION_INFO* pstSessionInfo, DWORD nBufPoolSize);

// const R_ANGLE_0   -1  //不旋转
// const R_ANGLE_L90  0  //向左旋转90度
// const R_ANGLE_R90  1  //向右旋转90度
// const R_ANGLE_180  2  //旋转180度

// PLAYM4_API BOOL __stdcall PlayM4_SetRotateAngle(LONG nPort, DWORD nRegionNum, DWORD dwType);

// #ifndef PLAYM4_ADDITION_INFO_TAG
// const PLAYM4_ADDITION_INFO_TAG
// typedef struct _PLAYM4_ADDITION_INFO_     //交互信息结构
// {
//     BYTE*   pData;            //附件数据
//     DWORD   dwDatalen;        //附件数据长度
//     DWORD    dwDataType;        //数据类型
//     DWORD    dwTimeStamp;    //相对时间戳
// } PLAYM4_ADDITION_INFO;
// #endif

// //dwGroupIndex 暂约定取值0~3，第一版本取消同步只能同个closestream处理
// PLAYM4_API BOOL __stdcall PlayM4_SetSycGroup(LONG nPort, DWORD dwGroupIndex);
// //暂不实现此函数，同个组设置的起始时间不一致，以最小的时间作为播放起点，同一组可只设一路
// PLAYM4_API BOOL __stdcall PlayM4_SetSycStartTime(LONG nPort, PLAYM4_SYSTEM_TIME *pstSystemTime);

// // 以下实现鱼眼相关的接口
// #ifndef FISH_EYE_TAG
// const FISH_EYE_TAG

// // 安装类型
// typedef enum tagFECPlaceType
// {
//     FEC_PLACE_WALL    = 0x1,        // 壁装方式        (法线水平)
//     FEC_PLACE_FLOOR   = 0x2,        // 地面安装        (法线向上)
//     FEC_PLACE_CEILING = 0x3,        // 顶装方式        (法线向下)

// }FECPLACETYPE;

// typedef enum tagFECCorrectType
// {
//     FEC_CORRECT_PTZ = 0x100,        // PTZ
//     FEC_CORRECT_180 = 0x200,        // 180度矫正  （对应2P）
//     FEC_CORRECT_360 = 0x300,        // 360全景矫正 （对应1P）
//     FEC_CORRECT_LAT = 0x400         //纬度展开

// }FECCORRECTTYPE;

// typedef struct tagCycleParam
// {
//     float    fRadiusLeft;    // 圆的最左边X坐标
//     float    fRadiusRight;    // 圆的最右边X坐标
//     float   fRadiusTop;        // 圆的最上边Y坐标
//     float   fRadiusBottom;    // 圆的最下边Y坐标

// }CYCLEPARAM;

// typedef struct tagPTZParam
// {
//     float fPTZPositionX;        // PTZ 显示的中心位置 X坐标
//     float fPTZPositionY;        // PTZ 显示的中心位置 Y坐标

// }PTZPARAM;

// // 错误码
// /*********************************************

//  ********************************************/

// // 更新标记变量定义

// const         FEC_UPDATE_RADIUS             0x1
// const         FEC_UPDATE_PTZZOOM             0x2
// const         FEC_UPDATE_WIDESCANOFFSET     0x4
// const         FEC_UPDATE_PTZPARAM             0x8

// typedef struct tagFECParam
// {

//     unsigned int     nUpDateType;            // 更新的类型

//     unsigned int    nPlaceAndCorrect;        // 安装方式和矫正方式，只能用于获取，SetParam的时候无效,该值表示安装方式和矫正方式的和

//     PTZPARAM        stPTZParam;                // PTZ 校正的参数

//     CYCLEPARAM        stCycleParam;            // 鱼眼图像圆心参数

//     float            fZoom;                    // PTZ 显示的范围参数

//     float            fWideScanOffset;        // 180或者360度校正的偏移角度

//     int                nResver[16];            // 保留字段

// }FISHEYEPARAM;

// typedef void (__stdcall * FISHEYE_CallBack )( void* pUser  , unsigned int  nSubPort , unsigned int nCBType , void * hDC ,   unsigned int nWidth , unsigned int nHeight);

// #endif
// // 启用鱼眼
// PLAYM4_API BOOL __stdcall PlayM4_FEC_Enable(LONG nPort);

// // 关闭鱼眼模块
// PLAYM4_API BOOL __stdcall PlayM4_FEC_Disable(LONG nPort);

// // 获取鱼眼矫正处理子端口 [1~31]
// PLAYM4_API BOOL  __stdcall PlayM4_FEC_GetPort(LONG nPort, unsigned int* nSubPort,FECPLACETYPE emPlaceType,FECCORRECTTYPE emCorrectType);

// // 删除鱼眼矫正处理子端口
// PLAYM4_API BOOL __stdcall PlayM4_FEC_DelPort(LONG nPort , unsigned int nSubPort);

// // 设置鱼眼矫正参数
// PLAYM4_API BOOL __stdcall PlayM4_FEC_SetParam(LONG nPort , unsigned int nSubPort , FISHEYEPARAM * pPara);

// // 获取鱼眼矫正参数
// PLAYM4_API BOOL __stdcall PlayM4_FEC_GetParam(LONG nPort , unsigned int nSubPort , FISHEYEPARAM * pPara);

// // 设置显示窗口，可以随时切换
// PLAYM4_API BOOL __stdcall PlayM4_FEC_SetWnd(LONG nPort , unsigned int nSubPort , void * hWnd);

// // 设置鱼眼窗口的绘图回调
// PLAYM4_API BOOL __stdcall PlayM4_FEC_SetCallBack(LONG nPort , unsigned int nSubPort , FISHEYE_CallBack cbFunc , void * pUser);

// //motionflow
// PLAYM4_API BOOL __stdcall PlayM4_MotionFlow(LONG nPort, DWORD dwAdjustType);

// //图像增强相关
// #ifndef PLAYM4_HIKVIE_TAG
// const PLAYM4_HIKVIE_TAG

// typedef struct _PLAYM4_VIE_DYNPARAM_
// {
//     int moduFlag;      //启用的算法处理模块，在PLAYM4_VIE_MODULES中定义
//     //如 PLAYM4_VIE_MODU_ADJ | PLAYM4_VIE_MODU_EHAN
//     //模块启用后，必须设置相应的参数；
//     //PLAYM4_VIE_MODU_ADJ
//     int brightVal;     //亮度调节值，[-255, 255]
//     int contrastVal;   //对比度调节值，[-256, 255]
//     int colorVal;      //饱和度调节值，[-256, 255]
//     //PLAYM4_VIE_MODU_EHAN
//     int toneScale;     //滤波范围，[0, 100]
//     int toneGain;      //对比度调节，全局对比度增益值，[-256, 255]
//     int toneOffset;    //亮度调节，亮度平均值偏移，[-255, 255]
//     int toneColor;     //颜色调节，颜色保真值，[-256, 255]
//     //PLAYM4_VIE_MODU_DEHAZE
//     int dehazeLevel;   //去雾强度，[0, 255]
//     int dehazeTrans;   //透射值，[0, 255]
//     int dehazeBright;  //亮度补偿，[0, 255]
//     //PLAYM4_VIE_MODU_DENOISE
//     int denoiseLevel;  //去噪强度，[0, 255]
//     //PLAYM4_VIE_MODU_SHARPEN
//     int usmAmount;     //锐化强度，[0, 255]
//     int usmRadius;     //锐化半径，[1, 15]
//     int usmThreshold;  //锐化阈值，[0, 255]
//     //PLAYM4_VIE_MODU_DEBLOCK
//     int deblockLevel;  //去块强度，[0, 100]
//     //PLAYM4_VIE_MODU_LENS
//     int lensWarp;      //畸变量，[-256, 255]
//     int lensZoom;      //缩放量，[-256, 255]
//     //PLAYM4_VIE_MODU_CRB
//     //无响应参数
// } PLAYM4_VIE_PARACONFIG;

// typedef enum _PLAYM4_VIE_MODULES
// {
//     PLAYM4_VIE_MODU_ADJ      = 0x00000001, //图像基本调节
//     PLAYM4_VIE_MODU_EHAN     = 0x00000002, //局部增强模块
//     PLAYM4_VIE_MODU_DEHAZE   = 0x00000004, //去雾模块
//     PLAYM4_VIE_MODU_DENOISE  = 0x00000008, //去噪模块
//     PLAYM4_VIE_MODU_SHARPEN  = 0x00000010, //锐化模块
//     PLAYM4_VIE_MODU_DEBLOCK  = 0x00000020, //去块滤波模块
//     PLAYM4_VIE_MODU_CRB      = 0x00000040, //色彩平衡模块
//     PLAYM4_VIE_MODU_LENS     = 0x00000080, //镜头畸变矫正模块
// }PLAYM4_VIE_MODULES;
// #endif

// //设置关闭/开启模块
// //dwModuFlag对应PLAYM4_VIE_MODULES宏,可组合
// //先设置模块开启，再设置模块参数；期间采用默认的参数;
// //关闭模块后，上次设置的参数清空
// //其他接口调用，必须在该接口开启模块后；否则，返回错误
// PLAYM4_API BOOL __stdcall PlayM4_VIE_SetModuConfig(LONG lPort,int nModuFlag,BOOL bEnable);

// //设置图像增强区域，NULL全图；超过全图，采用全图；最小区域16*16像素
// //可支持设置区域，最多比较说4个，第一个版本可以只支持一个。多个区域要求不能重叠，有重叠就报错
// PLAYM4_API BOOL __stdcall PlayM4_VIE_SetRegion(LONG lPort,LONG lRegNum,RECT* pRect);

// //获取开启模块
// PLAYM4_API BOOL __stdcall PlayM4_VIE_GetModuConfig(LONG lPort,int* pdwModuFlag);

// //设置参数
// //未开启模块的参数设置被忽略
// PLAYM4_API BOOL __stdcall PlayM4_VIE_SetParaConfig(LONG lPort,PLAYM4_VIE_PARACONFIG* pParaConfig);

// //获取开启模块的参数
// PLAYM4_API BOOL __stdcall PlayM4_VIE_GetParaConfig(LONG lPort,PLAYM4_VIE_PARACONFIG* pParaConfig);

// //音视频同步接口
// PLAYM4_API BOOL __stdcall PlayM4_SyncToAudio(LONG nPort, BOOL bSyncToAudio);

// // 私有信息模块类型
// typedef enum _PLAYM4_PRIDATA_RENDER
// {
//     PLAYM4_RENDER_ANA_INTEL_DATA   = 0x00000001, //智能分析
//     PLAYM4_RENDER_MD               = 0x00000002, //移动侦测
//     PLAYM4_RENDER_ADD_POS          = 0x00000004, //POS信息后叠加
//     PLAYM4_RENDER_ADD_PIC          = 0x00000008, //图片叠加
//     PLAYM4_RENDER_FIRE_DETCET      = 0x00000010,  //热成像信息
//     PLAYM4_RENDER_TEM              = 0x00000020,   //温度信息
//     PLAYM4_RENDER_TRACK_TEM        = 0x00000040,    //轨迹信息
//     PLAYM4_RENDER_THERMAL          = 0x00000080 //废气检测和烟火屏蔽信息
// }PLAYM4_PRIDATA_RENDER;

// typedef enum _PLAYM4_THERMAL_FLAG
// {
//     PLAYM4_THERMAL_FIREMASK = 0x00000001, //烟火屏蔽
//     PLAYM4_THERMAL_RULEGAS = 0x00000002, //规则废气检测
//     PLAYM4_THERMAL_TARGETGAS = 0x00000004 //目标废气检测
// }PLAYM4_THERMAL_FLAG;

// typedef enum _PLAYM4_FIRE_ALARM{
//     PLAYM4_FIRE_FRAME_DIS           = 0x00000001, //火点框显示
//     PLAYM4_FIRE_MAX_TEMP            = 0x00000002, //最高温度
//     PLAYM4_FIRE_MAX_TEMP_POSITION   = 0x00000004, //最高温度位置显示
//     PLAYM4_FIRE_DISTANCE            = 0x00000008, //最高温度距离}PLAYM4_FIRE_ALARM
// }PLAYM4_FIRE_ALARM;

// typedef enum _PLAYM4_TEM_FLAG{
//     PLAYM4_TEM_REGION_BOX           = 0x00000001, //框测温
//     PLAYM4_TEM_REGION_LINE          = 0x00000002, //线测温
//     PLAYM4_TEM_REGION_POINT         = 0x00000004, //点测温}PLAYM4_TEM_FLAG
// }PLAYM4_TEM_FLAG;

// typedef enum _PLAYM4_TRACK_FLAG
// {
//     PLAYM4_TRACK_PEOPLE = 0x00000001, //人轨迹
//     PLAYM4_TRACK_VEHICLE = 0x00000002, //车轨迹
// }PLAYM4_TRACK_FLAG;

// typedef struct TI_PTZ_INFO_
// {
//     unsigned short dwDefVer;    //结构体版本
//     unsigned short dwLength;    //PTZ_info长度，以8字节为单位
//     DWORD          dwP;    //P（0~3600）
//     DWORD          dwT;         //T（0~3600）
//     DWORD          dwZ;         //Z（0~3600）
//     BYTE        chFSMState; //跟踪状态
//     BYTE           bClearFocusState; //聚焦清晰状态（0,1）
//     BYTE        reserved[6]; //6个字节保留
// }PTZ_INFO;

// // 智能信息开关
// PLAYM4_API BOOL __stdcall PlayM4_RenderPrivateData(LONG nPort, int nIntelType, BOOL bTrue);

// PLAYM4_API BOOL __stdcall PlayM4_RenderPrivateDataEx(LONG nPort, int nIntelType, int nSubType, BOOL bTrue);

// // 加密码流回调,nType=0表示码流加密标记位发生变化就回调，nType=1表示码流有加密位发生回调
// PLAYM4_API BOOL __stdcall PlayM4_SetEncryptTypeCallBack(LONG nPort, DWORD nType,
//                                                         void (CALLBACK* EncryptTypeCBFun)(long nPort, ENCRYPT_INFO* pEncryptInfo, long nUser, long nReserved2), long nUser);
// //lType: 1 表示获取当前显示帧PTZ信息。以特定结构体形式存储在pInfo内，plLen返回长度信息
// PLAYM4_API BOOL __stdcall PlayM4_GetStreamAdditionalInfo(LONG nPort, LONG lType, BYTE* pInfo, LONG* plLen);

// #endif //_PLAYM4_H_
