export default {
  tx: {
    txid: '0xa1e6d3cd2e4c5bc06af21835065a44eb2d207962ebf36b9e24a366eb20e906da',
    vin: [
      {
        n: 0,
        addresses: ['0xBdd8CdB1158ba84DE117C8670BB27b80376Def1B'],
        isAddress: true,
      },
    ],
    vout: [
      {
        value: '6412730000000000',
        n: 0,
        addresses: ['0xC145990E84155416144C532E31f89B840Ca8c2cE'],
        isAddress: true,
      },
    ],
    blockHash: '0x41f15066c42d1d393ba30ff0729521631080711e978d3596ee140ec7bd412369',
    blockHeight: 12604164,
    confirmations: 183654,
    blockTime: 1623293292,
    value: '6412730000000000',
    fees: '1555280000000000',
    ethereumSpecific: {
      status: 1,
      nonce: 205,
      gasLimit: 90000,
      gasUsed: 38882,
      gasPrice: '40000000000',
      data:
        '0x574da717000000000000000000000000fc0cc6e85dff3d75e3985e0cb83b090cfd498dd100000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000016c8580db0c40000000000000000000000000000000000000000000000000000000000000000800000000000000000000000000000000000000000000000000000000000000047524546554e443a3835314234393937434638463946424138303642333738304530433137384343423137334145373845334644353035364637333735423035394232324244334100000000000000000000000000000000000000000000000000',
    },
  },
  internalTxs: [
    {
      blockNumber: '12604164',
      timeStamp: '1623293292',
      hash: '0xa1e6d3cd2e4c5bc06af21835065a44eb2d207962ebf36b9e24a366eb20e906da',
      from: '0xC145990E84155416144C532E31f89B840Ca8c2cE', // normalized to checksum format
      to: '0xfc0Cc6E85dFf3D75e3985e0CB83B090cfD498dd1', // normalized to checksum format
      value: '6412730000000000',
      contractAddress: '',
      input: '',
      type: 'call',
      gas: '54729',
      gasUsed: '0',
      traceId: '0',
      isError: '0',
      errCode: '',
    },
  ],
}
