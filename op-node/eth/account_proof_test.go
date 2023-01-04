package eth

import (
	"encoding/json"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/stretchr/testify/require"
)

func TestAccountResult_Verify(t *testing.T) {
	// Example account result: a SystemConfig storage read, to proof the value of slot 103 (0x67) equals some address,
	// which would be RLP-encoded when retrieved as value from a trie node.
	resultData := `
{
    "accountProof": [
      "0xf90211a03d0bbcace6414d254eb5ee34923da6dadda532025554251fccb0f3e401f97f64a00b0060669e7d8bee64b69cd327edd83aec3839bcdc2c51df61a87d26c06a6e6da0c0325f24cc335b26d107c0704a0c88d96aac42cb807f732d45196ba2b3ec6f4fa0be4d5280fb18316c250432acd5272e0558a850e2b17803e2262a0f12ec7293cda0584c838d31e2b8b7b1b61ea86ec96c14de6f8e48b5085ab88f26022e46c434f3a033ce2881dfbf590f36e88d999b52b5f88fa3e2845aa069f354fffc4657269db0a0bc0d89addbe9e9dd1b4570e6403b9540c631d4b018d6a30dee0d6b1416d79f7da0b208a7628b6e33e21638856ae0cf77f0245aab58300a3ea459b049fcbcf7c95aa01d16509886c5fef9deee332e82ca26739b35acfc192c4fd5f31310caf9996304a062a1114d555e0657a953a9e79156ecf6323629961afe756c846a1222c415999aa0e3ae9195242e45330d273187de63ccd74c2ab787dc50e0ff3edbc8ef47cd3b2da0c1b261a54efdfa3596cec84a23784ed051bc346b80dbf0694dc7d39356c212eea04512835bd07815d2bcc00f25ea70dee0f0a770b16d16250f7d5c884b0d440581a0ee25f14adcbd0dc55d15a2148deae12c3ab9c6a001c6e899e22d3e0890999083a01c0c47189ac3d6931aa05fbe3cd4140a1794f68b724b1804f46d2705435f08caa0d9a14e2b4e6778a2405bc7931934b5e97b8e3a2033eedb30409896420a2db44b80",
      "0xf90211a0e65e35c2066e63b933bd300ad97883dbcf4e2d2df8a540aa3964f5354e2bb8dba0814e82409c1813007af146f3b47f0e0d6da4c2885de8f34f034826bcb4c4c9e7a07d726afe2a20f4928e8715d80bc5aff5fc30c05b1e953372d564cba93d3b45bda02fc8df4a480ea0402f48aa206047ef0acbedf1a7033153c60d5ec4dfb9ee30cca05ecda76647515fb8342434b5a36faec2f17fc729ea369edcf7a42cd235c51fa2a0ab51933d2bde083801bd7fedb17ec22ae5231e508c62d66e9dccbd598c1d98fda0e2af2dc4a1b87147a5e388388ee5d1aafb415498e1011e87df127cc8d57616eca00bb9b95cfda6476e321ca441f523f80cb13519e537679ad6139600b35c91b5b1a0e3aae75bdc65365e2f8bb9c8646a4118d302076e0d9de95fc4425e82680ed912a0124bb32fc404a8647d7192f438a2bca8c52877b02a921b014fe4a192982bcf9ea0da2df7e5f41d6e38cc48fd9d0ca91bc96ff6ffa336b75784b3e23cd8dfddc0b1a0f7fc4a6792cdf964423f705935f0338128d099c83a74c2b76c688ecda8f23ffea009b2347ed0222d713137ad280310fad2ca92f007b5fd3ed1b768a6b52b5040d6a07b5bfe1bbd07575f26d97b9a0e8c67cd11d251f3d8bc823ade462daee2309e4ba02d2e464a4ad06d6dca6cb54e31fb5aac85dc9c5bab420a2c7c1ca986991f4d1da0cd5f723720652dec2a8bbba3bad3f73fb16b83a31d28a707fd0e0f85370eae6380",
      "0xf90151a04df2503da3b2491b68d13d2735005e451ca9cae77d336061df9250e0d3592e0ca01580374cce4fa8401d9fdeebb037e96c3f9c03873aec0637b43ba809fa8e53a9a0101d14d18ee6c9bf21f8292b155e40462605d15d315337e0f1098479d9ba6d1b80a0d0baf298792955a012fc458a97de8b3b2319911c3a82e427b5424fb544ac1ea080a09b2cc83f7e69b3c46bde63e4805d1aa9a064c0220532de547693122bba9b24458080a0a44850b8b25ced24d074182424ab7aa668b78acaa64e2da1854d04093332153b80a00b513b7944399ac7e2291ab7cdc1b99f0445256cc3f0c3b6b5c2d50eab4f9887a06b979f3c7a2a1d95af02c44d015ac42202cd0df8c19b629f6089a6a2a181d69aa0eb0216e118a1b797b90165cea439923e58bc672b81573ae44156cf38632fa36880a0e16e6a4603316f2fa41cd7d0a8fc9a8dca7ca36763abe39b99ef57b83adf501180",
      "0xf8689f3b8dff764734b5790096e811fc30835b13a5bea290831c7c6722f325a5f818b846f8448080a0dd6d0b4dba95e79439c501e2109ded26989a53ca4641e1f48bf702e7361637e9a01f958654ab06a152993e7a0ae7b6dbb0d4b19265cc9337b8789fe1353bd9dc35"
    ],
    "address": "0x6900000000000000000000000000000000000009",
    "balance": "0x0",
    "codeHash": "0x1f958654ab06a152993e7a0ae7b6dbb0d4b19265cc9337b8789fe1353bd9dc35",
    "nonce": "0x0",
    "storageHash": "0xdd6d0b4dba95e79439c501e2109ded26989a53ca4641e1f48bf702e7361637e9",
    "storageProof": [
      {
        "key": "0x0000000000000000000000000000000000000000000000000000000000000067",
        "value": "0x9965507d1a55bcc2695c58ba16fb37d819b0a4dc",
        "proof": [
          "0xf8f18080a04fc5f13ab2f9ba0c2da88b0151ab0e7cf4d85d08cca45ccd923c6ab76323eb2880a0f57febb7b16455e051f412a56e54016c676a3d4aa515d2e77a90520dfe36162ea0558f72e6d0e3b401856defa90b07dd0442282592b3ca718e2dc919ea53b8e69280a04f893abcf66ae78abb4ac986ddf78bdf95a7b15079be06cc5756f78d772271eba0c1529c7d0f249fd7060e930515ac4980103920979274f56b07419bf33be4d3d7a0a055722fdc9281d825dfc17c2ae775aba5b283954f5c484fc7e0e5a148131e2ea02833bc13e1f58010a678009d7d5982b892b3ba1432ca6b06f8a849b71491b51e808080808080",
          "0xf7a03787eeb91fe3101235e4a76063c7023ecb40f923f97916639c598592fa30d6ae95949965507d1a55bcc2695c58ba16fb37d819b0a4dc"
        ]
      }
    ]
  }
`
	var result AccountResult
	require.NoError(t, json.Unmarshal([]byte(resultData), &result))
	require.NoError(t, result.Verify(common.HexToHash("0xb3a98a923c23cf25cbe04485f55243b37b29b7e12760bd24368ace23bf370e7a")), "verifies against good state root")
	require.NotNil(t, result.Verify(common.HexToHash("0xb3a98a923c23cf25cbe04485f55243b37b29b7e12760bd24368ace23bf370e7b")), "does not verify against other state root")
}
