use core::block::Block;

pub struct Blockchain {
  blocks: Vec<Block>,
}
pub fn new() -> Blockchain {
  println!("New blockchain\n");

  let genesis = Block { data: 32 };

  let blockchain = Blockchain {
    blocks: vec![genesis],
  };
  return blockchain;
}

impl Blockchain {
  pub fn add_block(&mut self, block: Block) {
    self.blocks.push(block);
  }

  pub fn get_blocks(&mut self) -> &Vec<Block> {
    &self.blocks
  }
}
