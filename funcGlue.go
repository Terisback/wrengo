package wrengo

/*
#include "wren.h"
extern void f0(WrenVM*);
extern void f1(WrenVM*);
extern void f2(WrenVM*);
extern void f3(WrenVM*);
extern void f4(WrenVM*);
extern void f5(WrenVM*);
extern void f6(WrenVM*);
extern void f7(WrenVM*);
extern void f8(WrenVM*);
extern void f9(WrenVM*);
extern void f10(WrenVM*);
extern void f11(WrenVM*);
extern void f12(WrenVM*);
extern void f13(WrenVM*);
extern void f14(WrenVM*);
extern void f15(WrenVM*);
extern void f16(WrenVM*);
extern void f17(WrenVM*);
extern void f18(WrenVM*);
extern void f19(WrenVM*);
extern void f20(WrenVM*);
extern void f21(WrenVM*);
extern void f22(WrenVM*);
extern void f23(WrenVM*);
extern void f24(WrenVM*);
extern void f25(WrenVM*);
extern void f26(WrenVM*);
extern void f27(WrenVM*);
extern void f28(WrenVM*);
extern void f29(WrenVM*);
extern void f30(WrenVM*);
extern void f31(WrenVM*);
extern void f32(WrenVM*);
extern void f33(WrenVM*);
extern void f34(WrenVM*);
extern void f35(WrenVM*);
extern void f36(WrenVM*);
extern void f37(WrenVM*);
extern void f38(WrenVM*);
extern void f39(WrenVM*);
extern void f40(WrenVM*);
extern void f41(WrenVM*);
extern void f42(WrenVM*);
extern void f43(WrenVM*);
extern void f44(WrenVM*);
extern void f45(WrenVM*);
extern void f46(WrenVM*);
extern void f47(WrenVM*);
extern void f48(WrenVM*);
extern void f49(WrenVM*);
extern void f50(WrenVM*);
extern void f51(WrenVM*);
extern void f52(WrenVM*);
extern void f53(WrenVM*);
extern void f54(WrenVM*);
extern void f55(WrenVM*);
extern void f56(WrenVM*);
extern void f57(WrenVM*);
extern void f58(WrenVM*);
extern void f59(WrenVM*);
extern void f60(WrenVM*);
extern void f61(WrenVM*);
extern void f62(WrenVM*);
extern void f63(WrenVM*);
extern void f64(WrenVM*);
extern void f65(WrenVM*);
extern void f66(WrenVM*);
extern void f67(WrenVM*);
extern void f68(WrenVM*);
extern void f69(WrenVM*);
extern void f70(WrenVM*);
extern void f71(WrenVM*);
extern void f72(WrenVM*);
extern void f73(WrenVM*);
extern void f74(WrenVM*);
extern void f75(WrenVM*);
extern void f76(WrenVM*);
extern void f77(WrenVM*);
extern void f78(WrenVM*);
extern void f79(WrenVM*);
extern void f80(WrenVM*);
extern void f81(WrenVM*);
extern void f82(WrenVM*);
extern void f83(WrenVM*);
extern void f84(WrenVM*);
extern void f85(WrenVM*);
extern void f86(WrenVM*);
extern void f87(WrenVM*);
extern void f88(WrenVM*);
extern void f89(WrenVM*);
extern void f90(WrenVM*);
extern void f91(WrenVM*);
extern void f92(WrenVM*);
extern void f93(WrenVM*);
extern void f94(WrenVM*);
extern void f95(WrenVM*);
extern void f96(WrenVM*);
extern void f97(WrenVM*);
extern void f98(WrenVM*);
extern void f99(WrenVM*);
extern void f100(WrenVM*);
extern void f101(WrenVM*);
extern void f102(WrenVM*);
extern void f103(WrenVM*);
extern void f104(WrenVM*);
extern void f105(WrenVM*);
extern void f106(WrenVM*);
extern void f107(WrenVM*);
extern void f108(WrenVM*);
extern void f109(WrenVM*);
extern void f110(WrenVM*);
extern void f111(WrenVM*);
extern void f112(WrenVM*);
extern void f113(WrenVM*);
extern void f114(WrenVM*);
extern void f115(WrenVM*);
extern void f116(WrenVM*);
extern void f117(WrenVM*);
extern void f118(WrenVM*);
extern void f119(WrenVM*);
extern void f120(WrenVM*);
extern void f121(WrenVM*);
extern void f122(WrenVM*);
extern void f123(WrenVM*);
extern void f124(WrenVM*);
extern void f125(WrenVM*);
extern void f126(WrenVM*);
extern void f127(WrenVM*);
extern void f128(WrenVM*);
extern void f129(WrenVM*);
extern void f130(WrenVM*);
extern void f131(WrenVM*);
extern void f132(WrenVM*);
extern void f133(WrenVM*);
extern void f134(WrenVM*);
extern void f135(WrenVM*);
extern void f136(WrenVM*);
extern void f137(WrenVM*);
extern void f138(WrenVM*);
extern void f139(WrenVM*);
extern void f140(WrenVM*);
extern void f141(WrenVM*);
extern void f142(WrenVM*);
extern void f143(WrenVM*);
extern void f144(WrenVM*);
extern void f145(WrenVM*);
extern void f146(WrenVM*);
extern void f147(WrenVM*);
extern void f148(WrenVM*);
extern void f149(WrenVM*);
extern void f150(WrenVM*);
extern void f151(WrenVM*);
extern void f152(WrenVM*);
extern void f153(WrenVM*);
extern void f154(WrenVM*);
extern void f155(WrenVM*);
extern void f156(WrenVM*);
extern void f157(WrenVM*);
extern void f158(WrenVM*);
extern void f159(WrenVM*);
extern void f160(WrenVM*);
extern void f161(WrenVM*);
extern void f162(WrenVM*);
extern void f163(WrenVM*);
extern void f164(WrenVM*);
extern void f165(WrenVM*);
extern void f166(WrenVM*);
extern void f167(WrenVM*);
extern void f168(WrenVM*);
extern void f169(WrenVM*);
extern void f170(WrenVM*);
extern void f171(WrenVM*);
extern void f172(WrenVM*);
extern void f173(WrenVM*);
extern void f174(WrenVM*);
extern void f175(WrenVM*);
extern void f176(WrenVM*);
extern void f177(WrenVM*);
extern void f178(WrenVM*);
extern void f179(WrenVM*);
extern void f180(WrenVM*);
extern void f181(WrenVM*);
extern void f182(WrenVM*);
extern void f183(WrenVM*);
extern void f184(WrenVM*);
extern void f185(WrenVM*);
extern void f186(WrenVM*);
extern void f187(WrenVM*);
extern void f188(WrenVM*);
extern void f189(WrenVM*);
extern void f190(WrenVM*);
extern void f191(WrenVM*);
extern void f192(WrenVM*);
extern void f193(WrenVM*);
extern void f194(WrenVM*);
extern void f195(WrenVM*);
extern void f196(WrenVM*);
extern void f197(WrenVM*);
extern void f198(WrenVM*);
extern void f199(WrenVM*);
extern void f200(WrenVM*);
extern void f201(WrenVM*);
extern void f202(WrenVM*);
extern void f203(WrenVM*);
extern void f204(WrenVM*);
extern void f205(WrenVM*);
extern void f206(WrenVM*);
extern void f207(WrenVM*);
extern void f208(WrenVM*);
extern void f209(WrenVM*);
extern void f210(WrenVM*);
extern void f211(WrenVM*);
extern void f212(WrenVM*);
extern void f213(WrenVM*);
extern void f214(WrenVM*);
extern void f215(WrenVM*);
extern void f216(WrenVM*);
extern void f217(WrenVM*);
extern void f218(WrenVM*);
extern void f219(WrenVM*);
extern void f220(WrenVM*);
extern void f221(WrenVM*);
extern void f222(WrenVM*);
extern void f223(WrenVM*);
extern void f224(WrenVM*);
extern void f225(WrenVM*);
extern void f226(WrenVM*);
extern void f227(WrenVM*);
extern void f228(WrenVM*);
extern void f229(WrenVM*);
extern void f230(WrenVM*);
extern void f231(WrenVM*);
extern void f232(WrenVM*);
extern void f233(WrenVM*);
extern void f234(WrenVM*);
extern void f235(WrenVM*);
extern void f236(WrenVM*);
extern void f237(WrenVM*);
extern void f238(WrenVM*);
extern void f239(WrenVM*);
extern void f240(WrenVM*);
extern void f241(WrenVM*);
extern void f242(WrenVM*);
extern void f243(WrenVM*);
extern void f244(WrenVM*);
extern void f245(WrenVM*);
extern void f246(WrenVM*);
extern void f247(WrenVM*);
extern void f248(WrenVM*);
extern void f249(WrenVM*);
extern void f250(WrenVM*);
extern void f251(WrenVM*);
extern void f252(WrenVM*);
extern void f253(WrenVM*);
extern void f254(WrenVM*);
extern void f255(WrenVM*);

static inline void* get_f(int i) {
	switch (i) {
		case 0: return f0;
		case 1: return f1;
		case 2: return f2;
		case 3: return f3;
		case 4: return f4;
		case 5: return f5;
		case 6: return f6;
		case 7: return f7;
		case 8: return f8;
		case 9: return f9;
		case 10: return f10;
		case 11: return f11;
		case 12: return f12;
		case 13: return f13;
		case 14: return f14;
		case 15: return f15;
		case 16: return f16;
		case 17: return f17;
		case 18: return f18;
		case 19: return f19;
		case 20: return f20;
		case 21: return f21;
		case 22: return f22;
		case 23: return f23;
		case 24: return f24;
		case 25: return f25;
		case 26: return f26;
		case 27: return f27;
		case 28: return f28;
		case 29: return f29;
		case 30: return f30;
		case 31: return f31;
		case 32: return f32;
		case 33: return f33;
		case 34: return f34;
		case 35: return f35;
		case 36: return f36;
		case 37: return f37;
		case 38: return f38;
		case 39: return f39;
		case 40: return f40;
		case 41: return f41;
		case 42: return f42;
		case 43: return f43;
		case 44: return f44;
		case 45: return f45;
		case 46: return f46;
		case 47: return f47;
		case 48: return f48;
		case 49: return f49;
		case 50: return f50;
		case 51: return f51;
		case 52: return f52;
		case 53: return f53;
		case 54: return f54;
		case 55: return f55;
		case 56: return f56;
		case 57: return f57;
		case 58: return f58;
		case 59: return f59;
		case 60: return f60;
		case 61: return f61;
		case 62: return f62;
		case 63: return f63;
		case 64: return f64;
		case 65: return f65;
		case 66: return f66;
		case 67: return f67;
		case 68: return f68;
		case 69: return f69;
		case 70: return f70;
		case 71: return f71;
		case 72: return f72;
		case 73: return f73;
		case 74: return f74;
		case 75: return f75;
		case 76: return f76;
		case 77: return f77;
		case 78: return f78;
		case 79: return f79;
		case 80: return f80;
		case 81: return f81;
		case 82: return f82;
		case 83: return f83;
		case 84: return f84;
		case 85: return f85;
		case 86: return f86;
		case 87: return f87;
		case 88: return f88;
		case 89: return f89;
		case 90: return f90;
		case 91: return f91;
		case 92: return f92;
		case 93: return f93;
		case 94: return f94;
		case 95: return f95;
		case 96: return f96;
		case 97: return f97;
		case 98: return f98;
		case 99: return f99;
		case 100: return f100;
		case 101: return f101;
		case 102: return f102;
		case 103: return f103;
		case 104: return f104;
		case 105: return f105;
		case 106: return f106;
		case 107: return f107;
		case 108: return f108;
		case 109: return f109;
		case 110: return f110;
		case 111: return f111;
		case 112: return f112;
		case 113: return f113;
		case 114: return f114;
		case 115: return f115;
		case 116: return f116;
		case 117: return f117;
		case 118: return f118;
		case 119: return f119;
		case 120: return f120;
		case 121: return f121;
		case 122: return f122;
		case 123: return f123;
		case 124: return f124;
		case 125: return f125;
		case 126: return f126;
		case 127: return f127;
		case 128: return f128;
		case 129: return f129;
		case 130: return f130;
		case 131: return f131;
		case 132: return f132;
		case 133: return f133;
		case 134: return f134;
		case 135: return f135;
		case 136: return f136;
		case 137: return f137;
		case 138: return f138;
		case 139: return f139;
		case 140: return f140;
		case 141: return f141;
		case 142: return f142;
		case 143: return f143;
		case 144: return f144;
		case 145: return f145;
		case 146: return f146;
		case 147: return f147;
		case 148: return f148;
		case 149: return f149;
		case 150: return f150;
		case 151: return f151;
		case 152: return f152;
		case 153: return f153;
		case 154: return f154;
		case 155: return f155;
		case 156: return f156;
		case 157: return f157;
		case 158: return f158;
		case 159: return f159;
		case 160: return f160;
		case 161: return f161;
		case 162: return f162;
		case 163: return f163;
		case 164: return f164;
		case 165: return f165;
		case 166: return f166;
		case 167: return f167;
		case 168: return f168;
		case 169: return f169;
		case 170: return f170;
		case 171: return f171;
		case 172: return f172;
		case 173: return f173;
		case 174: return f174;
		case 175: return f175;
		case 176: return f176;
		case 177: return f177;
		case 178: return f178;
		case 179: return f179;
		case 180: return f180;
		case 181: return f181;
		case 182: return f182;
		case 183: return f183;
		case 184: return f184;
		case 185: return f185;
		case 186: return f186;
		case 187: return f187;
		case 188: return f188;
		case 189: return f189;
		case 190: return f190;
		case 191: return f191;
		case 192: return f192;
		case 193: return f193;
		case 194: return f194;
		case 195: return f195;
		case 196: return f196;
		case 197: return f197;
		case 198: return f198;
		case 199: return f199;
		case 200: return f200;
		case 201: return f201;
		case 202: return f202;
		case 203: return f203;
		case 204: return f204;
		case 205: return f205;
		case 206: return f206;
		case 207: return f207;
		case 208: return f208;
		case 209: return f209;
		case 210: return f210;
		case 211: return f211;
		case 212: return f212;
		case 213: return f213;
		case 214: return f214;
		case 215: return f215;
		case 216: return f216;
		case 217: return f217;
		case 218: return f218;
		case 219: return f219;
		case 220: return f220;
		case 221: return f221;
		case 222: return f222;
		case 223: return f223;
		case 224: return f224;
		case 225: return f225;
		case 226: return f226;
		case 227: return f227;
		case 228: return f228;
		case 229: return f229;
		case 230: return f230;
		case 231: return f231;
		case 232: return f232;
		case 233: return f233;
		case 234: return f234;
		case 235: return f235;
		case 236: return f236;
		case 237: return f237;
		case 238: return f238;
		case 239: return f239;
		case 240: return f240;
		case 241: return f241;
		case 242: return f242;
		case 243: return f243;
		case 244: return f244;
		case 245: return f245;
		case 246: return f246;
		case 247: return f247;
		case 248: return f248;
		case 249: return f249;
		case 250: return f250;
		case 251: return f251;
		case 252: return f252;
		case 253: return f253;
		case 254: return f254;
		case 255: return f255;
		default: return (void*)(0);
	}
}
*/
import "C"
import (
	"errors"
	"sync"
	"unsafe"
)

