use core::transaction::Transaction;

pub struct Block {
  pub transactions: Vec<Transaction>,
}

pub fn new() -> Block {
  Block {
    transactions: Vec::new(),
  }
}

pub fn create_genesis() -> Block {
  Block {
    transactions: Vec::new(),
  }
}
