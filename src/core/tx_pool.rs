use core::transaction::Transaction;
pub struct Tx_pool {
  pub transactions: Vec<Transaction>,
}

pub fn new() -> Tx_pool {
  Tx_pool {
    transactions: Vec::new(),
  }
}

impl Tx_pool {
  pub fn add(&mut self, transaction: Transaction) {
    self.transactions.push(transaction);
  }

  pub fn sanatize(&mut self) {
    self.transactions.clear();
  }
}
