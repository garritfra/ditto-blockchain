use core::transaction::Transaction;

pub struct Block {
  pub transactions: Vec<Transaction>,
}

pub fn new() -> Block {
  Block {
    transactions: Vec::new(),
  }
}

impl Block {
  pub fn add_transaction(&mut self, transaction: Transaction) {
    self.transactions.push(transaction);
  }
}

pub fn create_genesis() -> Block {
  Block {
    transactions: Vec::new(),
  }
}