const MAX_FUNC_REGISTRATIONS = 256

var (
	fMap = make(map[int]func(*VM))
	fMapGuard sync.Mutex
	fCounter int
)


//export f0
func f0(vm *C.WrenVM) {
	f := fMap[0]
	if f == nil {
		panic("function 0 not registered")
	}
	f(vmMap[vm])
}

//export f1
func f1(vm *C.WrenVM) {
	f := fMap[1]
	if f == nil {
		panic("function 1 not registered")
	}
	f(vmMap[vm])
}

//export f2
func f2(vm *C.WrenVM) {
	f := fMap[2]
	if f == nil {
		panic("function 2 not registered")
	}
	f(vmMap[vm])
}

//export f3
func f3(vm *C.WrenVM) {
	f := fMap[3]
	if f == nil {
		panic("function 3 not registered")
	}
	f(vmMap[vm])
}

//export f4
func f4(vm *C.WrenVM) {
	f := fMap[4]
	if f == nil {
		panic("function 4 not registered")
	}
	f(vmMap[vm])
}

//export f5
func f5(vm *C.WrenVM) {
	f := fMap[5]
	if f == nil {
		panic("function 5 not registered")
	}
	f(vmMap[vm])
}

//export f6
func f6(vm *C.WrenVM) {
	f := fMap[6]
	if f == nil {
		panic("function 6 not registered")
	}
	f(vmMap[vm])
}

