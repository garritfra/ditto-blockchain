use core::transaction::Transaction;

#[derive(Serialize, Deserialize)]
pub struct Block {
  pub hash: [u8; 32],
  pub prevHash: [u8; 32],
  pub transactions: Vec<Transaction>,
}

impl Block {
  fn hash(&mut self) {}
}

pub fn new(prevHash: [u8; 32]) -> Block {
  let block = Block {
    prevHash: prevHash,
    hash: std::slice::,
    transactions: Vec::new(),
  };

  let mut data = bincode::serialize(&block);

  return block;
}

pub fn create_genesis() -> Block {
  Block {
    transactions: Vec::new(),
  }
}
