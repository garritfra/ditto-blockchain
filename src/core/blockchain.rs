use core::block;
use core::transaction::Transaction;
use core::tx_pool;

pub struct Blockchain {
  pub blocks: Vec<block::Block>,
  pub tx_pool: tx_pool::Tx_pool,
}
pub fn new() -> Blockchain {
  println!("New blockchain\n");

  let genesis = block::create_genesis();

  let blockchain = Blockchain {
    blocks: vec![genesis],
    tx_pool: tx_pool::new(),
  };
  return blockchain;
}

impl Blockchain {
  pub fn mine_block(&mut self) -> &block::Block {
    let mut block = block::new();
    block.transactions = self.tx_pool.transactions.clone();
    self.clear_transactions();
    self.blocks.push(block);

    &self.blocks.last().unwrap()
  }

  fn clear_transactions(&mut self) {
    self.tx_pool.sanatize()
  }

  pub fn add_pending_transaction(&mut self, transaction: Transaction) {
    self.tx_pool.add(transaction)
  }

  pub fn get_blocks(&mut self) -> &Vec<block::Block> {
    &self.blocks
  }
}