//export f7
func f7(vm *C.WrenVM) {
	f := fMap[7]
	if f == nil {
		panic("function 7 not registered")
	}
	f(vmMap[vm])
}

//export f8
func f8(vm *C.WrenVM) {
	f := fMap[8]
	if f == nil {
		panic("function 8 not registered")
	}
	f(vmMap[vm])
}

//export f9
func f9(vm *C.WrenVM) {
	f := fMap[9]
	if f == nil {
		panic("function 9 not registered")
	}
	f(vmMap[vm])
}

//export f10
func f10(vm *C.WrenVM) {
	f := fMap[10]
	if f == nil {
		panic("function 10 not registered")
	}
	f(vmMap[vm])
}

//export f11
func f11(vm *C.WrenVM) {
	f := fMap[11]
	if f == nil {
		panic("function 11 not registered")
	}
	f(vmMap[vm])
}

//export f12
func f12(vm *C.WrenVM) {
	f := fMap[12]
	if f == nil {
		panic("function 12 not registered")
	}
	f(vmMap[vm])
}

//export f13
func f13(vm *C.WrenVM) {
	f := fMap[13]
	if f == nil {
		panic("function 13 not registered")
	}
	f(vmMap[vm])
}

//export f14
func f14(vm *C.WrenVM) {
	f := fMap[14]
	if f == nil {
		panic("function 14 not registered")
	}
	f(vmMap[vm])
}

//export f15
func f15(vm *C.WrenVM) {
	f := fMap[15]
	if f == nil {
		panic("function 15 not registered")
	}
	f(vmMap[vm])
}

//export f16
func f16(vm *C.WrenVM) {
	f := fMap[16]
	if f == nil {
		panic("function 16 not registered")
	}
	f(vmMap[vm])
}

//export f17
func f17(vm *C.WrenVM) {
	f := fMap[17]
	if f == nil {
		panic("function 17 not registered")
	}
	f(vmMap[vm])
}

//export f18
func f18(vm *C.WrenVM) {
	f := fMap[18]
	if f == nil {
		panic("function 18 not registered")
	}
	f(vmMap[vm])
}

//export f19
func f19(vm *C.WrenVM) {
	f := fMap[19]
	if f == nil {
		panic("function 19 not registered")
	}
	f(vmMap[vm])
}

//export f20
func f20(vm *C.WrenVM) {
	f := fMap[20]
	if f == nil {
		panic("function 20 not registered")
	}
	f(vmMap[vm])
}

//export f21
func f21(vm *C.WrenVM) {
	f := fMap[21]
	if f == nil {
		panic("function 21 not registered")
	}
	f(vmMap[vm])
}

//export f22
func f22(vm *C.WrenVM) {
	f := fMap[22]
	if f == nil {
		panic("function 22 not registered")
	}
	f(vmMap[vm])
}

//export f23
func f23(vm *C.WrenVM) {
	f := fMap[23]
	if f == nil {
		panic("function 23 not registered")
	}
	f(vmMap[vm])
}

//export f24
func f24(vm *C.WrenVM) {
	f := fMap[24]
	if f == nil {
		panic("function 24 not registered")
	}
	f(vmMap[vm])
}

//export f25
func f25(vm *C.WrenVM) {
	f := fMap[25]
	if f == nil {
		panic("function 25 not registered")
	}
	f(vmMap[vm])
}

//export f26
func f26(vm *C.WrenVM) {
	f := fMap[26]
	if f == nil {
		panic("function 26 not registered")
	}
	f(vmMap[vm])
}

//export f27
func f27(vm *C.WrenVM) {
	f := fMap[27]
	if f == nil {
		panic("function 27 not registered")
	}
	f(vmMap[vm])
}

//export f28
func f28(vm *C.WrenVM) {
	f := fMap[28]
	if f == nil {
		panic("function 28 not registered")
	}
	f(vmMap[vm])
}

//export f29
func f29(vm *C.WrenVM) {
	f := fMap[29]
	if f == nil {
		panic("function 29 not registered")
	}
	f(vmMap[vm])
}

//export f30
func f30(vm *C.WrenVM) {
	f := fMap[30]
	if f == nil {
		panic("function 30 not registered")
	}
	f(vmMap[vm])
}

//export f31
func f31(vm *C.WrenVM) {
	f := fMap[31]
	if f == nil {
		panic("function 31 not registered")
	}
	f(vmMap[vm])
}

//export f32
func f32(vm *C.WrenVM) {
	f := fMap[32]
	if f == nil {
		panic("function 32 not registered")
	}
	f(vmMap[vm])
}

//export f33
func f33(vm *C.WrenVM) {
	f := fMap[33]
	if f == nil {
		panic("function 33 not registered")
	}
	f(vmMap[vm])
}

//export f34
func f34(vm *C.WrenVM) {
	f := fMap[34]
	if f == nil {
		panic("function 34 not registered")
	}
	f(vmMap[vm])
}

//export f35
func f35(vm *C.WrenVM) {
	f := fMap[35]
	if f == nil {
		panic("function 35 not registered")
	}
	f(vmMap[vm])
}

//export f36
func f36(vm *C.WrenVM) {
	f := fMap[36]
	if f == nil {
		panic("function 36 not registered")
	}
	f(vmMap[vm])
}

//export f37
func f37(vm *C.WrenVM) {
	f := fMap[37]
	if f == nil {
		panic("function 37 not registered")
	}
	f(vmMap[vm])
}

//export f38
func f38(vm *C.WrenVM) {
	f := fMap[38]
	if f == nil {
		panic("function 38 not registered")
	}
	f(vmMap[vm])
}

//export f39
func f39(vm *C.WrenVM) {
	f := fMap[39]
	if f == nil {
		panic("function 39 not registered")
	}
	f(vmMap[vm])
}

