pragma solidity >=0.5.0;

contract ECC {

  uint8 constant keyTypeAdmin = 1;
  uint8 constant keyTypeDevice = 2;
  uint8 constant keyTypeMf = 3;

  int64 constant proposalDuration = 1 * 60 * 60; // 1 hour

  mapping (address => bool) public devices;

  mapping (address => bool) public admins;

  mapping (address => bool) public mf;

  struct Timestamp {
    int64 unixTime;
    bytes32 hash;
    address device;
  }

  mapping (bytes32 => Timestamp) public timestamps;

  struct Firmware {
    int64 unixTime;
    bytes32 hash;
    address admin;
    address mf;
    bytes mfSig;
  }

  mapping (bytes32 => Firmware) public firmwares;

  struct Proposal {
    int64 unixTime;
    uint8 keyType;
    address key;
    address proposer;
    mapping (address => bool) voted;
    int64 voteCount;
  }

  mapping (address => Proposal) public addKeyProposals;

  mapping (address => Proposal) public removeKeyProposals;

  constructor() public {

    address[] memory clientkeys = new address[](400);

    clientkeys[0] = 0x17Dc6cA2e1C84Ae4107975a48DFD05831b8AdDfF;
    clientkeys[1] = 0xAC5580aD28A3c0e044A52541785bFd34c753D3Bf;
    clientkeys[2] = 0xD0A73Fe9D44184e9F1264ce2097064212E67EbFE;
    clientkeys[3] = 0xd3CF32A24e4527Cf84EF0c8d571f8a77F6481984;
    clientkeys[4] = 0x763a9766405878E73fA45f26c95fE6d43913ff44;
    clientkeys[5] = 0x0e88FA918b9dd6d1a11285864D7350e6d923De40;
    clientkeys[6] = 0x00a80682712dbf93Df472233Ea027b38cEF73F70;
    clientkeys[7] = 0xB7797b168aE2c33a68E19756543A8Aca7C40053f;
    clientkeys[8] = 0x0e0a2C1D9484436499cd900D5Fc0278E5D274Ca7;
    clientkeys[9] = 0x62B0165c43f6378e497A3D91100e99C559e17d1B;
    clientkeys[10] = 0xe504799d748393a53Fa1009097f296c9bD8D8DD1;
    clientkeys[11] = 0xee994E752F37FbB06997C8d9E4F78EB5621cd1E0;
    clientkeys[12] = 0xc28CFE6e78E37a034e18BB87D458244EDdFF8167;
    clientkeys[13] = 0x6c902D9a10836968F3721a4c1E635f28bCc27dda;
    clientkeys[14] = 0x152adCCE72D9F0a8664847158138b5516357Ff40;
    clientkeys[15] = 0x9c8c770954A0E4f4313E884368EB33A48Ba79547;
    clientkeys[16] = 0x936a8FEB84A5B5403C70678e5DA1710Afd803558;
    clientkeys[17] = 0xAC6C0DAE19DCe2d1c8952C159c2a018D68E22360;
    clientkeys[18] = 0x125936228a489Db7E2E9b8aA90b7820E4FE1bFB0;
    clientkeys[19] = 0xC6bCCA3683BF39B1Cf55D75c82Ef57f432f7B9E2;
    clientkeys[20] = 0xB3372b8312679B4DB4252fc16da5f2e02B8624dB;
    clientkeys[21] = 0x0D7f9eF3A873b7e1b89E9fe31503072Fc3b5c557;
    clientkeys[22] = 0xe2Cf703dF8F2fDCAF2a5A54e5E8b88e640f6792a;
    clientkeys[23] = 0xE0b471a7F7A8BE9085937c3499488Fb155a67007;
    clientkeys[24] = 0xC65B5d83bd61BA3AA50A67A2649Bc85D303f4116;
    clientkeys[25] = 0xe503721F3E82A505Ffe7cc83DDDC768142Df93a3;
    clientkeys[26] = 0x5AC36deA5C219ee49037fb82c9e6D6A58bA4b3a8;
    clientkeys[27] = 0xAE4A0660813aE865a36E79fFaCAA907aED78B4b4;
    clientkeys[28] = 0x15ee24641049e10EDab9bF2862Fa603d9636ABEA;
    clientkeys[29] = 0xfB0dbFd4dbfA13d048e2c8A356F583E630D562DE;
    clientkeys[30] = 0x0350E4A3B51f2388a4BB53214196BF87dB88cDef;
    clientkeys[31] = 0xF8795c4fd7BDEd7E2F30559759E8faC0c8E11370;
    clientkeys[32] = 0x66C9C684fa5599C639ef9BD608B220EA0FE1BDcB;
    clientkeys[33] = 0xE9E280Ad18CD055AEfC5FB5a9A8903cc650ef379;
    clientkeys[34] = 0xF1b7280708E359DeBd871F38477c3535E2A5f145;
    clientkeys[35] = 0x6Eb19E67D3625BFA049efBFDB9D008B7438440F2;
    clientkeys[36] = 0xd19ABa0d1e2B30e764eD315ef26CaDEB0f430FB5;
    clientkeys[37] = 0x484b896ee33C0e0B75f83a7fB240A81c0a2e3170;
    clientkeys[38] = 0x62f2379E46D96082D99B22FB32CB32Cd03f02e56;
    clientkeys[39] = 0x3363460f013d0D4d5f7a95fD1c71D084749634e7;
    clientkeys[40] = 0x8b0cFE38d019537626BB3E470aa445f2828eda35;
    clientkeys[41] = 0x1536B7f6d1dB9343C777Ae6E76e6D1aa5b2032bE;
    clientkeys[42] = 0x443bdf162b3b72F1Cd84070eB722aB7907b198fD;
    clientkeys[43] = 0x28Aa44764D975E82D56B05Bebd3aFa2e8bA41D85;
    clientkeys[44] = 0x1F0905f7f26CD26C3557E75Aad7A456F97A16935;
    clientkeys[45] = 0x98Ab5765061F54694950b5a4D8ceE38614bFFf45;
    clientkeys[46] = 0x2b02Ed8fAab38E609106DbeCC61deaE11F5A0D40;
    clientkeys[47] = 0x6df136da0fe91c4bCBA6275DfA072823d6710901;
    clientkeys[48] = 0x358Aa690691335bFcFEeFbe1c49BD6e6C405F5D9;
    clientkeys[49] = 0x09A83011fD8c10FF12687C3055F1a3DD280fb8D0;
    clientkeys[50] = 0xc151B858c2e080FA4CD20bdD8F93D0B759407E2B;
    // clientkeys[51] = 0x3Dcea5dd31bab4A2CAD0e7D90E7cCb492E3A467D;
    // clientkeys[52] = 0x46b1F48347630f2d1e66E14f38716056471737e3;
    // clientkeys[53] = 0x2829dCB8776E66D5367FB527368C35E867AaeE57;
    // clientkeys[54] = 0xA8E101B57503aC4aD33Df2b048C1D0e4498EC956;
    // clientkeys[55] = 0xE432B1C0888B3113D54d5B609c3514d34C341142;
    // clientkeys[56] = 0x092Ec593CA8F3e13F1ef75aFCEC3B8B67924b01d;
    // clientkeys[57] = 0x5f0bA72Bb25ac3e58221B6277010b849A82A60c6;
    // clientkeys[58] = 0x94D88c06fBbcD0b8d896b84BFbb91FD1A3Cd8174;
    // clientkeys[59] = 0xFb2c09fE403EA74F9ef139551713C56cBE15382c;
    // clientkeys[60] = 0x383d81533d17A6e51414b3671CC4A64e0de0Fd6e;
    // clientkeys[61] = 0x39C6DAc8B9BDd11F262D0E3DfAA1dCc11F458638;
    // clientkeys[62] = 0x645E95545839FE2646B8B808587F7321C9b70bCf;
    // clientkeys[63] = 0x16660fe3fc856eD85167d4387f5F676ce03645aF;
    // clientkeys[64] = 0x6fae1cE8703225ef623c118F03f9A768FC9c9a83;
    // clientkeys[65] = 0xB540Bd13a89AC50581FC44a118AC019b064F5A72;
    // clientkeys[66] = 0xB0bbAdac649EE5E0288D0C18b5Eee8F95DF0069a;
    // clientkeys[67] = 0x292Cb3D175f3dD715B84E45C75F1F21F22635E41;
    // clientkeys[68] = 0x7321e0419b6353016B098D3e3e2c139a15671C64;
    // clientkeys[69] = 0xf679A2e90001Efd8e917EC4bf0Cda5Dac71AE1Ac;
    // clientkeys[70] = 0x812384Bb2D58dD5575c0A81ed8025b5852e43918;
    // clientkeys[71] = 0x9545C1cBf0EFE74130E2733484e7383393d19713;
    // clientkeys[72] = 0xF910bA3fF2504eF50347fA2e3b5E6da08b453Faa;
    // clientkeys[73] = 0xdf2Af3311CEBD7c162D9564731085815884A2198;
    // clientkeys[74] = 0x439aD3fDE4d3Ad72d179B75bB7e80497555AAa1E;
    // clientkeys[75] = 0xfCe8365ed959Ea49b0Ee85DFbA30ACeaBC1f6B10;
    // clientkeys[76] = 0xC6770c3E08Dd85f4CAcf6c81c293B7DDd3ECD8A2;
    // clientkeys[77] = 0x0BeCbA1162BF329E4e662b4B111a99f7858dC1F2;
    // clientkeys[78] = 0x5B43dA48Dc2b8570961e75a73a5b22943B780759;
    // clientkeys[79] = 0xcCb638B13F58079a639EfCa1A6fcAd1983ad86bF;
    // clientkeys[80] = 0x23d5EC2D8e3133949cfFbF4f53AA726B53c68329;
    // clientkeys[81] = 0x237A7A03806A4B7bB5e71bF8C0374A5c464DfFC8;
    // clientkeys[82] = 0x0e77BC518F1B1b68F45956dE47a7c0eF6597e380;
    // clientkeys[83] = 0xa0A0CAB896122D9Ca1f4892aD0d772EDFCA70cec;
    // clientkeys[84] = 0xba7D503Bd6CF8cfBA6b6771855Ad266Dd3EB6236;
    // clientkeys[85] = 0x10a0A17AA4BeAaB8aa7268fD57F58B56b6f998c5;
    // clientkeys[86] = 0xAD1d0EF785646d6f5efBF89BBE276bfe4b8D6f3e;
    // clientkeys[87] = 0x65A4149018d2eA79d70185d027adD5DE4dB46001;
    // clientkeys[88] = 0x337E084dD32F510444d4E2C17F94113854738795;
    // clientkeys[89] = 0x9Fa765B75ee6331225675f50864830F5b1Adf9a5;
    // clientkeys[90] = 0xF70539d1f6023DA7c9D16053477fCd0fc153144b;
    // clientkeys[91] = 0x3a087fD58b01D50146a06E2365C79e0C5E80e6AE;
    // clientkeys[92] = 0xb2e9F4ff47ab236671c6Cd940f56b401e4cB8ef2;
    // clientkeys[93] = 0xB75ccbf0451b705A415BbfD2da0c17fa754107Aa;
    // clientkeys[94] = 0x99C95e5347262F530269EE089d9493202977e776;
    // clientkeys[95] = 0x5E1303455842c0b733755DdA55A445b8ca0C40b9;
    // clientkeys[96] = 0xB325C5Aa8BA2f49Da9275373C7463b8B41635638;
    // clientkeys[97] = 0x8205D88900b228Fb2d244f169024a72548d61c13;
    // clientkeys[98] = 0xD48804e442944Eac02862EF57A7B84d0B71f97D3;
    // clientkeys[99] = 0x93D5818C7dbf11a6613eAB73bc20E7BFcFEEbFd4;
    // clientkeys[100] = 0xad43899633BFc63C964c930D2894443f9a89e1ac;
    // clientkeys[101] = 0x8916DE28fec2D244Df09B9fb0e5c9B90284b60A4;
    // clientkeys[102] = 0x6ac47C9547A1F4B7a231738d2E8007ea0B8B71A7;
    // clientkeys[103] = 0x3a99c1FB5D5FcDf61EBB9Ab1421D6A9f44b34cC3;
    // clientkeys[104] = 0x54BC0C0e2Ca8E77cEEe4d8fe0Ff4B5075974044f;
    // clientkeys[105] = 0x5ED209902Dee08593dd3F8421F6eF04aE2c21218;
    // clientkeys[106] = 0x0524A7338A8f1D5B4C99E83EAbCc1a6E99c61B78;
    // clientkeys[107] = 0xc88cf4c6eAd682DFF72a68a8273445091FbD440A;
    // clientkeys[108] = 0x041f93055D4962ff9757208bab2c673B38F20fcD;
    // clientkeys[109] = 0x4C1e3A73A65C2A4D3cD3de2717107C08D5e41221;
    // clientkeys[110] = 0xE8fC477aEF940Bd3a2E535da70775Bb6474e3E20;
    // clientkeys[111] = 0xeC13f0DBc21A6cCc85a4FED6E56454bd0fB4eA35;
    // clientkeys[112] = 0x19af195c0f84b8C35092892875aB220748a9e857;
    // clientkeys[113] = 0x6e2d2D665621CbF77A125eCb30F5bf7661EE3856;
    // clientkeys[114] = 0x2a347C40230DE10f8B86428183Bbd67bC6E0D81A;
    // clientkeys[115] = 0x0117856f25cBDa99F38D9186B3A909226099c9B2;
    // clientkeys[116] = 0x7D398F1cDf47041e8dEfc9F7c6bF0e080970C1a7;
    // clientkeys[117] = 0xd7Bf346952a3f4604d65F87d29c84070818dB062;
    // clientkeys[118] = 0x0958A3608809915B7822688717858C6461fa5fA0;
    // clientkeys[119] = 0xeEef891Ac2C344C7A1782393219babb8Ddb30851;
    // clientkeys[120] = 0x508B9FA11d30d569f7F00536FDc4Ce2936e2A154;
    // clientkeys[121] = 0x35013EE537eC7fE5A545763Ff51e7A83994017dF;
    // clientkeys[122] = 0xEcF96cA2cD8A5AE66613fa723e7dA358645644f1;
    // clientkeys[123] = 0xBCA0329cb2Afd2C111CEaC30c23FB0515f089cB8;
    // clientkeys[124] = 0xc16E12E3036F25c9fa6701E83A6D47aBB333C238;
    // clientkeys[125] = 0x68644658773612C028d6b2c0E4fF6708AD3780d0;
    // clientkeys[126] = 0x2fB982465D572AdDfeAA5FcF299A8F7eE3924FFD;
    // clientkeys[127] = 0x5901F544f19d7928efE96B5cB42F44973ee7829f;
    // clientkeys[128] = 0x355b346D34C8D3FA799e7e93499f167D2bada7aB;
    // clientkeys[129] = 0x61D165BE11E515642A66614d87968F49A6b1bED0;
    // clientkeys[130] = 0x7423Fb2a94d80B15CC0FED3524Da9C303A5A15B4;
    // clientkeys[131] = 0x9a5c79d07D872f8AA6206102F36073DFc7a6343c;
    // clientkeys[132] = 0xbB411EBC9e5c26069e58F7b3dB0ab12B2aeC5EC1;
    // clientkeys[133] = 0x133158BcD8732907E2b761013Ba2D5876a181528;
    // clientkeys[134] = 0xf3E7791ea9915cC3C26b4bE200845bB932928861;
    // clientkeys[135] = 0x76BE8343806570a89b793a88F396Ce5de1e9f3C5;
    // clientkeys[136] = 0x554E6849e5e934D30f013dE979DaF2EA7179fdA3;
    // clientkeys[137] = 0x26eF6B835a4cfDD0067C3733b24998Ec2f1139a5;
    // clientkeys[138] = 0xfBE72623dA26567908eE4A35A65896CaBb4C4884;
    // clientkeys[139] = 0x03E8b6F45486751B73C00a34eA924c6848547912;
    // clientkeys[140] = 0x91c6072e8cfB014F3B5B48dB1fB14A4e8d130491;
    // clientkeys[141] = 0x6B5A5173872865F277C8cFdD890b59f5DB855799;
    // clientkeys[142] = 0x8f4182F2602975038E62CaD68414C0bA99237189;
    // clientkeys[143] = 0xe0B37724937EF062BFFB71Eb63f5681bD9A322F9;
    // clientkeys[144] = 0x511976736A63a3aBf75D7E81cF23dF8828FE39a9;
    // clientkeys[145] = 0xB3dfb57a6e7C52D15Ff241e89265847cd45ad688;
    // clientkeys[146] = 0x1E9054A5E8FD163D8098635F4044462C28c0396A;
    // clientkeys[147] = 0x6e038f00568031cB7A5957BBa823e09b9cA5B1A0;
    // clientkeys[148] = 0x65Da27430298F14f6E02E1Ef73B0E4FFCd386514;
    // clientkeys[149] = 0xeEddd28660a23A9a98edF272cB3A4e075E1cAbF0;
    // clientkeys[150] = 0xEe606be1edBA57AD1277fecB219FB01700CadE21;
    // clientkeys[151] = 0xe9a39Ddb55327C566729b3464E61cF994063786e;
    // clientkeys[152] = 0x903CCc117C7792244391cFAD9bBCC61FB02FDCc1;
    // clientkeys[153] = 0xBfB668Ef113b87F3A741e30238261E081FB29786;
    // clientkeys[154] = 0x0bcDAFc342c41BfA4566444F637dC34714Ba3E91;
    // clientkeys[155] = 0x34732a57Ac7a69fFaaeA3f0Ec5ADa92fB6314b18;
    // clientkeys[156] = 0xAE09805321264ddbf884eD6617edC8447a2CE545;
    // clientkeys[157] = 0xd309ed7D304C12A76dC260b0D500B6601cC52abB;
    // clientkeys[158] = 0xa92b43764044cECC5d301E7F6AdC9cCEF11A330b;
    // clientkeys[159] = 0x3340c7E66D3dB445Aea319f9643C1CF70c6A9220;
    // clientkeys[160] = 0x606ca1e6aa6Fa44Cc84f4Bf87282F6f5292dED12;
    // clientkeys[161] = 0xcc63d1523184a8cd0C1d686C68e49438f493e55d;
    // clientkeys[162] = 0x1eaE82a55d713BC62e489e2B7c4E341DfB9843b8;
    // clientkeys[163] = 0x5A0722eBA955B2DBA75F33B9ccfa9753c297EA43;
    // clientkeys[164] = 0x88cA5e026BC4e563dc2EfA79a2aC4d899914DcB6;
    // clientkeys[165] = 0xc902aE32F0f68C6029FAD49788e4bE5f85115c60;
    // clientkeys[166] = 0xE0d475493FabAe5EBf282c30c32f5523c52b90C6;
    // clientkeys[167] = 0xF0083f7ADbE43814DcA7D151E89FF2344B43EEfb;
    // clientkeys[168] = 0x3D9e3f9C04b3C5648cb3f53756a8A076793aDBc8;
    // clientkeys[169] = 0x1749c2188240D6B441B85Ab75e7969Cd11b070AF;
    // clientkeys[170] = 0x9F0949b57B93844080A401790a4375202d63dE46;
    // clientkeys[171] = 0x9D62c5057859EC336BC26472a40f3bb1c584308b;
    // clientkeys[172] = 0x6FeDaA04e3aDcbe521eefd3dbbdc3660974cA2E6;
    // clientkeys[173] = 0x42A1deB80A13B69D7660d3ac1AcCD4688548a987;
    // clientkeys[174] = 0x7366eD8F6885ec96Ff0150f78e73ABA9ad1100Be;
    // clientkeys[175] = 0x00441DB699133E79684A8A159130EA41E3f3ccB2;
    // clientkeys[176] = 0x1699427D4042042AB46525560563F21bc89B0e42;
    // clientkeys[177] = 0x04A33F5d43f4e1aaa27899f0b8Ec116417Bf99a4;
    // clientkeys[178] = 0x391A45BE93711D5C690556f33f826474246A2e90;
    // clientkeys[179] = 0x3CD37Be64C2B4254B815b5CCa971529F765AA5bd;
    // clientkeys[180] = 0x171Eb6E0C17269534f7493c023804661A6Ae0Bd7;
    // clientkeys[181] = 0x7287cD8ca60BB44b72465b722a3B139f97870958;
    // clientkeys[182] = 0x8bFf1f45A52F4438071466C1D462a98CfbF30c98;
    // clientkeys[183] = 0x1133E780e5DbB8b05d7EE2Ee116D4B2e21dE350E;
    // clientkeys[184] = 0x6806815c5D2Ae49e043ce182Cd0b4d5286cD7686;
    // clientkeys[185] = 0x28BC49Aa7630880034E17099A8D821C77bfc98ca;
    // clientkeys[186] = 0x12bB975E1E52d5d719E536Bfc78eeC6fB451D9CD;
    // clientkeys[187] = 0xc7dde47F84De437cF913227387c4F2a7aE3B3799;
    // clientkeys[188] = 0x2d3Db4F73e2833af88B36519E1ea5c3348eC3F77;
    // clientkeys[189] = 0x759629262c87B4c5226F3C7078D69A14c54A32E2;
    // clientkeys[190] = 0x3891338217fFe35A20756784e0592C851a31Fdf7;
    // clientkeys[191] = 0x4BdaBd5651aDc37Cc98E7B92A52A591E135a3b7a;
    // clientkeys[192] = 0xd8B79D3874D9fC46773e91EC87E18d2e78071d00;
    // clientkeys[193] = 0x8afCC51530b8466735C78855c6A4bca1A587848c;
    // clientkeys[194] = 0x5d2A3836Fd4D6144b39d7d0c60689bf835002b09;
    // clientkeys[195] = 0x5ef0d1d9B07b61C60505C47C1a561E6BA117B6c9;
    // clientkeys[196] = 0x408CfaC9EC59BF99730009e7226AA3eb657709A8;
    // clientkeys[197] = 0x056A05a9B6D94f1D26aA16660c60320Ed41a5f30;
    // clientkeys[198] = 0xEc985d646df4d40CD56635678a94F733AB03E8C8;
    // clientkeys[199] = 0x9e3f6a6d69B158Eb72a9b8F03f7d770841f681a6;
    // clientkeys[200] = 0xc428e924B95d6DCa31EA01C75F862c6B63551fD6;
    // clientkeys[201] = 0x94397d89134Ead0aE19cFB7ab482dced9346d30d;
    // clientkeys[202] = 0x1f2c2E8B2Fc3e440F7d60AC4750cC020cc4c0f3a;
    // clientkeys[203] = 0x62a93a03FA4F3f1731244Eb09ec22353cb4C4902;
    // clientkeys[204] = 0x33e6f34594722F4BF6f554A8a5d562a34Ae92fc0;
    // clientkeys[205] = 0x082A3A283E971A7Ab99B91807CE616Db756f044d;
    // clientkeys[206] = 0xf45fd3D7A08EABEc453161F3F575348eB7e79Df9;
    // clientkeys[207] = 0x007d58650342a185a3DC3970373aF4b8a4F94Fdc;
    // clientkeys[208] = 0x064938430d857414bBD8B917845a65748BD4A582;
    // clientkeys[209] = 0xBFc979339f3f27f5ec2a9Da2C227A610F90cA78f;
    // clientkeys[210] = 0x0B5e71695225806fDF2A8855C51737137992F8E6;
    // clientkeys[211] = 0x41799C847caD4ABcA98443fb1521dc07da56c268;
    // clientkeys[212] = 0x82662Db352844254F67c77dd7f9a507371a0e607;
    // clientkeys[213] = 0xB56995724B5FA2DAaaeC36A3aa9c3949410aDd6B;
    // clientkeys[214] = 0x90e17644135BB1C6391c4Fd56618D518f2fBE879;
    // clientkeys[215] = 0x0f96DD25eFa41D42CD75a075c8967daDED215aE8;
    // clientkeys[216] = 0xa3912281Cb263a7Be42114Aaf0772411107B2f5f;
    // clientkeys[217] = 0x47942DB04FCDD56371694398650d97618Ff9B097;
    // clientkeys[218] = 0x2825F90Fe2DEBC13A96cE15E728b50a5AE060EE5;
    // clientkeys[219] = 0x5bc858E28F320591761b2d83157E0b001D7B0AA1;
    // clientkeys[220] = 0x185B2Bc8471155d2Ee7553bA92D64C0CaDDcFCbB;
    // clientkeys[221] = 0xA66738b53434B408e4e43DE2B4552AF26B895ae2;
    // clientkeys[222] = 0x743975ED92EC56Cab8Dd07fE0cE77ddE9CAcdACc;
    // clientkeys[223] = 0x7bc1C36c1F29A78Df00aA8b1731E9bA132C9a833;
    // clientkeys[224] = 0x513653932600Ed72e03BC595B51F72ce363A8448;
    // clientkeys[225] = 0x0d02A5250E76B2c09a7b68cef85Ec278E15D5861;
    // clientkeys[226] = 0xF27FBd43FB7d4741f2A9c260751268b7c99dBBa0;
    // clientkeys[227] = 0x4A5f6F9d5fB5160e3264aD5b3cb9B582c1A9A849;
    // clientkeys[228] = 0x0e25FacE826725120A37798504887AF259e9CD4B;
    // clientkeys[229] = 0x2fbC710Ea81512a54da9cB79cE6D621b37A68Cb5;
    // clientkeys[230] = 0x7a954B36a345421B9D781d37331d23A62eA0ACF7;
    // clientkeys[231] = 0x268667d0D9A2F9e649FB6eE58EaDf377BE9e220a;
    // clientkeys[232] = 0x15Bb7DeC53C991cFb93d4E6fA80C7F684568f013;
    // clientkeys[233] = 0x1A5C495981CE962Ff4cE252Efaaf6Abd4a6c58B9;
    // clientkeys[234] = 0x4C546eAf8030eEEFAd745C80C52257f844234a7B;
    // clientkeys[235] = 0xce59486EE89c685531369c2d9681D304BCCD9c52;
    // clientkeys[236] = 0x70C6b8c295FADCAAc2d68032C799b41a928FBa7d;
    // clientkeys[237] = 0x913A4D3FBFc5E0F7dc9F43a7f1df009E973440a8;
    // clientkeys[238] = 0xc4602bfA56f2265aF07F5381Bd7A20992fEF8339;
    // clientkeys[239] = 0x72c5AE1Ef5CB1E918AB2b7F6d5eAF6CdDC6438bA;
    // clientkeys[240] = 0xd49B63485448D9E9eb17dc8AaBD943F2C7f54391;
    // clientkeys[241] = 0x19F1D30462023893b2026a26623d2fB775B7b0F5;
    // clientkeys[242] = 0xCf60F3CA7cF63f86c763f2863b05AFC936Be4Dd7;
    // clientkeys[243] = 0x110547535E5468D4B6668ba65246397aE1807b46;
    // clientkeys[244] = 0xC01F21188DcA6c9e37486dD483066B1c99c85188;
    // clientkeys[245] = 0xbF32f392e9B96Cdd5742A5447387b867F2b9e72D;
    // clientkeys[246] = 0x8Bf25e55Fe97fBb3a2ddbd8166Dd5Eae0F6B564f;
    // clientkeys[247] = 0x8a87B443c3972A38B9d5504896393b988EC48787;
    // clientkeys[248] = 0xE3455E7B79B67802e74801A98E35255087327a9d;
    // clientkeys[249] = 0xB8Be310b9d45a041FE3B2ce7F202BB2747386b33;
    // clientkeys[250] = 0xE9cb924B8Adc1769f44283fe8f279021E857c410;
    // clientkeys[251] = 0x1c28082F4Cd5F8aC4F26E66b2b640622e960052D;
    // clientkeys[252] = 0xA5D70a197fc39766D70190b5e36BB3adA3B3bB8D;
    // clientkeys[253] = 0x9c4054DdbA97710B10965d4b4D1e41AcF04b6285;
    // clientkeys[254] = 0xC4E7a9Ac87528a39B975ebae5e2458be682a3D6E;
    // clientkeys[255] = 0xB53547D0143453710802880c59B8CCf0cE018e01;
    // clientkeys[256] = 0x46a26750B464669EFf8D71a9E13ebc9E393FAFC1;
    // clientkeys[257] = 0x917176C0660140DD42E8ea53970955403c37d583;
    // clientkeys[258] = 0x2fF65d00e92Ae1986d966baA49f14D33242c88A5;
    // clientkeys[259] = 0x8064C86149F7f644a85Aa212E8c3bC20D5f94110;
    // clientkeys[260] = 0x0182D0842C9cd77E9bB346215EaD3A0756dDac2a;
    // clientkeys[261] = 0x28E924BB434c4457Abc1C685989e102d1B690ED1;
    // clientkeys[262] = 0xEe98E362D678732Bd0e73853fbd36838F97d6C15;
    // clientkeys[263] = 0x2b80E50E5bC8DC4D4b02A7F6Aa5de24001e54ee4;
    // clientkeys[264] = 0x6b4BD1BDfd1A36F736C776d1B5F9C95e2f00291C;
    // clientkeys[265] = 0x30Cb2f4951Febe189C84bb5091fBC2c898dB3908;
    // clientkeys[266] = 0x04e455D28A4382e1E8dE4C15397B7190023a7262;
    // clientkeys[267] = 0x71a86C1954A327B8Edbf206Ea5C4BE11847E73a1;
    // clientkeys[268] = 0x0FB575db244a6052a21Ba3b8AD8580d691dE9834;
    // clientkeys[269] = 0x4409A9b57734b17B6a3e5BB10D4906123deB7f90;
    // clientkeys[270] = 0x24caf3bB3Ecbe3D6bfEC659a3B5d1990B0d4d5FB;
    // clientkeys[271] = 0xe9A6E26FBAA13856A9080508f3aD60eB8A8FdE18;
    // clientkeys[272] = 0x649E10Cb7504Aa1E3a5aF4D07a5D65ea527977E1;
    // clientkeys[273] = 0x67cd8C88CE78d66Fb223D128d5253F40e4B5D868;
    // clientkeys[274] = 0x8f0f4A1E68c7D50Ba842b2dCC3Af2dB99ebC6467;
    // clientkeys[275] = 0x9a148aBa0beCb6B53Cfae6d1c09738fFb4F02968;
    // clientkeys[276] = 0x24fE06c07f88a2f2946b82cCCd684Bf3c70214Fe;
    // clientkeys[277] = 0x98605BF860DB030c4f2A1178021F40AA3779A295;
    // clientkeys[278] = 0x1b643B0CeCDf4dA5D684F47359c3827320d72AB6;
    // clientkeys[279] = 0x30bCa1A51E10e3760f3eFAfbb7663e6FCA6913CE;
    // clientkeys[280] = 0xf00eaF4724dfaA0549D8e1C4Ed25F22F263f28AF;
    // clientkeys[281] = 0xf1E227fFadB684C8E7cD262337469B2419Af1dee;
    // clientkeys[282] = 0xf06bEA5c0F862e90EA9DfD1ca602c416496F1904;
    // clientkeys[283] = 0xB6234bc45c3ed595bE21c9DdC9539666deF711AA;
    // clientkeys[284] = 0x7f1CF1DCf5f2BeC35EF6D3cd709341137bf5B898;
    // clientkeys[285] = 0xFD86d332c3029021d32d7FAE058c3b5080C5878D;
    // clientkeys[286] = 0x8561B4a9ab872F13Ec193d47815310a979D65406;
    // clientkeys[287] = 0x13513266aD3c2003184b4a8c361b0C699cDB38a5;
    // clientkeys[288] = 0x495F24338dfdAA6a2737AB208C88f03a4EA8477F;
    // clientkeys[289] = 0xEeBDC0ED0b05B2F3C7388C39B030232b8ddE9b31;
    // clientkeys[290] = 0x938D4A4d4a35BCDb2dB564eDbaFF42ab0B46Ef30;
    // clientkeys[291] = 0x0D460Da5421f111Ce92FE589d0A940DBCa6B5367;
    // clientkeys[292] = 0x706C065eF7eD278880733704cf478f7978a69594;
    // clientkeys[293] = 0x8a6d42E9b13De9fBF2966B13A05E5061DC70FD40;
    // clientkeys[294] = 0x04042847D07b10077DB7f98938349B0ae67bBE6d;
    // clientkeys[295] = 0xff67031206B5d6898Fe34699BB1C636524cAE596;
    // clientkeys[296] = 0x02107E6633aD135a3a5DB2e78D5d9fd91510E4EE;
    // clientkeys[297] = 0x024D145004a6C1b8c7Bb185b94f4027C86E850Ef;
    // clientkeys[298] = 0x750a283D5983a5DE86A0F16ED852427c9e79301b;
    // clientkeys[299] = 0x8f512440E548d9A397E09A2ea29Af749896F5a3D;
    // clientkeys[300] = 0x805c37e69652D8710E1b1fC0d353FadC3cA2c588;
    // clientkeys[301] = 0xeF38e8EE8F514900473dD6d8d3862FE4264f9DaB;
    // clientkeys[302] = 0xaC7e7b3B4F0E62f4F7B3159ea093853691C5f3B4;
    // clientkeys[303] = 0x1bb1A2e0278c1B68d2B43f10DB34fAE864D93210;
    // clientkeys[304] = 0x2149C5eba168D08516291e3c52b5b6c0ed216d3f;
    // clientkeys[305] = 0xE407C4a62E6a7E8f44392D7Cc299EDb1B2f6D72E;
    // clientkeys[306] = 0x0e229087C19d3832E94539b30CAE0CB59Ea2EB13;
    // clientkeys[307] = 0xc3ef9842F8328391E88952D613A1f65b3CB07D3b;
    // clientkeys[308] = 0xE6e8694240C4c1f286eBef081343B38a14dbDefd;
    // clientkeys[309] = 0x5E1ED1afe1687a4216ccB4b4ea39d21cEe16FE37;
    // clientkeys[310] = 0x9F41F7Ae9fA8602142fc9b7C27C15487Eb4EDc41;
    // clientkeys[311] = 0xdcf15CAc114a60df371F8524Cbf23E61648aF49d;
    // clientkeys[312] = 0xe03Ac2E2C8fF5539A309fb6508542e322C30A52D;
    // clientkeys[313] = 0xD0082Afd94F996FFb1184D09eB6f09769342d8c1;
    // clientkeys[314] = 0x6105e4Ba978237EEf0464347FeAa63f1E286937a;
    // clientkeys[315] = 0xe9b8eb0a4057737Ff16d1706F4368C254cbCb9D8;
    // clientkeys[316] = 0x3f867f8f5b10C98Dcef049f7dF13Ae891b570E73;
    // clientkeys[317] = 0x90313726ff6fcdf9A912cDc1FDC417D8f88aFbFa;
    // clientkeys[318] = 0x867BA96E2e5065673d756Da61D36aA905361b50B;
    // clientkeys[319] = 0x7D3d4D18bcFC17a4575Be149A77e975700Ea9F13;
    // clientkeys[320] = 0x9398cd22eF3802A0924eDAd3cfa75A422bE75318;
    // clientkeys[321] = 0x68637B0e1B1b430CCD30381b582D061251536F49;
    // clientkeys[322] = 0xA6356FFf60D915a869ae47fea31a186eb1f3862f;
    // clientkeys[323] = 0xBAe04eDf1d6450B36518c25C388766ec2aeDcea3;
    // clientkeys[324] = 0x53AbDc885a4059fc26c3c631fa5e54e08abcBa1F;
    // clientkeys[325] = 0x73e0703A1F3D06A02C02df4c1228316523a297FC;
    // clientkeys[326] = 0x5D52165943d192Ef41E371502F99f51af90d59dE;
    // clientkeys[327] = 0x87862B48c3Fd8D1514230DF23B1Df1fb034A7532;
    // clientkeys[328] = 0x61bf05D5D06b8600F4a16F042fe9dDe6f53654ae;
    // clientkeys[329] = 0xc69C412D6f57611d9BFaf5E4E416C07E7650166d;
    // clientkeys[330] = 0x151c52f2f43Fe4B53349fB4fB0B8382555b5F0c5;
    // clientkeys[331] = 0x9C978F2de5C5F54f0D3172af2b180dc806BF71e1;
    // clientkeys[332] = 0x34Afc62a17671124fA9dED563CFea7Ce2AC85dD3;
    // clientkeys[333] = 0x9DC5F6Ac266Cbd0800b897C8d30c7f1244a19583;
    // clientkeys[334] = 0xcfeC1615dEc5989D95b15F133Da1Ce650E5CFB12;
    // clientkeys[335] = 0x9b61F7C92a6E83027D3Bd807F528c1C692966745;
    // clientkeys[336] = 0x63D299d71ee5aA72D3D0621ae2Fb14d65069700f;
    // clientkeys[337] = 0xdcAAf629Cf9F18240A087A4298cD86cA0Dd432D1;
    // clientkeys[338] = 0x2c2c0ebfA282590Aa114359b3057d1965b6ea623;
    // clientkeys[339] = 0x780302C498eE2110Da251a51c0BcAa5eE15AC78C;
    // clientkeys[340] = 0xb5c1E7A8866556310c9f39dbD14fFa8A59457F0c;
    // clientkeys[341] = 0x1f486308c634aa56D1b7194480cD3e3F2BA1BB9d;
    // clientkeys[342] = 0xFb551C496472d8CD3BB4b0a369e0e1a0c3b5eA94;
    // clientkeys[343] = 0x858792847582008Ef8c4C9ad9854E1332D95852C;
    // clientkeys[344] = 0xB3C0305fCbc7A4f893EaF3f3a558537459a4089E;
    // clientkeys[345] = 0x21757eBe75223A64106CDe7CF9Daea06386291Ed;
    // clientkeys[346] = 0xe30b6c65964b1C71E135d616dE09cDa1bAa54036;
    // clientkeys[347] = 0x35Ec3AC6556CA202dd7FBfB16e8d2a697F12A836;
    // clientkeys[348] = 0x00a0ecC418ad1E98E747571F228495a334f7CBf3;
    // clientkeys[349] = 0x7394C2f34AEd3B0aC553e9108D2eBB7A7b8Ee591;
    // clientkeys[350] = 0x032C6D9d94e28ee25BBe1d58FA5E65Afb9de0f3c;
    // clientkeys[351] = 0xb8284552410f57fBBFCD11E5d394eCaBd57cfA26;
    // clientkeys[352] = 0xC60E0A184f96115f5Ef21521c69138BfD6f58887;
    // clientkeys[353] = 0xcC58a1a9491EB41FBB0f3cCA8E0A03490e98652a;
    // clientkeys[354] = 0x36cBAd88C605E0D89a3BCB7a1A6a3eDAe726A110;
    // clientkeys[355] = 0xb022DF08530264E16e616416F7a9C83b5b602F14;
    // clientkeys[356] = 0x258bb8213d25e62aaC0d944632297979C489e605;
    // clientkeys[357] = 0x910a1F54C73DbCaA3b91a3cAbaf163d5A3575F68;
    // clientkeys[358] = 0xB630F6cEE89020e70B8EE0426217c9647c3f917e;
    // clientkeys[359] = 0xEcf682A9BA130CFBD17aD550fE42c48f6Bea665e;
    // clientkeys[360] = 0xe1a4391153CaA0C8671Fbe888C160Aa62A2254Af;
    // clientkeys[361] = 0xb31c8b8FEe0924dfC5917D01248AAFA6BD25E9DD;
    // clientkeys[362] = 0xE7A3D8f152dd78Ac478D355B3BD8158475c60c68;
    // clientkeys[363] = 0xbDb5971CA03Ba03ad9Adaf13FBbdb9B58F9b7f01;
    // clientkeys[364] = 0x081745d364DaBb0F5282d79eF7Ae924dd675C861;
    // clientkeys[365] = 0x41Da51e170d10e0300C55F701Ee2a89f5eE17eb4;
    // clientkeys[366] = 0x0B7CBD10Fd85bD5184fa52A7872315af7e6898F3;
    // clientkeys[367] = 0x5353782CA0c5E6A5c50BaF95AF899Cf6200BB14f;
    // clientkeys[368] = 0x2DB856f7b6DCb61297aFAeACc17345286A4AD1c0;
    // clientkeys[369] = 0x26654282A705b6F62B3B2C89A4d1faBf800b9c4B;
    // clientkeys[370] = 0xAF49f53e652b83a92081B884093f3Fb68fe79127;
    // clientkeys[371] = 0x0ce6423dAf24A3E6E513E59CdD9F709B4198383F;
    // clientkeys[372] = 0x7d31AD761b24E2Ead7449e549cB8B70b92298702;
    // clientkeys[373] = 0xD6c9F76ba27044d8c61178042ff711E9F3911D12;
    // clientkeys[374] = 0x9f7Ee13cf3DbF30a369D56575554ac80f94586C3;
    // clientkeys[375] = 0xD14AD853B04F5Eedf7476E2647d07f11b6746061;
    // clientkeys[376] = 0x70E60A36D6051A2EEf6171543781a26ac82D7728;
    // clientkeys[377] = 0xb822270E8A796CBF5E272c546fa1D21D7F8b5BCE;
    // clientkeys[378] = 0xb14547A12c6A8CaF2BCF87a54746b7f578f43eE7;
    // clientkeys[379] = 0xA3FC9EeDA6748179Cca635d0625Dd723743ec525;
    // clientkeys[380] = 0x44C0Daf177F785c7d9cDDd8837D34F41dd3Ab3F9;
    // clientkeys[381] = 0x5efd15B76DD0D0208c7D767Dd38BCC0Fab5a4304;
    // clientkeys[382] = 0x582066181527F17f58Cf7786BE395E717Daa278E;
    // clientkeys[383] = 0xC22F0cE160c14193889e73eBAb6b9FBd722dD545;
    // clientkeys[384] = 0x5Be9f3B30D0842B33c34e42b56cBF89e08665891;
    // clientkeys[385] = 0x754B1198bb0A62a5D62629f140458E72DA1dB74c;
    // clientkeys[386] = 0x89c66b2BA3b886500223A89055d584A8FAE88305;
    // clientkeys[387] = 0x17EB0Ef293F363BC958bDE6f950490Aa40b0d407;
    // clientkeys[388] = 0x3212057Fe42D60935deDdB870fD337E5A2FAeE59;
    // clientkeys[389] = 0x914B4145a2B6712F42cb7A6d19f43134A3bcD1c2;
    // clientkeys[390] = 0x46B87192AE6b1D413aC92C8C1375BCba391dBB82;
    // clientkeys[391] = 0x25961F06929dd83D676D548Cf181F3e39bBFaE51;
    // clientkeys[392] = 0x44d92c0D2beC5333C039eF5a4192F49a10691Fa5;
    // clientkeys[393] = 0xF77998a5a704fe8D7cC54326eD7ab201bE083110;
    // clientkeys[394] = 0xDc08B6072952bDE6E54Ceb8cbCa91ee8B8C806E5;
    // clientkeys[395] = 0x0b14F624BECd527f743A43cD33c9Dd8a94635423;
    // clientkeys[396] = 0x4E94A7e8867687c69075e8EF4c167aAEF75293eb;
    // clientkeys[397] = 0xCc421859B98f73fB9AC76c777a5676D3D3BF6ac6;
    // clientkeys[398] = 0xdF9d702021869811023639611CaB7d325069fD50;
    // clientkeys[399] = 0x5A91CAD00BBFb7224cE9BCDf7d83a8742E26c681;

    for (uint i = 0; i < clientkeys.length; i++) {
      admins[clientkeys[i]] = true;
      devices[clientkeys[i]] = true;
      mf[clientkeys[i]] = true;
    }
  }

  modifier onlyClient () {
    require(devices[msg.sender], "Client is not registered");
    _;
  }

  modifier onlyAdmin () {
    require(admins[msg.sender], "Sender must be admin");
    _;
  }

  function timestampHash(bytes32 hash) public onlyClient {

    timestamps[hash] = Timestamp(
      int64(block.timestamp),
      hash,
      msg.sender
    );

  }

  function updateFirmware(bytes32 hash, bytes memory mfSig) public onlyAdmin {

    address mfAddr = recoverSig(hash, mfSig);

    require(mf[mfAddr], "Manufacturer is not registered");

    firmwares[hash] = Firmware(
      int64(block.timestamp),
      hash,
      msg.sender,
      mfAddr,
      mfSig
    );

  }

  function recoverSig(
    bytes32 hash, bytes memory sig
  ) internal pure returns (address) {

    if (sig.length != 65) {
      return address(0);
    }

    bytes32 r;
    bytes32 s;
    uint8 v;

    // solium-disable-next-line security/no-inline-assembly
    assembly {
      r := mload(add(sig, 32))
      s := mload(add(sig, 64))
      v := mload(add(sig, 65))
    }
    v += 27;
    return ecrecover(hash, v, r, s);
  }

  function proposeAddKey(address key, uint8 keyType) public onlyAdmin {

    require(checkKey(key, keyType) == false, "Key already existed");
    addKeyProposals[key] = newProposal(key, keyType);
    voteAddKey(key);

  }

  function proposeRemoveKey(address key, uint8 keyType) public onlyAdmin {
    
    require(checkKey(key, keyType) == true, "Key does not exist");
    removeKeyProposals[key] = newProposal(key, keyType);
    voteRemoveKey(key);

  }

  function checkKey(address key, uint8 keyType) private view returns(bool) {
    require(keyType >= 1 && keyType <= 3, "Invalid Key Type");

    if (keyType == keyTypeAdmin) {
      return admins[key];
    }
    if (keyType == keyTypeDevice) {
      return devices[key];
    }
    if (keyType == keyTypeMf) {
      return mf[key];
    }
  }

  function newProposal(
    address key, uint8 keyType
  ) private view returns(Proposal memory) {

    Proposal memory p;
    p.unixTime = int64(block.timestamp);
    p.keyType = keyType;
    p.key = key;
    p.proposer = msg.sender;

    return p;
  }

  function voteAddKey(address key) public onlyAdmin {
    require(
      isExpired(addKeyProposals[key]) == false, 
      "Proposal does not exist or expired"
    );

    require(
      addKeyProposals[key].voted[msg.sender] == false,
      "Already voted with the same key"
    );
    addKeyProposals[key].voted[msg.sender] = true;
    addKeyProposals[key].voteCount++;

    if (isVotedEnough(addKeyProposals[key])) {
      setKey(key, addKeyProposals[key].keyType, true);
    }
  }

  function voteRemoveKey(address key) public onlyAdmin {
    require(
      isExpired(removeKeyProposals[key]) == false, 
      "Proposal does not exist or expired"
    );

    require(
      removeKeyProposals[key].voted[msg.sender] == false,
      "Already voted with the same key"
    );
    removeKeyProposals[key].voted[msg.sender] = true;
    removeKeyProposals[key].voteCount++;

    if (isVotedEnough(addKeyProposals[key])) {
      setKey(key, removeKeyProposals[key].keyType, false);
    }
  }
  
  function isVotedEnough(Proposal memory p) private pure returns(bool) {
    return p.voteCount >= 2;
  }

  function isExpired(Proposal memory p) private view returns(bool) {
    return int64(block.timestamp) > p.unixTime + proposalDuration;
  }

  function setKey(
    address key, uint8 keyType, bool value
  ) private {

    if (keyType == keyTypeAdmin) {
      admins[key] = value;
    } else if (keyType == keyTypeDevice) {
      devices[key] = value;
    } else if (keyType == keyTypeMf) {
      mf[key] = value;
    }

  }

}
