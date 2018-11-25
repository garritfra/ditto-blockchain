import Block from "./Block";

export default class Blockchain {
  blocks: any[];
  genesisBlock: Block;
  constructor() {
    this.blocks = [];
    this.genesisBlock = Block.createGenesisBlock();

    this.addBlock(this.genesisBlock);
  }

  addBlock(block: Block) {
    this.blocks.push(block);
  }

  getLastBlock() {
    let lastBlock = this.blocks[this.blocks.length - 1];
    return lastBlock;
  }
}
