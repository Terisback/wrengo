package wrengo

/*
#include "wren.h"
extern void c0(void *vm);
extern void c1(void *vm);
extern void c2(void *vm);
extern void c3(void *vm);
extern void c4(void *vm);
extern void c5(void *vm);
extern void c6(void *vm);
extern void c7(void *vm);
extern void c8(void *vm);
extern void c9(void *vm);
extern void c10(void *vm);
extern void c11(void *vm);
extern void c12(void *vm);
extern void c13(void *vm);
extern void c14(void *vm);
extern void c15(void *vm);
extern void c16(void *vm);
extern void c17(void *vm);
extern void c18(void *vm);
extern void c19(void *vm);
extern void c20(void *vm);
extern void c21(void *vm);
extern void c22(void *vm);
extern void c23(void *vm);
extern void c24(void *vm);
extern void c25(void *vm);
extern void c26(void *vm);
extern void c27(void *vm);
extern void c28(void *vm);
extern void c29(void *vm);
extern void c30(void *vm);
extern void c31(void *vm);
extern void c32(void *vm);
extern void c33(void *vm);
extern void c34(void *vm);
extern void c35(void *vm);
extern void c36(void *vm);
extern void c37(void *vm);
extern void c38(void *vm);
extern void c39(void *vm);
extern void c40(void *vm);
extern void c41(void *vm);
extern void c42(void *vm);
extern void c43(void *vm);
extern void c44(void *vm);
extern void c45(void *vm);
extern void c46(void *vm);
extern void c47(void *vm);
extern void c48(void *vm);
extern void c49(void *vm);
extern void c50(void *vm);
extern void c51(void *vm);
extern void c52(void *vm);
extern void c53(void *vm);
extern void c54(void *vm);
extern void c55(void *vm);
extern void c56(void *vm);
extern void c57(void *vm);
extern void c58(void *vm);
extern void c59(void *vm);
extern void c60(void *vm);
extern void c61(void *vm);
extern void c62(void *vm);
extern void c63(void *vm);
extern void c64(void *vm);
extern void c65(void *vm);
extern void c66(void *vm);
extern void c67(void *vm);
extern void c68(void *vm);
extern void c69(void *vm);
extern void c70(void *vm);
extern void c71(void *vm);
extern void c72(void *vm);
extern void c73(void *vm);
extern void c74(void *vm);
extern void c75(void *vm);
extern void c76(void *vm);
extern void c77(void *vm);
extern void c78(void *vm);
extern void c79(void *vm);
extern void c80(void *vm);
extern void c81(void *vm);
extern void c82(void *vm);
extern void c83(void *vm);
extern void c84(void *vm);
extern void c85(void *vm);
extern void c86(void *vm);
extern void c87(void *vm);
extern void c88(void *vm);
extern void c89(void *vm);
extern void c90(void *vm);
extern void c91(void *vm);
extern void c92(void *vm);
extern void c93(void *vm);
extern void c94(void *vm);
extern void c95(void *vm);
extern void c96(void *vm);
extern void c97(void *vm);
extern void c98(void *vm);
extern void c99(void *vm);
extern void c100(void *vm);
extern void c101(void *vm);
extern void c102(void *vm);
extern void c103(void *vm);
extern void c104(void *vm);
extern void c105(void *vm);
extern void c106(void *vm);
extern void c107(void *vm);
extern void c108(void *vm);
extern void c109(void *vm);
extern void c110(void *vm);
extern void c111(void *vm);
extern void c112(void *vm);
extern void c113(void *vm);
extern void c114(void *vm);
extern void c115(void *vm);
extern void c116(void *vm);
extern void c117(void *vm);
extern void c118(void *vm);
extern void c119(void *vm);
extern void c120(void *vm);
extern void c121(void *vm);
extern void c122(void *vm);
extern void c123(void *vm);
extern void c124(void *vm);
extern void c125(void *vm);
extern void c126(void *vm);
extern void c127(void *vm);
extern void c128(void *vm);
extern void c129(void *vm);
extern void c130(void *vm);
extern void c131(void *vm);
extern void c132(void *vm);
extern void c133(void *vm);
extern void c134(void *vm);
extern void c135(void *vm);
extern void c136(void *vm);
extern void c137(void *vm);
extern void c138(void *vm);
extern void c139(void *vm);
extern void c140(void *vm);
extern void c141(void *vm);
extern void c142(void *vm);
extern void c143(void *vm);
extern void c144(void *vm);
extern void c145(void *vm);
extern void c146(void *vm);
extern void c147(void *vm);
extern void c148(void *vm);
extern void c149(void *vm);
extern void c150(void *vm);
extern void c151(void *vm);
extern void c152(void *vm);
extern void c153(void *vm);
extern void c154(void *vm);
extern void c155(void *vm);
extern void c156(void *vm);
extern void c157(void *vm);
extern void c158(void *vm);
extern void c159(void *vm);
extern void c160(void *vm);
extern void c161(void *vm);
extern void c162(void *vm);
extern void c163(void *vm);
extern void c164(void *vm);
extern void c165(void *vm);
extern void c166(void *vm);
extern void c167(void *vm);
extern void c168(void *vm);
extern void c169(void *vm);
extern void c170(void *vm);
extern void c171(void *vm);
extern void c172(void *vm);
extern void c173(void *vm);
extern void c174(void *vm);
extern void c175(void *vm);
extern void c176(void *vm);
extern void c177(void *vm);
extern void c178(void *vm);
extern void c179(void *vm);
extern void c180(void *vm);
extern void c181(void *vm);
extern void c182(void *vm);
extern void c183(void *vm);
extern void c184(void *vm);
extern void c185(void *vm);
extern void c186(void *vm);
extern void c187(void *vm);
extern void c188(void *vm);
extern void c189(void *vm);
extern void c190(void *vm);
extern void c191(void *vm);
extern void c192(void *vm);
extern void c193(void *vm);
extern void c194(void *vm);
extern void c195(void *vm);
extern void c196(void *vm);
extern void c197(void *vm);
extern void c198(void *vm);
extern void c199(void *vm);
extern void c200(void *vm);
extern void c201(void *vm);
extern void c202(void *vm);
extern void c203(void *vm);
extern void c204(void *vm);
extern void c205(void *vm);
extern void c206(void *vm);
extern void c207(void *vm);
extern void c208(void *vm);
extern void c209(void *vm);
extern void c210(void *vm);
extern void c211(void *vm);
extern void c212(void *vm);
extern void c213(void *vm);
extern void c214(void *vm);
extern void c215(void *vm);
extern void c216(void *vm);
extern void c217(void *vm);
extern void c218(void *vm);
extern void c219(void *vm);
extern void c220(void *vm);
extern void c221(void *vm);
extern void c222(void *vm);
extern void c223(void *vm);
extern void c224(void *vm);
extern void c225(void *vm);
extern void c226(void *vm);
extern void c227(void *vm);
extern void c228(void *vm);
extern void c229(void *vm);
extern void c230(void *vm);
extern void c231(void *vm);
extern void c232(void *vm);
extern void c233(void *vm);
extern void c234(void *vm);
extern void c235(void *vm);
extern void c236(void *vm);
extern void c237(void *vm);
extern void c238(void *vm);
extern void c239(void *vm);
extern void c240(void *vm);
extern void c241(void *vm);
extern void c242(void *vm);
extern void c243(void *vm);
extern void c244(void *vm);
extern void c245(void *vm);
extern void c246(void *vm);
extern void c247(void *vm);
extern void c248(void *vm);
extern void c249(void *vm);
extern void c250(void *vm);
extern void c251(void *vm);
extern void c252(void *vm);
extern void c253(void *vm);
extern void c254(void *vm);
extern void c255(void *vm);

static inline void* get_c(int i) {
	switch (i) {
		case 0: return c0;
		case 1: return c1;
		case 2: return c2;
		case 3: return c3;
		case 4: return c4;
		case 5: return c5;
		case 6: return c6;
		case 7: return c7;
		case 8: return c8;
		case 9: return c9;
		case 10: return c10;
		case 11: return c11;
		case 12: return c12;
		case 13: return c13;
		case 14: return c14;
		case 15: return c15;
		case 16: return c16;
		case 17: return c17;
		case 18: return c18;
		case 19: return c19;
		case 20: return c20;
		case 21: return c21;
		case 22: return c22;
		case 23: return c23;
		case 24: return c24;
		case 25: return c25;
		case 26: return c26;
		case 27: return c27;
		case 28: return c28;
		case 29: return c29;
		case 30: return c30;
		case 31: return c31;
		case 32: return c32;
		case 33: return c33;
		case 34: return c34;
		case 35: return c35;
		case 36: return c36;
		case 37: return c37;
		case 38: return c38;
		case 39: return c39;
		case 40: return c40;
		case 41: return c41;
		case 42: return c42;
		case 43: return c43;
		case 44: return c44;
		case 45: return c45;
		case 46: return c46;
		case 47: return c47;
		case 48: return c48;
		case 49: return c49;
		case 50: return c50;
		case 51: return c51;
		case 52: return c52;
		case 53: return c53;
		case 54: return c54;
		case 55: return c55;
		case 56: return c56;
		case 57: return c57;
		case 58: return c58;
		case 59: return c59;
		case 60: return c60;
		case 61: return c61;
		case 62: return c62;
		case 63: return c63;
		case 64: return c64;
		case 65: return c65;
		case 66: return c66;
		case 67: return c67;
		case 68: return c68;
		case 69: return c69;
		case 70: return c70;
		case 71: return c71;
		case 72: return c72;
		case 73: return c73;
		case 74: return c74;
		case 75: return c75;
		case 76: return c76;
		case 77: return c77;
		case 78: return c78;
		case 79: return c79;
		case 80: return c80;
		case 81: return c81;
		case 82: return c82;
		case 83: return c83;
		case 84: return c84;
		case 85: return c85;
		case 86: return c86;
		case 87: return c87;
		case 88: return c88;
		case 89: return c89;
		case 90: return c90;
		case 91: return c91;
		case 92: return c92;
		case 93: return c93;
		case 94: return c94;
		case 95: return c95;
		case 96: return c96;
		case 97: return c97;
		case 98: return c98;
		case 99: return c99;
		case 100: return c100;
		case 101: return c101;
		case 102: return c102;
		case 103: return c103;
		case 104: return c104;
		case 105: return c105;
		case 106: return c106;
		case 107: return c107;
		case 108: return c108;
		case 109: return c109;
		case 110: return c110;
		case 111: return c111;
		case 112: return c112;
		case 113: return c113;
		case 114: return c114;
		case 115: return c115;
		case 116: return c116;
		case 117: return c117;
		case 118: return c118;
		case 119: return c119;
		case 120: return c120;
		case 121: return c121;
		case 122: return c122;
		case 123: return c123;
		case 124: return c124;
		case 125: return c125;
		case 126: return c126;
		case 127: return c127;
		case 128: return c128;
		case 129: return c129;
		case 130: return c130;
		case 131: return c131;
		case 132: return c132;
		case 133: return c133;
		case 134: return c134;
		case 135: return c135;
		case 136: return c136;
		case 137: return c137;
		case 138: return c138;
		case 139: return c139;
		case 140: return c140;
		case 141: return c141;
		case 142: return c142;
		case 143: return c143;
		case 144: return c144;
		case 145: return c145;
		case 146: return c146;
		case 147: return c147;
		case 148: return c148;
		case 149: return c149;
		case 150: return c150;
		case 151: return c151;
		case 152: return c152;
		case 153: return c153;
		case 154: return c154;
		case 155: return c155;
		case 156: return c156;
		case 157: return c157;
		case 158: return c158;
		case 159: return c159;
		case 160: return c160;
		case 161: return c161;
		case 162: return c162;
		case 163: return c163;
		case 164: return c164;
		case 165: return c165;
		case 166: return c166;
		case 167: return c167;
		case 168: return c168;
		case 169: return c169;
		case 170: return c170;
		case 171: return c171;
		case 172: return c172;
		case 173: return c173;
		case 174: return c174;
		case 175: return c175;
		case 176: return c176;
		case 177: return c177;
		case 178: return c178;
		case 179: return c179;
		case 180: return c180;
		case 181: return c181;
		case 182: return c182;
		case 183: return c183;
		case 184: return c184;
		case 185: return c185;
		case 186: return c186;
		case 187: return c187;
		case 188: return c188;
		case 189: return c189;
		case 190: return c190;
		case 191: return c191;
		case 192: return c192;
		case 193: return c193;
		case 194: return c194;
		case 195: return c195;
		case 196: return c196;
		case 197: return c197;
		case 198: return c198;
		case 199: return c199;
		case 200: return c200;
		case 201: return c201;
		case 202: return c202;
		case 203: return c203;
		case 204: return c204;
		case 205: return c205;
		case 206: return c206;
		case 207: return c207;
		case 208: return c208;
		case 209: return c209;
		case 210: return c210;
		case 211: return c211;
		case 212: return c212;
		case 213: return c213;
		case 214: return c214;
		case 215: return c215;
		case 216: return c216;
		case 217: return c217;
		case 218: return c218;
		case 219: return c219;
		case 220: return c220;
		case 221: return c221;
		case 222: return c222;
		case 223: return c223;
		case 224: return c224;
		case 225: return c225;
		case 226: return c226;
		case 227: return c227;
		case 228: return c228;
		case 229: return c229;
		case 230: return c230;
		case 231: return c231;
		case 232: return c232;
		case 233: return c233;
		case 234: return c234;
		case 235: return c235;
		case 236: return c236;
		case 237: return c237;
		case 238: return c238;
		case 239: return c239;
		case 240: return c240;
		case 241: return c241;
		case 242: return c242;
		case 243: return c243;
		case 244: return c244;
		case 245: return c245;
		case 246: return c246;
		case 247: return c247;
		case 248: return c248;
		case 249: return c249;
		case 250: return c250;
		case 251: return c251;
		case 252: return c252;
		case 253: return c253;
		case 254: return c254;
		case 255: return c255;
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

const MAX_CLASS_REGISTRATIONS = 256

var (
	cMap = make(map[int]func())
	cMapGuard sync.Mutex
	cCounter int
)


//export c0
func c0(vm unsafe.Pointer) {
	f := cMap[0]
	if f == nil {
		panic("function 0 not registered")
	}
	f()
}

//export c1
func c1(vm unsafe.Pointer) {
	f := cMap[1]
	if f == nil {
		panic("function 1 not registered")
	}
	f()
}

//export c2
func c2(vm unsafe.Pointer) {
	f := cMap[2]
	if f == nil {
		panic("function 2 not registered")
	}
	f()
}

//export c3
func c3(vm unsafe.Pointer) {
	f := cMap[3]
	if f == nil {
		panic("function 3 not registered")
	}
	f()
}

//export c4
func c4(vm unsafe.Pointer) {
	f := cMap[4]
	if f == nil {
		panic("function 4 not registered")
	}
	f()
}

//export c5
func c5(vm unsafe.Pointer) {
	f := cMap[5]
	if f == nil {
		panic("function 5 not registered")
	}
	f()
}

//export c6
func c6(vm unsafe.Pointer) {
	f := cMap[6]
	if f == nil {
		panic("function 6 not registered")
	}
	f()
}

//export c7
func c7(vm unsafe.Pointer) {
	f := cMap[7]
	if f == nil {
		panic("function 7 not registered")
	}
	f()
}

//export c8
func c8(vm unsafe.Pointer) {
	f := cMap[8]
	if f == nil {
		panic("function 8 not registered")
	}
	f()
}

//export c9
func c9(vm unsafe.Pointer) {
	f := cMap[9]
	if f == nil {
		panic("function 9 not registered")
	}
	f()
}

//export c10
func c10(vm unsafe.Pointer) {
	f := cMap[10]
	if f == nil {
		panic("function 10 not registered")
	}
	f()
}

//export c11
func c11(vm unsafe.Pointer) {
	f := cMap[11]
	if f == nil {
		panic("function 11 not registered")
	}
	f()
}

//export c12
func c12(vm unsafe.Pointer) {
	f := cMap[12]
	if f == nil {
		panic("function 12 not registered")
	}
	f()
}

//export c13
func c13(vm unsafe.Pointer) {
	f := cMap[13]
	if f == nil {
		panic("function 13 not registered")
	}
	f()
}

//export c14
func c14(vm unsafe.Pointer) {
	f := cMap[14]
	if f == nil {
		panic("function 14 not registered")
	}
	f()
}

//export c15
func c15(vm unsafe.Pointer) {
	f := cMap[15]
	if f == nil {
		panic("function 15 not registered")
	}
	f()
}

//export c16
func c16(vm unsafe.Pointer) {
	f := cMap[16]
	if f == nil {
		panic("function 16 not registered")
	}
	f()
}

//export c17
func c17(vm unsafe.Pointer) {
	f := cMap[17]
	if f == nil {
		panic("function 17 not registered")
	}
	f()
}

//export c18
func c18(vm unsafe.Pointer) {
	f := cMap[18]
	if f == nil {
		panic("function 18 not registered")
	}
	f()
}

//export c19
func c19(vm unsafe.Pointer) {
	f := cMap[19]
	if f == nil {
		panic("function 19 not registered")
	}
	f()
}

//export c20
func c20(vm unsafe.Pointer) {
	f := cMap[20]
	if f == nil {
		panic("function 20 not registered")
	}
	f()
}

//export c21
func c21(vm unsafe.Pointer) {
	f := cMap[21]
	if f == nil {
		panic("function 21 not registered")
	}
	f()
}

//export c22
func c22(vm unsafe.Pointer) {
	f := cMap[22]
	if f == nil {
		panic("function 22 not registered")
	}
	f()
}

//export c23
func c23(vm unsafe.Pointer) {
	f := cMap[23]
	if f == nil {
		panic("function 23 not registered")
	}
	f()
}

//export c24
func c24(vm unsafe.Pointer) {
	f := cMap[24]
	if f == nil {
		panic("function 24 not registered")
	}
	f()
}

//export c25
func c25(vm unsafe.Pointer) {
	f := cMap[25]
	if f == nil {
		panic("function 25 not registered")
	}
	f()
}

//export c26
func c26(vm unsafe.Pointer) {
	f := cMap[26]
	if f == nil {
		panic("function 26 not registered")
	}
	f()
}

//export c27
func c27(vm unsafe.Pointer) {
	f := cMap[27]
	if f == nil {
		panic("function 27 not registered")
	}
	f()
}

//export c28
func c28(vm unsafe.Pointer) {
	f := cMap[28]
	if f == nil {
		panic("function 28 not registered")
	}
	f()
}

//export c29
func c29(vm unsafe.Pointer) {
	f := cMap[29]
	if f == nil {
		panic("function 29 not registered")
	}
	f()
}

//export c30
func c30(vm unsafe.Pointer) {
	f := cMap[30]
	if f == nil {
		panic("function 30 not registered")
	}
	f()
}

//export c31
func c31(vm unsafe.Pointer) {
	f := cMap[31]
	if f == nil {
		panic("function 31 not registered")
	}
	f()
}

//export c32
func c32(vm unsafe.Pointer) {
	f := cMap[32]
	if f == nil {
		panic("function 32 not registered")
	}
	f()
}

//export c33
func c33(vm unsafe.Pointer) {
	f := cMap[33]
	if f == nil {
		panic("function 33 not registered")
	}
	f()
}

//export c34
func c34(vm unsafe.Pointer) {
	f := cMap[34]
	if f == nil {
		panic("function 34 not registered")
	}
	f()
}

//export c35
func c35(vm unsafe.Pointer) {
	f := cMap[35]
	if f == nil {
		panic("function 35 not registered")
	}
	f()
}

//export c36
func c36(vm unsafe.Pointer) {
	f := cMap[36]
	if f == nil {
		panic("function 36 not registered")
	}
	f()
}

//export c37
func c37(vm unsafe.Pointer) {
	f := cMap[37]
	if f == nil {
		panic("function 37 not registered")
	}
	f()
}

//export c38
func c38(vm unsafe.Pointer) {
	f := cMap[38]
	if f == nil {
		panic("function 38 not registered")
	}
	f()
}

//export c39
func c39(vm unsafe.Pointer) {
	f := cMap[39]
	if f == nil {
		panic("function 39 not registered")
	}
	f()
}

//export c40
func c40(vm unsafe.Pointer) {
	f := cMap[40]
	if f == nil {
		panic("function 40 not registered")
	}
	f()
}

//export c41
func c41(vm unsafe.Pointer) {
	f := cMap[41]
	if f == nil {
		panic("function 41 not registered")
	}
	f()
}

//export c42
func c42(vm unsafe.Pointer) {
	f := cMap[42]
	if f == nil {
		panic("function 42 not registered")
	}
	f()
}

//export c43
func c43(vm unsafe.Pointer) {
	f := cMap[43]
	if f == nil {
		panic("function 43 not registered")
	}
	f()
}

//export c44
func c44(vm unsafe.Pointer) {
	f := cMap[44]
	if f == nil {
		panic("function 44 not registered")
	}
	f()
}

//export c45
func c45(vm unsafe.Pointer) {
	f := cMap[45]
	if f == nil {
		panic("function 45 not registered")
	}
	f()
}

//export c46
func c46(vm unsafe.Pointer) {
	f := cMap[46]
	if f == nil {
		panic("function 46 not registered")
	}
	f()
}

//export c47
func c47(vm unsafe.Pointer) {
	f := cMap[47]
	if f == nil {
		panic("function 47 not registered")
	}
	f()
}

//export c48
func c48(vm unsafe.Pointer) {
	f := cMap[48]
	if f == nil {
		panic("function 48 not registered")
	}
	f()
}

//export c49
func c49(vm unsafe.Pointer) {
	f := cMap[49]
	if f == nil {
		panic("function 49 not registered")
	}
	f()
}

//export c50
func c50(vm unsafe.Pointer) {
	f := cMap[50]
	if f == nil {
		panic("function 50 not registered")
	}
	f()
}

//export c51
func c51(vm unsafe.Pointer) {
	f := cMap[51]
	if f == nil {
		panic("function 51 not registered")
	}
	f()
}

//export c52
func c52(vm unsafe.Pointer) {
	f := cMap[52]
	if f == nil {
		panic("function 52 not registered")
	}
	f()
}

//export c53
func c53(vm unsafe.Pointer) {
	f := cMap[53]
	if f == nil {
		panic("function 53 not registered")
	}
	f()
}

//export c54
func c54(vm unsafe.Pointer) {
	f := cMap[54]
	if f == nil {
		panic("function 54 not registered")
	}
	f()
}

//export c55
func c55(vm unsafe.Pointer) {
	f := cMap[55]
	if f == nil {
		panic("function 55 not registered")
	}
	f()
}

//export c56
func c56(vm unsafe.Pointer) {
	f := cMap[56]
	if f == nil {
		panic("function 56 not registered")
	}
	f()
}

//export c57
func c57(vm unsafe.Pointer) {
	f := cMap[57]
	if f == nil {
		panic("function 57 not registered")
	}
	f()
}

//export c58
func c58(vm unsafe.Pointer) {
	f := cMap[58]
	if f == nil {
		panic("function 58 not registered")
	}
	f()
}

//export c59
func c59(vm unsafe.Pointer) {
	f := cMap[59]
	if f == nil {
		panic("function 59 not registered")
	}
	f()
}

//export c60
func c60(vm unsafe.Pointer) {
	f := cMap[60]
	if f == nil {
		panic("function 60 not registered")
	}
	f()
}

//export c61
func c61(vm unsafe.Pointer) {
	f := cMap[61]
	if f == nil {
		panic("function 61 not registered")
	}
	f()
}

//export c62
func c62(vm unsafe.Pointer) {
	f := cMap[62]
	if f == nil {
		panic("function 62 not registered")
	}
	f()
}

//export c63
func c63(vm unsafe.Pointer) {
	f := cMap[63]
	if f == nil {
		panic("function 63 not registered")
	}
	f()
}

//export c64
func c64(vm unsafe.Pointer) {
	f := cMap[64]
	if f == nil {
		panic("function 64 not registered")
	}
	f()
}

//export c65
func c65(vm unsafe.Pointer) {
	f := cMap[65]
	if f == nil {
		panic("function 65 not registered")
	}
	f()
}

//export c66
func c66(vm unsafe.Pointer) {
	f := cMap[66]
	if f == nil {
		panic("function 66 not registered")
	}
	f()
}

//export c67
func c67(vm unsafe.Pointer) {
	f := cMap[67]
	if f == nil {
		panic("function 67 not registered")
	}
	f()
}

//export c68
func c68(vm unsafe.Pointer) {
	f := cMap[68]
	if f == nil {
		panic("function 68 not registered")
	}
	f()
}

//export c69
func c69(vm unsafe.Pointer) {
	f := cMap[69]
	if f == nil {
		panic("function 69 not registered")
	}
	f()
}

//export c70
func c70(vm unsafe.Pointer) {
	f := cMap[70]
	if f == nil {
		panic("function 70 not registered")
	}
	f()
}

//export c71
func c71(vm unsafe.Pointer) {
	f := cMap[71]
	if f == nil {
		panic("function 71 not registered")
	}
	f()
}

//export c72
func c72(vm unsafe.Pointer) {
	f := cMap[72]
	if f == nil {
		panic("function 72 not registered")
	}
	f()
}

//export c73
func c73(vm unsafe.Pointer) {
	f := cMap[73]
	if f == nil {
		panic("function 73 not registered")
	}
	f()
}

//export c74
func c74(vm unsafe.Pointer) {
	f := cMap[74]
	if f == nil {
		panic("function 74 not registered")
	}
	f()
}

//export c75
func c75(vm unsafe.Pointer) {
	f := cMap[75]
	if f == nil {
		panic("function 75 not registered")
	}
	f()
}

//export c76
func c76(vm unsafe.Pointer) {
	f := cMap[76]
	if f == nil {
		panic("function 76 not registered")
	}
	f()
}

//export c77
func c77(vm unsafe.Pointer) {
	f := cMap[77]
	if f == nil {
		panic("function 77 not registered")
	}
	f()
}

//export c78
func c78(vm unsafe.Pointer) {
	f := cMap[78]
	if f == nil {
		panic("function 78 not registered")
	}
	f()
}

//export c79
func c79(vm unsafe.Pointer) {
	f := cMap[79]
	if f == nil {
		panic("function 79 not registered")
	}
	f()
}

//export c80
func c80(vm unsafe.Pointer) {
	f := cMap[80]
	if f == nil {
		panic("function 80 not registered")
	}
	f()
}

//export c81
func c81(vm unsafe.Pointer) {
	f := cMap[81]
	if f == nil {
		panic("function 81 not registered")
	}
	f()
}

//export c82
func c82(vm unsafe.Pointer) {
	f := cMap[82]
	if f == nil {
		panic("function 82 not registered")
	}
	f()
}

//export c83
func c83(vm unsafe.Pointer) {
	f := cMap[83]
	if f == nil {
		panic("function 83 not registered")
	}
	f()
}

//export c84
func c84(vm unsafe.Pointer) {
	f := cMap[84]
	if f == nil {
		panic("function 84 not registered")
	}
	f()
}

//export c85
func c85(vm unsafe.Pointer) {
	f := cMap[85]
	if f == nil {
		panic("function 85 not registered")
	}
	f()
}

//export c86
func c86(vm unsafe.Pointer) {
	f := cMap[86]
	if f == nil {
		panic("function 86 not registered")
	}
	f()
}

//export c87
func c87(vm unsafe.Pointer) {
	f := cMap[87]
	if f == nil {
		panic("function 87 not registered")
	}
	f()
}

//export c88
func c88(vm unsafe.Pointer) {
	f := cMap[88]
	if f == nil {
		panic("function 88 not registered")
	}
	f()
}

//export c89
func c89(vm unsafe.Pointer) {
	f := cMap[89]
	if f == nil {
		panic("function 89 not registered")
	}
	f()
}

//export c90
func c90(vm unsafe.Pointer) {
	f := cMap[90]
	if f == nil {
		panic("function 90 not registered")
	}
	f()
}

//export c91
func c91(vm unsafe.Pointer) {
	f := cMap[91]
	if f == nil {
		panic("function 91 not registered")
	}
	f()
}

//export c92
func c92(vm unsafe.Pointer) {
	f := cMap[92]
	if f == nil {
		panic("function 92 not registered")
	}
	f()
}

//export c93
func c93(vm unsafe.Pointer) {
	f := cMap[93]
	if f == nil {
		panic("function 93 not registered")
	}
	f()
}

//export c94
func c94(vm unsafe.Pointer) {
	f := cMap[94]
	if f == nil {
		panic("function 94 not registered")
	}
	f()
}

//export c95
func c95(vm unsafe.Pointer) {
	f := cMap[95]
	if f == nil {
		panic("function 95 not registered")
	}
	f()
}

//export c96
func c96(vm unsafe.Pointer) {
	f := cMap[96]
	if f == nil {
		panic("function 96 not registered")
	}
	f()
}

//export c97
func c97(vm unsafe.Pointer) {
	f := cMap[97]
	if f == nil {
		panic("function 97 not registered")
	}
	f()
}

//export c98
func c98(vm unsafe.Pointer) {
	f := cMap[98]
	if f == nil {
		panic("function 98 not registered")
	}
	f()
}

//export c99
func c99(vm unsafe.Pointer) {
	f := cMap[99]
	if f == nil {
		panic("function 99 not registered")
	}
	f()
}

//export c100
func c100(vm unsafe.Pointer) {
	f := cMap[100]
	if f == nil {
		panic("function 100 not registered")
	}
	f()
}

//export c101
func c101(vm unsafe.Pointer) {
	f := cMap[101]
	if f == nil {
		panic("function 101 not registered")
	}
	f()
}

//export c102
func c102(vm unsafe.Pointer) {
	f := cMap[102]
	if f == nil {
		panic("function 102 not registered")
	}
	f()
}

//export c103
func c103(vm unsafe.Pointer) {
	f := cMap[103]
	if f == nil {
		panic("function 103 not registered")
	}
	f()
}

//export c104
func c104(vm unsafe.Pointer) {
	f := cMap[104]
	if f == nil {
		panic("function 104 not registered")
	}
	f()
}

//export c105
func c105(vm unsafe.Pointer) {
	f := cMap[105]
	if f == nil {
		panic("function 105 not registered")
	}
	f()
}

//export c106
func c106(vm unsafe.Pointer) {
	f := cMap[106]
	if f == nil {
		panic("function 106 not registered")
	}
	f()
}

//export c107
func c107(vm unsafe.Pointer) {
	f := cMap[107]
	if f == nil {
		panic("function 107 not registered")
	}
	f()
}

//export c108
func c108(vm unsafe.Pointer) {
	f := cMap[108]
	if f == nil {
		panic("function 108 not registered")
	}
	f()
}

//export c109
func c109(vm unsafe.Pointer) {
	f := cMap[109]
	if f == nil {
		panic("function 109 not registered")
	}
	f()
}

//export c110
func c110(vm unsafe.Pointer) {
	f := cMap[110]
	if f == nil {
		panic("function 110 not registered")
	}
	f()
}

//export c111
func c111(vm unsafe.Pointer) {
	f := cMap[111]
	if f == nil {
		panic("function 111 not registered")
	}
	f()
}

//export c112
func c112(vm unsafe.Pointer) {
	f := cMap[112]
	if f == nil {
		panic("function 112 not registered")
	}
	f()
}

//export c113
func c113(vm unsafe.Pointer) {
	f := cMap[113]
	if f == nil {
		panic("function 113 not registered")
	}
	f()
}

//export c114
func c114(vm unsafe.Pointer) {
	f := cMap[114]
	if f == nil {
		panic("function 114 not registered")
	}
	f()
}

//export c115
func c115(vm unsafe.Pointer) {
	f := cMap[115]
	if f == nil {
		panic("function 115 not registered")
	}
	f()
}

//export c116
func c116(vm unsafe.Pointer) {
	f := cMap[116]
	if f == nil {
		panic("function 116 not registered")
	}
	f()
}

//export c117
func c117(vm unsafe.Pointer) {
	f := cMap[117]
	if f == nil {
		panic("function 117 not registered")
	}
	f()
}

//export c118
func c118(vm unsafe.Pointer) {
	f := cMap[118]
	if f == nil {
		panic("function 118 not registered")
	}
	f()
}

//export c119
func c119(vm unsafe.Pointer) {
	f := cMap[119]
	if f == nil {
		panic("function 119 not registered")
	}
	f()
}

//export c120
func c120(vm unsafe.Pointer) {
	f := cMap[120]
	if f == nil {
		panic("function 120 not registered")
	}
	f()
}

//export c121
func c121(vm unsafe.Pointer) {
	f := cMap[121]
	if f == nil {
		panic("function 121 not registered")
	}
	f()
}

//export c122
func c122(vm unsafe.Pointer) {
	f := cMap[122]
	if f == nil {
		panic("function 122 not registered")
	}
	f()
}

//export c123
func c123(vm unsafe.Pointer) {
	f := cMap[123]
	if f == nil {
		panic("function 123 not registered")
	}
	f()
}

//export c124
func c124(vm unsafe.Pointer) {
	f := cMap[124]
	if f == nil {
		panic("function 124 not registered")
	}
	f()
}

//export c125
func c125(vm unsafe.Pointer) {
	f := cMap[125]
	if f == nil {
		panic("function 125 not registered")
	}
	f()
}

//export c126
func c126(vm unsafe.Pointer) {
	f := cMap[126]
	if f == nil {
		panic("function 126 not registered")
	}
	f()
}

//export c127
func c127(vm unsafe.Pointer) {
	f := cMap[127]
	if f == nil {
		panic("function 127 not registered")
	}
	f()
}

//export c128
func c128(vm unsafe.Pointer) {
	f := cMap[128]
	if f == nil {
		panic("function 128 not registered")
	}
	f()
}

//export c129
func c129(vm unsafe.Pointer) {
	f := cMap[129]
	if f == nil {
		panic("function 129 not registered")
	}
	f()
}

//export c130
func c130(vm unsafe.Pointer) {
	f := cMap[130]
	if f == nil {
		panic("function 130 not registered")
	}
	f()
}

//export c131
func c131(vm unsafe.Pointer) {
	f := cMap[131]
	if f == nil {
		panic("function 131 not registered")
	}
	f()
}

//export c132
func c132(vm unsafe.Pointer) {
	f := cMap[132]
	if f == nil {
		panic("function 132 not registered")
	}
	f()
}

//export c133
func c133(vm unsafe.Pointer) {
	f := cMap[133]
	if f == nil {
		panic("function 133 not registered")
	}
	f()
}

//export c134
func c134(vm unsafe.Pointer) {
	f := cMap[134]
	if f == nil {
		panic("function 134 not registered")
	}
	f()
}

//export c135
func c135(vm unsafe.Pointer) {
	f := cMap[135]
	if f == nil {
		panic("function 135 not registered")
	}
	f()
}

//export c136
func c136(vm unsafe.Pointer) {
	f := cMap[136]
	if f == nil {
		panic("function 136 not registered")
	}
	f()
}

//export c137
func c137(vm unsafe.Pointer) {
	f := cMap[137]
	if f == nil {
		panic("function 137 not registered")
	}
	f()
}

//export c138
func c138(vm unsafe.Pointer) {
	f := cMap[138]
	if f == nil {
		panic("function 138 not registered")
	}
	f()
}

//export c139
func c139(vm unsafe.Pointer) {
	f := cMap[139]
	if f == nil {
		panic("function 139 not registered")
	}
	f()
}

//export c140
func c140(vm unsafe.Pointer) {
	f := cMap[140]
	if f == nil {
		panic("function 140 not registered")
	}
	f()
}

//export c141
func c141(vm unsafe.Pointer) {
	f := cMap[141]
	if f == nil {
		panic("function 141 not registered")
	}
	f()
}

//export c142
func c142(vm unsafe.Pointer) {
	f := cMap[142]
	if f == nil {
		panic("function 142 not registered")
	}
	f()
}

//export c143
func c143(vm unsafe.Pointer) {
	f := cMap[143]
	if f == nil {
		panic("function 143 not registered")
	}
	f()
}

//export c144
func c144(vm unsafe.Pointer) {
	f := cMap[144]
	if f == nil {
		panic("function 144 not registered")
	}
	f()
}

//export c145
func c145(vm unsafe.Pointer) {
	f := cMap[145]
	if f == nil {
		panic("function 145 not registered")
	}
	f()
}

//export c146
func c146(vm unsafe.Pointer) {
	f := cMap[146]
	if f == nil {
		panic("function 146 not registered")
	}
	f()
}

//export c147
func c147(vm unsafe.Pointer) {
	f := cMap[147]
	if f == nil {
		panic("function 147 not registered")
	}
	f()
}

//export c148
func c148(vm unsafe.Pointer) {
	f := cMap[148]
	if f == nil {
		panic("function 148 not registered")
	}
	f()
}

//export c149
func c149(vm unsafe.Pointer) {
	f := cMap[149]
	if f == nil {
		panic("function 149 not registered")
	}
	f()
}

//export c150
func c150(vm unsafe.Pointer) {
	f := cMap[150]
	if f == nil {
		panic("function 150 not registered")
	}
	f()
}

//export c151
func c151(vm unsafe.Pointer) {
	f := cMap[151]
	if f == nil {
		panic("function 151 not registered")
	}
	f()
}

//export c152
func c152(vm unsafe.Pointer) {
	f := cMap[152]
	if f == nil {
		panic("function 152 not registered")
	}
	f()
}

//export c153
func c153(vm unsafe.Pointer) {
	f := cMap[153]
	if f == nil {
		panic("function 153 not registered")
	}
	f()
}

//export c154
func c154(vm unsafe.Pointer) {
	f := cMap[154]
	if f == nil {
		panic("function 154 not registered")
	}
	f()
}

//export c155
func c155(vm unsafe.Pointer) {
	f := cMap[155]
	if f == nil {
		panic("function 155 not registered")
	}
	f()
}

//export c156
func c156(vm unsafe.Pointer) {
	f := cMap[156]
	if f == nil {
		panic("function 156 not registered")
	}
	f()
}

//export c157
func c157(vm unsafe.Pointer) {
	f := cMap[157]
	if f == nil {
		panic("function 157 not registered")
	}
	f()
}

//export c158
func c158(vm unsafe.Pointer) {
	f := cMap[158]
	if f == nil {
		panic("function 158 not registered")
	}
	f()
}

//export c159
func c159(vm unsafe.Pointer) {
	f := cMap[159]
	if f == nil {
		panic("function 159 not registered")
	}
	f()
}

//export c160
func c160(vm unsafe.Pointer) {
	f := cMap[160]
	if f == nil {
		panic("function 160 not registered")
	}
	f()
}

//export c161
func c161(vm unsafe.Pointer) {
	f := cMap[161]
	if f == nil {
		panic("function 161 not registered")
	}
	f()
}

//export c162
func c162(vm unsafe.Pointer) {
	f := cMap[162]
	if f == nil {
		panic("function 162 not registered")
	}
	f()
}

//export c163
func c163(vm unsafe.Pointer) {
	f := cMap[163]
	if f == nil {
		panic("function 163 not registered")
	}
	f()
}

//export c164
func c164(vm unsafe.Pointer) {
	f := cMap[164]
	if f == nil {
		panic("function 164 not registered")
	}
	f()
}

//export c165
func c165(vm unsafe.Pointer) {
	f := cMap[165]
	if f == nil {
		panic("function 165 not registered")
	}
	f()
}

//export c166
func c166(vm unsafe.Pointer) {
	f := cMap[166]
	if f == nil {
		panic("function 166 not registered")
	}
	f()
}

//export c167
func c167(vm unsafe.Pointer) {
	f := cMap[167]
	if f == nil {
		panic("function 167 not registered")
	}
	f()
}

//export c168
func c168(vm unsafe.Pointer) {
	f := cMap[168]
	if f == nil {
		panic("function 168 not registered")
	}
	f()
}

//export c169
func c169(vm unsafe.Pointer) {
	f := cMap[169]
	if f == nil {
		panic("function 169 not registered")
	}
	f()
}

//export c170
func c170(vm unsafe.Pointer) {
	f := cMap[170]
	if f == nil {
		panic("function 170 not registered")
	}
	f()
}

//export c171
func c171(vm unsafe.Pointer) {
	f := cMap[171]
	if f == nil {
		panic("function 171 not registered")
	}
	f()
}

//export c172
func c172(vm unsafe.Pointer) {
	f := cMap[172]
	if f == nil {
		panic("function 172 not registered")
	}
	f()
}

//export c173
func c173(vm unsafe.Pointer) {
	f := cMap[173]
	if f == nil {
		panic("function 173 not registered")
	}
	f()
}

//export c174
func c174(vm unsafe.Pointer) {
	f := cMap[174]
	if f == nil {
		panic("function 174 not registered")
	}
	f()
}

//export c175
func c175(vm unsafe.Pointer) {
	f := cMap[175]
	if f == nil {
		panic("function 175 not registered")
	}
	f()
}

//export c176
func c176(vm unsafe.Pointer) {
	f := cMap[176]
	if f == nil {
		panic("function 176 not registered")
	}
	f()
}

//export c177
func c177(vm unsafe.Pointer) {
	f := cMap[177]
	if f == nil {
		panic("function 177 not registered")
	}
	f()
}

//export c178
func c178(vm unsafe.Pointer) {
	f := cMap[178]
	if f == nil {
		panic("function 178 not registered")
	}
	f()
}

//export c179
func c179(vm unsafe.Pointer) {
	f := cMap[179]
	if f == nil {
		panic("function 179 not registered")
	}
	f()
}

//export c180
func c180(vm unsafe.Pointer) {
	f := cMap[180]
	if f == nil {
		panic("function 180 not registered")
	}
	f()
}

//export c181
func c181(vm unsafe.Pointer) {
	f := cMap[181]
	if f == nil {
		panic("function 181 not registered")
	}
	f()
}

//export c182
func c182(vm unsafe.Pointer) {
	f := cMap[182]
	if f == nil {
		panic("function 182 not registered")
	}
	f()
}

//export c183
func c183(vm unsafe.Pointer) {
	f := cMap[183]
	if f == nil {
		panic("function 183 not registered")
	}
	f()
}

//export c184
func c184(vm unsafe.Pointer) {
	f := cMap[184]
	if f == nil {
		panic("function 184 not registered")
	}
	f()
}

//export c185
func c185(vm unsafe.Pointer) {
	f := cMap[185]
	if f == nil {
		panic("function 185 not registered")
	}
	f()
}

//export c186
func c186(vm unsafe.Pointer) {
	f := cMap[186]
	if f == nil {
		panic("function 186 not registered")
	}
	f()
}

//export c187
func c187(vm unsafe.Pointer) {
	f := cMap[187]
	if f == nil {
		panic("function 187 not registered")
	}
	f()
}

//export c188
func c188(vm unsafe.Pointer) {
	f := cMap[188]
	if f == nil {
		panic("function 188 not registered")
	}
	f()
}

//export c189
func c189(vm unsafe.Pointer) {
	f := cMap[189]
	if f == nil {
		panic("function 189 not registered")
	}
	f()
}

//export c190
func c190(vm unsafe.Pointer) {
	f := cMap[190]
	if f == nil {
		panic("function 190 not registered")
	}
	f()
}

//export c191
func c191(vm unsafe.Pointer) {
	f := cMap[191]
	if f == nil {
		panic("function 191 not registered")
	}
	f()
}

//export c192
func c192(vm unsafe.Pointer) {
	f := cMap[192]
	if f == nil {
		panic("function 192 not registered")
	}
	f()
}

//export c193
func c193(vm unsafe.Pointer) {
	f := cMap[193]
	if f == nil {
		panic("function 193 not registered")
	}
	f()
}

//export c194
func c194(vm unsafe.Pointer) {
	f := cMap[194]
	if f == nil {
		panic("function 194 not registered")
	}
	f()
}

//export c195
func c195(vm unsafe.Pointer) {
	f := cMap[195]
	if f == nil {
		panic("function 195 not registered")
	}
	f()
}

//export c196
func c196(vm unsafe.Pointer) {
	f := cMap[196]
	if f == nil {
		panic("function 196 not registered")
	}
	f()
}

//export c197
func c197(vm unsafe.Pointer) {
	f := cMap[197]
	if f == nil {
		panic("function 197 not registered")
	}
	f()
}

//export c198
func c198(vm unsafe.Pointer) {
	f := cMap[198]
	if f == nil {
		panic("function 198 not registered")
	}
	f()
}

//export c199
func c199(vm unsafe.Pointer) {
	f := cMap[199]
	if f == nil {
		panic("function 199 not registered")
	}
	f()
}

//export c200
func c200(vm unsafe.Pointer) {
	f := cMap[200]
	if f == nil {
		panic("function 200 not registered")
	}
	f()
}

//export c201
func c201(vm unsafe.Pointer) {
	f := cMap[201]
	if f == nil {
		panic("function 201 not registered")
	}
	f()
}

//export c202
func c202(vm unsafe.Pointer) {
	f := cMap[202]
	if f == nil {
		panic("function 202 not registered")
	}
	f()
}

//export c203
func c203(vm unsafe.Pointer) {
	f := cMap[203]
	if f == nil {
		panic("function 203 not registered")
	}
	f()
}

//export c204
func c204(vm unsafe.Pointer) {
	f := cMap[204]
	if f == nil {
		panic("function 204 not registered")
	}
	f()
}

//export c205
func c205(vm unsafe.Pointer) {
	f := cMap[205]
	if f == nil {
		panic("function 205 not registered")
	}
	f()
}

//export c206
func c206(vm unsafe.Pointer) {
	f := cMap[206]
	if f == nil {
		panic("function 206 not registered")
	}
	f()
}

//export c207
func c207(vm unsafe.Pointer) {
	f := cMap[207]
	if f == nil {
		panic("function 207 not registered")
	}
	f()
}

//export c208
func c208(vm unsafe.Pointer) {
	f := cMap[208]
	if f == nil {
		panic("function 208 not registered")
	}
	f()
}

//export c209
func c209(vm unsafe.Pointer) {
	f := cMap[209]
	if f == nil {
		panic("function 209 not registered")
	}
	f()
}

//export c210
func c210(vm unsafe.Pointer) {
	f := cMap[210]
	if f == nil {
		panic("function 210 not registered")
	}
	f()
}

//export c211
func c211(vm unsafe.Pointer) {
	f := cMap[211]
	if f == nil {
		panic("function 211 not registered")
	}
	f()
}

//export c212
func c212(vm unsafe.Pointer) {
	f := cMap[212]
	if f == nil {
		panic("function 212 not registered")
	}
	f()
}

//export c213
func c213(vm unsafe.Pointer) {
	f := cMap[213]
	if f == nil {
		panic("function 213 not registered")
	}
	f()
}

//export c214
func c214(vm unsafe.Pointer) {
	f := cMap[214]
	if f == nil {
		panic("function 214 not registered")
	}
	f()
}

//export c215
func c215(vm unsafe.Pointer) {
	f := cMap[215]
	if f == nil {
		panic("function 215 not registered")
	}
	f()
}

//export c216
func c216(vm unsafe.Pointer) {
	f := cMap[216]
	if f == nil {
		panic("function 216 not registered")
	}
	f()
}

//export c217
func c217(vm unsafe.Pointer) {
	f := cMap[217]
	if f == nil {
		panic("function 217 not registered")
	}
	f()
}

//export c218
func c218(vm unsafe.Pointer) {
	f := cMap[218]
	if f == nil {
		panic("function 218 not registered")
	}
	f()
}

//export c219
func c219(vm unsafe.Pointer) {
	f := cMap[219]
	if f == nil {
		panic("function 219 not registered")
	}
	f()
}

//export c220
func c220(vm unsafe.Pointer) {
	f := cMap[220]
	if f == nil {
		panic("function 220 not registered")
	}
	f()
}

//export c221
func c221(vm unsafe.Pointer) {
	f := cMap[221]
	if f == nil {
		panic("function 221 not registered")
	}
	f()
}

//export c222
func c222(vm unsafe.Pointer) {
	f := cMap[222]
	if f == nil {
		panic("function 222 not registered")
	}
	f()
}

//export c223
func c223(vm unsafe.Pointer) {
	f := cMap[223]
	if f == nil {
		panic("function 223 not registered")
	}
	f()
}

//export c224
func c224(vm unsafe.Pointer) {
	f := cMap[224]
	if f == nil {
		panic("function 224 not registered")
	}
	f()
}

//export c225
func c225(vm unsafe.Pointer) {
	f := cMap[225]
	if f == nil {
		panic("function 225 not registered")
	}
	f()
}

//export c226
func c226(vm unsafe.Pointer) {
	f := cMap[226]
	if f == nil {
		panic("function 226 not registered")
	}
	f()
}

//export c227
func c227(vm unsafe.Pointer) {
	f := cMap[227]
	if f == nil {
		panic("function 227 not registered")
	}
	f()
}

//export c228
func c228(vm unsafe.Pointer) {
	f := cMap[228]
	if f == nil {
		panic("function 228 not registered")
	}
	f()
}

//export c229
func c229(vm unsafe.Pointer) {
	f := cMap[229]
	if f == nil {
		panic("function 229 not registered")
	}
	f()
}

//export c230
func c230(vm unsafe.Pointer) {
	f := cMap[230]
	if f == nil {
		panic("function 230 not registered")
	}
	f()
}

//export c231
func c231(vm unsafe.Pointer) {
	f := cMap[231]
	if f == nil {
		panic("function 231 not registered")
	}
	f()
}

//export c232
func c232(vm unsafe.Pointer) {
	f := cMap[232]
	if f == nil {
		panic("function 232 not registered")
	}
	f()
}

//export c233
func c233(vm unsafe.Pointer) {
	f := cMap[233]
	if f == nil {
		panic("function 233 not registered")
	}
	f()
}

//export c234
func c234(vm unsafe.Pointer) {
	f := cMap[234]
	if f == nil {
		panic("function 234 not registered")
	}
	f()
}

//export c235
func c235(vm unsafe.Pointer) {
	f := cMap[235]
	if f == nil {
		panic("function 235 not registered")
	}
	f()
}

//export c236
func c236(vm unsafe.Pointer) {
	f := cMap[236]
	if f == nil {
		panic("function 236 not registered")
	}
	f()
}

//export c237
func c237(vm unsafe.Pointer) {
	f := cMap[237]
	if f == nil {
		panic("function 237 not registered")
	}
	f()
}

//export c238
func c238(vm unsafe.Pointer) {
	f := cMap[238]
	if f == nil {
		panic("function 238 not registered")
	}
	f()
}

//export c239
func c239(vm unsafe.Pointer) {
	f := cMap[239]
	if f == nil {
		panic("function 239 not registered")
	}
	f()
}

//export c240
func c240(vm unsafe.Pointer) {
	f := cMap[240]
	if f == nil {
		panic("function 240 not registered")
	}
	f()
}

//export c241
func c241(vm unsafe.Pointer) {
	f := cMap[241]
	if f == nil {
		panic("function 241 not registered")
	}
	f()
}

//export c242
func c242(vm unsafe.Pointer) {
	f := cMap[242]
	if f == nil {
		panic("function 242 not registered")
	}
	f()
}

//export c243
func c243(vm unsafe.Pointer) {
	f := cMap[243]
	if f == nil {
		panic("function 243 not registered")
	}
	f()
}

//export c244
func c244(vm unsafe.Pointer) {
	f := cMap[244]
	if f == nil {
		panic("function 244 not registered")
	}
	f()
}

//export c245
func c245(vm unsafe.Pointer) {
	f := cMap[245]
	if f == nil {
		panic("function 245 not registered")
	}
	f()
}

//export c246
func c246(vm unsafe.Pointer) {
	f := cMap[246]
	if f == nil {
		panic("function 246 not registered")
	}
	f()
}

//export c247
func c247(vm unsafe.Pointer) {
	f := cMap[247]
	if f == nil {
		panic("function 247 not registered")
	}
	f()
}

//export c248
func c248(vm unsafe.Pointer) {
	f := cMap[248]
	if f == nil {
		panic("function 248 not registered")
	}
	f()
}

//export c249
func c249(vm unsafe.Pointer) {
	f := cMap[249]
	if f == nil {
		panic("function 249 not registered")
	}
	f()
}

//export c250
func c250(vm unsafe.Pointer) {
	f := cMap[250]
	if f == nil {
		panic("function 250 not registered")
	}
	f()
}

//export c251
func c251(vm unsafe.Pointer) {
	f := cMap[251]
	if f == nil {
		panic("function 251 not registered")
	}
	f()
}

//export c252
func c252(vm unsafe.Pointer) {
	f := cMap[252]
	if f == nil {
		panic("function 252 not registered")
	}
	f()
}

//export c253
func c253(vm unsafe.Pointer) {
	f := cMap[253]
	if f == nil {
		panic("function 253 not registered")
	}
	f()
}

//export c254
func c254(vm unsafe.Pointer) {
	f := cMap[254]
	if f == nil {
		panic("function 254 not registered")
	}
	f()
}

//export c255
func c255(vm unsafe.Pointer) {
	f := cMap[255]
	if f == nil {
		panic("function 255 not registered")
	}
	f()
}


func registerClass(name string, f func()) (unsafe.Pointer, error) {
	if (cCounter+1) >= MAX_CLASS_REGISTRATIONS {
		return nil, errors.New("maximum function registration reached")
	}

	cMapGuard.Lock()
	defer cMapGuard.Unlock()

	cMap[cCounter] = f
	ptr := C.get_c(C.int(cCounter))
	cCounter++
	return ptr, nil
}
