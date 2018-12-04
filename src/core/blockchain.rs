use core::block;

pub struct Blockchain {
  pub blocks: Vec<block::Block>,
}
pub fn new() -> Blockchain {
  println!("New blockchain\n");

  let genesis = block::create_genesis();

  let blockchain = Blockchain {
    blocks: vec![genesis],
  };
  return blockchain;
}

impl Blockchain {
  pub fn add_block(&mut self, block: block::Block) {
    self.blocks.push(block);
  }
}