//export f40
func f40(vm *C.WrenVM) {
	f := fMap[40]
	if f == nil {
		panic("function 40 not registered")
	}
	f(vmMap[vm])
}

//export f41
func f41(vm *C.WrenVM) {
	f := fMap[41]
	if f == nil {
		panic("function 41 not registered")
	}
	f(vmMap[vm])
}

//export f42
func f42(vm *C.WrenVM) {
	f := fMap[42]
	if f == nil {
		panic("function 42 not registered")
	}
	f(vmMap[vm])
}

//export f43
func f43(vm *C.WrenVM) {
	f := fMap[43]
	if f == nil {
		panic("function 43 not registered")
	}
	f(vmMap[vm])
}

//export f44
func f44(vm *C.WrenVM) {
	f := fMap[44]
	if f == nil {
		panic("function 44 not registered")
	}
	f(vmMap[vm])
}

//export f45
func f45(vm *C.WrenVM) {
	f := fMap[45]
	if f == nil {
		panic("function 45 not registered")
	}
	f(vmMap[vm])
}

//export f46
func f46(vm *C.WrenVM) {
	f := fMap[46]
	if f == nil {
		panic("function 46 not registered")
	}
	f(vmMap[vm])
}

//export f47
func f47(vm *C.WrenVM) {
	f := fMap[47]
	if f == nil {
		panic("function 47 not registered")
	}
	f(vmMap[vm])
}

//export f48
func f48(vm *C.WrenVM) {
	f := fMap[48]
	if f == nil {
		panic("function 48 not registered")
	}
	f(vmMap[vm])
}

//export f49
func f49(vm *C.WrenVM) {
	f := fMap[49]
	if f == nil {
		panic("function 49 not registered")
	}
	f(vmMap[vm])
}

//export f50
func f50(vm *C.WrenVM) {
	f := fMap[50]
	if f == nil {
		panic("function 50 not registered")
	}
	f(vmMap[vm])
}

//export f51
func f51(vm *C.WrenVM) {
	f := fMap[51]
	if f == nil {
		panic("function 51 not registered")
	}
	f(vmMap[vm])
}

//export f52
func f52(vm *C.WrenVM) {
	f := fMap[52]
	if f == nil {
		panic("function 52 not registered")
	}
	f(vmMap[vm])
}

//export f53
func f53(vm *C.WrenVM) {
	f := fMap[53]
	if f == nil {
		panic("function 53 not registered")
	}
	f(vmMap[vm])
}

//export f54
func f54(vm *C.WrenVM) {
	f := fMap[54]
	if f == nil {
		panic("function 54 not registered")
	}
	f(vmMap[vm])
}

//export f55
func f55(vm *C.WrenVM) {
	f := fMap[55]
	if f == nil {
		panic("function 55 not registered")
	}
	f(vmMap[vm])
}

//export f56
func f56(vm *C.WrenVM) {
	f := fMap[56]
	if f == nil {
		panic("function 56 not registered")
	}
	f(vmMap[vm])
}

//export f57
func f57(vm *C.WrenVM) {
	f := fMap[57]
	if f == nil {
		panic("function 57 not registered")
	}
	f(vmMap[vm])
}

//export f58
func f58(vm *C.WrenVM) {
	f := fMap[58]
	if f == nil {
		panic("function 58 not registered")
	}
	f(vmMap[vm])
}

//export f59
func f59(vm *C.WrenVM) {
	f := fMap[59]
	if f == nil {
		panic("function 59 not registered")
	}
	f(vmMap[vm])
}

//export f60
func f60(vm *C.WrenVM) {
	f := fMap[60]
	if f == nil {
		panic("function 60 not registered")
	}
	f(vmMap[vm])
}

//export f61
func f61(vm *C.WrenVM) {
	f := fMap[61]
	if f == nil {
		panic("function 61 not registered")
	}
	f(vmMap[vm])
}

//export f62
func f62(vm *C.WrenVM) {
	f := fMap[62]
	if f == nil {
		panic("function 62 not registered")
	}
	f(vmMap[vm])
}

//export f63
func f63(vm *C.WrenVM) {
	f := fMap[63]
	if f == nil {
		panic("function 63 not registered")
	}
	f(vmMap[vm])
}

//export f64
func f64(vm *C.WrenVM) {
	f := fMap[64]
	if f == nil {
		panic("function 64 not registered")
	}
	f(vmMap[vm])
}

//export f65
func f65(vm *C.WrenVM) {
	f := fMap[65]
	if f == nil {
		panic("function 65 not registered")
	}
	f(vmMap[vm])
}

//export f66
func f66(vm *C.WrenVM) {
	f := fMap[66]
	if f == nil {
		panic("function 66 not registered")
	}
	f(vmMap[vm])
}

//export f67
func f67(vm *C.WrenVM) {
	f := fMap[67]
	if f == nil {
		panic("function 67 not registered")
	}
	f(vmMap[vm])
}

//export f68
func f68(vm *C.WrenVM) {
	f := fMap[68]
	if f == nil {
		panic("function 68 not registered")
	}
	f(vmMap[vm])
}

//export f69
func f69(vm *C.WrenVM) {
	f := fMap[69]
	if f == nil {
		panic("function 69 not registered")
	}
	f(vmMap[vm])
}

//export f70
func f70(vm *C.WrenVM) {
	f := fMap[70]
	if f == nil {
		panic("function 70 not registered")
	}
	f(vmMap[vm])
}

//export f71
func f71(vm *C.WrenVM) {
	f := fMap[71]
	if f == nil {
		panic("function 71 not registered")
	}
	f(vmMap[vm])
}

//export f72
func f72(vm *C.WrenVM) {
	f := fMap[72]
	if f == nil {
		panic("function 72 not registered")
	}
	f(vmMap[vm])
}

//export f73
func f73(vm *C.WrenVM) {
	f := fMap[73]
	if f == nil {
		panic("function 73 not registered")
	}
	f(vmMap[vm])
}

//export f74
func f74(vm *C.WrenVM) {
	f := fMap[74]
	if f == nil {
		panic("function 74 not registered")
	}
	f(vmMap[vm])
}

//export f75
func f75(vm *C.WrenVM) {
	f := fMap[75]
	if f == nil {
		panic("function 75 not registered")
	}
	f(vmMap[vm])
}

//export f76
func f76(vm *C.WrenVM) {
	f := fMap[76]
	if f == nil {
		panic("function 76 not registered")
	}
	f(vmMap[vm])
}

//export f77
func f77(vm *C.WrenVM) {
	f := fMap[77]
	if f == nil {
		panic("function 77 not registered")
	}
	f(vmMap[vm])
}

//export f78
func f78(vm *C.WrenVM) {
	f := fMap[78]
	if f == nil {
		panic("function 78 not registered")
	}
	f(vmMap[vm])
}

//export f79
func f79(vm *C.WrenVM) {
	f := fMap[79]
	if f == nil {
		panic("function 79 not registered")
	}
	f(vmMap[vm])
}

//export f80
func f80(vm *C.WrenVM) {
	f := fMap[80]
	if f == nil {
		panic("function 80 not registered")
	}
	f(vmMap[vm])
}

//export f81
func f81(vm *C.WrenVM) {
	f := fMap[81]
	if f == nil {
		panic("function 81 not registered")
	}
	f(vmMap[vm])
}

//export f82
func f82(vm *C.WrenVM) {
	f := fMap[82]
	if f == nil {
		panic("function 82 not registered")
	}
	f(vmMap[vm])
}

//export f83
func f83(vm *C.WrenVM) {
	f := fMap[83]
	if f == nil {
		panic("function 83 not registered")
	}
	f(vmMap[vm])
}

