mod core;

use core::block;
use core::transaction;

fn main() {
    let mut blockchain = core::blockchain::new();

    let transaction = transaction::new_with_amount("Me".to_string(), "You".to_string(), 10);
    let mut block = block::new();
    block.add_transaction(transaction);
    blockchain.add_block(block);
    for block in blockchain.blocks {
        for transaction in block.transactions {
            println!("{}", transaction.from);
        }
    }
}
