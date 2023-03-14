package hcnetsdk

import (
	"github.com/moonfdd/ffmpeg-go/ffcommon"
	"github.com/moonfdd/hcnetsdk-go/hcnetsdkcommon"
)

// #ifndef DATA_TYPE_H
// const DATA_TYPE_H
const FRAME_HEAD_MAGIC = 0x03211546
const SYSTEM_SYNC_ID = 2

// #ifdef __LINUX__
// typedef unsigned char            UCHAR;
// typedef unsigned char*          PBYTE;
// typedef char*                    LPTSTR;
// typedef    unsigned short             USHORT;
// typedef int                     HANDLE;
// typedef unsigned long           ULONG;
// typedef unsigned long           DWORD;
// #endif //#ifdef __LINUX__

type TMFRAME_HEADER struct {
	SyncId      ffcommon.FUnsignedLong /* 00000000000000000000000000010b */
	Magic       ffcommon.FUnsignedLong
	FrameType   hcnetsdkcommon.FUnsignedShort /* I frames , P frames or BBP frames Audio frames or dsp status etc */
	Length      ffcommon.FUnsignedLong        /*lenth include this header */
	FrameNumber ffcommon.FUnsignedLong        /* serial number of this frame */
	Breakable   hcnetsdkcommon.FUnsignedChar  /* indicate if stream breakable, you could restart new file(with PktSysHeader) if true */
	/*ULONG   Ack;*/
	PTS ffcommon.FUnsignedLong /* system clock when this frames is processed */
}

type VideoStandard_t ffcommon.FEnum

const (
	StandardNone  = 0x80000000
	StandardNTSC  = 0x00000001
	StandardPAL   = 0x00000002
	StandardSECAM = 0x00000004
)

type FrameType_t ffcommon.FEnum

const (
	PktError           = 0
	PktIFrames         = 0x0001
	PktPFrames         = 0x0002
	PktBBPFrames       = 0x0004
	PktAudioFrames     = 0x0008
	PktMotionDetection = 0x00010
	PktDspStatus       = 0x00020
	PktOrigImage       = 0x00040
	PktSysHeader       = 0x00080
	PktBPFrames        = 0x00100
	PktSFrames         = 0x00200
	PktSubIFrames      = 0x00400
	PktSubPFrames      = 0x00800
	PktSubBBPFrames    = 0x01000
	PktSubSysHeader    = 0x02000
)

type VERSION_INFO struct {
	DspVersion, DspBuildNum       ffcommon.FUnsignedLong
	DriverVersion, DriverBuildNum ffcommon.FUnsignedLong
	SDKVersion, SDKBuildNum       ffcommon.FUnsignedLong
}

type PictureFormat_t = ffcommon.FEnum

const (
	ENC_CIF_FORMAT      = 0
	ENC_QCIF_FORMAT     = 1
	ENC_2CIF_FORMAT     = 2
	ENC_4CIF_FORMAT     = 3
	ENC_QQCIF_FORMAT    = 4
	ENC_CIFQCIF_FORMAT  = 5
	ENC_CIFQQCIF_FORMAT = 6
	ENC_DCIF_FORMAT     = 7
	ENC_VGA_FORMAT      = 8
)

type MOTION_DATA_HEADER struct {
	PicFormat      PictureFormat_t
	HorizeBlocks   ffcommon.FUnsignedLong
	VerticalBlocks ffcommon.FUnsignedLong
	BlockSize      ffcommon.FUnsignedLong
}

const _OSD_BASE = 0x9000                 /*base address of special character*/
const _OSD_YEAR4 = (_OSD_BASE + 0)       /*show year time by length of 4 , for example: 2005*/
const _OSD_YEAR2 = (_OSD_BASE + 1)       /*show year time by length of 2, for example: 05    */
const _OSD_MONTH3 = (_OSD_BASE + 2)      /*show month time in English, for example: Jan*/
const _OSD_MONTH2 = (_OSD_BASE + 3)      /*show month time by two Arabic numerals, for example: 07*/
const _OSD_DAY = (_OSD_BASE + 4)         /*show day time by two Arabic numerals, for example: 31*/
const _OSD_WEEK3 = (_OSD_BASE + 5)       /*show week time in English: MON～SUN*/
const _OSD_CWEEK1 = (_OSD_BASE + 6)      /*show week time in Chinese GB code*/
const _OSD_HOUR24 = (_OSD_BASE + 7)      /*show 24 hours clock: 00～23    */
const _OSD_HOUR12 = (_OSD_BASE + 8)      /*show 12 hours clock: 00～12*/
const _OSD_MINUTE = (_OSD_BASE + 9)      /*show minute time by length of 2: 00～59*/
const _OSD_SECOND = (_OSD_BASE + 10)     /*show second time by length of 2: 00～59*/
const _OSD_MILISECOND = (_OSD_BASE + 11) /*show millisecond time by length of 3: 000~999*/
const _OSD_APM = (_OSD_BASE + 14)        /*show a.m. or p.m. by length of 2 bit, AM or PM*/

// For new added APIs to set OSD: SetEncoderOsdDisplayMode, SetDecoderOsdDisplayMode and SetDisplayOsdDisplayMode in v6.0 SDK, we use new basic address of special character.
const _OSD_BASE_EX = 0x90000                   /*base address of special character*/
const _OSD_YEAR4_EX = (_OSD_BASE_EX + 0)       /*the same as _OSD_YEAR4*/
const _OSD_YEAR2_EX = (_OSD_BASE_EX + 1)       /*the same as _OSD_YEAR2*/
const _OSD_MONTH3_EX = (_OSD_BASE_EX + 2)      /*the same as _OSD_MONTH3*/
const _OSD_MONTH2_EX = (_OSD_BASE_EX + 3)      /*the same as _OSD_MONTH2*/
const _OSD_DAY_EX = (_OSD_BASE_EX + 4)         /*the same as _OSD_DAY*/
const _OSD_WEEK3_EX = (_OSD_BASE_EX + 5)       /*the same as _OSD_WEEK3*/
const _OSD_CWEEK1_EX = (_OSD_BASE_EX + 6)      /*the same as _OSD_CWEEK1*/
const _OSD_HOUR24_EX = (_OSD_BASE_EX + 7)      /*the same as _OSD_HOUR24*/
const _OSD_HOUR12_EX = (_OSD_BASE_EX + 8)      /*the same as _OSD_HOUR12*/
const _OSD_MINUTE_EX = (_OSD_BASE_EX + 9)      /*the same as _OSD_MINUTE*/
const _OSD_SECOND_EX = (_OSD_BASE_EX + 10)     /*the same as _OSD_SECOND*/
const _OSD_MILISECOND_EX = (_OSD_BASE_EX + 11) /*the same as _OSD_MILISECOND*/
const _OSD_APM_EX = (_OSD_BASE_EX + 14)        /*the same as _OSD_APM*/

// #endif