//export f84
func f84(vm *C.WrenVM) {
	f := fMap[84]
	if f == nil {
		panic("function 84 not registered")
	}
	f(vmMap[vm])
}

//export f85
func f85(vm *C.WrenVM) {
	f := fMap[85]
	if f == nil {
		panic("function 85 not registered")
	}
	f(vmMap[vm])
}

//export f86
func f86(vm *C.WrenVM) {
	f := fMap[86]
	if f == nil {
		panic("function 86 not registered")
	}
	f(vmMap[vm])
}

//export f87
func f87(vm *C.WrenVM) {
	f := fMap[87]
	if f == nil {
		panic("function 87 not registered")
	}
	f(vmMap[vm])
}

//export f88
func f88(vm *C.WrenVM) {
	f := fMap[88]
	if f == nil {
		panic("function 88 not registered")
	}
	f(vmMap[vm])
}

//export f89
func f89(vm *C.WrenVM) {
	f := fMap[89]
	if f == nil {
		panic("function 89 not registered")
	}
	f(vmMap[vm])
}

//export f90
func f90(vm *C.WrenVM) {
	f := fMap[90]
	if f == nil {
		panic("function 90 not registered")
	}
	f(vmMap[vm])
}

//export f91
func f91(vm *C.WrenVM) {
	f := fMap[91]
	if f == nil {
		panic("function 91 not registered")
	}
	f(vmMap[vm])
}

//export f92
func f92(vm *C.WrenVM) {
	f := fMap[92]
	if f == nil {
		panic("function 92 not registered")
	}
	f(vmMap[vm])
}

//export f93
func f93(vm *C.WrenVM) {
	f := fMap[93]
	if f == nil {
		panic("function 93 not registered")
	}
	f(vmMap[vm])
}

//export f94
func f94(vm *C.WrenVM) {
	f := fMap[94]
	if f == nil {
		panic("function 94 not registered")
	}
	f(vmMap[vm])
}

//export f95
func f95(vm *C.WrenVM) {
	f := fMap[95]
	if f == nil {
		panic("function 95 not registered")
	}
	f(vmMap[vm])
}

//export f96
func f96(vm *C.WrenVM) {
	f := fMap[96]
	if f == nil {
		panic("function 96 not registered")
	}
	f(vmMap[vm])
}

//export f97
func f97(vm *C.WrenVM) {
	f := fMap[97]
	if f == nil {
		panic("function 97 not registered")
	}
	f(vmMap[vm])
}

//export f98
func f98(vm *C.WrenVM) {
	f := fMap[98]
	if f == nil {
		panic("function 98 not registered")
	}
	f(vmMap[vm])
}

//export f99
func f99(vm *C.WrenVM) {
	f := fMap[99]
	if f == nil {
		panic("function 99 not registered")
	}
	f(vmMap[vm])
}

//export f100
func f100(vm *C.WrenVM) {
	f := fMap[100]
	if f == nil {
		panic("function 100 not registered")
	}
	f(vmMap[vm])
}

//export f101
func f101(vm *C.WrenVM) {
	f := fMap[101]
	if f == nil {
		panic("function 101 not registered")
	}
	f(vmMap[vm])
}

//export f102
func f102(vm *C.WrenVM) {
	f := fMap[102]
	if f == nil {
		panic("function 102 not registered")
	}
	f(vmMap[vm])
}

//export f103
func f103(vm *C.WrenVM) {
	f := fMap[103]
	if f == nil {
		panic("function 103 not registered")
	}
	f(vmMap[vm])
}

//export f104
func f104(vm *C.WrenVM) {
	f := fMap[104]
	if f == nil {
		panic("function 104 not registered")
	}
	f(vmMap[vm])
}

//export f105
func f105(vm *C.WrenVM) {
	f := fMap[105]
	if f == nil {
		panic("function 105 not registered")
	}
	f(vmMap[vm])
}

//export f106
func f106(vm *C.WrenVM) {
	f := fMap[106]
	if f == nil {
		panic("function 106 not registered")
	}
	f(vmMap[vm])
}

//export f107
func f107(vm *C.WrenVM) {
	f := fMap[107]
	if f == nil {
		panic("function 107 not registered")
	}
	f(vmMap[vm])
}

//export f108
func f108(vm *C.WrenVM) {
	f := fMap[108]
	if f == nil {
		panic("function 108 not registered")
	}
	f(vmMap[vm])
}

//export f109
func f109(vm *C.WrenVM) {
	f := fMap[109]
	if f == nil {
		panic("function 109 not registered")
	}
	f(vmMap[vm])
}

//export f110
func f110(vm *C.WrenVM) {
	f := fMap[110]
	if f == nil {
		panic("function 110 not registered")
	}
	f(vmMap[vm])
}

//export f111
func f111(vm *C.WrenVM) {
	f := fMap[111]
	if f == nil {
		panic("function 111 not registered")
	}
	f(vmMap[vm])
}

//export f112
func f112(vm *C.WrenVM) {
	f := fMap[112]
	if f == nil {
		panic("function 112 not registered")
	}
	f(vmMap[vm])
}

//export f113
func f113(vm *C.WrenVM) {
	f := fMap[113]
	if f == nil {
		panic("function 113 not registered")
	}
	f(vmMap[vm])
}

//export f114
func f114(vm *C.WrenVM) {
	f := fMap[114]
	if f == nil {
		panic("function 114 not registered")
	}
	f(vmMap[vm])
}

//export f115
func f115(vm *C.WrenVM) {
	f := fMap[115]
	if f == nil {
		panic("function 115 not registered")
	}
	f(vmMap[vm])
}

//export f116
func f116(vm *C.WrenVM) {
	f := fMap[116]
	if f == nil {
		panic("function 116 not registered")
	}
	f(vmMap[vm])
}

//export f117
func f117(vm *C.WrenVM) {
	f := fMap[117]
	if f == nil {
		panic("function 117 not registered")
	}
	f(vmMap[vm])
}

//export f118
func f118(vm *C.WrenVM) {
	f := fMap[118]
	if f == nil {
		panic("function 118 not registered")
	}
	f(vmMap[vm])
}

//export f119
func f119(vm *C.WrenVM) {
	f := fMap[119]
	if f == nil {
		panic("function 119 not registered")
	}
	f(vmMap[vm])
}

//export f120
func f120(vm *C.WrenVM) {
	f := fMap[120]
	if f == nil {
		panic("function 120 not registered")
	}
	f(vmMap[vm])
}

//export f121
func f121(vm *C.WrenVM) {
	f := fMap[121]
	if f == nil {
		panic("function 121 not registered")
	}
	f(vmMap[vm])
}

//export f122
func f122(vm *C.WrenVM) {
	f := fMap[122]
	if f == nil {
		panic("function 122 not registered")
	}
	f(vmMap[vm])
}

//export f123
func f123(vm *C.WrenVM) {
	f := fMap[123]
	if f == nil {
		panic("function 123 not registered")
	}
	f(vmMap[vm])
}

//export f124
func f124(vm *C.WrenVM) {
	f := fMap[124]
	if f == nil {
		panic("function 124 not registered")
	}
	f(vmMap[vm])
}

//export f125
func f125(vm *C.WrenVM) {
	f := fMap[125]
	if f == nil {
		panic("function 125 not registered")
	}
	f(vmMap[vm])
}

//export f126
func f126(vm *C.WrenVM) {
	f := fMap[126]
	if f == nil {
		panic("function 126 not registered")
	}
	f(vmMap[vm])
}

//export f127
func f127(vm *C.WrenVM) {
	f := fMap[127]
	if f == nil {
		panic("function 127 not registered")
	}
	f(vmMap[vm])
}

