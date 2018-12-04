pub struct Transaction {
  pub from: String,
  pub to: String,
  pub amount: i32,
  pub message: String,
}

pub fn new_with_amount(from: String, to: String, amount: i32) -> Transaction {
  Transaction {
    from: from,
    to: to,
    amount: amount,
    message: "".to_string(),
  }
}