//export f128
func f128(vm *C.WrenVM) {
	f := fMap[128]
	if f == nil {
		panic("function 128 not registered")
	}
	f(vmMap[vm])
}

//export f129
func f129(vm *C.WrenVM) {
	f := fMap[129]
	if f == nil {
		panic("function 129 not registered")
	}
	f(vmMap[vm])
}

//export f130
func f130(vm *C.WrenVM) {
	f := fMap[130]
	if f == nil {
		panic("function 130 not registered")
	}
	f(vmMap[vm])
}

//export f131
func f131(vm *C.WrenVM) {
	f := fMap[131]
	if f == nil {
		panic("function 131 not registered")
	}
	f(vmMap[vm])
}

//export f132
func f132(vm *C.WrenVM) {
	f := fMap[132]
	if f == nil {
		panic("function 132 not registered")
	}
	f(vmMap[vm])
}

//export f133
func f133(vm *C.WrenVM) {
	f := fMap[133]
	if f == nil {
		panic("function 133 not registered")
	}
	f(vmMap[vm])
}

//export f134
func f134(vm *C.WrenVM) {
	f := fMap[134]
	if f == nil {
		panic("function 134 not registered")
	}
	f(vmMap[vm])
}

//export f135
func f135(vm *C.WrenVM) {
	f := fMap[135]
	if f == nil {
		panic("function 135 not registered")
	}
	f(vmMap[vm])
}

//export f136
func f136(vm *C.WrenVM) {
	f := fMap[136]
	if f == nil {
		panic("function 136 not registered")
	}
	f(vmMap[vm])
}

//export f137
func f137(vm *C.WrenVM) {
	f := fMap[137]
	if f == nil {
		panic("function 137 not registered")
	}
	f(vmMap[vm])
}

//export f138
func f138(vm *C.WrenVM) {
	f := fMap[138]
	if f == nil {
		panic("function 138 not registered")
	}
	f(vmMap[vm])
}

//export f139
func f139(vm *C.WrenVM) {
	f := fMap[139]
	if f == nil {
		panic("function 139 not registered")
	}
	f(vmMap[vm])
}

//export f140
func f140(vm *C.WrenVM) {
	f := fMap[140]
	if f == nil {
		panic("function 140 not registered")
	}
	f(vmMap[vm])
}

//export f141
func f141(vm *C.WrenVM) {
	f := fMap[141]
	if f == nil {
		panic("function 141 not registered")
	}
	f(vmMap[vm])
}

//export f142
func f142(vm *C.WrenVM) {
	f := fMap[142]
	if f == nil {
		panic("function 142 not registered")
	}
	f(vmMap[vm])
}

//export f143
func f143(vm *C.WrenVM) {
	f := fMap[143]
	if f == nil {
		panic("function 143 not registered")
	}
	f(vmMap[vm])
}

//export f144
func f144(vm *C.WrenVM) {
	f := fMap[144]
	if f == nil {
		panic("function 144 not registered")
	}
	f(vmMap[vm])
}

//export f145
func f145(vm *C.WrenVM) {
	f := fMap[145]
	if f == nil {
		panic("function 145 not registered")
	}
	f(vmMap[vm])
}

//export f146
func f146(vm *C.WrenVM) {
	f := fMap[146]
	if f == nil {
		panic("function 146 not registered")
	}
	f(vmMap[vm])
}

//export f147
func f147(vm *C.WrenVM) {
	f := fMap[147]
	if f == nil {
		panic("function 147 not registered")
	}
	f(vmMap[vm])
}

//export f148
func f148(vm *C.WrenVM) {
	f := fMap[148]
	if f == nil {
		panic("function 148 not registered")
	}
	f(vmMap[vm])
}

//export f149
func f149(vm *C.WrenVM) {
	f := fMap[149]
	if f == nil {
		panic("function 149 not registered")
	}
	f(vmMap[vm])
}

//export f150
func f150(vm *C.WrenVM) {
	f := fMap[150]
	if f == nil {
		panic("function 150 not registered")
	}
	f(vmMap[vm])
}

//export f151
func f151(vm *C.WrenVM) {
	f := fMap[151]
	if f == nil {
		panic("function 151 not registered")
	}
	f(vmMap[vm])
}

//export f152
func f152(vm *C.WrenVM) {
	f := fMap[152]
	if f == nil {
		panic("function 152 not registered")
	}
	f(vmMap[vm])
}

//export f153
func f153(vm *C.WrenVM) {
	f := fMap[153]
	if f == nil {
		panic("function 153 not registered")
	}
	f(vmMap[vm])
}

//export f154
func f154(vm *C.WrenVM) {
	f := fMap[154]
	if f == nil {
		panic("function 154 not registered")
	}
	f(vmMap[vm])
}

//export f155
func f155(vm *C.WrenVM) {
	f := fMap[155]
	if f == nil {
		panic("function 155 not registered")
	}
	f(vmMap[vm])
}

//export f156
func f156(vm *C.WrenVM) {
	f := fMap[156]
	if f == nil {
		panic("function 156 not registered")
	}
	f(vmMap[vm])
}

//export f157
func f157(vm *C.WrenVM) {
	f := fMap[157]
	if f == nil {
		panic("function 157 not registered")
	}
	f(vmMap[vm])
}

//export f158
func f158(vm *C.WrenVM) {
	f := fMap[158]
	if f == nil {
		panic("function 158 not registered")
	}
	f(vmMap[vm])
}

//export f159
func f159(vm *C.WrenVM) {
	f := fMap[159]
	if f == nil {
		panic("function 159 not registered")
	}
	f(vmMap[vm])
}

//export f160
func f160(vm *C.WrenVM) {
	f := fMap[160]
	if f == nil {
		panic("function 160 not registered")
	}
	f(vmMap[vm])
}

//export f161
func f161(vm *C.WrenVM) {
	f := fMap[161]
	if f == nil {
		panic("function 161 not registered")
	}
	f(vmMap[vm])
}

//export f162
func f162(vm *C.WrenVM) {
	f := fMap[162]
	if f == nil {
		panic("function 162 not registered")
	}
	f(vmMap[vm])
}

//export f163
func f163(vm *C.WrenVM) {
	f := fMap[163]
	if f == nil {
		panic("function 163 not registered")
	}
	f(vmMap[vm])
}

//export f164
func f164(vm *C.WrenVM) {
	f := fMap[164]
	if f == nil {
		panic("function 164 not registered")
	}
	f(vmMap[vm])
}

//export f165
func f165(vm *C.WrenVM) {
	f := fMap[165]
	if f == nil {
		panic("function 165 not registered")
	}
	f(vmMap[vm])
}

//export f166
func f166(vm *C.WrenVM) {
	f := fMap[166]
	if f == nil {
		panic("function 166 not registered")
	}
	f(vmMap[vm])
}

//export f167
func f167(vm *C.WrenVM) {
	f := fMap[167]
	if f == nil {
		panic("function 167 not registered")
	}
	f(vmMap[vm])
}

//export f168
func f168(vm *C.WrenVM) {
	f := fMap[168]
	if f == nil {
		panic("function 168 not registered")
	}
	f(vmMap[vm])
}

//export f169
func f169(vm *C.WrenVM) {
	f := fMap[169]
	if f == nil {
		panic("function 169 not registered")
	}
	f(vmMap[vm])
}

//export f170
func f170(vm *C.WrenVM) {
	f := fMap[170]
	if f == nil {
		panic("function 170 not registered")
	}
	f(vmMap[vm])
}

//export f171
func f171(vm *C.WrenVM) {
	f := fMap[171]
	if f == nil {
		panic("function 171 not registered")
	}
	f(vmMap[vm])
}

//export f172
func f172(vm *C.WrenVM) {
	f := fMap[172]
	if f == nil {
		panic("function 172 not registered")
	}
	f(vmMap[vm])
}

//export f173
func f173(vm *C.WrenVM) {
	f := fMap[173]
	if f == nil {
		panic("function 173 not registered")
	}
	f(vmMap[vm])
}

//export f174
func f174(vm *C.WrenVM) {
	f := fMap[174]
	if f == nil {
		panic("function 174 not registered")
	}
	f(vmMap[vm])
}

//export f175
func f175(vm *C.WrenVM) {
	f := fMap[175]
	if f == nil {
		panic("function 175 not registered")
	}
	f(vmMap[vm])
}

//export f176
func f176(vm *C.WrenVM) {
	f := fMap[176]
	if f == nil {
		panic("function 176 not registered")
	}
	f(vmMap[vm])
}

//export f177
func f177(vm *C.WrenVM) {
	f := fMap[177]
	if f == nil {
		panic("function 177 not registered")
	}
	f(vmMap[vm])
}

//export f178
func f178(vm *C.WrenVM) {
	f := fMap[178]
	if f == nil {
		panic("function 178 not registered")
	}
	f(vmMap[vm])
}

//export f179
func f179(vm *C.WrenVM) {
	f := fMap[179]
	if f == nil {
		panic("function 179 not registered")
	}
	f(vmMap[vm])
}

//export f180
func f180(vm *C.WrenVM) {
	f := fMap[180]
	if f == nil {
		panic("function 180 not registered")
	}
	f(vmMap[vm])
}

//export f181
func f181(vm *C.WrenVM) {
	f := fMap[181]
	if f == nil {
		panic("function 181 not registered")
	}
	f(vmMap[vm])
}

//export f182
func f182(vm *C.WrenVM) {
	f := fMap[182]
	if f == nil {
		panic("function 182 not registered")
	}
	f(vmMap[vm])
}

//export f183
func f183(vm *C.WrenVM) {
	f := fMap[183]
	if f == nil {
		panic("function 183 not registered")
	}
	f(vmMap[vm])
}

//export f184
func f184(vm *C.WrenVM) {
	f := fMap[184]
	if f == nil {
		panic("function 184 not registered")
	}
	f(vmMap[vm])
}

//export f185
func f185(vm *C.WrenVM) {
	f := fMap[185]
	if f == nil {
		panic("function 185 not registered")
	}
	f(vmMap[vm])
}

//export f186
func f186(vm *C.WrenVM) {
	f := fMap[186]
	if f == nil {
		panic("function 186 not registered")
	}
	f(vmMap[vm])
}

//export f187
func f187(vm *C.WrenVM) {
	f := fMap[187]
	if f == nil {
		panic("function 187 not registered")
	}
	f(vmMap[vm])
}

//export f188
func f188(vm *C.WrenVM) {
	f := fMap[188]
	if f == nil {
		panic("function 188 not registered")
	}
	f(vmMap[vm])
}

//export f189
func f189(vm *C.WrenVM) {
	f := fMap[189]
	if f == nil {
		panic("function 189 not registered")
	}
	f(vmMap[vm])
}

//export f190
func f190(vm *C.WrenVM) {
	f := fMap[190]
	if f == nil {
		panic("function 190 not registered")
	}
	f(vmMap[vm])
}

//export f191
func f191(vm *C.WrenVM) {
	f := fMap[191]
	if f == nil {
		panic("function 191 not registered")
	}
	f(vmMap[vm])
}

//export f192
func f192(vm *C.WrenVM) {
	f := fMap[192]
	if f == nil {
		panic("function 192 not registered")
	}
	f(vmMap[vm])
}

//export f193
func f193(vm *C.WrenVM) {
	f := fMap[193]
	if f == nil {
		panic("function 193 not registered")
	}
	f(vmMap[vm])
}

//export f194
func f194(vm *C.WrenVM) {
	f := fMap[194]
	if f == nil {
		panic("function 194 not registered")
	}
	f(vmMap[vm])
}

//export f195
func f195(vm *C.WrenVM) {
	f := fMap[195]
	if f == nil {
		panic("function 195 not registered")
	}
	f(vmMap[vm])
}

//export f196
func f196(vm *C.WrenVM) {
	f := fMap[196]
	if f == nil {
		panic("function 196 not registered")
	}
	f(vmMap[vm])
}

//export f197
func f197(vm *C.WrenVM) {
	f := fMap[197]
	if f == nil {
		panic("function 197 not registered")
	}
	f(vmMap[vm])
}

//export f198
func f198(vm *C.WrenVM) {
	f := fMap[198]
	if f == nil {
		panic("function 198 not registered")
	}
	f(vmMap[vm])
}

//export f199
func f199(vm *C.WrenVM) {
	f := fMap[199]
	if f == nil {
		panic("function 199 not registered")
	}
	f(vmMap[vm])
}

//export f200
func f200(vm *C.WrenVM) {
	f := fMap[200]
	if f == nil {
		panic("function 200 not registered")
	}
	f(vmMap[vm])
}

//export f201
func f201(vm *C.WrenVM) {
	f := fMap[201]
	if f == nil {
		panic("function 201 not registered")
	}
	f(vmMap[vm])
}

//export f202
func f202(vm *C.WrenVM) {
	f := fMap[202]
	if f == nil {
		panic("function 202 not registered")
	}
	f(vmMap[vm])
}

//export f203
func f203(vm *C.WrenVM) {
	f := fMap[203]
	if f == nil {
		panic("function 203 not registered")
	}
	f(vmMap[vm])
}

//export f204
func f204(vm *C.WrenVM) {
	f := fMap[204]
	if f == nil {
		panic("function 204 not registered")
	}
	f(vmMap[vm])
}

//export f205
func f205(vm *C.WrenVM) {
	f := fMap[205]
	if f == nil {
		panic("function 205 not registered")
	}
	f(vmMap[vm])
}

//export f206
func f206(vm *C.WrenVM) {
	f := fMap[206]
	if f == nil {
		panic("function 206 not registered")
	}
	f(vmMap[vm])
}

//export f207
func f207(vm *C.WrenVM) {
	f := fMap[207]
	if f == nil {
		panic("function 207 not registered")
	}
	f(vmMap[vm])
}

//export f208
func f208(vm *C.WrenVM) {
	f := fMap[208]
	if f == nil {
		panic("function 208 not registered")
	}
	f(vmMap[vm])
}

//export f209
func f209(vm *C.WrenVM) {
	f := fMap[209]
	if f == nil {
		panic("function 209 not registered")
	}
	f(vmMap[vm])
}

//export f210
func f210(vm *C.WrenVM) {
	f := fMap[210]
	if f == nil {
		panic("function 210 not registered")
	}
	f(vmMap[vm])
}

//export f211
func f211(vm *C.WrenVM) {
	f := fMap[211]
	if f == nil {
		panic("function 211 not registered")
	}
	f(vmMap[vm])
}

//export f212
func f212(vm *C.WrenVM) {
	f := fMap[212]
	if f == nil {
		panic("function 212 not registered")
	}
	f(vmMap[vm])
}

//export f213
func f213(vm *C.WrenVM) {
	f := fMap[213]
	if f == nil {
		panic("function 213 not registered")
	}
	f(vmMap[vm])
}

//export f214
func f214(vm *C.WrenVM) {
	f := fMap[214]
	if f == nil {
		panic("function 214 not registered")
	}
	f(vmMap[vm])
}

//export f215
func f215(vm *C.WrenVM) {
	f := fMap[215]
	if f == nil {
		panic("function 215 not registered")
	}
	f(vmMap[vm])
}

//export f216
func f216(vm *C.WrenVM) {
	f := fMap[216]
	if f == nil {
		panic("function 216 not registered")
	}
	f(vmMap[vm])
}

//export f217
func f217(vm *C.WrenVM) {
	f := fMap[217]
	if f == nil {
		panic("function 217 not registered")
	}
	f(vmMap[vm])
}

//export f218
func f218(vm *C.WrenVM) {
	f := fMap[218]
	if f == nil {
		panic("function 218 not registered")
	}
	f(vmMap[vm])
}

//export f219
func f219(vm *C.WrenVM) {
	f := fMap[219]
	if f == nil {
		panic("function 219 not registered")
	}
	f(vmMap[vm])
}

//export f220
func f220(vm *C.WrenVM) {
	f := fMap[220]
	if f == nil {
		panic("function 220 not registered")
	}
	f(vmMap[vm])
}

//export f221
func f221(vm *C.WrenVM) {
	f := fMap[221]
	if f == nil {
		panic("function 221 not registered")
	}
	f(vmMap[vm])
}

//export f222
func f222(vm *C.WrenVM) {
	f := fMap[222]
	if f == nil {
		panic("function 222 not registered")
	}
	f(vmMap[vm])
}

//export f223
func f223(vm *C.WrenVM) {
	f := fMap[223]
	if f == nil {
		panic("function 223 not registered")
	}
	f(vmMap[vm])
}

//export f224
func f224(vm *C.WrenVM) {
	f := fMap[224]
	if f == nil {
		panic("function 224 not registered")
	}
	f(vmMap[vm])
}

//export f225
func f225(vm *C.WrenVM) {
	f := fMap[225]
	if f == nil {
		panic("function 225 not registered")
	}
	f(vmMap[vm])
}

//export f226
func f226(vm *C.WrenVM) {
	f := fMap[226]
	if f == nil {
		panic("function 226 not registered")
	}
	f(vmMap[vm])
}

//export f227
func f227(vm *C.WrenVM) {
	f := fMap[227]
	if f == nil {
		panic("function 227 not registered")
	}
	f(vmMap[vm])
}

//export f228
func f228(vm *C.WrenVM) {
	f := fMap[228]
	if f == nil {
		panic("function 228 not registered")
	}
	f(vmMap[vm])
}

//export f229
func f229(vm *C.WrenVM) {
	f := fMap[229]
	if f == nil {
		panic("function 229 not registered")
	}
	f(vmMap[vm])
}

//export f230
func f230(vm *C.WrenVM) {
	f := fMap[230]
	if f == nil {
		panic("function 230 not registered")
	}
	f(vmMap[vm])
}

//export f231
func f231(vm *C.WrenVM) {
	f := fMap[231]
	if f == nil {
		panic("function 231 not registered")
	}
	f(vmMap[vm])
}

//export f232
func f232(vm *C.WrenVM) {
	f := fMap[232]
	if f == nil {
		panic("function 232 not registered")
	}
	f(vmMap[vm])
}

//export f233
func f233(vm *C.WrenVM) {
	f := fMap[233]
	if f == nil {
		panic("function 233 not registered")
	}
	f(vmMap[vm])
}

//export f234
func f234(vm *C.WrenVM) {
	f := fMap[234]
	if f == nil {
		panic("function 234 not registered")
	}
	f(vmMap[vm])
}

//export f235
func f235(vm *C.WrenVM) {
	f := fMap[235]
	if f == nil {
		panic("function 235 not registered")
	}
	f(vmMap[vm])
}

//export f236
func f236(vm *C.WrenVM) {
	f := fMap[236]
	if f == nil {
		panic("function 236 not registered")
	}
	f(vmMap[vm])
}

//export f237
func f237(vm *C.WrenVM) {
	f := fMap[237]
	if f == nil {
		panic("function 237 not registered")
	}
	f(vmMap[vm])
}

//export f238
func f238(vm *C.WrenVM) {
	f := fMap[238]
	if f == nil {
		panic("function 238 not registered")
	}
	f(vmMap[vm])
}

//export f239
func f239(vm *C.WrenVM) {
	f := fMap[239]
	if f == nil {
		panic("function 239 not registered")
	}
	f(vmMap[vm])
}

//export f240
func f240(vm *C.WrenVM) {
	f := fMap[240]
	if f == nil {
		panic("function 240 not registered")
	}
	f(vmMap[vm])
}

//export f241
func f241(vm *C.WrenVM) {
	f := fMap[241]
	if f == nil {
		panic("function 241 not registered")
	}
	f(vmMap[vm])
}

//export f242
func f242(vm *C.WrenVM) {
	f := fMap[242]
	if f == nil {
		panic("function 242 not registered")
	}
	f(vmMap[vm])
}

//export f243
func f243(vm *C.WrenVM) {
	f := fMap[243]
	if f == nil {
		panic("function 243 not registered")
	}
	f(vmMap[vm])
}

//export f244
func f244(vm *C.WrenVM) {
	f := fMap[244]
	if f == nil {
		panic("function 244 not registered")
	}
	f(vmMap[vm])
}

//export f245
func f245(vm *C.WrenVM) {
	f := fMap[245]
	if f == nil {
		panic("function 245 not registered")
	}
	f(vmMap[vm])
}

//export f246
func f246(vm *C.WrenVM) {
	f := fMap[246]
	if f == nil {
		panic("function 246 not registered")
	}
	f(vmMap[vm])
}

//export f247
func f247(vm *C.WrenVM) {
	f := fMap[247]
	if f == nil {
		panic("function 247 not registered")
	}
	f(vmMap[vm])
}

//export f248
func f248(vm *C.WrenVM) {
	f := fMap[248]
	if f == nil {
		panic("function 248 not registered")
	}
	f(vmMap[vm])
}

//export f249
func f249(vm *C.WrenVM) {
	f := fMap[249]
	if f == nil {
		panic("function 249 not registered")
	}
	f(vmMap[vm])
}

//export f250
func f250(vm *C.WrenVM) {
	f := fMap[250]
	if f == nil {
		panic("function 250 not registered")
	}
	f(vmMap[vm])
}

//export f251
func f251(vm *C.WrenVM) {
	f := fMap[251]
	if f == nil {
		panic("function 251 not registered")
	}
	f(vmMap[vm])
}

//export f252
func f252(vm *C.WrenVM) {
	f := fMap[252]
	if f == nil {
		panic("function 252 not registered")
	}
	f(vmMap[vm])
}

//export f253
func f253(vm *C.WrenVM) {
	f := fMap[253]
	if f == nil {
		panic("function 253 not registered")
	}
	f(vmMap[vm])
}

//export f254
func f254(vm *C.WrenVM) {
	f := fMap[254]
	if f == nil {
		panic("function 254 not registered")
	}
	f(vmMap[vm])
}

//export f255
func f255(vm *C.WrenVM) {
	f := fMap[255]
	if f == nil {
		panic("function 255 not registered")
	}
	f(vmMap[vm])
}


func registerFunc(name string, f func(*VM)) (unsafe.Pointer, error) {
	if (fCounter+1) >= MAX_FUNC_REGISTRATIONS {
		return nil, errors.New("maximum function registration reached")
	}

	fMapGuard.Lock()
	defer fMapGuard.Unlock()

	fMap[fCounter] = f
	ptr := C.get_f(C.int(fCounter))
	fCounter++
	return ptr, nil
}
